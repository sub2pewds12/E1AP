package e1ap_ies

// M6Configuration represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1352
type M6Configuration struct {
	M6reportInterval M6reportInterval `asn1:"mandatory,ext"`
	M6LinksToLog     LinksToLog       `asn1:"mandatory,ext"`
}
