package e1ap_ies

// TransportLayerAddressInfo From: 9_4_5_Information_Element_Definitions.txt:2320
type TransportLayerAddressInfo struct {
	TransportUPLayerAddressesInfoToAddList    []TransportUPLayerAddressesInfoToAddItem    `asn1:"optional,ext"`
	TransportUPLayerAddressesInfoToRemoveList []TransportUPLayerAddressesInfoToRemoveItem `asn1:"optional,ext"`
}
