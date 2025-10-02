package e1ap_ies

import "github.com/lvdund/ngap/aper"

// Dynamic5QIDescriptorDelayCritical From: 9_4_5_Information_Element_Definitions.txt:940
const (
	Dynamic5QIDescriptorDelayCriticalDelayCritical    aper.Enumerated = 0
	Dynamic5QIDescriptorDelayCriticalNonDelayCritical aper.Enumerated = 1
)

type Dynamic5QIDescriptorDelayCritical struct {
	Value aper.Enumerated
}
