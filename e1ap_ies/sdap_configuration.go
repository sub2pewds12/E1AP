package e1ap_ies

// SDAPConfiguration represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:2196
type SDAPConfiguration struct {
	DefaultDRB   DefaultDRB   `asn1:"mandatory,ext"`
	SDAPHeaderUL SDAPHeaderUL `asn1:"mandatory,ext"`
	SDAPHeaderDL SDAPHeaderDL `asn1:"mandatory,ext"`
}
