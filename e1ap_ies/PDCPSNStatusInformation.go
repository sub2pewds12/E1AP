package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// PDCPSNStatusInformation is a generated SEQUENCE type.
type PDCPSNStatusInformation struct {
	PdcpStatusTransferUL DRBBStatusTransfer          `aper:"mandatory,ext"`
	PdcpStatusTransferDL PDCPCount                   `aper:"mandatory,ext"`
	IEExtension          *ProtocolExtensionContainer `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *PDCPSNStatusInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.IEExtension != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(1), &aper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
		return fmt.Errorf("encode optionality bitmap failed: %w", err)
	}
	if err = s.PdcpStatusTransferUL.Encode(w); err != nil {
		return fmt.Errorf("encode PdcpStatusTransferUL failed: %w", err)
	}
	if err = s.PdcpStatusTransferDL.Encode(w); err != nil {
		return fmt.Errorf("encode PdcpStatusTransferDL failed: %w", err)
	}
	if s.IEExtension != nil {
		if err = s.IEExtension.Encode(w); err != nil {
			return fmt.Errorf("encode IEExtension failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *PDCPSNStatusInformation) Decode(r *aper.AperReader) (err error) {
	isExtensible, err := r.ReadBool()
	if err != nil {
		return fmt.Errorf("read extensibility bool failed: %w", err)
	}
	_ = isExtensible
	optionalityBitmap, _, err := r.ReadBitString(&aper.Constraint{Lb: 1, Ub: 1}, false)
	if err != nil {
		return fmt.Errorf("read optionality bitmap failed: %w", err)
	}
	if err = s.PdcpStatusTransferUL.Decode(r); err != nil {
		return fmt.Errorf("decode PdcpStatusTransferUL failed: %w", err)
	}
	if err = s.PdcpStatusTransferDL.Decode(r); err != nil {
		return fmt.Errorf("decode PdcpStatusTransferDL failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.IEExtension = new(ProtocolExtensionContainer)
		if err = s.IEExtension.Decode(r); err != nil {
			return fmt.Errorf("decode IEExtension failed: %w", err)
		}
	}
	if isExtensible { /* TODO: Implement extension skipping for PDCPSNStatusInformation */
	}
	return nil
}
