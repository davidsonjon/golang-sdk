/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
)

// checks if the SourceHealthDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SourceHealthDto{}

// SourceHealthDto Dto for source health data
type SourceHealthDto struct {
	// the id of the Source
	Id *string `json:"id,omitempty"`
	// Specifies the type of system being managed e.g. Active Directory, Workday, etc.. If you are creating a Delimited File source, you must set the `provisionasCsv` query parameter to `true`. 
	Type *string `json:"type,omitempty"`
	// the name of the source
	Name *string `json:"name,omitempty"`
	// source's org
	Org *string `json:"org,omitempty"`
	// Is the source authoritative
	IsAuthoritative *bool `json:"isAuthoritative,omitempty"`
	// Is the source in a cluster
	IsCluster *bool `json:"isCluster,omitempty"`
	// source's hostname
	Hostname *string `json:"hostname,omitempty"`
	// source's pod
	Pod *string `json:"pod,omitempty"`
	// The version of the iqService
	IqServiceVersion *string `json:"iqServiceVersion,omitempty"`
	// connection test result
	Status *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SourceHealthDto SourceHealthDto

// NewSourceHealthDto instantiates a new SourceHealthDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceHealthDto() *SourceHealthDto {
	this := SourceHealthDto{}
	return &this
}

// NewSourceHealthDtoWithDefaults instantiates a new SourceHealthDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceHealthDtoWithDefaults() *SourceHealthDto {
	this := SourceHealthDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SourceHealthDto) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceHealthDto) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SourceHealthDto) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SourceHealthDto) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SourceHealthDto) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceHealthDto) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SourceHealthDto) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SourceHealthDto) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SourceHealthDto) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceHealthDto) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SourceHealthDto) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SourceHealthDto) SetName(v string) {
	o.Name = &v
}

// GetOrg returns the Org field value if set, zero value otherwise.
func (o *SourceHealthDto) GetOrg() string {
	if o == nil || isNil(o.Org) {
		var ret string
		return ret
	}
	return *o.Org
}

// GetOrgOk returns a tuple with the Org field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceHealthDto) GetOrgOk() (*string, bool) {
	if o == nil || isNil(o.Org) {
		return nil, false
	}
	return o.Org, true
}

// HasOrg returns a boolean if a field has been set.
func (o *SourceHealthDto) HasOrg() bool {
	if o != nil && !isNil(o.Org) {
		return true
	}

	return false
}

// SetOrg gets a reference to the given string and assigns it to the Org field.
func (o *SourceHealthDto) SetOrg(v string) {
	o.Org = &v
}

// GetIsAuthoritative returns the IsAuthoritative field value if set, zero value otherwise.
func (o *SourceHealthDto) GetIsAuthoritative() bool {
	if o == nil || isNil(o.IsAuthoritative) {
		var ret bool
		return ret
	}
	return *o.IsAuthoritative
}

// GetIsAuthoritativeOk returns a tuple with the IsAuthoritative field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceHealthDto) GetIsAuthoritativeOk() (*bool, bool) {
	if o == nil || isNil(o.IsAuthoritative) {
		return nil, false
	}
	return o.IsAuthoritative, true
}

// HasIsAuthoritative returns a boolean if a field has been set.
func (o *SourceHealthDto) HasIsAuthoritative() bool {
	if o != nil && !isNil(o.IsAuthoritative) {
		return true
	}

	return false
}

// SetIsAuthoritative gets a reference to the given bool and assigns it to the IsAuthoritative field.
func (o *SourceHealthDto) SetIsAuthoritative(v bool) {
	o.IsAuthoritative = &v
}

// GetIsCluster returns the IsCluster field value if set, zero value otherwise.
func (o *SourceHealthDto) GetIsCluster() bool {
	if o == nil || isNil(o.IsCluster) {
		var ret bool
		return ret
	}
	return *o.IsCluster
}

// GetIsClusterOk returns a tuple with the IsCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceHealthDto) GetIsClusterOk() (*bool, bool) {
	if o == nil || isNil(o.IsCluster) {
		return nil, false
	}
	return o.IsCluster, true
}

