package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/asn1go/per"
)

// GNBCUCPTNLAToUpdateItem is a generated SEQUENCE type.
type GNBCUCPTNLAToUpdateItem struct {
	TNLAssociationTransportLayerAddress CPTNLInformation
	TNLAssociationUsage                 *TNLAssociationUsage
	IEExtensions                        *GNBCUCPTNLAToUpdateItemExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *GNBCUCPTNLAToUpdateItem) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: false,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "tNLAssociationTransportLayerAddress", Optional: false},
			per.ComponentInfo{Name: "tNLAssociationUsage", Optional: true},
			per.ComponentInfo{Name: "iE-Extensions", Optional: true},
		},
	}
	seqEncoder := w.NewSequenceEncoder(c)
	if err := seqEncoder.EncodeExtensionBit(false); err != nil {
		return err
	}

	optionalBitmap := make([]bool, 0)

	if s.TNLAssociationUsage != nil {
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

	if err = s.TNLAssociationTransportLayerAddress.Encode(w); err != nil {
		return fmt.Errorf("encode TNLAssociationTransportLayerAddress failed: %w", err)
	}

	if s.TNLAssociationUsage != nil {

		{
			enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 3), ExtValues: nil}
			if err = w.EncodeEnumerated(int64((*s.TNLAssociationUsage).Value), enumC); err != nil {
				return fmt.Errorf("encode TNLAssociationUsage failed: %w", err)
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
func (s *GNBCUCPTNLAToUpdateItem) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: false,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "tNLAssociationTransportLayerAddress", Optional: false},
			per.ComponentInfo{Name: "tNLAssociationUsage", Optional: true},
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

	if err = s.TNLAssociationTransportLayerAddress.Decode(r); err != nil {
		return fmt.Errorf("Decode TNLAssociationTransportLayerAddress failed: %w", err)
	}

	if seqDecoder.IsComponentPresent(1) {
		s.TNLAssociationUsage = new(TNLAssociationUsage)

		{
			enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 3), ExtValues: nil}
			val, err := r.DecodeEnumerated(enumC)
			if err != nil {
				return fmt.Errorf("decode TNLAssociationUsage failed: %w", err)
			}
			s.TNLAssociationUsage.Value = val
		}
	}

	if seqDecoder.IsComponentPresent(2) {
		s.IEExtensions = new(GNBCUCPTNLAToUpdateItemExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	return nil
}
