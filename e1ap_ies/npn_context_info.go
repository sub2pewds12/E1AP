package e1ap_ies

// NPNContextInfo is a generated CHOICE type.
type NPNContextInfo struct {
	Choice          uint64
	SNPN            *NPNContextInfoSNPN
	ChoiceExtension *ProtocolIESingleContainer
}

const (
	NPNContextInfoPresentNothing uint64 = iota
	NPNContextInfoPresentSNPN
	NPNContextInfoPresentChoiceExtension
)
