package e1ap_ies

// GTPTunnel represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1208
type GTPTunnel struct {
	TransportLayerAddress []byte `asn1:"lb:1,ub:160,mandatory,ext"`
	GTPTEID               []byte `asn1:"mandatory,ext"`
}
