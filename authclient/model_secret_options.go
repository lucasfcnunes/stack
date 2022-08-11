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

// SecretOptions struct for SecretOptions
type SecretOptions struct {
	Name string `json:"name"`
}

// NewSecretOptions instantiates a new SecretOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecretOptions(name string) *SecretOptions {
	this := SecretOptions{}
	this.Name = name
	return &this
}

// NewSecretOptionsWithDefaults instantiates a new SecretOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecretOptionsWithDefaults() *SecretOptions {
	this := SecretOptions{}
	return &this
}

// GetName returns the Name field value
func (o *SecretOptions) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SecretOptions) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SecretOptions) SetName(v string) {
	o.Name = v
}

func (o SecretOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableSecretOptions struct {
	value *SecretOptions
	isSet bool
}

func (v NullableSecretOptions) Get() *SecretOptions {
	return v.value
}

func (v *NullableSecretOptions) Set(val *SecretOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableSecretOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableSecretOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecretOptions(val *SecretOptions) *NullableSecretOptions {
	return &NullableSecretOptions{value: val, isSet: true}
}

func (v NullableSecretOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecretOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


