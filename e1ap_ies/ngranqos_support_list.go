package e1ap_ies

// NGRANQOSSupportList From: 9_4_5_Information_Element_Definitions.txt:1426
type NGRANQOSSupportList []NGRANQOSSupportItem

// NGRANQOSSupportItem From: 9_4_5_Information_Element_Definitions.txt:1428
type NGRANQOSSupportItem struct {
	NonDynamic5QIDescriptor NonDynamic5QIDescriptor `asn1:"mandatory"`
}
