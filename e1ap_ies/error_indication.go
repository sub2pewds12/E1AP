package e1ap_ies

import "github.com/lvdund/ngap/aper"

// ErrorIndication From: 9_4_4_PDU_Definitions.txt:319
// ASN.1 Data Type: SEQUENCE
type ErrorIndication struct {
	TransactionID          aper.Integer            `aper:"mandatory,reject,ext"`
	GNBCUCPUEE1APID        *aper.Integer           `aper:"lb:0,ub:4294967295,optional,ignore,ext"`
	GNBCUUPUEE1APID        *aper.Integer           `aper:"lb:0,ub:4294967295,optional,ignore,ext"`
	Cause                  *Cause                  `aper:"optional,ignore,ext"`
	CriticalityDiagnostics *CriticalityDiagnostics `aper:"optional,ignore,ext"`
}
