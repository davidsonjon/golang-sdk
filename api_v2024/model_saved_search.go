/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the SavedSearch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SavedSearch{}

// SavedSearch struct for SavedSearch
type SavedSearch struct {
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
	// Sort by index. This takes precedence over the `sort` property. 
	OrderBy map[string][]string `json:"orderBy,omitempty"`
	// The fields to be used to sort the search results. 
	Sort []string `json:"sort,omitempty"`
	Filters NullableSavedSearchDetailFilters `json:"filters,omitempty"`
	// The saved search ID. 
	Id *string `json:"id,omitempty"`
	Owner *TypedReference `json:"owner,omitempty"`
	// The ID of the identity that owns this saved search.
	OwnerId *string `json:"ownerId,omitempty"`
	// Whether this saved search is visible to anyone but the owner. This field will always be false as there is no way to set a saved search as public at this time.
	Public *bool `json:"public,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SavedSearch SavedSearch

// NewSavedSearch instantiates a new SavedSearch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSavedSearch(indices []Index, query string) *SavedSearch {
	this := SavedSearch{}
	this.Indices = indices
	this.Query = query
	var public bool = false
	this.Public = &public
	return &this
}

// NewSavedSearchWithDefaults instantiates a new SavedSearch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSavedSearchWithDefaults() *SavedSearch {
	this := SavedSearch{}
	var public bool = false
	this.Public = &public
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SavedSearch) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedSearch) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SavedSearch) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SavedSearch) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SavedSearch) GetDescription() string {
	if o == nil || isNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SavedSearch) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *SavedSearch) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *SavedSearch) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *SavedSearch) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *SavedSearch) UnsetDescription() {
	o.Description.Unset()
}

// GetCreated returns the Created field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SavedSearch) GetCreated() time.Time {
	if o == nil || isNil(o.Created.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SavedSearch) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// HasCreated returns a boolean if a field has been set.
func (o *SavedSearch) HasCreated() bool {
	if o != nil && o.Created.IsSet() {
		return true
	}

	return false
}

// SetCreated gets a reference to the given NullableTime and assigns it to the Created field.
func (o *SavedSearch) SetCreated(v time.Time) {
	o.Created.Set(&v)
}
// SetCreatedNil sets the value for Created to be an explicit nil
func (o *SavedSearch) SetCreatedNil() {
	o.Created.Set(nil)
}

// UnsetCreated ensures that no value is present for Created, not even an explicit nil
func (o *SavedSearch) UnsetCreated() {
	o.Created.Unset()
}

// GetModified returns the Modified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SavedSearch) GetModified() time.Time {
	if o == nil || isNil(o.Modified.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Modified.Get()
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SavedSearch) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Modified.Get(), o.Modified.IsSet()
}

// HasModified returns a boolean if a field has been set.
func (o *SavedSearch) HasModified() bool {
	if o != nil && o.Modified.IsSet() {
		return true
	}

	return false
}

// SetModified gets a reference to the given NullableTime and assigns it to the Modified field.
func (o *SavedSearch) SetModified(v time.Time) {
	o.Modified.Set(&v)
}
// SetModifiedNil sets the value for Modified to be an explicit nil
func (o *SavedSearch) SetModifiedNil() {
	o.Modified.Set(nil)
}

// UnsetModified ensures that no value is present for Modified, not even an explicit nil
func (o *SavedSearch) UnsetModified() {
	o.Modified.Unset()
}

// GetIndices returns the Indices field value
func (o *SavedSearch) GetIndices() []Index {
	if o == nil {
		var ret []Index
		return ret
	}

	return o.Indices
}

// GetIndicesOk returns a tuple with the Indices field value
// and a boolean to check if the value has been set.
func (o *SavedSearch) GetIndicesOk() ([]Index, bool) {
	if o == nil {
		return nil, false
	}
	return o.Indices, true
}

// SetIndices sets field value
func (o *SavedSearch) SetIndices(v []Index) {
	o.Indices = v
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *SavedSearch) GetColumns() map[string][]Column {
	if o == nil || isNil(o.Columns) {
		var ret map[string][]Column
		return ret
	}
	return *o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedSearch) GetColumnsOk() (*map[string][]Column, bool) {
	if o == nil || isNil(o.Columns) {
		return nil, false
	}
	return o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *SavedSearch) HasColumns() bool {
	if o != nil && !isNil(o.Columns) {
		return true
	}

	return false
}

// SetColumns gets a reference to the given map[string][]Column and assigns it to the Columns field.
func (o *SavedSearch) SetColumns(v map[string][]Column) {
	o.Columns = &v
}

// GetQuery returns the Query field value
func (o *SavedSearch) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *SavedSearch) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *SavedSearch) SetQuery(v string) {
	o.Query = v
}

// GetFields returns the Fields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SavedSearch) GetFields() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SavedSearch) GetFieldsOk() ([]string, bool) {
	if o == nil || isNil(o.Fields) {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *SavedSearch) HasFields() bool {
	if o != nil && isNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given []string and assigns it to the Fields field.
func (o *SavedSearch) SetFields(v []string) {
	o.Fields = v
}

// GetOrderBy returns the OrderBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SavedSearch) GetOrderBy() map[string][]string {
	if o == nil {
		var ret map[string][]string
		return ret
	}
	return o.OrderBy
}

// GetOrderByOk returns a tuple with the OrderBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SavedSearch) GetOrderByOk() (*map[string][]string, bool) {
	if o == nil || isNil(o.OrderBy) {
		return nil, false
	}
	return &o.OrderBy, true
}

// HasOrderBy returns a boolean if a field has been set.
func (o *SavedSearch) HasOrderBy() bool {
	if o != nil && isNil(o.OrderBy) {
		return true
	}

	return false
}

// SetOrderBy gets a reference to the given map[string][]string and assigns it to the OrderBy field.
func (o *SavedSearch) SetOrderBy(v map[string][]string) {
	o.OrderBy = v
}

// GetSort returns the Sort field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SavedSearch) GetSort() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SavedSearch) GetSortOk() ([]string, bool) {
	if o == nil || isNil(o.Sort) {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *SavedSearch) HasSort() bool {
	if o != nil && isNil(o.Sort) {
		return true
	}

	return false
}

// SetSort gets a reference to the given []string and assigns it to the Sort field.
func (o *SavedSearch) SetSort(v []string) {
	o.Sort = v
}

// GetFilters returns the Filters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SavedSearch) GetFilters() SavedSearchDetailFilters {
	if o == nil || isNil(o.Filters.Get()) {
		var ret SavedSearchDetailFilters
		return ret
	}
	return *o.Filters.Get()
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SavedSearch) GetFiltersOk() (*SavedSearchDetailFilters, bool) {
	if o == nil {
		return nil, false
	}
	return o.Filters.Get(), o.Filters.IsSet()
}

// HasFilters returns a boolean if a field has been set.
func (o *SavedSearch) HasFilters() bool {
	if o != nil && o.Filters.IsSet() {
		return true
	}

	return false
}

// SetFilters gets a reference to the given NullableSavedSearchDetailFilters and assigns it to the Filters field.
func (o *SavedSearch) SetFilters(v SavedSearchDetailFilters) {
	o.Filters.Set(&v)
}
// SetFiltersNil sets the value for Filters to be an explicit nil
func (o *SavedSearch) SetFiltersNil() {
	o.Filters.Set(nil)
}

// UnsetFilters ensures that no value is present for Filters, not even an explicit nil
func (o *SavedSearch) UnsetFilters() {
	o.Filters.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SavedSearch) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedSearch) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SavedSearch) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SavedSearch) SetId(v string) {
	o.Id = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *SavedSearch) GetOwner() TypedReference {
	if o == nil || isNil(o.Owner) {
		var ret TypedReference
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedSearch) GetOwnerOk() (*TypedReference, bool) {
	if o == nil || isNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *SavedSearch) HasOwner() bool {
	if o != nil && !isNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given TypedReference and assigns it to the Owner field.
func (o *SavedSearch) SetOwner(v TypedReference) {
	o.Owner = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *SavedSearch) GetOwnerId() string {
	if o == nil || isNil(o.OwnerId) {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedSearch) GetOwnerIdOk() (*string, bool) {
	if o == nil || isNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *SavedSearch) HasOwnerId() bool {
	if o != nil && !isNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *SavedSearch) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetPublic returns the Public field value if set, zero value otherwise.
func (o *SavedSearch) GetPublic() bool {
	if o == nil || isNil(o.Public) {
		var ret bool
		return ret
	}
	return *o.Public
}

// GetPublicOk returns a tuple with the Public field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedSearch) GetPublicOk() (*bool, bool) {
	if o == nil || isNil(o.Public) {
		return nil, false
	}
	return o.Public, true
}

// HasPublic returns a boolean if a field has been set.
func (o *SavedSearch) HasPublic() bool {
	if o != nil && !isNil(o.Public) {
		return true
	}

	return false
}

// SetPublic gets a reference to the given bool and assigns it to the Public field.
func (o *SavedSearch) SetPublic(v bool) {
	o.Public = &v
}

func (o SavedSearch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SavedSearch) ToMap() (map[string]interface{}, error) {
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
	if o.OrderBy != nil {
		toSerialize["orderBy"] = o.OrderBy
	}
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}
	if o.Filters.IsSet() {
		toSerialize["filters"] = o.Filters.Get()
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !isNil(o.OwnerId) {
		toSerialize["ownerId"] = o.OwnerId
	}
	if !isNil(o.Public) {
		toSerialize["public"] = o.Public
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SavedSearch) UnmarshalJSON(bytes []byte) (err error) {
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

	varSavedSearch := _SavedSearch{}

	if err = json.Unmarshal(bytes, &varSavedSearch); err == nil {
	*o = SavedSearch(varSavedSearch)
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
		delete(additionalProperties, "orderBy")
		delete(additionalProperties, "sort")
		delete(additionalProperties, "filters")
		delete(additionalProperties, "id")
		delete(additionalProperties, "owner")
		delete(additionalProperties, "ownerId")
		delete(additionalProperties, "public")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSavedSearch struct {
	value *SavedSearch
	isSet bool
}

func (v NullableSavedSearch) Get() *SavedSearch {
	return v.value
}

func (v *NullableSavedSearch) Set(val *SavedSearch) {
	v.value = val
	v.isSet = true
}

func (v NullableSavedSearch) IsSet() bool {
	return v.isSet
}

func (v *NullableSavedSearch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSavedSearch(val *SavedSearch) *NullableSavedSearch {
	return &NullableSavedSearch{value: val, isSet: true}
}

func (v NullableSavedSearch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSavedSearch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


