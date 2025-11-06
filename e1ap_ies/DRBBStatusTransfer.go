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
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.ReceiveStatusofPDCPSDU != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.IEExtension != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(2), &aper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return fmt.Errorf("Encode optionality bitmap failed: %w", err)
	}
	if s.ReceiveStatusofPDCPSDU != nil {
		if err = w.WriteBitString(s.ReceiveStatusofPDCPSDU.Value.Bytes, uint(s.ReceiveStatusofPDCPSDU.Value.NumBits), &aper.Constraint{Lb: 1, Ub: 131072}, false); err != nil {
			return fmt.Errorf("Encode ReceiveStatusofPDCPSDU failed: %w", err)
		}
	}
	if err = s.CountValue.Encode(w); err != nil {
		return fmt.Errorf("Encode CountValue failed: %w", err)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *DRBBStatusTransfer) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.ReceiveStatusofPDCPSDU = new(DRBBStatusTransferReceiveStatusofPDCPSDU)
		if err = s.ReceiveStatusofPDCPSDU.Decode(r); err != nil {
			return fmt.Errorf("Decode ReceiveStatusofPDCPSDU failed: %w", err)
		}
	}
	if err = s.CountValue.Decode(r); err != nil {
		return fmt.Errorf("Decode CountValue failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.IEExtension = new(ProtocolExtensionContainer)
		if err = s.IEExtension.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtension failed: %w", err)
		}
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for DRBBStatusTransfer")
	}
	return nil
}
