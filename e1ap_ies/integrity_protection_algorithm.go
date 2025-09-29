package e1ap_ies

// IntegrityProtectionAlgorithm represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1248
type IntegrityProtectionAlgorithm int32

const (
	IntegrityProtectionAlgorithm_NIA0     IntegrityProtectionAlgorithm = 0
	IntegrityProtectionAlgorithm_I128NIA1 IntegrityProtectionAlgorithm = 1
	IntegrityProtectionAlgorithm_I128NIA2 IntegrityProtectionAlgorithm = 2
	IntegrityProtectionAlgorithm_I128NIA3 IntegrityProtectionAlgorithm = 3
)
