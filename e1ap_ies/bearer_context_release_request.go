package e1ap_ies

import "github.com/lvdund/ngap/aper"

// BearerContextReleaseRequest From: 9_4_4_PDU_Definitions.txt:1056
// ASN.1 Data Type: SEQUENCE
type BearerContextReleaseRequest struct {
	GNBCUCPUEE1APID aper.Integer    `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	GNBCUUPUEE1APID aper.Integer    `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	DRBStatusList   []DRBStatusItem `aper:"lb:1,ub:MaxnoofDRBs,optional,ignore,ext"`
	Cause           Cause           `aper:"mandatory,ignore,ext"`
}
