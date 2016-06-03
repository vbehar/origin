// +build !ignore_autogenerated

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package api

import (
	api "k8s.io/kubernetes/pkg/api"
	unversioned "k8s.io/kubernetes/pkg/api/unversioned"
	conversion "k8s.io/kubernetes/pkg/conversion"
)

func init() {
	if err := api.Scheme.AddGeneratedDeepCopyFuncs(
		DeepCopy_api_PodSecurityPolicyReview,
		DeepCopy_api_PodSecurityPolicyReviewSpec,
		DeepCopy_api_PodSecurityPolicyReviewStatus,
		DeepCopy_api_PodSecurityPolicySelfSubjectReview,
		DeepCopy_api_PodSecurityPolicySelfSubjectReviewSpec,
		DeepCopy_api_PodSecurityPolicySubjectReview,
		DeepCopy_api_PodSecurityPolicySubjectReviewSpec,
		DeepCopy_api_PodSecurityPolicySubjectReviewStatus,
		DeepCopy_api_ServiceAccountPodSecurityPolicyReviewStatus,
	); err != nil {
		// if one of the deep copy functions is malformed, detect it immediately.
		panic(err)
	}
}

func DeepCopy_api_PodSecurityPolicyReview(in PodSecurityPolicyReview, out *PodSecurityPolicyReview, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := DeepCopy_api_PodSecurityPolicyReviewSpec(in.Spec, &out.Spec, c); err != nil {
		return err
	}
	if err := DeepCopy_api_PodSecurityPolicyReviewStatus(in.Status, &out.Status, c); err != nil {
		return err
	}
	return nil
}

func DeepCopy_api_PodSecurityPolicyReviewSpec(in PodSecurityPolicyReviewSpec, out *PodSecurityPolicyReviewSpec, c *conversion.Cloner) error {
	if err := api.DeepCopy_api_PodSpec(in.PodSpec, &out.PodSpec, c); err != nil {
		return err
	}
	if in.ServiceAccountNames != nil {
		in, out := in.ServiceAccountNames, &out.ServiceAccountNames
		*out = make([]string, len(in))
		copy(*out, in)
	} else {
		out.ServiceAccountNames = nil
	}
	return nil
}

func DeepCopy_api_PodSecurityPolicyReviewStatus(in PodSecurityPolicyReviewStatus, out *PodSecurityPolicyReviewStatus, c *conversion.Cloner) error {
	if in.AllowedServiceAccounts != nil {
		in, out := in.AllowedServiceAccounts, &out.AllowedServiceAccounts
		*out = make([]ServiceAccountPodSecurityPolicyReviewStatus, len(in))
		for i := range in {
			if err := DeepCopy_api_ServiceAccountPodSecurityPolicyReviewStatus(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.AllowedServiceAccounts = nil
	}
	return nil
}

func DeepCopy_api_PodSecurityPolicySelfSubjectReview(in PodSecurityPolicySelfSubjectReview, out *PodSecurityPolicySelfSubjectReview, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := DeepCopy_api_PodSecurityPolicySelfSubjectReviewSpec(in.Spec, &out.Spec, c); err != nil {
		return err
	}
	if err := DeepCopy_api_PodSecurityPolicySubjectReviewStatus(in.Status, &out.Status, c); err != nil {
		return err
	}
	return nil
}

func DeepCopy_api_PodSecurityPolicySelfSubjectReviewSpec(in PodSecurityPolicySelfSubjectReviewSpec, out *PodSecurityPolicySelfSubjectReviewSpec, c *conversion.Cloner) error {
	if err := api.DeepCopy_api_PodSpec(in.PodSpec, &out.PodSpec, c); err != nil {
		return err
	}
	return nil
}

func DeepCopy_api_PodSecurityPolicySubjectReview(in PodSecurityPolicySubjectReview, out *PodSecurityPolicySubjectReview, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := DeepCopy_api_PodSecurityPolicySubjectReviewSpec(in.Spec, &out.Spec, c); err != nil {
		return err
	}
	if err := DeepCopy_api_PodSecurityPolicySubjectReviewStatus(in.Status, &out.Status, c); err != nil {
		return err
	}
	return nil
}

func DeepCopy_api_PodSecurityPolicySubjectReviewSpec(in PodSecurityPolicySubjectReviewSpec, out *PodSecurityPolicySubjectReviewSpec, c *conversion.Cloner) error {
	if err := api.DeepCopy_api_PodSpec(in.PodSpec, &out.PodSpec, c); err != nil {
		return err
	}
	out.User = in.User
	if in.Groups != nil {
		in, out := in.Groups, &out.Groups
		*out = make([]string, len(in))
		copy(*out, in)
	} else {
		out.Groups = nil
	}
	return nil
}

func DeepCopy_api_PodSecurityPolicySubjectReviewStatus(in PodSecurityPolicySubjectReviewStatus, out *PodSecurityPolicySubjectReviewStatus, c *conversion.Cloner) error {
	if in.AllowedBy != nil {
		in, out := in.AllowedBy, &out.AllowedBy
		*out = new(api.ObjectReference)
		if err := api.DeepCopy_api_ObjectReference(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.AllowedBy = nil
	}
	out.Reason = in.Reason
	if err := api.DeepCopy_api_PodSpec(in.PodSpec, &out.PodSpec, c); err != nil {
		return err
	}
	return nil
}

func DeepCopy_api_ServiceAccountPodSecurityPolicyReviewStatus(in ServiceAccountPodSecurityPolicyReviewStatus, out *ServiceAccountPodSecurityPolicyReviewStatus, c *conversion.Cloner) error {
	if err := DeepCopy_api_PodSecurityPolicySubjectReviewStatus(in.PodSecurityPolicySubjectReviewStatus, &out.PodSecurityPolicySubjectReviewStatus, c); err != nil {
		return err
	}
	out.Name = in.Name
	return nil
}
