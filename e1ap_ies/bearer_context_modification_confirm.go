package e1ap_ies

// BearerContextModificationConfirm From: 9_4_4_PDU_Definitions.txt:966
type BearerContextModificationConfirm struct {
	GNBCUCPUEE1APID                        int64                                   `asn1:"lb:0,ub:4294967295,mandatory,reject,ext"`
	GNBCUUPUEE1APID                        int64                                   `asn1:"lb:0,ub:4294967295,mandatory,reject,ext"`
	SystemBearerContextModificationConfirm *SystemBearerContextModificationConfirm `asn1:"optional,ignore,ext"`
}
