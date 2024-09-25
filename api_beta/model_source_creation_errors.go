/*
Identity Security Cloud Beta API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
	"time"
)

// checks if the SourceCreationErrors type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SourceCreationErrors{}

// SourceCreationErrors struct for SourceCreationErrors
type SourceCreationErrors struct {
	// Multi-Host Integration ID.
	MultihostId *string `json:"multihost_id,omitempty"`
	// Source's human-readable name.
	SourceName *string `json:"source_name,omitempty"`
	// Source's human-readable description.
	SourceError *string `json:"source_error,omitempty"`
	// Date-time when the source was created
	Created *time.Time `json:"created,omitempty"`
	// Date-time when the source was last modified.
	Modified *time.Time `json:"modified,omitempty"`
	// operation category (e.g. DELETE).
	Operation NullableString `json:"operation,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SourceCreationErrors SourceCreationErrors

// NewSourceCreationErrors instantiates a new SourceCreationErrors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceCreationErrors() *SourceCreationErrors {
	this := SourceCreationErrors{}
	return &this
}

// NewSourceCreationErrorsWithDefaults instantiates a new SourceCreationErrors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceCreationErrorsWithDefaults() *SourceCreationErrors {
	this := SourceCreationErrors{}
	return &this
}

// GetMultihostId returns the MultihostId field value if set, zero value otherwise.
func (o *SourceCreationErrors) GetMultihostId() string {
	if o == nil || IsNil(o.MultihostId) {
		var ret string
		return ret
	}
	return *o.MultihostId
}

// GetMultihostIdOk returns a tuple with the MultihostId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceCreationErrors) GetMultihostIdOk() (*string, bool) {
	if o == nil || IsNil(o.MultihostId) {
		return nil, false
	}
	return o.MultihostId, true
}

// HasMultihostId returns a boolean if a field has been set.
func (o *SourceCreationErrors) HasMultihostId() bool {
	if o != nil && !IsNil(o.MultihostId) {
		return true
	}

	return false
}

// SetMultihostId gets a reference to the given string and assigns it to the MultihostId field.
func (o *SourceCreationErrors) SetMultihostId(v string) {
	o.MultihostId = &v
}

// GetSourceName returns the SourceName field value if set, zero value otherwise.
func (o *SourceCreationErrors) GetSourceName() string {
	if o == nil || IsNil(o.SourceName) {
		var ret string
		return ret
	}
	return *o.SourceName
}

// GetSourceNameOk returns a tuple with the SourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceCreationErrors) GetSourceNameOk() (*string, bool) {
	if o == nil || IsNil(o.SourceName) {
		return nil, false
	}
	return o.SourceName, true
}

// HasSourceName returns a boolean if a field has been set.
func (o *SourceCreationErrors) HasSourceName() bool {
	if o != nil && !IsNil(o.SourceName) {
		return true
	}

	return false
}

// SetSourceName gets a reference to the given string and assigns it to the SourceName field.
func (o *SourceCreationErrors) SetSourceName(v string) {
	o.SourceName = &v
}

// GetSourceError returns the SourceError field value if set, zero value otherwise.
func (o *SourceCreationErrors) GetSourceError() string {
	if o == nil || IsNil(o.SourceError) {
		var ret string
		return ret
	}
	return *o.SourceError
}

// GetSourceErrorOk returns a tuple with the SourceError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceCreationErrors) GetSourceErrorOk() (*string, bool) {
	if o == nil || IsNil(o.SourceError) {
		return nil, false
	}
	return o.SourceError, true
}

// HasSourceError returns a boolean if a field has been set.
func (o *SourceCreationErrors) HasSourceError() bool {
	if o != nil && !IsNil(o.SourceError) {
		return true
	}

	return false
}

// SetSourceError gets a reference to the given string and assigns it to the SourceError field.
func (o *SourceCreationErrors) SetSourceError(v string) {
	o.SourceError = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *SourceCreationErrors) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceCreationErrors) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *SourceCreationErrors) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *SourceCreationErrors) SetCreated(v time.Time) {
	o.Created = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *SourceCreationErrors) GetModified() time.Time {
	if o == nil || IsNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceCreationErrors) GetModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *SourceCreationErrors) HasModified() bool {
	if o != nil && !IsNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *SourceCreationErrors) SetModified(v time.Time) {
	o.Modified = &v
}

// GetOperation returns the Operation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourceCreationErrors) GetOperation() string {
	if o == nil || IsNil(o.Operation.Get()) {
		var ret string
		return ret
	}
	return *o.Operation.Get()
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SourceCreationErrors) GetOperationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Operation.Get(), o.Operation.IsSet()
}

// HasOperation returns a boolean if a field has been set.
func (o *SourceCreationErrors) HasOperation() bool {
	if o != nil && o.Operation.IsSet() {
		return true
	}

	return false
}

// SetOperation gets a reference to the given NullableString and assigns it to the Operation field.
func (o *SourceCreationErrors) SetOperation(v string) {
	o.Operation.Set(&v)
}
// SetOperationNil sets the value for Operation to be an explicit nil
func (o *SourceCreationErrors) SetOperationNil() {
	o.Operation.Set(nil)
}

// UnsetOperation ensures that no value is present for Operation, not even an explicit nil
func (o *SourceCreationErrors) UnsetOperation() {
	o.Operation.Unset()
}

func (o SourceCreationErrors) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SourceCreationErrors) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MultihostId) {
		toSerialize["multihost_id"] = o.MultihostId
	}
	if !IsNil(o.SourceName) {
		toSerialize["source_name"] = o.SourceName
	}
	if !IsNil(o.SourceError) {
		toSerialize["source_error"] = o.SourceError
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Modified) {
		toSerialize["modified"] = o.Modified
	}
	if o.Operation.IsSet() {
		toSerialize["operation"] = o.Operation.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SourceCreationErrors) UnmarshalJSON(data []byte) (err error) {
	varSourceCreationErrors := _SourceCreationErrors{}

	err = json.Unmarshal(data, &varSourceCreationErrors)

	if err != nil {
		return err
	}

	*o = SourceCreationErrors(varSourceCreationErrors)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "multihost_id")
		delete(additionalProperties, "source_name")
		delete(additionalProperties, "source_error")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "operation")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSourceCreationErrors struct {
	value *SourceCreationErrors
	isSet bool
}

func (v NullableSourceCreationErrors) Get() *SourceCreationErrors {
	return v.value
}

func (v *NullableSourceCreationErrors) Set(val *SourceCreationErrors) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceCreationErrors) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceCreationErrors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceCreationErrors(val *SourceCreationErrors) *NullableSourceCreationErrors {
	return &NullableSourceCreationErrors{value: val, isSet: true}
}

func (v NullableSourceCreationErrors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceCreationErrors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


