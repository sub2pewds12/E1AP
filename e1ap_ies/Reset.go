package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// Reset is a generated SEQUENCE type.
type Reset struct {
	TransactionID TransactionID `aper:"lb:0,ub:255,mandatory,ext"`
	Cause         Cause         `aper:"mandatory,ext"`
	ResetType     ResetType     `aper:"mandatory,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *Reset) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.TransactionID), &aper.Constraint{Lb: 0, Ub: 255}, true); err != nil {
		return fmt.Errorf("Encode TransactionID failed: %w", err)
	}
	if err = s.Cause.Encode(w); err != nil {
		return fmt.Errorf("Encode Cause failed: %w", err)
	}
	if err = s.ResetType.Encode(w); err != nil {
		return fmt.Errorf("Encode ResetType failed: %w", err)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *Reset) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}

	{
		var val int64
		if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 255}, true); err != nil {
			return fmt.Errorf("Decode TransactionID failed: %w", err)
		}
		s.TransactionID = TransactionID(val)
	}

	if err = s.Cause.Decode(r); err != nil {
		return fmt.Errorf("Decode Cause failed: %w", err)
	}
	if err = s.ResetType.Decode(r); err != nil {
		return fmt.Errorf("Decode ResetType failed: %w", err)
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for Reset")
	}
	return nil
}
