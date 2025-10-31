package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// GTPTunnel is a generated SEQUENCE type.
type GTPTunnel struct {
	TransportLayerAddress TransportLayerAddress       `aper:"lb:1,ub:160,mandatory,ext"`
	GTPTEID               GTPTEID                     `aper:"lb:4,ub:4,mandatory,ext"`
	IEExtensions          *ProtocolExtensionContainer `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *GTPTunnel) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(1), &aper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
		return fmt.Errorf("Encode optionality bitmap failed: %w", err)
	}
	if err = w.WriteBitString(s.TransportLayerAddress.Bytes, uint(s.TransportLayerAddress.NumBits), &aper.Constraint{Lb: 1, Ub: 160}, false); err != nil {
		return fmt.Errorf("Encode TransportLayerAddress failed: %w", err)
	}
	if err = w.WriteOctetString([]byte(s.GTPTEID), &aper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
		return fmt.Errorf("Encode GTPTEID failed: %w", err)
	}
	if s.IEExtensions != nil {
		if err = s.IEExtensions.Encode(w); err != nil {
			return fmt.Errorf("Encode IEExtensions failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *GTPTunnel) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}

	{
		var val []byte
		if val, err = r.ReadOctetString(&aper.Constraint{Lb: 1, Ub: 160}, false); err != nil {
			return fmt.Errorf("Decode TransportLayerAddress failed: %w", err)
		}
		s.TransportLayerAddress = TransportLayerAddress(val)
	}

	{
		var val []byte
		if val, err = r.ReadOctetString(&aper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
			return fmt.Errorf("Decode GTPTEID failed: %w", err)
		}
		s.GTPTEID = GTPTEID(val)
	}

	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for GTPTunnel")
	}
	return nil
}
