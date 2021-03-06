// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

// NsSnapVol - Select fields containing volume info.
// Export NsSnapVolFields for advance operations like search filter etc.
var NsSnapVolFields *NsSnapVol

func init() {
	VolIdfield := "vol_id"
	SnapNamefield := "snap_name"
	SnapDescriptionfield := "snap_description"
	Cookiefield := "cookie"
	AppUuidfield := "app_uuid"

	NsSnapVolFields = &NsSnapVol{
		VolId:           &VolIdfield,
		SnapName:        &SnapNamefield,
		SnapDescription: &SnapDescriptionfield,
		Cookie:          &Cookiefield,
		AppUuid:         &AppUuidfield,
	}
}

type NsSnapVol struct {
	// VolId - ID of volume.
	VolId *string `json:"vol_id,omitempty"`
	// SnapName - Snapshot name.
	SnapName *string `json:"snap_name,omitempty"`
	// SnapDescription - Snapshot description.
	SnapDescription *string `json:"snap_description,omitempty"`
	// Cookie - A cookie.
	Cookie *string `json:"cookie,omitempty"`
	// Online - Snapshot is online.
	Online *bool `json:"online,omitempty"`
	// Writable - Snapshot is writable.
	Writable *bool `json:"writable,omitempty"`
	// AppUuid - Application identifier of snapshots.
	AppUuid *string `json:"app_uuid,omitempty"`
	// AgentType - External management agent type.
	AgentType *NsAgentType `json:"agent_type,omitempty"`
	// Metadata - Key-value pairs that augment a snapshot's attributes.
	Metadata []*NsKeyValue `json:"metadata,omitempty"`
}
