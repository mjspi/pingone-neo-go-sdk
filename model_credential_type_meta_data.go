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

// checks if the CredentialTypeMetaData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CredentialTypeMetaData{}

// CredentialTypeMetaData struct for CredentialTypeMetaData
type CredentialTypeMetaData struct {
	// A string that specifies the name of the credential.
	Name *string `json:"name,omitempty"`
	// A string that specifies the description of the credential.
	Description *string `json:"description,omitempty"`
	// A string that specifies the color of the text to show on the credential.
	TextColor *string `json:"textColor,omitempty"`
	// A string that specifies the color to show on the credential.
	CardColor *string `json:"cardColor,omitempty"`
	// A string that specifies the URL to an image of the background to show in the credential.
	BackgroundImage *string `json:"backgroundImage,omitempty"`
	// A string that specifies the URL to an image of the logo to show in the credential.
	LogoImage *string `json:"logoImage,omitempty"`
	// A string that specifies the percent opacity of the background image in the credential. High percentage opacity may make reading text difficult.
	BgOpacityPercent *int32 `json:"bgOpacityPercent,omitempty"`
	// An array of objects that specifies the fields on the credential.
	Fields []CredentialTypeMetaDataFieldsInner `json:"fields,omitempty"`
	// An integer that specifies theersion of this credential. If not provided, the service assigns a version.
	Version *int32 `json:"version,omitempty"`
}

// NewCredentialTypeMetaData instantiates a new CredentialTypeMetaData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialTypeMetaData() *CredentialTypeMetaData {
	this := CredentialTypeMetaData{}
	return &this
}

// NewCredentialTypeMetaDataWithDefaults instantiates a new CredentialTypeMetaData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialTypeMetaDataWithDefaults() *CredentialTypeMetaData {
	this := CredentialTypeMetaData{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CredentialTypeMetaData) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialTypeMetaData) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CredentialTypeMetaData) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CredentialTypeMetaData) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CredentialTypeMetaData) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialTypeMetaData) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CredentialTypeMetaData) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CredentialTypeMetaData) SetDescription(v string) {
	o.Description = &v
}

// GetTextColor returns the TextColor field value if set, zero value otherwise.
func (o *CredentialTypeMetaData) GetTextColor() string {
	if o == nil || IsNil(o.TextColor) {
		var ret string
		return ret
	}
	return *o.TextColor
}

// GetTextColorOk returns a tuple with the TextColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialTypeMetaData) GetTextColorOk() (*string, bool) {
	if o == nil || IsNil(o.TextColor) {
		return nil, false
	}
	return o.TextColor, true
}

// HasTextColor returns a boolean if a field has been set.
func (o *CredentialTypeMetaData) HasTextColor() bool {
	if o != nil && !IsNil(o.TextColor) {
		return true
	}

	return false
}

// SetTextColor gets a reference to the given string and assigns it to the TextColor field.
func (o *CredentialTypeMetaData) SetTextColor(v string) {
	o.TextColor = &v
}

// GetCardColor returns the CardColor field value if set, zero value otherwise.
func (o *CredentialTypeMetaData) GetCardColor() string {
	if o == nil || IsNil(o.CardColor) {
		var ret string
		return ret
	}
	return *o.CardColor
}

// GetCardColorOk returns a tuple with the CardColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialTypeMetaData) GetCardColorOk() (*string, bool) {
	if o == nil || IsNil(o.CardColor) {
		return nil, false
	}
	return o.CardColor, true
}

// HasCardColor returns a boolean if a field has been set.
func (o *CredentialTypeMetaData) HasCardColor() bool {
	if o != nil && !IsNil(o.CardColor) {
		return true
	}

	return false
}

// SetCardColor gets a reference to the given string and assigns it to the CardColor field.
func (o *CredentialTypeMetaData) SetCardColor(v string) {
	o.CardColor = &v
}

// GetBackgroundImage returns the BackgroundImage field value if set, zero value otherwise.
func (o *CredentialTypeMetaData) GetBackgroundImage() string {
	if o == nil || IsNil(o.BackgroundImage) {
		var ret string
		return ret
	}
	return *o.BackgroundImage
}

// GetBackgroundImageOk returns a tuple with the BackgroundImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialTypeMetaData) GetBackgroundImageOk() (*string, bool) {
	if o == nil || IsNil(o.BackgroundImage) {
		return nil, false
	}
	return o.BackgroundImage, true
}

// HasBackgroundImage returns a boolean if a field has been set.
func (o *CredentialTypeMetaData) HasBackgroundImage() bool {
	if o != nil && !IsNil(o.BackgroundImage) {
		return true
	}

	return false
}

