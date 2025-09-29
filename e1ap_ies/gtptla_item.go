package e1ap_ies

// GTPTLAItem represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1198
type GTPTLAItem struct {
	GTPTransportLayerAddresses []byte `asn1:"lb:1,ub:160,mandatory,ext"`
}
