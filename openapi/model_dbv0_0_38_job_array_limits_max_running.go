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

// checks if the Dbv0038JobArrayLimitsMaxRunning type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0038JobArrayLimitsMaxRunning{}

// Dbv0038JobArrayLimitsMaxRunning Limits on array settings
type Dbv0038JobArrayLimitsMaxRunning struct {
	// Max running tasks in array at any one time
	Tasks *int32 `json:"tasks,omitempty"`
}

// NewDbv0038JobArrayLimitsMaxRunning instantiates a new Dbv0038JobArrayLimitsMaxRunning object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0038JobArrayLimitsMaxRunning() *Dbv0038JobArrayLimitsMaxRunning {
	this := Dbv0038JobArrayLimitsMaxRunning{}
	return &this
}

// NewDbv0038JobArrayLimitsMaxRunningWithDefaults instantiates a new Dbv0038JobArrayLimitsMaxRunning object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0038JobArrayLimitsMaxRunningWithDefaults() *Dbv0038JobArrayLimitsMaxRunning {
	this := Dbv0038JobArrayLimitsMaxRunning{}
	return &this
}

// GetTasks returns the Tasks field value if set, zero value otherwise.
func (o *Dbv0038JobArrayLimitsMaxRunning) GetTasks() int32 {
	if o == nil || IsNil(o.Tasks) {
		var ret int32
		return ret
	}
	return *o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038JobArrayLimitsMaxRunning) GetTasksOk() (*int32, bool) {
	if o == nil || IsNil(o.Tasks) {
		return nil, false
	}
	return o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *Dbv0038JobArrayLimitsMaxRunning) HasTasks() bool {
	if o != nil && !IsNil(o.Tasks) {
		return true
	}

	return false
}

// SetTasks gets a reference to the given int32 and assigns it to the Tasks field.
func (o *Dbv0038JobArrayLimitsMaxRunning) SetTasks(v int32) {
	o.Tasks = &v
}

func (o Dbv0038JobArrayLimitsMaxRunning) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0038JobArrayLimitsMaxRunning) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tasks) {
		toSerialize["tasks"] = o.Tasks
	}
	return toSerialize, nil
}

type NullableDbv0038JobArrayLimitsMaxRunning struct {
	value *Dbv0038JobArrayLimitsMaxRunning
	isSet bool
}

func (v NullableDbv0038JobArrayLimitsMaxRunning) Get() *Dbv0038JobArrayLimitsMaxRunning {
	return v.value
}

func (v *NullableDbv0038JobArrayLimitsMaxRunning) Set(val *Dbv0038JobArrayLimitsMaxRunning) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0038JobArrayLimitsMaxRunning) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0038JobArrayLimitsMaxRunning) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0038JobArrayLimitsMaxRunning(val *Dbv0038JobArrayLimitsMaxRunning) *NullableDbv0038JobArrayLimitsMaxRunning {
	return &NullableDbv0038JobArrayLimitsMaxRunning{value: val, isSet: true}
}

func (v NullableDbv0038JobArrayLimitsMaxRunning) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0038JobArrayLimitsMaxRunning) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


