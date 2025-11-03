package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// Dynamic5QIDescriptor is a generated SEQUENCE type.
type Dynamic5QIDescriptor struct {
	QoSPriorityLevel   QoSPriorityLevel                   `aper:"lb:0,ub:127,mandatory,ext"`
	PacketDelayBudget  PacketDelayBudget                  `aper:"lb:0,ub:1023,mandatory,ext"`
	PacketErrorRate    PacketErrorRate                    `aper:"mandatory,ext"`
	FiveQI             *Dynamic5QIDescriptorFiveQI        `aper:"lb:0,ub:255,optional,ext"`
	DelayCritical      *Dynamic5QIDescriptorDelayCritical `aper:"optional,ext"`
	AveragingWindow    *AveragingWindow                   `aper:"lb:0,ub:4095,optional,ext"`
	MaxDataBurstVolume *MaxDataBurstVolume                `aper:"lb:0,ub:4095,optional,ext"`
	IEExtensions       *Dynamic5QIDescriptorExtensions    `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *Dynamic5QIDescriptor) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.FiveQI != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.DelayCritical != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if s.AveragingWindow != nil {
		optionalityBitmap[0] |= 1 << 5
	}
	if s.MaxDataBurstVolume != nil {
		optionalityBitmap[0] |= 1 << 4
	}
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 3
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(5), &aper.Constraint{Lb: 5, Ub: 5}, false); err != nil {
		return fmt.Errorf("Encode optionality bitmap failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.QoSPriorityLevel.Value), &aper.Constraint{Lb: 0, Ub: 127}, true); err != nil {
		return fmt.Errorf("Encode QoSPriorityLevel failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.PacketDelayBudget.Value), &aper.Constraint{Lb: 0, Ub: 1023}, true); err != nil {
		return fmt.Errorf("Encode PacketDelayBudget failed: %w", err)
	}
	if err = s.PacketErrorRate.Encode(w); err != nil {
		return fmt.Errorf("Encode PacketErrorRate failed: %w", err)
	}
	if s.FiveQI != nil {
		if err = w.WriteInteger(int64(s.FiveQI.Value), &aper.Constraint{Lb: 0, Ub: 255}, true); err != nil {
			return fmt.Errorf("Encode FiveQI failed: %w", err)
		}
	}
	if s.DelayCritical != nil {
		if err = w.WriteEnumerate(uint64(s.DelayCritical.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			return fmt.Errorf("Encode DelayCritical failed: %w", err)
		}
	}
	if s.AveragingWindow != nil {
		if err = w.WriteInteger(int64(s.AveragingWindow.Value), &aper.Constraint{Lb: 0, Ub: 4095}, true); err != nil {
			return fmt.Errorf("Encode AveragingWindow failed: %w", err)
		}
	}
	if s.MaxDataBurstVolume != nil {
		if err = w.WriteInteger(int64(s.MaxDataBurstVolume.Value), &aper.Constraint{Lb: 0, Ub: 4095}, true); err != nil {
			return fmt.Errorf("Encode MaxDataBurstVolume failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *Dynamic5QIDescriptor) Decode(r *aper.AperReader) (err error) {
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
		if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 127}, true); err != nil {
			return fmt.Errorf("Decode QoSPriorityLevel failed: %w", err)
		}
		s.QoSPriorityLevel.Value = aper.Integer(val)
	}

	{
		var val int64
		if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 1023}, true); err != nil {
			return fmt.Errorf("Decode PacketDelayBudget failed: %w", err)
		}
		s.PacketDelayBudget.Value = aper.Integer(val)
	}
	if err = s.PacketErrorRate.Decode(r); err != nil {
		return fmt.Errorf("Decode PacketErrorRate failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {

		{
			var val int64
			if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 255}, true); err != nil {
				return fmt.Errorf("Decode FiveQI failed: %w", err)
			}
			s.FiveQI = new(Dynamic5QIDescriptorFiveQI)
			s.FiveQI.Value = aper.Integer(val)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.DelayCritical = new(Dynamic5QIDescriptorDelayCritical)

		{
			var val uint64
			if val, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
				return fmt.Errorf("Decode DelayCritical failed: %w", err)
			}
			s.DelayCritical.Value = aper.Enumerated(val)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<5) > 0 {

		{
			var val int64
			if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4095}, true); err != nil {
				return fmt.Errorf("Decode AveragingWindow failed: %w", err)
			}
			s.AveragingWindow = new(AveragingWindow)
			s.AveragingWindow.Value = aper.Integer(val)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<4) > 0 {

		{
			var val int64
			if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4095}, true); err != nil {
				return fmt.Errorf("Decode MaxDataBurstVolume failed: %w", err)
			}
			s.MaxDataBurstVolume = new(MaxDataBurstVolume)
			s.MaxDataBurstVolume.Value = aper.Integer(val)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<3) > 0 {
		s.IEExtensions = new(Dynamic5QIDescriptorExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for Dynamic5QIDescriptor")
	}
	return nil
}
