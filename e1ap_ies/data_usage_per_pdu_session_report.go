package e1ap_ies

// DataUsagePerPDUSessionReport is a generated SEQUENCE type.
type DataUsagePerPDUSessionReport struct {
	SecondaryRATType          DataUsagePerPDUSessionReportSecondaryRATType `aper:"mandatory,ext"`
	PDUSessionTimedReportList []MRDCDataUsageReportItem                    `aper:"lb:1,ub:Maxnooftimeperiods,mandatory,ext"`
	IEExtensions              *ProtocolExtensionContainer                  `aper:"optional,ext"`
}
