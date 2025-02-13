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

// checks if the WorkflowInstanceHistoryStageOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowInstanceHistoryStageOutput{}

// WorkflowInstanceHistoryStageOutput struct for WorkflowInstanceHistoryStageOutput
type WorkflowInstanceHistoryStageOutput struct {
	GetAccount *AccountResponse `json:"GetAccount,omitempty"`
	CreateTransaction *TransactionsResponse `json:"CreateTransaction,omitempty"`
	RevertTransaction *TransactionResponse `json:"RevertTransaction,omitempty"`
	GetPayment *PaymentResponse `json:"GetPayment,omitempty"`
	DebitWallet *DebitWalletResponse `json:"DebitWallet,omitempty"`
	GetWallet *GetWalletResponse `json:"GetWallet,omitempty"`
}

// NewWorkflowInstanceHistoryStageOutput instantiates a new WorkflowInstanceHistoryStageOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowInstanceHistoryStageOutput() *WorkflowInstanceHistoryStageOutput {
	this := WorkflowInstanceHistoryStageOutput{}
	return &this
}

// NewWorkflowInstanceHistoryStageOutputWithDefaults instantiates a new WorkflowInstanceHistoryStageOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowInstanceHistoryStageOutputWithDefaults() *WorkflowInstanceHistoryStageOutput {
	this := WorkflowInstanceHistoryStageOutput{}
	return &this
}

// GetGetAccount returns the GetAccount field value if set, zero value otherwise.
func (o *WorkflowInstanceHistoryStageOutput) GetGetAccount() AccountResponse {
	if o == nil || IsNil(o.GetAccount) {
		var ret AccountResponse
		return ret
	}
	return *o.GetAccount
}

// GetGetAccountOk returns a tuple with the GetAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowInstanceHistoryStageOutput) GetGetAccountOk() (*AccountResponse, bool) {
	if o == nil || IsNil(o.GetAccount) {
		return nil, false
	}
	return o.GetAccount, true
}

// HasGetAccount returns a boolean if a field has been set.
func (o *WorkflowInstanceHistoryStageOutput) HasGetAccount() bool {
	if o != nil && !IsNil(o.GetAccount) {
		return true
	}

	return false
}

// SetGetAccount gets a reference to the given AccountResponse and assigns it to the GetAccount field.
func (o *WorkflowInstanceHistoryStageOutput) SetGetAccount(v AccountResponse) {
	o.GetAccount = &v
}

// GetCreateTransaction returns the CreateTransaction field value if set, zero value otherwise.
func (o *WorkflowInstanceHistoryStageOutput) GetCreateTransaction() TransactionsResponse {
	if o == nil || IsNil(o.CreateTransaction) {
		var ret TransactionsResponse
		return ret
	}
	return *o.CreateTransaction
}

// GetCreateTransactionOk returns a tuple with the CreateTransaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowInstanceHistoryStageOutput) GetCreateTransactionOk() (*TransactionsResponse, bool) {
	if o == nil || IsNil(o.CreateTransaction) {
		return nil, false
	}
	return o.CreateTransaction, true
}

// HasCreateTransaction returns a boolean if a field has been set.
func (o *WorkflowInstanceHistoryStageOutput) HasCreateTransaction() bool {
	if o != nil && !IsNil(o.CreateTransaction) {
		return true
	}

	return false
}

// SetCreateTransaction gets a reference to the given TransactionsResponse and assigns it to the CreateTransaction field.
func (o *WorkflowInstanceHistoryStageOutput) SetCreateTransaction(v TransactionsResponse) {
	o.CreateTransaction = &v
}

// GetRevertTransaction returns the RevertTransaction field value if set, zero value otherwise.
func (o *WorkflowInstanceHistoryStageOutput) GetRevertTransaction() TransactionResponse {
	if o == nil || IsNil(o.RevertTransaction) {
		var ret TransactionResponse
		return ret
	}
	return *o.RevertTransaction
}

// GetRevertTransactionOk returns a tuple with the RevertTransaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowInstanceHistoryStageOutput) GetRevertTransactionOk() (*TransactionResponse, bool) {
	if o == nil || IsNil(o.RevertTransaction) {
		return nil, false
	}
	return o.RevertTransaction, true
}

// HasRevertTransaction returns a boolean if a field has been set.
func (o *WorkflowInstanceHistoryStageOutput) HasRevertTransaction() bool {
	if o != nil && !IsNil(o.RevertTransaction) {
		return true
	}

	return false
}

// SetRevertTransaction gets a reference to the given TransactionResponse and assigns it to the RevertTransaction field.
func (o *WorkflowInstanceHistoryStageOutput) SetRevertTransaction(v TransactionResponse) {
	o.RevertTransaction = &v
}

