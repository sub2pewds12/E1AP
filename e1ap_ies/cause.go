package e1ap_ies

// Cause From: 9_4_5_Information_Element_Definitions.txt:141
const (
	CausePresentNothing uint64 = iota
	CausePresentRadioNetwork
	CausePresentTransport
	CausePresentProtocol
	CausePresentMisc
)

type Cause struct {
	Choice       uint64
	RadioNetwork *CauseRadioNetwork
	Transport    *CauseTransport
	Protocol     *CauseProtocol
	Misc         *CauseMisc
}
