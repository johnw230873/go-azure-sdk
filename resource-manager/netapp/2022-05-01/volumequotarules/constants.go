package volumequotarules

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisioningState string

const (
	ProvisioningStateAccepted  ProvisioningState = "Accepted"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateMoving    ProvisioningState = "Moving"
	ProvisioningStatePatching  ProvisioningState = "Patching"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateAccepted),
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateMoving),
		string(ProvisioningStatePatching),
		string(ProvisioningStateSucceeded),
	}
}

func parseProvisioningState(input string) (*ProvisioningState, error) {
	vals := map[string]ProvisioningState{
		"accepted":  ProvisioningStateAccepted,
		"creating":  ProvisioningStateCreating,
		"deleting":  ProvisioningStateDeleting,
		"failed":    ProvisioningStateFailed,
		"moving":    ProvisioningStateMoving,
		"patching":  ProvisioningStatePatching,
		"succeeded": ProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningState(input)
	return &out, nil
}

type Type string

const (
	TypeDefaultGroupQuota    Type = "DefaultGroupQuota"
	TypeDefaultUserQuota     Type = "DefaultUserQuota"
	TypeIndividualGroupQuota Type = "IndividualGroupQuota"
	TypeIndividualUserQuota  Type = "IndividualUserQuota"
)

func PossibleValuesForType() []string {
	return []string{
		string(TypeDefaultGroupQuota),
		string(TypeDefaultUserQuota),
		string(TypeIndividualGroupQuota),
		string(TypeIndividualUserQuota),
	}
}

func parseType(input string) (*Type, error) {
	vals := map[string]Type{
		"defaultgroupquota":    TypeDefaultGroupQuota,
		"defaultuserquota":     TypeDefaultUserQuota,
		"individualgroupquota": TypeIndividualGroupQuota,
		"individualuserquota":  TypeIndividualUserQuota,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Type(input)
	return &out, nil
}
