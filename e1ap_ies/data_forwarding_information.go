package e1ap_ies

// DataForwardingInformation represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:328
type DataForwardingInformation struct {
	ULDataForwarding *UPTNLInformation `asn1:"optional,ext"`
	DLDataForwarding *UPTNLInformation `asn1:"optional,ext"`
}
