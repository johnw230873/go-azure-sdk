package services

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ServicePlacementPolicy = ServicePlacementNonPartiallyPlaceServicePolicy{}

type ServicePlacementNonPartiallyPlaceServicePolicy struct {

	// Fields inherited from ServicePlacementPolicy
}

var _ json.Marshaler = ServicePlacementNonPartiallyPlaceServicePolicy{}

func (s ServicePlacementNonPartiallyPlaceServicePolicy) MarshalJSON() ([]byte, error) {
	type wrapper ServicePlacementNonPartiallyPlaceServicePolicy
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ServicePlacementNonPartiallyPlaceServicePolicy: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ServicePlacementNonPartiallyPlaceServicePolicy: %+v", err)
	}
	decoded["type"] = "NonPartiallyPlaceService"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ServicePlacementNonPartiallyPlaceServicePolicy: %+v", err)
	}

	return encoded, nil
}
