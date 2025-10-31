package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// CNSupport is a generated ENUMERATED type.
type CNSupport struct {
	Value aper.Enumerated
}

const (
	CNSupportCEpc aper.Enumerated = 0
	CNSupportC5gc aper.Enumerated = 1
	CNSupportBoth aper.Enumerated = 2
)

func (e *CNSupport) Encode(w *aper.AperWriter) error {
	// Encode logic for enum CNSupport to be generated here.
	return fmt.Errorf("Encode not implemented for enum CNSupport")
}

func (e *CNSupport) Decode(r *aper.AperReader) error {
	// Decode logic for enum CNSupport to be generated here.
	return fmt.Errorf("Decode not implemented for enum CNSupport")
}
