package e1ap_ies

// M4Configuration From: 9_4_5_Information_Element_Definitions.txt:1339
// ASN.1 Data Type: SEQUENCE
type M4Configuration struct {
	M4period     M4period   `aper:"mandatory,ext"`
	M4LinksToLog LinksToLog `aper:"mandatory,ext"`
}
