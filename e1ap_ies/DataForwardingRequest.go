package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// DataForwardingRequest is a generated ENUMERATED type.
type DataForwardingRequest struct {
	Value aper.Enumerated
}

const (
	DataForwardingRequestUL   aper.Enumerated = 0
	DataForwardingRequestDL   aper.Enumerated = 1
	DataForwardingRequestBoth aper.Enumerated = 2
)

func (e *DataForwardingRequest) Encode(w *aper.AperWriter) error {
	// Encode logic for enum DataForwardingRequest to be generated here.
	return fmt.Errorf("Encode not implemented for enum DataForwardingRequest")
}

func (e *DataForwardingRequest) Decode(r *aper.AperReader) error {
	// Decode logic for enum DataForwardingRequest to be generated here.
	return fmt.Errorf("Decode not implemented for enum DataForwardingRequest")
}
