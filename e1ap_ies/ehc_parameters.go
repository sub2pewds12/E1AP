package e1ap_ies

// EHCParameters From: 9_4_5_Information_Element_Definitions.txt:1004
// ASN.1 Data Type: SEQUENCE
type EHCParameters struct {
	EhcCommon   EHCCommonParameters    `aper:"mandatory"`
	EhcDownlink *EHCDownlinkParameters `aper:"optional"`
	EhcUplink   *EHCUplinkParameters   `aper:"optional"`
}
