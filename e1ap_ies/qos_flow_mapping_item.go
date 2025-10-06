package e1ap_ies

import "github.com/lvdund/ngap/aper"

// QOSFlowMappingItem is a generated SEQUENCE type.
type QOSFlowMappingItem struct {
	QOSFlowIdentifier        aper.Integer              `aper:"lb:0,ub:63,mandatory,ext"`
	QoSFlowMappingIndication *QOSFlowMappingIndication `aper:"optional,ext"`
}
