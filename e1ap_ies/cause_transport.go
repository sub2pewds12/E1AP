package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// CauseTransport From: 9_4_5_Information_Element_Definitions.txt:207
// ASN.1 Data Type: ENUMERATED
const (
	CauseTransportUnspecified                  aper.Enumerated = 0
	CauseTransportTransportResourceUnavailable aper.Enumerated = 1
	CauseTransportUnknownTNLAddressForIAB      aper.Enumerated = 2
)

type CauseTransport struct {
	Value aper.Enumerated
}
