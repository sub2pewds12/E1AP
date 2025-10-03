package e1ap_ies

import "github.com/lvdund/ngap/aper"

// NPNContextInfoSNPN From: 9_4_5_Information_Element_Definitions.txt:1481
// ASN.1 Data Type: SEQUENCE
type NPNContextInfoSNPN struct {
	NID aper.BitString `aper:"lb:44,ub:44,mandatory"`
}
