package e1ap_ies

// BearerContextModificationFailure From: 9_4_4_PDU_Definitions.txt:896
type BearerContextModificationFailure struct {
	GNBCUCPUEE1APID        int64                   `asn1:"lb:0,ub:4294967295,mandatory,reject,ext"`
	GNBCUUPUEE1APID        int64                   `asn1:"lb:0,ub:4294967295,mandatory,reject,ext"`
	Cause                  Cause                   `asn1:"mandatory,ignore,ext"`
	CriticalityDiagnostics *CriticalityDiagnostics `asn1:"optional,ignore,ext"`
}
