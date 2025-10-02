package e1ap_ies

// ErrorIndication From: 9_4_4_PDU_Definitions.txt:319
type ErrorIndication struct {
	TransactionID          int64                   `asn1:"mandatory,reject,ext"`
	GNBCUCPUEE1APID        *int64                  `asn1:"lb:0,ub:4294967295,optional,ignore,ext"`
	GNBCUUPUEE1APID        *int64                  `asn1:"lb:0,ub:4294967295,optional,ignore,ext"`
	Cause                  *Cause                  `asn1:"optional,ignore,ext"`
	CriticalityDiagnostics *CriticalityDiagnostics `asn1:"optional,ignore,ext"`
}
