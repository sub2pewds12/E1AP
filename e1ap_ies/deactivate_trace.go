package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DeactivateTrace From: 9_4_4_PDU_Definitions.txt:1285
// ASN.1 Data Type: SEQUENCE
type DeactivateTrace struct {
	GNBCUCPUEE1APID aper.Integer     `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	GNBCUUPUEE1APID aper.Integer     `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	TraceID         aper.OctetString `aper:"lb:8,ub:8,mandatory,ignore,ext"`
}
