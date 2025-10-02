package e1ap_ies

// DRBBStatusTransfer From: 9_4_5_Information_Element_Definitions.txt:1634
type DRBBStatusTransfer struct {
	ReceiveStatusofPDCPSDU *[]byte                     `asn1:"optional,ext"`
	CountValue             PDCPCount                   `asn1:"mandatory,ext"`
	IEExtension            *ProtocolExtensionContainer `asn1:"optional,ext"`
}
