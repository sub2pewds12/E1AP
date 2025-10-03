package e1ap_ies

// NGRANQOSSupportItem From: 9_4_5_Information_Element_Definitions.txt:1428
// ASN.1 Data Type: SEQUENCE
type NGRANQOSSupportItem struct {
	NonDynamic5QIDescriptor NonDynamic5QIDescriptor `aper:"mandatory"`
}
