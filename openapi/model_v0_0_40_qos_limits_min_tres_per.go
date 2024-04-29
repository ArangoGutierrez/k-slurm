/*
Slurm Rest API

API to access and control Slurm DB.

API version: Slurm-23.11.5&openapi/dbv0.0.39&openapi/slurmctld&openapi/v0.0.39&openapi/dbv0.0.38&openapi/slurmdbd&openapi/v0.0.38
Contact: sales@schedmd.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the V0040QosLimitsMinTresPer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040QosLimitsMinTresPer{}

// V0040QosLimitsMinTresPer struct for V0040QosLimitsMinTresPer
type V0040QosLimitsMinTresPer struct {
	Job []V0040Tres `json:"job,omitempty"`
}

// NewV0040QosLimitsMinTresPer instantiates a new V0040QosLimitsMinTresPer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040QosLimitsMinTresPer() *V0040QosLimitsMinTresPer {
	this := V0040QosLimitsMinTresPer{}
	return &this
}

// NewV0040QosLimitsMinTresPerWithDefaults instantiates a new V0040QosLimitsMinTresPer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040QosLimitsMinTresPerWithDefaults() *V0040QosLimitsMinTresPer {
	this := V0040QosLimitsMinTresPer{}
	return &this
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *V0040QosLimitsMinTresPer) GetJob() []V0040Tres {
	if o == nil || IsNil(o.Job) {
		var ret []V0040Tres
		return ret
	}
	return o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040QosLimitsMinTresPer) GetJobOk() ([]V0040Tres, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *V0040QosLimitsMinTresPer) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given []V0040Tres and assigns it to the Job field.
func (o *V0040QosLimitsMinTresPer) SetJob(v []V0040Tres) {
	o.Job = v
}

func (o V0040QosLimitsMinTresPer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040QosLimitsMinTresPer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	return toSerialize, nil
}

type NullableV0040QosLimitsMinTresPer struct {
	value *V0040QosLimitsMinTresPer
	isSet bool
}

func (v NullableV0040QosLimitsMinTresPer) Get() *V0040QosLimitsMinTresPer {
	return v.value
}

func (v *NullableV0040QosLimitsMinTresPer) Set(val *V0040QosLimitsMinTresPer) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040QosLimitsMinTresPer) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040QosLimitsMinTresPer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040QosLimitsMinTresPer(val *V0040QosLimitsMinTresPer) *NullableV0040QosLimitsMinTresPer {
	return &NullableV0040QosLimitsMinTresPer{value: val, isSet: true}
}

func (v NullableV0040QosLimitsMinTresPer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040QosLimitsMinTresPer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

