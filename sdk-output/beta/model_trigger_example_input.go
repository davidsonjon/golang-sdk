/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package beta

import (
	"encoding/json"
	"fmt"
)

// TriggerExampleInput - An example of the JSON payload that will be sent by the trigger to the subscribed service.
type TriggerExampleInput struct {
	AccessRequestDynamicApprover *AccessRequestDynamicApprover
	AccessRequestPostApproval *AccessRequestPostApproval
	AccessRequestPreApproval *AccessRequestPreApproval
	AccountAggregationCompleted *AccountAggregationCompleted
	AccountAttributesChanged *AccountAttributesChanged
	AccountCorrelated *AccountCorrelated
	AccountUncorrelated *AccountUncorrelated
	AccountsCollectedForAggregation *AccountsCollectedForAggregation
	CampaignActivated *CampaignActivated
	CampaignEnded *CampaignEnded
	CampaignGenerated *CampaignGenerated
	CertificationSignedOff *CertificationSignedOff
	IdentityAttributesChanged *IdentityAttributesChanged
	IdentityCreated *IdentityCreated
	IdentityDeleted *IdentityDeleted
	ProvisioningCompleted *ProvisioningCompleted
	SavedSearchComplete *SavedSearchComplete
	SourceAccount *SourceAccount
	SourceCreated *SourceCreated
	SourceDeleted *SourceDeleted
	SourceUpdated *SourceUpdated
	VAClusterStatusChangeEvent *VAClusterStatusChangeEvent
}

// AccessRequestDynamicApproverAsTriggerExampleInput is a convenience function that returns AccessRequestDynamicApprover wrapped in TriggerExampleInput
func AccessRequestDynamicApproverAsTriggerExampleInput(v *AccessRequestDynamicApprover) TriggerExampleInput {
	return TriggerExampleInput{
		AccessRequestDynamicApprover: v,
	}
}

// AccessRequestPostApprovalAsTriggerExampleInput is a convenience function that returns AccessRequestPostApproval wrapped in TriggerExampleInput
func AccessRequestPostApprovalAsTriggerExampleInput(v *AccessRequestPostApproval) TriggerExampleInput {
	return TriggerExampleInput{
		AccessRequestPostApproval: v,
	}
}

// AccessRequestPreApprovalAsTriggerExampleInput is a convenience function that returns AccessRequestPreApproval wrapped in TriggerExampleInput
func AccessRequestPreApprovalAsTriggerExampleInput(v *AccessRequestPreApproval) TriggerExampleInput {
	return TriggerExampleInput{
		AccessRequestPreApproval: v,
	}
}

// AccountAggregationCompletedAsTriggerExampleInput is a convenience function that returns AccountAggregationCompleted wrapped in TriggerExampleInput
func AccountAggregationCompletedAsTriggerExampleInput(v *AccountAggregationCompleted) TriggerExampleInput {
	return TriggerExampleInput{
		AccountAggregationCompleted: v,
	}
}

// AccountAttributesChangedAsTriggerExampleInput is a convenience function that returns AccountAttributesChanged wrapped in TriggerExampleInput
func AccountAttributesChangedAsTriggerExampleInput(v *AccountAttributesChanged) TriggerExampleInput {
	return TriggerExampleInput{
		AccountAttributesChanged: v,
	}
}

// AccountCorrelatedAsTriggerExampleInput is a convenience function that returns AccountCorrelated wrapped in TriggerExampleInput
func AccountCorrelatedAsTriggerExampleInput(v *AccountCorrelated) TriggerExampleInput {
	return TriggerExampleInput{
		AccountCorrelated: v,
	}
}

// AccountUncorrelatedAsTriggerExampleInput is a convenience function that returns AccountUncorrelated wrapped in TriggerExampleInput
func AccountUncorrelatedAsTriggerExampleInput(v *AccountUncorrelated) TriggerExampleInput {
	return TriggerExampleInput{
		AccountUncorrelated: v,
	}
}

// AccountsCollectedForAggregationAsTriggerExampleInput is a convenience function that returns AccountsCollectedForAggregation wrapped in TriggerExampleInput
func AccountsCollectedForAggregationAsTriggerExampleInput(v *AccountsCollectedForAggregation) TriggerExampleInput {
	return TriggerExampleInput{
		AccountsCollectedForAggregation: v,
	}
}

