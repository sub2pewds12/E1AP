package e1ap_ies

// MDTConfiguration From: 9_4_5_Information_Element_Definitions.txt:1384
// ASN.1 Data Type: SEQUENCE
type MDTConfiguration struct {
	MdtActivation MDTActivation `aper:"mandatory,ext"`
	MDTMode       MDTMode       `aper:"mandatory,ext"`
}
