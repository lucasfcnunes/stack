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

// checks if the Script type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Script{}

// Script struct for Script
type Script struct {
	Plain string `json:"plain"`
	Vars map[string]interface{} `json:"vars,omitempty"`
	// Reference to attach to the generated transaction
	Reference *string `json:"reference,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewScript instantiates a new Script object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScript(plain string) *Script {
	this := Script{}
	this.Plain = plain
	return &this
}

// NewScriptWithDefaults instantiates a new Script object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScriptWithDefaults() *Script {
	this := Script{}
	return &this
}

// GetPlain returns the Plain field value
func (o *Script) GetPlain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Plain
}

// GetPlainOk returns a tuple with the Plain field value
// and a boolean to check if the value has been set.
func (o *Script) GetPlainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Plain, true
}

// SetPlain sets field value
func (o *Script) SetPlain(v string) {
	o.Plain = v
}

// GetVars returns the Vars field value if set, zero value otherwise.
func (o *Script) GetVars() map[string]interface{} {
	if o == nil || IsNil(o.Vars) {
		var ret map[string]interface{}
		return ret
	}
	return o.Vars
}

// GetVarsOk returns a tuple with the Vars field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Script) GetVarsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Vars) {
		return map[string]interface{}{}, false
	}
	return o.Vars, true
}

// HasVars returns a boolean if a field has been set.
func (o *Script) HasVars() bool {
	if o != nil && !IsNil(o.Vars) {
		return true
	}

	return false
}

// SetVars gets a reference to the given map[string]interface{} and assigns it to the Vars field.
func (o *Script) SetVars(v map[string]interface{}) {
	o.Vars = v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *Script) GetReference() string {
	if o == nil || IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Script) GetReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *Script) HasReference() bool {
	if o != nil && !IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *Script) SetReference(v string) {
	o.Reference = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Script) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Script) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Script) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *Script) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o Script) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Script) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["plain"] = o.Plain
	if !IsNil(o.Vars) {
		toSerialize["vars"] = o.Vars
	}
	if !IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableScript struct {
	value *Script
	isSet bool
}

func (v NullableScript) Get() *Script {
	return v.value
}

func (v *NullableScript) Set(val *Script) {
	v.value = val
	v.isSet = true
}

func (v NullableScript) IsSet() bool {
	return v.isSet
}

func (v *NullableScript) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScript(val *Script) *NullableScript {
	return &NullableScript{value: val, isSet: true}
}

func (v NullableScript) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScript) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


