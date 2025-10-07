package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DRBRemovedItem is a generated SEQUENCE type.
type DRBRemovedItem struct {
	DRBID                     aper.Integer                        `aper:"lb:1,ub:32,mandatory,ext"`
	DRBReleasedInSession      *DRBRemovedItemDRBReleasedInSession `aper:"optional,ext"`
	DRBAccumulatedSessionTime *aper.OctetString                   `aper:"lb:5,ub:5,optional,ext"`
	QOSFlowRemovedList        []QOSFlowRemovedItem                `aper:"optional,ext"`
	IEExtensions              *ProtocolExtensionContainer         `aper:"optional,ext"`
}
