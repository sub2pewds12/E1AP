package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DRBStatusItem From: 9_4_5_Information_Element_Definitions.txt:684
// ASN.1 Data Type: SEQUENCE
type DRBStatusItem struct {
	DRBID       aper.Integer `aper:"mandatory,ext"`
	PDCPDLCount *PDCPCount   `aper:"optional,ext"`
	PDCPULCount *PDCPCount   `aper:"optional,ext"`
}
