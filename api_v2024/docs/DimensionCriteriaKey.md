# DimensionCriteriaKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**DimensionCriteriaKeyType**](DimensionCriteriaKeyType.md) |  | 
**Property** | **string** | The name of the identity attribute to which the associated criteria applies. | 

## Methods

### NewDimensionCriteriaKey

`func NewDimensionCriteriaKey(type_ DimensionCriteriaKeyType, property string, ) *DimensionCriteriaKey`

NewDimensionCriteriaKey instantiates a new DimensionCriteriaKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDimensionCriteriaKeyWithDefaults

`func NewDimensionCriteriaKeyWithDefaults() *DimensionCriteriaKey`

NewDimensionCriteriaKeyWithDefaults instantiates a new DimensionCriteriaKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DimensionCriteriaKey) GetType() DimensionCriteriaKeyType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DimensionCriteriaKey) GetTypeOk() (*DimensionCriteriaKeyType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DimensionCriteriaKey) SetType(v DimensionCriteriaKeyType)`

SetType sets Type field to given value.


### GetProperty

`func (o *DimensionCriteriaKey) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *DimensionCriteriaKey) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *DimensionCriteriaKey) SetProperty(v string)`

SetProperty sets Property field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

