package e1ap_ies

import "github.com/lvdund/ngap/aper"

// PacketErrorRate From: 9_4_5_Information_Element_Definitions.txt:1525
// ASN.1 Data Type: SEQUENCE
type PacketErrorRate struct {
	PERScalar   aper.Integer `aper:"mandatory,ext"`
	PERExponent aper.Integer `aper:"mandatory,ext"`
}
