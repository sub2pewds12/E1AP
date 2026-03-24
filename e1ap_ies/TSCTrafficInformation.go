package e1ap_ies

import (
	"fmt"

	"asn1go/per"
)

// TSCTrafficInformation is a generated SEQUENCE type.
type TSCTrafficInformation struct {
	Periodicity      Periodicity
	BurstArrivalTime *BurstArrivalTime
	IEExtensions     *TSCTrafficInformationExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *TSCTrafficInformation) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: false,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "periodicity", Optional: false},
			per.ComponentInfo{Name: "burstArrivalTime", Optional: true},
			per.ComponentInfo{Name: "iE-Extensions", Optional: true},
		},
	}
	seqEncoder := w.NewSequenceEncoder(c)
	if err := seqEncoder.EncodeExtensionBit(false); err != nil {
		return err
	}

	optionalBitmap := make([]bool, 0)

	if s.BurstArrivalTime != nil {
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

	if err = w.EncodeInteger(int64(s.Periodicity.Value), per.ConstrainedExtensible(1, 640000)); err != nil {
		return fmt.Errorf("encode Periodicity failed: %w", err)
	}

	if s.BurstArrivalTime != nil {
		if err = w.EncodeOctetString([]byte((*s.BurstArrivalTime).Value), per.SizeConstraints{Extensible: false, Min: int64Ptr(0), Max: int64Ptr(0)}); err != nil {
			return fmt.Errorf("encode BurstArrivalTime failed: %w", err)
		}
	}

	if s.IEExtensions != nil {
		if err = s.IEExtensions.Encode(w); err != nil {
			return fmt.Errorf("encode IEExtensions failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *TSCTrafficInformation) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: false,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "periodicity", Optional: false},
			per.ComponentInfo{Name: "burstArrivalTime", Optional: true},
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
		val, err := r.DecodeInteger(per.ConstrainedExtensible(1, 640000))
		if err != nil {
			return fmt.Errorf("decode Periodicity failed: %w", err)
		}
		s.Periodicity.Value = val
	}

	if seqDecoder.IsComponentPresent(1) {
		s.BurstArrivalTime = new(BurstArrivalTime)

		{
			val, err := r.DecodeOctetString(per.SizeConstraints{Extensible: false, Min: int64Ptr(0), Max: int64Ptr(0)})
			if err != nil {
				return fmt.Errorf("decode BurstArrivalTime failed: %w", err)
			}
			s.BurstArrivalTime.Value = val
		}
	}

	if seqDecoder.IsComponentPresent(2) {
		s.IEExtensions = new(TSCTrafficInformationExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	return nil
}
