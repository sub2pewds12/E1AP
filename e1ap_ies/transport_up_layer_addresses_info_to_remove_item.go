package e1ap_ies

import "github.com/lvdund/ngap/aper"

// TransportUPLayerAddressesInfoToRemoveItem From: 9_4_5_Information_Element_Definitions.txt:2346
// ASN.1 Data Type: SEQUENCE
type TransportUPLayerAddressesInfoToRemoveItem struct {
	IPSecTransportLayerAddress         aper.BitString `aper:"mandatory,ignore,ext"`
	GTPTransportLayerAddressesToRemove []GTPTLAItem   `aper:"optional,ext"`
}
