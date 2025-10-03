package e1ap_ies

import "github.com/lvdund/ngap/aper"

// GTPTLAItem From: 9_4_5_Information_Element_Definitions.txt:1198
// ASN.1 Data Type: SEQUENCE
type GTPTLAItem struct {
	GTPTransportLayerAddresses aper.BitString `aper:"mandatory,ignore,ext"`
}
