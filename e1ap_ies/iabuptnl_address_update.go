package e1ap_ies

// IABUPTNLAddressUpdate From: 9_4_4_PDU_Definitions.txt:1430
type IABUPTNLAddressUpdate struct {
	TransactionID              int64                        `asn1:"mandatory,reject,ext"`
	DLUPTNLAddressToUpdateList []DLUPTNLAddressToUpdateItem `asn1:"optional,ignore,ext"`
}