// CampaignActivatedAsTriggerExampleInput is a convenience function that returns CampaignActivated wrapped in TriggerExampleInput
func CampaignActivatedAsTriggerExampleInput(v *CampaignActivated) TriggerExampleInput {
	return TriggerExampleInput{
		CampaignActivated: v,
	}
}

// CampaignEndedAsTriggerExampleInput is a convenience function that returns CampaignEnded wrapped in TriggerExampleInput
func CampaignEndedAsTriggerExampleInput(v *CampaignEnded) TriggerExampleInput {
	return TriggerExampleInput{
		CampaignEnded: v,
	}
}

// CampaignGeneratedAsTriggerExampleInput is a convenience function that returns CampaignGenerated wrapped in TriggerExampleInput
func CampaignGeneratedAsTriggerExampleInput(v *CampaignGenerated) TriggerExampleInput {
	return TriggerExampleInput{
		CampaignGenerated: v,
	}
}

// CertificationSignedOffAsTriggerExampleInput is a convenience function that returns CertificationSignedOff wrapped in TriggerExampleInput
func CertificationSignedOffAsTriggerExampleInput(v *CertificationSignedOff) TriggerExampleInput {
	return TriggerExampleInput{
		CertificationSignedOff: v,
	}
}

// IdentityAttributesChangedAsTriggerExampleInput is a convenience function that returns IdentityAttributesChanged wrapped in TriggerExampleInput
func IdentityAttributesChangedAsTriggerExampleInput(v *IdentityAttributesChanged) TriggerExampleInput {
	return TriggerExampleInput{
		IdentityAttributesChanged: v,
	}
}

// IdentityCreatedAsTriggerExampleInput is a convenience function that returns IdentityCreated wrapped in TriggerExampleInput
func IdentityCreatedAsTriggerExampleInput(v *IdentityCreated) TriggerExampleInput {
	return TriggerExampleInput{
		IdentityCreated: v,
	}
}

// IdentityDeletedAsTriggerExampleInput is a convenience function that returns IdentityDeleted wrapped in TriggerExampleInput
func IdentityDeletedAsTriggerExampleInput(v *IdentityDeleted) TriggerExampleInput {
	return TriggerExampleInput{
		IdentityDeleted: v,
	}
}

// ProvisioningCompletedAsTriggerExampleInput is a convenience function that returns ProvisioningCompleted wrapped in TriggerExampleInput
func ProvisioningCompletedAsTriggerExampleInput(v *ProvisioningCompleted) TriggerExampleInput {
	return TriggerExampleInput{
		ProvisioningCompleted: v,
	}
}

// SavedSearchCompleteAsTriggerExampleInput is a convenience function that returns SavedSearchComplete wrapped in TriggerExampleInput
func SavedSearchCompleteAsTriggerExampleInput(v *SavedSearchComplete) TriggerExampleInput {
	return TriggerExampleInput{
		SavedSearchComplete: v,
	}
}

// SourceAccountAsTriggerExampleInput is a convenience function that returns SourceAccount wrapped in TriggerExampleInput
func SourceAccountAsTriggerExampleInput(v *SourceAccount) TriggerExampleInput {
	return TriggerExampleInput{
		SourceAccount: v,
	}
}

// SourceCreatedAsTriggerExampleInput is a convenience function that returns SourceCreated wrapped in TriggerExampleInput
func SourceCreatedAsTriggerExampleInput(v *SourceCreated) TriggerExampleInput {
	return TriggerExampleInput{
		SourceCreated: v,
	}
}

// SourceDeletedAsTriggerExampleInput is a convenience function that returns SourceDeleted wrapped in TriggerExampleInput
func SourceDeletedAsTriggerExampleInput(v *SourceDeleted) TriggerExampleInput {
	return TriggerExampleInput{
		SourceDeleted: v,
	}
}

// SourceUpdatedAsTriggerExampleInput is a convenience function that returns SourceUpdated wrapped in TriggerExampleInput
func SourceUpdatedAsTriggerExampleInput(v *SourceUpdated) TriggerExampleInput {
	return TriggerExampleInput{
		SourceUpdated: v,
	}
}

