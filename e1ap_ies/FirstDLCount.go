package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// FirstDLCount is a generated SEQUENCE type.
type FirstDLCount struct {
	FirstDLCountVal PDCPCount                   `aper:"mandatory"`
	IEExtensions    *ProtocolExtensionContainer `aper:"optional"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *FirstDLCount) Encode(w *aper.AperWriter) (err error) {
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
	if err = s.FirstDLCountVal.Encode(w); err != nil {
		return fmt.Errorf("Encode FirstDLCountVal failed: %w", err)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *FirstDLCount) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}
	if err = s.FirstDLCountVal.Decode(r); err != nil {
		return fmt.Errorf("Decode FirstDLCountVal failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible { /* TODO: Implement extension skipping for FirstDLCount */
	}
	return nil
}
