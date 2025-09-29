package e1ap_ies

// QOSMappingInformation represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:2049
type QOSMappingInformation struct {
	Dscp      *[]byte `asn1:"lb:6,ub:6,optional,ext"`
	FlowLabel *[]byte `asn1:"lb:20,ub:20,optional,ext"`
}
