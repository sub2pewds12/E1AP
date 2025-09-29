package e1ap_ies

// Cause represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:141
// Cause represents a CHOICE type.
type Cause struct {
	RadioNetwork CauseRadioNetwork `asn1:"mandatory"`
	Transport    CauseTransport    `asn1:"mandatory"`
	Protocol     CauseProtocol     `asn1:"mandatory"`
	Misc         CauseMisc         `asn1:"mandatory"`
}
