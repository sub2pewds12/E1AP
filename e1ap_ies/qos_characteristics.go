package e1ap_ies

// QOSCharacteristics represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1936
// QOSCharacteristics represents a CHOICE type.
type QOSCharacteristics struct {
	NonDynamic5QI NonDynamic5QIDescriptor `asn1:"mandatory"`
	Dynamic5QI    Dynamic5QIDescriptor    `asn1:"mandatory"`
}
