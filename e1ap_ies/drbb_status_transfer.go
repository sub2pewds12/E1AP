package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DRBBStatusTransfer From: 9_4_5_Information_Element_Definitions.txt:1634
// ASN.1 Data Type: SEQUENCE
type DRBBStatusTransfer struct {
	ReceiveStatusofPDCPSDU *aper.BitString             `aper:"lb:1,ub:131072,optional,ext"`
	CountValue             PDCPCount                   `aper:"mandatory,ext"`
	IEExtension            *ProtocolExtensionContainer `aper:"optional,ext"`
}
