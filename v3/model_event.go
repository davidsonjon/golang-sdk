/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
	"time"
)

// checks if the Event type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Event{}

// Event Event
type Event struct {
	Id string `json:"id"`
	Name string `json:"name"`
	DocumentType DocumentType `json:"_type"`
	// A date-time in ISO-8601 format
	Created NullableTime `json:"created,omitempty"`
	// A date-time in ISO-8601 format
	Synced NullableTime `json:"synced,omitempty"`
	// The action that was performed
	Action *string `json:"action,omitempty"`
	// The type of event
	Type *string `json:"type,omitempty"`
	Actor *NameType `json:"actor,omitempty"`
	Target *NameType `json:"target,omitempty"`
	Stack *string `json:"stack,omitempty"`
	TrackingNumber *string `json:"trackingNumber,omitempty"`
	IpAddress *string `json:"ipAddress,omitempty"`
	Details *string `json:"details,omitempty"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Objects []string `json:"objects,omitempty"`
	Operation *string `json:"operation,omitempty"`
	Status *string `json:"status,omitempty"`
	TechnicalName *string `json:"technicalName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Event Event

// NewEvent instantiates a new Event object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEvent(id string, name string, type_ DocumentType) *Event {
	this := Event{}
	this.Id = id
	this.Name = name
	this.DocumentType = type_
	return &this
}

// NewEventWithDefaults instantiates a new Event object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventWithDefaults() *Event {
	this := Event{}
	return &this
}

// GetId returns the Id field value
func (o *Event) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Event) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Event) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Event) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Event) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Event) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *Event) GetDocumentType() DocumentType {
	if o == nil {
		var ret DocumentType
		return ret
	}

	return o.DocumentType
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Event) GetDocumentTypeOk() (*DocumentType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DocumentType, true
}

// SetType sets field value
func (o *Event) SetDocumentType(v DocumentType) {
	o.DocumentType = v
}

// GetCreated returns the Created field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Event) GetCreated() time.Time {
	if o == nil || isNil(o.Created.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Event) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// HasCreated returns a boolean if a field has been set.
func (o *Event) HasCreated() bool {
	if o != nil && o.Created.IsSet() {
		return true
	}

	return false
}

// SetCreated gets a reference to the given NullableTime and assigns it to the Created field.
func (o *Event) SetCreated(v time.Time) {
	o.Created.Set(&v)
}
// SetCreatedNil sets the value for Created to be an explicit nil
func (o *Event) SetCreatedNil() {
	o.Created.Set(nil)
}

// UnsetCreated ensures that no value is present for Created, not even an explicit nil
func (o *Event) UnsetCreated() {
	o.Created.Unset()
}

// GetSynced returns the Synced field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Event) GetSynced() time.Time {
	if o == nil || isNil(o.Synced.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Synced.Get()
}

// GetSyncedOk returns a tuple with the Synced field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Event) GetSyncedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Synced.Get(), o.Synced.IsSet()
}

// HasSynced returns a boolean if a field has been set.
func (o *Event) HasSynced() bool {
	if o != nil && o.Synced.IsSet() {
		return true
	}

	return false
}

// SetSynced gets a reference to the given NullableTime and assigns it to the Synced field.
func (o *Event) SetSynced(v time.Time) {
	o.Synced.Set(&v)
}
// SetSyncedNil sets the value for Synced to be an explicit nil
func (o *Event) SetSyncedNil() {
	o.Synced.Set(nil)
}

// UnsetSynced ensures that no value is present for Synced, not even an explicit nil
func (o *Event) UnsetSynced() {
	o.Synced.Unset()
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *Event) GetAction() string {
	if o == nil || isNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetActionOk() (*string, bool) {
	if o == nil || isNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *Event) HasAction() bool {
	if o != nil && !isNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *Event) SetAction(v string) {
	o.Action = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Event) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Event) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Event) SetType(v string) {
	o.Type = &v
}

// GetActor returns the Actor field value if set, zero value otherwise.
func (o *Event) GetActor() NameType {
	if o == nil || isNil(o.Actor) {
		var ret NameType
		return ret
	}
	return *o.Actor
}

// GetActorOk returns a tuple with the Actor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetActorOk() (*NameType, bool) {
	if o == nil || isNil(o.Actor) {
		return nil, false
	}
	return o.Actor, true
}

// HasActor returns a boolean if a field has been set.
func (o *Event) HasActor() bool {
	if o != nil && !isNil(o.Actor) {
		return true
	}

	return false
}

