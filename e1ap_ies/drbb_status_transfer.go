package e1ap_ies

// DRBBStatusTransfer represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1634
type DRBBStatusTransfer struct {
	ReceiveStatusofPDCPSDU *[]byte                     `asn1:"lb:1,ub:131072,optional,ext"`
	CountValue             PDCPCount                   `asn1:"mandatory,ext"`
	IEExtension            *ProtocolExtensionContainer `asn1:"optional,ext"`
}
