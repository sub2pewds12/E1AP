package e1ap_ies

// EHCDownlinkParameters represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:986
type EHCDownlinkParameters struct {
	DRBContinueEHCDL EHCDownlinkParametersDRBContinueEHCDL `asn1:"mandatory,ext"`
}
