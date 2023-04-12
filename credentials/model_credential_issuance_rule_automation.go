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

// checks if the CredentialIssuanceRuleAutomation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CredentialIssuanceRuleAutomation{}

// CredentialIssuanceRuleAutomation struct for CredentialIssuanceRuleAutomation
type CredentialIssuanceRuleAutomation struct {
	Issue EnumCredentialIssuanceRuleAutomationMethod `json:"issue"`
	Update EnumCredentialIssuanceRuleAutomationMethod `json:"update"`
	Revoke EnumCredentialIssuanceRuleAutomationMethod `json:"revoke"`
}

// NewCredentialIssuanceRuleAutomation instantiates a new CredentialIssuanceRuleAutomation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialIssuanceRuleAutomation(issue EnumCredentialIssuanceRuleAutomationMethod, update EnumCredentialIssuanceRuleAutomationMethod, revoke EnumCredentialIssuanceRuleAutomationMethod) *CredentialIssuanceRuleAutomation {
	this := CredentialIssuanceRuleAutomation{}
	this.Issue = issue
	this.Update = update
	this.Revoke = revoke
	return &this
}

// NewCredentialIssuanceRuleAutomationWithDefaults instantiates a new CredentialIssuanceRuleAutomation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialIssuanceRuleAutomationWithDefaults() *CredentialIssuanceRuleAutomation {
	this := CredentialIssuanceRuleAutomation{}
	return &this
}

// GetIssue returns the Issue field value
func (o *CredentialIssuanceRuleAutomation) GetIssue() EnumCredentialIssuanceRuleAutomationMethod {
	if o == nil {
		var ret EnumCredentialIssuanceRuleAutomationMethod
		return ret
	}

	return o.Issue
}

// GetIssueOk returns a tuple with the Issue field value
// and a boolean to check if the value has been set.
func (o *CredentialIssuanceRuleAutomation) GetIssueOk() (*EnumCredentialIssuanceRuleAutomationMethod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Issue, true
}

// SetIssue sets field value
func (o *CredentialIssuanceRuleAutomation) SetIssue(v EnumCredentialIssuanceRuleAutomationMethod) {
	o.Issue = v
}

// GetUpdate returns the Update field value
func (o *CredentialIssuanceRuleAutomation) GetUpdate() EnumCredentialIssuanceRuleAutomationMethod {
	if o == nil {
		var ret EnumCredentialIssuanceRuleAutomationMethod
		return ret
	}

	return o.Update
}

// GetUpdateOk returns a tuple with the Update field value
// and a boolean to check if the value has been set.
func (o *CredentialIssuanceRuleAutomation) GetUpdateOk() (*EnumCredentialIssuanceRuleAutomationMethod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Update, true
}

// SetUpdate sets field value
func (o *CredentialIssuanceRuleAutomation) SetUpdate(v EnumCredentialIssuanceRuleAutomationMethod) {
	o.Update = v
}

// GetRevoke returns the Revoke field value
func (o *CredentialIssuanceRuleAutomation) GetRevoke() EnumCredentialIssuanceRuleAutomationMethod {
	if o == nil {
		var ret EnumCredentialIssuanceRuleAutomationMethod
		return ret
	}

	return o.Revoke
}

// GetRevokeOk returns a tuple with the Revoke field value
// and a boolean to check if the value has been set.
func (o *CredentialIssuanceRuleAutomation) GetRevokeOk() (*EnumCredentialIssuanceRuleAutomationMethod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Revoke, true
}

// SetRevoke sets field value
func (o *CredentialIssuanceRuleAutomation) SetRevoke(v EnumCredentialIssuanceRuleAutomationMethod) {
	o.Revoke = v
}

func (o CredentialIssuanceRuleAutomation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CredentialIssuanceRuleAutomation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["issue"] = o.Issue
	toSerialize["update"] = o.Update
	toSerialize["revoke"] = o.Revoke
	return toSerialize, nil
}

type NullableCredentialIssuanceRuleAutomation struct {
	value *CredentialIssuanceRuleAutomation
	isSet bool
}

func (v NullableCredentialIssuanceRuleAutomation) Get() *CredentialIssuanceRuleAutomation {
	return v.value
}

func (v *NullableCredentialIssuanceRuleAutomation) Set(val *CredentialIssuanceRuleAutomation) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialIssuanceRuleAutomation) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialIssuanceRuleAutomation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialIssuanceRuleAutomation(val *CredentialIssuanceRuleAutomation) *NullableCredentialIssuanceRuleAutomation {
	return &NullableCredentialIssuanceRuleAutomation{value: val, isSet: true}
}

func (v NullableCredentialIssuanceRuleAutomation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialIssuanceRuleAutomation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

