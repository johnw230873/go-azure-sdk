package provisionedclusterinstances

import (
	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AgentPool struct {
	ExtendedLocation *ExtendedLocation      `json:"extendedLocation,omitempty"`
	Id               *string                `json:"id,omitempty"`
	Name             *string                `json:"name,omitempty"`
	Properties       *AgentPoolProperties   `json:"properties,omitempty"`
	SystemData       *systemdata.SystemData `json:"systemData,omitempty"`
	Tags             *map[string]string     `json:"tags,omitempty"`
	Type             *string                `json:"type,omitempty"`
}
