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

// checks if the V0040PartitionInfoTres type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040PartitionInfoTres{}

// V0040PartitionInfoTres struct for V0040PartitionInfoTres
type V0040PartitionInfoTres struct {
	BillingWeights *string `json:"billing_weights,omitempty"`
	Configured *string `json:"configured,omitempty"`
}

// NewV0040PartitionInfoTres instantiates a new V0040PartitionInfoTres object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040PartitionInfoTres() *V0040PartitionInfoTres {
	this := V0040PartitionInfoTres{}
	return &this
}

// NewV0040PartitionInfoTresWithDefaults instantiates a new V0040PartitionInfoTres object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040PartitionInfoTresWithDefaults() *V0040PartitionInfoTres {
	this := V0040PartitionInfoTres{}
	return &this
}

// GetBillingWeights returns the BillingWeights field value if set, zero value otherwise.
func (o *V0040PartitionInfoTres) GetBillingWeights() string {
	if o == nil || IsNil(o.BillingWeights) {
		var ret string
		return ret
	}
	return *o.BillingWeights
}

// GetBillingWeightsOk returns a tuple with the BillingWeights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040PartitionInfoTres) GetBillingWeightsOk() (*string, bool) {
	if o == nil || IsNil(o.BillingWeights) {
		return nil, false
	}
	return o.BillingWeights, true
}

// HasBillingWeights returns a boolean if a field has been set.
func (o *V0040PartitionInfoTres) HasBillingWeights() bool {
	if o != nil && !IsNil(o.BillingWeights) {
		return true
	}

	return false
}

// SetBillingWeights gets a reference to the given string and assigns it to the BillingWeights field.
func (o *V0040PartitionInfoTres) SetBillingWeights(v string) {
	o.BillingWeights = &v
}

// GetConfigured returns the Configured field value if set, zero value otherwise.
func (o *V0040PartitionInfoTres) GetConfigured() string {
	if o == nil || IsNil(o.Configured) {
		var ret string
		return ret
	}
	return *o.Configured
}

// GetConfiguredOk returns a tuple with the Configured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040PartitionInfoTres) GetConfiguredOk() (*string, bool) {
	if o == nil || IsNil(o.Configured) {
		return nil, false
	}
	return o.Configured, true
}

// HasConfigured returns a boolean if a field has been set.
func (o *V0040PartitionInfoTres) HasConfigured() bool {
	if o != nil && !IsNil(o.Configured) {
		return true
	}

	return false
}

// SetConfigured gets a reference to the given string and assigns it to the Configured field.
func (o *V0040PartitionInfoTres) SetConfigured(v string) {
	o.Configured = &v
}

func (o V0040PartitionInfoTres) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040PartitionInfoTres) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BillingWeights) {
		toSerialize["billing_weights"] = o.BillingWeights
	}
	if !IsNil(o.Configured) {
		toSerialize["configured"] = o.Configured
	}
	return toSerialize, nil
}

type NullableV0040PartitionInfoTres struct {
	value *V0040PartitionInfoTres
	isSet bool
}

func (v NullableV0040PartitionInfoTres) Get() *V0040PartitionInfoTres {
	return v.value
}

func (v *NullableV0040PartitionInfoTres) Set(val *V0040PartitionInfoTres) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040PartitionInfoTres) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040PartitionInfoTres) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040PartitionInfoTres(val *V0040PartitionInfoTres) *NullableV0040PartitionInfoTres {
	return &NullableV0040PartitionInfoTres{value: val, isSet: true}
}

func (v NullableV0040PartitionInfoTres) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040PartitionInfoTres) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

