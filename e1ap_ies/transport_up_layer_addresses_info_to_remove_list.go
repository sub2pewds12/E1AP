package e1ap_ies

// TransportUPLayerAddressesInfoToRemoveList From: 9_4_5_Information_Element_Definitions.txt:2344
type TransportUPLayerAddressesInfoToRemoveList []TransportUPLayerAddressesInfoToRemoveItem

// TransportUPLayerAddressesInfoToRemoveItem From: 9_4_5_Information_Element_Definitions.txt:2346
type TransportUPLayerAddressesInfoToRemoveItem struct {
	IPSecTransportLayerAddress         []byte       `asn1:"mandatory,ext"`
	GTPTransportLayerAddressesToRemove []GTPTLAItem `asn1:"optional,ext"`
}
