package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DRBsSubjectToCounterCheckItemNGRAN From: 9_4_5_Information_Element_Definitions.txt:712
// ASN.1 Data Type: SEQUENCE
type DRBsSubjectToCounterCheckItemNGRAN struct {
	PDUSessionID aper.Integer `aper:"lb:0,ub:255,mandatory,ext"`
	DRBID        aper.Integer `aper:"mandatory,ext"`
	PDCPULCount  PDCPCount    `aper:"mandatory,ext"`
	PDCPDLCount  PDCPCount    `aper:"mandatory,ext"`
}
