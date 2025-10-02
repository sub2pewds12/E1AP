package e1ap_ies

// ResetAcknowledge From: 9_4_4_PDU_Definitions.txt:291
type ResetAcknowledge struct {
	TransactionID                             int64                       `asn1:"mandatory,reject,ext"`
	UEAssociatedLogicalE1ConnectionListResAck []ProtocolIESingleContainer `asn1:"optional,ignore,ext"`
	CriticalityDiagnostics                    *CriticalityDiagnostics     `asn1:"optional,ignore,ext"`
}
