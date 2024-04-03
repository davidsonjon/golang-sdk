/*
Identity Security Cloud V3 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v3

import (
	"encoding/json"
	"fmt"
)

// checks if the InnerHit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InnerHit{}

// InnerHit Inner Hit query object that will cause the specified nested type to be returned as the result matching the supplied query.
type InnerHit struct {
	// The search query using the Elasticsearch [Query String Query](https://www.elastic.co/guide/en/elasticsearch/reference/5.2/query-dsl-query-string-query.html#query-string) syntax from the Query DSL extended by SailPoint to support Nested queries.
	Query string `json:"query"`
	// The nested type to use in the inner hits query.  The nested type [Nested Type](https://www.elastic.co/guide/en/elasticsearch/reference/current/nested.html) refers to a document \"nested\" within another document. For example, an identity can have nested documents for access, accounts, and apps.
	Type string `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _InnerHit InnerHit

// NewInnerHit instantiates a new InnerHit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInnerHit(query string, type_ string) *InnerHit {
	this := InnerHit{}
	this.Query = query
	this.Type = type_
	return &this
}

// NewInnerHitWithDefaults instantiates a new InnerHit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInnerHitWithDefaults() *InnerHit {
	this := InnerHit{}
	return &this
}

// GetQuery returns the Query field value
func (o *InnerHit) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *InnerHit) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *InnerHit) SetQuery(v string) {
	o.Query = v
}

// GetType returns the Type field value
func (o *InnerHit) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InnerHit) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InnerHit) SetType(v string) {
	o.Type = v
}

func (o InnerHit) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InnerHit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["query"] = o.Query
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InnerHit) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"query",
		"type",
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

	varInnerHit := _InnerHit{}

	if err = json.Unmarshal(bytes, &varInnerHit); err == nil {
	*o = InnerHit(varInnerHit)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "query")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInnerHit struct {
	value *InnerHit
	isSet bool
}

func (v NullableInnerHit) Get() *InnerHit {
	return v.value
}

func (v *NullableInnerHit) Set(val *InnerHit) {
	v.value = val
	v.isSet = true
}

func (v NullableInnerHit) IsSet() bool {
	return v.isSet
}

func (v *NullableInnerHit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInnerHit(val *InnerHit) *NullableInnerHit {
	return &NullableInnerHit{value: val, isSet: true}
}

func (v NullableInnerHit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInnerHit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


