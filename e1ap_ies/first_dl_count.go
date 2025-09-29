package e1ap_ies

// FirstDLCount represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1069
type FirstDLCount struct {
	FirstDLCountVal PDCPCount `asn1:"mandatory"`
}
