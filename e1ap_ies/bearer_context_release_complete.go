package e1ap_ies

// BearerContextReleaseComplete From: 9_4_4_PDU_Definitions.txt:1030
type BearerContextReleaseComplete struct {
	GNBCUCPUEE1APID               int64                   `asn1:"lb:0,ub:4294967295,mandatory,reject,ext"`
	GNBCUUPUEE1APID               int64                   `asn1:"lb:0,ub:4294967295,mandatory,reject,ext"`
	CriticalityDiagnostics        *CriticalityDiagnostics `asn1:"optional,ignore,ext"`
	RetainabilityMeasurementsInfo []DRBRemovedItem        `asn1:"optional,ignore,ext"`
}
