package e1ap_ies

// SDAPConfiguration From: 9_4_5_Information_Element_Definitions.txt:2196
// ASN.1 Data Type: SEQUENCE
type SDAPConfiguration struct {
	DefaultDRB   DefaultDRB   `aper:"mandatory,ext"`
	SDAPHeaderUL SDAPHeaderUL `aper:"mandatory,ext"`
	SDAPHeaderDL SDAPHeaderDL `aper:"mandatory,ext"`
}
