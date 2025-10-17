package e1ap_ies

// ImmediateMDT is a generated SEQUENCE type.
type ImmediateMDT struct {
	MeasurementsToActivate MeasurementsToActivate      `aper:"lb:8,ub:8,mandatory,ext"`
	MeasurementFour        *M4Configuration            `aper:"optional,ext"`
	MeasurementSix         *M6Configuration            `aper:"optional,ext"`
	MeasurementSeven       *M7Configuration            `aper:"optional,ext"`
	IEExtensions           *ProtocolExtensionContainer `aper:"optional,ext"`
}
