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

// checks if the Dbv0038JobStepStatisticsEnergy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0038JobStepStatisticsEnergy{}

// Dbv0038JobStepStatisticsEnergy Statistics of energy
type Dbv0038JobStepStatisticsEnergy struct {
	// Energy consumed during step
	Consumed *int32 `json:"consumed,omitempty"`
}

// NewDbv0038JobStepStatisticsEnergy instantiates a new Dbv0038JobStepStatisticsEnergy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0038JobStepStatisticsEnergy() *Dbv0038JobStepStatisticsEnergy {
	this := Dbv0038JobStepStatisticsEnergy{}
	return &this
}

// NewDbv0038JobStepStatisticsEnergyWithDefaults instantiates a new Dbv0038JobStepStatisticsEnergy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0038JobStepStatisticsEnergyWithDefaults() *Dbv0038JobStepStatisticsEnergy {
	this := Dbv0038JobStepStatisticsEnergy{}
	return &this
}

// GetConsumed returns the Consumed field value if set, zero value otherwise.
func (o *Dbv0038JobStepStatisticsEnergy) GetConsumed() int32 {
	if o == nil || IsNil(o.Consumed) {
		var ret int32
		return ret
	}
	return *o.Consumed
}

// GetConsumedOk returns a tuple with the Consumed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038JobStepStatisticsEnergy) GetConsumedOk() (*int32, bool) {
	if o == nil || IsNil(o.Consumed) {
		return nil, false
	}
	return o.Consumed, true
}

// HasConsumed returns a boolean if a field has been set.
func (o *Dbv0038JobStepStatisticsEnergy) HasConsumed() bool {
	if o != nil && !IsNil(o.Consumed) {
		return true
	}

	return false
}

// SetConsumed gets a reference to the given int32 and assigns it to the Consumed field.
func (o *Dbv0038JobStepStatisticsEnergy) SetConsumed(v int32) {
	o.Consumed = &v
}

func (o Dbv0038JobStepStatisticsEnergy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0038JobStepStatisticsEnergy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Consumed) {
		toSerialize["consumed"] = o.Consumed
	}
	return toSerialize, nil
}

type NullableDbv0038JobStepStatisticsEnergy struct {
	value *Dbv0038JobStepStatisticsEnergy
	isSet bool
}

func (v NullableDbv0038JobStepStatisticsEnergy) Get() *Dbv0038JobStepStatisticsEnergy {
	return v.value
}

func (v *NullableDbv0038JobStepStatisticsEnergy) Set(val *Dbv0038JobStepStatisticsEnergy) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0038JobStepStatisticsEnergy) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0038JobStepStatisticsEnergy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0038JobStepStatisticsEnergy(val *Dbv0038JobStepStatisticsEnergy) *NullableDbv0038JobStepStatisticsEnergy {
	return &NullableDbv0038JobStepStatisticsEnergy{value: val, isSet: true}
}

func (v NullableDbv0038JobStepStatisticsEnergy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0038JobStepStatisticsEnergy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


