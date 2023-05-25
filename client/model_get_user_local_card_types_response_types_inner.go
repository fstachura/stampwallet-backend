/*
StampWallet API Server

StampWallet API Server REST Specification

API version: 0.1.0
Contact: fbstachura@gmail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the GetUserLocalCardTypesResponseTypesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetUserLocalCardTypesResponseTypesInner{}

// GetUserLocalCardTypesResponseTypesInner struct for GetUserLocalCardTypesResponseTypesInner
type GetUserLocalCardTypesResponseTypesInner struct {
	PublicId *string `json:"publicId,omitempty"`
	Name     *string `json:"name,omitempty"`
	Code     *string `json:"code,omitempty"`
}

// NewGetUserLocalCardTypesResponseTypesInner instantiates a new GetUserLocalCardTypesResponseTypesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUserLocalCardTypesResponseTypesInner() *GetUserLocalCardTypesResponseTypesInner {
	this := GetUserLocalCardTypesResponseTypesInner{}
	return &this
}

// NewGetUserLocalCardTypesResponseTypesInnerWithDefaults instantiates a new GetUserLocalCardTypesResponseTypesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUserLocalCardTypesResponseTypesInnerWithDefaults() *GetUserLocalCardTypesResponseTypesInner {
	this := GetUserLocalCardTypesResponseTypesInner{}
	return &this
}

// GetPublicId returns the PublicId field value if set, zero value otherwise.
func (o *GetUserLocalCardTypesResponseTypesInner) GetPublicId() string {
	if o == nil || isNil(o.PublicId) {
		var ret string
		return ret
	}
	return *o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserLocalCardTypesResponseTypesInner) GetPublicIdOk() (*string, bool) {
	if o == nil || isNil(o.PublicId) {
		return nil, false
	}
	return o.PublicId, true
}

// HasPublicId returns a boolean if a field has been set.
func (o *GetUserLocalCardTypesResponseTypesInner) HasPublicId() bool {
	if o != nil && !isNil(o.PublicId) {
		return true
	}

	return false
}

// SetPublicId gets a reference to the given string and assigns it to the PublicId field.
func (o *GetUserLocalCardTypesResponseTypesInner) SetPublicId(v string) {
	o.PublicId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetUserLocalCardTypesResponseTypesInner) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserLocalCardTypesResponseTypesInner) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetUserLocalCardTypesResponseTypesInner) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetUserLocalCardTypesResponseTypesInner) SetName(v string) {
	o.Name = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *GetUserLocalCardTypesResponseTypesInner) GetCode() string {
	if o == nil || isNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserLocalCardTypesResponseTypesInner) GetCodeOk() (*string, bool) {
	if o == nil || isNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *GetUserLocalCardTypesResponseTypesInner) HasCode() bool {
	if o != nil && !isNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *GetUserLocalCardTypesResponseTypesInner) SetCode(v string) {
	o.Code = &v
}

func (o GetUserLocalCardTypesResponseTypesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUserLocalCardTypesResponseTypesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PublicId) {
		toSerialize["publicId"] = o.PublicId
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	return toSerialize, nil
}

type NullableGetUserLocalCardTypesResponseTypesInner struct {
	value *GetUserLocalCardTypesResponseTypesInner
	isSet bool
}

func (v NullableGetUserLocalCardTypesResponseTypesInner) Get() *GetUserLocalCardTypesResponseTypesInner {
	return v.value
}

func (v *NullableGetUserLocalCardTypesResponseTypesInner) Set(val *GetUserLocalCardTypesResponseTypesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUserLocalCardTypesResponseTypesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUserLocalCardTypesResponseTypesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUserLocalCardTypesResponseTypesInner(val *GetUserLocalCardTypesResponseTypesInner) *NullableGetUserLocalCardTypesResponseTypesInner {
	return &NullableGetUserLocalCardTypesResponseTypesInner{value: val, isSet: true}
}

func (v NullableGetUserLocalCardTypesResponseTypesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUserLocalCardTypesResponseTypesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
