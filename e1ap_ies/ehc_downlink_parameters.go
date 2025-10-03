package e1ap_ies

// EHCDownlinkParameters From: 9_4_5_Information_Element_Definitions.txt:986
// ASN.1 Data Type: SEQUENCE
type EHCDownlinkParameters struct {
	DRBContinueEHCDL EHCDownlinkParametersDRBContinueEHCDL `aper:"mandatory,ext"`
}
