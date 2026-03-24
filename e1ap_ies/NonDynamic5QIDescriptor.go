package e1ap_ies

import (
	"fmt"

	"asn1go/per"
)

// NonDynamic5QIDescriptor is a generated SEQUENCE type.
type NonDynamic5QIDescriptor struct {
	FiveQI             NonDynamic5QIDescriptorFiveQI
	QoSPriorityLevel   *QoSPriorityLevel
	AveragingWindow    *AveragingWindow
	MaxDataBurstVolume *MaxDataBurstVolume
	IEExtensions       *NonDynamic5QIDescriptorExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *NonDynamic5QIDescriptor) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "fiveQI", Optional: false},
			per.ComponentInfo{Name: "qoSPriorityLevel", Optional: true},
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

	if s.QoSPriorityLevel != nil {
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

	if err = w.EncodeInteger(int64(s.FiveQI.Value), per.ConstrainedExtensible(0, 255)); err != nil {
		return fmt.Errorf("encode FiveQI failed: %w", err)
	}

	if s.QoSPriorityLevel != nil {
		if err = w.EncodeInteger(int64((*s.QoSPriorityLevel).Value), per.ConstrainedExtensible(0, 127)); err != nil {
			return fmt.Errorf("encode QoSPriorityLevel failed: %w", err)
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
func (s *NonDynamic5QIDescriptor) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "fiveQI", Optional: false},
			per.ComponentInfo{Name: "qoSPriorityLevel", Optional: true},
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
		val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 255))
		if err != nil {
			return fmt.Errorf("decode FiveQI failed: %w", err)
		}
		s.FiveQI.Value = val
	}

	if seqDecoder.IsComponentPresent(1) {
		s.QoSPriorityLevel = new(QoSPriorityLevel)

		{
			val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 127))
			if err != nil {
				return fmt.Errorf("decode QoSPriorityLevel failed: %w", err)
			}
			s.QoSPriorityLevel.Value = val
		}
	}

	if seqDecoder.IsComponentPresent(2) {
		s.AveragingWindow = new(AveragingWindow)

		{
			val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 4095))
			if err != nil {
				return fmt.Errorf("decode AveragingWindow failed: %w", err)
			}
			s.AveragingWindow.Value = val
		}
	}

	if seqDecoder.IsComponentPresent(3) {
		s.MaxDataBurstVolume = new(MaxDataBurstVolume)

		{
			val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 4095))
			if err != nil {
				return fmt.Errorf("decode MaxDataBurstVolume failed: %w", err)
			}
			s.MaxDataBurstVolume.Value = val
		}
	}

	if seqDecoder.IsComponentPresent(4) {
		s.IEExtensions = new(NonDynamic5QIDescriptorExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
