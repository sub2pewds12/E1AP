package e1ap_ies

import "github.com/lvdund/ngap/aper"

// BearerContextSetupResponse From: 9_4_4_PDU_Definitions.txt:726
// ASN.1 Data Type: SEQUENCE
type BearerContextSetupResponse struct {
	GNBCUCPUEE1APID                  aper.Integer                     `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	GNBCUUPUEE1APID                  aper.Integer                     `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	SystemBearerContextSetupResponse SystemBearerContextSetupResponse `aper:"mandatory,ignore,ext"`
}
