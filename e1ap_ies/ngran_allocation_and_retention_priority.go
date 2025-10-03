package e1ap_ies

import "github.com/lvdund/ngap/aper"

// NGRANAllocationAndRetentionPriority From: 9_4_5_Information_Element_Definitions.txt:1415
// ASN.1 Data Type: SEQUENCE
type NGRANAllocationAndRetentionPriority struct {
	PriorityLevel           aper.Integer            `aper:"mandatory"`
	PreEmptionCapability    PreEmptionCapability    `aper:"mandatory"`
	PreEmptionVulnerability PreEmptionVulnerability `aper:"mandatory"`
}
