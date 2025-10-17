package e1ap_ies

// DRBStatusItem is a generated SEQUENCE type.
type DRBStatusItem struct {
	DRBID        DRBID                       `aper:"lb:1,ub:32,mandatory,ext"`
	PDCPDLCount  *PDCPCount                  `aper:"optional,ext"`
	PDCPULCount  *PDCPCount                  `aper:"optional,ext"`
	IEExtensions *ProtocolExtensionContainer `aper:"optional,ext"`
}
