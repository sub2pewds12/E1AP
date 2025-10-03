package e1ap_ies

import "github.com/lvdund/ngap/aper"

// ImmediateMDT From: 9_4_5_Information_Element_Definitions.txt:1268
// ASN.1 Data Type: SEQUENCE
type ImmediateMDT struct {
	MeasurementsToActivate aper.BitString   `aper:"lb:8,ub:8,mandatory,ext"`
	MeasurementFour        *M4Configuration `aper:"optional,ext"`
	MeasurementSix         *M6Configuration `aper:"optional,ext"`
	MeasurementSeven       *M7Configuration `aper:"optional,ext"`
}
