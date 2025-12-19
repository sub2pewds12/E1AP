package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// TraceActivation is a generated SEQUENCE type.
type TraceActivation struct {
	TraceID                        TraceID                    `aper:"lb:8,ub:8,mandatory,ext"`
	InterfacesToTrace              InterfacesToTrace          `aper:"lb:8,ub:8,mandatory,ext"`
	TraceDepth                     TraceDepth                 `aper:"mandatory,ext"`
	TraceCollectionEntityIPAddress TransportLayerAddress      `aper:"lb:1,ub:160,mandatory,ext"`
	IEExtensions                   *TraceActivationExtensions `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *TraceActivation) Encode(w *aper.AperWriter) (err error) {
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
	if err = w.WriteOctetString([]byte(s.TraceID.Value), &aper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
		return fmt.Errorf("Encode TraceID failed: %w", err)
	}
	if err = w.WriteBitString(s.InterfacesToTrace.Value.Bytes, uint(s.InterfacesToTrace.Value.NumBits), &aper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
		return fmt.Errorf("Encode InterfacesToTrace failed: %w", err)
	}
	if err = w.WriteEnumerate(uint64(s.TraceDepth.Value), aper.Constraint{Lb: 0, Ub: 5}, true); err != nil {
		return fmt.Errorf("Encode TraceDepth failed: %w", err)
	}
	if err = w.WriteBitString(s.TraceCollectionEntityIPAddress.Value.Bytes, uint(s.TraceCollectionEntityIPAddress.Value.NumBits), &aper.Constraint{Lb: 1, Ub: 160}, false); err != nil {
		return fmt.Errorf("Encode TraceCollectionEntityIPAddress failed: %w", err)
	}
	if s.IEExtensions != nil {
		if err = s.IEExtensions.Encode(w); err != nil {
			return fmt.Errorf("Encode IEExtensions failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *TraceActivation) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}
	if err = s.TraceID.Decode(r); err != nil {
		return fmt.Errorf("Decode TraceID failed: %w", err)
	}
	if err = s.InterfacesToTrace.Decode(r); err != nil {
		return fmt.Errorf("Decode InterfacesToTrace failed: %w", err)
	}
	if err = s.TraceDepth.Decode(r); err != nil {
		return fmt.Errorf("Decode TraceDepth failed: %w", err)
	}
	if err = s.TraceCollectionEntityIPAddress.Decode(r); err != nil {
		return fmt.Errorf("Decode TraceCollectionEntityIPAddress failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.IEExtensions = new(TraceActivationExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible { /* TODO: Implement extension skipping for TraceActivation */
	}
	return nil
}
