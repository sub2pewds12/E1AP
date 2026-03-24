package e1ap_ies

import (
	"fmt"

	"asn1go/per"
)

// CellGroupInformationItem is a generated SEQUENCE type.
type CellGroupInformationItem struct {
	CellGroupID     CellGroupID
	ULConfiguration *ULConfiguration
	DLTXStop        *DLTXStop
	RATType         *RATType
	IEExtensions    *CellGroupInformationItemExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *CellGroupInformationItem) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "cell-Group-ID", Optional: false},
			per.ComponentInfo{Name: "uL-Configuration", Optional: true},
			per.ComponentInfo{Name: "dL-TX-Stop", Optional: true},
			per.ComponentInfo{Name: "rAT-Type", Optional: true},
			per.ComponentInfo{Name: "iE-Extensions", Optional: true},
		},
	}
	seqEncoder := w.NewSequenceEncoder(c)
	if err := seqEncoder.EncodeExtensionBit(false); err != nil {
		return err
	}

	optionalBitmap := make([]bool, 0)

	if s.ULConfiguration != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.DLTXStop != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.RATType != nil {
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

	if err = w.EncodeInteger(int64(s.CellGroupID.Value), per.ConstrainedExtensible(0, 3)); err != nil {
		return fmt.Errorf("encode CellGroupID failed: %w", err)
	}

	if s.ULConfiguration != nil {

		{
			enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 3), ExtValues: nil}
			if err = w.EncodeEnumerated(int64((*s.ULConfiguration).Value), enumC); err != nil {
				return fmt.Errorf("encode ULConfiguration failed: %w", err)
			}
		}
	}

	if s.DLTXStop != nil {

		{
			enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 2), ExtValues: nil}
			if err = w.EncodeEnumerated(int64((*s.DLTXStop).Value), enumC); err != nil {
				return fmt.Errorf("encode DLTXStop failed: %w", err)
			}
		}
	}

	if s.RATType != nil {

		{
			enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 2), ExtValues: nil}
			if err = w.EncodeEnumerated(int64((*s.RATType).Value), enumC); err != nil {
				return fmt.Errorf("encode RATType failed: %w", err)
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
func (s *CellGroupInformationItem) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "cell-Group-ID", Optional: false},
			per.ComponentInfo{Name: "uL-Configuration", Optional: true},
			per.ComponentInfo{Name: "dL-TX-Stop", Optional: true},
			per.ComponentInfo{Name: "rAT-Type", Optional: true},
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
		val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 3))
		if err != nil {
			return fmt.Errorf("decode CellGroupID failed: %w", err)
		}
		s.CellGroupID.Value = val
	}

	if seqDecoder.IsComponentPresent(1) {
		s.ULConfiguration = new(ULConfiguration)

		{
			enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 3), ExtValues: nil}
			val, err := r.DecodeEnumerated(enumC)
			if err != nil {
				return fmt.Errorf("decode ULConfiguration failed: %w", err)
			}
			s.ULConfiguration.Value = val
		}
	}

	if seqDecoder.IsComponentPresent(2) {
		s.DLTXStop = new(DLTXStop)

		{
			enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 2), ExtValues: nil}
			val, err := r.DecodeEnumerated(enumC)
			if err != nil {
				return fmt.Errorf("decode DLTXStop failed: %w", err)
			}
			s.DLTXStop.Value = val
		}
	}

	if seqDecoder.IsComponentPresent(3) {
		s.RATType = new(RATType)

		{
			enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 2), ExtValues: nil}
			val, err := r.DecodeEnumerated(enumC)
			if err != nil {
				return fmt.Errorf("decode RATType failed: %w", err)
			}
			s.RATType.Value = val
		}
	}

	if seqDecoder.IsComponentPresent(4) {
		s.IEExtensions = new(CellGroupInformationItemExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
