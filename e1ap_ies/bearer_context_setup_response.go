package e1ap_ies

// BearerContextSetupResponse From: 9_4_4_PDU_Definitions.txt:726
type BearerContextSetupResponse struct {
	GNBCUCPUEE1APID                  int64                            `asn1:"lb:0,ub:4294967295,mandatory,reject,ext"`
	GNBCUUPUEE1APID                  int64                            `asn1:"lb:0,ub:4294967295,mandatory,reject,ext"`
	SystemBearerContextSetupResponse SystemBearerContextSetupResponse `asn1:"mandatory,ignore,ext"`
}
