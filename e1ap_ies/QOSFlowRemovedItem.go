package e1ap_ies

import (
	"fmt"

	"asn1go/per"
)

// QOSFlowRemovedItem is a generated SEQUENCE type.
type QOSFlowRemovedItem struct {
	QOSFlowIdentifier             QOSFlowIdentifier
	QOSFlowReleasedInSession      *QOSFlowRemovedItemQOSFlowReleasedInSession
	QOSFlowAccumulatedSessionTime *QOSFlowRemovedItemQOSFlowAccumulatedSessionTime
	IEExtensions                  *QOSFlowRemovedItemExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *QOSFlowRemovedItem) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "qoS-Flow-Identifier", Optional: false},
			per.ComponentInfo{Name: "qoS-Flow-Released-In-Session", Optional: true},
			per.ComponentInfo{Name: "qoS-Flow-Accumulated-Session-Time", Optional: true},
			per.ComponentInfo{Name: "iE-Extensions", Optional: true},
		},
	}
	seqEncoder := w.NewSequenceEncoder(c)
	if err := seqEncoder.EncodeExtensionBit(false); err != nil {
		return err
	}

	optionalBitmap := make([]bool, 0)

	if s.QOSFlowReleasedInSession != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.QOSFlowAccumulatedSessionTime != nil {
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

	if s.QOSFlowReleasedInSession != nil {

		{
			enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 2), ExtValues: nil}
			if err = w.EncodeEnumerated(int64((*s.QOSFlowReleasedInSession).Value), enumC); err != nil {
				return fmt.Errorf("encode QOSFlowReleasedInSession failed: %w", err)
			}
		}
	}

	if s.QOSFlowAccumulatedSessionTime != nil {
		if err = w.EncodeOctetString([]byte((*s.QOSFlowAccumulatedSessionTime).Value), per.SizeConstraints{Extensible: false, Min: int64Ptr(5), Max: int64Ptr(5)}); err != nil {
			return fmt.Errorf("encode QOSFlowAccumulatedSessionTime failed: %w", err)
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
func (s *QOSFlowRemovedItem) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "qoS-Flow-Identifier", Optional: false},
			per.ComponentInfo{Name: "qoS-Flow-Released-In-Session", Optional: true},
			per.ComponentInfo{Name: "qoS-Flow-Accumulated-Session-Time", Optional: true},
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

	if seqDecoder.IsComponentPresent(1) {
		s.QOSFlowReleasedInSession = new(QOSFlowRemovedItemQOSFlowReleasedInSession)

		{
			enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 2), ExtValues: nil}
			val, err := r.DecodeEnumerated(enumC)
			if err != nil {
				return fmt.Errorf("decode QOSFlowReleasedInSession failed: %w", err)
			}
			s.QOSFlowReleasedInSession.Value = val
		}
	}

	if seqDecoder.IsComponentPresent(2) {
		s.QOSFlowAccumulatedSessionTime = new(QOSFlowRemovedItemQOSFlowAccumulatedSessionTime)

		{
			val, err := r.DecodeOctetString(per.SizeConstraints{Extensible: false, Min: int64Ptr(5), Max: int64Ptr(5)})
			if err != nil {
				return fmt.Errorf("decode QOSFlowAccumulatedSessionTime failed: %w", err)
			}
			s.QOSFlowAccumulatedSessionTime.Value = val
		}
	}

	if seqDecoder.IsComponentPresent(3) {
		s.IEExtensions = new(QOSFlowRemovedItemExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
