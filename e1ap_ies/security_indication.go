package e1ap_ies

// SecurityIndication From: 9_4_5_Information_Element_Definitions.txt:2140
// ASN.1 Data Type: SEQUENCE
type SecurityIndication struct {
	IntegrityProtectionIndication       IntegrityProtectionIndication       `aper:"mandatory,ext"`
	ConfidentialityProtectionIndication ConfidentialityProtectionIndication `aper:"mandatory,ext"`
	MaximumIPdatarate                   *MaximumIPdatarate                  `aper:"optional,ext"`
}
