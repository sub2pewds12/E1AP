package e1ap_ies

import "github.com/lvdund/ngap/aper"

// E1ReleaseRequest From: 9_4_4_PDU_Definitions.txt:633
// ASN.1 Data Type: SEQUENCE
type E1ReleaseRequest struct {
	TransactionID aper.Integer `aper:"mandatory,reject,ext"`
	Cause         Cause        `aper:"mandatory,ignore,ext"`
}
