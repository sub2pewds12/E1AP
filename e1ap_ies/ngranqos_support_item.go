package e1ap_ies

// NGRANQOSSupportItem represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1428
type NGRANQOSSupportItem struct {
	NonDynamic5QIDescriptor NonDynamic5QIDescriptor `asn1:"mandatory"`
}
