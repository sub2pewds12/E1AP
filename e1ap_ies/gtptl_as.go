package e1ap_ies

// GTPTLAs From: 9_4_5_Information_Element_Definitions.txt:1196
type GTPTLAs []GTPTLAItem

// GTPTLAItem From: 9_4_5_Information_Element_Definitions.txt:1198
type GTPTLAItem struct {
	GTPTransportLayerAddresses []byte `asn1:"mandatory,ext"`
}
