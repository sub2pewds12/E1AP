package e1ap_ies

// RedundantPDUSessionInformation From: 9_4_5_Information_Element_Definitions.txt:2067
// ASN.1 Data Type: SEQUENCE
type RedundantPDUSessionInformation struct {
	RSN RSN `aper:"mandatory,ext"`
}
