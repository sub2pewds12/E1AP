package e1ap_ies

// DRBActivityItem is a generated SEQUENCE type.
type DRBActivityItem struct {
	DRBID        DRBID                       `aper:"lb:1,ub:32,mandatory,ext"`
	DRBActivity  DRBActivity                 `aper:"mandatory,ext"`
	IEExtensions *ProtocolExtensionContainer `aper:"optional,ext"`
}
