package e1ap_ies

// EHCUplinkParameters represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:995
type EHCUplinkParameters struct {
	DRBContinueEHCUL EHCUplinkParametersDRBContinueEHCUL `asn1:"mandatory,ext"`
}
