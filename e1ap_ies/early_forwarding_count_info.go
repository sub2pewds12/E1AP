package e1ap_ies

// EarlyForwardingCOUNTInfo is a generated CHOICE type.
type EarlyForwardingCOUNTInfo struct {
	Choice            uint64
	FirstDLCount      *FirstDLCount
	DLDiscardingCount *DLDiscarding
	ChoiceExtension   *ProtocolIESingleContainer
}

const (
	EarlyForwardingCOUNTInfoPresentNothing uint64 = iota
	EarlyForwardingCOUNTInfoPresentFirstDLCount
	EarlyForwardingCOUNTInfoPresentDLDiscardingCount
	EarlyForwardingCOUNTInfoPresentChoiceExtension
)
