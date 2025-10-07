package e1ap_ies

import "github.com/lvdund/ngap/aper"

// QOSFlowFailedItem is a generated SEQUENCE type.
type QOSFlowFailedItem struct {
	QOSFlowIdentifier aper.Integer                `aper:"lb:0,ub:63,mandatory,ext"`
	Cause             Cause                       `aper:"mandatory,ignore,ext"`
	IEExtensions      *ProtocolExtensionContainer `aper:"optional,ext"`
}
