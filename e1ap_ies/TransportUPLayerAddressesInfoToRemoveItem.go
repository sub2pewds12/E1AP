package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// TransportUPLayerAddressesInfoToRemoveItem is a generated SEQUENCE type.
type TransportUPLayerAddressesInfoToRemoveItem struct {
	IPSecTransportLayerAddress         TransportLayerAddress       `aper:"lb:1,ub:160,mandatory,ext"`
	GTPTransportLayerAddressesToRemove *GTPTLAs                    `aper:"ub:MaxnoofGTPTLAs,optional,ext"`
	IEExtensions                       *ProtocolExtensionContainer `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *TransportUPLayerAddressesInfoToRemoveItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.GTPTransportLayerAddressesToRemove != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(2), &aper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return fmt.Errorf("Encode optionality bitmap failed: %w", err)
	}
	if err = w.WriteBitString(s.IPSecTransportLayerAddress.Value.Bytes, uint(s.IPSecTransportLayerAddress.Value.NumBits), &aper.Constraint{Lb: 1, Ub: 160}, false); err != nil {
		return fmt.Errorf("Encode IPSecTransportLayerAddress failed: %w", err)
	}
	if s.GTPTransportLayerAddressesToRemove != nil {
		{
			itemPointers := make([]aper.AperMarshaller, len(s.GTPTransportLayerAddressesToRemove.Value))
			for i := 0; i < len(s.GTPTransportLayerAddressesToRemove.Value); i++ {
				itemPointers[i] = &(s.GTPTransportLayerAddressesToRemove.Value[i])
			}
			if err = aper.WriteSequenceOf(itemPointers, w, &aper.Constraint{Lb: 0, Ub: MaxnoofGTPTLAs}, false); err != nil {
				return fmt.Errorf("Encode GTPTransportLayerAddressesToRemove failed: %w", err)
			}
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *TransportUPLayerAddressesInfoToRemoveItem) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}
	if err = s.IPSecTransportLayerAddress.Decode(r); err != nil {
		return fmt.Errorf("Decode IPSecTransportLayerAddress failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.GTPTransportLayerAddressesToRemove = new(GTPTLAs)
		if err = s.GTPTransportLayerAddressesToRemove.Decode(r); err != nil {
			return fmt.Errorf("Decode GTPTransportLayerAddressesToRemove failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible { /* TODO: Implement extension skipping for TransportUPLayerAddressesInfoToRemoveItem */
	}
	return nil
}
