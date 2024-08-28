/*
Identity Security Cloud V2024 API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: v2024
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2024

import (
	"encoding/json"
	"gopkg.in/validator.v2"
	"fmt"
)

// ReportDetailsArguments - The string-object map(dictionary) with the arguments needed for report processing.
type ReportDetailsArguments struct {
	AccountsExportReportArguments *AccountsExportReportArguments
	IdentitiesDetailsReportArguments *IdentitiesDetailsReportArguments
	IdentitiesReportArguments *IdentitiesReportArguments
	IdentityProfileIdentityErrorReportArguments *IdentityProfileIdentityErrorReportArguments
	OrphanUncorrelatedReportArguments *OrphanUncorrelatedReportArguments
	SearchExportReportArguments *SearchExportReportArguments
}

// AccountsExportReportArgumentsAsReportDetailsArguments is a convenience function that returns AccountsExportReportArguments wrapped in ReportDetailsArguments
func AccountsExportReportArgumentsAsReportDetailsArguments(v *AccountsExportReportArguments) ReportDetailsArguments {
	return ReportDetailsArguments{
		AccountsExportReportArguments: v,
	}
}

// IdentitiesDetailsReportArgumentsAsReportDetailsArguments is a convenience function that returns IdentitiesDetailsReportArguments wrapped in ReportDetailsArguments
func IdentitiesDetailsReportArgumentsAsReportDetailsArguments(v *IdentitiesDetailsReportArguments) ReportDetailsArguments {
	return ReportDetailsArguments{
		IdentitiesDetailsReportArguments: v,
	}
}

// IdentitiesReportArgumentsAsReportDetailsArguments is a convenience function that returns IdentitiesReportArguments wrapped in ReportDetailsArguments
func IdentitiesReportArgumentsAsReportDetailsArguments(v *IdentitiesReportArguments) ReportDetailsArguments {
	return ReportDetailsArguments{
		IdentitiesReportArguments: v,
	}
}

// IdentityProfileIdentityErrorReportArgumentsAsReportDetailsArguments is a convenience function that returns IdentityProfileIdentityErrorReportArguments wrapped in ReportDetailsArguments
func IdentityProfileIdentityErrorReportArgumentsAsReportDetailsArguments(v *IdentityProfileIdentityErrorReportArguments) ReportDetailsArguments {
	return ReportDetailsArguments{
		IdentityProfileIdentityErrorReportArguments: v,
	}
}

// OrphanUncorrelatedReportArgumentsAsReportDetailsArguments is a convenience function that returns OrphanUncorrelatedReportArguments wrapped in ReportDetailsArguments
func OrphanUncorrelatedReportArgumentsAsReportDetailsArguments(v *OrphanUncorrelatedReportArguments) ReportDetailsArguments {
	return ReportDetailsArguments{
		OrphanUncorrelatedReportArguments: v,
	}
}

// SearchExportReportArgumentsAsReportDetailsArguments is a convenience function that returns SearchExportReportArguments wrapped in ReportDetailsArguments
func SearchExportReportArgumentsAsReportDetailsArguments(v *SearchExportReportArguments) ReportDetailsArguments {
	return ReportDetailsArguments{
		SearchExportReportArguments: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ReportDetailsArguments) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AccountsExportReportArguments
	err = newStrictDecoder(data).Decode(&dst.AccountsExportReportArguments)
	if err == nil {
		jsonAccountsExportReportArguments, _ := json.Marshal(dst.AccountsExportReportArguments)
		if string(jsonAccountsExportReportArguments) == "{}" { // empty struct
			dst.AccountsExportReportArguments = nil
		} else {
			if err = validator.Validate(dst.AccountsExportReportArguments); err != nil {
				dst.AccountsExportReportArguments = nil
			} else {
				match++
			}
		}
	} else {
		dst.AccountsExportReportArguments = nil
	}

	// try to unmarshal data into IdentitiesDetailsReportArguments
	err = newStrictDecoder(data).Decode(&dst.IdentitiesDetailsReportArguments)
	if err == nil {
		jsonIdentitiesDetailsReportArguments, _ := json.Marshal(dst.IdentitiesDetailsReportArguments)
		if string(jsonIdentitiesDetailsReportArguments) == "{}" { // empty struct
			dst.IdentitiesDetailsReportArguments = nil
		} else {
			if err = validator.Validate(dst.IdentitiesDetailsReportArguments); err != nil {
				dst.IdentitiesDetailsReportArguments = nil
			} else {
				match++
			}
		}
	} else {
		dst.IdentitiesDetailsReportArguments = nil
	}

	// try to unmarshal data into IdentitiesReportArguments
	err = newStrictDecoder(data).Decode(&dst.IdentitiesReportArguments)
	if err == nil {
		jsonIdentitiesReportArguments, _ := json.Marshal(dst.IdentitiesReportArguments)
		if string(jsonIdentitiesReportArguments) == "{}" { // empty struct
			dst.IdentitiesReportArguments = nil
		} else {
			if err = validator.Validate(dst.IdentitiesReportArguments); err != nil {
				dst.IdentitiesReportArguments = nil
			} else {
				match++
			}
		}
	} else {
		dst.IdentitiesReportArguments = nil
	}

	// try to unmarshal data into IdentityProfileIdentityErrorReportArguments
	err = newStrictDecoder(data).Decode(&dst.IdentityProfileIdentityErrorReportArguments)
	if err == nil {
		jsonIdentityProfileIdentityErrorReportArguments, _ := json.Marshal(dst.IdentityProfileIdentityErrorReportArguments)
		if string(jsonIdentityProfileIdentityErrorReportArguments) == "{}" { // empty struct
			dst.IdentityProfileIdentityErrorReportArguments = nil
		} else {
			if err = validator.Validate(dst.IdentityProfileIdentityErrorReportArguments); err != nil {
				dst.IdentityProfileIdentityErrorReportArguments = nil
			} else {
				match++
			}
		}
	} else {
		dst.IdentityProfileIdentityErrorReportArguments = nil
	}

	// try to unmarshal data into OrphanUncorrelatedReportArguments
	err = newStrictDecoder(data).Decode(&dst.OrphanUncorrelatedReportArguments)
	if err == nil {
		jsonOrphanUncorrelatedReportArguments, _ := json.Marshal(dst.OrphanUncorrelatedReportArguments)
		if string(jsonOrphanUncorrelatedReportArguments) == "{}" { // empty struct
			dst.OrphanUncorrelatedReportArguments = nil
		} else {
			if err = validator.Validate(dst.OrphanUncorrelatedReportArguments); err != nil {
				dst.OrphanUncorrelatedReportArguments = nil
			} else {
				match++
			}
		}
	} else {
		dst.OrphanUncorrelatedReportArguments = nil
	}

	// try to unmarshal data into SearchExportReportArguments
	err = newStrictDecoder(data).Decode(&dst.SearchExportReportArguments)
	if err == nil {
		jsonSearchExportReportArguments, _ := json.Marshal(dst.SearchExportReportArguments)
		if string(jsonSearchExportReportArguments) == "{}" { // empty struct
			dst.SearchExportReportArguments = nil
		} else {
			if err = validator.Validate(dst.SearchExportReportArguments); err != nil {
				dst.SearchExportReportArguments = nil
			} else {
				match++
			}
		}
	} else {
		dst.SearchExportReportArguments = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AccountsExportReportArguments = nil
		dst.IdentitiesDetailsReportArguments = nil
		dst.IdentitiesReportArguments = nil
		dst.IdentityProfileIdentityErrorReportArguments = nil
		dst.OrphanUncorrelatedReportArguments = nil
		dst.SearchExportReportArguments = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ReportDetailsArguments)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ReportDetailsArguments)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ReportDetailsArguments) MarshalJSON() ([]byte, error) {
	if src.AccountsExportReportArguments != nil {
		return json.Marshal(&src.AccountsExportReportArguments)
	}

	if src.IdentitiesDetailsReportArguments != nil {
		return json.Marshal(&src.IdentitiesDetailsReportArguments)
	}

	if src.IdentitiesReportArguments != nil {
		return json.Marshal(&src.IdentitiesReportArguments)
	}

	if src.IdentityProfileIdentityErrorReportArguments != nil {
		return json.Marshal(&src.IdentityProfileIdentityErrorReportArguments)
	}

	if src.OrphanUncorrelatedReportArguments != nil {
		return json.Marshal(&src.OrphanUncorrelatedReportArguments)
	}

	if src.SearchExportReportArguments != nil {
		return json.Marshal(&src.SearchExportReportArguments)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ReportDetailsArguments) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AccountsExportReportArguments != nil {
		return obj.AccountsExportReportArguments
	}

	if obj.IdentitiesDetailsReportArguments != nil {
		return obj.IdentitiesDetailsReportArguments
	}

	if obj.IdentitiesReportArguments != nil {
		return obj.IdentitiesReportArguments
	}

	if obj.IdentityProfileIdentityErrorReportArguments != nil {
		return obj.IdentityProfileIdentityErrorReportArguments
	}

	if obj.OrphanUncorrelatedReportArguments != nil {
		return obj.OrphanUncorrelatedReportArguments
	}

	if obj.SearchExportReportArguments != nil {
		return obj.SearchExportReportArguments
	}

	// all schemas are nil
	return nil
}

type NullableReportDetailsArguments struct {
	value *ReportDetailsArguments
	isSet bool
}

func (v NullableReportDetailsArguments) Get() *ReportDetailsArguments {
	return v.value
}

func (v *NullableReportDetailsArguments) Set(val *ReportDetailsArguments) {
	v.value = val
	v.isSet = true
}

func (v NullableReportDetailsArguments) IsSet() bool {
	return v.isSet
}

func (v *NullableReportDetailsArguments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportDetailsArguments(val *ReportDetailsArguments) *NullableReportDetailsArguments {
	return &NullableReportDetailsArguments{value: val, isSet: true}
}

func (v NullableReportDetailsArguments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportDetailsArguments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


