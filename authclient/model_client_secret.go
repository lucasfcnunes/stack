/*
Auth API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: AUTH_VERSION
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authclient

import (
	"encoding/json"
)

// ClientSecret struct for ClientSecret
type ClientSecret struct {
	LastDigits float32 `json:"lastDigits"`
	Name string `json:"name"`
	Id string `json:"id"`
}

// NewClientSecret instantiates a new ClientSecret object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientSecret(lastDigits float32, name string, id string) *ClientSecret {
	this := ClientSecret{}
	this.LastDigits = lastDigits
	this.Name = name
	this.Id = id
	return &this
}

// NewClientSecretWithDefaults instantiates a new ClientSecret object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientSecretWithDefaults() *ClientSecret {
	this := ClientSecret{}
	return &this
}

// GetLastDigits returns the LastDigits field value
func (o *ClientSecret) GetLastDigits() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.LastDigits
}

// GetLastDigitsOk returns a tuple with the LastDigits field value
// and a boolean to check if the value has been set.
func (o *ClientSecret) GetLastDigitsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastDigits, true
}

// SetLastDigits sets field value
func (o *ClientSecret) SetLastDigits(v float32) {
	o.LastDigits = v
}

// GetName returns the Name field value
func (o *ClientSecret) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ClientSecret) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ClientSecret) SetName(v string) {
	o.Name = v
}

// GetId returns the Id field value
func (o *ClientSecret) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ClientSecret) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ClientSecret) SetId(v string) {
	o.Id = v
}

func (o ClientSecret) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["lastDigits"] = o.LastDigits
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableClientSecret struct {
	value *ClientSecret
	isSet bool
}

func (v NullableClientSecret) Get() *ClientSecret {
	return v.value
}

func (v *NullableClientSecret) Set(val *ClientSecret) {
	v.value = val
	v.isSet = true
}

func (v NullableClientSecret) IsSet() bool {
	return v.isSet
}

func (v *NullableClientSecret) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientSecret(val *ClientSecret) *NullableClientSecret {
	return &NullableClientSecret{value: val, isSet: true}
}

func (v NullableClientSecret) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientSecret) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


