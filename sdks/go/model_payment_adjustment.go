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
	"time"
)

// checks if the PaymentAdjustment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentAdjustment{}

// PaymentAdjustment struct for PaymentAdjustment
type PaymentAdjustment struct {
	Status PaymentStatus `json:"status"`
	Amount int64 `json:"amount"`
	Date time.Time `json:"date"`
	Raw map[string]interface{} `json:"raw"`
	Absolute bool `json:"absolute"`
}

// NewPaymentAdjustment instantiates a new PaymentAdjustment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentAdjustment(status PaymentStatus, amount int64, date time.Time, raw map[string]interface{}, absolute bool) *PaymentAdjustment {
	this := PaymentAdjustment{}
	this.Status = status
	this.Amount = amount
	this.Date = date
	this.Raw = raw
	this.Absolute = absolute
	return &this
}

// NewPaymentAdjustmentWithDefaults instantiates a new PaymentAdjustment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentAdjustmentWithDefaults() *PaymentAdjustment {
	this := PaymentAdjustment{}
	return &this
}

// GetStatus returns the Status field value
func (o *PaymentAdjustment) GetStatus() PaymentStatus {
	if o == nil {
		var ret PaymentStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PaymentAdjustment) GetStatusOk() (*PaymentStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PaymentAdjustment) SetStatus(v PaymentStatus) {
	o.Status = v
}

// GetAmount returns the Amount field value
func (o *PaymentAdjustment) GetAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *PaymentAdjustment) GetAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *PaymentAdjustment) SetAmount(v int64) {
	o.Amount = v
}

// GetDate returns the Date field value
func (o *PaymentAdjustment) GetDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *PaymentAdjustment) GetDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *PaymentAdjustment) SetDate(v time.Time) {
	o.Date = v
}

// GetRaw returns the Raw field value
func (o *PaymentAdjustment) GetRaw() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Raw
}

// GetRawOk returns a tuple with the Raw field value
// and a boolean to check if the value has been set.
func (o *PaymentAdjustment) GetRawOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Raw, true
}

// SetRaw sets field value
func (o *PaymentAdjustment) SetRaw(v map[string]interface{}) {
	o.Raw = v
}

// GetAbsolute returns the Absolute field value
func (o *PaymentAdjustment) GetAbsolute() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Absolute
}

// GetAbsoluteOk returns a tuple with the Absolute field value
// and a boolean to check if the value has been set.
func (o *PaymentAdjustment) GetAbsoluteOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Absolute, true
}

// SetAbsolute sets field value
func (o *PaymentAdjustment) SetAbsolute(v bool) {
	o.Absolute = v
}

func (o PaymentAdjustment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentAdjustment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["amount"] = o.Amount
	toSerialize["date"] = o.Date
	toSerialize["raw"] = o.Raw
	toSerialize["absolute"] = o.Absolute
	return toSerialize, nil
}

type NullablePaymentAdjustment struct {
	value *PaymentAdjustment
	isSet bool
}

func (v NullablePaymentAdjustment) Get() *PaymentAdjustment {
	return v.value
}

func (v *NullablePaymentAdjustment) Set(val *PaymentAdjustment) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentAdjustment) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentAdjustment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentAdjustment(val *PaymentAdjustment) *NullablePaymentAdjustment {
	return &NullablePaymentAdjustment{value: val, isSet: true}
}

func (v NullablePaymentAdjustment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentAdjustment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


