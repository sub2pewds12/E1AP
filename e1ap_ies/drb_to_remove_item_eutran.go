package e1ap_ies

// DRBToRemoveItemEUTRAN is a generated SEQUENCE type.
type DRBToRemoveItemEUTRAN struct {
	DRBID        DRBID                       `aper:"lb:1,ub:32,mandatory,ext"`
	IEExtensions *ProtocolExtensionContainer `aper:"optional,ext"`
}
