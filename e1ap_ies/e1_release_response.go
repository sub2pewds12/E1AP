package e1ap_ies

import "github.com/lvdund/ngap/aper"

// E1ReleaseResponse From: 9_4_4_PDU_Definitions.txt:650
// ASN.1 Data Type: SEQUENCE
type E1ReleaseResponse struct {
	TransactionID aper.Integer `aper:"mandatory,reject,ext"`
}
