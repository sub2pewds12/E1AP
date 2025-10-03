package e1ap_ies

import "github.com/lvdund/ngap/aper"

// QOSFlowMappingItem From: 9_4_5_Information_Element_Definitions.txt:1976
// ASN.1 Data Type: SEQUENCE
type QOSFlowMappingItem struct {
	QOSFlowIdentifier        aper.Integer              `aper:"lb:0,ub:63,mandatory,ext"`
	QoSFlowMappingIndication *QOSFlowMappingIndication `aper:"optional,ext"`
}