// SetBackgroundImage gets a reference to the given string and assigns it to the BackgroundImage field.
func (o *CredentialTypeMetaData) SetBackgroundImage(v string) {
	o.BackgroundImage = &v
}

// GetLogoImage returns the LogoImage field value if set, zero value otherwise.
func (o *CredentialTypeMetaData) GetLogoImage() string {
	if o == nil || IsNil(o.LogoImage) {
		var ret string
		return ret
	}
	return *o.LogoImage
}

// GetLogoImageOk returns a tuple with the LogoImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialTypeMetaData) GetLogoImageOk() (*string, bool) {
	if o == nil || IsNil(o.LogoImage) {
		return nil, false
	}
	return o.LogoImage, true
}

// HasLogoImage returns a boolean if a field has been set.
func (o *CredentialTypeMetaData) HasLogoImage() bool {
	if o != nil && !IsNil(o.LogoImage) {
		return true
	}

	return false
}

// SetLogoImage gets a reference to the given string and assigns it to the LogoImage field.
func (o *CredentialTypeMetaData) SetLogoImage(v string) {
	o.LogoImage = &v
}

// GetBgOpacityPercent returns the BgOpacityPercent field value if set, zero value otherwise.
func (o *CredentialTypeMetaData) GetBgOpacityPercent() int32 {
	if o == nil || IsNil(o.BgOpacityPercent) {
		var ret int32
		return ret
	}
	return *o.BgOpacityPercent
}

// GetBgOpacityPercentOk returns a tuple with the BgOpacityPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialTypeMetaData) GetBgOpacityPercentOk() (*int32, bool) {
	if o == nil || IsNil(o.BgOpacityPercent) {
		return nil, false
	}
	return o.BgOpacityPercent, true
}

// HasBgOpacityPercent returns a boolean if a field has been set.
func (o *CredentialTypeMetaData) HasBgOpacityPercent() bool {
	if o != nil && !IsNil(o.BgOpacityPercent) {
		return true
	}

	return false
}

// SetBgOpacityPercent gets a reference to the given int32 and assigns it to the BgOpacityPercent field.
func (o *CredentialTypeMetaData) SetBgOpacityPercent(v int32) {
	o.BgOpacityPercent = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *CredentialTypeMetaData) GetFields() []CredentialTypeMetaDataFieldsInner {
	if o == nil || IsNil(o.Fields) {
		var ret []CredentialTypeMetaDataFieldsInner
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialTypeMetaData) GetFieldsOk() ([]CredentialTypeMetaDataFieldsInner, bool) {
	if o == nil || IsNil(o.Fields) {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *CredentialTypeMetaData) HasFields() bool {
	if o != nil && !IsNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given []CredentialTypeMetaDataFieldsInner and assigns it to the Fields field.
func (o *CredentialTypeMetaData) SetFields(v []CredentialTypeMetaDataFieldsInner) {
	o.Fields = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *CredentialTypeMetaData) GetVersion() int32 {
	if o == nil || IsNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialTypeMetaData) GetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *CredentialTypeMetaData) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *CredentialTypeMetaData) SetVersion(v int32) {
	o.Version = &v
}

func (o CredentialTypeMetaData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CredentialTypeMetaData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.TextColor) {
		toSerialize["textColor"] = o.TextColor
	}
	if !IsNil(o.CardColor) {
		toSerialize["cardColor"] = o.CardColor
	}
	if !IsNil(o.BackgroundImage) {
		toSerialize["backgroundImage"] = o.BackgroundImage
	}
	if !IsNil(o.LogoImage) {
		toSerialize["logoImage"] = o.LogoImage
	}
	if !IsNil(o.BgOpacityPercent) {
		toSerialize["bgOpacityPercent"] = o.BgOpacityPercent
	}
	if !IsNil(o.Fields) {
		toSerialize["fields"] = o.Fields
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableCredentialTypeMetaData struct {
	value *CredentialTypeMetaData
	isSet bool
}

func (v NullableCredentialTypeMetaData) Get() *CredentialTypeMetaData {
	return v.value
}

func (v *NullableCredentialTypeMetaData) Set(val *CredentialTypeMetaData) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialTypeMetaData) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialTypeMetaData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialTypeMetaData(val *CredentialTypeMetaData) *NullableCredentialTypeMetaData {
	return &NullableCredentialTypeMetaData{value: val, isSet: true}
}

func (v NullableCredentialTypeMetaData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialTypeMetaData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

