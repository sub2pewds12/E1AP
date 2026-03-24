package e1ap_ies

import (
	"fmt"

	"asn1go/per"
)

// QoSFlowLevelQoSParameters is a generated SEQUENCE type.
type QoSFlowLevelQoSParameters struct {
	QOSCharacteristics               QOSCharacteristics
	NGRANallocationRetentionPriority NGRANAllocationAndRetentionPriority
	GBRQOSFlowInformation            *GBRQoSFlowInformation
	ReflectiveQOSAttribute           *QoSFlowLevelQoSParametersReflectiveQOSAttribute
	AdditionalQOSInformation         *QoSFlowLevelQoSParametersAdditionalQOSInformation
	PagingPolicyIndicator            *QoSFlowLevelQoSParametersPagingPolicyIndicator
	ReflectiveQOSIndicator           *QoSFlowLevelQoSParametersReflectiveQOSIndicator
	IEExtensions                     *QoSFlowLevelQoSParametersExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *QoSFlowLevelQoSParameters) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "qoS-Characteristics", Optional: false},
			per.ComponentInfo{Name: "nGRANallocationRetentionPriority", Optional: false},
			per.ComponentInfo{Name: "gBR-QoS-Flow-Information", Optional: true},
			per.ComponentInfo{Name: "reflective-QoS-Attribute", Optional: true},
			per.ComponentInfo{Name: "additional-QoS-Information", Optional: true},
			per.ComponentInfo{Name: "paging-Policy-Indicator", Optional: true},
			per.ComponentInfo{Name: "reflective-QoS-Indicator", Optional: true},
			per.ComponentInfo{Name: "iE-Extensions", Optional: true},
		},
	}
	seqEncoder := w.NewSequenceEncoder(c)
	if err := seqEncoder.EncodeExtensionBit(false); err != nil {
		return err
	}

	optionalBitmap := make([]bool, 0)

	if s.GBRQOSFlowInformation != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.ReflectiveQOSAttribute != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.AdditionalQOSInformation != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.PagingPolicyIndicator != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.ReflectiveQOSIndicator != nil {
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

	if err = s.QOSCharacteristics.Encode(w); err != nil {
		return fmt.Errorf("encode QOSCharacteristics failed: %w", err)
	}
	if err = s.NGRANallocationRetentionPriority.Encode(w); err != nil {
		return fmt.Errorf("encode NGRANallocationRetentionPriority failed: %w", err)
	}

	if s.GBRQOSFlowInformation != nil {
		if err = s.GBRQOSFlowInformation.Encode(w); err != nil {
			return fmt.Errorf("encode GBRQOSFlowInformation failed: %w", err)
		}
	}

	if s.ReflectiveQOSAttribute != nil {

		{
			enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 1), ExtValues: nil}
			if err = w.EncodeEnumerated(int64((*s.ReflectiveQOSAttribute).Value), enumC); err != nil {
				return fmt.Errorf("encode ReflectiveQOSAttribute failed: %w", err)
			}
		}
	}

	if s.AdditionalQOSInformation != nil {

		{
			enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 1), ExtValues: nil}
			if err = w.EncodeEnumerated(int64((*s.AdditionalQOSInformation).Value), enumC); err != nil {
				return fmt.Errorf("encode AdditionalQOSInformation failed: %w", err)
			}
		}
	}

	if s.PagingPolicyIndicator != nil {
		if err = w.EncodeInteger(int64((*s.PagingPolicyIndicator).Value), per.ConstrainedExtensible(1, 8)); err != nil {
			return fmt.Errorf("encode PagingPolicyIndicator failed: %w", err)
		}
	}

	if s.ReflectiveQOSIndicator != nil {

		{
			enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 1), ExtValues: nil}
			if err = w.EncodeEnumerated(int64((*s.ReflectiveQOSIndicator).Value), enumC); err != nil {
				return fmt.Errorf("encode ReflectiveQOSIndicator failed: %w", err)
			}
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
func (s *QoSFlowLevelQoSParameters) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "qoS-Characteristics", Optional: false},
			per.ComponentInfo{Name: "nGRANallocationRetentionPriority", Optional: false},
			per.ComponentInfo{Name: "gBR-QoS-Flow-Information", Optional: true},
			per.ComponentInfo{Name: "reflective-QoS-Attribute", Optional: true},
			per.ComponentInfo{Name: "additional-QoS-Information", Optional: true},
			per.ComponentInfo{Name: "paging-Policy-Indicator", Optional: true},
			per.ComponentInfo{Name: "reflective-QoS-Indicator", Optional: true},
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

	if err = s.QOSCharacteristics.Decode(r); err != nil {
		return fmt.Errorf("Decode QOSCharacteristics failed: %w", err)
	}
	if err = s.NGRANallocationRetentionPriority.Decode(r); err != nil {
		return fmt.Errorf("Decode NGRANallocationRetentionPriority failed: %w", err)
	}

	if seqDecoder.IsComponentPresent(2) {
		s.GBRQOSFlowInformation = new(GBRQoSFlowInformation)
		if err = s.GBRQOSFlowInformation.Decode(r); err != nil {
			return fmt.Errorf("Decode GBRQOSFlowInformation failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(3) {
		s.ReflectiveQOSAttribute = new(QoSFlowLevelQoSParametersReflectiveQOSAttribute)

		{
			enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 1), ExtValues: nil}
			val, err := r.DecodeEnumerated(enumC)
			if err != nil {
				return fmt.Errorf("decode ReflectiveQOSAttribute failed: %w", err)
			}
			s.ReflectiveQOSAttribute.Value = val
		}
	}

	if seqDecoder.IsComponentPresent(4) {
		s.AdditionalQOSInformation = new(QoSFlowLevelQoSParametersAdditionalQOSInformation)

		{
			enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 1), ExtValues: nil}
			val, err := r.DecodeEnumerated(enumC)
			if err != nil {
				return fmt.Errorf("decode AdditionalQOSInformation failed: %w", err)
			}
			s.AdditionalQOSInformation.Value = val
		}
	}

	if seqDecoder.IsComponentPresent(5) {
		s.PagingPolicyIndicator = new(QoSFlowLevelQoSParametersPagingPolicyIndicator)

		{
			val, err := r.DecodeInteger(per.ConstrainedExtensible(1, 8))
			if err != nil {
				return fmt.Errorf("decode PagingPolicyIndicator failed: %w", err)
			}
			s.PagingPolicyIndicator.Value = val
		}
	}

	if seqDecoder.IsComponentPresent(6) {
		s.ReflectiveQOSIndicator = new(QoSFlowLevelQoSParametersReflectiveQOSIndicator)

		{
			enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 1), ExtValues: nil}
			val, err := r.DecodeEnumerated(enumC)
			if err != nil {
				return fmt.Errorf("decode ReflectiveQOSIndicator failed: %w", err)
			}
			s.ReflectiveQOSIndicator.Value = val
		}
	}

	if seqDecoder.IsComponentPresent(7) {
		s.IEExtensions = new(QoSFlowLevelQoSParametersExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
