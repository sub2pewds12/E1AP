package e1ap_ies

import "github.com/lvdund/ngap/aper"

// UPSecuritykey is a generated SEQUENCE type.
type UPSecuritykey struct {
	EncryptionKey          aper.OctetString  `aper:"mandatory,ext"`
	IntegrityProtectionKey *aper.OctetString `aper:"optional,ext"`
}
