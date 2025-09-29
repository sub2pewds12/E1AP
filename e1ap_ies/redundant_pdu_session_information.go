package e1ap_ies

// RedundantPDUSessionInformation represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:2067
type RedundantPDUSessionInformation struct {
	RSN RSN `asn1:"mandatory,ext"`
}
