package e1ap_ies

import "github.com/lvdund/ngap/aper"

// SNSSAI From: 9_4_5_Information_Element_Definitions.txt:2185
// ASN.1 Data Type: SEQUENCE
type SNSSAI struct {
	SST aper.OctetString  `aper:"lb:1,ub:1,mandatory,ext"`
	SD  *aper.OctetString `aper:"lb:3,ub:3,optional,ext"`
}
