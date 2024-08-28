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

// checks if the IdentityProfileExportedObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityProfileExportedObject{}

// IdentityProfileExportedObject Identity profile exported object.
type IdentityProfileExportedObject struct {
	// Version or object from the target service.
	Version *int32 `json:"version,omitempty"`
	Self *IdentityProfileExportedObjectSelf `json:"self,omitempty"`
	Object *IdentityProfile `json:"object,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentityProfileExportedObject IdentityProfileExportedObject

// NewIdentityProfileExportedObject instantiates a new IdentityProfileExportedObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityProfileExportedObject() *IdentityProfileExportedObject {
	this := IdentityProfileExportedObject{}
	return &this
}

// NewIdentityProfileExportedObjectWithDefaults instantiates a new IdentityProfileExportedObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityProfileExportedObjectWithDefaults() *IdentityProfileExportedObject {
	this := IdentityProfileExportedObject{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *IdentityProfileExportedObject) GetVersion() int32 {
	if o == nil || IsNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProfileExportedObject) GetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *IdentityProfileExportedObject) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *IdentityProfileExportedObject) SetVersion(v int32) {
	o.Version = &v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *IdentityProfileExportedObject) GetSelf() IdentityProfileExportedObjectSelf {
	if o == nil || IsNil(o.Self) {
		var ret IdentityProfileExportedObjectSelf
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProfileExportedObject) GetSelfOk() (*IdentityProfileExportedObjectSelf, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *IdentityProfileExportedObject) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given IdentityProfileExportedObjectSelf and assigns it to the Self field.
func (o *IdentityProfileExportedObject) SetSelf(v IdentityProfileExportedObjectSelf) {
	o.Self = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *IdentityProfileExportedObject) GetObject() IdentityProfile {
	if o == nil || IsNil(o.Object) {
		var ret IdentityProfile
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProfileExportedObject) GetObjectOk() (*IdentityProfile, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *IdentityProfileExportedObject) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given IdentityProfile and assigns it to the Object field.
func (o *IdentityProfileExportedObject) SetObject(v IdentityProfile) {
	o.Object = &v
}

func (o IdentityProfileExportedObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityProfileExportedObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IdentityProfileExportedObject) UnmarshalJSON(data []byte) (err error) {
	varIdentityProfileExportedObject := _IdentityProfileExportedObject{}

	err = json.Unmarshal(data, &varIdentityProfileExportedObject)

	if err != nil {
		return err
	}

	*o = IdentityProfileExportedObject(varIdentityProfileExportedObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "version")
		delete(additionalProperties, "self")
		delete(additionalProperties, "object")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentityProfileExportedObject struct {
	value *IdentityProfileExportedObject
	isSet bool
}

func (v NullableIdentityProfileExportedObject) Get() *IdentityProfileExportedObject {
	return v.value
}

func (v *NullableIdentityProfileExportedObject) Set(val *IdentityProfileExportedObject) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityProfileExportedObject) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityProfileExportedObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityProfileExportedObject(val *IdentityProfileExportedObject) *NullableIdentityProfileExportedObject {
	return &NullableIdentityProfileExportedObject{value: val, isSet: true}
}

func (v NullableIdentityProfileExportedObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityProfileExportedObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


