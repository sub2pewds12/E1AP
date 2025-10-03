package e1ap_ies

// EHCUplinkParameters From: 9_4_5_Information_Element_Definitions.txt:995
// ASN.1 Data Type: SEQUENCE
type EHCUplinkParameters struct {
	DRBContinueEHCUL EHCUplinkParametersDRBContinueEHCUL `aper:"mandatory,ext"`
}
