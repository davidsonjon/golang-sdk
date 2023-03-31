/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package beta

import (
	"encoding/json"
)

// checks if the RemediationItemDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemediationItemDetails{}

// RemediationItemDetails struct for RemediationItemDetails
type RemediationItemDetails struct {
	// The ID of the certification
	Id *string `json:"id,omitempty"`
	// The ID of the certification target
	TargetId *string `json:"targetId,omitempty"`
	// The name of the certification target
	TargetName *string `json:"targetName,omitempty"`
	// The display name of the certification target
	TargetDisplayName *string `json:"targetDisplayName,omitempty"`
	// The name of the application/source
	ApplicationName *string `json:"applicationName,omitempty"`
	// The name of the attribute being certified
	AttributeName *string `json:"attributeName,omitempty"`
	// The operation of the certification on the attribute
	AttributeOperation *string `json:"attributeOperation,omitempty"`
	// The value of the attribute being certified
	AttributeValue *string `json:"attributeValue,omitempty"`
	// The native identity of the target
	NativeIdentity *string `json:"nativeIdentity,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RemediationItemDetails RemediationItemDetails

// NewRemediationItemDetails instantiates a new RemediationItemDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemediationItemDetails() *RemediationItemDetails {
	this := RemediationItemDetails{}
	return &this
}

// NewRemediationItemDetailsWithDefaults instantiates a new RemediationItemDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemediationItemDetailsWithDefaults() *RemediationItemDetails {
	this := RemediationItemDetails{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RemediationItemDetails) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationItemDetails) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RemediationItemDetails) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RemediationItemDetails) SetId(v string) {
	o.Id = &v
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *RemediationItemDetails) GetTargetId() string {
	if o == nil || isNil(o.TargetId) {
		var ret string
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationItemDetails) GetTargetIdOk() (*string, bool) {
	if o == nil || isNil(o.TargetId) {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *RemediationItemDetails) HasTargetId() bool {
	if o != nil && !isNil(o.TargetId) {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given string and assigns it to the TargetId field.
func (o *RemediationItemDetails) SetTargetId(v string) {
	o.TargetId = &v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *RemediationItemDetails) GetTargetName() string {
	if o == nil || isNil(o.TargetName) {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationItemDetails) GetTargetNameOk() (*string, bool) {
	if o == nil || isNil(o.TargetName) {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *RemediationItemDetails) HasTargetName() bool {
	if o != nil && !isNil(o.TargetName) {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *RemediationItemDetails) SetTargetName(v string) {
	o.TargetName = &v
}

// GetTargetDisplayName returns the TargetDisplayName field value if set, zero value otherwise.
func (o *RemediationItemDetails) GetTargetDisplayName() string {
	if o == nil || isNil(o.TargetDisplayName) {
		var ret string
		return ret
	}
	return *o.TargetDisplayName
}

// GetTargetDisplayNameOk returns a tuple with the TargetDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationItemDetails) GetTargetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.TargetDisplayName) {
		return nil, false
	}
	return o.TargetDisplayName, true
}

// HasTargetDisplayName returns a boolean if a field has been set.
func (o *RemediationItemDetails) HasTargetDisplayName() bool {
	if o != nil && !isNil(o.TargetDisplayName) {
		return true
	}

	return false
}

// SetTargetDisplayName gets a reference to the given string and assigns it to the TargetDisplayName field.
func (o *RemediationItemDetails) SetTargetDisplayName(v string) {
	o.TargetDisplayName = &v
}

// GetApplicationName returns the ApplicationName field value if set, zero value otherwise.
func (o *RemediationItemDetails) GetApplicationName() string {
	if o == nil || isNil(o.ApplicationName) {
		var ret string
		return ret
	}
	return *o.ApplicationName
}

// GetApplicationNameOk returns a tuple with the ApplicationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationItemDetails) GetApplicationNameOk() (*string, bool) {
	if o == nil || isNil(o.ApplicationName) {
		return nil, false
	}
	return o.ApplicationName, true
}

// HasApplicationName returns a boolean if a field has been set.
func (o *RemediationItemDetails) HasApplicationName() bool {
	if o != nil && !isNil(o.ApplicationName) {
		return true
	}

	return false
}

// SetApplicationName gets a reference to the given string and assigns it to the ApplicationName field.
func (o *RemediationItemDetails) SetApplicationName(v string) {
	o.ApplicationName = &v
}

// GetAttributeName returns the AttributeName field value if set, zero value otherwise.
func (o *RemediationItemDetails) GetAttributeName() string {
	if o == nil || isNil(o.AttributeName) {
		var ret string
		return ret
	}
	return *o.AttributeName
}

// GetAttributeNameOk returns a tuple with the AttributeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationItemDetails) GetAttributeNameOk() (*string, bool) {
	if o == nil || isNil(o.AttributeName) {
		return nil, false
	}
	return o.AttributeName, true
}

// HasAttributeName returns a boolean if a field has been set.
func (o *RemediationItemDetails) HasAttributeName() bool {
	if o != nil && !isNil(o.AttributeName) {
		return true
	}

	return false
}

// SetAttributeName gets a reference to the given string and assigns it to the AttributeName field.
func (o *RemediationItemDetails) SetAttributeName(v string) {
	o.AttributeName = &v
}

// GetAttributeOperation returns the AttributeOperation field value if set, zero value otherwise.
func (o *RemediationItemDetails) GetAttributeOperation() string {
	if o == nil || isNil(o.AttributeOperation) {
		var ret string
		return ret
	}
	return *o.AttributeOperation
}

// GetAttributeOperationOk returns a tuple with the AttributeOperation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationItemDetails) GetAttributeOperationOk() (*string, bool) {
	if o == nil || isNil(o.AttributeOperation) {
		return nil, false
	}
	return o.AttributeOperation, true
}

// HasAttributeOperation returns a boolean if a field has been set.
func (o *RemediationItemDetails) HasAttributeOperation() bool {
	if o != nil && !isNil(o.AttributeOperation) {
		return true
	}

	return false
}

// SetAttributeOperation gets a reference to the given string and assigns it to the AttributeOperation field.
func (o *RemediationItemDetails) SetAttributeOperation(v string) {
	o.AttributeOperation = &v
}

// GetAttributeValue returns the AttributeValue field value if set, zero value otherwise.
func (o *RemediationItemDetails) GetAttributeValue() string {
	if o == nil || isNil(o.AttributeValue) {
		var ret string
		return ret
	}
	return *o.AttributeValue
}

// GetAttributeValueOk returns a tuple with the AttributeValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationItemDetails) GetAttributeValueOk() (*string, bool) {
	if o == nil || isNil(o.AttributeValue) {
		return nil, false
	}
	return o.AttributeValue, true
}

// HasAttributeValue returns a boolean if a field has been set.
func (o *RemediationItemDetails) HasAttributeValue() bool {
	if o != nil && !isNil(o.AttributeValue) {
		return true
	}

	return false
}

// SetAttributeValue gets a reference to the given string and assigns it to the AttributeValue field.
func (o *RemediationItemDetails) SetAttributeValue(v string) {
	o.AttributeValue = &v
}

// GetNativeIdentity returns the NativeIdentity field value if set, zero value otherwise.
func (o *RemediationItemDetails) GetNativeIdentity() string {
	if o == nil || isNil(o.NativeIdentity) {
		var ret string
		return ret
	}
	return *o.NativeIdentity
}

// GetNativeIdentityOk returns a tuple with the NativeIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationItemDetails) GetNativeIdentityOk() (*string, bool) {
	if o == nil || isNil(o.NativeIdentity) {
		return nil, false
	}
	return o.NativeIdentity, true
}

// HasNativeIdentity returns a boolean if a field has been set.
func (o *RemediationItemDetails) HasNativeIdentity() bool {
	if o != nil && !isNil(o.NativeIdentity) {
		return true
	}

	return false
}

// SetNativeIdentity gets a reference to the given string and assigns it to the NativeIdentity field.
func (o *RemediationItemDetails) SetNativeIdentity(v string) {
	o.NativeIdentity = &v
}

func (o RemediationItemDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemediationItemDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.TargetId) {
		toSerialize["targetId"] = o.TargetId
	}
	if !isNil(o.TargetName) {
		toSerialize["targetName"] = o.TargetName
	}
	if !isNil(o.TargetDisplayName) {
		toSerialize["targetDisplayName"] = o.TargetDisplayName
	}
	if !isNil(o.ApplicationName) {
		toSerialize["applicationName"] = o.ApplicationName
	}
	if !isNil(o.AttributeName) {
		toSerialize["attributeName"] = o.AttributeName
	}
	if !isNil(o.AttributeOperation) {
		toSerialize["attributeOperation"] = o.AttributeOperation
	}
	if !isNil(o.AttributeValue) {
		toSerialize["attributeValue"] = o.AttributeValue
	}
	if !isNil(o.NativeIdentity) {
		toSerialize["nativeIdentity"] = o.NativeIdentity
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RemediationItemDetails) UnmarshalJSON(bytes []byte) (err error) {
	varRemediationItemDetails := _RemediationItemDetails{}

	if err = json.Unmarshal(bytes, &varRemediationItemDetails); err == nil {
		*o = RemediationItemDetails(varRemediationItemDetails)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "targetId")
		delete(additionalProperties, "targetName")
		delete(additionalProperties, "targetDisplayName")
		delete(additionalProperties, "applicationName")
		delete(additionalProperties, "attributeName")
		delete(additionalProperties, "attributeOperation")
		delete(additionalProperties, "attributeValue")
		delete(additionalProperties, "nativeIdentity")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRemediationItemDetails struct {
	value *RemediationItemDetails
	isSet bool
}

func (v NullableRemediationItemDetails) Get() *RemediationItemDetails {
	return v.value
}

func (v *NullableRemediationItemDetails) Set(val *RemediationItemDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableRemediationItemDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableRemediationItemDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemediationItemDetails(val *RemediationItemDetails) *NullableRemediationItemDetails {
	return &NullableRemediationItemDetails{value: val, isSet: true}
}

func (v NullableRemediationItemDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemediationItemDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


