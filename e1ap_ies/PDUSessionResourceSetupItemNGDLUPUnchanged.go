package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// PDUSessionResourceSetupItemNGDLUPUnchanged is a generated ENUMERATED type.
type PDUSessionResourceSetupItemNGDLUPUnchanged struct {
	Value aper.Enumerated
}

const (
	PDUSessionResourceSetupItemNGDLUPUnchangedTrue aper.Enumerated = 0
)

func (e *PDUSessionResourceSetupItemNGDLUPUnchanged) Encode(w *aper.AperWriter) error {
	// Encode logic for enum PDUSessionResourceSetupItemNGDLUPUnchanged to be generated here.
	return fmt.Errorf("Encode not implemented for enum PDUSessionResourceSetupItemNGDLUPUnchanged")
}

func (e *PDUSessionResourceSetupItemNGDLUPUnchanged) Decode(r *aper.AperReader) error {
	// Decode logic for enum PDUSessionResourceSetupItemNGDLUPUnchanged to be generated here.
	return fmt.Errorf("Decode not implemented for enum PDUSessionResourceSetupItemNGDLUPUnchanged")
}
