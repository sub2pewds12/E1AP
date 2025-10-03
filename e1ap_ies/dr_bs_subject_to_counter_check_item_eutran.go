package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DRBsSubjectToCounterCheckItemEUTRAN From: 9_4_5_Information_Element_Definitions.txt:698
// ASN.1 Data Type: SEQUENCE
type DRBsSubjectToCounterCheckItemEUTRAN struct {
	DRBID       aper.Integer `aper:"mandatory,ext"`
	PDCPULCount PDCPCount    `aper:"mandatory,ext"`
	PDCPDLCount PDCPCount    `aper:"mandatory,ext"`
}
