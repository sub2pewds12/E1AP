package e1ap_ies

// QOSParametersSupportList From: 9_4_5_Information_Element_Definitions.txt:1989
// ASN.1 Data Type: SEQUENCE
type QOSParametersSupportList struct {
	EUTRANQOSSupportList []EUTRANQOSSupportItem `aper:"optional,ext"`
	NGRANQOSSupportList  []NGRANQOSSupportItem  `aper:"optional,ext"`
}
