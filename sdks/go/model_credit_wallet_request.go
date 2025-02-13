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

// checks if the CreditWalletRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreditWalletRequest{}

// CreditWalletRequest struct for CreditWalletRequest
type CreditWalletRequest struct {
	Amount Monetary `json:"amount"`
	// Metadata associated with the wallet.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Reference *string `json:"reference,omitempty"`
	Sources []Subject `json:"sources"`
	// The balance to credit
	Balance *string `json:"balance,omitempty"`
}

// NewCreditWalletRequest instantiates a new CreditWalletRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditWalletRequest(amount Monetary, sources []Subject) *CreditWalletRequest {
	this := CreditWalletRequest{}
	this.Amount = amount
	this.Sources = sources
	return &this
}

// NewCreditWalletRequestWithDefaults instantiates a new CreditWalletRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditWalletRequestWithDefaults() *CreditWalletRequest {
	this := CreditWalletRequest{}
	return &this
}

// GetAmount returns the Amount field value
func (o *CreditWalletRequest) GetAmount() Monetary {
	if o == nil {
		var ret Monetary
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *CreditWalletRequest) GetAmountOk() (*Monetary, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *CreditWalletRequest) SetAmount(v Monetary) {
	o.Amount = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CreditWalletRequest) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditWalletRequest) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CreditWalletRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *CreditWalletRequest) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *CreditWalletRequest) GetReference() string {
	if o == nil || IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditWalletRequest) GetReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *CreditWalletRequest) HasReference() bool {
	if o != nil && !IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *CreditWalletRequest) SetReference(v string) {
	o.Reference = &v
}

// GetSources returns the Sources field value
func (o *CreditWalletRequest) GetSources() []Subject {
	if o == nil {
		var ret []Subject
		return ret
	}

	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value
// and a boolean to check if the value has been set.
func (o *CreditWalletRequest) GetSourcesOk() ([]Subject, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sources, true
}

// SetSources sets field value
func (o *CreditWalletRequest) SetSources(v []Subject) {
	o.Sources = v
}

// GetBalance returns the Balance field value if set, zero value otherwise.
func (o *CreditWalletRequest) GetBalance() string {
	if o == nil || IsNil(o.Balance) {
		var ret string
		return ret
	}
	return *o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditWalletRequest) GetBalanceOk() (*string, bool) {
	if o == nil || IsNil(o.Balance) {
		return nil, false
	}
	return o.Balance, true
}

// HasBalance returns a boolean if a field has been set.
func (o *CreditWalletRequest) HasBalance() bool {
	if o != nil && !IsNil(o.Balance) {
		return true
	}

	return false
}

// SetBalance gets a reference to the given string and assigns it to the Balance field.
func (o *CreditWalletRequest) SetBalance(v string) {
	o.Balance = &v
}

func (o CreditWalletRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreditWalletRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	toSerialize["sources"] = o.Sources
	if !IsNil(o.Balance) {
		toSerialize["balance"] = o.Balance
	}
	return toSerialize, nil
}

type NullableCreditWalletRequest struct {
	value *CreditWalletRequest
	isSet bool
}

func (v NullableCreditWalletRequest) Get() *CreditWalletRequest {
	return v.value
}

func (v *NullableCreditWalletRequest) Set(val *CreditWalletRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditWalletRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditWalletRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditWalletRequest(val *CreditWalletRequest) *NullableCreditWalletRequest {
	return &NullableCreditWalletRequest{value: val, isSet: true}
}

func (v NullableCreditWalletRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditWalletRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


