package e1ap_ies

// TransportUPLayerAddressesInfoToAddItem represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:2333
type TransportUPLayerAddressesInfoToAddItem struct {
	IPSecTransportLayerAddress      []byte       `asn1:"lb:1,ub:160,mandatory,ext"`
	GTPTransportLayerAddressesToAdd []GTPTLAItem `asn1:"lb:1,ub:Item,optional,ext"`
}
