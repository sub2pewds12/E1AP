package e1ap_ies

import "github.com/lvdund/ngap/aper"

// TransportUPLayerAddressesInfoToAddItem From: 9_4_5_Information_Element_Definitions.txt:2333
// ASN.1 Data Type: SEQUENCE
type TransportUPLayerAddressesInfoToAddItem struct {
	IPSecTransportLayerAddress      aper.BitString `aper:"mandatory,ignore,ext"`
	GTPTransportLayerAddressesToAdd []GTPTLAItem   `aper:"optional,ext"`
}
