package e1ap_ies

// TransportUPLayerAddressesInfoToAddList From: 9_4_5_Information_Element_Definitions.txt:2331
type TransportUPLayerAddressesInfoToAddList []TransportUPLayerAddressesInfoToAddItem

// TransportUPLayerAddressesInfoToAddItem From: 9_4_5_Information_Element_Definitions.txt:2333
type TransportUPLayerAddressesInfoToAddItem struct {
	IPSecTransportLayerAddress      []byte       `asn1:"mandatory,ext"`
	GTPTransportLayerAddressesToAdd []GTPTLAItem `asn1:"optional,ext"`
}
