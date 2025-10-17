package e1ap_ies

// DRBsSubjectToCounterCheckItemNGRAN is a generated SEQUENCE type.
type DRBsSubjectToCounterCheckItemNGRAN struct {
	PDUSessionID PDUSessionID                `aper:"lb:0,ub:255,mandatory,ext"`
	DRBID        DRBID                       `aper:"lb:1,ub:32,mandatory,ext"`
	PDCPULCount  PDCPCount                   `aper:"mandatory,ext"`
	PDCPDLCount  PDCPCount                   `aper:"mandatory,ext"`
	IEExtensions *ProtocolExtensionContainer `aper:"optional,ext"`
}
