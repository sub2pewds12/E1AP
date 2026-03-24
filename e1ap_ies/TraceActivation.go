package e1ap_ies

import (
	"fmt"

	"asn1go/per"
)

// TraceActivation is a generated SEQUENCE type.
type TraceActivation struct {
	TraceID                        TraceID
	InterfacesToTrace              InterfacesToTrace
	TraceDepth                     TraceDepth
	TraceCollectionEntityIPAddress TransportLayerAddress
	IEExtensions                   *TraceActivationExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *TraceActivation) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "traceID", Optional: false},
			per.ComponentInfo{Name: "interfacesToTrace", Optional: false},
			per.ComponentInfo{Name: "traceDepth", Optional: false},
			per.ComponentInfo{Name: "traceCollectionEntityIPAddress", Optional: false},
			per.ComponentInfo{Name: "iE-Extensions", Optional: true},
		},
	}
	seqEncoder := w.NewSequenceEncoder(c)
	if err := seqEncoder.EncodeExtensionBit(false); err != nil {
		return err
	}

	optionalBitmap := make([]bool, 0)

	if s.IEExtensions != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if err := seqEncoder.EncodePreamble(optionalBitmap); err != nil {
		return err
	}

	if err = w.EncodeOctetString([]byte(s.TraceID.Value), per.SizeConstraints{Extensible: false, Min: int64Ptr(8), Max: int64Ptr(8)}); err != nil {
		return fmt.Errorf("encode TraceID failed: %w", err)
	}
	if err = w.EncodeBitString(s.InterfacesToTrace.Value, per.SizeConstraints{Extensible: false, Min: int64Ptr(8), Max: int64Ptr(8)}); err != nil {
		return fmt.Errorf("encode InterfacesToTrace failed: %w", err)
	}

	{
		enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 6), ExtValues: nil}
		if err = w.EncodeEnumerated(int64(s.TraceDepth.Value), enumC); err != nil {
			return fmt.Errorf("encode TraceDepth failed: %w", err)
		}
	}
	if err = w.EncodeBitString(s.TraceCollectionEntityIPAddress.Value, per.SizeConstraints{Extensible: false, Min: int64Ptr(1), Max: int64Ptr(160)}); err != nil {
		return fmt.Errorf("encode TraceCollectionEntityIPAddress failed: %w", err)
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
func (s *TraceActivation) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "traceID", Optional: false},
			per.ComponentInfo{Name: "interfacesToTrace", Optional: false},
			per.ComponentInfo{Name: "traceDepth", Optional: false},
			per.ComponentInfo{Name: "traceCollectionEntityIPAddress", Optional: false},
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
		val, err := r.DecodeOctetString(per.SizeConstraints{Extensible: false, Min: int64Ptr(8), Max: int64Ptr(8)})
		if err != nil {
			return fmt.Errorf("decode TraceID failed: %w", err)
		}
		s.TraceID.Value = val
	}

	{
		val, err := r.DecodeBitString(per.SizeConstraints{Extensible: false, Min: int64Ptr(8), Max: int64Ptr(8)})
		if err != nil {
			return fmt.Errorf("decode InterfacesToTrace failed: %w", err)
		}
		s.InterfacesToTrace.Value = val
	}

	{
		enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 6), ExtValues: nil}
		val, err := r.DecodeEnumerated(enumC)
		if err != nil {
			return fmt.Errorf("decode TraceDepth failed: %w", err)
		}
		s.TraceDepth.Value = val
	}

	{
		val, err := r.DecodeBitString(per.SizeConstraints{Extensible: false, Min: int64Ptr(1), Max: int64Ptr(160)})
		if err != nil {
			return fmt.Errorf("decode TraceCollectionEntityIPAddress failed: %w", err)
		}
		s.TraceCollectionEntityIPAddress.Value = val
	}

	if seqDecoder.IsComponentPresent(4) {
		s.IEExtensions = new(TraceActivationExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
