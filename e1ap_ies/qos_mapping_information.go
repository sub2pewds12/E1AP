package e1ap_ies

// QOSMappingInformation From: 9_4_5_Information_Element_Definitions.txt:2049
type QOSMappingInformation struct {
	Dscp      *[]byte `asn1:"optional,ext"`
	FlowLabel *[]byte `asn1:"optional,ext"`
}
