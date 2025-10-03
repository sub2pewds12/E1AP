package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DataUsagePerQOSFlowItem From: 9_4_5_Information_Element_Definitions.txt:359
// ASN.1 Data Type: SEQUENCE
type DataUsagePerQOSFlowItem struct {
	QOSFlowIdentifier      aper.Integer                            `aper:"lb:0,ub:63,mandatory,ext"`
	SecondaryRATType       DataUsagePerQOSFlowItemSecondaryRATType `aper:"mandatory,ext"`
	QOSFlowTimedReportList SEQUENCE                                `aper:"mandatory,ext"`
}
