package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// EndpointIPAddressAndPort is a generated SEQUENCE type.
type EndpointIPAddressAndPort struct {
	EndpointIPAddress TransportLayerAddress       `aper:"lb:1,ub:160,mandatory"`
	PortNumber        PortNumber                  `aper:"lb:16,ub:16,mandatory"`
	IEExtensions      *ProtocolExtensionContainer `aper:"optional"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *EndpointIPAddressAndPort) Encode(w *aper.AperWriter) (err error) {
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
	if err = w.WriteBitString(s.EndpointIPAddress.Value.Bytes, uint(s.EndpointIPAddress.Value.NumBits), &aper.Constraint{Lb: 1, Ub: 160}, false); err != nil {
		return fmt.Errorf("Encode EndpointIPAddress failed: %w", err)
	}
	if err = w.WriteBitString(s.PortNumber.Value.Bytes, uint(s.PortNumber.Value.NumBits), &aper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
		return fmt.Errorf("Encode PortNumber failed: %w", err)
	}
	if s.IEExtensions != nil {
		if err = s.IEExtensions.Encode(w); err != nil {
			return fmt.Errorf("Encode IEExtensions failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *EndpointIPAddressAndPort) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}
	if err = s.EndpointIPAddress.Decode(r); err != nil {
		return fmt.Errorf("Decode EndpointIPAddress failed: %w", err)
	}
	if err = s.PortNumber.Decode(r); err != nil {
		return fmt.Errorf("Decode PortNumber failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible { /* TODO: Implement extension skipping for EndpointIPAddressAndPort */
	}
	return nil
}
