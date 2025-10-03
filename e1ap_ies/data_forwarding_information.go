package e1ap_ies

// DataForwardingInformation From: 9_4_5_Information_Element_Definitions.txt:328
// ASN.1 Data Type: SEQUENCE
type DataForwardingInformation struct {
	ULDataForwarding *UPTNLInformation `aper:"optional,ext"`
	DLDataForwarding *UPTNLInformation `aper:"optional,ext"`
}
