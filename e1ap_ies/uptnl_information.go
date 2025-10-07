package e1ap_ies

// UPTNLInformation is a generated CHOICE type.
type UPTNLInformation struct {
	Choice          uint64
	GTPTunnel       *GTPTunnel
	ChoiceExtension *ProtocolIESingleContainer
}

const (
	UPTNLInformationPresentNothing uint64 = iota
	UPTNLInformationPresentGTPTunnel
	UPTNLInformationPresentChoiceExtension
)
