package e1ap_ies

// DRBFailedToModifyItemEUTRAN is a generated SEQUENCE type.
type DRBFailedToModifyItemEUTRAN struct {
	DRBID        DRBID                       `aper:"lb:1,ub:32,mandatory,ext"`
	Cause        Cause                       `aper:"mandatory,ignore,ext"`
	IEExtensions *ProtocolExtensionContainer `aper:"optional,ext"`
}
