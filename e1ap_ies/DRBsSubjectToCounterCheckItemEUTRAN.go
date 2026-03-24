package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/asn1go/per"
)

// DRBsSubjectToCounterCheckItemEUTRAN is a generated SEQUENCE type.
type DRBsSubjectToCounterCheckItemEUTRAN struct {
	DRBID        DRBID
	PDCPULCount  PDCPCount
	PDCPDLCount  PDCPCount
	IEExtensions *DRBsSubjectToCounterCheckItemEUTRANExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *DRBsSubjectToCounterCheckItemEUTRAN) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "dRB-ID", Optional: false},
			per.ComponentInfo{Name: "pDCP-UL-Count", Optional: false},
			per.ComponentInfo{Name: "pDCP-DL-Count", Optional: false},
			per.ComponentInfo{Name: "iE-Extensions", Optional: true},
		},
	}
	seqEncoder := w.NewSequenceEncoder(c)
	if err := seqEncoder.EncodeExtensionBit(false); err != nil {
		return err
	}

	optionalBitmap := make([]bool, 0)

	if s.IEExtensions != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if err := seqEncoder.EncodePreamble(optionalBitmap); err != nil {
		return err
	}

	if err = w.EncodeInteger(int64(s.DRBID.Value), per.ConstrainedExtensible(1, 32)); err != nil {
		return fmt.Errorf("encode DRBID failed: %w", err)
	}
	if err = s.PDCPULCount.Encode(w); err != nil {
		return fmt.Errorf("encode PDCPULCount failed: %w", err)
	}
	if err = s.PDCPDLCount.Encode(w); err != nil {
		return fmt.Errorf("encode PDCPDLCount failed: %w", err)
	}

	if s.IEExtensions != nil {
		if err = s.IEExtensions.Encode(w); err != nil {
			return fmt.Errorf("encode IEExtensions failed: %w", err)
		}
	}

	if err := seqEncoder.EncodeExtensionAdditions([]bool{}, [][]byte{}); err != nil {
		return err
	}

	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *DRBsSubjectToCounterCheckItemEUTRAN) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "dRB-ID", Optional: false},
			per.ComponentInfo{Name: "pDCP-UL-Count", Optional: false},
			per.ComponentInfo{Name: "pDCP-DL-Count", Optional: false},
			per.ComponentInfo{Name: "iE-Extensions", Optional: true},
		},
	}
	seqDecoder := r.NewSequenceDecoder(c)
	if err := seqDecoder.DecodeExtensionBit(); err != nil {
		return err
	}

	if err := seqDecoder.DecodePreamble(); err != nil {
		return err
	}

	{
		val, err := r.DecodeInteger(per.ConstrainedExtensible(1, 32))
		if err != nil {
			return fmt.Errorf("decode DRBID failed: %w", err)
		}
		s.DRBID.Value = val
	}
	if err = s.PDCPULCount.Decode(r); err != nil {
		return fmt.Errorf("Decode PDCPULCount failed: %w", err)
	}
	if err = s.PDCPDLCount.Decode(r); err != nil {
		return fmt.Errorf("Decode PDCPDLCount failed: %w", err)
	}

	if seqDecoder.IsComponentPresent(3) {
		s.IEExtensions = new(DRBsSubjectToCounterCheckItemEUTRANExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
