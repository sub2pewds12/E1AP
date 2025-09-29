package e1ap_ies

// EarlyForwardingCOUNTReq represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:975
type EarlyForwardingCOUNTReq int32

const (
	EarlyForwardingCOUNTReq_FirstDlCount EarlyForwardingCOUNTReq = 0
	EarlyForwardingCOUNTReq_DlDiscarding EarlyForwardingCOUNTReq = 1
)
