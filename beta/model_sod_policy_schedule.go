/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package beta

import (
	"encoding/json"
	"time"
)

// checks if the SodPolicySchedule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SodPolicySchedule{}

// SodPolicySchedule struct for SodPolicySchedule
type SodPolicySchedule struct {
	// SOD Policy schedule name
	Name *string `json:"name,omitempty"`
	// The time when this SOD policy schedule is created.
	Created *time.Time `json:"created,omitempty"`
	// The time when this SOD policy schedule is modified.
	Modified *time.Time `json:"modified,omitempty"`
	// SOD Policy schedule description
	Description *string `json:"description,omitempty"`
	Schedule *Schedule1 `json:"schedule,omitempty"`
	Recipients []BaseReferenceDto `json:"recipients,omitempty"`
	// Indicates if empty results need to be emailed
	EmailEmptyResults *bool `json:"emailEmptyResults,omitempty"`
	// Policy's creator ID
	CreatorId *string `json:"creatorId,omitempty"`
	// Policy's modifier ID
	ModifierId *string `json:"modifierId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SodPolicySchedule SodPolicySchedule

// NewSodPolicySchedule instantiates a new SodPolicySchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSodPolicySchedule() *SodPolicySchedule {
	this := SodPolicySchedule{}
	return &this
}

// NewSodPolicyScheduleWithDefaults instantiates a new SodPolicySchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSodPolicyScheduleWithDefaults() *SodPolicySchedule {
	this := SodPolicySchedule{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SodPolicySchedule) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodPolicySchedule) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SodPolicySchedule) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SodPolicySchedule) SetName(v string) {
	o.Name = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *SodPolicySchedule) GetCreated() time.Time {
	if o == nil || isNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodPolicySchedule) GetCreatedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *SodPolicySchedule) HasCreated() bool {
	if o != nil && !isNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *SodPolicySchedule) SetCreated(v time.Time) {
	o.Created = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *SodPolicySchedule) GetModified() time.Time {
	if o == nil || isNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodPolicySchedule) GetModifiedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *SodPolicySchedule) HasModified() bool {
	if o != nil && !isNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *SodPolicySchedule) SetModified(v time.Time) {
	o.Modified = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SodPolicySchedule) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodPolicySchedule) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SodPolicySchedule) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SodPolicySchedule) SetDescription(v string) {
	o.Description = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *SodPolicySchedule) GetSchedule() Schedule1 {
	if o == nil || isNil(o.Schedule) {
		var ret Schedule1
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodPolicySchedule) GetScheduleOk() (*Schedule1, bool) {
	if o == nil || isNil(o.Schedule) {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *SodPolicySchedule) HasSchedule() bool {
	if o != nil && !isNil(o.Schedule) {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given Schedule1 and assigns it to the Schedule field.
func (o *SodPolicySchedule) SetSchedule(v Schedule1) {
	o.Schedule = &v
}

// GetRecipients returns the Recipients field value if set, zero value otherwise.
func (o *SodPolicySchedule) GetRecipients() []BaseReferenceDto {
	if o == nil || isNil(o.Recipients) {
		var ret []BaseReferenceDto
		return ret
	}
	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodPolicySchedule) GetRecipientsOk() ([]BaseReferenceDto, bool) {
	if o == nil || isNil(o.Recipients) {
		return nil, false
	}
	return o.Recipients, true
}

// HasRecipients returns a boolean if a field has been set.
func (o *SodPolicySchedule) HasRecipients() bool {
	if o != nil && !isNil(o.Recipients) {
		return true
	}

	return false
}

// SetRecipients gets a reference to the given []BaseReferenceDto and assigns it to the Recipients field.
func (o *SodPolicySchedule) SetRecipients(v []BaseReferenceDto) {
	o.Recipients = v
}

// GetEmailEmptyResults returns the EmailEmptyResults field value if set, zero value otherwise.
func (o *SodPolicySchedule) GetEmailEmptyResults() bool {
	if o == nil || isNil(o.EmailEmptyResults) {
		var ret bool
		return ret
	}
	return *o.EmailEmptyResults
}

// GetEmailEmptyResultsOk returns a tuple with the EmailEmptyResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodPolicySchedule) GetEmailEmptyResultsOk() (*bool, bool) {
	if o == nil || isNil(o.EmailEmptyResults) {
		return nil, false
	}
	return o.EmailEmptyResults, true
}

// HasEmailEmptyResults returns a boolean if a field has been set.
func (o *SodPolicySchedule) HasEmailEmptyResults() bool {
	if o != nil && !isNil(o.EmailEmptyResults) {
		return true
	}

	return false
}

// SetEmailEmptyResults gets a reference to the given bool and assigns it to the EmailEmptyResults field.
func (o *SodPolicySchedule) SetEmailEmptyResults(v bool) {
	o.EmailEmptyResults = &v
}

// GetCreatorId returns the CreatorId field value if set, zero value otherwise.
func (o *SodPolicySchedule) GetCreatorId() string {
	if o == nil || isNil(o.CreatorId) {
		var ret string
		return ret
	}
	return *o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodPolicySchedule) GetCreatorIdOk() (*string, bool) {
	if o == nil || isNil(o.CreatorId) {
		return nil, false
	}
	return o.CreatorId, true
}

// HasCreatorId returns a boolean if a field has been set.
func (o *SodPolicySchedule) HasCreatorId() bool {
	if o != nil && !isNil(o.CreatorId) {
		return true
	}

	return false
}

// SetCreatorId gets a reference to the given string and assigns it to the CreatorId field.
func (o *SodPolicySchedule) SetCreatorId(v string) {
	o.CreatorId = &v
}

// GetModifierId returns the ModifierId field value if set, zero value otherwise.
func (o *SodPolicySchedule) GetModifierId() string {
	if o == nil || isNil(o.ModifierId) {
		var ret string
		return ret
	}
	return *o.ModifierId
}

// GetModifierIdOk returns a tuple with the ModifierId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodPolicySchedule) GetModifierIdOk() (*string, bool) {
	if o == nil || isNil(o.ModifierId) {
		return nil, false
	}
	return o.ModifierId, true
}

// HasModifierId returns a boolean if a field has been set.
func (o *SodPolicySchedule) HasModifierId() bool {
	if o != nil && !isNil(o.ModifierId) {
		return true
	}

	return false
}

// SetModifierId gets a reference to the given string and assigns it to the ModifierId field.
func (o *SodPolicySchedule) SetModifierId(v string) {
	o.ModifierId = &v
}

func (o SodPolicySchedule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SodPolicySchedule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !isNil(o.Modified) {
		toSerialize["modified"] = o.Modified
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Schedule) {
		toSerialize["schedule"] = o.Schedule
	}
	if !isNil(o.Recipients) {
		toSerialize["recipients"] = o.Recipients
	}
	if !isNil(o.EmailEmptyResults) {
		toSerialize["emailEmptyResults"] = o.EmailEmptyResults
	}
	if !isNil(o.CreatorId) {
		toSerialize["creatorId"] = o.CreatorId
	}
	if !isNil(o.ModifierId) {
		toSerialize["modifierId"] = o.ModifierId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SodPolicySchedule) UnmarshalJSON(bytes []byte) (err error) {
	varSodPolicySchedule := _SodPolicySchedule{}

	if err = json.Unmarshal(bytes, &varSodPolicySchedule); err == nil {
		*o = SodPolicySchedule(varSodPolicySchedule)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "description")
		delete(additionalProperties, "schedule")
		delete(additionalProperties, "recipients")
		delete(additionalProperties, "emailEmptyResults")
		delete(additionalProperties, "creatorId")
		delete(additionalProperties, "modifierId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSodPolicySchedule struct {
	value *SodPolicySchedule
	isSet bool
}

func (v NullableSodPolicySchedule) Get() *SodPolicySchedule {
	return v.value
}

func (v *NullableSodPolicySchedule) Set(val *SodPolicySchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableSodPolicySchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableSodPolicySchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSodPolicySchedule(val *SodPolicySchedule) *NullableSodPolicySchedule {
	return &NullableSodPolicySchedule{value: val, isSet: true}
}

func (v NullableSodPolicySchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSodPolicySchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


