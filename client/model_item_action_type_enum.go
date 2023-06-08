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
	"fmt"
)

// ItemActionTypeEnum the model 'ItemActionTypeEnum'
type ItemActionTypeEnum string

// List of ItemActionTypeEnum
const (
	NO_ACTION ItemActionTypeEnum = "NO_ACTION"
	REDEEMED  ItemActionTypeEnum = "REDEEMED"
	RECALLED  ItemActionTypeEnum = "RECALLED"
	CANCELLED ItemActionTypeEnum = "CANCELLED"
)

// All allowed values of ItemActionTypeEnum enum
var AllowedItemActionTypeEnumEnumValues = []ItemActionTypeEnum{
	"NO_ACTION",
	"REDEEMED",
	"RECALLED",
	"CANCELLED",
}

func (v *ItemActionTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ItemActionTypeEnum(value)
	for _, existing := range AllowedItemActionTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ItemActionTypeEnum", value)
}

// NewItemActionTypeEnumFromValue returns a pointer to a valid ItemActionTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewItemActionTypeEnumFromValue(v string) (*ItemActionTypeEnum, error) {
	ev := ItemActionTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ItemActionTypeEnum: valid values are %v", v, AllowedItemActionTypeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ItemActionTypeEnum) IsValid() bool {
	for _, existing := range AllowedItemActionTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ItemActionTypeEnum value
func (v ItemActionTypeEnum) Ptr() *ItemActionTypeEnum {
	return &v
}

type NullableItemActionTypeEnum struct {
	value *ItemActionTypeEnum
	isSet bool
}

func (v NullableItemActionTypeEnum) Get() *ItemActionTypeEnum {
	return v.value
}

func (v *NullableItemActionTypeEnum) Set(val *ItemActionTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableItemActionTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableItemActionTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemActionTypeEnum(val *ItemActionTypeEnum) *NullableItemActionTypeEnum {
	return &NullableItemActionTypeEnum{value: val, isSet: true}
}

func (v NullableItemActionTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemActionTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
