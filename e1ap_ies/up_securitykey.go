package e1ap_ies

// UPSecuritykey represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:2410
type UPSecuritykey struct {
	EncryptionKey          []byte  `asn1:"mandatory,ext"`
	IntegrityProtectionKey *[]byte `asn1:"optional,ext"`
}
