package e1ap_ies

// MDTMode represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1394
// MDTMode represents a CHOICE type.
type MDTMode struct {
	ImmediateMDT ImmediateMDT `asn1:"mandatory"`
}
