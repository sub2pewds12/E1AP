package e1ap_ies

// EHCParameters is a generated SEQUENCE type.
type EHCParameters struct {
	EhcCommon   EHCCommonParameters    `aper:"mandatory"`
	EhcDownlink *EHCDownlinkParameters `aper:"optional"`
	EhcUplink   *EHCUplinkParameters   `aper:"optional"`
}
