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

// checks if the TransactionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionsResponse{}

// TransactionsResponse struct for TransactionsResponse
type TransactionsResponse struct {
	Data []Transaction `json:"data"`
}

// NewTransactionsResponse instantiates a new TransactionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionsResponse(data []Transaction) *TransactionsResponse {
	this := TransactionsResponse{}
	this.Data = data
	return &this
}

// NewTransactionsResponseWithDefaults instantiates a new TransactionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionsResponseWithDefaults() *TransactionsResponse {
	this := TransactionsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *TransactionsResponse) GetData() []Transaction {
	if o == nil {
		var ret []Transaction
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *TransactionsResponse) GetDataOk() ([]Transaction, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *TransactionsResponse) SetData(v []Transaction) {
	o.Data = v
}

func (o TransactionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableTransactionsResponse struct {
	value *TransactionsResponse
	isSet bool
}

func (v NullableTransactionsResponse) Get() *TransactionsResponse {
	return v.value
}

func (v *NullableTransactionsResponse) Set(val *TransactionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionsResponse(val *TransactionsResponse) *NullableTransactionsResponse {
	return &NullableTransactionsResponse{value: val, isSet: true}
}

func (v NullableTransactionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


