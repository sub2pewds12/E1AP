package e1ap_ies

// CPTNLInformation is a generated CHOICE type.
type CPTNLInformation struct {
	Choice            uint64
	EndpointIPAddress *TransportLayerAddress
	ChoiceExtension   *ProtocolIESingleContainer
}

const (
	CPTNLInformationPresentNothing uint64 = iota
	CPTNLInformationPresentEndpointIPAddress
	CPTNLInformationPresentChoiceExtension
)
