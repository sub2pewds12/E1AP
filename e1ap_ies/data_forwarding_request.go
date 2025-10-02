package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DataForwardingRequest From: 9_4_5_Information_Element_Definitions.txt:339
const (
	DataForwardingRequestUL   aper.Enumerated = 0
	DataForwardingRequestDL   aper.Enumerated = 1
	DataForwardingRequestBoth aper.Enumerated = 2
)

type DataForwardingRequest struct {
	Value aper.Enumerated
}
