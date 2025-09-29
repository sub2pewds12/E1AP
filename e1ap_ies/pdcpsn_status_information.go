package e1ap_ies

// PDCPSNStatusInformation represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1616
type PDCPSNStatusInformation struct {
	PdcpStatusTransferUL DRBBStatusTransfer          `asn1:"mandatory,ext"`
	PdcpStatusTransferDL PDCPCount                   `asn1:"mandatory,ext"`
	IEExtension          *ProtocolExtensionContainer `asn1:"optional,ext"`
}
