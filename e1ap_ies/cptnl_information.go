package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// CPTNLInformation is a generated CHOICE type.
type CPTNLInformation struct {
	Choice            uint64
	EndpointIPAddress *aper.BitString
	ChoiceExtension   *ProtocolIESingleContainer
}

const (
	CPTNLInformationPresentNothing uint64 = iota
	CPTNLInformationPresentEndpointIPAddress
	CPTNLInformationPresentChoiceExtension
)
