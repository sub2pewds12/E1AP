package e1ap_ies

import "github.com/lvdund/ngap/aper"

// ROHC From: 9_4_5_Information_Element_Definitions.txt:2115
// ASN.1 Data Type: SEQUENCE
type ROHC struct {
	MaxCID       aper.Integer      `aper:"mandatory,ext"`
	ROHCProfiles aper.Integer      `aper:"mandatory,ext"`
	ContinueROHC *ROHCContinueROHC `aper:"optional,ext"`
}
