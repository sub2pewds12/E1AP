package e1ap_ies

// DRBConfirmModifiedItemEUTRAN is a generated SEQUENCE type.
type DRBConfirmModifiedItemEUTRAN struct {
	DRBID                DRBID                       `aper:"lb:1,ub:32,mandatory,ext"`
	CellGroupInformation []CellGroupInformationItem  `aper:"optional,ext"`
	IEExtensions         *ProtocolExtensionContainer `aper:"optional,ext"`
}
