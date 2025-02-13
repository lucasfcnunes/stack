/*
Membership API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package membershipclient

import (
	"encoding/json"
)

// checks if the ListOrganizationExpandedResponseDataInnerAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListOrganizationExpandedResponseDataInnerAllOf{}

// ListOrganizationExpandedResponseDataInnerAllOf struct for ListOrganizationExpandedResponseDataInnerAllOf
type ListOrganizationExpandedResponseDataInnerAllOf struct {
	TotalStacks *int32 `json:"totalStacks,omitempty"`
	TotalUsers  *int32 `json:"totalUsers,omitempty"`
	Owner       *User  `json:"owner,omitempty"`
}

// NewListOrganizationExpandedResponseDataInnerAllOf instantiates a new ListOrganizationExpandedResponseDataInnerAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListOrganizationExpandedResponseDataInnerAllOf() *ListOrganizationExpandedResponseDataInnerAllOf {
	this := ListOrganizationExpandedResponseDataInnerAllOf{}
	return &this
}

// NewListOrganizationExpandedResponseDataInnerAllOfWithDefaults instantiates a new ListOrganizationExpandedResponseDataInnerAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListOrganizationExpandedResponseDataInnerAllOfWithDefaults() *ListOrganizationExpandedResponseDataInnerAllOf {
	this := ListOrganizationExpandedResponseDataInnerAllOf{}
	return &this
}

// GetTotalStacks returns the TotalStacks field value if set, zero value otherwise.
func (o *ListOrganizationExpandedResponseDataInnerAllOf) GetTotalStacks() int32 {
	if o == nil || IsNil(o.TotalStacks) {
		var ret int32
		return ret
	}
	return *o.TotalStacks
}

// GetTotalStacksOk returns a tuple with the TotalStacks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOrganizationExpandedResponseDataInnerAllOf) GetTotalStacksOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalStacks) {
		return nil, false
	}
	return o.TotalStacks, true
}

// HasTotalStacks returns a boolean if a field has been set.
func (o *ListOrganizationExpandedResponseDataInnerAllOf) HasTotalStacks() bool {
	if o != nil && !IsNil(o.TotalStacks) {
		return true
	}

	return false
}

// SetTotalStacks gets a reference to the given int32 and assigns it to the TotalStacks field.
func (o *ListOrganizationExpandedResponseDataInnerAllOf) SetTotalStacks(v int32) {
	o.TotalStacks = &v
}

// GetTotalUsers returns the TotalUsers field value if set, zero value otherwise.
func (o *ListOrganizationExpandedResponseDataInnerAllOf) GetTotalUsers() int32 {
	if o == nil || IsNil(o.TotalUsers) {
		var ret int32
		return ret
	}
	return *o.TotalUsers
}

// GetTotalUsersOk returns a tuple with the TotalUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOrganizationExpandedResponseDataInnerAllOf) GetTotalUsersOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalUsers) {
		return nil, false
	}
	return o.TotalUsers, true
}

// HasTotalUsers returns a boolean if a field has been set.
func (o *ListOrganizationExpandedResponseDataInnerAllOf) HasTotalUsers() bool {
	if o != nil && !IsNil(o.TotalUsers) {
		return true
	}

	return false
}

// SetTotalUsers gets a reference to the given int32 and assigns it to the TotalUsers field.
func (o *ListOrganizationExpandedResponseDataInnerAllOf) SetTotalUsers(v int32) {
	o.TotalUsers = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *ListOrganizationExpandedResponseDataInnerAllOf) GetOwner() User {
	if o == nil || IsNil(o.Owner) {
		var ret User
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOrganizationExpandedResponseDataInnerAllOf) GetOwnerOk() (*User, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *ListOrganizationExpandedResponseDataInnerAllOf) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given User and assigns it to the Owner field.
func (o *ListOrganizationExpandedResponseDataInnerAllOf) SetOwner(v User) {
	o.Owner = &v
}

func (o ListOrganizationExpandedResponseDataInnerAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListOrganizationExpandedResponseDataInnerAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalStacks) {
		toSerialize["totalStacks"] = o.TotalStacks
	}
	if !IsNil(o.TotalUsers) {
		toSerialize["totalUsers"] = o.TotalUsers
	}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	return toSerialize, nil
}

type NullableListOrganizationExpandedResponseDataInnerAllOf struct {
	value *ListOrganizationExpandedResponseDataInnerAllOf
	isSet bool
}

func (v NullableListOrganizationExpandedResponseDataInnerAllOf) Get() *ListOrganizationExpandedResponseDataInnerAllOf {
	return v.value
}

func (v *NullableListOrganizationExpandedResponseDataInnerAllOf) Set(val *ListOrganizationExpandedResponseDataInnerAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableListOrganizationExpandedResponseDataInnerAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableListOrganizationExpandedResponseDataInnerAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListOrganizationExpandedResponseDataInnerAllOf(val *ListOrganizationExpandedResponseDataInnerAllOf) *NullableListOrganizationExpandedResponseDataInnerAllOf {
	return &NullableListOrganizationExpandedResponseDataInnerAllOf{value: val, isSet: true}
}

func (v NullableListOrganizationExpandedResponseDataInnerAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListOrganizationExpandedResponseDataInnerAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
