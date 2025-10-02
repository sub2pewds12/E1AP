package e1ap_ies

// QOSParametersSupportList From: 9_4_5_Information_Element_Definitions.txt:1989
type QOSParametersSupportList struct {
	EUTRANQOSSupportList []EUTRANQOSSupportItem `asn1:"optional,ext"`
	NGRANQOSSupportList  []NGRANQOSSupportItem  `asn1:"optional,ext"`
}
