package e1ap_ies

import "github.com/lvdund/ngap/aper"

// QOSFlowFailedItem From: 9_4_5_Information_Element_Definitions.txt:1963
// ASN.1 Data Type: SEQUENCE
type QOSFlowFailedItem struct {
	QOSFlowIdentifier aper.Integer `aper:"lb:0,ub:63,mandatory,ext"`
	Cause             Cause        `aper:"mandatory,ignore,ext"`
}
