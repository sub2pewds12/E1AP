package e1ap_ies

// DLDiscarding From: 9_4_5_Information_Element_Definitions.txt:393
// ASN.1 Data Type: SEQUENCE
type DLDiscarding struct {
	DLDiscardingCountVal PDCPCount `aper:"mandatory"`
}
