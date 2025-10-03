package e1ap_ies

import "github.com/lvdund/ngap/aper"

// NPNSupportInfoSNPN From: 9_4_5_Information_Element_Definitions.txt:1462
// ASN.1 Data Type: SEQUENCE
type NPNSupportInfoSNPN struct {
	NID aper.BitString `aper:"lb:44,ub:44,mandatory"`
}
