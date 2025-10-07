package e1ap_ies

// Cause is a generated CHOICE type.
type Cause struct {
	Choice          uint64
	RadioNetwork    *CauseRadioNetwork
	Transport       *CauseTransport
	Protocol        *CauseProtocol
	Misc            *CauseMisc
	ChoiceExtension *ProtocolIESingleContainer
}

const (
	CausePresentNothing uint64 = iota
	CausePresentRadioNetwork
	CausePresentTransport
	CausePresentProtocol
	CausePresentMisc
	CausePresentChoiceExtension
)
