package e1ap_ies

// SecurityIndication represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:2140
type SecurityIndication struct {
	IntegrityProtectionIndication       IntegrityProtectionIndication       `asn1:"mandatory,ext"`
	ConfidentialityProtectionIndication ConfidentialityProtectionIndication `asn1:"mandatory,ext"`
	MaximumIPdatarate                   *MaximumIPdatarate                  `asn1:"optional,ext"`
}
