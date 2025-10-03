package e1ap_ies

import "github.com/lvdund/ngap/aper"

// UPSecuritykey From: 9_4_5_Information_Element_Definitions.txt:2410
// ASN.1 Data Type: SEQUENCE
type UPSecuritykey struct {
	EncryptionKey          aper.OctetString  `aper:"mandatory,ext"`
	IntegrityProtectionKey *aper.OctetString `aper:"optional,ext"`
}