// HasIsCluster returns a boolean if a field has been set.
func (o *SourceHealthDto) HasIsCluster() bool {
	if o != nil && !isNil(o.IsCluster) {
		return true
	}

	return false
}

// SetIsCluster gets a reference to the given bool and assigns it to the IsCluster field.
func (o *SourceHealthDto) SetIsCluster(v bool) {
	o.IsCluster = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *SourceHealthDto) GetHostname() string {
	if o == nil || isNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceHealthDto) GetHostnameOk() (*string, bool) {
	if o == nil || isNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *SourceHealthDto) HasHostname() bool {
	if o != nil && !isNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *SourceHealthDto) SetHostname(v string) {
	o.Hostname = &v
}

// GetPod returns the Pod field value if set, zero value otherwise.
func (o *SourceHealthDto) GetPod() string {
	if o == nil || isNil(o.Pod) {
		var ret string
		return ret
	}
	return *o.Pod
}

// GetPodOk returns a tuple with the Pod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceHealthDto) GetPodOk() (*string, bool) {
	if o == nil || isNil(o.Pod) {
		return nil, false
	}
	return o.Pod, true
}

// HasPod returns a boolean if a field has been set.
func (o *SourceHealthDto) HasPod() bool {
	if o != nil && !isNil(o.Pod) {
		return true
	}

	return false
}

// SetPod gets a reference to the given string and assigns it to the Pod field.
func (o *SourceHealthDto) SetPod(v string) {
	o.Pod = &v
}

// GetIqServiceVersion returns the IqServiceVersion field value if set, zero value otherwise.
func (o *SourceHealthDto) GetIqServiceVersion() string {
	if o == nil || isNil(o.IqServiceVersion) {
		var ret string
		return ret
	}
	return *o.IqServiceVersion
}

// GetIqServiceVersionOk returns a tuple with the IqServiceVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceHealthDto) GetIqServiceVersionOk() (*string, bool) {
	if o == nil || isNil(o.IqServiceVersion) {
		return nil, false
	}
	return o.IqServiceVersion, true
}

// HasIqServiceVersion returns a boolean if a field has been set.
func (o *SourceHealthDto) HasIqServiceVersion() bool {
	if o != nil && !isNil(o.IqServiceVersion) {
		return true
	}

	return false
}

// SetIqServiceVersion gets a reference to the given string and assigns it to the IqServiceVersion field.
func (o *SourceHealthDto) SetIqServiceVersion(v string) {
	o.IqServiceVersion = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SourceHealthDto) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceHealthDto) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SourceHealthDto) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SourceHealthDto) SetStatus(v string) {
	o.Status = &v
}

func (o SourceHealthDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SourceHealthDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Org) {
		toSerialize["org"] = o.Org
	}
	if !isNil(o.IsAuthoritative) {
		toSerialize["isAuthoritative"] = o.IsAuthoritative
	}
	if !isNil(o.IsCluster) {
		toSerialize["isCluster"] = o.IsCluster
	}
	if !isNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !isNil(o.Pod) {
		toSerialize["pod"] = o.Pod
	}
	if !isNil(o.IqServiceVersion) {
		toSerialize["iqServiceVersion"] = o.IqServiceVersion
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SourceHealthDto) UnmarshalJSON(bytes []byte) (err error) {
	varSourceHealthDto := _SourceHealthDto{}

	if err = json.Unmarshal(bytes, &varSourceHealthDto); err == nil {
	*o = SourceHealthDto(varSourceHealthDto)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "org")
		delete(additionalProperties, "isAuthoritative")
		delete(additionalProperties, "isCluster")
		delete(additionalProperties, "hostname")
		delete(additionalProperties, "pod")
		delete(additionalProperties, "iqServiceVersion")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSourceHealthDto struct {
	value *SourceHealthDto
	isSet bool
}

func (v NullableSourceHealthDto) Get() *SourceHealthDto {
	return v.value
}

func (v *NullableSourceHealthDto) Set(val *SourceHealthDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceHealthDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceHealthDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceHealthDto(val *SourceHealthDto) *NullableSourceHealthDto {
	return &NullableSourceHealthDto{value: val, isSet: true}
}

func (v NullableSourceHealthDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceHealthDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


