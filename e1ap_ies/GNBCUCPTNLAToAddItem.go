package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// GNBCUCPTNLAToAddItem is a generated SEQUENCE type.
type GNBCUCPTNLAToAddItem struct {
	TNLAssociationTransportLayerAddress CPTNLInformation            `aper:"mandatory"`
	TNLAssociationUsage                 TNLAssociationUsage         `aper:"mandatory"`
	IEExtensions                        *ProtocolExtensionContainer `aper:"optional"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *GNBCUCPTNLAToAddItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(false); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(1), &aper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
		return fmt.Errorf("Encode optionality bitmap failed: %w", err)
	}
	if err = s.TNLAssociationTransportLayerAddress.Encode(w); err != nil {
		return fmt.Errorf("Encode TNLAssociationTransportLayerAddress failed: %w", err)
	}
	if err = s.TNLAssociationUsage.Encode(w); err != nil {
		return fmt.Errorf("Encode TNLAssociationUsage failed: %w", err)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *GNBCUCPTNLAToAddItem) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}
	if err = s.TNLAssociationTransportLayerAddress.Decode(r); err != nil {
		return fmt.Errorf("Decode TNLAssociationTransportLayerAddress failed: %w", err)
	}
	if err = s.TNLAssociationUsage.Decode(r); err != nil {
		return fmt.Errorf("Decode TNLAssociationUsage failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for GNBCUCPTNLAToAddItem")
	}
	return nil
}
