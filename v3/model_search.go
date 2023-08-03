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

// checks if the Search type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Search{}

// Search struct for Search
type Search struct {
	// The names of the Elasticsearch indices in which to search. If none are provided, then all indices will be searched.
	Indices []Index `json:"indices,omitempty"`
	QueryType *QueryType `json:"queryType,omitempty"`
	QueryVersion *string `json:"queryVersion,omitempty"`
	Query *Query `json:"query,omitempty"`
	// The search query using the Elasticsearch [Query DSL](https://www.elastic.co/guide/en/elasticsearch/reference/7.10/query-dsl.html) syntax.
	QueryDsl map[string]interface{} `json:"queryDsl,omitempty"`
	TextQuery *TextQuery `json:"textQuery,omitempty"`
	TypeAheadQuery *TypeAheadQuery `json:"typeAheadQuery,omitempty"`
	// Indicates whether nested objects from returned search results should be included.
	IncludeNested *bool `json:"includeNested,omitempty"`
	QueryResultFilter *QueryResultFilter `json:"queryResultFilter,omitempty"`
	AggregationType *AggregationType `json:"aggregationType,omitempty"`
	AggregationsVersion *string `json:"aggregationsVersion,omitempty"`
	// The aggregation search query using Elasticsearch [Aggregations](https://www.elastic.co/guide/en/elasticsearch/reference/5.2/search-aggregations.html) syntax.
	AggregationsDsl map[string]interface{} `json:"aggregationsDsl,omitempty"`
	Aggregations *SearchAggregationSpecification `json:"aggregations,omitempty"`
	// The fields to be used to sort the search results. Use + or - to specify the sort direction.
	Sort []string `json:"sort,omitempty"`
	// Used to begin the search window at the values specified. This parameter consists of the last values of the sorted fields in the current record set. This is used to expand the Elasticsearch limit of 10K records by shifting the 10K window to begin at this value. It is recommended that you always include the ID of the object in addition to any other fields on this parameter in order to ensure you don't get duplicate results while paging. For example, when searching for identities, if you are sorting by displayName you will also want to include ID, for example [\"displayName\", \"id\"].  If the last identity ID in the search result is 2c91808375d8e80a0175e1f88a575221 and the last displayName is \"John Doe\", then using that displayName and ID will start a new search after this identity. The searchAfter value will look like [\"John Doe\",\"2c91808375d8e80a0175e1f88a575221\"]
	SearchAfter []string `json:"searchAfter,omitempty"`
	// The filters to be applied for each filtered field name.
	Filters *map[string]Filter `json:"filters,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Search Search

// NewSearch instantiates a new Search object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearch() *Search {
	this := Search{}
	var queryType QueryType = QUERYTYPE_SAILPOINT
	this.QueryType = &queryType
	var includeNested bool = true
	this.IncludeNested = &includeNested
	var aggregationType AggregationType = AGGREGATIONTYPE_DSL
	this.AggregationType = &aggregationType
	return &this
}

// NewSearchWithDefaults instantiates a new Search object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchWithDefaults() *Search {
	this := Search{}
	var queryType QueryType = QUERYTYPE_SAILPOINT
	this.QueryType = &queryType
	var includeNested bool = true
	this.IncludeNested = &includeNested
	var aggregationType AggregationType = AGGREGATIONTYPE_DSL
	this.AggregationType = &aggregationType
	return &this
}

// GetIndices returns the Indices field value if set, zero value otherwise.
func (o *Search) GetIndices() []Index {
	if o == nil || isNil(o.Indices) {
		var ret []Index
		return ret
	}
	return o.Indices
}

// GetIndicesOk returns a tuple with the Indices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Search) GetIndicesOk() ([]Index, bool) {
	if o == nil || isNil(o.Indices) {
		return nil, false
	}
	return o.Indices, true
}

// HasIndices returns a boolean if a field has been set.
func (o *Search) HasIndices() bool {
	if o != nil && !isNil(o.Indices) {
		return true
	}

	return false
}

// SetIndices gets a reference to the given []Index and assigns it to the Indices field.
func (o *Search) SetIndices(v []Index) {
	o.Indices = v
}

// GetQueryType returns the QueryType field value if set, zero value otherwise.
func (o *Search) GetQueryType() QueryType {
	if o == nil || isNil(o.QueryType) {
		var ret QueryType
		return ret
	}
	return *o.QueryType
}

// GetQueryTypeOk returns a tuple with the QueryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Search) GetQueryTypeOk() (*QueryType, bool) {
	if o == nil || isNil(o.QueryType) {
		return nil, false
	}
	return o.QueryType, true
}

// HasQueryType returns a boolean if a field has been set.
func (o *Search) HasQueryType() bool {
	if o != nil && !isNil(o.QueryType) {
		return true
	}

	return false
}

// SetQueryType gets a reference to the given QueryType and assigns it to the QueryType field.
func (o *Search) SetQueryType(v QueryType) {
	o.QueryType = &v
}

// GetQueryVersion returns the QueryVersion field value if set, zero value otherwise.
func (o *Search) GetQueryVersion() string {
	if o == nil || isNil(o.QueryVersion) {
		var ret string
		return ret
	}
	return *o.QueryVersion
}

// GetQueryVersionOk returns a tuple with the QueryVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Search) GetQueryVersionOk() (*string, bool) {
	if o == nil || isNil(o.QueryVersion) {
		return nil, false
	}
	return o.QueryVersion, true
}

// HasQueryVersion returns a boolean if a field has been set.
func (o *Search) HasQueryVersion() bool {
	if o != nil && !isNil(o.QueryVersion) {
		return true
	}

	return false
}

// SetQueryVersion gets a reference to the given string and assigns it to the QueryVersion field.
func (o *Search) SetQueryVersion(v string) {
	o.QueryVersion = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *Search) GetQuery() Query {
	if o == nil || isNil(o.Query) {
		var ret Query
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Search) GetQueryOk() (*Query, bool) {
	if o == nil || isNil(o.Query) {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *Search) HasQuery() bool {
	if o != nil && !isNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given Query and assigns it to the Query field.
func (o *Search) SetQuery(v Query) {
	o.Query = &v
}

// GetQueryDsl returns the QueryDsl field value if set, zero value otherwise.
func (o *Search) GetQueryDsl() map[string]interface{} {
	if o == nil || isNil(o.QueryDsl) {
		var ret map[string]interface{}
		return ret
	}
	return o.QueryDsl
}

// GetQueryDslOk returns a tuple with the QueryDsl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Search) GetQueryDslOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.QueryDsl) {
		return map[string]interface{}{}, false
	}
	return o.QueryDsl, true
}

// HasQueryDsl returns a boolean if a field has been set.
func (o *Search) HasQueryDsl() bool {
	if o != nil && !isNil(o.QueryDsl) {
		return true
	}

	return false
}

// SetQueryDsl gets a reference to the given map[string]interface{} and assigns it to the QueryDsl field.
func (o *Search) SetQueryDsl(v map[string]interface{}) {
	o.QueryDsl = v
}

// GetTextQuery returns the TextQuery field value if set, zero value otherwise.
func (o *Search) GetTextQuery() TextQuery {
	if o == nil || isNil(o.TextQuery) {
		var ret TextQuery
		return ret
	}
	return *o.TextQuery
}

// GetTextQueryOk returns a tuple with the TextQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Search) GetTextQueryOk() (*TextQuery, bool) {
	if o == nil || isNil(o.TextQuery) {
		return nil, false
	}
	return o.TextQuery, true
}

// HasTextQuery returns a boolean if a field has been set.
func (o *Search) HasTextQuery() bool {
	if o != nil && !isNil(o.TextQuery) {
		return true
	}

	return false
}

// SetTextQuery gets a reference to the given TextQuery and assigns it to the TextQuery field.
func (o *Search) SetTextQuery(v TextQuery) {
	o.TextQuery = &v
}

// GetTypeAheadQuery returns the TypeAheadQuery field value if set, zero value otherwise.
func (o *Search) GetTypeAheadQuery() TypeAheadQuery {
	if o == nil || isNil(o.TypeAheadQuery) {
		var ret TypeAheadQuery
		return ret
	}
	return *o.TypeAheadQuery
}

// GetTypeAheadQueryOk returns a tuple with the TypeAheadQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Search) GetTypeAheadQueryOk() (*TypeAheadQuery, bool) {
	if o == nil || isNil(o.TypeAheadQuery) {
		return nil, false
	}
	return o.TypeAheadQuery, true
}

// HasTypeAheadQuery returns a boolean if a field has been set.
func (o *Search) HasTypeAheadQuery() bool {
	if o != nil && !isNil(o.TypeAheadQuery) {
		return true
	}

	return false
}

// SetTypeAheadQuery gets a reference to the given TypeAheadQuery and assigns it to the TypeAheadQuery field.
func (o *Search) SetTypeAheadQuery(v TypeAheadQuery) {
	o.TypeAheadQuery = &v
}

// GetIncludeNested returns the IncludeNested field value if set, zero value otherwise.
func (o *Search) GetIncludeNested() bool {
	if o == nil || isNil(o.IncludeNested) {
		var ret bool
		return ret
	}
	return *o.IncludeNested
}

// GetIncludeNestedOk returns a tuple with the IncludeNested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Search) GetIncludeNestedOk() (*bool, bool) {
	if o == nil || isNil(o.IncludeNested) {
		return nil, false
	}
	return o.IncludeNested, true
}

// HasIncludeNested returns a boolean if a field has been set.
func (o *Search) HasIncludeNested() bool {
	if o != nil && !isNil(o.IncludeNested) {
		return true
	}

	return false
}

// SetIncludeNested gets a reference to the given bool and assigns it to the IncludeNested field.
func (o *Search) SetIncludeNested(v bool) {
	o.IncludeNested = &v
}

// GetQueryResultFilter returns the QueryResultFilter field value if set, zero value otherwise.
func (o *Search) GetQueryResultFilter() QueryResultFilter {
	if o == nil || isNil(o.QueryResultFilter) {
		var ret QueryResultFilter
		return ret
	}
	return *o.QueryResultFilter
}

// GetQueryResultFilterOk returns a tuple with the QueryResultFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Search) GetQueryResultFilterOk() (*QueryResultFilter, bool) {
	if o == nil || isNil(o.QueryResultFilter) {
		return nil, false
	}
	return o.QueryResultFilter, true
}

// HasQueryResultFilter returns a boolean if a field has been set.
func (o *Search) HasQueryResultFilter() bool {
	if o != nil && !isNil(o.QueryResultFilter) {
		return true
	}

	return false
}

// SetQueryResultFilter gets a reference to the given QueryResultFilter and assigns it to the QueryResultFilter field.
func (o *Search) SetQueryResultFilter(v QueryResultFilter) {
	o.QueryResultFilter = &v
}

// GetAggregationType returns the AggregationType field value if set, zero value otherwise.
func (o *Search) GetAggregationType() AggregationType {
	if o == nil || isNil(o.AggregationType) {
		var ret AggregationType
		return ret
	}
	return *o.AggregationType
}

// GetAggregationTypeOk returns a tuple with the AggregationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Search) GetAggregationTypeOk() (*AggregationType, bool) {
	if o == nil || isNil(o.AggregationType) {
		return nil, false
	}
	return o.AggregationType, true
}

// HasAggregationType returns a boolean if a field has been set.
func (o *Search) HasAggregationType() bool {
	if o != nil && !isNil(o.AggregationType) {
		return true
	}

	return false
}

// SetAggregationType gets a reference to the given AggregationType and assigns it to the AggregationType field.
func (o *Search) SetAggregationType(v AggregationType) {
	o.AggregationType = &v
}

// GetAggregationsVersion returns the AggregationsVersion field value if set, zero value otherwise.
func (o *Search) GetAggregationsVersion() string {
	if o == nil || isNil(o.AggregationsVersion) {
		var ret string
		return ret
	}
	return *o.AggregationsVersion
}

// GetAggregationsVersionOk returns a tuple with the AggregationsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Search) GetAggregationsVersionOk() (*string, bool) {
	if o == nil || isNil(o.AggregationsVersion) {
		return nil, false
	}
	return o.AggregationsVersion, true
}

// HasAggregationsVersion returns a boolean if a field has been set.
func (o *Search) HasAggregationsVersion() bool {
	if o != nil && !isNil(o.AggregationsVersion) {
		return true
	}

	return false
}

// SetAggregationsVersion gets a reference to the given string and assigns it to the AggregationsVersion field.
func (o *Search) SetAggregationsVersion(v string) {
	o.AggregationsVersion = &v
}

// GetAggregationsDsl returns the AggregationsDsl field value if set, zero value otherwise.
func (o *Search) GetAggregationsDsl() map[string]interface{} {
	if o == nil || isNil(o.AggregationsDsl) {
		var ret map[string]interface{}
		return ret
	}
	return o.AggregationsDsl
}

// GetAggregationsDslOk returns a tuple with the AggregationsDsl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Search) GetAggregationsDslOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.AggregationsDsl) {
		return map[string]interface{}{}, false
	}
	return o.AggregationsDsl, true
}

// HasAggregationsDsl returns a boolean if a field has been set.
func (o *Search) HasAggregationsDsl() bool {
	if o != nil && !isNil(o.AggregationsDsl) {
		return true
	}

	return false
}

// SetAggregationsDsl gets a reference to the given map[string]interface{} and assigns it to the AggregationsDsl field.
func (o *Search) SetAggregationsDsl(v map[string]interface{}) {
	o.AggregationsDsl = v
}

// GetAggregations returns the Aggregations field value if set, zero value otherwise.
func (o *Search) GetAggregations() SearchAggregationSpecification {
	if o == nil || isNil(o.Aggregations) {
		var ret SearchAggregationSpecification
		return ret
	}
	return *o.Aggregations
}

// GetAggregationsOk returns a tuple with the Aggregations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Search) GetAggregationsOk() (*SearchAggregationSpecification, bool) {
	if o == nil || isNil(o.Aggregations) {
		return nil, false
	}
	return o.Aggregations, true
}

// HasAggregations returns a boolean if a field has been set.
func (o *Search) HasAggregations() bool {
	if o != nil && !isNil(o.Aggregations) {
		return true
	}

	return false
}

// SetAggregations gets a reference to the given SearchAggregationSpecification and assigns it to the Aggregations field.
func (o *Search) SetAggregations(v SearchAggregationSpecification) {
	o.Aggregations = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *Search) GetSort() []string {
	if o == nil || isNil(o.Sort) {
		var ret []string
		return ret
	}
	return o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Search) GetSortOk() ([]string, bool) {
	if o == nil || isNil(o.Sort) {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *Search) HasSort() bool {
	if o != nil && !isNil(o.Sort) {
		return true
	}

	return false
}

// SetSort gets a reference to the given []string and assigns it to the Sort field.
func (o *Search) SetSort(v []string) {
	o.Sort = v
}

// GetSearchAfter returns the SearchAfter field value if set, zero value otherwise.
func (o *Search) GetSearchAfter() []string {
	if o == nil || isNil(o.SearchAfter) {
		var ret []string
		return ret
	}
	return o.SearchAfter
}

// GetSearchAfterOk returns a tuple with the SearchAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Search) GetSearchAfterOk() ([]string, bool) {
	if o == nil || isNil(o.SearchAfter) {
		return nil, false
	}
	return o.SearchAfter, true
}

// HasSearchAfter returns a boolean if a field has been set.
func (o *Search) HasSearchAfter() bool {
	if o != nil && !isNil(o.SearchAfter) {
		return true
	}

	return false
}

// SetSearchAfter gets a reference to the given []string and assigns it to the SearchAfter field.
func (o *Search) SetSearchAfter(v []string) {
	o.SearchAfter = v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *Search) GetFilters() map[string]Filter {
	if o == nil || isNil(o.Filters) {
		var ret map[string]Filter
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Search) GetFiltersOk() (*map[string]Filter, bool) {
	if o == nil || isNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *Search) HasFilters() bool {
	if o != nil && !isNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given map[string]Filter and assigns it to the Filters field.
func (o *Search) SetFilters(v map[string]Filter) {
	o.Filters = &v
}

func (o Search) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Search) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Indices) {
		toSerialize["indices"] = o.Indices
	}
	if !isNil(o.QueryType) {
		toSerialize["queryType"] = o.QueryType
	}
	if !isNil(o.QueryVersion) {
		toSerialize["queryVersion"] = o.QueryVersion
	}
	if !isNil(o.Query) {
		toSerialize["query"] = o.Query
	}
	if !isNil(o.QueryDsl) {
		toSerialize["queryDsl"] = o.QueryDsl
	}
	if !isNil(o.TextQuery) {
		toSerialize["textQuery"] = o.TextQuery
	}
	if !isNil(o.TypeAheadQuery) {
		toSerialize["typeAheadQuery"] = o.TypeAheadQuery
	}
	if !isNil(o.IncludeNested) {
		toSerialize["includeNested"] = o.IncludeNested
	}
	if !isNil(o.QueryResultFilter) {
		toSerialize["queryResultFilter"] = o.QueryResultFilter
	}
	if !isNil(o.AggregationType) {
		toSerialize["aggregationType"] = o.AggregationType
	}
	if !isNil(o.AggregationsVersion) {
		toSerialize["aggregationsVersion"] = o.AggregationsVersion
	}
	if !isNil(o.AggregationsDsl) {
		toSerialize["aggregationsDsl"] = o.AggregationsDsl
	}
	if !isNil(o.Aggregations) {
		toSerialize["aggregations"] = o.Aggregations
	}
	if !isNil(o.Sort) {
		toSerialize["sort"] = o.Sort
	}
	if !isNil(o.SearchAfter) {
		toSerialize["searchAfter"] = o.SearchAfter
	}
	if !isNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Search) UnmarshalJSON(bytes []byte) (err error) {
	varSearch := _Search{}

	if err = json.Unmarshal(bytes, &varSearch); err == nil {
		*o = Search(varSearch)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "indices")
		delete(additionalProperties, "queryType")
		delete(additionalProperties, "queryVersion")
		delete(additionalProperties, "query")
		delete(additionalProperties, "queryDsl")
		delete(additionalProperties, "textQuery")
		delete(additionalProperties, "typeAheadQuery")
		delete(additionalProperties, "includeNested")
		delete(additionalProperties, "queryResultFilter")
		delete(additionalProperties, "aggregationType")
		delete(additionalProperties, "aggregationsVersion")
		delete(additionalProperties, "aggregationsDsl")
		delete(additionalProperties, "aggregations")
		delete(additionalProperties, "sort")
		delete(additionalProperties, "searchAfter")
		delete(additionalProperties, "filters")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSearch struct {
	value *Search
	isSet bool
}

func (v NullableSearch) Get() *Search {
	return v.value
}

func (v *NullableSearch) Set(val *Search) {
	v.value = val
	v.isSet = true
}

func (v NullableSearch) IsSet() bool {
	return v.isSet
}

func (v *NullableSearch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearch(val *Search) *NullableSearch {
	return &NullableSearch{value: val, isSet: true}
}

func (v NullableSearch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


