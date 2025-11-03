package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// AlternativeQoSParaSetItem is a generated SEQUENCE type.
type AlternativeQoSParaSetItem struct {
	AlternativeQoSParameterIndex AlternativeQoSParaSetItemAlternativeQoSParameterIndex `aper:"lb:1,ub:8,mandatory,ext"`
	GuaranteedFlowBitRateDL      *BitRate                                              `aper:"lb:0,ub:4000000000000,optional,ext"`
	GuaranteedFlowBitRateUL      *BitRate                                              `aper:"lb:0,ub:4000000000000,optional,ext"`
	PacketDelayBudget            *PacketDelayBudget                                    `aper:"lb:0,ub:1023,optional,ext"`
	PacketErrorRate              *PacketErrorRate                                      `aper:"optional,ext"`
	IEExtensions                 *ProtocolExtensionContainer                           `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *AlternativeQoSParaSetItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.GuaranteedFlowBitRateDL != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.GuaranteedFlowBitRateUL != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if s.PacketDelayBudget != nil {
		optionalityBitmap[0] |= 1 << 5
	}
	if s.PacketErrorRate != nil {
		optionalityBitmap[0] |= 1 << 4
	}
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 3
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(5), &aper.Constraint{Lb: 5, Ub: 5}, false); err != nil {
		return fmt.Errorf("Encode optionality bitmap failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.AlternativeQoSParameterIndex.Value), &aper.Constraint{Lb: 1, Ub: 8}, true); err != nil {
		return fmt.Errorf("Encode AlternativeQoSParameterIndex failed: %w", err)
	}
	if s.GuaranteedFlowBitRateDL != nil {
		if err = w.WriteInteger(int64(s.GuaranteedFlowBitRateDL.Value), &aper.Constraint{Lb: 0, Ub: 4000000000000}, true); err != nil {
			return fmt.Errorf("Encode GuaranteedFlowBitRateDL failed: %w", err)
		}
	}
	if s.GuaranteedFlowBitRateUL != nil {
		if err = w.WriteInteger(int64(s.GuaranteedFlowBitRateUL.Value), &aper.Constraint{Lb: 0, Ub: 4000000000000}, true); err != nil {
			return fmt.Errorf("Encode GuaranteedFlowBitRateUL failed: %w", err)
		}
	}
	if s.PacketDelayBudget != nil {
		if err = w.WriteInteger(int64(s.PacketDelayBudget.Value), &aper.Constraint{Lb: 0, Ub: 1023}, true); err != nil {
			return fmt.Errorf("Encode PacketDelayBudget failed: %w", err)
		}
	}
	if s.PacketErrorRate != nil {
		if err = s.PacketErrorRate.Encode(w); err != nil {
			return fmt.Errorf("Encode PacketErrorRate failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *AlternativeQoSParaSetItem) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 5, Ub: 5}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}

	{
		var val int64
		if val, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 8}, true); err != nil {
			return fmt.Errorf("Decode AlternativeQoSParameterIndex failed: %w", err)
		}
		s.AlternativeQoSParameterIndex.Value = aper.Integer(val)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {

		{
			var val int64
			if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4000000000000}, true); err != nil {
				return fmt.Errorf("Decode GuaranteedFlowBitRateDL failed: %w", err)
			}
			s.GuaranteedFlowBitRateDL = new(BitRate)
			s.GuaranteedFlowBitRateDL.Value = aper.Integer(val)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {

		{
			var val int64
			if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4000000000000}, true); err != nil {
				return fmt.Errorf("Decode GuaranteedFlowBitRateUL failed: %w", err)
			}
			s.GuaranteedFlowBitRateUL = new(BitRate)
			s.GuaranteedFlowBitRateUL.Value = aper.Integer(val)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<5) > 0 {

		{
			var val int64
			if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 1023}, true); err != nil {
				return fmt.Errorf("Decode PacketDelayBudget failed: %w", err)
			}
			s.PacketDelayBudget = new(PacketDelayBudget)
			s.PacketDelayBudget.Value = aper.Integer(val)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<4) > 0 {
		s.PacketErrorRate = new(PacketErrorRate)
		if err = s.PacketErrorRate.Decode(r); err != nil {
			return fmt.Errorf("Decode PacketErrorRate failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<3) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for AlternativeQoSParaSetItem")
	}
	return nil
}
