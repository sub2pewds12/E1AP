package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DRBsSubjectToEarlyForwardingItem From: 9_4_5_Information_Element_Definitions.txt:727
// ASN.1 Data Type: SEQUENCE
type DRBsSubjectToEarlyForwardingItem struct {
	DRBID        aper.Integer `aper:"mandatory,ext"`
	DLCountValue PDCPCount    `aper:"mandatory,ext"`
}
