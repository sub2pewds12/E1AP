package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// CPTNLInformation From: 9_4_5_Information_Element_Definitions.txt:266
// ASN.1 Data Type: CHOICE
const (
	CPTNLInformationPresentNothing uint64 = iota
	CPTNLInformationPresentEndpointIPAddress
)

type CPTNLInformation struct {
	Choice            uint64
	EndpointIPAddress *aper.BitString
}
