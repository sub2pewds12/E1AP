package e1ap_ies

// EHCParameters represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1004
type EHCParameters struct {
	EhcCommon   EHCCommonParameters    `asn1:"mandatory"`
	EhcDownlink *EHCDownlinkParameters `asn1:"optional"`
	EhcUplink   *EHCUplinkParameters   `asn1:"optional"`
}
