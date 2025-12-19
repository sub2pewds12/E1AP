package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// DRBBStatusTransfer is a generated SEQUENCE type.
type DRBBStatusTransfer struct {
	ReceiveStatusofPDCPSDU *DRBBStatusTransferReceiveStatusofPDCPSDU `aper:"lb:1,ub:131072,optional,ext"`
	CountValue             PDCPCount                                 `aper:"mandatory,ext"`
	IEExtension            *ProtocolExtensionContainer               `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *DRBBStatusTransfer) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.ReceiveStatusofPDCPSDU != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.IEExtension != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(2), &aper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return fmt.Errorf("encode optionality bitmap failed: %w", err)
	}
	if s.ReceiveStatusofPDCPSDU != nil {
		if err = w.WriteBitString((*s.ReceiveStatusofPDCPSDU).Value.Bytes, uint((*s.ReceiveStatusofPDCPSDU).Value.NumBits), &aper.Constraint{Lb: 1, Ub: 131072}, false); err != nil {
			return fmt.Errorf("encode ReceiveStatusofPDCPSDU failed: %w", err)
		}
	}
	if err = s.CountValue.Encode(w); err != nil {
		return fmt.Errorf("encode CountValue failed: %w", err)
	}
	if s.IEExtension != nil {
		if err = s.IEExtension.Encode(w); err != nil {
			return fmt.Errorf("encode IEExtension failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *DRBBStatusTransfer) Decode(r *aper.AperReader) (err error) {
	isExtensible, err := r.ReadBool()
	if err != nil {
		return fmt.Errorf("read extensibility bool failed: %w", err)
	}
	_ = isExtensible
	optionalityBitmap, _, err := r.ReadBitString(&aper.Constraint{Lb: 2, Ub: 2}, false)
	if err != nil {
		return fmt.Errorf("read optionality bitmap failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.ReceiveStatusofPDCPSDU = new(DRBBStatusTransferReceiveStatusofPDCPSDU)
		if err = s.ReceiveStatusofPDCPSDU.Decode(r); err != nil {
			return fmt.Errorf("decode ReceiveStatusofPDCPSDU failed: %w", err)
		}
	}
	if err = s.CountValue.Decode(r); err != nil {
		return fmt.Errorf("decode CountValue failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.IEExtension = new(ProtocolExtensionContainer)
		if err = s.IEExtension.Decode(r); err != nil {
			return fmt.Errorf("decode IEExtension failed: %w", err)
		}
	}
	if isExtensible { /* TODO: Implement extension skipping for DRBBStatusTransfer */
	}
	return nil
}
