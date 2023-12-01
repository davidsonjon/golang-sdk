/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
)

// checks if the AccountToggleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountToggleRequest{}

// AccountToggleRequest Request used for account enable/disable
type AccountToggleRequest struct {
	// If set, an external process validates that the user wants to proceed with this request.
	ExternalVerificationId *string `json:"externalVerificationId,omitempty"`
	// If set, provisioning updates the account attribute at the source.   This option is used when the account is not synced to ensure the attribute is updated. Providing 'true' for an unlocked account will add and process 'Unlock' operation by the workflow.
	ForceProvisioning *bool `json:"forceProvisioning,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountToggleRequest AccountToggleRequest

// NewAccountToggleRequest instantiates a new AccountToggleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountToggleRequest() *AccountToggleRequest {
	this := AccountToggleRequest{}
	return &this
}

// NewAccountToggleRequestWithDefaults instantiates a new AccountToggleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountToggleRequestWithDefaults() *AccountToggleRequest {
	this := AccountToggleRequest{}
	return &this
}

// GetExternalVerificationId returns the ExternalVerificationId field value if set, zero value otherwise.
func (o *AccountToggleRequest) GetExternalVerificationId() string {
	if o == nil || isNil(o.ExternalVerificationId) {
		var ret string
		return ret
	}
	return *o.ExternalVerificationId
}

// GetExternalVerificationIdOk returns a tuple with the ExternalVerificationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountToggleRequest) GetExternalVerificationIdOk() (*string, bool) {
	if o == nil || isNil(o.ExternalVerificationId) {
		return nil, false
	}
	return o.ExternalVerificationId, true
}

// HasExternalVerificationId returns a boolean if a field has been set.
func (o *AccountToggleRequest) HasExternalVerificationId() bool {
	if o != nil && !isNil(o.ExternalVerificationId) {
		return true
	}

	return false
}

// SetExternalVerificationId gets a reference to the given string and assigns it to the ExternalVerificationId field.
func (o *AccountToggleRequest) SetExternalVerificationId(v string) {
	o.ExternalVerificationId = &v
}

// GetForceProvisioning returns the ForceProvisioning field value if set, zero value otherwise.
func (o *AccountToggleRequest) GetForceProvisioning() bool {
	if o == nil || isNil(o.ForceProvisioning) {
		var ret bool
		return ret
	}
	return *o.ForceProvisioning
}

// GetForceProvisioningOk returns a tuple with the ForceProvisioning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountToggleRequest) GetForceProvisioningOk() (*bool, bool) {
	if o == nil || isNil(o.ForceProvisioning) {
		return nil, false
	}
	return o.ForceProvisioning, true
}

// HasForceProvisioning returns a boolean if a field has been set.
func (o *AccountToggleRequest) HasForceProvisioning() bool {
	if o != nil && !isNil(o.ForceProvisioning) {
		return true
	}

	return false
}

// SetForceProvisioning gets a reference to the given bool and assigns it to the ForceProvisioning field.
func (o *AccountToggleRequest) SetForceProvisioning(v bool) {
	o.ForceProvisioning = &v
}

func (o AccountToggleRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountToggleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ExternalVerificationId) {
		toSerialize["externalVerificationId"] = o.ExternalVerificationId
	}
	if !isNil(o.ForceProvisioning) {
		toSerialize["forceProvisioning"] = o.ForceProvisioning
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountToggleRequest) UnmarshalJSON(bytes []byte) (err error) {
	varAccountToggleRequest := _AccountToggleRequest{}

	if err = json.Unmarshal(bytes, &varAccountToggleRequest); err == nil {
	*o = AccountToggleRequest(varAccountToggleRequest)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "externalVerificationId")
		delete(additionalProperties, "forceProvisioning")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountToggleRequest struct {
	value *AccountToggleRequest
	isSet bool
}

func (v NullableAccountToggleRequest) Get() *AccountToggleRequest {
	return v.value
}

func (v *NullableAccountToggleRequest) Set(val *AccountToggleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountToggleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountToggleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountToggleRequest(val *AccountToggleRequest) *NullableAccountToggleRequest {
	return &NullableAccountToggleRequest{value: val, isSet: true}
}

func (v NullableAccountToggleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountToggleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


