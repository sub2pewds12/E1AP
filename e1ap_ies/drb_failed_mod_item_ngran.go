package e1ap_ies

// DRBFailedModItemNGRAN is a generated SEQUENCE type.
type DRBFailedModItemNGRAN struct {
	DRBID        DRBID                       `aper:"lb:1,ub:32,mandatory,ext"`
	Cause        Cause                       `aper:"mandatory,ignore,ext"`
	IEExtensions *ProtocolExtensionContainer `aper:"optional,ext"`
}
