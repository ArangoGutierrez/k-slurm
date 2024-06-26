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

// checks if the Dbv0038ConfigInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0038ConfigInfo{}

// Dbv0038ConfigInfo struct for Dbv0038ConfigInfo
type Dbv0038ConfigInfo struct {
	Meta *Dbv0038Meta `json:"meta,omitempty"`
	// Slurm errors
	Errors []Dbv0038Error `json:"errors,omitempty"`
	// Array of TRES
	Tres [][]Dbv0038TresListInner `json:"tres,omitempty"`
	// Array of accounts
	Accounts []Dbv0038Account `json:"accounts,omitempty"`
	// Array of associations
	Associations []Dbv0038Association `json:"associations,omitempty"`
	// Array of users
	Users []Dbv0038User `json:"users,omitempty"`
	// Array of qos
	Qos []Dbv0038Qos `json:"qos,omitempty"`
	// Array of wckeys
	Wckeys []Dbv0038Wckey `json:"wckeys,omitempty"`
}

// NewDbv0038ConfigInfo instantiates a new Dbv0038ConfigInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0038ConfigInfo() *Dbv0038ConfigInfo {
	this := Dbv0038ConfigInfo{}
	return &this
}

// NewDbv0038ConfigInfoWithDefaults instantiates a new Dbv0038ConfigInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0038ConfigInfoWithDefaults() *Dbv0038ConfigInfo {
	this := Dbv0038ConfigInfo{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *Dbv0038ConfigInfo) GetMeta() Dbv0038Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Dbv0038Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038ConfigInfo) GetMetaOk() (*Dbv0038Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *Dbv0038ConfigInfo) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Dbv0038Meta and assigns it to the Meta field.
func (o *Dbv0038ConfigInfo) SetMeta(v Dbv0038Meta) {
	o.Meta = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *Dbv0038ConfigInfo) GetErrors() []Dbv0038Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Dbv0038Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038ConfigInfo) GetErrorsOk() ([]Dbv0038Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *Dbv0038ConfigInfo) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Dbv0038Error and assigns it to the Errors field.
func (o *Dbv0038ConfigInfo) SetErrors(v []Dbv0038Error) {
	o.Errors = v
}

// GetTres returns the Tres field value if set, zero value otherwise.
func (o *Dbv0038ConfigInfo) GetTres() [][]Dbv0038TresListInner {
	if o == nil || IsNil(o.Tres) {
		var ret [][]Dbv0038TresListInner
		return ret
	}
	return o.Tres
}

// GetTresOk returns a tuple with the Tres field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038ConfigInfo) GetTresOk() ([][]Dbv0038TresListInner, bool) {
	if o == nil || IsNil(o.Tres) {
		return nil, false
	}
	return o.Tres, true
}

// HasTres returns a boolean if a field has been set.
func (o *Dbv0038ConfigInfo) HasTres() bool {
	if o != nil && !IsNil(o.Tres) {
		return true
	}

	return false
}

// SetTres gets a reference to the given [][]Dbv0038TresListInner and assigns it to the Tres field.
func (o *Dbv0038ConfigInfo) SetTres(v [][]Dbv0038TresListInner) {
	o.Tres = v
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *Dbv0038ConfigInfo) GetAccounts() []Dbv0038Account {
	if o == nil || IsNil(o.Accounts) {
		var ret []Dbv0038Account
		return ret
	}
	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038ConfigInfo) GetAccountsOk() ([]Dbv0038Account, bool) {
	if o == nil || IsNil(o.Accounts) {
		return nil, false
	}
	return o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *Dbv0038ConfigInfo) HasAccounts() bool {
	if o != nil && !IsNil(o.Accounts) {
		return true
	}

	return false
}

// SetAccounts gets a reference to the given []Dbv0038Account and assigns it to the Accounts field.
func (o *Dbv0038ConfigInfo) SetAccounts(v []Dbv0038Account) {
	o.Accounts = v
}

// GetAssociations returns the Associations field value if set, zero value otherwise.
func (o *Dbv0038ConfigInfo) GetAssociations() []Dbv0038Association {
	if o == nil || IsNil(o.Associations) {
		var ret []Dbv0038Association
		return ret
	}
	return o.Associations
}

