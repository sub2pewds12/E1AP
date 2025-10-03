package e1ap_ies

import "github.com/lvdund/ngap/aper"

// BearerContextModificationRequired From: 9_4_4_PDU_Definitions.txt:921
// ASN.1 Data Type: SEQUENCE
type BearerContextModificationRequired struct {
	GNBCUCPUEE1APID                         aper.Integer                            `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	GNBCUUPUEE1APID                         aper.Integer                            `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	SystemBearerContextModificationRequired SystemBearerContextModificationRequired `aper:"mandatory,reject,ext"`
}