// SetActor gets a reference to the given NameType and assigns it to the Actor field.
func (o *Event) SetActor(v NameType) {
	o.Actor = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *Event) GetTarget() NameType {
	if o == nil || isNil(o.Target) {
		var ret NameType
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetTargetOk() (*NameType, bool) {
	if o == nil || isNil(o.Target) {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *Event) HasTarget() bool {
	if o != nil && !isNil(o.Target) {
		return true
	}

	return false
}

// SetTarget gets a reference to the given NameType and assigns it to the Target field.
func (o *Event) SetTarget(v NameType) {
	o.Target = &v
}

// GetStack returns the Stack field value if set, zero value otherwise.
func (o *Event) GetStack() string {
	if o == nil || isNil(o.Stack) {
		var ret string
		return ret
	}
	return *o.Stack
}

// GetStackOk returns a tuple with the Stack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetStackOk() (*string, bool) {
	if o == nil || isNil(o.Stack) {
		return nil, false
	}
	return o.Stack, true
}

// HasStack returns a boolean if a field has been set.
func (o *Event) HasStack() bool {
	if o != nil && !isNil(o.Stack) {
		return true
	}

	return false
}

// SetStack gets a reference to the given string and assigns it to the Stack field.
func (o *Event) SetStack(v string) {
	o.Stack = &v
}

// GetTrackingNumber returns the TrackingNumber field value if set, zero value otherwise.
func (o *Event) GetTrackingNumber() string {
	if o == nil || isNil(o.TrackingNumber) {
		var ret string
		return ret
	}
	return *o.TrackingNumber
}

// GetTrackingNumberOk returns a tuple with the TrackingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetTrackingNumberOk() (*string, bool) {
	if o == nil || isNil(o.TrackingNumber) {
		return nil, false
	}
	return o.TrackingNumber, true
}

// HasTrackingNumber returns a boolean if a field has been set.
func (o *Event) HasTrackingNumber() bool {
	if o != nil && !isNil(o.TrackingNumber) {
		return true
	}

	return false
}

// SetTrackingNumber gets a reference to the given string and assigns it to the TrackingNumber field.
func (o *Event) SetTrackingNumber(v string) {
	o.TrackingNumber = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *Event) GetIpAddress() string {
	if o == nil || isNil(o.IpAddress) {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetIpAddressOk() (*string, bool) {
	if o == nil || isNil(o.IpAddress) {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *Event) HasIpAddress() bool {
	if o != nil && !isNil(o.IpAddress) {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *Event) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *Event) GetDetails() string {
	if o == nil || isNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetDetailsOk() (*string, bool) {
	if o == nil || isNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *Event) HasDetails() bool {
	if o != nil && !isNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *Event) SetDetails(v string) {
	o.Details = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *Event) GetAttributes() map[string]interface{} {
	if o == nil || isNil(o.Attributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Attributes) {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *Event) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *Event) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetObjects returns the Objects field value if set, zero value otherwise.
func (o *Event) GetObjects() []string {
	if o == nil || isNil(o.Objects) {
		var ret []string
		return ret
	}
	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetObjectsOk() ([]string, bool) {
	if o == nil || isNil(o.Objects) {
		return nil, false
	}
	return o.Objects, true
}

// HasObjects returns a boolean if a field has been set.
func (o *Event) HasObjects() bool {
	if o != nil && !isNil(o.Objects) {
		return true
	}

	return false
}

// SetObjects gets a reference to the given []string and assigns it to the Objects field.
func (o *Event) SetObjects(v []string) {
	o.Objects = v
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *Event) GetOperation() string {
	if o == nil || isNil(o.Operation) {
		var ret string
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetOperationOk() (*string, bool) {
	if o == nil || isNil(o.Operation) {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *Event) HasOperation() bool {
	if o != nil && !isNil(o.Operation) {
		return true
	}

	return false
}

// SetOperation gets a reference to the given string and assigns it to the Operation field.
func (o *Event) SetOperation(v string) {
	o.Operation = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Event) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Event) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Event) SetStatus(v string) {
	o.Status = &v
}

// GetTechnicalName returns the TechnicalName field value if set, zero value otherwise.
func (o *Event) GetTechnicalName() string {
	if o == nil || isNil(o.TechnicalName) {
		var ret string
		return ret
	}
	return *o.TechnicalName
}

// GetTechnicalNameOk returns a tuple with the TechnicalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetTechnicalNameOk() (*string, bool) {
	if o == nil || isNil(o.TechnicalName) {
		return nil, false
	}
	return o.TechnicalName, true
}

// HasTechnicalName returns a boolean if a field has been set.
func (o *Event) HasTechnicalName() bool {
	if o != nil && !isNil(o.TechnicalName) {
		return true
	}

	return false
}

// SetTechnicalName gets a reference to the given string and assigns it to the TechnicalName field.
func (o *Event) SetTechnicalName(v string) {
	o.TechnicalName = &v
}

func (o Event) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Event) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["_type"] = o.Type
	if o.Created.IsSet() {
		toSerialize["created"] = o.Created.Get()
	}
	if o.Synced.IsSet() {
		toSerialize["synced"] = o.Synced.Get()
	}
	if !isNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Actor) {
		toSerialize["actor"] = o.Actor
	}
	if !isNil(o.Target) {
		toSerialize["target"] = o.Target
	}
	if !isNil(o.Stack) {
		toSerialize["stack"] = o.Stack
	}
	if !isNil(o.TrackingNumber) {
		toSerialize["trackingNumber"] = o.TrackingNumber
	}
	if !isNil(o.IpAddress) {
		toSerialize["ipAddress"] = o.IpAddress
	}
	if !isNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !isNil(o.Objects) {
		toSerialize["objects"] = o.Objects
	}
	if !isNil(o.Operation) {
		toSerialize["operation"] = o.Operation
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.TechnicalName) {
		toSerialize["technicalName"] = o.TechnicalName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Event) UnmarshalJSON(bytes []byte) (err error) {
	varEvent := _Event{}

	if err = json.Unmarshal(bytes, &varEvent); err == nil {
		*o = Event(varEvent)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "_type")
		delete(additionalProperties, "created")
		delete(additionalProperties, "synced")
		delete(additionalProperties, "action")
		delete(additionalProperties, "type")
		delete(additionalProperties, "actor")
		delete(additionalProperties, "target")
		delete(additionalProperties, "stack")
		delete(additionalProperties, "trackingNumber")
		delete(additionalProperties, "ipAddress")
		delete(additionalProperties, "details")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "objects")
		delete(additionalProperties, "operation")
		delete(additionalProperties, "status")
		delete(additionalProperties, "technicalName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEvent struct {
	value *Event
	isSet bool
}

func (v NullableEvent) Get() *Event {
	return v.value
}

func (v *NullableEvent) Set(val *Event) {
	v.value = val
	v.isSet = true
}

func (v NullableEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvent(val *Event) *NullableEvent {
	return &NullableEvent{value: val, isSet: true}
}

func (v NullableEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


