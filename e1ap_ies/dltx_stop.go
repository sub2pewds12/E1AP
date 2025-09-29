package e1ap_ies

// DLTXStop represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:413
type DLTXStop int32

const (
	DLTXStop_Stop   DLTXStop = 0
	DLTXStop_Resume DLTXStop = 1
)
