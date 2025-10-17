package e1ap_ies

// DRBToRemoveItemNGRAN is a generated SEQUENCE type.
type DRBToRemoveItemNGRAN struct {
	DRBID        DRBID                       `aper:"lb:1,ub:32,mandatory,ext"`
	IEExtensions *ProtocolExtensionContainer `aper:"optional,ext"`
}
