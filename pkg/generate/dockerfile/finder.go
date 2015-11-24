package dockerfile

import (
	"os"
	"path/filepath"
	"strings"
)

var DefaultDockerfileName string = "Dockerfile"

type Tester interface {
	Has(dir string) (string, bool, error)
}

type tester struct {
	statFunc       StatFunc
	dockerfileName *string
}

type StatFunc func(path string) (os.FileInfo, error)

func (t tester) Has(dir string) (string, bool, error) {
	path := filepath.Join(dir, *t.dockerfileName)
	_, err := t.statFunc(path)
	if os.IsNotExist(err) {
		return "", false, nil
	}
	if err != nil {
		return "", false, err
	}
	return path, true, nil
}

func NewTester(dockerfileName *string) Tester {
	return tester{
		statFunc:       StatFunc(os.Stat),
		dockerfileName: dockerfileName,
	}
}

// Finder allows searching for Dockerfiles in a given directory
// with a given name
type Finder interface {
	Find(dir string, dockerfileName string) ([]string, error)
}

type finder struct {
	fsWalk func(dir string, fn filepath.WalkFunc) error
}

// NewFinder creates a new Dockerfile Finder
func NewFinder() Finder {
	return &finder{fsWalk: filepath.Walk}
}

// Find returns the relative paths of Dockerfiles in the given directory dir.
func (f *finder) Find(dir string, dockerfileName string) ([]string, error) {
	dockerfiles := []string{}
	err := f.fsWalk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Skip hidden directories.
		if info.IsDir() && strings.HasPrefix(info.Name(), ".") {
			return filepath.SkipDir
		}
		// Add relative path to Dockerfile.
		if isDockerfile(info, dockerfileName) {
			relpath, err := filepath.Rel(dir, path)
			if err != nil {
				return err
			}
			dockerfiles = append(dockerfiles, relpath)
		}
		return nil
	})
	return dockerfiles, err
}

// isDockerfile returns true if info looks like a Dockerfile. It must be named
// dockerfileName and be either a regular file or a symlink.
func isDockerfile(info os.FileInfo, dockerfileName string) bool {
	return info.Name() == dockerfileName && (info.Mode().IsRegular() || info.Mode()&os.ModeSymlink != 0)
}
