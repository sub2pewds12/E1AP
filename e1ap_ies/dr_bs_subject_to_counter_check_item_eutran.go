package e1ap_ies

// DRBsSubjectToCounterCheckItemEUTRAN is a generated SEQUENCE type.
type DRBsSubjectToCounterCheckItemEUTRAN struct {
	DRBID        DRBID                       `aper:"lb:1,ub:32,mandatory,ext"`
	PDCPULCount  PDCPCount                   `aper:"mandatory,ext"`
	PDCPDLCount  PDCPCount                   `aper:"mandatory,ext"`
	IEExtensions *ProtocolExtensionContainer `aper:"optional,ext"`
}
