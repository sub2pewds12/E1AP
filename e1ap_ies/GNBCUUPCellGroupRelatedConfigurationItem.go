package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/asn1go/per"
)

// GNBCUUPCellGroupRelatedConfigurationItem is a generated SEQUENCE type.
type GNBCUUPCellGroupRelatedConfigurationItem struct {
	CellGroupID      CellGroupID
	UPTNLInformation UPTNLInformation
	ULConfiguration  *ULConfiguration
	IEExtensions     *GNBCUUPCellGroupRelatedConfigurationItemExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *GNBCUUPCellGroupRelatedConfigurationItem) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: false,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "cell-Group-ID", Optional: false},
			per.ComponentInfo{Name: "uP-TNL-Information", Optional: false},
			per.ComponentInfo{Name: "uL-Configuration", Optional: true},
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
	if err = s.UPTNLInformation.Encode(w); err != nil {
		return fmt.Errorf("encode UPTNLInformation failed: %w", err)
	}

	if s.ULConfiguration != nil {

		{
			enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 3), ExtValues: nil}
			if err = w.EncodeEnumerated(int64((*s.ULConfiguration).Value), enumC); err != nil {
				return fmt.Errorf("encode ULConfiguration failed: %w", err)
			}
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
func (s *GNBCUUPCellGroupRelatedConfigurationItem) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: false,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "cell-Group-ID", Optional: false},
			per.ComponentInfo{Name: "uP-TNL-Information", Optional: false},
			per.ComponentInfo{Name: "uL-Configuration", Optional: true},
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
	if err = s.UPTNLInformation.Decode(r); err != nil {
		return fmt.Errorf("Decode UPTNLInformation failed: %w", err)
	}

	if seqDecoder.IsComponentPresent(2) {
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

	if seqDecoder.IsComponentPresent(3) {
		s.IEExtensions = new(GNBCUUPCellGroupRelatedConfigurationItemExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	return nil
}
