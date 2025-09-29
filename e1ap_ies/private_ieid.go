package e1ap_ies

// PrivateIEID represents the ASN.1 definition from 9_4_6_Common_Definitions.txt:38
// PrivateIEID represents a CHOICE type.
type PrivateIEID struct {
	Local  int64  `asn1:"lb:0,ub:MaxPrivateIEs,mandatory"`
	Global string `asn1:"mandatory"`
}
