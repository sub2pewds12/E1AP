package e1ap_ies

// MDTConfiguration represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1384
type MDTConfiguration struct {
	MdtActivation MDTActivation `asn1:"mandatory,ext"`
	MDTMode       MDTMode       `asn1:"mandatory,ext"`
}
