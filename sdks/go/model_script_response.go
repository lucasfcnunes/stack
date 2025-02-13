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

// checks if the ScriptResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScriptResponse{}

// ScriptResponse struct for ScriptResponse
type ScriptResponse struct {
	ErrorCode *ErrorsEnum `json:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty"`
	Details *string `json:"details,omitempty"`
	Transaction *Transaction `json:"transaction,omitempty"`
}

// NewScriptResponse instantiates a new ScriptResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScriptResponse() *ScriptResponse {
	this := ScriptResponse{}
	return &this
}

// NewScriptResponseWithDefaults instantiates a new ScriptResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScriptResponseWithDefaults() *ScriptResponse {
	this := ScriptResponse{}
	return &this
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *ScriptResponse) GetErrorCode() ErrorsEnum {
	if o == nil || IsNil(o.ErrorCode) {
		var ret ErrorsEnum
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScriptResponse) GetErrorCodeOk() (*ErrorsEnum, bool) {
	if o == nil || IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *ScriptResponse) HasErrorCode() bool {
	if o != nil && !IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given ErrorsEnum and assigns it to the ErrorCode field.
func (o *ScriptResponse) SetErrorCode(v ErrorsEnum) {
	o.ErrorCode = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *ScriptResponse) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScriptResponse) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *ScriptResponse) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *ScriptResponse) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *ScriptResponse) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScriptResponse) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *ScriptResponse) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *ScriptResponse) SetDetails(v string) {
	o.Details = &v
}

// GetTransaction returns the Transaction field value if set, zero value otherwise.
func (o *ScriptResponse) GetTransaction() Transaction {
	if o == nil || IsNil(o.Transaction) {
		var ret Transaction
		return ret
	}
	return *o.Transaction
}

// GetTransactionOk returns a tuple with the Transaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScriptResponse) GetTransactionOk() (*Transaction, bool) {
	if o == nil || IsNil(o.Transaction) {
		return nil, false
	}
	return o.Transaction, true
}

// HasTransaction returns a boolean if a field has been set.
func (o *ScriptResponse) HasTransaction() bool {
	if o != nil && !IsNil(o.Transaction) {
		return true
	}

	return false
}

// SetTransaction gets a reference to the given Transaction and assigns it to the Transaction field.
func (o *ScriptResponse) SetTransaction(v Transaction) {
	o.Transaction = &v
}

func (o ScriptResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScriptResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ErrorCode) {
		toSerialize["errorCode"] = o.ErrorCode
	}
	if !IsNil(o.ErrorMessage) {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if !IsNil(o.Transaction) {
		toSerialize["transaction"] = o.Transaction
	}
	return toSerialize, nil
}

type NullableScriptResponse struct {
	value *ScriptResponse
	isSet bool
}

func (v NullableScriptResponse) Get() *ScriptResponse {
	return v.value
}

func (v *NullableScriptResponse) Set(val *ScriptResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableScriptResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableScriptResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScriptResponse(val *ScriptResponse) *NullableScriptResponse {
	return &NullableScriptResponse{value: val, isSet: true}
}

func (v NullableScriptResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScriptResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


