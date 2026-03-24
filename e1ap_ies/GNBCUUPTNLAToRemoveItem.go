package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/asn1go/per"
)

// GNBCUUPTNLAToRemoveItem is a generated SEQUENCE type.
type GNBCUUPTNLAToRemoveItem struct {
	TNLAssociationTransportLayerAddress        CPTNLInformation
	TNLAssociationTransportLayerAddressgNBCUCP *CPTNLInformation
	IEExtensions                               *GNBCUUPTNLAToRemoveItemExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *GNBCUUPTNLAToRemoveItem) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: false,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "tNLAssociationTransportLayerAddress", Optional: false},
			per.ComponentInfo{Name: "tNLAssociationTransportLayerAddressgNBCUCP", Optional: true},
			per.ComponentInfo{Name: "iE-Extensions", Optional: true},
		},
	}
	seqEncoder := w.NewSequenceEncoder(c)
	if err := seqEncoder.EncodeExtensionBit(false); err != nil {
		return err
	}

	optionalBitmap := make([]bool, 0)

	if s.TNLAssociationTransportLayerAddressgNBCUCP != nil {
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

	if s.TNLAssociationTransportLayerAddressgNBCUCP != nil {
		if err = s.TNLAssociationTransportLayerAddressgNBCUCP.Encode(w); err != nil {
			return fmt.Errorf("encode TNLAssociationTransportLayerAddressgNBCUCP failed: %w", err)
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
func (s *GNBCUUPTNLAToRemoveItem) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: false,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "tNLAssociationTransportLayerAddress", Optional: false},
			per.ComponentInfo{Name: "tNLAssociationTransportLayerAddressgNBCUCP", Optional: true},
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
		s.TNLAssociationTransportLayerAddressgNBCUCP = new(CPTNLInformation)
		if err = s.TNLAssociationTransportLayerAddressgNBCUCP.Decode(r); err != nil {
			return fmt.Errorf("Decode TNLAssociationTransportLayerAddressgNBCUCP failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(2) {
		s.IEExtensions = new(GNBCUUPTNLAToRemoveItemExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	return nil
}
