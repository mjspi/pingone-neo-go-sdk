/*
PingOne Platform API - Credentials

The PingOne Platform API covering the PingOne Credentials service

API version: 2023-03-30
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the CredentialDigitalWalletDigitalWalletApplication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CredentialDigitalWalletDigitalWalletApplication{}

// CredentialDigitalWalletDigitalWalletApplication struct for CredentialDigitalWalletDigitalWalletApplication
type CredentialDigitalWalletDigitalWalletApplication struct {
	// A string that specifies the identifier (UUID) of the customer's digital wallet app that will interact with the user's digital wallet.
	Id string `json:"id"`
}

// NewCredentialDigitalWalletDigitalWalletApplication instantiates a new CredentialDigitalWalletDigitalWalletApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialDigitalWalletDigitalWalletApplication(id string) *CredentialDigitalWalletDigitalWalletApplication {
	this := CredentialDigitalWalletDigitalWalletApplication{}
	this.Id = id
	return &this
}

// NewCredentialDigitalWalletDigitalWalletApplicationWithDefaults instantiates a new CredentialDigitalWalletDigitalWalletApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialDigitalWalletDigitalWalletApplicationWithDefaults() *CredentialDigitalWalletDigitalWalletApplication {
	this := CredentialDigitalWalletDigitalWalletApplication{}
	return &this
}

// GetId returns the Id field value
func (o *CredentialDigitalWalletDigitalWalletApplication) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CredentialDigitalWalletDigitalWalletApplication) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CredentialDigitalWalletDigitalWalletApplication) SetId(v string) {
	o.Id = v
}

func (o CredentialDigitalWalletDigitalWalletApplication) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CredentialDigitalWalletDigitalWalletApplication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableCredentialDigitalWalletDigitalWalletApplication struct {
	value *CredentialDigitalWalletDigitalWalletApplication
	isSet bool
}

func (v NullableCredentialDigitalWalletDigitalWalletApplication) Get() *CredentialDigitalWalletDigitalWalletApplication {
	return v.value
}

func (v *NullableCredentialDigitalWalletDigitalWalletApplication) Set(val *CredentialDigitalWalletDigitalWalletApplication) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialDigitalWalletDigitalWalletApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialDigitalWalletDigitalWalletApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialDigitalWalletDigitalWalletApplication(val *CredentialDigitalWalletDigitalWalletApplication) *NullableCredentialDigitalWalletDigitalWalletApplication {
	return &NullableCredentialDigitalWalletDigitalWalletApplication{value: val, isSet: true}
}

func (v NullableCredentialDigitalWalletDigitalWalletApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialDigitalWalletDigitalWalletApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


