package e1ap_ies

// SDAPHeaderUL represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:2214
type SDAPHeaderUL int32

const (
	SDAPHeaderUL_Present SDAPHeaderUL = 0
	SDAPHeaderUL_Absent  SDAPHeaderUL = 1
)
