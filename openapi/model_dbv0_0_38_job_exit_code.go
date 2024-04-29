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

// checks if the Dbv0038JobExitCode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0038JobExitCode{}

// Dbv0038JobExitCode struct for Dbv0038JobExitCode
type Dbv0038JobExitCode struct {
	// Job exit status
	Status *string `json:"status,omitempty"`
	// Return code from parent process
	ReturnCode *int32 `json:"return_code,omitempty"`
	Signal *Dbv0038JobExitCodeSignal `json:"signal,omitempty"`
}

// NewDbv0038JobExitCode instantiates a new Dbv0038JobExitCode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0038JobExitCode() *Dbv0038JobExitCode {
	this := Dbv0038JobExitCode{}
	return &this
}

// NewDbv0038JobExitCodeWithDefaults instantiates a new Dbv0038JobExitCode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0038JobExitCodeWithDefaults() *Dbv0038JobExitCode {
	this := Dbv0038JobExitCode{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Dbv0038JobExitCode) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038JobExitCode) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Dbv0038JobExitCode) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Dbv0038JobExitCode) SetStatus(v string) {
	o.Status = &v
}

// GetReturnCode returns the ReturnCode field value if set, zero value otherwise.
func (o *Dbv0038JobExitCode) GetReturnCode() int32 {
	if o == nil || IsNil(o.ReturnCode) {
		var ret int32
		return ret
	}
	return *o.ReturnCode
}

// GetReturnCodeOk returns a tuple with the ReturnCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038JobExitCode) GetReturnCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.ReturnCode) {
		return nil, false
	}
	return o.ReturnCode, true
}

// HasReturnCode returns a boolean if a field has been set.
func (o *Dbv0038JobExitCode) HasReturnCode() bool {
	if o != nil && !IsNil(o.ReturnCode) {
		return true
	}

	return false
}

// SetReturnCode gets a reference to the given int32 and assigns it to the ReturnCode field.
func (o *Dbv0038JobExitCode) SetReturnCode(v int32) {
	o.ReturnCode = &v
}

// GetSignal returns the Signal field value if set, zero value otherwise.
func (o *Dbv0038JobExitCode) GetSignal() Dbv0038JobExitCodeSignal {
	if o == nil || IsNil(o.Signal) {
		var ret Dbv0038JobExitCodeSignal
		return ret
	}
	return *o.Signal
}

// GetSignalOk returns a tuple with the Signal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038JobExitCode) GetSignalOk() (*Dbv0038JobExitCodeSignal, bool) {
	if o == nil || IsNil(o.Signal) {
		return nil, false
	}
	return o.Signal, true
}

// HasSignal returns a boolean if a field has been set.
func (o *Dbv0038JobExitCode) HasSignal() bool {
	if o != nil && !IsNil(o.Signal) {
		return true
	}

	return false
}

// SetSignal gets a reference to the given Dbv0038JobExitCodeSignal and assigns it to the Signal field.
func (o *Dbv0038JobExitCode) SetSignal(v Dbv0038JobExitCodeSignal) {
	o.Signal = &v
}

func (o Dbv0038JobExitCode) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0038JobExitCode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.ReturnCode) {
		toSerialize["return_code"] = o.ReturnCode
	}
	if !IsNil(o.Signal) {
		toSerialize["signal"] = o.Signal
	}
	return toSerialize, nil
}

type NullableDbv0038JobExitCode struct {
	value *Dbv0038JobExitCode
	isSet bool
}

func (v NullableDbv0038JobExitCode) Get() *Dbv0038JobExitCode {
	return v.value
}

func (v *NullableDbv0038JobExitCode) Set(val *Dbv0038JobExitCode) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0038JobExitCode) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0038JobExitCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0038JobExitCode(val *Dbv0038JobExitCode) *NullableDbv0038JobExitCode {
	return &NullableDbv0038JobExitCode{value: val, isSet: true}
}

func (v NullableDbv0038JobExitCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0038JobExitCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

