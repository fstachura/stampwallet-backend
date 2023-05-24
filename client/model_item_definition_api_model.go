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
	"time"
)

// checks if the ItemDefinitionAPIModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemDefinitionAPIModel{}

// ItemDefinitionAPIModel struct for ItemDefinitionAPIModel
type ItemDefinitionAPIModel struct {
	PublicId    *string      `json:"publicId,omitempty"`
	Name        *string      `json:"name,omitempty"`
	Price       *int32       `json:"price,omitempty"`
	Description *string      `json:"description,omitempty"`
	ImageId     *string      `json:"imageId,omitempty"`
	StartDate   NullableTime `json:"startDate,omitempty"`
	EndDate     NullableTime `json:"endDate,omitempty"`
	MaxAmount   *int32       `json:"maxAmount,omitempty"`
	Available   *bool        `json:"available,omitempty"`
}

// NewItemDefinitionAPIModel instantiates a new ItemDefinitionAPIModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemDefinitionAPIModel() *ItemDefinitionAPIModel {
	this := ItemDefinitionAPIModel{}
	return &this
}

// NewItemDefinitionAPIModelWithDefaults instantiates a new ItemDefinitionAPIModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemDefinitionAPIModelWithDefaults() *ItemDefinitionAPIModel {
	this := ItemDefinitionAPIModel{}
	return &this
}

// GetPublicId returns the PublicId field value if set, zero value otherwise.
func (o *ItemDefinitionAPIModel) GetPublicId() string {
	if o == nil || isNil(o.PublicId) {
		var ret string
		return ret
	}
	return *o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemDefinitionAPIModel) GetPublicIdOk() (*string, bool) {
	if o == nil || isNil(o.PublicId) {
		return nil, false
	}
	return o.PublicId, true
}

// HasPublicId returns a boolean if a field has been set.
func (o *ItemDefinitionAPIModel) HasPublicId() bool {
	if o != nil && !isNil(o.PublicId) {
		return true
	}

	return false
}

// SetPublicId gets a reference to the given string and assigns it to the PublicId field.
func (o *ItemDefinitionAPIModel) SetPublicId(v string) {
	o.PublicId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ItemDefinitionAPIModel) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemDefinitionAPIModel) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ItemDefinitionAPIModel) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ItemDefinitionAPIModel) SetName(v string) {
	o.Name = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *ItemDefinitionAPIModel) GetPrice() int32 {
	if o == nil || isNil(o.Price) {
		var ret int32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemDefinitionAPIModel) GetPriceOk() (*int32, bool) {
	if o == nil || isNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *ItemDefinitionAPIModel) HasPrice() bool {
	if o != nil && !isNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given int32 and assigns it to the Price field.
func (o *ItemDefinitionAPIModel) SetPrice(v int32) {
	o.Price = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ItemDefinitionAPIModel) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemDefinitionAPIModel) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ItemDefinitionAPIModel) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ItemDefinitionAPIModel) SetDescription(v string) {
	o.Description = &v
}

// GetImageId returns the ImageId field value if set, zero value otherwise.
func (o *ItemDefinitionAPIModel) GetImageId() string {
	if o == nil || isNil(o.ImageId) {
		var ret string
		return ret
	}
	return *o.ImageId
}

// GetImageIdOk returns a tuple with the ImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemDefinitionAPIModel) GetImageIdOk() (*string, bool) {
	if o == nil || isNil(o.ImageId) {
		return nil, false
	}
	return o.ImageId, true
}

// HasImageId returns a boolean if a field has been set.
func (o *ItemDefinitionAPIModel) HasImageId() bool {
	if o != nil && !isNil(o.ImageId) {
		return true
	}

	return false
}

// SetImageId gets a reference to the given string and assigns it to the ImageId field.
func (o *ItemDefinitionAPIModel) SetImageId(v string) {
	o.ImageId = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ItemDefinitionAPIModel) GetStartDate() time.Time {
	if o == nil || isNil(o.StartDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.StartDate.Get()
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemDefinitionAPIModel) GetStartDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartDate.Get(), o.StartDate.IsSet()
}

