package e1ap_ies

// DLDiscarding represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:393
type DLDiscarding struct {
	DLDiscardingCountVal PDCPCount `asn1:"mandatory"`
}
