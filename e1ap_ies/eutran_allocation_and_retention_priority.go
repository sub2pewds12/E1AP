package e1ap_ies

import "github.com/lvdund/ngap/aper"

// EUTRANAllocationAndRetentionPriority From: 9_4_5_Information_Element_Definitions.txt:1027
// ASN.1 Data Type: SEQUENCE
type EUTRANAllocationAndRetentionPriority struct {
	PriorityLevel           aper.Integer            `aper:"mandatory,ext"`
	PreEmptionCapability    PreEmptionCapability    `aper:"mandatory,ext"`
	PreEmptionVulnerability PreEmptionVulnerability `aper:"mandatory,ext"`
}
