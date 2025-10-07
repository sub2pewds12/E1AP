package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DataUsagePerQOSFlowItem is a generated SEQUENCE type.
type DataUsagePerQOSFlowItem struct {
	QOSFlowIdentifier      aper.Integer                            `aper:"lb:0,ub:63,mandatory,ext"`
	SecondaryRATType       DataUsagePerQOSFlowItemSecondaryRATType `aper:"mandatory,ext"`
	QOSFlowTimedReportList []MRDCDataUsageReportItem               `aper:"lb:1,ub:Maxnooftimeperiods,mandatory,ext"`
	IEExtensions           *ProtocolExtensionContainer             `aper:"optional,ext"`
}
