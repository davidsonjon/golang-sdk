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

// checks if the RequestedItemStatusRequesterComment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestedItemStatusRequesterComment{}

// RequestedItemStatusRequesterComment struct for RequestedItemStatusRequesterComment
type RequestedItemStatusRequesterComment struct {
	// Comment content.
	Comment NullableString `json:"comment,omitempty"`
	// Date and time comment was created.
	Created *time.Time `json:"created,omitempty"`
	Author *CommentDto1Author `json:"author,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RequestedItemStatusRequesterComment RequestedItemStatusRequesterComment

// NewRequestedItemStatusRequesterComment instantiates a new RequestedItemStatusRequesterComment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestedItemStatusRequesterComment() *RequestedItemStatusRequesterComment {
	this := RequestedItemStatusRequesterComment{}
	return &this
}

// NewRequestedItemStatusRequesterCommentWithDefaults instantiates a new RequestedItemStatusRequesterComment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestedItemStatusRequesterCommentWithDefaults() *RequestedItemStatusRequesterComment {
	this := RequestedItemStatusRequesterComment{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestedItemStatusRequesterComment) GetComment() string {
	if o == nil || IsNil(o.Comment.Get()) {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestedItemStatusRequesterComment) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *RequestedItemStatusRequesterComment) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *RequestedItemStatusRequesterComment) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *RequestedItemStatusRequesterComment) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *RequestedItemStatusRequesterComment) UnsetComment() {
	o.Comment.Unset()
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *RequestedItemStatusRequesterComment) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestedItemStatusRequesterComment) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *RequestedItemStatusRequesterComment) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *RequestedItemStatusRequesterComment) SetCreated(v time.Time) {
	o.Created = &v
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *RequestedItemStatusRequesterComment) GetAuthor() CommentDto1Author {
	if o == nil || IsNil(o.Author) {
		var ret CommentDto1Author
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestedItemStatusRequesterComment) GetAuthorOk() (*CommentDto1Author, bool) {
	if o == nil || IsNil(o.Author) {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *RequestedItemStatusRequesterComment) HasAuthor() bool {
	if o != nil && !IsNil(o.Author) {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given CommentDto1Author and assigns it to the Author field.
func (o *RequestedItemStatusRequesterComment) SetAuthor(v CommentDto1Author) {
	o.Author = &v
}

func (o RequestedItemStatusRequesterComment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestedItemStatusRequesterComment) ToMap() (map[string]interface{}, error) {
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

func (o *RequestedItemStatusRequesterComment) UnmarshalJSON(data []byte) (err error) {
	varRequestedItemStatusRequesterComment := _RequestedItemStatusRequesterComment{}

	err = json.Unmarshal(data, &varRequestedItemStatusRequesterComment)

	if err != nil {
		return err
	}

	*o = RequestedItemStatusRequesterComment(varRequestedItemStatusRequesterComment)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "comment")
		delete(additionalProperties, "created")
		delete(additionalProperties, "author")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestedItemStatusRequesterComment struct {
	value *RequestedItemStatusRequesterComment
	isSet bool
}

func (v NullableRequestedItemStatusRequesterComment) Get() *RequestedItemStatusRequesterComment {
	return v.value
}

func (v *NullableRequestedItemStatusRequesterComment) Set(val *RequestedItemStatusRequesterComment) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestedItemStatusRequesterComment) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestedItemStatusRequesterComment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestedItemStatusRequesterComment(val *RequestedItemStatusRequesterComment) *NullableRequestedItemStatusRequesterComment {
	return &NullableRequestedItemStatusRequesterComment{value: val, isSet: true}
}

func (v NullableRequestedItemStatusRequesterComment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestedItemStatusRequesterComment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


