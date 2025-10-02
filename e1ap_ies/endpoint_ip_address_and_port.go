package e1ap_ies

// EndpointIPAddressAndPort From: 9_4_5_Information_Element_Definitions.txt:1017
type EndpointIPAddressAndPort struct {
	EndpointIPAddress []byte `asn1:"mandatory"`
	PortNumber        []byte `asn1:"lb:16,ub:16,mandatory"`
}
