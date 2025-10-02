package e1ap_ies

// BearerContextModificationResponse From: 9_4_4_PDU_Definitions.txt:848
type BearerContextModificationResponse struct {
	GNBCUCPUEE1APID                         int64                                    `asn1:"lb:0,ub:4294967295,mandatory,reject,ext"`
	GNBCUUPUEE1APID                         int64                                    `asn1:"lb:0,ub:4294967295,mandatory,reject,ext"`
	SystemBearerContextModificationResponse *SystemBearerContextModificationResponse `asn1:"optional,ignore,ext"`
}
