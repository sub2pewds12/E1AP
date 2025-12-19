package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// TransportUPLayerAddressesInfoToAddItem is a generated SEQUENCE type.
type TransportUPLayerAddressesInfoToAddItem struct {
	IPSecTransportLayerAddress      TransportLayerAddress       `aper:"lb:1,ub:160,mandatory,ext"`
	GTPTransportLayerAddressesToAdd *GTPTLAs                    `aper:"ub:MaxnoofGTPTLAs,optional,ext"`
	IEExtensions                    *ProtocolExtensionContainer `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *TransportUPLayerAddressesInfoToAddItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.GTPTransportLayerAddressesToAdd != nil {
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
	if s.GTPTransportLayerAddressesToAdd != nil {
		if err = s.GTPTransportLayerAddressesToAdd.Encode(w); err != nil {
			return fmt.Errorf("Encode GTPTransportLayerAddressesToAdd failed: %w", err)
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
func (s *TransportUPLayerAddressesInfoToAddItem) Decode(r *aper.AperReader) (err error) {
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
		s.GTPTransportLayerAddressesToAdd = new(GTPTLAs)
		if err = s.GTPTransportLayerAddressesToAdd.Decode(r); err != nil {
			return fmt.Errorf("Decode GTPTransportLayerAddressesToAdd failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible { /* TODO: Implement extension skipping for TransportUPLayerAddressesInfoToAddItem */
	}
	return nil
}
