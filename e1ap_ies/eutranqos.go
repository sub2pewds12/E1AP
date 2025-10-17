package e1ap_ies

// EUTRANQOS is a generated SEQUENCE type.
type EUTRANQOS struct {
	QCI                                  QCI                                  `aper:"lb:0,ub:255,mandatory,ext"`
	EUTRANallocationAndRetentionPriority EUTRANAllocationAndRetentionPriority `aper:"mandatory,ext"`
	GbrQosInformation                    *GBRQosInformation                   `aper:"optional,ext"`
	IEExtensions                         *ProtocolExtensionContainer          `aper:"optional,ext"`
}
