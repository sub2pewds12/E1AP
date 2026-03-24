package e1ap_ies

import (
	"fmt"

	"asn1go/per"
)

// Dynamic5QIDescriptor is a generated SEQUENCE type.
type Dynamic5QIDescriptor struct {
	QoSPriorityLevel   QoSPriorityLevel
	PacketDelayBudget  PacketDelayBudget
	PacketErrorRate    PacketErrorRate
	FiveQI             *Dynamic5QIDescriptorFiveQI
	DelayCritical      *Dynamic5QIDescriptorDelayCritical
	AveragingWindow    *AveragingWindow
	MaxDataBurstVolume *MaxDataBurstVolume
	IEExtensions       *Dynamic5QIDescriptorExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *Dynamic5QIDescriptor) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "qoSPriorityLevel", Optional: false},
			per.ComponentInfo{Name: "packetDelayBudget", Optional: false},
			per.ComponentInfo{Name: "packetErrorRate", Optional: false},
			per.ComponentInfo{Name: "fiveQI", Optional: true},
			per.ComponentInfo{Name: "delayCritical", Optional: true},
			per.ComponentInfo{Name: "averagingWindow", Optional: true},
			per.ComponentInfo{Name: "maxDataBurstVolume", Optional: true},
			per.ComponentInfo{Name: "iE-Extensions", Optional: true},
		},
	}
	seqEncoder := w.NewSequenceEncoder(c)
	if err := seqEncoder.EncodeExtensionBit(false); err != nil {
		return err
	}

	optionalBitmap := make([]bool, 0)

	if s.FiveQI != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.DelayCritical != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.AveragingWindow != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.MaxDataBurstVolume != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.IEExtensions != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if err := seqEncoder.EncodePreamble(optionalBitmap); err != nil {
		return err
	}

	if err = w.EncodeInteger(int64(s.QoSPriorityLevel.Value), per.ConstrainedExtensible(0, 127)); err != nil {
		return fmt.Errorf("encode QoSPriorityLevel failed: %w", err)
	}
	if err = w.EncodeInteger(int64(s.PacketDelayBudget.Value), per.ConstrainedExtensible(0, 1023)); err != nil {
		return fmt.Errorf("encode PacketDelayBudget failed: %w", err)
	}
	if err = s.PacketErrorRate.Encode(w); err != nil {
		return fmt.Errorf("encode PacketErrorRate failed: %w", err)
	}

	if s.FiveQI != nil {
		if err = w.EncodeInteger(int64((*s.FiveQI).Value), per.ConstrainedExtensible(0, 255)); err != nil {
			return fmt.Errorf("encode FiveQI failed: %w", err)
		}
	}

	if s.DelayCritical != nil {

		{
			enumC := per.EnumeratedConstraints{Extensible: false, RootValues: make([]int64, 2), ExtValues: nil}
			if err = w.EncodeEnumerated(int64((*s.DelayCritical).Value), enumC); err != nil {
				return fmt.Errorf("encode DelayCritical failed: %w", err)
			}
		}
	}

	if s.AveragingWindow != nil {
		if err = w.EncodeInteger(int64((*s.AveragingWindow).Value), per.ConstrainedExtensible(0, 4095)); err != nil {
			return fmt.Errorf("encode AveragingWindow failed: %w", err)
		}
	}

	if s.MaxDataBurstVolume != nil {
		if err = w.EncodeInteger(int64((*s.MaxDataBurstVolume).Value), per.ConstrainedExtensible(0, 4095)); err != nil {
			return fmt.Errorf("encode MaxDataBurstVolume failed: %w", err)
		}
	}

	if s.IEExtensions != nil {
		if err = s.IEExtensions.Encode(w); err != nil {
			return fmt.Errorf("encode IEExtensions failed: %w", err)
		}
	}

	if err := seqEncoder.EncodeExtensionAdditions([]bool{}, [][]byte{}); err != nil {
		return err
	}

	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *Dynamic5QIDescriptor) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "qoSPriorityLevel", Optional: false},
			per.ComponentInfo{Name: "packetDelayBudget", Optional: false},
			per.ComponentInfo{Name: "packetErrorRate", Optional: false},
			per.ComponentInfo{Name: "fiveQI", Optional: true},
			per.ComponentInfo{Name: "delayCritical", Optional: true},
			per.ComponentInfo{Name: "averagingWindow", Optional: true},
			per.ComponentInfo{Name: "maxDataBurstVolume", Optional: true},
			per.ComponentInfo{Name: "iE-Extensions", Optional: true},
		},
	}
	seqDecoder := r.NewSequenceDecoder(c)
	if err := seqDecoder.DecodeExtensionBit(); err != nil {
		return err
	}

	if err := seqDecoder.DecodePreamble(); err != nil {
		return err
	}

	{
		val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 127))
		if err != nil {
			return fmt.Errorf("decode QoSPriorityLevel failed: %w", err)
		}
		s.QoSPriorityLevel.Value = val
	}

	{
		val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 1023))
		if err != nil {
			return fmt.Errorf("decode PacketDelayBudget failed: %w", err)
		}
		s.PacketDelayBudget.Value = val
	}
	if err = s.PacketErrorRate.Decode(r); err != nil {
		return fmt.Errorf("Decode PacketErrorRate failed: %w", err)
	}

	if seqDecoder.IsComponentPresent(3) {
		s.FiveQI = new(Dynamic5QIDescriptorFiveQI)

		{
			val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 255))
			if err != nil {
				return fmt.Errorf("decode FiveQI failed: %w", err)
			}
			s.FiveQI.Value = val
		}
	}

	if seqDecoder.IsComponentPresent(4) {
		s.DelayCritical = new(Dynamic5QIDescriptorDelayCritical)

		{
			enumC := per.EnumeratedConstraints{Extensible: false, RootValues: make([]int64, 2), ExtValues: nil}
			val, err := r.DecodeEnumerated(enumC)
			if err != nil {
				return fmt.Errorf("decode DelayCritical failed: %w", err)
			}
			s.DelayCritical.Value = val
		}
	}

	if seqDecoder.IsComponentPresent(5) {
		s.AveragingWindow = new(AveragingWindow)

		{
			val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 4095))
			if err != nil {
				return fmt.Errorf("decode AveragingWindow failed: %w", err)
			}
			s.AveragingWindow.Value = val
		}
	}

	if seqDecoder.IsComponentPresent(6) {
		s.MaxDataBurstVolume = new(MaxDataBurstVolume)

		{
			val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 4095))
			if err != nil {
				return fmt.Errorf("decode MaxDataBurstVolume failed: %w", err)
			}
			s.MaxDataBurstVolume.Value = val
		}
	}

	if seqDecoder.IsComponentPresent(7) {
		s.IEExtensions = new(Dynamic5QIDescriptorExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
