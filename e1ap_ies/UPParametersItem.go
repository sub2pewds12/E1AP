package e1ap_ies

import (
	"fmt"

	"asn1go/per"
)

// UPParametersItem is a generated SEQUENCE type.
type UPParametersItem struct {
	UPTNLInformation UPTNLInformation
	CellGroupID      CellGroupID
	IEExtensions     *UPParametersItemExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *UPParametersItem) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "uP-TNL-Information", Optional: false},
			per.ComponentInfo{Name: "cell-Group-ID", Optional: false},
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

	if err = s.UPTNLInformation.Encode(w); err != nil {
		return fmt.Errorf("encode UPTNLInformation failed: %w", err)
	}
	if err = w.EncodeInteger(int64(s.CellGroupID.Value), per.ConstrainedExtensible(0, 3)); err != nil {
		return fmt.Errorf("encode CellGroupID failed: %w", err)
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
func (s *UPParametersItem) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "uP-TNL-Information", Optional: false},
			per.ComponentInfo{Name: "cell-Group-ID", Optional: false},
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

	if err = s.UPTNLInformation.Decode(r); err != nil {
		return fmt.Errorf("Decode UPTNLInformation failed: %w", err)
	}

	{
		val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 3))
		if err != nil {
			return fmt.Errorf("decode CellGroupID failed: %w", err)
		}
		s.CellGroupID.Value = val
	}

	if seqDecoder.IsComponentPresent(2) {
		s.IEExtensions = new(UPParametersItemExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
