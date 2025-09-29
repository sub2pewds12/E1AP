package e1ap_ies

// TransportLayerAddressInfo represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:2320
type TransportLayerAddressInfo struct {
	TransportUPLayerAddressesInfoToAddList    []TransportUPLayerAddressesInfoToAddItem    `asn1:"lb:1,ub:Item,optional,ext"`
	TransportUPLayerAddressesInfoToRemoveList []TransportUPLayerAddressesInfoToRemoveItem `asn1:"lb:1,ub:Item,optional,ext"`
}
