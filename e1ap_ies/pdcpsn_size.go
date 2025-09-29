package e1ap_ies

// PDCPSNSize represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1610
type PDCPSNSize int32

const (
	PDCPSNSize_S12 PDCPSNSize = 0
	PDCPSNSize_S18 PDCPSNSize = 1
)
