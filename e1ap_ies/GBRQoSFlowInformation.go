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
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
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
		return fmt.Errorf("Encode optionality bitmap failed: %w", err)
	}
	if s.MaxFlowBitRateDownlink != nil {
		if err = w.WriteInteger(int64(s.MaxFlowBitRateDownlink.Value), &aper.Constraint{Lb: 0, Ub: 4000000000000}, true); err != nil {
			return fmt.Errorf("Encode MaxFlowBitRateDownlink failed: %w", err)
		}
	}
	if s.MaxFlowBitRateUplink != nil {
		if err = w.WriteInteger(int64(s.MaxFlowBitRateUplink.Value), &aper.Constraint{Lb: 0, Ub: 4000000000000}, true); err != nil {
			return fmt.Errorf("Encode MaxFlowBitRateUplink failed: %w", err)
		}
	}
	if s.GuaranteedFlowBitRateDownlink != nil {
		if err = w.WriteInteger(int64(s.GuaranteedFlowBitRateDownlink.Value), &aper.Constraint{Lb: 0, Ub: 4000000000000}, true); err != nil {
			return fmt.Errorf("Encode GuaranteedFlowBitRateDownlink failed: %w", err)
		}
	}
	if s.GuaranteedFlowBitRateUplink != nil {
		if err = w.WriteInteger(int64(s.GuaranteedFlowBitRateUplink.Value), &aper.Constraint{Lb: 0, Ub: 4000000000000}, true); err != nil {
			return fmt.Errorf("Encode GuaranteedFlowBitRateUplink failed: %w", err)
		}
	}
	if s.MaxPacketLossRateDownlink != nil {
		if err = w.WriteInteger(int64(s.MaxPacketLossRateDownlink.Value), &aper.Constraint{Lb: 0, Ub: 1000}, true); err != nil {
			return fmt.Errorf("Encode MaxPacketLossRateDownlink failed: %w", err)
		}
	}
	if s.MaxPacketLossRateUplink != nil {
		if err = w.WriteInteger(int64(s.MaxPacketLossRateUplink.Value), &aper.Constraint{Lb: 0, Ub: 1000}, true); err != nil {
			return fmt.Errorf("Encode MaxPacketLossRateUplink failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *GBRQoSFlowInformation) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 7, Ub: 7}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.MaxFlowBitRateDownlink = new(BitRate)
		if err = s.MaxFlowBitRateDownlink.Decode(r); err != nil {
			return fmt.Errorf("Decode MaxFlowBitRateDownlink failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.MaxFlowBitRateUplink = new(BitRate)
		if err = s.MaxFlowBitRateUplink.Decode(r); err != nil {
			return fmt.Errorf("Decode MaxFlowBitRateUplink failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<5) > 0 {
		s.GuaranteedFlowBitRateDownlink = new(BitRate)
		if err = s.GuaranteedFlowBitRateDownlink.Decode(r); err != nil {
			return fmt.Errorf("Decode GuaranteedFlowBitRateDownlink failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<4) > 0 {
		s.GuaranteedFlowBitRateUplink = new(BitRate)
		if err = s.GuaranteedFlowBitRateUplink.Decode(r); err != nil {
			return fmt.Errorf("Decode GuaranteedFlowBitRateUplink failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<3) > 0 {
		s.MaxPacketLossRateDownlink = new(MaxPacketLossRate)
		if err = s.MaxPacketLossRateDownlink.Decode(r); err != nil {
			return fmt.Errorf("Decode MaxPacketLossRateDownlink failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<2) > 0 {
		s.MaxPacketLossRateUplink = new(MaxPacketLossRate)
		if err = s.MaxPacketLossRateUplink.Decode(r); err != nil {
			return fmt.Errorf("Decode MaxPacketLossRateUplink failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<1) > 0 {
		s.IEExtensions = new(GBRQoSFlowInformationExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for GBRQoSFlowInformation")
	}
	return nil
}
