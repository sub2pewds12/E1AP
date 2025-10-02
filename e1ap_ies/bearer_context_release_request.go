package e1ap_ies

// BearerContextReleaseRequest From: 9_4_4_PDU_Definitions.txt:1056
type BearerContextReleaseRequest struct {
	GNBCUCPUEE1APID int64           `asn1:"lb:0,ub:4294967295,mandatory,reject,ext"`
	GNBCUUPUEE1APID int64           `asn1:"lb:0,ub:4294967295,mandatory,reject,ext"`
	DRBStatusList   []DRBStatusItem `asn1:"lb:1,ub:MaxnoofDRBs,optional,ignore,ext"`
	Cause           Cause           `asn1:"mandatory,ignore,ext"`
}
