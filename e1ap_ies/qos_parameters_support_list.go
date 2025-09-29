package e1ap_ies

// QOSParametersSupportList represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1989
type QOSParametersSupportList struct {
	EUTRANQOSSupportList []EUTRANQOSSupportItem `asn1:"lb:1,ub:Item,optional,ext"`
	NGRANQOSSupportList  []NGRANQOSSupportItem  `asn1:"lb:1,ub:Item,optional,ext"`
}