// GetAssociationsOk returns a tuple with the Associations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038ConfigInfo) GetAssociationsOk() ([]Dbv0038Association, bool) {
	if o == nil || IsNil(o.Associations) {
		return nil, false
	}
	return o.Associations, true
}

// HasAssociations returns a boolean if a field has been set.
func (o *Dbv0038ConfigInfo) HasAssociations() bool {
	if o != nil && !IsNil(o.Associations) {
		return true
	}

	return false
}

// SetAssociations gets a reference to the given []Dbv0038Association and assigns it to the Associations field.
func (o *Dbv0038ConfigInfo) SetAssociations(v []Dbv0038Association) {
	o.Associations = v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *Dbv0038ConfigInfo) GetUsers() []Dbv0038User {
	if o == nil || IsNil(o.Users) {
		var ret []Dbv0038User
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038ConfigInfo) GetUsersOk() ([]Dbv0038User, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *Dbv0038ConfigInfo) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []Dbv0038User and assigns it to the Users field.
func (o *Dbv0038ConfigInfo) SetUsers(v []Dbv0038User) {
	o.Users = v
}

// GetQos returns the Qos field value if set, zero value otherwise.
func (o *Dbv0038ConfigInfo) GetQos() []Dbv0038Qos {
	if o == nil || IsNil(o.Qos) {
		var ret []Dbv0038Qos
		return ret
	}
	return o.Qos
}

// GetQosOk returns a tuple with the Qos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038ConfigInfo) GetQosOk() ([]Dbv0038Qos, bool) {
	if o == nil || IsNil(o.Qos) {
		return nil, false
	}
	return o.Qos, true
}

// HasQos returns a boolean if a field has been set.
func (o *Dbv0038ConfigInfo) HasQos() bool {
	if o != nil && !IsNil(o.Qos) {
		return true
	}

	return false
}

// SetQos gets a reference to the given []Dbv0038Qos and assigns it to the Qos field.
func (o *Dbv0038ConfigInfo) SetQos(v []Dbv0038Qos) {
	o.Qos = v
}

// GetWckeys returns the Wckeys field value if set, zero value otherwise.
func (o *Dbv0038ConfigInfo) GetWckeys() []Dbv0038Wckey {
	if o == nil || IsNil(o.Wckeys) {
		var ret []Dbv0038Wckey
		return ret
	}
	return o.Wckeys
}

// GetWckeysOk returns a tuple with the Wckeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038ConfigInfo) GetWckeysOk() ([]Dbv0038Wckey, bool) {
	if o == nil || IsNil(o.Wckeys) {
		return nil, false
	}
	return o.Wckeys, true
}

// HasWckeys returns a boolean if a field has been set.
func (o *Dbv0038ConfigInfo) HasWckeys() bool {
	if o != nil && !IsNil(o.Wckeys) {
		return true
	}

	return false
}

// SetWckeys gets a reference to the given []Dbv0038Wckey and assigns it to the Wckeys field.
func (o *Dbv0038ConfigInfo) SetWckeys(v []Dbv0038Wckey) {
	o.Wckeys = v
}

func (o Dbv0038ConfigInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0038ConfigInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.Tres) {
		toSerialize["tres"] = o.Tres
	}
	if !IsNil(o.Accounts) {
		toSerialize["accounts"] = o.Accounts
	}
	if !IsNil(o.Associations) {
		toSerialize["associations"] = o.Associations
	}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	if !IsNil(o.Qos) {
		toSerialize["qos"] = o.Qos
	}
	if !IsNil(o.Wckeys) {
		toSerialize["wckeys"] = o.Wckeys
	}
	return toSerialize, nil
}

type NullableDbv0038ConfigInfo struct {
	value *Dbv0038ConfigInfo
	isSet bool
}

func (v NullableDbv0038ConfigInfo) Get() *Dbv0038ConfigInfo {
	return v.value
}

func (v *NullableDbv0038ConfigInfo) Set(val *Dbv0038ConfigInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0038ConfigInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0038ConfigInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0038ConfigInfo(val *Dbv0038ConfigInfo) *NullableDbv0038ConfigInfo {
	return &NullableDbv0038ConfigInfo{value: val, isSet: true}
}

func (v NullableDbv0038ConfigInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0038ConfigInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


