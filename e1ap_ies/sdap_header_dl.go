package e1ap_ies

// SDAPHeaderDL represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:2208
type SDAPHeaderDL int32

const (
	SDAPHeaderDL_Present SDAPHeaderDL = 0
	SDAPHeaderDL_Absent  SDAPHeaderDL = 1
)
