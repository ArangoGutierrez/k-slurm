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

// checks if the Dbv0038JobMcs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0038JobMcs{}

// Dbv0038JobMcs Multi-Category Security
type Dbv0038JobMcs struct {
	// Assigned MCS label
	Label *string `json:"label,omitempty"`
}

// NewDbv0038JobMcs instantiates a new Dbv0038JobMcs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0038JobMcs() *Dbv0038JobMcs {
	this := Dbv0038JobMcs{}
	return &this
}

// NewDbv0038JobMcsWithDefaults instantiates a new Dbv0038JobMcs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0038JobMcsWithDefaults() *Dbv0038JobMcs {
	this := Dbv0038JobMcs{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *Dbv0038JobMcs) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038JobMcs) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *Dbv0038JobMcs) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *Dbv0038JobMcs) SetLabel(v string) {
	o.Label = &v
}

func (o Dbv0038JobMcs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0038JobMcs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	return toSerialize, nil
}

type NullableDbv0038JobMcs struct {
	value *Dbv0038JobMcs
	isSet bool
}

func (v NullableDbv0038JobMcs) Get() *Dbv0038JobMcs {
	return v.value
}

func (v *NullableDbv0038JobMcs) Set(val *Dbv0038JobMcs) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0038JobMcs) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0038JobMcs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0038JobMcs(val *Dbv0038JobMcs) *NullableDbv0038JobMcs {
	return &NullableDbv0038JobMcs{value: val, isSet: true}
}

func (v NullableDbv0038JobMcs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0038JobMcs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


