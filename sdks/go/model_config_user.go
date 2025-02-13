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

// checks if the ConfigUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigUser{}

// ConfigUser struct for ConfigUser
type ConfigUser struct {
	Endpoint string `json:"endpoint"`
	Secret *string `json:"secret,omitempty"`
	EventTypes []string `json:"eventTypes"`
}

// NewConfigUser instantiates a new ConfigUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigUser(endpoint string, eventTypes []string) *ConfigUser {
	this := ConfigUser{}
	this.Endpoint = endpoint
	this.EventTypes = eventTypes
	return &this
}

// NewConfigUserWithDefaults instantiates a new ConfigUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigUserWithDefaults() *ConfigUser {
	this := ConfigUser{}
	return &this
}

// GetEndpoint returns the Endpoint field value
func (o *ConfigUser) GetEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value
// and a boolean to check if the value has been set.
func (o *ConfigUser) GetEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Endpoint, true
}

// SetEndpoint sets field value
func (o *ConfigUser) SetEndpoint(v string) {
	o.Endpoint = v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *ConfigUser) GetSecret() string {
	if o == nil || IsNil(o.Secret) {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigUser) GetSecretOk() (*string, bool) {
	if o == nil || IsNil(o.Secret) {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *ConfigUser) HasSecret() bool {
	if o != nil && !IsNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *ConfigUser) SetSecret(v string) {
	o.Secret = &v
}

// GetEventTypes returns the EventTypes field value
func (o *ConfigUser) GetEventTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.EventTypes
}

// GetEventTypesOk returns a tuple with the EventTypes field value
// and a boolean to check if the value has been set.
func (o *ConfigUser) GetEventTypesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EventTypes, true
}

// SetEventTypes sets field value
func (o *ConfigUser) SetEventTypes(v []string) {
	o.EventTypes = v
}

func (o ConfigUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["endpoint"] = o.Endpoint
	if !IsNil(o.Secret) {
		toSerialize["secret"] = o.Secret
	}
	toSerialize["eventTypes"] = o.EventTypes
	return toSerialize, nil
}

type NullableConfigUser struct {
	value *ConfigUser
	isSet bool
}

func (v NullableConfigUser) Get() *ConfigUser {
	return v.value
}

func (v *NullableConfigUser) Set(val *ConfigUser) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigUser) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigUser(val *ConfigUser) *NullableConfigUser {
	return &NullableConfigUser{value: val, isSet: true}
}

func (v NullableConfigUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


