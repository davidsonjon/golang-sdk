/*
Identity Security Cloud Beta API

Use these APIs to interact with the Identity Security Cloud platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_beta

import (
	"encoding/json"
	"fmt"
)

// SourceFeature Optional features that can be supported by an source. * AUTHENTICATE: The source supports pass-through authentication. * COMPOSITE: The source supports composite source creation. * DIRECT_PERMISSIONS: The source supports returning DirectPermissions. * DISCOVER_SCHEMA: The source supports discovering schemas for users and groups. * ENABLE The source supports reading if an account is enabled or disabled. * MANAGER_LOOKUP: The source supports looking up managers as they are encountered in a feed. This is the opposite of NO_RANDOM_ACCESS. * NO_RANDOM_ACCESS: The source does not support random access and the getObject() methods should not be called and expected to perform. * PROXY: The source can serve as a proxy for another source. When an source has a proxy, all connector calls made with that source are redirected through the connector for the proxy source. * SEARCH * TEMPLATE * UNLOCK: The source supports reading if an account is locked or unlocked. * UNSTRUCTURED_TARGETS: The source supports returning unstructured Targets. * SHAREPOINT_TARGET: The source supports returning unstructured Target data for SharePoint. It will be typically used by AD, LDAP sources. * PROVISIONING: The source can both read and write accounts. Having this feature implies that the provision() method is implemented. It also means that direct and target permissions can also be provisioned if they can be returned by aggregation. * GROUP_PROVISIONING: The source can both read and write groups. Having this feature implies that the provision() method is implemented. * SYNC_PROVISIONING: The source can provision accounts synchronously. * PASSWORD: The source can provision password changes. Since sources can never read passwords, this is should only be used in conjunction with the PROVISIONING feature. * CURRENT_PASSWORD: Some source types support verification of the current password * ACCOUNT_ONLY_REQUEST: The source supports requesting accounts without entitlements. * ADDITIONAL_ACCOUNT_REQUEST: The source supports requesting additional accounts. * NO_AGGREGATION: A source that does not support aggregation. * GROUPS_HAVE_MEMBERS: The source models group memberships with a member attribute on the group object rather than a groups attribute on the account object. This effects the implementation of delta account aggregation. * NO_PERMISSIONS_PROVISIONING: Indicates that the connector cannot provision direct or target permissions for accounts. When DIRECT_PERMISSIONS and PROVISIONING features are present, it is assumed that the connector can also provision direct permissions. This feature disables that assumption and causes permission request to be converted to work items for accounts. * NO_GROUP_PERMISSIONS_PROVISIONING: Indicates that the connector cannot provision direct or target permissions for groups. When DIRECT_PERMISSIONS and PROVISIONING features are present, it is assumed that the connector can also provision direct permissions. This feature disables that assumption and causes permission request to be converted to work items for groups. * NO_UNSTRUCTURED_TARGETS_PROVISIONING: This string will be replaced by NO_GROUP_PERMISSIONS_PROVISIONING and NO_PERMISSIONS_PROVISIONING. * NO_DIRECT_PERMISSIONS_PROVISIONING: This string will be replaced by NO_GROUP_PERMISSIONS_PROVISIONING and NO_PERMISSIONS_PROVISIONING. * USES_UUID: Connectivity 2.0 flag used to indicate that the connector supports a compound naming structure. * PREFER_UUID: Used in ISC Provisioning AND Aggregation to decide if it should prefer account.uuid to account.nativeIdentity when data is read in through aggregation OR pushed out through provisioning. * ARM_SECURITY_EXTRACT: Indicates the application supports Security extracts for ARM * ARM_UTILIZATION_EXTRACT: Indicates the application supports Utilization extracts for ARM * ARM_CHANGELOG_EXTRACT: Indicates the application supports Change-log extracts for ARM
type SourceFeature string

// List of SourceFeature
const (
	SOURCEFEATURE_AUTHENTICATE SourceFeature = "AUTHENTICATE"
	SOURCEFEATURE_COMPOSITE SourceFeature = "COMPOSITE"
	SOURCEFEATURE_DIRECT_PERMISSIONS SourceFeature = "DIRECT_PERMISSIONS"
	SOURCEFEATURE_DISCOVER_SCHEMA SourceFeature = "DISCOVER_SCHEMA"
	SOURCEFEATURE_ENABLE SourceFeature = "ENABLE"
	SOURCEFEATURE_MANAGER_LOOKUP SourceFeature = "MANAGER_LOOKUP"
	SOURCEFEATURE_NO_RANDOM_ACCESS SourceFeature = "NO_RANDOM_ACCESS"
	SOURCEFEATURE_PROXY SourceFeature = "PROXY"
	SOURCEFEATURE_SEARCH SourceFeature = "SEARCH"
	SOURCEFEATURE_TEMPLATE SourceFeature = "TEMPLATE"
	SOURCEFEATURE_UNLOCK SourceFeature = "UNLOCK"
	SOURCEFEATURE_UNSTRUCTURED_TARGETS SourceFeature = "UNSTRUCTURED_TARGETS"
	SOURCEFEATURE_SHAREPOINT_TARGET SourceFeature = "SHAREPOINT_TARGET"
	SOURCEFEATURE_PROVISIONING SourceFeature = "PROVISIONING"
	SOURCEFEATURE_GROUP_PROVISIONING SourceFeature = "GROUP_PROVISIONING"
	SOURCEFEATURE_SYNC_PROVISIONING SourceFeature = "SYNC_PROVISIONING"
	SOURCEFEATURE_PASSWORD SourceFeature = "PASSWORD"
	SOURCEFEATURE_CURRENT_PASSWORD SourceFeature = "CURRENT_PASSWORD"
	SOURCEFEATURE_ACCOUNT_ONLY_REQUEST SourceFeature = "ACCOUNT_ONLY_REQUEST"
	SOURCEFEATURE_ADDITIONAL_ACCOUNT_REQUEST SourceFeature = "ADDITIONAL_ACCOUNT_REQUEST"
	SOURCEFEATURE_NO_AGGREGATION SourceFeature = "NO_AGGREGATION"
	SOURCEFEATURE_GROUPS_HAVE_MEMBERS SourceFeature = "GROUPS_HAVE_MEMBERS"
	SOURCEFEATURE_NO_PERMISSIONS_PROVISIONING SourceFeature = "NO_PERMISSIONS_PROVISIONING"
	SOURCEFEATURE_NO_GROUP_PERMISSIONS_PROVISIONING SourceFeature = "NO_GROUP_PERMISSIONS_PROVISIONING"
	SOURCEFEATURE_NO_UNSTRUCTURED_TARGETS_PROVISIONING SourceFeature = "NO_UNSTRUCTURED_TARGETS_PROVISIONING"
	SOURCEFEATURE_NO_DIRECT_PERMISSIONS_PROVISIONING SourceFeature = "NO_DIRECT_PERMISSIONS_PROVISIONING"
	SOURCEFEATURE_PREFER_UUID SourceFeature = "PREFER_UUID"
	SOURCEFEATURE_ARM_SECURITY_EXTRACT SourceFeature = "ARM_SECURITY_EXTRACT"
	SOURCEFEATURE_ARM_UTILIZATION_EXTRACT SourceFeature = "ARM_UTILIZATION_EXTRACT"
	SOURCEFEATURE_ARM_CHANGELOG_EXTRACT SourceFeature = "ARM_CHANGELOG_EXTRACT"
	SOURCEFEATURE_USES_UUID SourceFeature = "USES_UUID"
)

// All allowed values of SourceFeature enum
var AllowedSourceFeatureEnumValues = []SourceFeature{
	"AUTHENTICATE",
	"COMPOSITE",
	"DIRECT_PERMISSIONS",
	"DISCOVER_SCHEMA",
	"ENABLE",
	"MANAGER_LOOKUP",
	"NO_RANDOM_ACCESS",
	"PROXY",
	"SEARCH",
	"TEMPLATE",
	"UNLOCK",
	"UNSTRUCTURED_TARGETS",
	"SHAREPOINT_TARGET",
	"PROVISIONING",
	"GROUP_PROVISIONING",
	"SYNC_PROVISIONING",
	"PASSWORD",
	"CURRENT_PASSWORD",
	"ACCOUNT_ONLY_REQUEST",
	"ADDITIONAL_ACCOUNT_REQUEST",
	"NO_AGGREGATION",
	"GROUPS_HAVE_MEMBERS",
	"NO_PERMISSIONS_PROVISIONING",
	"NO_GROUP_PERMISSIONS_PROVISIONING",
	"NO_UNSTRUCTURED_TARGETS_PROVISIONING",
	"NO_DIRECT_PERMISSIONS_PROVISIONING",
	"PREFER_UUID",
	"ARM_SECURITY_EXTRACT",
	"ARM_UTILIZATION_EXTRACT",
	"ARM_CHANGELOG_EXTRACT",
	"USES_UUID",
}

func (v *SourceFeature) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SourceFeature(value)
	for _, existing := range AllowedSourceFeatureEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SourceFeature", value)
}

// NewSourceFeatureFromValue returns a pointer to a valid SourceFeature
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSourceFeatureFromValue(v string) (*SourceFeature, error) {
	ev := SourceFeature(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SourceFeature: valid values are %v", v, AllowedSourceFeatureEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SourceFeature) IsValid() bool {
	for _, existing := range AllowedSourceFeatureEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SourceFeature value
func (v SourceFeature) Ptr() *SourceFeature {
	return &v
}

type NullableSourceFeature struct {
	value *SourceFeature
	isSet bool
}

func (v NullableSourceFeature) Get() *SourceFeature {
	return v.value
}

func (v *NullableSourceFeature) Set(val *SourceFeature) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceFeature) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceFeature) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceFeature(val *SourceFeature) *NullableSourceFeature {
	return &NullableSourceFeature{value: val, isSet: true}
}

func (v NullableSourceFeature) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceFeature) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

