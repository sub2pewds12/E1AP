package e1ap_ies

import "github.com/lvdund/ngap/aper"

// UplinkOnlyROHC From: 9_4_5_Information_Element_Definitions.txt:2430
// ASN.1 Data Type: SEQUENCE
type UplinkOnlyROHC struct {
	MaxCID       aper.Integer                `aper:"mandatory,ext"`
	ROHCProfiles aper.Integer                `aper:"mandatory,ext"`
	ContinueROHC *UplinkOnlyROHCContinueROHC `aper:"optional,ext"`
}
