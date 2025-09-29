package e1ap_ies

// MDTActivation represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1378
type MDTActivation int32

const (
	MDTActivation_ImmediateMDTOnly     MDTActivation = 0
	MDTActivation_ImmediateMDTAndTrace MDTActivation = 1
)
