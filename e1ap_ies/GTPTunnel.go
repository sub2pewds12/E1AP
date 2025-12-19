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
		return fmt.Errorf("encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(1), &aper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
		return fmt.Errorf("encode optionality bitmap failed: %w", err)
	}
	if err = w.WriteBitString(s.TransportLayerAddress.Value.Bytes, uint(s.TransportLayerAddress.Value.NumBits), &aper.Constraint{Lb: 1, Ub: 160}, false); err != nil {
		return fmt.Errorf("encode TransportLayerAddress failed: %w", err)
	}
	if err = w.WriteOctetString([]byte(s.GTPTEID.Value), &aper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
		return fmt.Errorf("encode GTPTEID failed: %w", err)
	}
	if s.IEExtensions != nil {
		if err = s.IEExtensions.Encode(w); err != nil {
			return fmt.Errorf("encode IEExtensions failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *GTPTunnel) Decode(r *aper.AperReader) (err error) {
	isExtensible, err := r.ReadBool()
	if err != nil {
		return fmt.Errorf("read extensibility bool failed: %w", err)
	}
	_ = isExtensible
	optionalityBitmap, _, err := r.ReadBitString(&aper.Constraint{Lb: 1, Ub: 1}, false)
	if err != nil {
		return fmt.Errorf("read optionality bitmap failed: %w", err)
	}
	if err = s.TransportLayerAddress.Decode(r); err != nil {
		return fmt.Errorf("decode TransportLayerAddress failed: %w", err)
	}
	if err = s.GTPTEID.Decode(r); err != nil {
		return fmt.Errorf("decode GTPTEID failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible { /* TODO: Implement extension skipping for GTPTunnel */
	}
	return nil
}
