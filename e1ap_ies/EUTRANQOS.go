package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// EUTRANQOS is a generated SEQUENCE type.
type EUTRANQOS struct {
	QCI                                  QCI                                  `aper:"lb:0,ub:255,mandatory,ext"`
	EUTRANallocationAndRetentionPriority EUTRANAllocationAndRetentionPriority `aper:"mandatory,ext"`
	GbrQosInformation                    *GBRQosInformation                   `aper:"optional,ext"`
	IEExtensions                         *ProtocolExtensionContainer          `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *EUTRANQOS) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.GbrQosInformation != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(2), &aper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return fmt.Errorf("Encode optionality bitmap failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.QCI.Value), &aper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
		return fmt.Errorf("Encode QCI failed: %w", err)
	}
	if err = s.EUTRANallocationAndRetentionPriority.Encode(w); err != nil {
		return fmt.Errorf("Encode EUTRANallocationAndRetentionPriority failed: %w", err)
	}
	if s.GbrQosInformation != nil {
		if err = s.GbrQosInformation.Encode(w); err != nil {
			return fmt.Errorf("Encode GbrQosInformation failed: %w", err)
		}
	}
	if s.IEExtensions != nil {
		if err = s.IEExtensions.Encode(w); err != nil {
			return fmt.Errorf("Encode IEExtensions failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *EUTRANQOS) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}

	{
		var val int64
		if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
			return fmt.Errorf("Decode QCI failed: %w", err)
		}
		s.QCI.Value = aper.Integer(val)
	}
	if err = s.EUTRANallocationAndRetentionPriority.Decode(r); err != nil {
		return fmt.Errorf("Decode EUTRANallocationAndRetentionPriority failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.GbrQosInformation = new(GBRQosInformation)
		if err = s.GbrQosInformation.Decode(r); err != nil {
			return fmt.Errorf("Decode GbrQosInformation failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for EUTRANQOS")
	}
	return nil
}
