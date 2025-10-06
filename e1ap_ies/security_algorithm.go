package e1ap_ies

// SecurityAlgorithm is a generated SEQUENCE type.
type SecurityAlgorithm struct {
	CipheringAlgorithm           CipheringAlgorithm            `aper:"mandatory,ext"`
	IntegrityProtectionAlgorithm *IntegrityProtectionAlgorithm `aper:"optional,ext"`
}
