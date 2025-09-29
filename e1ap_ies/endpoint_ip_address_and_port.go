package e1ap_ies

// EndpointIPAddressAndPort represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1017
type EndpointIPAddressAndPort struct {
	EndpointIPAddress []byte `asn1:"lb:1,ub:160,mandatory"`
	PortNumber        []byte `asn1:"lb:16,ub:16,mandatory"`
}
