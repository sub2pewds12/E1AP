package e1ap_ies

// IABUPTNLAddressUpdateAcknowledge From: 9_4_4_PDU_Definitions.txt:1449
type IABUPTNLAddressUpdateAcknowledge struct {
	TransactionID              int64                        `asn1:"mandatory,reject,ext"`
	CriticalityDiagnostics     *CriticalityDiagnostics      `asn1:"optional,ignore,ext"`
	ULUPTNLAddressToUpdateList []ULUPTNLAddressToUpdateItem `asn1:"optional,ignore,ext"`
}
