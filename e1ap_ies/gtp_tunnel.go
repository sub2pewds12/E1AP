package e1ap_ies

import "github.com/lvdund/ngap/aper"

// GTPTunnel From: 9_4_5_Information_Element_Definitions.txt:1208
// ASN.1 Data Type: SEQUENCE
type GTPTunnel struct {
	TransportLayerAddress aper.BitString   `aper:"mandatory,ignore,ext"`
	GTPTEID               aper.OctetString `aper:"lb:4,ub:4,mandatory,ext"`
}
