/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the SimIntegrationDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SimIntegrationDetails{}

// SimIntegrationDetails struct for SimIntegrationDetails
type SimIntegrationDetails struct {
	// System-generated unique ID of the Object
	Id *string `json:"id,omitempty"`
	// Name of the Object
	Name string `json:"name"`
	// Creation date of the Object
	Created *time.Time `json:"created,omitempty"`
	// Last modification date of the Object
	Modified *time.Time `json:"modified,omitempty"`
	// The description of the integration
	Description *string `json:"description,omitempty"`
	// The integration type
	Type *string `json:"type,omitempty"`
	// The attributes map containing the credentials used to configure the integration.
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	// The list of sources (managed resources)
	Sources []string `json:"sources,omitempty"`
	// The cluster/proxy
	Cluster *string `json:"cluster,omitempty"`
	// Custom mapping between the integration result and the provisioning result
	StatusMap map[string]interface{} `json:"statusMap,omitempty"`
	// Request data to customize desc and body of the created ticket
	Request map[string]interface{} `json:"request,omitempty"`
	BeforeProvisioningRule *SimIntegrationDetailsAllOfBeforeProvisioningRule `json:"beforeProvisioningRule,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SimIntegrationDetails SimIntegrationDetails

// NewSimIntegrationDetails instantiates a new SimIntegrationDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimIntegrationDetails(name string) *SimIntegrationDetails {
	this := SimIntegrationDetails{}
	this.Name = name
	return &this
}

// NewSimIntegrationDetailsWithDefaults instantiates a new SimIntegrationDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimIntegrationDetailsWithDefaults() *SimIntegrationDetails {
	this := SimIntegrationDetails{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SimIntegrationDetails) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimIntegrationDetails) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SimIntegrationDetails) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SimIntegrationDetails) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *SimIntegrationDetails) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SimIntegrationDetails) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SimIntegrationDetails) SetName(v string) {
	o.Name = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *SimIntegrationDetails) GetCreated() time.Time {
	if o == nil || isNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimIntegrationDetails) GetCreatedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *SimIntegrationDetails) HasCreated() bool {
	if o != nil && !isNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *SimIntegrationDetails) SetCreated(v time.Time) {
	o.Created = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *SimIntegrationDetails) GetModified() time.Time {
	if o == nil || isNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimIntegrationDetails) GetModifiedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *SimIntegrationDetails) HasModified() bool {
	if o != nil && !isNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *SimIntegrationDetails) SetModified(v time.Time) {
	o.Modified = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SimIntegrationDetails) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimIntegrationDetails) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SimIntegrationDetails) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SimIntegrationDetails) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SimIntegrationDetails) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimIntegrationDetails) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SimIntegrationDetails) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SimIntegrationDetails) SetType(v string) {
	o.Type = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SimIntegrationDetails) GetAttributes() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SimIntegrationDetails) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Attributes) {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *SimIntegrationDetails) HasAttributes() bool {
	if o != nil && isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *SimIntegrationDetails) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetSources returns the Sources field value if set, zero value otherwise.
func (o *SimIntegrationDetails) GetSources() []string {
	if o == nil || isNil(o.Sources) {
		var ret []string
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimIntegrationDetails) GetSourcesOk() ([]string, bool) {
	if o == nil || isNil(o.Sources) {
		return nil, false
	}
	return o.Sources, true
}

// HasSources returns a boolean if a field has been set.
func (o *SimIntegrationDetails) HasSources() bool {
	if o != nil && !isNil(o.Sources) {
		return true
	}

	return false
}

// SetSources gets a reference to the given []string and assigns it to the Sources field.
func (o *SimIntegrationDetails) SetSources(v []string) {
	o.Sources = v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *SimIntegrationDetails) GetCluster() string {
	if o == nil || isNil(o.Cluster) {
		var ret string
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimIntegrationDetails) GetClusterOk() (*string, bool) {
	if o == nil || isNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *SimIntegrationDetails) HasCluster() bool {
	if o != nil && !isNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given string and assigns it to the Cluster field.
func (o *SimIntegrationDetails) SetCluster(v string) {
	o.Cluster = &v
}

// GetStatusMap returns the StatusMap field value if set, zero value otherwise.
func (o *SimIntegrationDetails) GetStatusMap() map[string]interface{} {
	if o == nil || isNil(o.StatusMap) {
		var ret map[string]interface{}
		return ret
	}
	return o.StatusMap
}

// GetStatusMapOk returns a tuple with the StatusMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimIntegrationDetails) GetStatusMapOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.StatusMap) {
		return map[string]interface{}{}, false
	}
	return o.StatusMap, true
}

// HasStatusMap returns a boolean if a field has been set.
func (o *SimIntegrationDetails) HasStatusMap() bool {
	if o != nil && !isNil(o.StatusMap) {
		return true
	}

	return false
}

// SetStatusMap gets a reference to the given map[string]interface{} and assigns it to the StatusMap field.
func (o *SimIntegrationDetails) SetStatusMap(v map[string]interface{}) {
	o.StatusMap = v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *SimIntegrationDetails) GetRequest() map[string]interface{} {
	if o == nil || isNil(o.Request) {
		var ret map[string]interface{}
		return ret
	}
	return o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimIntegrationDetails) GetRequestOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Request) {
		return map[string]interface{}{}, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *SimIntegrationDetails) HasRequest() bool {
	if o != nil && !isNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given map[string]interface{} and assigns it to the Request field.
func (o *SimIntegrationDetails) SetRequest(v map[string]interface{}) {
	o.Request = v
}

// GetBeforeProvisioningRule returns the BeforeProvisioningRule field value if set, zero value otherwise.
func (o *SimIntegrationDetails) GetBeforeProvisioningRule() SimIntegrationDetailsAllOfBeforeProvisioningRule {
	if o == nil || isNil(o.BeforeProvisioningRule) {
		var ret SimIntegrationDetailsAllOfBeforeProvisioningRule
		return ret
	}
	return *o.BeforeProvisioningRule
}

// GetBeforeProvisioningRuleOk returns a tuple with the BeforeProvisioningRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimIntegrationDetails) GetBeforeProvisioningRuleOk() (*SimIntegrationDetailsAllOfBeforeProvisioningRule, bool) {
	if o == nil || isNil(o.BeforeProvisioningRule) {
		return nil, false
	}
	return o.BeforeProvisioningRule, true
}

// HasBeforeProvisioningRule returns a boolean if a field has been set.
func (o *SimIntegrationDetails) HasBeforeProvisioningRule() bool {
	if o != nil && !isNil(o.BeforeProvisioningRule) {
		return true
	}

	return false
}

// SetBeforeProvisioningRule gets a reference to the given SimIntegrationDetailsAllOfBeforeProvisioningRule and assigns it to the BeforeProvisioningRule field.
func (o *SimIntegrationDetails) SetBeforeProvisioningRule(v SimIntegrationDetailsAllOfBeforeProvisioningRule) {
	o.BeforeProvisioningRule = &v
}

func (o SimIntegrationDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SimIntegrationDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	toSerialize["name"] = o.Name
	// skip: created is readOnly
	// skip: modified is readOnly
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if !isNil(o.Sources) {
		toSerialize["sources"] = o.Sources
	}
	if !isNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	if !isNil(o.StatusMap) {
		toSerialize["statusMap"] = o.StatusMap
	}
	if !isNil(o.Request) {
		toSerialize["request"] = o.Request
	}
	if !isNil(o.BeforeProvisioningRule) {
		toSerialize["beforeProvisioningRule"] = o.BeforeProvisioningRule
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SimIntegrationDetails) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varSimIntegrationDetails := _SimIntegrationDetails{}

	if err = json.Unmarshal(bytes, &varSimIntegrationDetails); err == nil {
	*o = SimIntegrationDetails(varSimIntegrationDetails)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "description")
		delete(additionalProperties, "type")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "sources")
		delete(additionalProperties, "cluster")
		delete(additionalProperties, "statusMap")
		delete(additionalProperties, "request")
		delete(additionalProperties, "beforeProvisioningRule")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSimIntegrationDetails struct {
	value *SimIntegrationDetails
	isSet bool
}

func (v NullableSimIntegrationDetails) Get() *SimIntegrationDetails {
	return v.value
}

func (v *NullableSimIntegrationDetails) Set(val *SimIntegrationDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableSimIntegrationDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableSimIntegrationDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimIntegrationDetails(val *SimIntegrationDetails) *NullableSimIntegrationDetails {
	return &NullableSimIntegrationDetails{value: val, isSet: true}
}

func (v NullableSimIntegrationDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimIntegrationDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


