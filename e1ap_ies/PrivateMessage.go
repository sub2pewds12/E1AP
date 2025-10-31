package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// PrivateMessage is a generated SEQUENCE type.
type PrivateMessage struct {
	PrivateIEs []PrivateIEField `aper:"ub:MaxPrivateIEs,mandatory,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *PrivateMessage) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	if err = s.PrivateIEs.Encode(w); err != nil {
		return fmt.Errorf("Encode PrivateIEs failed: %w", err)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *PrivateMessage) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	if err = s.PrivateIEs.Decode(r); err != nil {
		return fmt.Errorf("Decode PrivateIEs failed: %w", err)
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for PrivateMessage")
	}
	return nil
}
