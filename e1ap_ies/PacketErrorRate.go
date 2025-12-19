package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// PacketErrorRate is a generated SEQUENCE type.
type PacketErrorRate struct {
	PERScalar    PERScalar                   `aper:"lb:0,ub:9,mandatory,ext"`
	PERExponent  PERExponent                 `aper:"lb:0,ub:9,mandatory,ext"`
	IEExtensions *ProtocolExtensionContainer `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *PacketErrorRate) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(1), &aper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
		return fmt.Errorf("encode optionality bitmap failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.PERScalar.Value), &aper.Constraint{Lb: 0, Ub: 9}, true); err != nil {
		return fmt.Errorf("encode PERScalar failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.PERExponent.Value), &aper.Constraint{Lb: 0, Ub: 9}, true); err != nil {
		return fmt.Errorf("encode PERExponent failed: %w", err)
	}
	if s.IEExtensions != nil {
		if err = s.IEExtensions.Encode(w); err != nil {
			return fmt.Errorf("encode IEExtensions failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *PacketErrorRate) Decode(r *aper.AperReader) (err error) {
	isExtensible, err := r.ReadBool()
	if err != nil {
		return fmt.Errorf("read extensibility bool failed: %w", err)
	}
	_ = isExtensible
	optionalityBitmap, _, err := r.ReadBitString(&aper.Constraint{Lb: 1, Ub: 1}, false)
	if err != nil {
		return fmt.Errorf("read optionality bitmap failed: %w", err)
	}
	if err = s.PERScalar.Decode(r); err != nil {
		return fmt.Errorf("decode PERScalar failed: %w", err)
	}
	if err = s.PERExponent.Decode(r); err != nil {
		return fmt.Errorf("decode PERExponent failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible { /* TODO: Implement extension skipping for PacketErrorRate */
	}
	return nil
}
