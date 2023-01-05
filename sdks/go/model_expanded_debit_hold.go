/*
Formance Stack API

Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 

API version: SDK_VERSION
Contact: support@formance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package formance

import (
	"encoding/json"
)

// ExpandedDebitHold struct for ExpandedDebitHold
type ExpandedDebitHold struct {
	// The unique ID of the hold.
	Id string `json:"id"`
	// The ID of the wallet the hold is associated with.
	WalletID string `json:"walletID"`
	// Metadata associated with the hold.
	Metadata map[string]interface{} `json:"metadata"`
	Description string `json:"description"`
	// Remaining amount on hold
	Remaining int64 `json:"remaining"`
	// Original amount on hold
	OriginalAmount int64 `json:"originalAmount"`
}

// NewExpandedDebitHold instantiates a new ExpandedDebitHold object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExpandedDebitHold(id string, walletID string, metadata map[string]interface{}, description string, remaining int64, originalAmount int64) *ExpandedDebitHold {
	this := ExpandedDebitHold{}
	this.Id = id
	this.WalletID = walletID
	this.Metadata = metadata
	this.Description = description
	this.Remaining = remaining
	this.OriginalAmount = originalAmount
	return &this
}

// NewExpandedDebitHoldWithDefaults instantiates a new ExpandedDebitHold object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExpandedDebitHoldWithDefaults() *ExpandedDebitHold {
	this := ExpandedDebitHold{}
	return &this
}

// GetId returns the Id field value
func (o *ExpandedDebitHold) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ExpandedDebitHold) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ExpandedDebitHold) SetId(v string) {
	o.Id = v
}

// GetWalletID returns the WalletID field value
func (o *ExpandedDebitHold) GetWalletID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletID
}

// GetWalletIDOk returns a tuple with the WalletID field value
// and a boolean to check if the value has been set.
func (o *ExpandedDebitHold) GetWalletIDOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.WalletID, true
}

// SetWalletID sets field value
func (o *ExpandedDebitHold) SetWalletID(v string) {
	o.WalletID = v
}

// GetMetadata returns the Metadata field value
func (o *ExpandedDebitHold) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *ExpandedDebitHold) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil {
    return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// SetMetadata sets field value
func (o *ExpandedDebitHold) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetDescription returns the Description field value
func (o *ExpandedDebitHold) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ExpandedDebitHold) GetDescriptionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ExpandedDebitHold) SetDescription(v string) {
	o.Description = v
}

// GetRemaining returns the Remaining field value
func (o *ExpandedDebitHold) GetRemaining() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Remaining
}

// GetRemainingOk returns a tuple with the Remaining field value
// and a boolean to check if the value has been set.
func (o *ExpandedDebitHold) GetRemainingOk() (*int64, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Remaining, true
}

// SetRemaining sets field value
func (o *ExpandedDebitHold) SetRemaining(v int64) {
	o.Remaining = v
}

// GetOriginalAmount returns the OriginalAmount field value
func (o *ExpandedDebitHold) GetOriginalAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.OriginalAmount
}

// GetOriginalAmountOk returns a tuple with the OriginalAmount field value
// and a boolean to check if the value has been set.
func (o *ExpandedDebitHold) GetOriginalAmountOk() (*int64, bool) {
	if o == nil {
    return nil, false
	}
	return &o.OriginalAmount, true
}

// SetOriginalAmount sets field value
func (o *ExpandedDebitHold) SetOriginalAmount(v int64) {
	o.OriginalAmount = v
}

func (o ExpandedDebitHold) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["walletID"] = o.WalletID
	}
	if true {
		toSerialize["metadata"] = o.Metadata
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["remaining"] = o.Remaining
	}
	if true {
		toSerialize["originalAmount"] = o.OriginalAmount
	}
	return json.Marshal(toSerialize)
}

type NullableExpandedDebitHold struct {
	value *ExpandedDebitHold
	isSet bool
}

func (v NullableExpandedDebitHold) Get() *ExpandedDebitHold {
	return v.value
}

func (v *NullableExpandedDebitHold) Set(val *ExpandedDebitHold) {
	v.value = val
	v.isSet = true
}

func (v NullableExpandedDebitHold) IsSet() bool {
	return v.isSet
}

func (v *NullableExpandedDebitHold) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpandedDebitHold(val *ExpandedDebitHold) *NullableExpandedDebitHold {
	return &NullableExpandedDebitHold{value: val, isSet: true}
}

func (v NullableExpandedDebitHold) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpandedDebitHold) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


