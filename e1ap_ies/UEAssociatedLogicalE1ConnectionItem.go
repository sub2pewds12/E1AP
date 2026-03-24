package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/asn1go/per"
)

// UEAssociatedLogicalE1ConnectionItem is a generated SEQUENCE type.
type UEAssociatedLogicalE1ConnectionItem struct {
	GNBCUCPUEE1APID GNBCUCPUEE1APID
	GNBCUUPUEE1APID GNBCUUPUEE1APID
	IEExtensions    *UEAssociatedLogicalE1ConnectionItemExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *UEAssociatedLogicalE1ConnectionItem) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "gNB-CU-CP-UE-E1AP-ID", Optional: false},
			per.ComponentInfo{Name: "gNB-CU-UP-UE-E1AP-ID", Optional: false},
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

	if err = w.EncodeInteger(int64(s.GNBCUCPUEE1APID.Value), per.Unconstrained()); err != nil {
		return fmt.Errorf("encode GNBCUCPUEE1APID failed: %w", err)
	}
	if err = w.EncodeInteger(int64(s.GNBCUUPUEE1APID.Value), per.Unconstrained()); err != nil {
		return fmt.Errorf("encode GNBCUUPUEE1APID failed: %w", err)
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
func (s *UEAssociatedLogicalE1ConnectionItem) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "gNB-CU-CP-UE-E1AP-ID", Optional: false},
			per.ComponentInfo{Name: "gNB-CU-UP-UE-E1AP-ID", Optional: false},
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
		val, err := r.DecodeInteger(per.Unconstrained())
		if err != nil {
			return fmt.Errorf("decode GNBCUCPUEE1APID failed: %w", err)
		}
		s.GNBCUCPUEE1APID.Value = val
	}

	{
		val, err := r.DecodeInteger(per.Unconstrained())
		if err != nil {
			return fmt.Errorf("decode GNBCUUPUEE1APID failed: %w", err)
		}
		s.GNBCUUPUEE1APID.Value = val
	}

	if seqDecoder.IsComponentPresent(2) {
		s.IEExtensions = new(UEAssociatedLogicalE1ConnectionItemExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
