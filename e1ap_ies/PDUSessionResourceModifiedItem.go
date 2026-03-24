package e1ap_ies

import (
	"fmt"

	"asn1go/per"
)

// PDUSessionResourceModifiedItem is a generated SEQUENCE type.
type PDUSessionResourceModifiedItem struct {
	PDUSessionID                                PDUSessionID
	NGDLUPTNLInformation                        *UPTNLInformation
	SecurityResult                              *SecurityResult
	PDUSessionDataForwardingInformationResponse *DataForwardingInformation
	DRBSetupListNGRAN                           *DRBSetupListNGRAN
	DRBFailedListNGRAN                          *DRBFailedListNGRAN
	DRBModifiedListNGRAN                        *DRBModifiedListNGRAN
	DRBFailedToModifyListNGRAN                  *DRBFailedToModifyListNGRAN
	IEExtensions                                *PDUSessionResourceModifiedItemExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *PDUSessionResourceModifiedItem) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "pDU-Session-ID", Optional: false},
			per.ComponentInfo{Name: "nG-DL-UP-TNL-Information", Optional: true},
			per.ComponentInfo{Name: "securityResult", Optional: true},
			per.ComponentInfo{Name: "pDU-Session-Data-Forwarding-Information-Response", Optional: true},
			per.ComponentInfo{Name: "dRB-Setup-List-NG-RAN", Optional: true},
			per.ComponentInfo{Name: "dRB-Failed-List-NG-RAN", Optional: true},
			per.ComponentInfo{Name: "dRB-Modified-List-NG-RAN", Optional: true},
			per.ComponentInfo{Name: "dRB-Failed-To-Modify-List-NG-RAN", Optional: true},
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

	if s.SecurityResult != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.PDUSessionDataForwardingInformationResponse != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.DRBSetupListNGRAN != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.DRBFailedListNGRAN != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.DRBModifiedListNGRAN != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.DRBFailedToModifyListNGRAN != nil {
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

	if s.SecurityResult != nil {
		if err = s.SecurityResult.Encode(w); err != nil {
			return fmt.Errorf("encode SecurityResult failed: %w", err)
		}
	}

	if s.PDUSessionDataForwardingInformationResponse != nil {
		if err = s.PDUSessionDataForwardingInformationResponse.Encode(w); err != nil {
			return fmt.Errorf("encode PDUSessionDataForwardingInformationResponse failed: %w", err)
		}
	}

	if s.DRBSetupListNGRAN != nil {
		if err = s.DRBSetupListNGRAN.Encode(w); err != nil {
			return fmt.Errorf("encode DRBSetupListNGRAN failed: %w", err)
		}
	}

	if s.DRBFailedListNGRAN != nil {
		if err = s.DRBFailedListNGRAN.Encode(w); err != nil {
			return fmt.Errorf("encode DRBFailedListNGRAN failed: %w", err)
		}
	}

	if s.DRBModifiedListNGRAN != nil {
		if err = s.DRBModifiedListNGRAN.Encode(w); err != nil {
			return fmt.Errorf("encode DRBModifiedListNGRAN failed: %w", err)
		}
	}

	if s.DRBFailedToModifyListNGRAN != nil {
		if err = s.DRBFailedToModifyListNGRAN.Encode(w); err != nil {
			return fmt.Errorf("encode DRBFailedToModifyListNGRAN failed: %w", err)
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
func (s *PDUSessionResourceModifiedItem) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "pDU-Session-ID", Optional: false},
			per.ComponentInfo{Name: "nG-DL-UP-TNL-Information", Optional: true},
			per.ComponentInfo{Name: "securityResult", Optional: true},
			per.ComponentInfo{Name: "pDU-Session-Data-Forwarding-Information-Response", Optional: true},
			per.ComponentInfo{Name: "dRB-Setup-List-NG-RAN", Optional: true},
			per.ComponentInfo{Name: "dRB-Failed-List-NG-RAN", Optional: true},
			per.ComponentInfo{Name: "dRB-Modified-List-NG-RAN", Optional: true},
			per.ComponentInfo{Name: "dRB-Failed-To-Modify-List-NG-RAN", Optional: true},
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
		s.SecurityResult = new(SecurityResult)
		if err = s.SecurityResult.Decode(r); err != nil {
			return fmt.Errorf("Decode SecurityResult failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(3) {
		s.PDUSessionDataForwardingInformationResponse = new(DataForwardingInformation)
		if err = s.PDUSessionDataForwardingInformationResponse.Decode(r); err != nil {
			return fmt.Errorf("Decode PDUSessionDataForwardingInformationResponse failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(4) {
		s.DRBSetupListNGRAN = new(DRBSetupListNGRAN)
		if err = s.DRBSetupListNGRAN.Decode(r); err != nil {
			return fmt.Errorf("Decode DRBSetupListNGRAN failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(5) {
		s.DRBFailedListNGRAN = new(DRBFailedListNGRAN)
		if err = s.DRBFailedListNGRAN.Decode(r); err != nil {
			return fmt.Errorf("Decode DRBFailedListNGRAN failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(6) {
		s.DRBModifiedListNGRAN = new(DRBModifiedListNGRAN)
		if err = s.DRBModifiedListNGRAN.Decode(r); err != nil {
			return fmt.Errorf("Decode DRBModifiedListNGRAN failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(7) {
		s.DRBFailedToModifyListNGRAN = new(DRBFailedToModifyListNGRAN)
		if err = s.DRBFailedToModifyListNGRAN.Decode(r); err != nil {
			return fmt.Errorf("Decode DRBFailedToModifyListNGRAN failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(8) {
		s.IEExtensions = new(PDUSessionResourceModifiedItemExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
