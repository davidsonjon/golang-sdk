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

// checks if the Static type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Static{}

// Static struct for Static
type Static struct {
	// This must evaluate to a JSON string, either through a fixed value or through conditional logic using the Apache Velocity Template Language.
	Values string `json:"values"`
	// A value that indicates whether the transform logic should be re-evaluated every evening as part of the identity refresh process
	RequiresPeriodicRefresh *bool `json:"requiresPeriodicRefresh,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Static Static

// NewStatic instantiates a new Static object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatic(values string) *Static {
	this := Static{}
	this.Values = values
	var requiresPeriodicRefresh bool = false
	this.RequiresPeriodicRefresh = &requiresPeriodicRefresh
	return &this
}

// NewStaticWithDefaults instantiates a new Static object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStaticWithDefaults() *Static {
	this := Static{}
	var requiresPeriodicRefresh bool = false
	this.RequiresPeriodicRefresh = &requiresPeriodicRefresh
	return &this
}

// GetValues returns the Values field value
func (o *Static) GetValues() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *Static) GetValuesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Values, true
}

// SetValues sets field value
func (o *Static) SetValues(v string) {
	o.Values = v
}

// GetRequiresPeriodicRefresh returns the RequiresPeriodicRefresh field value if set, zero value otherwise.
func (o *Static) GetRequiresPeriodicRefresh() bool {
	if o == nil || isNil(o.RequiresPeriodicRefresh) {
		var ret bool
		return ret
	}
	return *o.RequiresPeriodicRefresh
}

// GetRequiresPeriodicRefreshOk returns a tuple with the RequiresPeriodicRefresh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Static) GetRequiresPeriodicRefreshOk() (*bool, bool) {
	if o == nil || isNil(o.RequiresPeriodicRefresh) {
		return nil, false
	}
	return o.RequiresPeriodicRefresh, true
}

// HasRequiresPeriodicRefresh returns a boolean if a field has been set.
func (o *Static) HasRequiresPeriodicRefresh() bool {
	if o != nil && !isNil(o.RequiresPeriodicRefresh) {
		return true
	}

	return false
}

// SetRequiresPeriodicRefresh gets a reference to the given bool and assigns it to the RequiresPeriodicRefresh field.
func (o *Static) SetRequiresPeriodicRefresh(v bool) {
	o.RequiresPeriodicRefresh = &v
}

func (o Static) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Static) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["values"] = o.Values
	if !isNil(o.RequiresPeriodicRefresh) {
		toSerialize["requiresPeriodicRefresh"] = o.RequiresPeriodicRefresh
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Static) UnmarshalJSON(bytes []byte) (err error) {
	varStatic := _Static{}

	if err = json.Unmarshal(bytes, &varStatic); err == nil {
		*o = Static(varStatic)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "values")
		delete(additionalProperties, "requiresPeriodicRefresh")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStatic struct {
	value *Static
	isSet bool
}

func (v NullableStatic) Get() *Static {
	return v.value
}

func (v *NullableStatic) Set(val *Static) {
	v.value = val
	v.isSet = true
}

func (v NullableStatic) IsSet() bool {
	return v.isSet
}

func (v *NullableStatic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatic(val *Static) *NullableStatic {
	return &NullableStatic{value: val, isSet: true}
}

func (v NullableStatic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


