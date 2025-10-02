package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DataDiscardRequired From: 9_4_5_Information_Element_Definitions.txt:958
const (
	DataDiscardRequiredRequired aper.Enumerated = 0
)

type DataDiscardRequired struct {
	Value aper.Enumerated
}
