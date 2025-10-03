package e1ap_ies

import "github.com/lvdund/ngap/aper"

// UEAssociatedLogicalE1ConnectionItem From: 9_4_5_Information_Element_Definitions.txt:2364
// ASN.1 Data Type: SEQUENCE
type UEAssociatedLogicalE1ConnectionItem struct {
	GNBCUCPUEE1APID *aper.Integer `aper:"lb:0,ub:4294967295,optional,reject,ext"`
	GNBCUUPUEE1APID *aper.Integer `aper:"lb:0,ub:4294967295,optional,reject,ext"`
}
