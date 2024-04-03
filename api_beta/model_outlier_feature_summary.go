/*
Identity Security Cloud Beta API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
)

// checks if the OutlierFeatureSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OutlierFeatureSummary{}

// OutlierFeatureSummary struct for OutlierFeatureSummary
type OutlierFeatureSummary struct {
	// Contributing feature name
	ContributingFeatureName *string `json:"contributingFeatureName,omitempty"`
	// Identity display name
	IdentityOutlierDisplayName *string `json:"identityOutlierDisplayName,omitempty"`
	OutlierFeatureDisplayValues []OutlierFeatureSummaryOutlierFeatureDisplayValuesInner `json:"outlierFeatureDisplayValues,omitempty"`
	// Definition of the feature
	FeatureDefinition *string `json:"featureDefinition,omitempty"`
	// Detailed explanation of the feature
	FeatureExplanation *string `json:"featureExplanation,omitempty"`
	// outlier's peer identity display name
	PeerDisplayName *string `json:"peerDisplayName,omitempty"`
	// outlier's peer identity id
	PeerIdentityId *string `json:"peerIdentityId,omitempty"`
	// Access Item reference
	AccessItemReference map[string]interface{} `json:"accessItemReference,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OutlierFeatureSummary OutlierFeatureSummary

// NewOutlierFeatureSummary instantiates a new OutlierFeatureSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutlierFeatureSummary() *OutlierFeatureSummary {
	this := OutlierFeatureSummary{}
	return &this
}

// NewOutlierFeatureSummaryWithDefaults instantiates a new OutlierFeatureSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutlierFeatureSummaryWithDefaults() *OutlierFeatureSummary {
	this := OutlierFeatureSummary{}
	return &this
}

// GetContributingFeatureName returns the ContributingFeatureName field value if set, zero value otherwise.
func (o *OutlierFeatureSummary) GetContributingFeatureName() string {
	if o == nil || isNil(o.ContributingFeatureName) {
		var ret string
		return ret
	}
	return *o.ContributingFeatureName
}

// GetContributingFeatureNameOk returns a tuple with the ContributingFeatureName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutlierFeatureSummary) GetContributingFeatureNameOk() (*string, bool) {
	if o == nil || isNil(o.ContributingFeatureName) {
		return nil, false
	}
	return o.ContributingFeatureName, true
}

// HasContributingFeatureName returns a boolean if a field has been set.
func (o *OutlierFeatureSummary) HasContributingFeatureName() bool {
	if o != nil && !isNil(o.ContributingFeatureName) {
		return true
	}

	return false
}

// SetContributingFeatureName gets a reference to the given string and assigns it to the ContributingFeatureName field.
func (o *OutlierFeatureSummary) SetContributingFeatureName(v string) {
	o.ContributingFeatureName = &v
}

// GetIdentityOutlierDisplayName returns the IdentityOutlierDisplayName field value if set, zero value otherwise.
func (o *OutlierFeatureSummary) GetIdentityOutlierDisplayName() string {
	if o == nil || isNil(o.IdentityOutlierDisplayName) {
		var ret string
		return ret
	}
	return *o.IdentityOutlierDisplayName
}

// GetIdentityOutlierDisplayNameOk returns a tuple with the IdentityOutlierDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutlierFeatureSummary) GetIdentityOutlierDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.IdentityOutlierDisplayName) {
		return nil, false
	}
	return o.IdentityOutlierDisplayName, true
}

// HasIdentityOutlierDisplayName returns a boolean if a field has been set.
func (o *OutlierFeatureSummary) HasIdentityOutlierDisplayName() bool {
	if o != nil && !isNil(o.IdentityOutlierDisplayName) {
		return true
	}

	return false
}

// SetIdentityOutlierDisplayName gets a reference to the given string and assigns it to the IdentityOutlierDisplayName field.
func (o *OutlierFeatureSummary) SetIdentityOutlierDisplayName(v string) {
	o.IdentityOutlierDisplayName = &v
}

// GetOutlierFeatureDisplayValues returns the OutlierFeatureDisplayValues field value if set, zero value otherwise.
func (o *OutlierFeatureSummary) GetOutlierFeatureDisplayValues() []OutlierFeatureSummaryOutlierFeatureDisplayValuesInner {
	if o == nil || isNil(o.OutlierFeatureDisplayValues) {
		var ret []OutlierFeatureSummaryOutlierFeatureDisplayValuesInner
		return ret
	}
	return o.OutlierFeatureDisplayValues
}

// GetOutlierFeatureDisplayValuesOk returns a tuple with the OutlierFeatureDisplayValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutlierFeatureSummary) GetOutlierFeatureDisplayValuesOk() ([]OutlierFeatureSummaryOutlierFeatureDisplayValuesInner, bool) {
	if o == nil || isNil(o.OutlierFeatureDisplayValues) {
		return nil, false
	}
	return o.OutlierFeatureDisplayValues, true
}

// HasOutlierFeatureDisplayValues returns a boolean if a field has been set.
func (o *OutlierFeatureSummary) HasOutlierFeatureDisplayValues() bool {
	if o != nil && !isNil(o.OutlierFeatureDisplayValues) {
		return true
	}

	return false
}

// SetOutlierFeatureDisplayValues gets a reference to the given []OutlierFeatureSummaryOutlierFeatureDisplayValuesInner and assigns it to the OutlierFeatureDisplayValues field.
func (o *OutlierFeatureSummary) SetOutlierFeatureDisplayValues(v []OutlierFeatureSummaryOutlierFeatureDisplayValuesInner) {
	o.OutlierFeatureDisplayValues = v
}

// GetFeatureDefinition returns the FeatureDefinition field value if set, zero value otherwise.
func (o *OutlierFeatureSummary) GetFeatureDefinition() string {
	if o == nil || isNil(o.FeatureDefinition) {
		var ret string
		return ret
	}
	return *o.FeatureDefinition
}

// GetFeatureDefinitionOk returns a tuple with the FeatureDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutlierFeatureSummary) GetFeatureDefinitionOk() (*string, bool) {
	if o == nil || isNil(o.FeatureDefinition) {
		return nil, false
	}
	return o.FeatureDefinition, true
}

// HasFeatureDefinition returns a boolean if a field has been set.
func (o *OutlierFeatureSummary) HasFeatureDefinition() bool {
	if o != nil && !isNil(o.FeatureDefinition) {
		return true
	}

	return false
}

// SetFeatureDefinition gets a reference to the given string and assigns it to the FeatureDefinition field.
func (o *OutlierFeatureSummary) SetFeatureDefinition(v string) {
	o.FeatureDefinition = &v
}

// GetFeatureExplanation returns the FeatureExplanation field value if set, zero value otherwise.
func (o *OutlierFeatureSummary) GetFeatureExplanation() string {
	if o == nil || isNil(o.FeatureExplanation) {
		var ret string
		return ret
	}
	return *o.FeatureExplanation
}

// GetFeatureExplanationOk returns a tuple with the FeatureExplanation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutlierFeatureSummary) GetFeatureExplanationOk() (*string, bool) {
	if o == nil || isNil(o.FeatureExplanation) {
		return nil, false
	}
	return o.FeatureExplanation, true
}

// HasFeatureExplanation returns a boolean if a field has been set.
func (o *OutlierFeatureSummary) HasFeatureExplanation() bool {
	if o != nil && !isNil(o.FeatureExplanation) {
		return true
	}

	return false
}

// SetFeatureExplanation gets a reference to the given string and assigns it to the FeatureExplanation field.
func (o *OutlierFeatureSummary) SetFeatureExplanation(v string) {
	o.FeatureExplanation = &v
}

// GetPeerDisplayName returns the PeerDisplayName field value if set, zero value otherwise.
func (o *OutlierFeatureSummary) GetPeerDisplayName() string {
	if o == nil || isNil(o.PeerDisplayName) {
		var ret string
		return ret
	}
	return *o.PeerDisplayName
}

// GetPeerDisplayNameOk returns a tuple with the PeerDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutlierFeatureSummary) GetPeerDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.PeerDisplayName) {
		return nil, false
	}
	return o.PeerDisplayName, true
}

// HasPeerDisplayName returns a boolean if a field has been set.
func (o *OutlierFeatureSummary) HasPeerDisplayName() bool {
	if o != nil && !isNil(o.PeerDisplayName) {
		return true
	}

	return false
}

// SetPeerDisplayName gets a reference to the given string and assigns it to the PeerDisplayName field.
func (o *OutlierFeatureSummary) SetPeerDisplayName(v string) {
	o.PeerDisplayName = &v
}

// GetPeerIdentityId returns the PeerIdentityId field value if set, zero value otherwise.
func (o *OutlierFeatureSummary) GetPeerIdentityId() string {
	if o == nil || isNil(o.PeerIdentityId) {
		var ret string
		return ret
	}
	return *o.PeerIdentityId
}

// GetPeerIdentityIdOk returns a tuple with the PeerIdentityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutlierFeatureSummary) GetPeerIdentityIdOk() (*string, bool) {
	if o == nil || isNil(o.PeerIdentityId) {
		return nil, false
	}
	return o.PeerIdentityId, true
}

// HasPeerIdentityId returns a boolean if a field has been set.
func (o *OutlierFeatureSummary) HasPeerIdentityId() bool {
	if o != nil && !isNil(o.PeerIdentityId) {
		return true
	}

	return false
}

// SetPeerIdentityId gets a reference to the given string and assigns it to the PeerIdentityId field.
func (o *OutlierFeatureSummary) SetPeerIdentityId(v string) {
	o.PeerIdentityId = &v
}

// GetAccessItemReference returns the AccessItemReference field value if set, zero value otherwise.
func (o *OutlierFeatureSummary) GetAccessItemReference() map[string]interface{} {
	if o == nil || isNil(o.AccessItemReference) {
		var ret map[string]interface{}
		return ret
	}
	return o.AccessItemReference
}

// GetAccessItemReferenceOk returns a tuple with the AccessItemReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutlierFeatureSummary) GetAccessItemReferenceOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.AccessItemReference) {
		return map[string]interface{}{}, false
	}
	return o.AccessItemReference, true
}

// HasAccessItemReference returns a boolean if a field has been set.
func (o *OutlierFeatureSummary) HasAccessItemReference() bool {
	if o != nil && !isNil(o.AccessItemReference) {
		return true
	}

	return false
}

// SetAccessItemReference gets a reference to the given map[string]interface{} and assigns it to the AccessItemReference field.
func (o *OutlierFeatureSummary) SetAccessItemReference(v map[string]interface{}) {
	o.AccessItemReference = v
}

func (o OutlierFeatureSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OutlierFeatureSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ContributingFeatureName) {
		toSerialize["contributingFeatureName"] = o.ContributingFeatureName
	}
	if !isNil(o.IdentityOutlierDisplayName) {
		toSerialize["identityOutlierDisplayName"] = o.IdentityOutlierDisplayName
	}
	if !isNil(o.OutlierFeatureDisplayValues) {
		toSerialize["outlierFeatureDisplayValues"] = o.OutlierFeatureDisplayValues
	}
	if !isNil(o.FeatureDefinition) {
		toSerialize["featureDefinition"] = o.FeatureDefinition
	}
	if !isNil(o.FeatureExplanation) {
		toSerialize["featureExplanation"] = o.FeatureExplanation
	}
	if !isNil(o.PeerDisplayName) {
		toSerialize["peerDisplayName"] = o.PeerDisplayName
	}
	if !isNil(o.PeerIdentityId) {
		toSerialize["peerIdentityId"] = o.PeerIdentityId
	}
	if !isNil(o.AccessItemReference) {
		toSerialize["accessItemReference"] = o.AccessItemReference
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OutlierFeatureSummary) UnmarshalJSON(bytes []byte) (err error) {
	varOutlierFeatureSummary := _OutlierFeatureSummary{}

	if err = json.Unmarshal(bytes, &varOutlierFeatureSummary); err == nil {
	*o = OutlierFeatureSummary(varOutlierFeatureSummary)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "contributingFeatureName")
		delete(additionalProperties, "identityOutlierDisplayName")
		delete(additionalProperties, "outlierFeatureDisplayValues")
		delete(additionalProperties, "featureDefinition")
		delete(additionalProperties, "featureExplanation")
		delete(additionalProperties, "peerDisplayName")
		delete(additionalProperties, "peerIdentityId")
		delete(additionalProperties, "accessItemReference")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOutlierFeatureSummary struct {
	value *OutlierFeatureSummary
	isSet bool
}

func (v NullableOutlierFeatureSummary) Get() *OutlierFeatureSummary {
	return v.value
}

func (v *NullableOutlierFeatureSummary) Set(val *OutlierFeatureSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableOutlierFeatureSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableOutlierFeatureSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutlierFeatureSummary(val *OutlierFeatureSummary) *NullableOutlierFeatureSummary {
	return &NullableOutlierFeatureSummary{value: val, isSet: true}
}

func (v NullableOutlierFeatureSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutlierFeatureSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


