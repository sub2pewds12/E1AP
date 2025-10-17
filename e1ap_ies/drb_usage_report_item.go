package e1ap_ies

// DRBUsageReportItem is a generated SEQUENCE type.
type DRBUsageReportItem struct {
	StartTimeStamp DRBUsageReportItemStartTimeStamp `aper:"lb:4,ub:4,mandatory,ext"`
	EndTimeStamp   DRBUsageReportItemEndTimeStamp   `aper:"lb:4,ub:4,mandatory,ext"`
	UsageCountUL   DRBUsageReportItemUsageCountUL   `aper:"lb:0,ub:18446744073709551615,mandatory,ext"`
	UsageCountDL   DRBUsageReportItemUsageCountDL   `aper:"lb:0,ub:18446744073709551615,mandatory,ext"`
	IEExtensions   *ProtocolExtensionContainer      `aper:"optional,ext"`
}
