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

// checks if the WorkgroupDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkgroupDto{}

// WorkgroupDto struct for WorkgroupDto
type WorkgroupDto struct {
	Owner *Owner `json:"owner,omitempty"`
	// ID of the object to which this reference applies
	Id *string `json:"id,omitempty"`
	// Description of the Governance Group
	Description *string `json:"description,omitempty"`
	// Number of members in the Governance Group.
	MemberCount *int64 `json:"memberCount,omitempty"`
	// Number of connections in the Governance Group.
	ConnectionCount *int64 `json:"connectionCount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkgroupDto WorkgroupDto

// NewWorkgroupDto instantiates a new WorkgroupDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkgroupDto() *WorkgroupDto {
	this := WorkgroupDto{}
	return &this
}

// NewWorkgroupDtoWithDefaults instantiates a new WorkgroupDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkgroupDtoWithDefaults() *WorkgroupDto {
	this := WorkgroupDto{}
	return &this
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *WorkgroupDto) GetOwner() Owner {
	if o == nil || isNil(o.Owner) {
		var ret Owner
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkgroupDto) GetOwnerOk() (*Owner, bool) {
	if o == nil || isNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *WorkgroupDto) HasOwner() bool {
	if o != nil && !isNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given Owner and assigns it to the Owner field.
func (o *WorkgroupDto) SetOwner(v Owner) {
	o.Owner = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WorkgroupDto) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkgroupDto) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WorkgroupDto) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WorkgroupDto) SetId(v string) {
	o.Id = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkgroupDto) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkgroupDto) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkgroupDto) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkgroupDto) SetDescription(v string) {
	o.Description = &v
}

// GetMemberCount returns the MemberCount field value if set, zero value otherwise.
func (o *WorkgroupDto) GetMemberCount() int64 {
	if o == nil || isNil(o.MemberCount) {
		var ret int64
		return ret
	}
	return *o.MemberCount
}

// GetMemberCountOk returns a tuple with the MemberCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkgroupDto) GetMemberCountOk() (*int64, bool) {
	if o == nil || isNil(o.MemberCount) {
		return nil, false
	}
	return o.MemberCount, true
}

// HasMemberCount returns a boolean if a field has been set.
func (o *WorkgroupDto) HasMemberCount() bool {
	if o != nil && !isNil(o.MemberCount) {
		return true
	}

	return false
}

// SetMemberCount gets a reference to the given int64 and assigns it to the MemberCount field.
func (o *WorkgroupDto) SetMemberCount(v int64) {
	o.MemberCount = &v
}

// GetConnectionCount returns the ConnectionCount field value if set, zero value otherwise.
func (o *WorkgroupDto) GetConnectionCount() int64 {
	if o == nil || isNil(o.ConnectionCount) {
		var ret int64
		return ret
	}
	return *o.ConnectionCount
}

// GetConnectionCountOk returns a tuple with the ConnectionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkgroupDto) GetConnectionCountOk() (*int64, bool) {
	if o == nil || isNil(o.ConnectionCount) {
		return nil, false
	}
	return o.ConnectionCount, true
}

// HasConnectionCount returns a boolean if a field has been set.
func (o *WorkgroupDto) HasConnectionCount() bool {
	if o != nil && !isNil(o.ConnectionCount) {
		return true
	}

	return false
}

// SetConnectionCount gets a reference to the given int64 and assigns it to the ConnectionCount field.
func (o *WorkgroupDto) SetConnectionCount(v int64) {
	o.ConnectionCount = &v
}

func (o WorkgroupDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkgroupDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.MemberCount) {
		toSerialize["memberCount"] = o.MemberCount
	}
	if !isNil(o.ConnectionCount) {
		toSerialize["connectionCount"] = o.ConnectionCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WorkgroupDto) UnmarshalJSON(bytes []byte) (err error) {
	varWorkgroupDto := _WorkgroupDto{}

	if err = json.Unmarshal(bytes, &varWorkgroupDto); err == nil {
		*o = WorkgroupDto(varWorkgroupDto)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "owner")
		delete(additionalProperties, "id")
		delete(additionalProperties, "description")
		delete(additionalProperties, "memberCount")
		delete(additionalProperties, "connectionCount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkgroupDto struct {
	value *WorkgroupDto
	isSet bool
}

func (v NullableWorkgroupDto) Get() *WorkgroupDto {
	return v.value
}

func (v *NullableWorkgroupDto) Set(val *WorkgroupDto) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkgroupDto) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkgroupDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkgroupDto(val *WorkgroupDto) *NullableWorkgroupDto {
	return &NullableWorkgroupDto{value: val, isSet: true}
}

func (v NullableWorkgroupDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkgroupDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


