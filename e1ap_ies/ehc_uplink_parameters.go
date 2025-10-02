package e1ap_ies

// EHCUplinkParameters From: 9_4_5_Information_Element_Definitions.txt:995
type EHCUplinkParameters struct {
	DRBContinueEHCUL EHCUplinkParametersDRBContinueEHCUL `asn1:"mandatory,ext"`
}
