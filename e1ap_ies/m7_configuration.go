package e1ap_ies

// M7Configuration From: 9_4_5_Information_Element_Definitions.txt:1365
type M7Configuration struct {
	M7period     int64      `asn1:"mandatory,ext"`
	M7LinksToLog LinksToLog `asn1:"mandatory,ext"`
}
