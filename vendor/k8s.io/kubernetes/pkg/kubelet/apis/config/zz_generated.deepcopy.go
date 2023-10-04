//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package config

import (
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	apiv1 "k8s.io/component-base/tracing/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CredentialProvider) DeepCopyInto(out *CredentialProvider) {
	*out = *in
	if in.MatchImages != nil {
		in, out := &in.MatchImages, &out.MatchImages
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DefaultCacheDuration != nil {
		in, out := &in.DefaultCacheDuration, &out.DefaultCacheDuration
		*out = new(v1.Duration)
		**out = **in
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]ExecEnvVar, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CredentialProvider.
func (in *CredentialProvider) DeepCopy() *CredentialProvider {
	if in == nil {
		return nil
	}
	out := new(CredentialProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CredentialProviderConfig) DeepCopyInto(out *CredentialProviderConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Providers != nil {
		in, out := &in.Providers, &out.Providers
		*out = make([]CredentialProvider, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CredentialProviderConfig.
func (in *CredentialProviderConfig) DeepCopy() *CredentialProviderConfig {
	if in == nil {
		return nil
	}
	out := new(CredentialProviderConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CredentialProviderConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExecEnvVar) DeepCopyInto(out *ExecEnvVar) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExecEnvVar.
func (in *ExecEnvVar) DeepCopy() *ExecEnvVar {
	if in == nil {
		return nil
	}
	out := new(ExecEnvVar)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeletAnonymousAuthentication) DeepCopyInto(out *KubeletAnonymousAuthentication) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeletAnonymousAuthentication.
func (in *KubeletAnonymousAuthentication) DeepCopy() *KubeletAnonymousAuthentication {
	if in == nil {
		return nil
	}
	out := new(KubeletAnonymousAuthentication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeletAuthentication) DeepCopyInto(out *KubeletAuthentication) {
	*out = *in
	out.X509 = in.X509
	out.Webhook = in.Webhook
	out.Anonymous = in.Anonymous
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeletAuthentication.
func (in *KubeletAuthentication) DeepCopy() *KubeletAuthentication {
	if in == nil {
		return nil
	}
	out := new(KubeletAuthentication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeletAuthorization) DeepCopyInto(out *KubeletAuthorization) {
	*out = *in
	out.Webhook = in.Webhook
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeletAuthorization.
func (in *KubeletAuthorization) DeepCopy() *KubeletAuthorization {
	if in == nil {
		return nil
	}
	out := new(KubeletAuthorization)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeletConfiguration) DeepCopyInto(out *KubeletConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.SyncFrequency = in.SyncFrequency
	out.FileCheckFrequency = in.FileCheckFrequency
	out.HTTPCheckFrequency = in.HTTPCheckFrequency
	if in.StaticPodURLHeader != nil {
		in, out := &in.StaticPodURLHeader, &out.StaticPodURLHeader
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	if in.TLSCipherSuites != nil {
		in, out := &in.TLSCipherSuites, &out.TLSCipherSuites
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.Authentication = in.Authentication
	out.Authorization = in.Authorization
	if in.ClusterDNS != nil {
		in, out := &in.ClusterDNS, &out.ClusterDNS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.StreamingConnectionIdleTimeout = in.StreamingConnectionIdleTimeout
	out.NodeStatusUpdateFrequency = in.NodeStatusUpdateFrequency
	out.NodeStatusReportFrequency = in.NodeStatusReportFrequency
	out.ImageMinimumGCAge = in.ImageMinimumGCAge
	out.VolumeStatsAggPeriod = in.VolumeStatsAggPeriod
	if in.CPUManagerPolicyOptions != nil {
		in, out := &in.CPUManagerPolicyOptions, &out.CPUManagerPolicyOptions
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.CPUManagerReconcilePeriod = in.CPUManagerReconcilePeriod
	if in.TopologyManagerPolicyOptions != nil {
		in, out := &in.TopologyManagerPolicyOptions, &out.TopologyManagerPolicyOptions
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.QOSReserved != nil {
		in, out := &in.QOSReserved, &out.QOSReserved
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.RuntimeRequestTimeout = in.RuntimeRequestTimeout
	out.CPUCFSQuotaPeriod = in.CPUCFSQuotaPeriod
	if in.EvictionHard != nil {
		in, out := &in.EvictionHard, &out.EvictionHard
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.EvictionSoft != nil {
		in, out := &in.EvictionSoft, &out.EvictionSoft
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.EvictionSoftGracePeriod != nil {
		in, out := &in.EvictionSoftGracePeriod, &out.EvictionSoftGracePeriod
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.EvictionPressureTransitionPeriod = in.EvictionPressureTransitionPeriod
	if in.EvictionMinimumReclaim != nil {
		in, out := &in.EvictionMinimumReclaim, &out.EvictionMinimumReclaim
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.FeatureGates != nil {
		in, out := &in.FeatureGates, &out.FeatureGates
		*out = make(map[string]bool, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.MemorySwap = in.MemorySwap
	if in.AllowedUnsafeSysctls != nil {
		in, out := &in.AllowedUnsafeSysctls, &out.AllowedUnsafeSysctls
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SystemReserved != nil {
		in, out := &in.SystemReserved, &out.SystemReserved
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.KubeReserved != nil {
		in, out := &in.KubeReserved, &out.KubeReserved
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.EnforceNodeAllocatable != nil {
		in, out := &in.EnforceNodeAllocatable, &out.EnforceNodeAllocatable
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Logging.DeepCopyInto(&out.Logging)
	out.ShutdownGracePeriod = in.ShutdownGracePeriod
	out.ShutdownGracePeriodCriticalPods = in.ShutdownGracePeriodCriticalPods
	if in.ShutdownGracePeriodByPodPriority != nil {
		in, out := &in.ShutdownGracePeriodByPodPriority, &out.ShutdownGracePeriodByPodPriority
		*out = make([]ShutdownGracePeriodByPodPriority, len(*in))
		copy(*out, *in)
	}
	if in.ReservedMemory != nil {
		in, out := &in.ReservedMemory, &out.ReservedMemory
		*out = make([]MemoryReservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MemoryThrottlingFactor != nil {
		in, out := &in.MemoryThrottlingFactor, &out.MemoryThrottlingFactor
		*out = new(float64)
		**out = **in
	}
	if in.RegisterWithTaints != nil {
		in, out := &in.RegisterWithTaints, &out.RegisterWithTaints
		*out = make([]corev1.Taint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tracing != nil {
		in, out := &in.Tracing, &out.Tracing
		*out = new(apiv1.TracingConfiguration)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeletConfiguration.
func (in *KubeletConfiguration) DeepCopy() *KubeletConfiguration {
	if in == nil {
		return nil
	}
	out := new(KubeletConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeletConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeletWebhookAuthentication) DeepCopyInto(out *KubeletWebhookAuthentication) {
	*out = *in
	out.CacheTTL = in.CacheTTL
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeletWebhookAuthentication.
func (in *KubeletWebhookAuthentication) DeepCopy() *KubeletWebhookAuthentication {
	if in == nil {
		return nil
	}
	out := new(KubeletWebhookAuthentication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeletWebhookAuthorization) DeepCopyInto(out *KubeletWebhookAuthorization) {
	*out = *in
	out.CacheAuthorizedTTL = in.CacheAuthorizedTTL
	out.CacheUnauthorizedTTL = in.CacheUnauthorizedTTL
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeletWebhookAuthorization.
func (in *KubeletWebhookAuthorization) DeepCopy() *KubeletWebhookAuthorization {
	if in == nil {
		return nil
	}
	out := new(KubeletWebhookAuthorization)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeletX509Authentication) DeepCopyInto(out *KubeletX509Authentication) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeletX509Authentication.
func (in *KubeletX509Authentication) DeepCopy() *KubeletX509Authentication {
	if in == nil {
		return nil
	}
	out := new(KubeletX509Authentication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemoryReservation) DeepCopyInto(out *MemoryReservation) {
	*out = *in
	if in.Limits != nil {
		in, out := &in.Limits, &out.Limits
		*out = make(corev1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemoryReservation.
func (in *MemoryReservation) DeepCopy() *MemoryReservation {
	if in == nil {
		return nil
	}
	out := new(MemoryReservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemorySwapConfiguration) DeepCopyInto(out *MemorySwapConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemorySwapConfiguration.
func (in *MemorySwapConfiguration) DeepCopy() *MemorySwapConfiguration {
	if in == nil {
		return nil
	}
	out := new(MemorySwapConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SerializedNodeConfigSource) DeepCopyInto(out *SerializedNodeConfigSource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.Source.DeepCopyInto(&out.Source)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SerializedNodeConfigSource.
func (in *SerializedNodeConfigSource) DeepCopy() *SerializedNodeConfigSource {
	if in == nil {
		return nil
	}
	out := new(SerializedNodeConfigSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SerializedNodeConfigSource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShutdownGracePeriodByPodPriority) DeepCopyInto(out *ShutdownGracePeriodByPodPriority) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShutdownGracePeriodByPodPriority.
func (in *ShutdownGracePeriodByPodPriority) DeepCopy() *ShutdownGracePeriodByPodPriority {
	if in == nil {
		return nil
	}
	out := new(ShutdownGracePeriodByPodPriority)
	in.DeepCopyInto(out)
	return out
}