// VAClusterStatusChangeEventAsTriggerExampleInput is a convenience function that returns VAClusterStatusChangeEvent wrapped in TriggerExampleInput
func VAClusterStatusChangeEventAsTriggerExampleInput(v *VAClusterStatusChangeEvent) TriggerExampleInput {
	return TriggerExampleInput{
		VAClusterStatusChangeEvent: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *TriggerExampleInput) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AccessRequestDynamicApprover
	err = newStrictDecoder(data).Decode(&dst.AccessRequestDynamicApprover)
	if err == nil {
		jsonAccessRequestDynamicApprover, _ := json.Marshal(dst.AccessRequestDynamicApprover)
		if string(jsonAccessRequestDynamicApprover) == "{}" { // empty struct
			dst.AccessRequestDynamicApprover = nil
		} else {
			match++
		}
	} else {
		dst.AccessRequestDynamicApprover = nil
	}

	// try to unmarshal data into AccessRequestPostApproval
	err = newStrictDecoder(data).Decode(&dst.AccessRequestPostApproval)
	if err == nil {
		jsonAccessRequestPostApproval, _ := json.Marshal(dst.AccessRequestPostApproval)
		if string(jsonAccessRequestPostApproval) == "{}" { // empty struct
			dst.AccessRequestPostApproval = nil
		} else {
			match++
		}
	} else {
		dst.AccessRequestPostApproval = nil
	}

	// try to unmarshal data into AccessRequestPreApproval
	err = newStrictDecoder(data).Decode(&dst.AccessRequestPreApproval)
	if err == nil {
		jsonAccessRequestPreApproval, _ := json.Marshal(dst.AccessRequestPreApproval)
		if string(jsonAccessRequestPreApproval) == "{}" { // empty struct
			dst.AccessRequestPreApproval = nil
		} else {
			match++
		}
	} else {
		dst.AccessRequestPreApproval = nil
	}

	// try to unmarshal data into AccountAggregationCompleted
	err = newStrictDecoder(data).Decode(&dst.AccountAggregationCompleted)
	if err == nil {
		jsonAccountAggregationCompleted, _ := json.Marshal(dst.AccountAggregationCompleted)
		if string(jsonAccountAggregationCompleted) == "{}" { // empty struct
			dst.AccountAggregationCompleted = nil
		} else {
			match++
		}
	} else {
		dst.AccountAggregationCompleted = nil
	}

	// try to unmarshal data into AccountAttributesChanged
	err = newStrictDecoder(data).Decode(&dst.AccountAttributesChanged)
	if err == nil {
		jsonAccountAttributesChanged, _ := json.Marshal(dst.AccountAttributesChanged)
		if string(jsonAccountAttributesChanged) == "{}" { // empty struct
			dst.AccountAttributesChanged = nil
		} else {
			match++
		}
	} else {
		dst.AccountAttributesChanged = nil
	}

	// try to unmarshal data into AccountCorrelated
	err = newStrictDecoder(data).Decode(&dst.AccountCorrelated)
	if err == nil {
		jsonAccountCorrelated, _ := json.Marshal(dst.AccountCorrelated)
		if string(jsonAccountCorrelated) == "{}" { // empty struct
			dst.AccountCorrelated = nil
		} else {
			match++
		}
	} else {
		dst.AccountCorrelated = nil
	}

	// try to unmarshal data into AccountUncorrelated
	err = newStrictDecoder(data).Decode(&dst.AccountUncorrelated)
	if err == nil {
		jsonAccountUncorrelated, _ := json.Marshal(dst.AccountUncorrelated)
		if string(jsonAccountUncorrelated) == "{}" { // empty struct
			dst.AccountUncorrelated = nil
		} else {
			match++
		}
	} else {
		dst.AccountUncorrelated = nil
	}

	// try to unmarshal data into AccountsCollectedForAggregation
	err = newStrictDecoder(data).Decode(&dst.AccountsCollectedForAggregation)
	if err == nil {
		jsonAccountsCollectedForAggregation, _ := json.Marshal(dst.AccountsCollectedForAggregation)
		if string(jsonAccountsCollectedForAggregation) == "{}" { // empty struct
			dst.AccountsCollectedForAggregation = nil
		} else {
			match++
		}
	} else {
		dst.AccountsCollectedForAggregation = nil
	}

	// try to unmarshal data into CampaignActivated
	err = newStrictDecoder(data).Decode(&dst.CampaignActivated)
	if err == nil {
		jsonCampaignActivated, _ := json.Marshal(dst.CampaignActivated)
		if string(jsonCampaignActivated) == "{}" { // empty struct
			dst.CampaignActivated = nil
		} else {
			match++
		}
	} else {
		dst.CampaignActivated = nil
	}

	// try to unmarshal data into CampaignEnded
	err = newStrictDecoder(data).Decode(&dst.CampaignEnded)
	if err == nil {
		jsonCampaignEnded, _ := json.Marshal(dst.CampaignEnded)
		if string(jsonCampaignEnded) == "{}" { // empty struct
			dst.CampaignEnded = nil
		} else {
			match++
		}
	} else {
		dst.CampaignEnded = nil
	}

	// try to unmarshal data into CampaignGenerated
	err = newStrictDecoder(data).Decode(&dst.CampaignGenerated)
	if err == nil {
		jsonCampaignGenerated, _ := json.Marshal(dst.CampaignGenerated)
		if string(jsonCampaignGenerated) == "{}" { // empty struct
			dst.CampaignGenerated = nil
		} else {
			match++
		}
	} else {
		dst.CampaignGenerated = nil
	}

	// try to unmarshal data into CertificationSignedOff
	err = newStrictDecoder(data).Decode(&dst.CertificationSignedOff)
	if err == nil {
		jsonCertificationSignedOff, _ := json.Marshal(dst.CertificationSignedOff)
		if string(jsonCertificationSignedOff) == "{}" { // empty struct
			dst.CertificationSignedOff = nil
		} else {
			match++
		}
	} else {
		dst.CertificationSignedOff = nil
	}

	// try to unmarshal data into IdentityAttributesChanged
	err = newStrictDecoder(data).Decode(&dst.IdentityAttributesChanged)
	if err == nil {
		jsonIdentityAttributesChanged, _ := json.Marshal(dst.IdentityAttributesChanged)
		if string(jsonIdentityAttributesChanged) == "{}" { // empty struct
			dst.IdentityAttributesChanged = nil
		} else {
			match++
		}
	} else {
		dst.IdentityAttributesChanged = nil
	}

	// try to unmarshal data into IdentityCreated
	err = newStrictDecoder(data).Decode(&dst.IdentityCreated)
	if err == nil {
		jsonIdentityCreated, _ := json.Marshal(dst.IdentityCreated)
		if string(jsonIdentityCreated) == "{}" { // empty struct
			dst.IdentityCreated = nil
		} else {
			match++
		}
	} else {
		dst.IdentityCreated = nil
	}

	// try to unmarshal data into IdentityDeleted
	err = newStrictDecoder(data).Decode(&dst.IdentityDeleted)
	if err == nil {
		jsonIdentityDeleted, _ := json.Marshal(dst.IdentityDeleted)
		if string(jsonIdentityDeleted) == "{}" { // empty struct
			dst.IdentityDeleted = nil
		} else {
			match++
		}
	} else {
		dst.IdentityDeleted = nil
	}

	// try to unmarshal data into ProvisioningCompleted
	err = newStrictDecoder(data).Decode(&dst.ProvisioningCompleted)
	if err == nil {
		jsonProvisioningCompleted, _ := json.Marshal(dst.ProvisioningCompleted)
		if string(jsonProvisioningCompleted) == "{}" { // empty struct
			dst.ProvisioningCompleted = nil
		} else {
			match++
		}
	} else {
		dst.ProvisioningCompleted = nil
	}

	// try to unmarshal data into SavedSearchComplete
	err = newStrictDecoder(data).Decode(&dst.SavedSearchComplete)
	if err == nil {
		jsonSavedSearchComplete, _ := json.Marshal(dst.SavedSearchComplete)
		if string(jsonSavedSearchComplete) == "{}" { // empty struct
			dst.SavedSearchComplete = nil
		} else {
			match++
		}
	} else {
		dst.SavedSearchComplete = nil
	}

	// try to unmarshal data into SourceAccount
	err = newStrictDecoder(data).Decode(&dst.SourceAccount)
	if err == nil {
		jsonSourceAccount, _ := json.Marshal(dst.SourceAccount)
		if string(jsonSourceAccount) == "{}" { // empty struct
			dst.SourceAccount = nil
		} else {
			match++
		}
	} else {
		dst.SourceAccount = nil
	}

	// try to unmarshal data into SourceCreated
	err = newStrictDecoder(data).Decode(&dst.SourceCreated)
	if err == nil {
		jsonSourceCreated, _ := json.Marshal(dst.SourceCreated)
		if string(jsonSourceCreated) == "{}" { // empty struct
			dst.SourceCreated = nil
		} else {
			match++
		}
	} else {
		dst.SourceCreated = nil
	}

	// try to unmarshal data into SourceDeleted
	err = newStrictDecoder(data).Decode(&dst.SourceDeleted)
	if err == nil {
		jsonSourceDeleted, _ := json.Marshal(dst.SourceDeleted)
		if string(jsonSourceDeleted) == "{}" { // empty struct
			dst.SourceDeleted = nil
		} else {
			match++
		}
	} else {
		dst.SourceDeleted = nil
	}

	// try to unmarshal data into SourceUpdated
	err = newStrictDecoder(data).Decode(&dst.SourceUpdated)
	if err == nil {
		jsonSourceUpdated, _ := json.Marshal(dst.SourceUpdated)
		if string(jsonSourceUpdated) == "{}" { // empty struct
			dst.SourceUpdated = nil
		} else {
			match++
		}
	} else {
		dst.SourceUpdated = nil
	}

	// try to unmarshal data into VAClusterStatusChangeEvent
	err = newStrictDecoder(data).Decode(&dst.VAClusterStatusChangeEvent)
	if err == nil {
		jsonVAClusterStatusChangeEvent, _ := json.Marshal(dst.VAClusterStatusChangeEvent)
		if string(jsonVAClusterStatusChangeEvent) == "{}" { // empty struct
			dst.VAClusterStatusChangeEvent = nil
		} else {
			match++
		}
	} else {
		dst.VAClusterStatusChangeEvent = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AccessRequestDynamicApprover = nil
		dst.AccessRequestPostApproval = nil
		dst.AccessRequestPreApproval = nil
		dst.AccountAggregationCompleted = nil
		dst.AccountAttributesChanged = nil
		dst.AccountCorrelated = nil
		dst.AccountUncorrelated = nil
		dst.AccountsCollectedForAggregation = nil
		dst.CampaignActivated = nil
		dst.CampaignEnded = nil
		dst.CampaignGenerated = nil
		dst.CertificationSignedOff = nil
		dst.IdentityAttributesChanged = nil
		dst.IdentityCreated = nil
		dst.IdentityDeleted = nil
		dst.ProvisioningCompleted = nil
		dst.SavedSearchComplete = nil
		dst.SourceAccount = nil
		dst.SourceCreated = nil
		dst.SourceDeleted = nil
		dst.SourceUpdated = nil
		dst.VAClusterStatusChangeEvent = nil

		return fmt.Errorf("data matches more than one schema in oneOf(TriggerExampleInput)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(TriggerExampleInput)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TriggerExampleInput) MarshalJSON() ([]byte, error) {
	if src.AccessRequestDynamicApprover != nil {
		return json.Marshal(&src.AccessRequestDynamicApprover)
	}

	if src.AccessRequestPostApproval != nil {
		return json.Marshal(&src.AccessRequestPostApproval)
	}

	if src.AccessRequestPreApproval != nil {
		return json.Marshal(&src.AccessRequestPreApproval)
	}

	if src.AccountAggregationCompleted != nil {
		return json.Marshal(&src.AccountAggregationCompleted)
	}

	if src.AccountAttributesChanged != nil {
		return json.Marshal(&src.AccountAttributesChanged)
	}

	if src.AccountCorrelated != nil {
		return json.Marshal(&src.AccountCorrelated)
	}

	if src.AccountUncorrelated != nil {
		return json.Marshal(&src.AccountUncorrelated)
	}

	if src.AccountsCollectedForAggregation != nil {
		return json.Marshal(&src.AccountsCollectedForAggregation)
	}

	if src.CampaignActivated != nil {
		return json.Marshal(&src.CampaignActivated)
	}

	if src.CampaignEnded != nil {
		return json.Marshal(&src.CampaignEnded)
	}

	if src.CampaignGenerated != nil {
		return json.Marshal(&src.CampaignGenerated)
	}

	if src.CertificationSignedOff != nil {
		return json.Marshal(&src.CertificationSignedOff)
	}

	if src.IdentityAttributesChanged != nil {
		return json.Marshal(&src.IdentityAttributesChanged)
	}

	if src.IdentityCreated != nil {
		return json.Marshal(&src.IdentityCreated)
	}

	if src.IdentityDeleted != nil {
		return json.Marshal(&src.IdentityDeleted)
	}

	if src.ProvisioningCompleted != nil {
		return json.Marshal(&src.ProvisioningCompleted)
	}

	if src.SavedSearchComplete != nil {
		return json.Marshal(&src.SavedSearchComplete)
	}

	if src.SourceAccount != nil {
		return json.Marshal(&src.SourceAccount)
	}

	if src.SourceCreated != nil {
		return json.Marshal(&src.SourceCreated)
	}

	if src.SourceDeleted != nil {
		return json.Marshal(&src.SourceDeleted)
	}

	if src.SourceUpdated != nil {
		return json.Marshal(&src.SourceUpdated)
	}

	if src.VAClusterStatusChangeEvent != nil {
		return json.Marshal(&src.VAClusterStatusChangeEvent)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TriggerExampleInput) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AccessRequestDynamicApprover != nil {
		return obj.AccessRequestDynamicApprover
	}

	if obj.AccessRequestPostApproval != nil {
		return obj.AccessRequestPostApproval
	}

	if obj.AccessRequestPreApproval != nil {
		return obj.AccessRequestPreApproval
	}

	if obj.AccountAggregationCompleted != nil {
		return obj.AccountAggregationCompleted
	}

	if obj.AccountAttributesChanged != nil {
		return obj.AccountAttributesChanged
	}

	if obj.AccountCorrelated != nil {
		return obj.AccountCorrelated
	}

	if obj.AccountUncorrelated != nil {
		return obj.AccountUncorrelated
	}

	if obj.AccountsCollectedForAggregation != nil {
		return obj.AccountsCollectedForAggregation
	}

	if obj.CampaignActivated != nil {
		return obj.CampaignActivated
	}

	if obj.CampaignEnded != nil {
		return obj.CampaignEnded
	}

	if obj.CampaignGenerated != nil {
		return obj.CampaignGenerated
	}

	if obj.CertificationSignedOff != nil {
		return obj.CertificationSignedOff
	}

	if obj.IdentityAttributesChanged != nil {
		return obj.IdentityAttributesChanged
	}

	if obj.IdentityCreated != nil {
		return obj.IdentityCreated
	}

	if obj.IdentityDeleted != nil {
		return obj.IdentityDeleted
	}

	if obj.ProvisioningCompleted != nil {
		return obj.ProvisioningCompleted
	}

	if obj.SavedSearchComplete != nil {
		return obj.SavedSearchComplete
	}

	if obj.SourceAccount != nil {
		return obj.SourceAccount
	}

	if obj.SourceCreated != nil {
		return obj.SourceCreated
	}

	if obj.SourceDeleted != nil {
		return obj.SourceDeleted
	}

	if obj.SourceUpdated != nil {
		return obj.SourceUpdated
	}

	if obj.VAClusterStatusChangeEvent != nil {
		return obj.VAClusterStatusChangeEvent
	}

	// all schemas are nil
	return nil
}

type NullableTriggerExampleInput struct {
	value *TriggerExampleInput
	isSet bool
}

func (v NullableTriggerExampleInput) Get() *TriggerExampleInput {
	return v.value
}

func (v *NullableTriggerExampleInput) Set(val *TriggerExampleInput) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerExampleInput) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerExampleInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerExampleInput(val *TriggerExampleInput) *NullableTriggerExampleInput {
	return &NullableTriggerExampleInput{value: val, isSet: true}
}

func (v NullableTriggerExampleInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerExampleInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


