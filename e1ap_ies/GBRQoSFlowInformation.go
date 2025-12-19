package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// GBRQoSFlowInformation is a generated SEQUENCE type.
type GBRQoSFlowInformation struct {
	MaxFlowBitRateDownlink        *BitRate                         `aper:"lb:0,ub:4000000000000,optional,ext"`
	MaxFlowBitRateUplink          *BitRate                         `aper:"lb:0,ub:4000000000000,optional,ext"`
	GuaranteedFlowBitRateDownlink *BitRate                         `aper:"lb:0,ub:4000000000000,optional,ext"`
	GuaranteedFlowBitRateUplink   *BitRate                         `aper:"lb:0,ub:4000000000000,optional,ext"`
	MaxPacketLossRateDownlink     *MaxPacketLossRate               `aper:"lb:0,ub:1000,optional,ext"`
	MaxPacketLossRateUplink       *MaxPacketLossRate               `aper:"lb:0,ub:1000,optional,ext"`
	IEExtensions                  *GBRQoSFlowInformationExtensions `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *GBRQoSFlowInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.MaxFlowBitRateDownlink != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.MaxFlowBitRateUplink != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if s.GuaranteedFlowBitRateDownlink != nil {
		optionalityBitmap[0] |= 1 << 5
	}
	if s.GuaranteedFlowBitRateUplink != nil {
		optionalityBitmap[0] |= 1 << 4
	}
	if s.MaxPacketLossRateDownlink != nil {
		optionalityBitmap[0] |= 1 << 3
	}
	if s.MaxPacketLossRateUplink != nil {
		optionalityBitmap[0] |= 1 << 2
	}
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 1
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(7), &aper.Constraint{Lb: 7, Ub: 7}, false); err != nil {
		return fmt.Errorf("encode optionality bitmap failed: %w", err)
	}
	if s.MaxFlowBitRateDownlink != nil {
		if err = w.WriteInteger(int64((*s.MaxFlowBitRateDownlink).Value), &aper.Constraint{Lb: 0, Ub: 4000000000000}, true); err != nil {
			return fmt.Errorf("encode MaxFlowBitRateDownlink failed: %w", err)
		}
	}
	if s.MaxFlowBitRateUplink != nil {
		if err = w.WriteInteger(int64((*s.MaxFlowBitRateUplink).Value), &aper.Constraint{Lb: 0, Ub: 4000000000000}, true); err != nil {
			return fmt.Errorf("encode MaxFlowBitRateUplink failed: %w", err)
		}
	}
	if s.GuaranteedFlowBitRateDownlink != nil {
		if err = w.WriteInteger(int64((*s.GuaranteedFlowBitRateDownlink).Value), &aper.Constraint{Lb: 0, Ub: 4000000000000}, true); err != nil {
			return fmt.Errorf("encode GuaranteedFlowBitRateDownlink failed: %w", err)
		}
	}
	if s.GuaranteedFlowBitRateUplink != nil {
		if err = w.WriteInteger(int64((*s.GuaranteedFlowBitRateUplink).Value), &aper.Constraint{Lb: 0, Ub: 4000000000000}, true); err != nil {
			return fmt.Errorf("encode GuaranteedFlowBitRateUplink failed: %w", err)
		}
	}
	if s.MaxPacketLossRateDownlink != nil {
		if err = w.WriteInteger(int64((*s.MaxPacketLossRateDownlink).Value), &aper.Constraint{Lb: 0, Ub: 1000}, true); err != nil {
			return fmt.Errorf("encode MaxPacketLossRateDownlink failed: %w", err)
		}
	}
	if s.MaxPacketLossRateUplink != nil {
		if err = w.WriteInteger(int64((*s.MaxPacketLossRateUplink).Value), &aper.Constraint{Lb: 0, Ub: 1000}, true); err != nil {
			return fmt.Errorf("encode MaxPacketLossRateUplink failed: %w", err)
		}
	}
	if s.IEExtensions != nil {
		if err = s.IEExtensions.Encode(w); err != nil {
			return fmt.Errorf("encode IEExtensions failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *GBRQoSFlowInformation) Decode(r *aper.AperReader) (err error) {
	isExtensible, err := r.ReadBool()
	if err != nil {
		return fmt.Errorf("read extensibility bool failed: %w", err)
	}
	_ = isExtensible
	optionalityBitmap, _, err := r.ReadBitString(&aper.Constraint{Lb: 7, Ub: 7}, false)
	if err != nil {
		return fmt.Errorf("read optionality bitmap failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.MaxFlowBitRateDownlink = new(BitRate)
		if err = s.MaxFlowBitRateDownlink.Decode(r); err != nil {
			return fmt.Errorf("decode MaxFlowBitRateDownlink failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.MaxFlowBitRateUplink = new(BitRate)
		if err = s.MaxFlowBitRateUplink.Decode(r); err != nil {
			return fmt.Errorf("decode MaxFlowBitRateUplink failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<5) > 0 {
		s.GuaranteedFlowBitRateDownlink = new(BitRate)
		if err = s.GuaranteedFlowBitRateDownlink.Decode(r); err != nil {
			return fmt.Errorf("decode GuaranteedFlowBitRateDownlink failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<4) > 0 {
		s.GuaranteedFlowBitRateUplink = new(BitRate)
		if err = s.GuaranteedFlowBitRateUplink.Decode(r); err != nil {
			return fmt.Errorf("decode GuaranteedFlowBitRateUplink failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<3) > 0 {
		s.MaxPacketLossRateDownlink = new(MaxPacketLossRate)
		if err = s.MaxPacketLossRateDownlink.Decode(r); err != nil {
			return fmt.Errorf("decode MaxPacketLossRateDownlink failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<2) > 0 {
		s.MaxPacketLossRateUplink = new(MaxPacketLossRate)
		if err = s.MaxPacketLossRateUplink.Decode(r); err != nil {
			return fmt.Errorf("decode MaxPacketLossRateUplink failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<1) > 0 {
		s.IEExtensions = new(GBRQoSFlowInformationExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible { /* TODO: Implement extension skipping for GBRQoSFlowInformation */
	}
	return nil
}
