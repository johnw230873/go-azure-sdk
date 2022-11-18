package dataconnectors

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TIDataConnectorProperties struct {
	DataTypes         *TIDataConnectorDataTypes `json:"dataTypes,omitempty"`
	TenantId          *string                   `json:"tenantId,omitempty"`
	TipLookbackPeriod *string                   `json:"tipLookbackPeriod,omitempty"`
}

func (o *TIDataConnectorProperties) GetTipLookbackPeriodAsTime() (*time.Time, error) {
	if o.TipLookbackPeriod == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.TipLookbackPeriod, "2006-01-02T15:04:05Z07:00")
}

func (o *TIDataConnectorProperties) SetTipLookbackPeriodAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.TipLookbackPeriod = &formatted
}