// GetGetPayment returns the GetPayment field value if set, zero value otherwise.
func (o *WorkflowInstanceHistoryStageOutput) GetGetPayment() PaymentResponse {
	if o == nil || IsNil(o.GetPayment) {
		var ret PaymentResponse
		return ret
	}
	return *o.GetPayment
}

// GetGetPaymentOk returns a tuple with the GetPayment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowInstanceHistoryStageOutput) GetGetPaymentOk() (*PaymentResponse, bool) {
	if o == nil || IsNil(o.GetPayment) {
		return nil, false
	}
	return o.GetPayment, true
}

// HasGetPayment returns a boolean if a field has been set.
func (o *WorkflowInstanceHistoryStageOutput) HasGetPayment() bool {
	if o != nil && !IsNil(o.GetPayment) {
		return true
	}

	return false
}

// SetGetPayment gets a reference to the given PaymentResponse and assigns it to the GetPayment field.
func (o *WorkflowInstanceHistoryStageOutput) SetGetPayment(v PaymentResponse) {
	o.GetPayment = &v
}

// GetDebitWallet returns the DebitWallet field value if set, zero value otherwise.
func (o *WorkflowInstanceHistoryStageOutput) GetDebitWallet() DebitWalletResponse {
	if o == nil || IsNil(o.DebitWallet) {
		var ret DebitWalletResponse
		return ret
	}
	return *o.DebitWallet
}

// GetDebitWalletOk returns a tuple with the DebitWallet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowInstanceHistoryStageOutput) GetDebitWalletOk() (*DebitWalletResponse, bool) {
	if o == nil || IsNil(o.DebitWallet) {
		return nil, false
	}
	return o.DebitWallet, true
}

// HasDebitWallet returns a boolean if a field has been set.
func (o *WorkflowInstanceHistoryStageOutput) HasDebitWallet() bool {
	if o != nil && !IsNil(o.DebitWallet) {
		return true
	}

	return false
}

// SetDebitWallet gets a reference to the given DebitWalletResponse and assigns it to the DebitWallet field.
func (o *WorkflowInstanceHistoryStageOutput) SetDebitWallet(v DebitWalletResponse) {
	o.DebitWallet = &v
}

// GetGetWallet returns the GetWallet field value if set, zero value otherwise.
func (o *WorkflowInstanceHistoryStageOutput) GetGetWallet() GetWalletResponse {
	if o == nil || IsNil(o.GetWallet) {
		var ret GetWalletResponse
		return ret
	}
	return *o.GetWallet
}

// GetGetWalletOk returns a tuple with the GetWallet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowInstanceHistoryStageOutput) GetGetWalletOk() (*GetWalletResponse, bool) {
	if o == nil || IsNil(o.GetWallet) {
		return nil, false
	}
	return o.GetWallet, true
}

// HasGetWallet returns a boolean if a field has been set.
func (o *WorkflowInstanceHistoryStageOutput) HasGetWallet() bool {
	if o != nil && !IsNil(o.GetWallet) {
		return true
	}

	return false
}

// SetGetWallet gets a reference to the given GetWalletResponse and assigns it to the GetWallet field.
func (o *WorkflowInstanceHistoryStageOutput) SetGetWallet(v GetWalletResponse) {
	o.GetWallet = &v
}

func (o WorkflowInstanceHistoryStageOutput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowInstanceHistoryStageOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GetAccount) {
		toSerialize["GetAccount"] = o.GetAccount
	}
	if !IsNil(o.CreateTransaction) {
		toSerialize["CreateTransaction"] = o.CreateTransaction
	}
	if !IsNil(o.RevertTransaction) {
		toSerialize["RevertTransaction"] = o.RevertTransaction
	}
	if !IsNil(o.GetPayment) {
		toSerialize["GetPayment"] = o.GetPayment
	}
	if !IsNil(o.DebitWallet) {
		toSerialize["DebitWallet"] = o.DebitWallet
	}
	if !IsNil(o.GetWallet) {
		toSerialize["GetWallet"] = o.GetWallet
	}
	return toSerialize, nil
}

type NullableWorkflowInstanceHistoryStageOutput struct {
	value *WorkflowInstanceHistoryStageOutput
	isSet bool
}

func (v NullableWorkflowInstanceHistoryStageOutput) Get() *WorkflowInstanceHistoryStageOutput {
	return v.value
}

func (v *NullableWorkflowInstanceHistoryStageOutput) Set(val *WorkflowInstanceHistoryStageOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowInstanceHistoryStageOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowInstanceHistoryStageOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowInstanceHistoryStageOutput(val *WorkflowInstanceHistoryStageOutput) *NullableWorkflowInstanceHistoryStageOutput {
	return &NullableWorkflowInstanceHistoryStageOutput{value: val, isSet: true}
}

func (v NullableWorkflowInstanceHistoryStageOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowInstanceHistoryStageOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


