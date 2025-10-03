package e1ap_ies

// TransportLayerAddressInfo From: 9_4_5_Information_Element_Definitions.txt:2320
// ASN.1 Data Type: SEQUENCE
type TransportLayerAddressInfo struct {
	TransportUPLayerAddressesInfoToAddList    []TransportUPLayerAddressesInfoToAddItem    `aper:"optional,ext"`
	TransportUPLayerAddressesInfoToRemoveList []TransportUPLayerAddressesInfoToRemoveItem `aper:"optional,ext"`
}
