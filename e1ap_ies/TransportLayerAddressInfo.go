package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// TransportLayerAddressInfo is a generated SEQUENCE type.
type TransportLayerAddressInfo struct {
	TransportUPLayerAddressesInfoToAddList    []TransportUPLayerAddressesInfoToAddItem    `aper:"optional,ext"`
	TransportUPLayerAddressesInfoToRemoveList []TransportUPLayerAddressesInfoToRemoveItem `aper:"optional,ext"`
	IEExtensions                              *ProtocolExtensionContainer                 `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *TransportLayerAddressInfo) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.TransportUPLayerAddressesInfoToAddList != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.TransportUPLayerAddressesInfoToRemoveList != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 5
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(3), &aper.Constraint{Lb: 3, Ub: 3}, false); err != nil {
		return fmt.Errorf("Encode optionality bitmap failed: %w", err)
	}
	if s.TransportUPLayerAddressesInfoToAddList != nil {
		if err = s.TransportUPLayerAddressesInfoToAddList.Encode(w); err != nil {
			return fmt.Errorf("Encode TransportUPLayerAddressesInfoToAddList failed: %w", err)
		}
	}
	if s.TransportUPLayerAddressesInfoToRemoveList != nil {
		if err = s.TransportUPLayerAddressesInfoToRemoveList.Encode(w); err != nil {
			return fmt.Errorf("Encode TransportUPLayerAddressesInfoToRemoveList failed: %w", err)
		}
	}
	if s.IEExtensions != nil {
		if err = s.IEExtensions.Encode(w); err != nil {
			return fmt.Errorf("Encode IEExtensions failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *TransportLayerAddressInfo) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 3, Ub: 3}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.TransportUPLayerAddressesInfoToAddList = new(TransportUPLayerAddressesInfoToAddList)
		if err = s.TransportUPLayerAddressesInfoToAddList.Decode(r); err != nil {
			return fmt.Errorf("Decode TransportUPLayerAddressesInfoToAddList failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.TransportUPLayerAddressesInfoToRemoveList = new(TransportUPLayerAddressesInfoToRemoveList)
		if err = s.TransportUPLayerAddressesInfoToRemoveList.Decode(r); err != nil {
			return fmt.Errorf("Decode TransportUPLayerAddressesInfoToRemoveList failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<5) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for TransportLayerAddressInfo")
	}
	return nil
}
