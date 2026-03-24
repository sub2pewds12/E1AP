package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/asn1go/per"
)

// DRBRequiredToModifyItemNGRAN is a generated SEQUENCE type.
type DRBRequiredToModifyItemNGRAN struct {
	DRBID                                DRBID
	GNBCUUPCellGroupRelatedConfiguration *GNBCUUPCellGroupRelatedConfiguration
	FlowToRemove                         *QOSFlowList
	Cause                                Cause
	IEExtensions                         *DRBRequiredToModifyItemNGRANExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *DRBRequiredToModifyItemNGRAN) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "dRB-ID", Optional: false},
			per.ComponentInfo{Name: "gNB-CU-UP-CellGroupRelatedConfiguration", Optional: true},
			per.ComponentInfo{Name: "flow-To-Remove", Optional: true},
			per.ComponentInfo{Name: "cause", Optional: false},
			per.ComponentInfo{Name: "iE-Extensions", Optional: true},
		},
	}
	seqEncoder := w.NewSequenceEncoder(c)
	if err := seqEncoder.EncodeExtensionBit(false); err != nil {
		return err
	}

	optionalBitmap := make([]bool, 0)

	if s.GNBCUUPCellGroupRelatedConfiguration != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.FlowToRemove != nil {
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

	if s.GNBCUUPCellGroupRelatedConfiguration != nil {
		if err = s.GNBCUUPCellGroupRelatedConfiguration.Encode(w); err != nil {
			return fmt.Errorf("encode GNBCUUPCellGroupRelatedConfiguration failed: %w", err)
		}
	}

	if s.FlowToRemove != nil {
		if err = s.FlowToRemove.Encode(w); err != nil {
			return fmt.Errorf("encode FlowToRemove failed: %w", err)
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
func (s *DRBRequiredToModifyItemNGRAN) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "dRB-ID", Optional: false},
			per.ComponentInfo{Name: "gNB-CU-UP-CellGroupRelatedConfiguration", Optional: true},
			per.ComponentInfo{Name: "flow-To-Remove", Optional: true},
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
		s.GNBCUUPCellGroupRelatedConfiguration = new(GNBCUUPCellGroupRelatedConfiguration)
		if err = s.GNBCUUPCellGroupRelatedConfiguration.Decode(r); err != nil {
			return fmt.Errorf("Decode GNBCUUPCellGroupRelatedConfiguration failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(2) {
		s.FlowToRemove = new(QOSFlowList)
		if err = s.FlowToRemove.Decode(r); err != nil {
			return fmt.Errorf("Decode FlowToRemove failed: %w", err)
		}
	}
	if err = s.Cause.Decode(r); err != nil {
		return fmt.Errorf("Decode Cause failed: %w", err)
	}

	if seqDecoder.IsComponentPresent(4) {
		s.IEExtensions = new(DRBRequiredToModifyItemNGRANExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
