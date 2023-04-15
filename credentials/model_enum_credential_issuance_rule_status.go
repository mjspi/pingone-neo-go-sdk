/*
PingOne Platform API - Credentials

The PingOne Platform API covering the PingOne Credentials service

API version: 2023-03-30
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package credentials

import (
	"encoding/json"
	"fmt"
)

// EnumCredentialIssuanceRuleStatus Specifies the ACTIVE or DISABLED status of the credential issuance rule.
type EnumCredentialIssuanceRuleStatus string

// List of EnumCredentialIssuanceRuleStatus
const (
	ENUMCREDENTIALISSUANCERULESTATUS_ACTIVE EnumCredentialIssuanceRuleStatus = "ACTIVE"
	ENUMCREDENTIALISSUANCERULESTATUS_DISABLED EnumCredentialIssuanceRuleStatus = "DISABLED"
)

// All allowed values of EnumCredentialIssuanceRuleStatus enum
var AllowedEnumCredentialIssuanceRuleStatusEnumValues = []EnumCredentialIssuanceRuleStatus{
	"ACTIVE",
	"DISABLED",
}

func (v *EnumCredentialIssuanceRuleStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumCredentialIssuanceRuleStatus(value)
	for _, existing := range AllowedEnumCredentialIssuanceRuleStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumCredentialIssuanceRuleStatus", value)
}

// NewEnumCredentialIssuanceRuleStatusFromValue returns a pointer to a valid EnumCredentialIssuanceRuleStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumCredentialIssuanceRuleStatusFromValue(v string) (*EnumCredentialIssuanceRuleStatus, error) {
	ev := EnumCredentialIssuanceRuleStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumCredentialIssuanceRuleStatus: valid values are %v", v, AllowedEnumCredentialIssuanceRuleStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumCredentialIssuanceRuleStatus) IsValid() bool {
	for _, existing := range AllowedEnumCredentialIssuanceRuleStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumCredentialIssuanceRuleStatus value
func (v EnumCredentialIssuanceRuleStatus) Ptr() *EnumCredentialIssuanceRuleStatus {
	return &v
}

type NullableEnumCredentialIssuanceRuleStatus struct {
	value *EnumCredentialIssuanceRuleStatus
	isSet bool
}

func (v NullableEnumCredentialIssuanceRuleStatus) Get() *EnumCredentialIssuanceRuleStatus {
	return v.value
}

func (v *NullableEnumCredentialIssuanceRuleStatus) Set(val *EnumCredentialIssuanceRuleStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumCredentialIssuanceRuleStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumCredentialIssuanceRuleStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumCredentialIssuanceRuleStatus(val *EnumCredentialIssuanceRuleStatus) *NullableEnumCredentialIssuanceRuleStatus {
	return &NullableEnumCredentialIssuanceRuleStatus{value: val, isSet: true}
}

func (v NullableEnumCredentialIssuanceRuleStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumCredentialIssuanceRuleStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

