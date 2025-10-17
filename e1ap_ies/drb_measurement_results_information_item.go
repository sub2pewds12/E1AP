package e1ap_ies

// DRBMeasurementResultsInformationItem is a generated SEQUENCE type.
type DRBMeasurementResultsInformationItem struct {
	DRBID        DRBID                                           `aper:"lb:1,ub:32,mandatory,ext"`
	ULD1Result   *DRBMeasurementResultsInformationItemULD1Result `aper:"lb:0,ub:10000,optional,ext"`
	IEExtensions *ProtocolExtensionContainer                     `aper:"optional,ext"`
}
