package e1ap_ies

// M6Configuration From: 9_4_5_Information_Element_Definitions.txt:1352
// ASN.1 Data Type: SEQUENCE
type M6Configuration struct {
	M6reportInterval M6reportInterval `aper:"mandatory,ext"`
	M6LinksToLog     LinksToLog       `aper:"mandatory,ext"`
}
