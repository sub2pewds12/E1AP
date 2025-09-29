package e1ap_ies

// CipheringAlgorithm represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:234
type CipheringAlgorithm int32

const (
	CipheringAlgorithm_NEA0     CipheringAlgorithm = 0
	CipheringAlgorithm_C128NEA1 CipheringAlgorithm = 1
	CipheringAlgorithm_C128NEA2 CipheringAlgorithm = 2
	CipheringAlgorithm_C128NEA3 CipheringAlgorithm = 3
)
