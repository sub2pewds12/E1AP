package e1ap_ies

import "github.com/lvdund/ngap/aper"

// BearerContextSetupFailure From: 9_4_4_PDU_Definitions.txt:768
// ASN.1 Data Type: SEQUENCE
type BearerContextSetupFailure struct {
	GNBCUCPUEE1APID        aper.Integer            `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	GNBCUUPUEE1APID        *aper.Integer           `aper:"lb:0,ub:4294967295,optional,ignore,ext"`
	Cause                  Cause                   `aper:"mandatory,ignore,ext"`
	CriticalityDiagnostics *CriticalityDiagnostics `aper:"optional,ignore,ext"`
}
