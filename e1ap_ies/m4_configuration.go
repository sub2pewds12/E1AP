package e1ap_ies

// M4Configuration represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1339
type M4Configuration struct {
	M4period     M4period   `asn1:"mandatory,ext"`
	M4LinksToLog LinksToLog `asn1:"mandatory,ext"`
}
