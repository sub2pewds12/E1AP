package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/asn1go/per"
)

// QOSParametersSupportList is a generated SEQUENCE type.
type QOSParametersSupportList struct {
	EUTRANQOSSupportList *EUTRANQOSSupportList
	NGRANQOSSupportList  *NGRANQOSSupportList
	IEExtensions         *QOSParametersSupportListExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *QOSParametersSupportList) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "eUTRAN-QoS-Support-List", Optional: true},
			per.ComponentInfo{Name: "nG-RAN-QoS-Support-List", Optional: true},
			per.ComponentInfo{Name: "iE-Extensions", Optional: true},
		},
	}
	seqEncoder := w.NewSequenceEncoder(c)
	if err := seqEncoder.EncodeExtensionBit(false); err != nil {
		return err
	}

	optionalBitmap := make([]bool, 0)

	if s.EUTRANQOSSupportList != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.NGRANQOSSupportList != nil {
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

	if s.EUTRANQOSSupportList != nil {
		if err = s.EUTRANQOSSupportList.Encode(w); err != nil {
			return fmt.Errorf("encode EUTRANQOSSupportList failed: %w", err)
		}
	}

	if s.NGRANQOSSupportList != nil {
		if err = s.NGRANQOSSupportList.Encode(w); err != nil {
			return fmt.Errorf("encode NGRANQOSSupportList failed: %w", err)
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
func (s *QOSParametersSupportList) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "eUTRAN-QoS-Support-List", Optional: true},
			per.ComponentInfo{Name: "nG-RAN-QoS-Support-List", Optional: true},
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
		s.EUTRANQOSSupportList = new(EUTRANQOSSupportList)
		if err = s.EUTRANQOSSupportList.Decode(r); err != nil {
			return fmt.Errorf("Decode EUTRANQOSSupportList failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(1) {
		s.NGRANQOSSupportList = new(NGRANQOSSupportList)
		if err = s.NGRANQOSSupportList.Decode(r); err != nil {
			return fmt.Errorf("Decode NGRANQOSSupportList failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(2) {
		s.IEExtensions = new(QOSParametersSupportListExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
