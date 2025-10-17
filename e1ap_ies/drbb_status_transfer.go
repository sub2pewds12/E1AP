package e1ap_ies

// DRBBStatusTransfer is a generated SEQUENCE type.
type DRBBStatusTransfer struct {
	ReceiveStatusofPDCPSDU *DRBBStatusTransferReceiveStatusofPDCPSDU `aper:"lb:1,ub:131072,optional,ext"`
	CountValue             PDCPCount                                 `aper:"mandatory,ext"`
	IEExtension            *ProtocolExtensionContainer               `aper:"optional,ext"`
}
