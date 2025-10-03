package e1ap_ies

import "github.com/lvdund/ngap/aper"

// BearerContextReleaseComplete From: 9_4_4_PDU_Definitions.txt:1030
// ASN.1 Data Type: SEQUENCE
type BearerContextReleaseComplete struct {
	GNBCUCPUEE1APID               aper.Integer            `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	GNBCUUPUEE1APID               aper.Integer            `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	CriticalityDiagnostics        *CriticalityDiagnostics `aper:"optional,ignore,ext"`
	RetainabilityMeasurementsInfo []DRBRemovedItem        `aper:"optional,ignore,ext"`
}
