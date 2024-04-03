/*
Identity Security Cloud Beta API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
	"fmt"
)

// checks if the ConnectorRuleValidationResponseDetailsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectorRuleValidationResponseDetailsInner{}

// ConnectorRuleValidationResponseDetailsInner CodeErrorDetail
type ConnectorRuleValidationResponseDetailsInner struct {
	// The line number where the issue occurred
	Line int32 `json:"line"`
	// the column number where the issue occurred
	Column int32 `json:"column"`
	// a description of the issue in the code
	Messsage *string `json:"messsage,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConnectorRuleValidationResponseDetailsInner ConnectorRuleValidationResponseDetailsInner

// NewConnectorRuleValidationResponseDetailsInner instantiates a new ConnectorRuleValidationResponseDetailsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorRuleValidationResponseDetailsInner(line int32, column int32) *ConnectorRuleValidationResponseDetailsInner {
	this := ConnectorRuleValidationResponseDetailsInner{}
	this.Line = line
	this.Column = column
	return &this
}

// NewConnectorRuleValidationResponseDetailsInnerWithDefaults instantiates a new ConnectorRuleValidationResponseDetailsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorRuleValidationResponseDetailsInnerWithDefaults() *ConnectorRuleValidationResponseDetailsInner {
	this := ConnectorRuleValidationResponseDetailsInner{}
	return &this
}

// GetLine returns the Line field value
func (o *ConnectorRuleValidationResponseDetailsInner) GetLine() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Line
}

// GetLineOk returns a tuple with the Line field value
// and a boolean to check if the value has been set.
func (o *ConnectorRuleValidationResponseDetailsInner) GetLineOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Line, true
}

// SetLine sets field value
func (o *ConnectorRuleValidationResponseDetailsInner) SetLine(v int32) {
	o.Line = v
}

// GetColumn returns the Column field value
func (o *ConnectorRuleValidationResponseDetailsInner) GetColumn() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Column
}

// GetColumnOk returns a tuple with the Column field value
// and a boolean to check if the value has been set.
func (o *ConnectorRuleValidationResponseDetailsInner) GetColumnOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Column, true
}

// SetColumn sets field value
func (o *ConnectorRuleValidationResponseDetailsInner) SetColumn(v int32) {
	o.Column = v
}

// GetMesssage returns the Messsage field value if set, zero value otherwise.
func (o *ConnectorRuleValidationResponseDetailsInner) GetMesssage() string {
	if o == nil || isNil(o.Messsage) {
		var ret string
		return ret
	}
	return *o.Messsage
}

// GetMesssageOk returns a tuple with the Messsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorRuleValidationResponseDetailsInner) GetMesssageOk() (*string, bool) {
	if o == nil || isNil(o.Messsage) {
		return nil, false
	}
	return o.Messsage, true
}

// HasMesssage returns a boolean if a field has been set.
func (o *ConnectorRuleValidationResponseDetailsInner) HasMesssage() bool {
	if o != nil && !isNil(o.Messsage) {
		return true
	}

	return false
}

// SetMesssage gets a reference to the given string and assigns it to the Messsage field.
func (o *ConnectorRuleValidationResponseDetailsInner) SetMesssage(v string) {
	o.Messsage = &v
}

func (o ConnectorRuleValidationResponseDetailsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectorRuleValidationResponseDetailsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["line"] = o.Line
	toSerialize["column"] = o.Column
	if !isNil(o.Messsage) {
		toSerialize["messsage"] = o.Messsage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConnectorRuleValidationResponseDetailsInner) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"line",
		"column",
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

	varConnectorRuleValidationResponseDetailsInner := _ConnectorRuleValidationResponseDetailsInner{}

	if err = json.Unmarshal(bytes, &varConnectorRuleValidationResponseDetailsInner); err == nil {
	*o = ConnectorRuleValidationResponseDetailsInner(varConnectorRuleValidationResponseDetailsInner)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "line")
		delete(additionalProperties, "column")
		delete(additionalProperties, "messsage")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConnectorRuleValidationResponseDetailsInner struct {
	value *ConnectorRuleValidationResponseDetailsInner
	isSet bool
}

func (v NullableConnectorRuleValidationResponseDetailsInner) Get() *ConnectorRuleValidationResponseDetailsInner {
	return v.value
}

func (v *NullableConnectorRuleValidationResponseDetailsInner) Set(val *ConnectorRuleValidationResponseDetailsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorRuleValidationResponseDetailsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorRuleValidationResponseDetailsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorRuleValidationResponseDetailsInner(val *ConnectorRuleValidationResponseDetailsInner) *NullableConnectorRuleValidationResponseDetailsInner {
	return &NullableConnectorRuleValidationResponseDetailsInner{value: val, isSet: true}
}

func (v NullableConnectorRuleValidationResponseDetailsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorRuleValidationResponseDetailsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


