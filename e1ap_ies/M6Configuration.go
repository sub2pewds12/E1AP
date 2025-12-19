package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// M6Configuration is a generated SEQUENCE type.
type M6Configuration struct {
	M6reportInterval M6reportInterval            `aper:"mandatory,ext"`
	M6LinksToLog     LinksToLog                  `aper:"mandatory,ext"`
	IEExtensions     *ProtocolExtensionContainer `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *M6Configuration) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(1), &aper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
		return fmt.Errorf("Encode optionality bitmap failed: %w", err)
	}
	if err = w.WriteEnumerate(uint64(s.M6reportInterval.Value), aper.Constraint{Lb: 0, Ub: 13}, true); err != nil {
		return fmt.Errorf("Encode M6reportInterval failed: %w", err)
	}
	if err = w.WriteEnumerate(uint64(s.M6LinksToLog.Value), aper.Constraint{Lb: 0, Ub: 2}, true); err != nil {
		return fmt.Errorf("Encode M6LinksToLog failed: %w", err)
	}
	if s.IEExtensions != nil {
		if err = s.IEExtensions.Encode(w); err != nil {
			return fmt.Errorf("Encode IEExtensions failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *M6Configuration) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}
	if err = s.M6reportInterval.Decode(r); err != nil {
		return fmt.Errorf("Decode M6reportInterval failed: %w", err)
	}
	if err = s.M6LinksToLog.Decode(r); err != nil {
		return fmt.Errorf("Decode M6LinksToLog failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible { /* TODO: Implement extension skipping for M6Configuration */
	}
	return nil
}
