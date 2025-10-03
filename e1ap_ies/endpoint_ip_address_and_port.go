package e1ap_ies

import "github.com/lvdund/ngap/aper"

// EndpointIPAddressAndPort From: 9_4_5_Information_Element_Definitions.txt:1017
// ASN.1 Data Type: SEQUENCE
type EndpointIPAddressAndPort struct {
	EndpointIPAddress aper.BitString `aper:"mandatory,ignore"`
	PortNumber        aper.BitString `aper:"lb:16,ub:16,mandatory"`
}
