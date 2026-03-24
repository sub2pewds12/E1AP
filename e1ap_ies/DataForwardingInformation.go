package e1ap_ies

import (
	"fmt"

	"asn1go/per"
)

// DataForwardingInformation is a generated SEQUENCE type.
type DataForwardingInformation struct {
	ULDataForwarding *UPTNLInformation
	DLDataForwarding *UPTNLInformation
	IEExtensions     *DataForwardingInformationExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *DataForwardingInformation) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "uL-Data-Forwarding", Optional: true},
			per.ComponentInfo{Name: "dL-Data-Forwarding", Optional: true},
			per.ComponentInfo{Name: "iE-Extensions", Optional: true},
		},
	}
	seqEncoder := w.NewSequenceEncoder(c)
	if err := seqEncoder.EncodeExtensionBit(false); err != nil {
		return err
	}

	optionalBitmap := make([]bool, 0)

	if s.ULDataForwarding != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.DLDataForwarding != nil {
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

	if s.ULDataForwarding != nil {
		if err = s.ULDataForwarding.Encode(w); err != nil {
			return fmt.Errorf("encode ULDataForwarding failed: %w", err)
		}
	}

	if s.DLDataForwarding != nil {
		if err = s.DLDataForwarding.Encode(w); err != nil {
			return fmt.Errorf("encode DLDataForwarding failed: %w", err)
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
func (s *DataForwardingInformation) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "uL-Data-Forwarding", Optional: true},
			per.ComponentInfo{Name: "dL-Data-Forwarding", Optional: true},
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

	if seqDecoder.IsComponentPresent(0) {
		s.ULDataForwarding = new(UPTNLInformation)
		if err = s.ULDataForwarding.Decode(r); err != nil {
			return fmt.Errorf("Decode ULDataForwarding failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(1) {
		s.DLDataForwarding = new(UPTNLInformation)
		if err = s.DLDataForwarding.Decode(r); err != nil {
			return fmt.Errorf("Decode DLDataForwarding failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(2) {
		s.IEExtensions = new(DataForwardingInformationExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
