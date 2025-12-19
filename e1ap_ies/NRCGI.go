package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// NRCGI is a generated SEQUENCE type.
type NRCGI struct {
	PLMNIdentity   PLMNIdentity                `aper:"lb:3,ub:3,mandatory"`
	NRCellIdentity NRCellIdentity              `aper:"lb:36,ub:36,mandatory"`
	IEExtensions   *ProtocolExtensionContainer `aper:"optional"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *NRCGI) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(false); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(1), &aper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
		return fmt.Errorf("Encode optionality bitmap failed: %w", err)
	}
	if err = w.WriteOctetString([]byte(s.PLMNIdentity.Value), &aper.Constraint{Lb: 3, Ub: 3}, false); err != nil {
		return fmt.Errorf("Encode PLMNIdentity failed: %w", err)
	}
	if err = w.WriteBitString(s.NRCellIdentity.Value.Bytes, uint(s.NRCellIdentity.Value.NumBits), &aper.Constraint{Lb: 36, Ub: 36}, false); err != nil {
		return fmt.Errorf("Encode NRCellIdentity failed: %w", err)
	}
	if s.IEExtensions != nil {
		if err = s.IEExtensions.Encode(w); err != nil {
			return fmt.Errorf("Encode IEExtensions failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *NRCGI) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}
	if err = s.PLMNIdentity.Decode(r); err != nil {
		return fmt.Errorf("Decode PLMNIdentity failed: %w", err)
	}
	if err = s.NRCellIdentity.Decode(r); err != nil {
		return fmt.Errorf("Decode NRCellIdentity failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible { /* TODO: Implement extension skipping for NRCGI */
	}
	return nil
}
