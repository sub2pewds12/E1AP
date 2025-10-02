package e1ap_ies

// ImmediateMDT From: 9_4_5_Information_Element_Definitions.txt:1268
type ImmediateMDT struct {
	MeasurementsToActivate []byte           `asn1:"lb:8,ub:8,mandatory,ext"`
	MeasurementFour        *M4Configuration `asn1:"optional,ext"`
	MeasurementSix         *M6Configuration `asn1:"optional,ext"`
	MeasurementSeven       *M7Configuration `asn1:"optional,ext"`
}
