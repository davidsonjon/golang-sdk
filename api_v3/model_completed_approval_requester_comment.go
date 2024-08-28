/*
Identity Security Cloud V3 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v3

import (
	"encoding/json"
	"time"
)

// checks if the CompletedApprovalRequesterComment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompletedApprovalRequesterComment{}

// CompletedApprovalRequesterComment struct for CompletedApprovalRequesterComment
type CompletedApprovalRequesterComment struct {
	// Comment content.
	Comment NullableString `json:"comment,omitempty"`
	// Date and time comment was created.
	Created *time.Time `json:"created,omitempty"`
	Author *CommentDtoAuthor `json:"author,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CompletedApprovalRequesterComment CompletedApprovalRequesterComment

// NewCompletedApprovalRequesterComment instantiates a new CompletedApprovalRequesterComment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompletedApprovalRequesterComment() *CompletedApprovalRequesterComment {
	this := CompletedApprovalRequesterComment{}
	return &this
}

// NewCompletedApprovalRequesterCommentWithDefaults instantiates a new CompletedApprovalRequesterComment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompletedApprovalRequesterCommentWithDefaults() *CompletedApprovalRequesterComment {
	this := CompletedApprovalRequesterComment{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompletedApprovalRequesterComment) GetComment() string {
	if o == nil || IsNil(o.Comment.Get()) {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompletedApprovalRequesterComment) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *CompletedApprovalRequesterComment) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *CompletedApprovalRequesterComment) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *CompletedApprovalRequesterComment) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *CompletedApprovalRequesterComment) UnsetComment() {
	o.Comment.Unset()
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *CompletedApprovalRequesterComment) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompletedApprovalRequesterComment) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *CompletedApprovalRequesterComment) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *CompletedApprovalRequesterComment) SetCreated(v time.Time) {
	o.Created = &v
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *CompletedApprovalRequesterComment) GetAuthor() CommentDtoAuthor {
	if o == nil || IsNil(o.Author) {
		var ret CommentDtoAuthor
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompletedApprovalRequesterComment) GetAuthorOk() (*CommentDtoAuthor, bool) {
	if o == nil || IsNil(o.Author) {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *CompletedApprovalRequesterComment) HasAuthor() bool {
	if o != nil && !IsNil(o.Author) {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given CommentDtoAuthor and assigns it to the Author field.
func (o *CompletedApprovalRequesterComment) SetAuthor(v CommentDtoAuthor) {
	o.Author = &v
}

func (o CompletedApprovalRequesterComment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompletedApprovalRequesterComment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Comment.IsSet() {
		toSerialize["comment"] = o.Comment.Get()
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Author) {
		toSerialize["author"] = o.Author
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CompletedApprovalRequesterComment) UnmarshalJSON(data []byte) (err error) {
	varCompletedApprovalRequesterComment := _CompletedApprovalRequesterComment{}

	err = json.Unmarshal(data, &varCompletedApprovalRequesterComment)

	if err != nil {
		return err
	}

	*o = CompletedApprovalRequesterComment(varCompletedApprovalRequesterComment)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "comment")
		delete(additionalProperties, "created")
		delete(additionalProperties, "author")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCompletedApprovalRequesterComment struct {
	value *CompletedApprovalRequesterComment
	isSet bool
}

func (v NullableCompletedApprovalRequesterComment) Get() *CompletedApprovalRequesterComment {
	return v.value
}

func (v *NullableCompletedApprovalRequesterComment) Set(val *CompletedApprovalRequesterComment) {
	v.value = val
	v.isSet = true
}

func (v NullableCompletedApprovalRequesterComment) IsSet() bool {
	return v.isSet
}

func (v *NullableCompletedApprovalRequesterComment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompletedApprovalRequesterComment(val *CompletedApprovalRequesterComment) *NullableCompletedApprovalRequesterComment {
	return &NullableCompletedApprovalRequesterComment{value: val, isSet: true}
}

func (v NullableCompletedApprovalRequesterComment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompletedApprovalRequesterComment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


