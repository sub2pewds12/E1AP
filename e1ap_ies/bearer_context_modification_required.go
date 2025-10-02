package e1ap_ies

// BearerContextModificationRequired From: 9_4_4_PDU_Definitions.txt:921
type BearerContextModificationRequired struct {
	GNBCUCPUEE1APID                         int64                                   `asn1:"lb:0,ub:4294967295,mandatory,reject,ext"`
	GNBCUUPUEE1APID                         int64                                   `asn1:"lb:0,ub:4294967295,mandatory,reject,ext"`
	SystemBearerContextModificationRequired SystemBearerContextModificationRequired `asn1:"mandatory,reject,ext"`
}
