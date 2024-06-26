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

// checks if the Dbv0038QosPreempt type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0038QosPreempt{}

// Dbv0038QosPreempt Preemption settings
type Dbv0038QosPreempt struct {
	// List of preemptable QOS
	List []string `json:"list,omitempty"`
	// List of preemption modes
	Mode []string `json:"mode,omitempty"`
	// Grace period (s) before jobs can preempted
	ExemptTime *int32 `json:"exempt_time,omitempty"`
}

// NewDbv0038QosPreempt instantiates a new Dbv0038QosPreempt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0038QosPreempt() *Dbv0038QosPreempt {
	this := Dbv0038QosPreempt{}
	return &this
}

// NewDbv0038QosPreemptWithDefaults instantiates a new Dbv0038QosPreempt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0038QosPreemptWithDefaults() *Dbv0038QosPreempt {
	this := Dbv0038QosPreempt{}
	return &this
}

// GetList returns the List field value if set, zero value otherwise.
func (o *Dbv0038QosPreempt) GetList() []string {
	if o == nil || IsNil(o.List) {
		var ret []string
		return ret
	}
	return o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038QosPreempt) GetListOk() ([]string, bool) {
	if o == nil || IsNil(o.List) {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *Dbv0038QosPreempt) HasList() bool {
	if o != nil && !IsNil(o.List) {
		return true
	}

	return false
}

// SetList gets a reference to the given []string and assigns it to the List field.
func (o *Dbv0038QosPreempt) SetList(v []string) {
	o.List = v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *Dbv0038QosPreempt) GetMode() []string {
	if o == nil || IsNil(o.Mode) {
		var ret []string
		return ret
	}
	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038QosPreempt) GetModeOk() ([]string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *Dbv0038QosPreempt) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given []string and assigns it to the Mode field.
func (o *Dbv0038QosPreempt) SetMode(v []string) {
	o.Mode = v
}

// GetExemptTime returns the ExemptTime field value if set, zero value otherwise.
func (o *Dbv0038QosPreempt) GetExemptTime() int32 {
	if o == nil || IsNil(o.ExemptTime) {
		var ret int32
		return ret
	}
	return *o.ExemptTime
}

// GetExemptTimeOk returns a tuple with the ExemptTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038QosPreempt) GetExemptTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.ExemptTime) {
		return nil, false
	}
	return o.ExemptTime, true
}

// HasExemptTime returns a boolean if a field has been set.
func (o *Dbv0038QosPreempt) HasExemptTime() bool {
	if o != nil && !IsNil(o.ExemptTime) {
		return true
	}

	return false
}

// SetExemptTime gets a reference to the given int32 and assigns it to the ExemptTime field.
func (o *Dbv0038QosPreempt) SetExemptTime(v int32) {
	o.ExemptTime = &v
}

func (o Dbv0038QosPreempt) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0038QosPreempt) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.List) {
		toSerialize["list"] = o.List
	}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !IsNil(o.ExemptTime) {
		toSerialize["exempt_time"] = o.ExemptTime
	}
	return toSerialize, nil
}

type NullableDbv0038QosPreempt struct {
	value *Dbv0038QosPreempt
	isSet bool
}

func (v NullableDbv0038QosPreempt) Get() *Dbv0038QosPreempt {
	return v.value
}

func (v *NullableDbv0038QosPreempt) Set(val *Dbv0038QosPreempt) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0038QosPreempt) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0038QosPreempt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0038QosPreempt(val *Dbv0038QosPreempt) *NullableDbv0038QosPreempt {
	return &NullableDbv0038QosPreempt{value: val, isSet: true}
}

func (v NullableDbv0038QosPreempt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0038QosPreempt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


