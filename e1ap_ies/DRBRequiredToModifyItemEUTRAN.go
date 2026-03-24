package e1ap_ies

import (
	"fmt"

	"asn1go/per"
)

// DRBRequiredToModifyItemEUTRAN is a generated SEQUENCE type.
type DRBRequiredToModifyItemEUTRAN struct {
	DRBID                                DRBID
	S1DLUPTNLInformation                 *UPTNLInformation
	GNBCUUPCellGroupRelatedConfiguration *GNBCUUPCellGroupRelatedConfiguration
	Cause                                Cause
	IEExtensions                         *DRBRequiredToModifyItemEUTRANExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *DRBRequiredToModifyItemEUTRAN) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "dRB-ID", Optional: false},
			per.ComponentInfo{Name: "s1-DL-UP-TNL-Information", Optional: true},
			per.ComponentInfo{Name: "gNB-CU-UP-CellGroupRelatedConfiguration", Optional: true},
			per.ComponentInfo{Name: "cause", Optional: false},
			per.ComponentInfo{Name: "iE-Extensions", Optional: true},
		},
	}
	seqEncoder := w.NewSequenceEncoder(c)
	if err := seqEncoder.EncodeExtensionBit(false); err != nil {
		return err
	}

	optionalBitmap := make([]bool, 0)

	if s.S1DLUPTNLInformation != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.GNBCUUPCellGroupRelatedConfiguration != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

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

	if s.S1DLUPTNLInformation != nil {
		if err = s.S1DLUPTNLInformation.Encode(w); err != nil {
			return fmt.Errorf("encode S1DLUPTNLInformation failed: %w", err)
		}
	}

	if s.GNBCUUPCellGroupRelatedConfiguration != nil {
		if err = s.GNBCUUPCellGroupRelatedConfiguration.Encode(w); err != nil {
			return fmt.Errorf("encode GNBCUUPCellGroupRelatedConfiguration failed: %w", err)
		}
	}
	if err = s.Cause.Encode(w); err != nil {
		return fmt.Errorf("encode Cause failed: %w", err)
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
func (s *DRBRequiredToModifyItemEUTRAN) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "dRB-ID", Optional: false},
			per.ComponentInfo{Name: "s1-DL-UP-TNL-Information", Optional: true},
			per.ComponentInfo{Name: "gNB-CU-UP-CellGroupRelatedConfiguration", Optional: true},
			per.ComponentInfo{Name: "cause", Optional: false},
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

	if seqDecoder.IsComponentPresent(1) {
		s.S1DLUPTNLInformation = new(UPTNLInformation)
		if err = s.S1DLUPTNLInformation.Decode(r); err != nil {
			return fmt.Errorf("Decode S1DLUPTNLInformation failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(2) {
		s.GNBCUUPCellGroupRelatedConfiguration = new(GNBCUUPCellGroupRelatedConfiguration)
		if err = s.GNBCUUPCellGroupRelatedConfiguration.Decode(r); err != nil {
			return fmt.Errorf("Decode GNBCUUPCellGroupRelatedConfiguration failed: %w", err)
		}
	}
	if err = s.Cause.Decode(r); err != nil {
		return fmt.Errorf("Decode Cause failed: %w", err)
	}

	if seqDecoder.IsComponentPresent(4) {
		s.IEExtensions = new(DRBRequiredToModifyItemEUTRANExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
