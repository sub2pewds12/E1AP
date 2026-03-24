package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/asn1go/per"
)

// QOSMappingInformation is a generated SEQUENCE type.
type QOSMappingInformation struct {
	Dscp      *QOSMappingInformationDscp
	FlowLabel *QOSMappingInformationFlowLabel
}

// Encode implements the aper.AperMarshaller interface.
func (s *QOSMappingInformation) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "dscp", Optional: true},
			per.ComponentInfo{Name: "flow-label", Optional: true},
		},
	}
	seqEncoder := w.NewSequenceEncoder(c)
	if err := seqEncoder.EncodeExtensionBit(false); err != nil {
		return err
	}

	optionalBitmap := make([]bool, 0)

	if s.Dscp != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.FlowLabel != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if err := seqEncoder.EncodePreamble(optionalBitmap); err != nil {
		return err
	}

	if s.Dscp != nil {
		if err = w.EncodeBitString((*s.Dscp).Value, per.SizeConstraints{Extensible: false, Min: int64Ptr(6), Max: int64Ptr(6)}); err != nil {
			return fmt.Errorf("encode Dscp failed: %w", err)
		}
	}

	if s.FlowLabel != nil {
		if err = w.EncodeBitString((*s.FlowLabel).Value, per.SizeConstraints{Extensible: false, Min: int64Ptr(20), Max: int64Ptr(20)}); err != nil {
			return fmt.Errorf("encode FlowLabel failed: %w", err)
		}
	}

	if err := seqEncoder.EncodeExtensionAdditions([]bool{}, [][]byte{}); err != nil {
		return err
	}

	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *QOSMappingInformation) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "dscp", Optional: true},
			per.ComponentInfo{Name: "flow-label", Optional: true},
		},
	}
	seqDecoder := r.NewSequenceDecoder(c)
	if err := seqDecoder.DecodeExtensionBit(); err != nil {
		return err
	}

	if err := seqDecoder.DecodePreamble(); err != nil {
		return err
	}

	if seqDecoder.IsComponentPresent(0) {
		s.Dscp = new(QOSMappingInformationDscp)

		{
			val, err := r.DecodeBitString(per.SizeConstraints{Extensible: false, Min: int64Ptr(6), Max: int64Ptr(6)})
			if err != nil {
				return fmt.Errorf("decode Dscp failed: %w", err)
			}
			s.Dscp.Value = val
		}
	}

	if seqDecoder.IsComponentPresent(1) {
		s.FlowLabel = new(QOSMappingInformationFlowLabel)

		{
			val, err := r.DecodeBitString(per.SizeConstraints{Extensible: false, Min: int64Ptr(20), Max: int64Ptr(20)})
			if err != nil {
				return fmt.Errorf("decode FlowLabel failed: %w", err)
			}
			s.FlowLabel.Value = val
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
