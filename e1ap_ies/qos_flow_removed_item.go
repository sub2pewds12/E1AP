package e1ap_ies

// QOSFlowRemovedItem is a generated SEQUENCE type.
type QOSFlowRemovedItem struct {
	QOSFlowIdentifier             QOSFlowIdentifier                                `aper:"lb:0,ub:63,mandatory,ext"`
	QOSFlowReleasedInSession      *QOSFlowRemovedItemQOSFlowReleasedInSession      `aper:"optional,ext"`
	QOSFlowAccumulatedSessionTime *QOSFlowRemovedItemQOSFlowAccumulatedSessionTime `aper:"lb:5,ub:5,optional,ext"`
	IEExtensions                  *ProtocolExtensionContainer                      `aper:"optional,ext"`
}
