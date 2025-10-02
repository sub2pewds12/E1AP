package e1ap_ies

import "github.com/lvdund/ngap/aper"

// QOSFlowMappingIndication From: 9_4_5_Information_Element_Definitions.txt:1987
const (
	QOSFlowMappingIndicationUl aper.Enumerated = 0
	QOSFlowMappingIndicationDl aper.Enumerated = 1
)

type QOSFlowMappingIndication struct {
	Value aper.Enumerated
}
