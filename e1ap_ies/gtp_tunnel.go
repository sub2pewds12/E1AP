package e1ap_ies

// GTPTunnel From: 9_4_5_Information_Element_Definitions.txt:1208
type GTPTunnel struct {
	TransportLayerAddress []byte `asn1:"mandatory,ext"`
	GTPTEID               []byte `asn1:"lb:4,ub:4,mandatory,ext"`
}
