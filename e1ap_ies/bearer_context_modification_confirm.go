package e1ap_ies

import "github.com/lvdund/ngap/aper"

// BearerContextModificationConfirm From: 9_4_4_PDU_Definitions.txt:966
// ASN.1 Data Type: SEQUENCE
type BearerContextModificationConfirm struct {
	GNBCUCPUEE1APID                        aper.Integer                            `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	GNBCUUPUEE1APID                        aper.Integer                            `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	SystemBearerContextModificationConfirm *SystemBearerContextModificationConfirm `aper:"optional,ignore,ext"`
}
