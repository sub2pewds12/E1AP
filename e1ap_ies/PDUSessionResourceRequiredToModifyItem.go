package e1ap_ies

import (
	"fmt"

	"asn1go/per"
)

// PDUSessionResourceRequiredToModifyItem is a generated SEQUENCE type.
type PDUSessionResourceRequiredToModifyItem struct {
	PDUSessionID                 PDUSessionID
	NGDLUPTNLInformation         *UPTNLInformation
	DRBRequiredToModifyListNGRAN *DRBRequiredToModifyListNGRAN
	DRBRequiredToRemoveListNGRAN *DRBRequiredToRemoveListNGRAN
	IEExtensions                 *PDUSessionResourceRequiredToModifyItemExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *PDUSessionResourceRequiredToModifyItem) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "pDU-Session-ID", Optional: false},
			per.ComponentInfo{Name: "nG-DL-UP-TNL-Information", Optional: true},
			per.ComponentInfo{Name: "dRB-Required-To-Modify-List-NG-RAN", Optional: true},
			per.ComponentInfo{Name: "dRB-Required-To-Remove-List-NG-RAN", Optional: true},
			per.ComponentInfo{Name: "iE-Extensions", Optional: true},
		},
	}
	seqEncoder := w.NewSequenceEncoder(c)
	if err := seqEncoder.EncodeExtensionBit(false); err != nil {
		return err
	}

	optionalBitmap := make([]bool, 0)

	if s.NGDLUPTNLInformation != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.DRBRequiredToModifyListNGRAN != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.DRBRequiredToRemoveListNGRAN != nil {
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

	if err = w.EncodeInteger(int64(s.PDUSessionID.Value), per.Constrained(0, 255)); err != nil {
		return fmt.Errorf("encode PDUSessionID failed: %w", err)
	}

	if s.NGDLUPTNLInformation != nil {
		if err = s.NGDLUPTNLInformation.Encode(w); err != nil {
			return fmt.Errorf("encode NGDLUPTNLInformation failed: %w", err)
		}
	}

	if s.DRBRequiredToModifyListNGRAN != nil {
		if err = s.DRBRequiredToModifyListNGRAN.Encode(w); err != nil {
			return fmt.Errorf("encode DRBRequiredToModifyListNGRAN failed: %w", err)
		}
	}

	if s.DRBRequiredToRemoveListNGRAN != nil {
		if err = s.DRBRequiredToRemoveListNGRAN.Encode(w); err != nil {
			return fmt.Errorf("encode DRBRequiredToRemoveListNGRAN failed: %w", err)
		}
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
func (s *PDUSessionResourceRequiredToModifyItem) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "pDU-Session-ID", Optional: false},
			per.ComponentInfo{Name: "nG-DL-UP-TNL-Information", Optional: true},
			per.ComponentInfo{Name: "dRB-Required-To-Modify-List-NG-RAN", Optional: true},
			per.ComponentInfo{Name: "dRB-Required-To-Remove-List-NG-RAN", Optional: true},
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
		val, err := r.DecodeInteger(per.Constrained(0, 255))
		if err != nil {
			return fmt.Errorf("decode PDUSessionID failed: %w", err)
		}
		s.PDUSessionID.Value = val
	}

	if seqDecoder.IsComponentPresent(1) {
		s.NGDLUPTNLInformation = new(UPTNLInformation)
		if err = s.NGDLUPTNLInformation.Decode(r); err != nil {
			return fmt.Errorf("Decode NGDLUPTNLInformation failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(2) {
		s.DRBRequiredToModifyListNGRAN = new(DRBRequiredToModifyListNGRAN)
		if err = s.DRBRequiredToModifyListNGRAN.Decode(r); err != nil {
			return fmt.Errorf("Decode DRBRequiredToModifyListNGRAN failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(3) {
		s.DRBRequiredToRemoveListNGRAN = new(DRBRequiredToRemoveListNGRAN)
		if err = s.DRBRequiredToRemoveListNGRAN.Decode(r); err != nil {
			return fmt.Errorf("Decode DRBRequiredToRemoveListNGRAN failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(4) {
		s.IEExtensions = new(PDUSessionResourceRequiredToModifyItemExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
