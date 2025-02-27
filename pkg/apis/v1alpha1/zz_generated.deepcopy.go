//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 The Arbiter Authors.

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

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionCondition) DeepCopyInto(out *ActionCondition) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionCondition.
func (in *ActionCondition) DeepCopy() *ActionCondition {
	if in == nil {
		return nil
	}
	out := new(ActionCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionInfo) DeepCopyInto(out *ActionInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionInfo.
func (in *ActionInfo) DeepCopy() *ActionInfo {
	if in == nil {
		return nil
	}
	out := new(ActionInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	in.LastHeartbeatTime.DeepCopyInto(&out.LastHeartbeatTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObservabilityActionPolicy) DeepCopyInto(out *ObservabilityActionPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObservabilityActionPolicy.
func (in *ObservabilityActionPolicy) DeepCopy() *ObservabilityActionPolicy {
	if in == nil {
		return nil
	}
	out := new(ObservabilityActionPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ObservabilityActionPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObservabilityActionPolicyList) DeepCopyInto(out *ObservabilityActionPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ObservabilityActionPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObservabilityActionPolicyList.
func (in *ObservabilityActionPolicyList) DeepCopy() *ObservabilityActionPolicyList {
	if in == nil {
		return nil
	}
	out := new(ObservabilityActionPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ObservabilityActionPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObservabilityActionPolicySpec) DeepCopyInto(out *ObservabilityActionPolicySpec) {
	*out = *in
	out.Condition = in.Condition
	if in.Executors != nil {
		in, out := &in.Executors, &out.Executors
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ActionData != nil {
		in, out := &in.ActionData, &out.ActionData
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObservabilityActionPolicySpec.
func (in *ObservabilityActionPolicySpec) DeepCopy() *ObservabilityActionPolicySpec {
	if in == nil {
		return nil
	}
	out := new(ObservabilityActionPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObservabilityActionPolicyStatus) DeepCopyInto(out *ObservabilityActionPolicyStatus) {
	*out = *in
	if in.ActionInfo != nil {
		in, out := &in.ActionInfo, &out.ActionInfo
		*out = make([]ActionInfo, len(*in))
		copy(*out, *in)
	}
	in.TimeInfo.DeepCopyInto(&out.TimeInfo)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObservabilityActionPolicyStatus.
func (in *ObservabilityActionPolicyStatus) DeepCopy() *ObservabilityActionPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(ObservabilityActionPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObservabilityIndicant) DeepCopyInto(out *ObservabilityIndicant) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObservabilityIndicant.
func (in *ObservabilityIndicant) DeepCopy() *ObservabilityIndicant {
	if in == nil {
		return nil
	}
	out := new(ObservabilityIndicant)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ObservabilityIndicant) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObservabilityIndicantList) DeepCopyInto(out *ObservabilityIndicantList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ObservabilityIndicant, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObservabilityIndicantList.
func (in *ObservabilityIndicantList) DeepCopy() *ObservabilityIndicantList {
	if in == nil {
		return nil
	}
	out := new(ObservabilityIndicantList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ObservabilityIndicantList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObservabilityIndicantSpec) DeepCopyInto(out *ObservabilityIndicantSpec) {
	*out = *in
	in.TargetRef.DeepCopyInto(&out.TargetRef)
	if in.Metric != nil {
		in, out := &in.Metric, &out.Metric
		*out = new(ObservabilityIndicantSpecMetric)
		(*in).DeepCopyInto(*out)
	}
	if in.Trace != nil {
		in, out := &in.Trace, &out.Trace
		*out = new(ObservabilityIndicantSpecTrace)
		**out = **in
	}
	if in.Log != nil {
		in, out := &in.Log, &out.Log
		*out = new(ObservabilityIndicantSpecLog)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObservabilityIndicantSpec.
func (in *ObservabilityIndicantSpec) DeepCopy() *ObservabilityIndicantSpec {
	if in == nil {
		return nil
	}
	out := new(ObservabilityIndicantSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObservabilityIndicantSpecLog) DeepCopyInto(out *ObservabilityIndicantSpecLog) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObservabilityIndicantSpecLog.
func (in *ObservabilityIndicantSpecLog) DeepCopy() *ObservabilityIndicantSpecLog {
	if in == nil {
		return nil
	}
	out := new(ObservabilityIndicantSpecLog)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObservabilityIndicantSpecMetric) DeepCopyInto(out *ObservabilityIndicantSpecMetric) {
	*out = *in
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = make(map[string]ObservabilityIndicantSpecMetricAbility, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObservabilityIndicantSpecMetric.
func (in *ObservabilityIndicantSpecMetric) DeepCopy() *ObservabilityIndicantSpecMetric {
	if in == nil {
		return nil
	}
	out := new(ObservabilityIndicantSpecMetric)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObservabilityIndicantSpecMetricAbility) DeepCopyInto(out *ObservabilityIndicantSpecMetricAbility) {
	*out = *in
	if in.Aggregations != nil {
		in, out := &in.Aggregations, &out.Aggregations
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObservabilityIndicantSpecMetricAbility.
func (in *ObservabilityIndicantSpecMetricAbility) DeepCopy() *ObservabilityIndicantSpecMetricAbility {
	if in == nil {
		return nil
	}
	out := new(ObservabilityIndicantSpecMetricAbility)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObservabilityIndicantSpecTargetRef) DeepCopyInto(out *ObservabilityIndicantSpecTargetRef) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObservabilityIndicantSpecTargetRef.
func (in *ObservabilityIndicantSpecTargetRef) DeepCopy() *ObservabilityIndicantSpecTargetRef {
	if in == nil {
		return nil
	}
	out := new(ObservabilityIndicantSpecTargetRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObservabilityIndicantSpecTrace) DeepCopyInto(out *ObservabilityIndicantSpecTrace) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObservabilityIndicantSpecTrace.
func (in *ObservabilityIndicantSpecTrace) DeepCopy() *ObservabilityIndicantSpecTrace {
	if in == nil {
		return nil
	}
	out := new(ObservabilityIndicantSpecTrace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObservabilityIndicantStatus) DeepCopyInto(out *ObservabilityIndicantStatus) {
	*out = *in
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = make(map[string][]ObservabilityIndicantStatusMetricInfo, len(*in))
		for key, val := range *in {
			var outVal []ObservabilityIndicantStatusMetricInfo
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]ObservabilityIndicantStatusMetricInfo, len(*in))
				for i := range *in {
					(*in)[i].DeepCopyInto(&(*out)[i])
				}
			}
			(*out)[key] = outVal
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObservabilityIndicantStatus.
func (in *ObservabilityIndicantStatus) DeepCopy() *ObservabilityIndicantStatus {
	if in == nil {
		return nil
	}
	out := new(ObservabilityIndicantStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObservabilityIndicantStatusAggregationItem) DeepCopyInto(out *ObservabilityIndicantStatusAggregationItem) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObservabilityIndicantStatusAggregationItem.
func (in *ObservabilityIndicantStatusAggregationItem) DeepCopy() *ObservabilityIndicantStatusAggregationItem {
	if in == nil {
		return nil
	}
	out := new(ObservabilityIndicantStatusAggregationItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObservabilityIndicantStatusMetricInfo) DeepCopyInto(out *ObservabilityIndicantStatusMetricInfo) {
	*out = *in
	if in.Records != nil {
		in, out := &in.Records, &out.Records
		*out = make([]Record, len(*in))
		copy(*out, *in)
	}
	in.StartTime.DeepCopyInto(&out.StartTime)
	in.EndTime.DeepCopyInto(&out.EndTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObservabilityIndicantStatusMetricInfo.
func (in *ObservabilityIndicantStatusMetricInfo) DeepCopy() *ObservabilityIndicantStatusMetricInfo {
	if in == nil {
		return nil
	}
	out := new(ObservabilityIndicantStatusMetricInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObservabilityIndicantStatusTimeInfo) DeepCopyInto(out *ObservabilityIndicantStatusTimeInfo) {
	*out = *in
	in.StartTime.DeepCopyInto(&out.StartTime)
	in.EndTime.DeepCopyInto(&out.EndTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObservabilityIndicantStatusTimeInfo.
func (in *ObservabilityIndicantStatusTimeInfo) DeepCopy() *ObservabilityIndicantStatusTimeInfo {
	if in == nil {
		return nil
	}
	out := new(ObservabilityIndicantStatusTimeInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObservabilityTimeInfo) DeepCopyInto(out *ObservabilityTimeInfo) {
	*out = *in
	in.StartTime.DeepCopyInto(&out.StartTime)
	in.EndTime.DeepCopyInto(&out.EndTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObservabilityTimeInfo.
func (in *ObservabilityTimeInfo) DeepCopy() *ObservabilityTimeInfo {
	if in == nil {
		return nil
	}
	out := new(ObservabilityTimeInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Record) DeepCopyInto(out *Record) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Record.
func (in *Record) DeepCopy() *Record {
	if in == nil {
		return nil
	}
	out := new(Record)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Scheduler) DeepCopyInto(out *Scheduler) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Scheduler.
func (in *Scheduler) DeepCopy() *Scheduler {
	if in == nil {
		return nil
	}
	out := new(Scheduler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Scheduler) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulerList) DeepCopyInto(out *SchedulerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Scheduler, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulerList.
func (in *SchedulerList) DeepCopy() *SchedulerList {
	if in == nil {
		return nil
	}
	out := new(SchedulerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SchedulerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulerSpec) DeepCopyInto(out *SchedulerSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulerSpec.
func (in *SchedulerSpec) DeepCopy() *SchedulerSpec {
	if in == nil {
		return nil
	}
	out := new(SchedulerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulerStatus) DeepCopyInto(out *SchedulerStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulerStatus.
func (in *SchedulerStatus) DeepCopy() *SchedulerStatus {
	if in == nil {
		return nil
	}
	out := new(SchedulerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Score) DeepCopyInto(out *Score) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Score.
func (in *Score) DeepCopy() *Score {
	if in == nil {
		return nil
	}
	out := new(Score)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Score) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScoreList) DeepCopyInto(out *ScoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Score, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScoreList.
func (in *ScoreList) DeepCopy() *ScoreList {
	if in == nil {
		return nil
	}
	out := new(ScoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScoreSpec) DeepCopyInto(out *ScoreSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScoreSpec.
func (in *ScoreSpec) DeepCopy() *ScoreSpec {
	if in == nil {
		return nil
	}
	out := new(ScoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScoreStatus) DeepCopyInto(out *ScoreStatus) {
	*out = *in
	in.ScheduleStartTime.DeepCopyInto(&out.ScheduleStartTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScoreStatus.
func (in *ScoreStatus) DeepCopy() *ScoreStatus {
	if in == nil {
		return nil
	}
	out := new(ScoreStatus)
	in.DeepCopyInto(out)
	return out
}
