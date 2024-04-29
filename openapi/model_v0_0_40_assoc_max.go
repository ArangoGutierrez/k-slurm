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

// checks if the V0040AssocMax type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040AssocMax{}

// V0040AssocMax struct for V0040AssocMax
type V0040AssocMax struct {
	Jobs *V0040AssocMaxJobs `json:"jobs,omitempty"`
	Tres *V0040AssocMaxTres `json:"tres,omitempty"`
	Per *V0040AssocMaxPer `json:"per,omitempty"`
}

// NewV0040AssocMax instantiates a new V0040AssocMax object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040AssocMax() *V0040AssocMax {
	this := V0040AssocMax{}
	return &this
}

// NewV0040AssocMaxWithDefaults instantiates a new V0040AssocMax object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040AssocMaxWithDefaults() *V0040AssocMax {
	this := V0040AssocMax{}
	return &this
}

// GetJobs returns the Jobs field value if set, zero value otherwise.
func (o *V0040AssocMax) GetJobs() V0040AssocMaxJobs {
	if o == nil || IsNil(o.Jobs) {
		var ret V0040AssocMaxJobs
		return ret
	}
	return *o.Jobs
}

// GetJobsOk returns a tuple with the Jobs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AssocMax) GetJobsOk() (*V0040AssocMaxJobs, bool) {
	if o == nil || IsNil(o.Jobs) {
		return nil, false
	}
	return o.Jobs, true
}

// HasJobs returns a boolean if a field has been set.
func (o *V0040AssocMax) HasJobs() bool {
	if o != nil && !IsNil(o.Jobs) {
		return true
	}

	return false
}

// SetJobs gets a reference to the given V0040AssocMaxJobs and assigns it to the Jobs field.
func (o *V0040AssocMax) SetJobs(v V0040AssocMaxJobs) {
	o.Jobs = &v
}

// GetTres returns the Tres field value if set, zero value otherwise.
func (o *V0040AssocMax) GetTres() V0040AssocMaxTres {
	if o == nil || IsNil(o.Tres) {
		var ret V0040AssocMaxTres
		return ret
	}
	return *o.Tres
}

// GetTresOk returns a tuple with the Tres field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AssocMax) GetTresOk() (*V0040AssocMaxTres, bool) {
	if o == nil || IsNil(o.Tres) {
		return nil, false
	}
	return o.Tres, true
}

// HasTres returns a boolean if a field has been set.
func (o *V0040AssocMax) HasTres() bool {
	if o != nil && !IsNil(o.Tres) {
		return true
	}

	return false
}

// SetTres gets a reference to the given V0040AssocMaxTres and assigns it to the Tres field.
func (o *V0040AssocMax) SetTres(v V0040AssocMaxTres) {
	o.Tres = &v
}

// GetPer returns the Per field value if set, zero value otherwise.
func (o *V0040AssocMax) GetPer() V0040AssocMaxPer {
	if o == nil || IsNil(o.Per) {
		var ret V0040AssocMaxPer
		return ret
	}
	return *o.Per
}

// GetPerOk returns a tuple with the Per field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AssocMax) GetPerOk() (*V0040AssocMaxPer, bool) {
	if o == nil || IsNil(o.Per) {
		return nil, false
	}
	return o.Per, true
}

// HasPer returns a boolean if a field has been set.
func (o *V0040AssocMax) HasPer() bool {
	if o != nil && !IsNil(o.Per) {
		return true
	}

	return false
}

// SetPer gets a reference to the given V0040AssocMaxPer and assigns it to the Per field.
func (o *V0040AssocMax) SetPer(v V0040AssocMaxPer) {
	o.Per = &v
}

func (o V0040AssocMax) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040AssocMax) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Jobs) {
		toSerialize["jobs"] = o.Jobs
	}
	if !IsNil(o.Tres) {
		toSerialize["tres"] = o.Tres
	}
	if !IsNil(o.Per) {
		toSerialize["per"] = o.Per
	}
	return toSerialize, nil
}

type NullableV0040AssocMax struct {
	value *V0040AssocMax
	isSet bool
}

func (v NullableV0040AssocMax) Get() *V0040AssocMax {
	return v.value
}

func (v *NullableV0040AssocMax) Set(val *V0040AssocMax) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040AssocMax) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040AssocMax) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040AssocMax(val *V0040AssocMax) *NullableV0040AssocMax {
	return &NullableV0040AssocMax{value: val, isSet: true}
}

func (v NullableV0040AssocMax) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040AssocMax) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

