package e1ap_ies

// EHCDownlinkParameters is a generated SEQUENCE type.
type EHCDownlinkParameters struct {
	DRBContinueEHCDL EHCDownlinkParametersDRBContinueEHCDL `aper:"mandatory,ext"`
	IEExtensions     *ProtocolExtensionContainer           `aper:"optional,ext"`
}
