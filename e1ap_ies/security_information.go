package e1ap_ies

// SecurityInformation is a generated SEQUENCE type.
type SecurityInformation struct {
	SecurityAlgorithm SecurityAlgorithm `aper:"mandatory,ext"`
	UPSecuritykey     UPSecuritykey     `aper:"mandatory,ext"`
}
