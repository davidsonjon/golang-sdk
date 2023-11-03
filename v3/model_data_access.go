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

// checks if the DataAccess type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataAccess{}

// DataAccess DAS data for the entitlement
type DataAccess struct {
	// List of classification policies that apply to resources the entitlement \\ groups has access to
	Policies []DataAccessPoliciesInner `json:"policies,omitempty"`
	// List of classification categories that apply to resources the entitlement \\ groups has access to
	Categories []DataAccessCategoriesInner `json:"categories,omitempty"`
	ImpactScore *DataAccessImpactScore `json:"impactScore,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DataAccess DataAccess

// NewDataAccess instantiates a new DataAccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataAccess() *DataAccess {
	this := DataAccess{}
	return &this
}

// NewDataAccessWithDefaults instantiates a new DataAccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataAccessWithDefaults() *DataAccess {
	this := DataAccess{}
	return &this
}

// GetPolicies returns the Policies field value if set, zero value otherwise.
func (o *DataAccess) GetPolicies() []DataAccessPoliciesInner {
	if o == nil || isNil(o.Policies) {
		var ret []DataAccessPoliciesInner
		return ret
	}
	return o.Policies
}

// GetPoliciesOk returns a tuple with the Policies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataAccess) GetPoliciesOk() ([]DataAccessPoliciesInner, bool) {
	if o == nil || isNil(o.Policies) {
		return nil, false
	}
	return o.Policies, true
}

// HasPolicies returns a boolean if a field has been set.
func (o *DataAccess) HasPolicies() bool {
	if o != nil && !isNil(o.Policies) {
		return true
	}

	return false
}

// SetPolicies gets a reference to the given []DataAccessPoliciesInner and assigns it to the Policies field.
func (o *DataAccess) SetPolicies(v []DataAccessPoliciesInner) {
	o.Policies = v
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *DataAccess) GetCategories() []DataAccessCategoriesInner {
	if o == nil || isNil(o.Categories) {
		var ret []DataAccessCategoriesInner
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataAccess) GetCategoriesOk() ([]DataAccessCategoriesInner, bool) {
	if o == nil || isNil(o.Categories) {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *DataAccess) HasCategories() bool {
	if o != nil && !isNil(o.Categories) {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []DataAccessCategoriesInner and assigns it to the Categories field.
func (o *DataAccess) SetCategories(v []DataAccessCategoriesInner) {
	o.Categories = v
}

// GetImpactScore returns the ImpactScore field value if set, zero value otherwise.
func (o *DataAccess) GetImpactScore() DataAccessImpactScore {
	if o == nil || isNil(o.ImpactScore) {
		var ret DataAccessImpactScore
		return ret
	}
	return *o.ImpactScore
}

// GetImpactScoreOk returns a tuple with the ImpactScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataAccess) GetImpactScoreOk() (*DataAccessImpactScore, bool) {
	if o == nil || isNil(o.ImpactScore) {
		return nil, false
	}
	return o.ImpactScore, true
}

// HasImpactScore returns a boolean if a field has been set.
func (o *DataAccess) HasImpactScore() bool {
	if o != nil && !isNil(o.ImpactScore) {
		return true
	}

	return false
}

// SetImpactScore gets a reference to the given DataAccessImpactScore and assigns it to the ImpactScore field.
func (o *DataAccess) SetImpactScore(v DataAccessImpactScore) {
	o.ImpactScore = &v
}

func (o DataAccess) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataAccess) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Policies) {
		toSerialize["policies"] = o.Policies
	}
	if !isNil(o.Categories) {
		toSerialize["categories"] = o.Categories
	}
	if !isNil(o.ImpactScore) {
		toSerialize["impactScore"] = o.ImpactScore
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DataAccess) UnmarshalJSON(bytes []byte) (err error) {
	varDataAccess := _DataAccess{}

	if err = json.Unmarshal(bytes, &varDataAccess); err == nil {
		*o = DataAccess(varDataAccess)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "policies")
		delete(additionalProperties, "categories")
		delete(additionalProperties, "impactScore")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDataAccess struct {
	value *DataAccess
	isSet bool
}

func (v NullableDataAccess) Get() *DataAccess {
	return v.value
}

func (v *NullableDataAccess) Set(val *DataAccess) {
	v.value = val
	v.isSet = true
}

func (v NullableDataAccess) IsSet() bool {
	return v.isSet
}

func (v *NullableDataAccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataAccess(val *DataAccess) *NullableDataAccess {
	return &NullableDataAccess{value: val, isSet: true}
}

func (v NullableDataAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataAccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


