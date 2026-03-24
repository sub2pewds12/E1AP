package e1ap_ies

import (
	"fmt"

	"asn1go/per"
)

// PDUSessionResourceSetupModItem is a generated SEQUENCE type.
type PDUSessionResourceSetupModItem struct {
	PDUSessionID                                PDUSessionID
	SecurityResult                              *SecurityResult
	NGDLUPTNLInformation                        UPTNLInformation
	PDUSessionDataForwardingInformationResponse *DataForwardingInformation
	DRBSetupModListNGRAN                        DRBSetupModListNGRAN
	DRBFailedModListNGRAN                       *DRBFailedModListNGRAN
	IEExtensions                                *PDUSessionResourceSetupModItemExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *PDUSessionResourceSetupModItem) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "pDU-Session-ID", Optional: false},
			per.ComponentInfo{Name: "securityResult", Optional: true},
			per.ComponentInfo{Name: "nG-DL-UP-TNL-Information", Optional: false},
			per.ComponentInfo{Name: "pDU-Session-Data-Forwarding-Information-Response", Optional: true},
			per.ComponentInfo{Name: "dRB-Setup-Mod-List-NG-RAN", Optional: false},
			per.ComponentInfo{Name: "dRB-Failed-Mod-List-NG-RAN", Optional: true},
			per.ComponentInfo{Name: "iE-Extensions", Optional: true},
		},
	}
	seqEncoder := w.NewSequenceEncoder(c)
	if err := seqEncoder.EncodeExtensionBit(false); err != nil {
		return err
	}

	optionalBitmap := make([]bool, 0)

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

	if s.DRBFailedModListNGRAN != nil {
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

	if s.SecurityResult != nil {
		if err = s.SecurityResult.Encode(w); err != nil {
			return fmt.Errorf("encode SecurityResult failed: %w", err)
		}
	}
	if err = s.NGDLUPTNLInformation.Encode(w); err != nil {
		return fmt.Errorf("encode NGDLUPTNLInformation failed: %w", err)
	}

	if s.PDUSessionDataForwardingInformationResponse != nil {
		if err = s.PDUSessionDataForwardingInformationResponse.Encode(w); err != nil {
			return fmt.Errorf("encode PDUSessionDataForwardingInformationResponse failed: %w", err)
		}
	}
	if err = s.DRBSetupModListNGRAN.Encode(w); err != nil {
		return fmt.Errorf("encode DRBSetupModListNGRAN failed: %w", err)
	}

	if s.DRBFailedModListNGRAN != nil {
		if err = s.DRBFailedModListNGRAN.Encode(w); err != nil {
			return fmt.Errorf("encode DRBFailedModListNGRAN failed: %w", err)
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
func (s *PDUSessionResourceSetupModItem) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "pDU-Session-ID", Optional: false},
			per.ComponentInfo{Name: "securityResult", Optional: true},
			per.ComponentInfo{Name: "nG-DL-UP-TNL-Information", Optional: false},
			per.ComponentInfo{Name: "pDU-Session-Data-Forwarding-Information-Response", Optional: true},
			per.ComponentInfo{Name: "dRB-Setup-Mod-List-NG-RAN", Optional: false},
			per.ComponentInfo{Name: "dRB-Failed-Mod-List-NG-RAN", Optional: true},
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
		s.SecurityResult = new(SecurityResult)
		if err = s.SecurityResult.Decode(r); err != nil {
			return fmt.Errorf("Decode SecurityResult failed: %w", err)
		}
	}
	if err = s.NGDLUPTNLInformation.Decode(r); err != nil {
		return fmt.Errorf("Decode NGDLUPTNLInformation failed: %w", err)
	}

	if seqDecoder.IsComponentPresent(3) {
		s.PDUSessionDataForwardingInformationResponse = new(DataForwardingInformation)
		if err = s.PDUSessionDataForwardingInformationResponse.Decode(r); err != nil {
			return fmt.Errorf("Decode PDUSessionDataForwardingInformationResponse failed: %w", err)
		}
	}
	if err = s.DRBSetupModListNGRAN.Decode(r); err != nil {
		return fmt.Errorf("Decode DRBSetupModListNGRAN failed: %w", err)
	}

	if seqDecoder.IsComponentPresent(5) {
		s.DRBFailedModListNGRAN = new(DRBFailedModListNGRAN)
		if err = s.DRBFailedModListNGRAN.Decode(r); err != nil {
			return fmt.Errorf("Decode DRBFailedModListNGRAN failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(6) {
		s.IEExtensions = new(PDUSessionResourceSetupModItemExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
