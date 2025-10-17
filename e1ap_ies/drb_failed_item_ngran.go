package e1ap_ies

// DRBFailedItemNGRAN is a generated SEQUENCE type.
type DRBFailedItemNGRAN struct {
	DRBID        DRBID                       `aper:"lb:1,ub:32,mandatory,ext"`
	Cause        Cause                       `aper:"mandatory,ignore,ext"`
	IEExtensions *ProtocolExtensionContainer `aper:"optional,ext"`
}
