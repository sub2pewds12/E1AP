package e1ap_ies

import "github.com/lvdund/ngap/aper"

// Reset From: 9_4_4_PDU_Definitions.txt:249
// ASN.1 Data Type: SEQUENCE
type Reset struct {
	TransactionID aper.Integer `aper:"mandatory,reject,ext"`
	Cause         Cause        `aper:"mandatory,ignore,ext"`
	ResetType     ResetType    `aper:"mandatory,reject,ext"`
}
