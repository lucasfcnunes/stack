/*
Membership API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package membershipclient

import (
	"encoding/json"
)

// checks if the BillingPortalResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingPortalResponse{}

// BillingPortalResponse struct for BillingPortalResponse
type BillingPortalResponse struct {
	Data *BillingPortal `json:"data,omitempty"`
}

// NewBillingPortalResponse instantiates a new BillingPortalResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingPortalResponse() *BillingPortalResponse {
	this := BillingPortalResponse{}
	return &this
}

// NewBillingPortalResponseWithDefaults instantiates a new BillingPortalResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingPortalResponseWithDefaults() *BillingPortalResponse {
	this := BillingPortalResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *BillingPortalResponse) GetData() BillingPortal {
	if o == nil || IsNil(o.Data) {
		var ret BillingPortal
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPortalResponse) GetDataOk() (*BillingPortal, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *BillingPortalResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given BillingPortal and assigns it to the Data field.
func (o *BillingPortalResponse) SetData(v BillingPortal) {
	o.Data = &v
}

func (o BillingPortalResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingPortalResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableBillingPortalResponse struct {
	value *BillingPortalResponse
	isSet bool
}

func (v NullableBillingPortalResponse) Get() *BillingPortalResponse {
	return v.value
}

func (v *NullableBillingPortalResponse) Set(val *BillingPortalResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingPortalResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingPortalResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingPortalResponse(val *BillingPortalResponse) *NullableBillingPortalResponse {
	return &NullableBillingPortalResponse{value: val, isSet: true}
}

func (v NullableBillingPortalResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingPortalResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
