/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
	"fmt"
)

// checks if the NonEmployeeSourceRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NonEmployeeSourceRequestBody{}

// NonEmployeeSourceRequestBody struct for NonEmployeeSourceRequestBody
type NonEmployeeSourceRequestBody struct {
	// Name of non-employee source.
	Name string `json:"name"`
	// Description of non-employee source.
	Description string `json:"description"`
	Owner NonEmployeeIdnUserRequest `json:"owner"`
	// The ID for the management workgroup that contains source sub-admins
	ManagementWorkgroup *string `json:"managementWorkgroup,omitempty"`
	// List of approvers.
	Approvers []NonEmployeeIdnUserRequest `json:"approvers,omitempty"`
	// List of account managers.
	AccountManagers []NonEmployeeIdnUserRequest `json:"accountManagers,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NonEmployeeSourceRequestBody NonEmployeeSourceRequestBody

// NewNonEmployeeSourceRequestBody instantiates a new NonEmployeeSourceRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNonEmployeeSourceRequestBody(name string, description string, owner NonEmployeeIdnUserRequest) *NonEmployeeSourceRequestBody {
	this := NonEmployeeSourceRequestBody{}
	this.Name = name
	this.Description = description
	this.Owner = owner
	return &this
}

// NewNonEmployeeSourceRequestBodyWithDefaults instantiates a new NonEmployeeSourceRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNonEmployeeSourceRequestBodyWithDefaults() *NonEmployeeSourceRequestBody {
	this := NonEmployeeSourceRequestBody{}
	return &this
}

// GetName returns the Name field value
func (o *NonEmployeeSourceRequestBody) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NonEmployeeSourceRequestBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NonEmployeeSourceRequestBody) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *NonEmployeeSourceRequestBody) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *NonEmployeeSourceRequestBody) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *NonEmployeeSourceRequestBody) SetDescription(v string) {
	o.Description = v
}

// GetOwner returns the Owner field value
func (o *NonEmployeeSourceRequestBody) GetOwner() NonEmployeeIdnUserRequest {
	if o == nil {
		var ret NonEmployeeIdnUserRequest
		return ret
	}

	return o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value
// and a boolean to check if the value has been set.
func (o *NonEmployeeSourceRequestBody) GetOwnerOk() (*NonEmployeeIdnUserRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Owner, true
}

// SetOwner sets field value
func (o *NonEmployeeSourceRequestBody) SetOwner(v NonEmployeeIdnUserRequest) {
	o.Owner = v
}

// GetManagementWorkgroup returns the ManagementWorkgroup field value if set, zero value otherwise.
func (o *NonEmployeeSourceRequestBody) GetManagementWorkgroup() string {
	if o == nil || isNil(o.ManagementWorkgroup) {
		var ret string
		return ret
	}
	return *o.ManagementWorkgroup
}

// GetManagementWorkgroupOk returns a tuple with the ManagementWorkgroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeSourceRequestBody) GetManagementWorkgroupOk() (*string, bool) {
	if o == nil || isNil(o.ManagementWorkgroup) {
		return nil, false
	}
	return o.ManagementWorkgroup, true
}

// HasManagementWorkgroup returns a boolean if a field has been set.
func (o *NonEmployeeSourceRequestBody) HasManagementWorkgroup() bool {
	if o != nil && !isNil(o.ManagementWorkgroup) {
		return true
	}

	return false
}

// SetManagementWorkgroup gets a reference to the given string and assigns it to the ManagementWorkgroup field.
func (o *NonEmployeeSourceRequestBody) SetManagementWorkgroup(v string) {
	o.ManagementWorkgroup = &v
}

// GetApprovers returns the Approvers field value if set, zero value otherwise.
func (o *NonEmployeeSourceRequestBody) GetApprovers() []NonEmployeeIdnUserRequest {
	if o == nil || isNil(o.Approvers) {
		var ret []NonEmployeeIdnUserRequest
		return ret
	}
	return o.Approvers
}

// GetApproversOk returns a tuple with the Approvers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeSourceRequestBody) GetApproversOk() ([]NonEmployeeIdnUserRequest, bool) {
	if o == nil || isNil(o.Approvers) {
		return nil, false
	}
	return o.Approvers, true
}

// HasApprovers returns a boolean if a field has been set.
func (o *NonEmployeeSourceRequestBody) HasApprovers() bool {
	if o != nil && !isNil(o.Approvers) {
		return true
	}

	return false
}

// SetApprovers gets a reference to the given []NonEmployeeIdnUserRequest and assigns it to the Approvers field.
func (o *NonEmployeeSourceRequestBody) SetApprovers(v []NonEmployeeIdnUserRequest) {
	o.Approvers = v
}

// GetAccountManagers returns the AccountManagers field value if set, zero value otherwise.
func (o *NonEmployeeSourceRequestBody) GetAccountManagers() []NonEmployeeIdnUserRequest {
	if o == nil || isNil(o.AccountManagers) {
		var ret []NonEmployeeIdnUserRequest
		return ret
	}
	return o.AccountManagers
}

// GetAccountManagersOk returns a tuple with the AccountManagers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeSourceRequestBody) GetAccountManagersOk() ([]NonEmployeeIdnUserRequest, bool) {
	if o == nil || isNil(o.AccountManagers) {
		return nil, false
	}
	return o.AccountManagers, true
}

// HasAccountManagers returns a boolean if a field has been set.
func (o *NonEmployeeSourceRequestBody) HasAccountManagers() bool {
	if o != nil && !isNil(o.AccountManagers) {
		return true
	}

	return false
}

// SetAccountManagers gets a reference to the given []NonEmployeeIdnUserRequest and assigns it to the AccountManagers field.
func (o *NonEmployeeSourceRequestBody) SetAccountManagers(v []NonEmployeeIdnUserRequest) {
	o.AccountManagers = v
}

func (o NonEmployeeSourceRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NonEmployeeSourceRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["owner"] = o.Owner
	if !isNil(o.ManagementWorkgroup) {
		toSerialize["managementWorkgroup"] = o.ManagementWorkgroup
	}
	if !isNil(o.Approvers) {
		toSerialize["approvers"] = o.Approvers
	}
	if !isNil(o.AccountManagers) {
		toSerialize["accountManagers"] = o.AccountManagers
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NonEmployeeSourceRequestBody) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"description",
		"owner",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varNonEmployeeSourceRequestBody := _NonEmployeeSourceRequestBody{}

	if err = json.Unmarshal(bytes, &varNonEmployeeSourceRequestBody); err == nil {
	*o = NonEmployeeSourceRequestBody(varNonEmployeeSourceRequestBody)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "owner")
		delete(additionalProperties, "managementWorkgroup")
		delete(additionalProperties, "approvers")
		delete(additionalProperties, "accountManagers")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNonEmployeeSourceRequestBody struct {
	value *NonEmployeeSourceRequestBody
	isSet bool
}

func (v NullableNonEmployeeSourceRequestBody) Get() *NonEmployeeSourceRequestBody {
	return v.value
}

func (v *NullableNonEmployeeSourceRequestBody) Set(val *NonEmployeeSourceRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableNonEmployeeSourceRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableNonEmployeeSourceRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNonEmployeeSourceRequestBody(val *NonEmployeeSourceRequestBody) *NullableNonEmployeeSourceRequestBody {
	return &NullableNonEmployeeSourceRequestBody{value: val, isSet: true}
}

func (v NullableNonEmployeeSourceRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNonEmployeeSourceRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


