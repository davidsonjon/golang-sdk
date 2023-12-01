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

// checks if the CompletedApprovalReviewerComment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompletedApprovalReviewerComment{}

// CompletedApprovalReviewerComment The approval's reviewer's comment.
type CompletedApprovalReviewerComment struct {
	// Comment content.
	Comment NullableString `json:"comment,omitempty"`
	Author *CommentDtoAuthor `json:"author,omitempty"`
	// Date and time comment was created.
	Created *time.Time `json:"created,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CompletedApprovalReviewerComment CompletedApprovalReviewerComment

// NewCompletedApprovalReviewerComment instantiates a new CompletedApprovalReviewerComment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompletedApprovalReviewerComment() *CompletedApprovalReviewerComment {
	this := CompletedApprovalReviewerComment{}
	return &this
}

// NewCompletedApprovalReviewerCommentWithDefaults instantiates a new CompletedApprovalReviewerComment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompletedApprovalReviewerCommentWithDefaults() *CompletedApprovalReviewerComment {
	this := CompletedApprovalReviewerComment{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompletedApprovalReviewerComment) GetComment() string {
	if o == nil || isNil(o.Comment.Get()) {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompletedApprovalReviewerComment) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *CompletedApprovalReviewerComment) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *CompletedApprovalReviewerComment) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *CompletedApprovalReviewerComment) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *CompletedApprovalReviewerComment) UnsetComment() {
	o.Comment.Unset()
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *CompletedApprovalReviewerComment) GetAuthor() CommentDtoAuthor {
	if o == nil || isNil(o.Author) {
		var ret CommentDtoAuthor
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompletedApprovalReviewerComment) GetAuthorOk() (*CommentDtoAuthor, bool) {
	if o == nil || isNil(o.Author) {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *CompletedApprovalReviewerComment) HasAuthor() bool {
	if o != nil && !isNil(o.Author) {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given CommentDtoAuthor and assigns it to the Author field.
func (o *CompletedApprovalReviewerComment) SetAuthor(v CommentDtoAuthor) {
	o.Author = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *CompletedApprovalReviewerComment) GetCreated() time.Time {
	if o == nil || isNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompletedApprovalReviewerComment) GetCreatedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *CompletedApprovalReviewerComment) HasCreated() bool {
	if o != nil && !isNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *CompletedApprovalReviewerComment) SetCreated(v time.Time) {
	o.Created = &v
}

func (o CompletedApprovalReviewerComment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompletedApprovalReviewerComment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Comment.IsSet() {
		toSerialize["comment"] = o.Comment.Get()
	}
	if !isNil(o.Author) {
		toSerialize["author"] = o.Author
	}
	if !isNil(o.Created) {
		toSerialize["created"] = o.Created
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CompletedApprovalReviewerComment) UnmarshalJSON(bytes []byte) (err error) {
	varCompletedApprovalReviewerComment := _CompletedApprovalReviewerComment{}

	if err = json.Unmarshal(bytes, &varCompletedApprovalReviewerComment); err == nil {
	*o = CompletedApprovalReviewerComment(varCompletedApprovalReviewerComment)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "comment")
		delete(additionalProperties, "author")
		delete(additionalProperties, "created")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCompletedApprovalReviewerComment struct {
	value *CompletedApprovalReviewerComment
	isSet bool
}

func (v NullableCompletedApprovalReviewerComment) Get() *CompletedApprovalReviewerComment {
	return v.value
}

func (v *NullableCompletedApprovalReviewerComment) Set(val *CompletedApprovalReviewerComment) {
	v.value = val
	v.isSet = true
}

func (v NullableCompletedApprovalReviewerComment) IsSet() bool {
	return v.isSet
}

func (v *NullableCompletedApprovalReviewerComment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompletedApprovalReviewerComment(val *CompletedApprovalReviewerComment) *NullableCompletedApprovalReviewerComment {
	return &NullableCompletedApprovalReviewerComment{value: val, isSet: true}
}

func (v NullableCompletedApprovalReviewerComment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompletedApprovalReviewerComment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