// HasStartDate returns a boolean if a field has been set.
func (o *ItemDefinitionAPIModel) HasStartDate() bool {
	if o != nil && o.StartDate.IsSet() {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given NullableTime and assigns it to the StartDate field.
func (o *ItemDefinitionAPIModel) SetStartDate(v time.Time) {
	o.StartDate.Set(&v)
}

// SetStartDateNil sets the value for StartDate to be an explicit nil
func (o *ItemDefinitionAPIModel) SetStartDateNil() {
	o.StartDate.Set(nil)
}

// UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
func (o *ItemDefinitionAPIModel) UnsetStartDate() {
	o.StartDate.Unset()
}

// GetEndDate returns the EndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ItemDefinitionAPIModel) GetEndDate() time.Time {
	if o == nil || isNil(o.EndDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemDefinitionAPIModel) GetEndDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// HasEndDate returns a boolean if a field has been set.
func (o *ItemDefinitionAPIModel) HasEndDate() bool {
	if o != nil && o.EndDate.IsSet() {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given NullableTime and assigns it to the EndDate field.
func (o *ItemDefinitionAPIModel) SetEndDate(v time.Time) {
	o.EndDate.Set(&v)
}

// SetEndDateNil sets the value for EndDate to be an explicit nil
func (o *ItemDefinitionAPIModel) SetEndDateNil() {
	o.EndDate.Set(nil)
}

// UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
func (o *ItemDefinitionAPIModel) UnsetEndDate() {
	o.EndDate.Unset()
}

// GetMaxAmount returns the MaxAmount field value if set, zero value otherwise.
func (o *ItemDefinitionAPIModel) GetMaxAmount() int32 {
	if o == nil || isNil(o.MaxAmount) {
		var ret int32
		return ret
	}
	return *o.MaxAmount
}

// GetMaxAmountOk returns a tuple with the MaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemDefinitionAPIModel) GetMaxAmountOk() (*int32, bool) {
	if o == nil || isNil(o.MaxAmount) {
		return nil, false
	}
	return o.MaxAmount, true
}

// HasMaxAmount returns a boolean if a field has been set.
func (o *ItemDefinitionAPIModel) HasMaxAmount() bool {
	if o != nil && !isNil(o.MaxAmount) {
		return true
	}

	return false
}

// SetMaxAmount gets a reference to the given int32 and assigns it to the MaxAmount field.
func (o *ItemDefinitionAPIModel) SetMaxAmount(v int32) {
	o.MaxAmount = &v
}

// GetAvailable returns the Available field value if set, zero value otherwise.
func (o *ItemDefinitionAPIModel) GetAvailable() bool {
	if o == nil || isNil(o.Available) {
		var ret bool
		return ret
	}
	return *o.Available
}

// GetAvailableOk returns a tuple with the Available field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemDefinitionAPIModel) GetAvailableOk() (*bool, bool) {
	if o == nil || isNil(o.Available) {
		return nil, false
	}
	return o.Available, true
}

// HasAvailable returns a boolean if a field has been set.
func (o *ItemDefinitionAPIModel) HasAvailable() bool {
	if o != nil && !isNil(o.Available) {
		return true
	}

	return false
}

// SetAvailable gets a reference to the given bool and assigns it to the Available field.
func (o *ItemDefinitionAPIModel) SetAvailable(v bool) {
	o.Available = &v
}

func (o ItemDefinitionAPIModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ItemDefinitionAPIModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PublicId) {
		toSerialize["publicId"] = o.PublicId
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.ImageId) {
		toSerialize["imageId"] = o.ImageId
	}
	if o.StartDate.IsSet() {
		toSerialize["startDate"] = o.StartDate.Get()
	}
	if o.EndDate.IsSet() {
		toSerialize["endDate"] = o.EndDate.Get()
	}
	if !isNil(o.MaxAmount) {
		toSerialize["maxAmount"] = o.MaxAmount
	}
	if !isNil(o.Available) {
		toSerialize["available"] = o.Available
	}
	return toSerialize, nil
}

type NullableItemDefinitionAPIModel struct {
	value *ItemDefinitionAPIModel
	isSet bool
}

func (v NullableItemDefinitionAPIModel) Get() *ItemDefinitionAPIModel {
	return v.value
}

func (v *NullableItemDefinitionAPIModel) Set(val *ItemDefinitionAPIModel) {
	v.value = val
	v.isSet = true
}

func (v NullableItemDefinitionAPIModel) IsSet() bool {
	return v.isSet
}

func (v *NullableItemDefinitionAPIModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemDefinitionAPIModel(val *ItemDefinitionAPIModel) *NullableItemDefinitionAPIModel {
	return &NullableItemDefinitionAPIModel{value: val, isSet: true}
}

func (v NullableItemDefinitionAPIModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemDefinitionAPIModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
