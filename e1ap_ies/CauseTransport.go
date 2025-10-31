package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// CauseTransport is a generated ENUMERATED type.
type CauseTransport struct {
	Value aper.Enumerated
}

const (
	CauseTransportUnspecified                  aper.Enumerated = 0
	CauseTransportTransportResourceUnavailable aper.Enumerated = 1
	CauseTransportUnknownTNLAddressForIAB      aper.Enumerated = 2
)

func (e *CauseTransport) Encode(w *aper.AperWriter) error {
	// Encode logic for enum CauseTransport to be generated here.
	return fmt.Errorf("Encode not implemented for enum CauseTransport")
}

func (e *CauseTransport) Decode(r *aper.AperReader) error {
	// Decode logic for enum CauseTransport to be generated here.
	return fmt.Errorf("Decode not implemented for enum CauseTransport")
}
