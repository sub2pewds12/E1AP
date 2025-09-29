package e1ap_ies

// SecurityAlgorithm represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:2129
type SecurityAlgorithm struct {
	CipheringAlgorithm           CipheringAlgorithm            `asn1:"mandatory,ext"`
	IntegrityProtectionAlgorithm *IntegrityProtectionAlgorithm `asn1:"optional,ext"`
}
