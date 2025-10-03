package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// QOSFlowRemovedItemQOSFlowReleasedInSession From: 9_4_5_Information_Element_Definitions.txt:2037
// ASN.1 Data Type: ENUMERATED
const (
	QOSFlowRemovedItemQOSFlowReleasedInSessionReleasedInSession    aper.Enumerated = 0
	QOSFlowRemovedItemQOSFlowReleasedInSessionNotReleasedInSession aper.Enumerated = 1
)

type QOSFlowRemovedItemQOSFlowReleasedInSession struct {
	Value aper.Enumerated
}
