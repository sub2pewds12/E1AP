package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// NonDynamic5QIDescriptor is a generated SEQUENCE type.
type NonDynamic5QIDescriptor struct {
	FiveQI             NonDynamic5QIDescriptorFiveQI      `aper:"lb:0,ub:255,mandatory,ext"`
	QoSPriorityLevel   *QoSPriorityLevel                  `aper:"lb:0,ub:127,optional,ext"`
	AveragingWindow    *AveragingWindow                   `aper:"lb:0,ub:4095,optional,ext"`
	MaxDataBurstVolume *MaxDataBurstVolume                `aper:"lb:0,ub:4095,optional,ext"`
	IEExtensions       *NonDynamic5QIDescriptorExtensions `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *NonDynamic5QIDescriptor) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.QoSPriorityLevel != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.AveragingWindow != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if s.MaxDataBurstVolume != nil {
		optionalityBitmap[0] |= 1 << 5
	}
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 4
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(4), &aper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
		return fmt.Errorf("Encode optionality bitmap failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.FiveQI.Value), &aper.Constraint{Lb: 0, Ub: 255}, true); err != nil {
		return fmt.Errorf("Encode FiveQI failed: %w", err)
	}
	if s.QoSPriorityLevel != nil {
		if err = w.WriteInteger(int64(s.QoSPriorityLevel.Value), &aper.Constraint{Lb: 0, Ub: 127}, true); err != nil {
			return fmt.Errorf("Encode QoSPriorityLevel failed: %w", err)
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
func (s *NonDynamic5QIDescriptor) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}
	if err = s.FiveQI.Decode(r); err != nil {
		return fmt.Errorf("Decode FiveQI failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.QoSPriorityLevel = new(QoSPriorityLevel)
		if err = s.QoSPriorityLevel.Decode(r); err != nil {
			return fmt.Errorf("Decode QoSPriorityLevel failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.AveragingWindow = new(AveragingWindow)
		if err = s.AveragingWindow.Decode(r); err != nil {
			return fmt.Errorf("Decode AveragingWindow failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<5) > 0 {
		s.MaxDataBurstVolume = new(MaxDataBurstVolume)
		if err = s.MaxDataBurstVolume.Decode(r); err != nil {
			return fmt.Errorf("Decode MaxDataBurstVolume failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<4) > 0 {
		s.IEExtensions = new(NonDynamic5QIDescriptorExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for NonDynamic5QIDescriptor")
	}
	return nil
}
