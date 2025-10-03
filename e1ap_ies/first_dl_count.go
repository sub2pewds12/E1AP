package e1ap_ies

// FirstDLCount From: 9_4_5_Information_Element_Definitions.txt:1069
// ASN.1 Data Type: SEQUENCE
type FirstDLCount struct {
	FirstDLCountVal PDCPCount `aper:"mandatory"`
}
