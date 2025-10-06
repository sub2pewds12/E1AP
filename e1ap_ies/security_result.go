package e1ap_ies

// SecurityResult is a generated SEQUENCE type.
type SecurityResult struct {
	IntegrityProtectionResult       IntegrityProtectionResult       `aper:"mandatory,ext"`
	ConfidentialityProtectionResult ConfidentialityProtectionResult `aper:"mandatory,ext"`
}
