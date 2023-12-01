/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
)

// checks if the Query type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Query{}

// Query Query parameters used to construct an Elasticsearch query object.
type Query struct {
	// The query using the Elasticsearch [Query String Query](https://www.elastic.co/guide/en/elasticsearch/reference/5.2/query-dsl-query-string-query.html#query-string) syntax from the Query DSL extended by SailPoint to support Nested queries.
	Query *string `json:"query,omitempty"`
	// The fields the query will be applied to.  Fields provide you with a simple way to add additional fields to search, without making the query too complicated.  For example, you can use the fields to specify that you want your query of \"a*\" to be applied to \"name\", \"firstName\", and the \"source.name\".  The response will include all results matching the \"a*\" query found in those three fields.  A field's availability depends on the indices being searched.  For example, if you are searching \"identities\", you can apply your search to the \"firstName\" field, but you couldn't use \"firstName\" with a search on \"access profiles\".  Refer to the response schema for the respective lists of available fields. 
	Fields []string `json:"fields,omitempty"`
	// The time zone to be applied to any range query related to dates.
	TimeZone *string `json:"timeZone,omitempty"`
	InnerHit *InnerHit `json:"innerHit,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Query Query

// NewQuery instantiates a new Query object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuery() *Query {
	this := Query{}
	return &this
}

// NewQueryWithDefaults instantiates a new Query object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryWithDefaults() *Query {
	this := Query{}
	return &this
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *Query) GetQuery() string {
	if o == nil || isNil(o.Query) {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetQueryOk() (*string, bool) {
	if o == nil || isNil(o.Query) {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *Query) HasQuery() bool {
	if o != nil && !isNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *Query) SetQuery(v string) {
	o.Query = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *Query) GetFields() []string {
	if o == nil || isNil(o.Fields) {
		var ret []string
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetFieldsOk() ([]string, bool) {
	if o == nil || isNil(o.Fields) {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *Query) HasFields() bool {
	if o != nil && !isNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given []string and assigns it to the Fields field.
func (o *Query) SetFields(v []string) {
	o.Fields = v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *Query) GetTimeZone() string {
	if o == nil || isNil(o.TimeZone) {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetTimeZoneOk() (*string, bool) {
	if o == nil || isNil(o.TimeZone) {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *Query) HasTimeZone() bool {
	if o != nil && !isNil(o.TimeZone) {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *Query) SetTimeZone(v string) {
	o.TimeZone = &v
}

// GetInnerHit returns the InnerHit field value if set, zero value otherwise.
func (o *Query) GetInnerHit() InnerHit {
	if o == nil || isNil(o.InnerHit) {
		var ret InnerHit
		return ret
	}
	return *o.InnerHit
}

// GetInnerHitOk returns a tuple with the InnerHit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetInnerHitOk() (*InnerHit, bool) {
	if o == nil || isNil(o.InnerHit) {
		return nil, false
	}
	return o.InnerHit, true
}

// HasInnerHit returns a boolean if a field has been set.
func (o *Query) HasInnerHit() bool {
	if o != nil && !isNil(o.InnerHit) {
		return true
	}

	return false
}

// SetInnerHit gets a reference to the given InnerHit and assigns it to the InnerHit field.
func (o *Query) SetInnerHit(v InnerHit) {
	o.InnerHit = &v
}

func (o Query) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Query) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Query) {
		toSerialize["query"] = o.Query
	}
	if !isNil(o.Fields) {
		toSerialize["fields"] = o.Fields
	}
	if !isNil(o.TimeZone) {
		toSerialize["timeZone"] = o.TimeZone
	}
	if !isNil(o.InnerHit) {
		toSerialize["innerHit"] = o.InnerHit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Query) UnmarshalJSON(bytes []byte) (err error) {
	varQuery := _Query{}

	if err = json.Unmarshal(bytes, &varQuery); err == nil {
	*o = Query(varQuery)
}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "query")
		delete(additionalProperties, "fields")
		delete(additionalProperties, "timeZone")
		delete(additionalProperties, "innerHit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQuery struct {
	value *Query
	isSet bool
}

func (v NullableQuery) Get() *Query {
	return v.value
}

func (v *NullableQuery) Set(val *Query) {
	v.value = val
	v.isSet = true
}

func (v NullableQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuery(val *Query) *NullableQuery {
	return &NullableQuery{value: val, isSet: true}
}

func (v NullableQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


