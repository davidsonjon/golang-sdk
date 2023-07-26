# AccountUncorrelatedAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **map[string]interface{}** | The type of object that is referenced | 
**NativeIdentity** | **string** | Unique ID of the account on the source. | 
**Uuid** | Pointer to **NullableString** | The source&#39;s unique identifier for the account. UUID is generated by the source system. | [optional] 
**Id** | **string** | ID of the object to which this reference applies | 
**Name** | **string** | Human-readable display name of the object to which this reference applies | 

## Methods

### NewAccountUncorrelatedAccount

`func NewAccountUncorrelatedAccount(type_ map[string]interface{}, nativeIdentity string, id string, name string, ) *AccountUncorrelatedAccount`

NewAccountUncorrelatedAccount instantiates a new AccountUncorrelatedAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountUncorrelatedAccountWithDefaults

`func NewAccountUncorrelatedAccountWithDefaults() *AccountUncorrelatedAccount`

NewAccountUncorrelatedAccountWithDefaults instantiates a new AccountUncorrelatedAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AccountUncorrelatedAccount) GetType() map[string]interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccountUncorrelatedAccount) GetTypeOk() (*map[string]interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccountUncorrelatedAccount) SetType(v map[string]interface{})`

SetType sets Type field to given value.


### GetNativeIdentity

`func (o *AccountUncorrelatedAccount) GetNativeIdentity() string`

GetNativeIdentity returns the NativeIdentity field if non-nil, zero value otherwise.

### GetNativeIdentityOk

`func (o *AccountUncorrelatedAccount) GetNativeIdentityOk() (*string, bool)`

GetNativeIdentityOk returns a tuple with the NativeIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeIdentity

`func (o *AccountUncorrelatedAccount) SetNativeIdentity(v string)`

SetNativeIdentity sets NativeIdentity field to given value.


### GetUuid

`func (o *AccountUncorrelatedAccount) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AccountUncorrelatedAccount) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AccountUncorrelatedAccount) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *AccountUncorrelatedAccount) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *AccountUncorrelatedAccount) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *AccountUncorrelatedAccount) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetId

`func (o *AccountUncorrelatedAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountUncorrelatedAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountUncorrelatedAccount) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AccountUncorrelatedAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountUncorrelatedAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountUncorrelatedAccount) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

