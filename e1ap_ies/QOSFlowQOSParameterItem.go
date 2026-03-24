package e1ap_ies

import (
	"fmt"

	"asn1go/per"
)

// QOSFlowQOSParameterItem is a generated SEQUENCE type.
type QOSFlowQOSParameterItem struct {
	QOSFlowIdentifier         QOSFlowIdentifier
	QoSFlowLevelQoSParameters QoSFlowLevelQoSParameters
	QoSFlowMappingIndication  *QOSFlowMappingIndication
	IEExtensions              *QOSFlowQOSParameterItemExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *QOSFlowQOSParameterItem) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "qoS-Flow-Identifier", Optional: false},
			per.ComponentInfo{Name: "qoSFlowLevelQoSParameters", Optional: false},
			per.ComponentInfo{Name: "qoSFlowMappingIndication", Optional: true},
			per.ComponentInfo{Name: "iE-Extensions", Optional: true},
		},
	}
	seqEncoder := w.NewSequenceEncoder(c)
	if err := seqEncoder.EncodeExtensionBit(false); err != nil {
		return err
	}

	optionalBitmap := make([]bool, 0)

	if s.QoSFlowMappingIndication != nil {
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

	if err = w.EncodeInteger(int64(s.QOSFlowIdentifier.Value), per.Constrained(0, 63)); err != nil {
		return fmt.Errorf("encode QOSFlowIdentifier failed: %w", err)
	}
	if err = s.QoSFlowLevelQoSParameters.Encode(w); err != nil {
		return fmt.Errorf("encode QoSFlowLevelQoSParameters failed: %w", err)
	}

	if s.QoSFlowMappingIndication != nil {

		{
			enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 2), ExtValues: nil}
			if err = w.EncodeEnumerated(int64((*s.QoSFlowMappingIndication).Value), enumC); err != nil {
				return fmt.Errorf("encode QoSFlowMappingIndication failed: %w", err)
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
func (s *QOSFlowQOSParameterItem) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "qoS-Flow-Identifier", Optional: false},
			per.ComponentInfo{Name: "qoSFlowLevelQoSParameters", Optional: false},
			per.ComponentInfo{Name: "qoSFlowMappingIndication", Optional: true},
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
		val, err := r.DecodeInteger(per.Constrained(0, 63))
		if err != nil {
			return fmt.Errorf("decode QOSFlowIdentifier failed: %w", err)
		}
		s.QOSFlowIdentifier.Value = val
	}
	if err = s.QoSFlowLevelQoSParameters.Decode(r); err != nil {
		return fmt.Errorf("Decode QoSFlowLevelQoSParameters failed: %w", err)
	}

	if seqDecoder.IsComponentPresent(2) {
		s.QoSFlowMappingIndication = new(QOSFlowMappingIndication)

		{
			enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 2), ExtValues: nil}
			val, err := r.DecodeEnumerated(enumC)
			if err != nil {
				return fmt.Errorf("decode QoSFlowMappingIndication failed: %w", err)
			}
			s.QoSFlowMappingIndication.Value = val
		}
	}

	if seqDecoder.IsComponentPresent(3) {
		s.IEExtensions = new(QOSFlowQOSParameterItemExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
