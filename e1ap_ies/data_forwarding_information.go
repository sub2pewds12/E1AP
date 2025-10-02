package e1ap_ies

// DataForwardingInformation From: 9_4_5_Information_Element_Definitions.txt:328
type DataForwardingInformation struct {
	ULDataForwarding *UPTNLInformation `asn1:"optional,ext"`
	DLDataForwarding *UPTNLInformation `asn1:"optional,ext"`
}
