package e1ap_ies

import "github.com/lvdund/ngap/aper"

// IABUPTNLAddressUpdateAcknowledge From: 9_4_4_PDU_Definitions.txt:1449
// ASN.1 Data Type: SEQUENCE
type IABUPTNLAddressUpdateAcknowledge struct {
	TransactionID              aper.Integer                 `aper:"mandatory,reject,ext"`
	CriticalityDiagnostics     *CriticalityDiagnostics      `aper:"optional,ignore,ext"`
	ULUPTNLAddressToUpdateList []ULUPTNLAddressToUpdateItem `aper:"optional,ignore,ext"`
}
