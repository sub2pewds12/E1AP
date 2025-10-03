package e1ap_ies

import "github.com/lvdund/ngap/aper"

// IABUPTNLAddressUpdate From: 9_4_4_PDU_Definitions.txt:1430
// ASN.1 Data Type: SEQUENCE
type IABUPTNLAddressUpdate struct {
	TransactionID              aper.Integer                 `aper:"mandatory,reject,ext"`
	DLUPTNLAddressToUpdateList []DLUPTNLAddressToUpdateItem `aper:"optional,ignore,ext"`
}
