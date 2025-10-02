package e1ap_ies

// EarlyForwardingCOUNTInfo From: 9_4_5_Information_Element_Definitions.txt:965
const (
	EarlyForwardingCOUNTInfoPresentNothing uint64 = iota
	EarlyForwardingCOUNTInfoPresentFirstDLCount
	EarlyForwardingCOUNTInfoPresentDLDiscardingCount
	EarlyForwardingCOUNTInfoPresentChoiceExtension
)

type EarlyForwardingCOUNTInfo struct {
	Choice            uint64
	FirstDLCount      *FirstDLCount
	DLDiscardingCount *DLDiscarding
	ChoiceExtension   *ProtocolIESingleContainer
}
