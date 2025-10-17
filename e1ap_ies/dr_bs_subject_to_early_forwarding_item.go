package e1ap_ies

// DRBsSubjectToEarlyForwardingItem is a generated SEQUENCE type.
type DRBsSubjectToEarlyForwardingItem struct {
	DRBID        DRBID                       `aper:"lb:1,ub:32,mandatory,ext"`
	DLCountValue PDCPCount                   `aper:"mandatory,ext"`
	IEExtensions *ProtocolExtensionContainer `aper:"optional,ext"`
}
