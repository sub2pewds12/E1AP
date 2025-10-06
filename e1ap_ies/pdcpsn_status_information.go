package e1ap_ies

// PDCPSNStatusInformation is a generated SEQUENCE type.
type PDCPSNStatusInformation struct {
	PdcpStatusTransferUL DRBBStatusTransfer          `aper:"mandatory,ext"`
	PdcpStatusTransferDL PDCPCount                   `aper:"mandatory,ext"`
	IEExtension          *ProtocolExtensionContainer `aper:"optional,ext"`
}
