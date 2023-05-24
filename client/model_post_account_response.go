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

// checks if the PostAccountResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostAccountResponse{}

// PostAccountResponse struct for PostAccountResponse
type PostAccountResponse struct {
	Token *string `json:"token,omitempty"`
}

// NewPostAccountResponse instantiates a new PostAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostAccountResponse() *PostAccountResponse {
	this := PostAccountResponse{}
	return &this
}

// NewPostAccountResponseWithDefaults instantiates a new PostAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostAccountResponseWithDefaults() *PostAccountResponse {
	this := PostAccountResponse{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *PostAccountResponse) GetToken() string {
	if o == nil || isNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAccountResponse) GetTokenOk() (*string, bool) {
	if o == nil || isNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *PostAccountResponse) HasToken() bool {
	if o != nil && !isNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *PostAccountResponse) SetToken(v string) {
	o.Token = &v
}

func (o PostAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostAccountResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	return toSerialize, nil
}

type NullablePostAccountResponse struct {
	value *PostAccountResponse
	isSet bool
}

func (v NullablePostAccountResponse) Get() *PostAccountResponse {
	return v.value
}

func (v *NullablePostAccountResponse) Set(val *PostAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePostAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePostAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostAccountResponse(val *PostAccountResponse) *NullablePostAccountResponse {
	return &NullablePostAccountResponse{value: val, isSet: true}
}

func (v NullablePostAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}