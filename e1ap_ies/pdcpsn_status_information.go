package e1ap_ies

// PDCPSNStatusInformation From: 9_4_5_Information_Element_Definitions.txt:1616
// ASN.1 Data Type: SEQUENCE
type PDCPSNStatusInformation struct {
	PdcpStatusTransferUL DRBBStatusTransfer          `aper:"mandatory,ext"`
	PdcpStatusTransferDL PDCPCount                   `aper:"mandatory,ext"`
	IEExtension          *ProtocolExtensionContainer `aper:"optional,ext"`
}
