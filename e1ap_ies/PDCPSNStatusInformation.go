package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/asn1go/per"
)

// PDCPSNStatusInformation is a generated SEQUENCE type.
type PDCPSNStatusInformation struct {
	PdcpStatusTransferUL DRBBStatusTransfer
	PdcpStatusTransferDL PDCPCount
	IEExtension          *PDCPSNStatusInformationExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *PDCPSNStatusInformation) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "pdcpStatusTransfer-UL", Optional: false},
			per.ComponentInfo{Name: "pdcpStatusTransfer-DL", Optional: false},
			per.ComponentInfo{Name: "iE-Extension", Optional: true},
		},
	}
	seqEncoder := w.NewSequenceEncoder(c)
	if err := seqEncoder.EncodeExtensionBit(false); err != nil {
		return err
	}

	optionalBitmap := make([]bool, 0)

	if s.IEExtension != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if err := seqEncoder.EncodePreamble(optionalBitmap); err != nil {
		return err
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

	if err := seqEncoder.EncodeExtensionAdditions([]bool{}, [][]byte{}); err != nil {
		return err
	}

	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *PDCPSNStatusInformation) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "pdcpStatusTransfer-UL", Optional: false},
			per.ComponentInfo{Name: "pdcpStatusTransfer-DL", Optional: false},
			per.ComponentInfo{Name: "iE-Extension", Optional: true},
		},
	}
	seqDecoder := r.NewSequenceDecoder(c)
	if err := seqDecoder.DecodeExtensionBit(); err != nil {
		return err
	}

	if err := seqDecoder.DecodePreamble(); err != nil {
		return err
	}

	if err = s.PdcpStatusTransferUL.Decode(r); err != nil {
		return fmt.Errorf("Decode PdcpStatusTransferUL failed: %w", err)
	}
	if err = s.PdcpStatusTransferDL.Decode(r); err != nil {
		return fmt.Errorf("Decode PdcpStatusTransferDL failed: %w", err)
	}

	if seqDecoder.IsComponentPresent(2) {
		s.IEExtension = new(PDCPSNStatusInformationExtensions)
		if err = s.IEExtension.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtension failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
