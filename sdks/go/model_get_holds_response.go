/*
Formance Stack API

Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 

API version: develop
Contact: support@formance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package formance

import (
	"encoding/json"
)

// GetHoldsResponse struct for GetHoldsResponse
type GetHoldsResponse struct {
	Cursor GetHoldsResponseCursor `json:"cursor"`
}

// NewGetHoldsResponse instantiates a new GetHoldsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetHoldsResponse(cursor GetHoldsResponseCursor) *GetHoldsResponse {
	this := GetHoldsResponse{}
	this.Cursor = cursor
	return &this
}

// NewGetHoldsResponseWithDefaults instantiates a new GetHoldsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetHoldsResponseWithDefaults() *GetHoldsResponse {
	this := GetHoldsResponse{}
	return &this
}

// GetCursor returns the Cursor field value
func (o *GetHoldsResponse) GetCursor() GetHoldsResponseCursor {
	if o == nil {
		var ret GetHoldsResponseCursor
		return ret
	}

	return o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value
// and a boolean to check if the value has been set.
func (o *GetHoldsResponse) GetCursorOk() (*GetHoldsResponseCursor, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Cursor, true
}

// SetCursor sets field value
func (o *GetHoldsResponse) SetCursor(v GetHoldsResponseCursor) {
	o.Cursor = v
}

func (o GetHoldsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cursor"] = o.Cursor
	}
	return json.Marshal(toSerialize)
}

type NullableGetHoldsResponse struct {
	value *GetHoldsResponse
	isSet bool
}

func (v NullableGetHoldsResponse) Get() *GetHoldsResponse {
	return v.value
}

func (v *NullableGetHoldsResponse) Set(val *GetHoldsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetHoldsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetHoldsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetHoldsResponse(val *GetHoldsResponse) *NullableGetHoldsResponse {
	return &NullableGetHoldsResponse{value: val, isSet: true}
}

func (v NullableGetHoldsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetHoldsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


