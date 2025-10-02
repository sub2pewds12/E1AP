package e1ap_ies

import "github.com/lvdund/ngap/aper"

// LinksToLog From: 9_4_5_Information_Element_Definitions.txt:1286
const (
	LinksToLogUplink                aper.Enumerated = 0
	LinksToLogDownlink              aper.Enumerated = 1
	LinksToLogBothUplinkAndDownlink aper.Enumerated = 2
)

type LinksToLog struct {
	Value aper.Enumerated
}
