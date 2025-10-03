package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DLDataNotification From: 9_4_4_PDU_Definitions.txt:1109
// ASN.1 Data Type: SEQUENCE
type DLDataNotification struct {
	GNBCUCPUEE1APID aper.Integer  `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	GNBCUUPUEE1APID aper.Integer  `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	PPI             *aper.Integer `aper:"optional,ignore,ext"`
}
