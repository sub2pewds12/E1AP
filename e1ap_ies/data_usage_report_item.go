package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DataUsageReportItem is a generated SEQUENCE type.
type DataUsageReportItem struct {
	DRBID              aper.Integer                `aper:"lb:1,ub:32,mandatory,ext"`
	RATType            RATType                     `aper:"mandatory,ext"`
	DRBUsageReportList []DRBUsageReportItem        `aper:"lb:1,ub:Maxnooftimeperiods,mandatory,ext"`
	IEExtensions       *ProtocolExtensionContainer `aper:"optional,ext"`
}
