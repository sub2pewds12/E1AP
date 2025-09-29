package e1ap_ies

// IntegrityProtectionIndication represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1241
type IntegrityProtectionIndication int32

const (
	IntegrityProtectionIndication_Required  IntegrityProtectionIndication = 0
	IntegrityProtectionIndication_Preferred IntegrityProtectionIndication = 1
	IntegrityProtectionIndication_NotNeeded IntegrityProtectionIndication = 2
)
