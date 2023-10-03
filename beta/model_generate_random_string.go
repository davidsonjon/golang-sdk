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

// checks if the GenerateRandomString type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenerateRandomString{}

// GenerateRandomString struct for GenerateRandomString
type GenerateRandomString struct {
	// This must always be set to \"Cloud Services Deployment Utility\"
	Name string `json:"name"`
	// The operation to perform `generateRandomString`
	Operation string `json:"operation"`
	// This must be either \"true\" or \"false\" to indicate whether the generator logic should include numbers
	IncludeNumbers bool `json:"includeNumbers"`
	// This must be either \"true\" or \"false\" to indicate whether the generator logic should include special characters
	IncludeSpecialChars bool `json:"includeSpecialChars"`
	// This specifies how long the randomly generated string needs to be   >NOTE Due to identity attribute data constraints, the maximum allowable value is 450 characters 
	Length string `json:"length"`
	// A value that indicates whether the transform logic should be re-evaluated every evening as part of the identity refresh process
	RequiresPeriodicRefresh *bool `json:"requiresPeriodicRefresh,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GenerateRandomString GenerateRandomString

// NewGenerateRandomString instantiates a new GenerateRandomString object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerateRandomString(name string, operation string, includeNumbers bool, includeSpecialChars bool, length string) *GenerateRandomString {
	this := GenerateRandomString{}
	this.Name = name
	this.Operation = operation
	this.IncludeNumbers = includeNumbers
	this.IncludeSpecialChars = includeSpecialChars
	this.Length = length
	return &this
}

// NewGenerateRandomStringWithDefaults instantiates a new GenerateRandomString object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateRandomStringWithDefaults() *GenerateRandomString {
	this := GenerateRandomString{}
	return &this
}

// GetName returns the Name field value
func (o *GenerateRandomString) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GenerateRandomString) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GenerateRandomString) SetName(v string) {
	o.Name = v
}

// GetOperation returns the Operation field value
func (o *GenerateRandomString) GetOperation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *GenerateRandomString) GetOperationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operation, true
}

// SetOperation sets field value
func (o *GenerateRandomString) SetOperation(v string) {
	o.Operation = v
}

// GetIncludeNumbers returns the IncludeNumbers field value
func (o *GenerateRandomString) GetIncludeNumbers() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IncludeNumbers
}

// GetIncludeNumbersOk returns a tuple with the IncludeNumbers field value
// and a boolean to check if the value has been set.
func (o *GenerateRandomString) GetIncludeNumbersOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncludeNumbers, true
}

// SetIncludeNumbers sets field value
func (o *GenerateRandomString) SetIncludeNumbers(v bool) {
	o.IncludeNumbers = v
}

// GetIncludeSpecialChars returns the IncludeSpecialChars field value
func (o *GenerateRandomString) GetIncludeSpecialChars() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IncludeSpecialChars
}

// GetIncludeSpecialCharsOk returns a tuple with the IncludeSpecialChars field value
// and a boolean to check if the value has been set.
func (o *GenerateRandomString) GetIncludeSpecialCharsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncludeSpecialChars, true
}

// SetIncludeSpecialChars sets field value
func (o *GenerateRandomString) SetIncludeSpecialChars(v bool) {
	o.IncludeSpecialChars = v
}

// GetLength returns the Length field value
func (o *GenerateRandomString) GetLength() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Length
}

// GetLengthOk returns a tuple with the Length field value
// and a boolean to check if the value has been set.
func (o *GenerateRandomString) GetLengthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Length, true
}

// SetLength sets field value
func (o *GenerateRandomString) SetLength(v string) {
	o.Length = v
}

// GetRequiresPeriodicRefresh returns the RequiresPeriodicRefresh field value if set, zero value otherwise.
func (o *GenerateRandomString) GetRequiresPeriodicRefresh() bool {
	if o == nil || isNil(o.RequiresPeriodicRefresh) {
		var ret bool
		return ret
	}
	return *o.RequiresPeriodicRefresh
}

// GetRequiresPeriodicRefreshOk returns a tuple with the RequiresPeriodicRefresh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateRandomString) GetRequiresPeriodicRefreshOk() (*bool, bool) {
	if o == nil || isNil(o.RequiresPeriodicRefresh) {
		return nil, false
	}
	return o.RequiresPeriodicRefresh, true
}

// HasRequiresPeriodicRefresh returns a boolean if a field has been set.
func (o *GenerateRandomString) HasRequiresPeriodicRefresh() bool {
	if o != nil && !isNil(o.RequiresPeriodicRefresh) {
		return true
	}

	return false
}

// SetRequiresPeriodicRefresh gets a reference to the given bool and assigns it to the RequiresPeriodicRefresh field.
func (o *GenerateRandomString) SetRequiresPeriodicRefresh(v bool) {
	o.RequiresPeriodicRefresh = &v
}

func (o GenerateRandomString) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GenerateRandomString) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["operation"] = o.Operation
	toSerialize["includeNumbers"] = o.IncludeNumbers
	toSerialize["includeSpecialChars"] = o.IncludeSpecialChars
	toSerialize["length"] = o.Length
	if !isNil(o.RequiresPeriodicRefresh) {
		toSerialize["requiresPeriodicRefresh"] = o.RequiresPeriodicRefresh
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GenerateRandomString) UnmarshalJSON(bytes []byte) (err error) {
	varGenerateRandomString := _GenerateRandomString{}

	if err = json.Unmarshal(bytes, &varGenerateRandomString); err == nil {
		*o = GenerateRandomString(varGenerateRandomString)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "operation")
		delete(additionalProperties, "includeNumbers")
		delete(additionalProperties, "includeSpecialChars")
		delete(additionalProperties, "length")
		delete(additionalProperties, "requiresPeriodicRefresh")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGenerateRandomString struct {
	value *GenerateRandomString
	isSet bool
}

func (v NullableGenerateRandomString) Get() *GenerateRandomString {
	return v.value
}

func (v *NullableGenerateRandomString) Set(val *GenerateRandomString) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateRandomString) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateRandomString) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateRandomString(val *GenerateRandomString) *NullableGenerateRandomString {
	return &NullableGenerateRandomString{value: val, isSet: true}
}

func (v NullableGenerateRandomString) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenerateRandomString) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


