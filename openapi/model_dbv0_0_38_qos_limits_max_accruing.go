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

// checks if the Dbv0038QosLimitsMaxAccruing type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0038QosLimitsMaxAccruing{}

// Dbv0038QosLimitsMaxAccruing Limits on accruing priority
type Dbv0038QosLimitsMaxAccruing struct {
	Per *Dbv0038QosLimitsMaxAccruingPer `json:"per,omitempty"`
}

// NewDbv0038QosLimitsMaxAccruing instantiates a new Dbv0038QosLimitsMaxAccruing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0038QosLimitsMaxAccruing() *Dbv0038QosLimitsMaxAccruing {
	this := Dbv0038QosLimitsMaxAccruing{}
	return &this
}

// NewDbv0038QosLimitsMaxAccruingWithDefaults instantiates a new Dbv0038QosLimitsMaxAccruing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0038QosLimitsMaxAccruingWithDefaults() *Dbv0038QosLimitsMaxAccruing {
	this := Dbv0038QosLimitsMaxAccruing{}
	return &this
}

// GetPer returns the Per field value if set, zero value otherwise.
func (o *Dbv0038QosLimitsMaxAccruing) GetPer() Dbv0038QosLimitsMaxAccruingPer {
	if o == nil || IsNil(o.Per) {
		var ret Dbv0038QosLimitsMaxAccruingPer
		return ret
	}
	return *o.Per
}

// GetPerOk returns a tuple with the Per field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038QosLimitsMaxAccruing) GetPerOk() (*Dbv0038QosLimitsMaxAccruingPer, bool) {
	if o == nil || IsNil(o.Per) {
		return nil, false
	}
	return o.Per, true
}

// HasPer returns a boolean if a field has been set.
func (o *Dbv0038QosLimitsMaxAccruing) HasPer() bool {
	if o != nil && !IsNil(o.Per) {
		return true
	}

	return false
}

// SetPer gets a reference to the given Dbv0038QosLimitsMaxAccruingPer and assigns it to the Per field.
func (o *Dbv0038QosLimitsMaxAccruing) SetPer(v Dbv0038QosLimitsMaxAccruingPer) {
	o.Per = &v
}

func (o Dbv0038QosLimitsMaxAccruing) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0038QosLimitsMaxAccruing) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Per) {
		toSerialize["per"] = o.Per
	}
	return toSerialize, nil
}

type NullableDbv0038QosLimitsMaxAccruing struct {
	value *Dbv0038QosLimitsMaxAccruing
	isSet bool
}

func (v NullableDbv0038QosLimitsMaxAccruing) Get() *Dbv0038QosLimitsMaxAccruing {
	return v.value
}

func (v *NullableDbv0038QosLimitsMaxAccruing) Set(val *Dbv0038QosLimitsMaxAccruing) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0038QosLimitsMaxAccruing) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0038QosLimitsMaxAccruing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0038QosLimitsMaxAccruing(val *Dbv0038QosLimitsMaxAccruing) *NullableDbv0038QosLimitsMaxAccruing {
	return &NullableDbv0038QosLimitsMaxAccruing{value: val, isSet: true}
}

func (v NullableDbv0038QosLimitsMaxAccruing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0038QosLimitsMaxAccruing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


