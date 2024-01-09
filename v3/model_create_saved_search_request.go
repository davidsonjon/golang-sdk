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
	"fmt"
)

// checks if the CreateSavedSearchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSavedSearchRequest{}

// CreateSavedSearchRequest struct for CreateSavedSearchRequest
type CreateSavedSearchRequest struct {
	// The name of the saved search. 
	Name *string `json:"name,omitempty"`
	// The description of the saved search. 
	Description NullableString `json:"description,omitempty"`
	// A date-time in ISO-8601 format
	Created NullableTime `json:"created,omitempty"`
	// A date-time in ISO-8601 format
	Modified NullableTime `json:"modified,omitempty"`
	// The names of the Elasticsearch indices in which to search. 
	Indices []Index `json:"indices"`
	// The columns to be returned (specifies the order in which they will be presented) for each document type.  The currently supported document types are: _accessprofile_, _accountactivity_, _account_, _aggregation_, _entitlement_, _event_, _identity_, and _role_. 
	Columns *map[string][]Column `json:"columns,omitempty"`
	// The search query using Elasticsearch [Query String Query](https://www.elastic.co/guide/en/elasticsearch/reference/5.2/query-dsl-query-string-query.html#query-string) syntax from the Query DSL. 
	Query string `json:"query"`
	// The fields to be searched against in a multi-field query. 
	Fields []string `json:"fields,omitempty"`
	// The fields to be used to sort the search results. 
	Sort []string `json:"sort,omitempty"`
	Filters NullableSavedSearchDetailFilters `json:"filters,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateSavedSearchRequest CreateSavedSearchRequest

// NewCreateSavedSearchRequest instantiates a new CreateSavedSearchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSavedSearchRequest(indices []Index, query string) *CreateSavedSearchRequest {
	this := CreateSavedSearchRequest{}
	this.Indices = indices
	this.Query = query
	return &this
}

// NewCreateSavedSearchRequestWithDefaults instantiates a new CreateSavedSearchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSavedSearchRequestWithDefaults() *CreateSavedSearchRequest {
	this := CreateSavedSearchRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateSavedSearchRequest) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSavedSearchRequest) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateSavedSearchRequest) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateSavedSearchRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateSavedSearchRequest) GetDescription() string {
	if o == nil || isNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateSavedSearchRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateSavedSearchRequest) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *CreateSavedSearchRequest) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *CreateSavedSearchRequest) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *CreateSavedSearchRequest) UnsetDescription() {
	o.Description.Unset()
}

// GetCreated returns the Created field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateSavedSearchRequest) GetCreated() time.Time {
	if o == nil || isNil(o.Created.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateSavedSearchRequest) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// HasCreated returns a boolean if a field has been set.
func (o *CreateSavedSearchRequest) HasCreated() bool {
	if o != nil && o.Created.IsSet() {
		return true
	}

	return false
}

// SetCreated gets a reference to the given NullableTime and assigns it to the Created field.
func (o *CreateSavedSearchRequest) SetCreated(v time.Time) {
	o.Created.Set(&v)
}
// SetCreatedNil sets the value for Created to be an explicit nil
func (o *CreateSavedSearchRequest) SetCreatedNil() {
	o.Created.Set(nil)
}

// UnsetCreated ensures that no value is present for Created, not even an explicit nil
func (o *CreateSavedSearchRequest) UnsetCreated() {
	o.Created.Unset()
}

// GetModified returns the Modified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateSavedSearchRequest) GetModified() time.Time {
	if o == nil || isNil(o.Modified.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Modified.Get()
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateSavedSearchRequest) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Modified.Get(), o.Modified.IsSet()
}

// HasModified returns a boolean if a field has been set.
func (o *CreateSavedSearchRequest) HasModified() bool {
	if o != nil && o.Modified.IsSet() {
		return true
	}

	return false
}

// SetModified gets a reference to the given NullableTime and assigns it to the Modified field.
func (o *CreateSavedSearchRequest) SetModified(v time.Time) {
	o.Modified.Set(&v)
}
// SetModifiedNil sets the value for Modified to be an explicit nil
func (o *CreateSavedSearchRequest) SetModifiedNil() {
	o.Modified.Set(nil)
}

// UnsetModified ensures that no value is present for Modified, not even an explicit nil
func (o *CreateSavedSearchRequest) UnsetModified() {
	o.Modified.Unset()
}

// GetIndices returns the Indices field value
func (o *CreateSavedSearchRequest) GetIndices() []Index {
	if o == nil {
		var ret []Index
		return ret
	}

	return o.Indices
}

// GetIndicesOk returns a tuple with the Indices field value
// and a boolean to check if the value has been set.
func (o *CreateSavedSearchRequest) GetIndicesOk() ([]Index, bool) {
	if o == nil {
		return nil, false
	}
	return o.Indices, true
}

// SetIndices sets field value
func (o *CreateSavedSearchRequest) SetIndices(v []Index) {
	o.Indices = v
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *CreateSavedSearchRequest) GetColumns() map[string][]Column {
	if o == nil || isNil(o.Columns) {
		var ret map[string][]Column
		return ret
	}
	return *o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSavedSearchRequest) GetColumnsOk() (*map[string][]Column, bool) {
	if o == nil || isNil(o.Columns) {
		return nil, false
	}
	return o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *CreateSavedSearchRequest) HasColumns() bool {
	if o != nil && !isNil(o.Columns) {
		return true
	}

	return false
}

// SetColumns gets a reference to the given map[string][]Column and assigns it to the Columns field.
func (o *CreateSavedSearchRequest) SetColumns(v map[string][]Column) {
	o.Columns = &v
}

// GetQuery returns the Query field value
func (o *CreateSavedSearchRequest) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *CreateSavedSearchRequest) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *CreateSavedSearchRequest) SetQuery(v string) {
	o.Query = v
}

// GetFields returns the Fields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateSavedSearchRequest) GetFields() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateSavedSearchRequest) GetFieldsOk() ([]string, bool) {
	if o == nil || isNil(o.Fields) {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *CreateSavedSearchRequest) HasFields() bool {
	if o != nil && isNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given []string and assigns it to the Fields field.
func (o *CreateSavedSearchRequest) SetFields(v []string) {
	o.Fields = v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *CreateSavedSearchRequest) GetSort() []string {
	if o == nil || isNil(o.Sort) {
		var ret []string
		return ret
	}
	return o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSavedSearchRequest) GetSortOk() ([]string, bool) {
	if o == nil || isNil(o.Sort) {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *CreateSavedSearchRequest) HasSort() bool {
	if o != nil && !isNil(o.Sort) {
		return true
	}

	return false
}

// SetSort gets a reference to the given []string and assigns it to the Sort field.
func (o *CreateSavedSearchRequest) SetSort(v []string) {
	o.Sort = v
}

// GetFilters returns the Filters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateSavedSearchRequest) GetFilters() SavedSearchDetailFilters {
	if o == nil || isNil(o.Filters.Get()) {
		var ret SavedSearchDetailFilters
		return ret
	}
	return *o.Filters.Get()
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateSavedSearchRequest) GetFiltersOk() (*SavedSearchDetailFilters, bool) {
	if o == nil {
		return nil, false
	}
	return o.Filters.Get(), o.Filters.IsSet()
}

// HasFilters returns a boolean if a field has been set.
func (o *CreateSavedSearchRequest) HasFilters() bool {
	if o != nil && o.Filters.IsSet() {
		return true
	}

	return false
}

// SetFilters gets a reference to the given NullableSavedSearchDetailFilters and assigns it to the Filters field.
func (o *CreateSavedSearchRequest) SetFilters(v SavedSearchDetailFilters) {
	o.Filters.Set(&v)
}
// SetFiltersNil sets the value for Filters to be an explicit nil
func (o *CreateSavedSearchRequest) SetFiltersNil() {
	o.Filters.Set(nil)
}

// UnsetFilters ensures that no value is present for Filters, not even an explicit nil
func (o *CreateSavedSearchRequest) UnsetFilters() {
	o.Filters.Unset()
}

func (o CreateSavedSearchRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSavedSearchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Created.IsSet() {
		toSerialize["created"] = o.Created.Get()
	}
	if o.Modified.IsSet() {
		toSerialize["modified"] = o.Modified.Get()
	}
	toSerialize["indices"] = o.Indices
	if !isNil(o.Columns) {
		toSerialize["columns"] = o.Columns
	}
	toSerialize["query"] = o.Query
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	if !isNil(o.Sort) {
		toSerialize["sort"] = o.Sort
	}
	if o.Filters.IsSet() {
		toSerialize["filters"] = o.Filters.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateSavedSearchRequest) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"indices",
		"query",
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

	varCreateSavedSearchRequest := _CreateSavedSearchRequest{}

	if err = json.Unmarshal(bytes, &varCreateSavedSearchRequest); err == nil {
	*o = CreateSavedSearchRequest(varCreateSavedSearchRequest)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "indices")
		delete(additionalProperties, "columns")
		delete(additionalProperties, "query")
		delete(additionalProperties, "fields")
		delete(additionalProperties, "sort")
		delete(additionalProperties, "filters")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateSavedSearchRequest struct {
	value *CreateSavedSearchRequest
	isSet bool
}

func (v NullableCreateSavedSearchRequest) Get() *CreateSavedSearchRequest {
	return v.value
}

func (v *NullableCreateSavedSearchRequest) Set(val *CreateSavedSearchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSavedSearchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSavedSearchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSavedSearchRequest(val *CreateSavedSearchRequest) *NullableCreateSavedSearchRequest {
	return &NullableCreateSavedSearchRequest{value: val, isSet: true}
}

func (v NullableCreateSavedSearchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSavedSearchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


