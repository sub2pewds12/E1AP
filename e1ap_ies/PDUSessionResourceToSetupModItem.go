package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/asn1go/per"
)

// PDUSessionResourceToSetupModItem is a generated SEQUENCE type.
type PDUSessionResourceToSetupModItem struct {
	PDUSessionID                               PDUSessionID
	PDUSessionType                             PDUSessionType
	SNSSAI                                     SNSSAI
	SecurityIndication                         SecurityIndication
	PDUSessionResourceAMBR                     *BitRate
	NGULUPTNLInformation                       UPTNLInformation
	PDUSessionDataForwardingInformationRequest *DataForwardingInformationRequest
	PDUSessionInactivityTimer                  *InactivityTimer
	DRBToSetupModListNGRAN                     DRBToSetupModListNGRAN
	IEExtensions                               *PDUSessionResourceToSetupModItemExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *PDUSessionResourceToSetupModItem) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "pDU-Session-ID", Optional: false},
			per.ComponentInfo{Name: "pDU-Session-Type", Optional: false},
			per.ComponentInfo{Name: "sNSSAI", Optional: false},
			per.ComponentInfo{Name: "securityIndication", Optional: false},
			per.ComponentInfo{Name: "pDU-Session-Resource-AMBR", Optional: true},
			per.ComponentInfo{Name: "nG-UL-UP-TNL-Information", Optional: false},
			per.ComponentInfo{Name: "pDU-Session-Data-Forwarding-Information-Request", Optional: true},
			per.ComponentInfo{Name: "pDU-Session-Inactivity-Timer", Optional: true},
			per.ComponentInfo{Name: "dRB-To-Setup-Mod-List-NG-RAN", Optional: false},
			per.ComponentInfo{Name: "iE-Extensions", Optional: true},
		},
	}
	seqEncoder := w.NewSequenceEncoder(c)
	if err := seqEncoder.EncodeExtensionBit(false); err != nil {
		return err
	}

	optionalBitmap := make([]bool, 0)

	if s.PDUSessionResourceAMBR != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.PDUSessionDataForwardingInformationRequest != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.PDUSessionInactivityTimer != nil {
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

	{
		enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 5), ExtValues: nil}
		if err = w.EncodeEnumerated(int64(s.PDUSessionType.Value), enumC); err != nil {
			return fmt.Errorf("encode PDUSessionType failed: %w", err)
		}
	}
	if err = s.SNSSAI.Encode(w); err != nil {
		return fmt.Errorf("encode SNSSAI failed: %w", err)
	}
	if err = s.SecurityIndication.Encode(w); err != nil {
		return fmt.Errorf("encode SecurityIndication failed: %w", err)
	}

	if s.PDUSessionResourceAMBR != nil {
		if err = w.EncodeInteger(int64((*s.PDUSessionResourceAMBR).Value), per.ConstrainedExtensible(0, 4000000000000)); err != nil {
			return fmt.Errorf("encode PDUSessionResourceAMBR failed: %w", err)
		}
	}
	if err = s.NGULUPTNLInformation.Encode(w); err != nil {
		return fmt.Errorf("encode NGULUPTNLInformation failed: %w", err)
	}

	if s.PDUSessionDataForwardingInformationRequest != nil {
		if err = s.PDUSessionDataForwardingInformationRequest.Encode(w); err != nil {
			return fmt.Errorf("encode PDUSessionDataForwardingInformationRequest failed: %w", err)
		}
	}

	if s.PDUSessionInactivityTimer != nil {
		if err = w.EncodeInteger(int64((*s.PDUSessionInactivityTimer).Value), per.ConstrainedExtensible(1, 7200)); err != nil {
			return fmt.Errorf("encode PDUSessionInactivityTimer failed: %w", err)
		}
	}
	if err = s.DRBToSetupModListNGRAN.Encode(w); err != nil {
		return fmt.Errorf("encode DRBToSetupModListNGRAN failed: %w", err)
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
func (s *PDUSessionResourceToSetupModItem) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "pDU-Session-ID", Optional: false},
			per.ComponentInfo{Name: "pDU-Session-Type", Optional: false},
			per.ComponentInfo{Name: "sNSSAI", Optional: false},
			per.ComponentInfo{Name: "securityIndication", Optional: false},
			per.ComponentInfo{Name: "pDU-Session-Resource-AMBR", Optional: true},
			per.ComponentInfo{Name: "nG-UL-UP-TNL-Information", Optional: false},
			per.ComponentInfo{Name: "pDU-Session-Data-Forwarding-Information-Request", Optional: true},
			per.ComponentInfo{Name: "pDU-Session-Inactivity-Timer", Optional: true},
			per.ComponentInfo{Name: "dRB-To-Setup-Mod-List-NG-RAN", Optional: false},
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

	{
		enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 5), ExtValues: nil}
		val, err := r.DecodeEnumerated(enumC)
		if err != nil {
			return fmt.Errorf("decode PDUSessionType failed: %w", err)
		}
		s.PDUSessionType.Value = val
	}
	if err = s.SNSSAI.Decode(r); err != nil {
		return fmt.Errorf("Decode SNSSAI failed: %w", err)
	}
	if err = s.SecurityIndication.Decode(r); err != nil {
		return fmt.Errorf("Decode SecurityIndication failed: %w", err)
	}

	if seqDecoder.IsComponentPresent(4) {
		s.PDUSessionResourceAMBR = new(BitRate)

		{
			val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 4000000000000))
			if err != nil {
				return fmt.Errorf("decode PDUSessionResourceAMBR failed: %w", err)
			}
			s.PDUSessionResourceAMBR.Value = val
		}
	}
	if err = s.NGULUPTNLInformation.Decode(r); err != nil {
		return fmt.Errorf("Decode NGULUPTNLInformation failed: %w", err)
	}

	if seqDecoder.IsComponentPresent(6) {
		s.PDUSessionDataForwardingInformationRequest = new(DataForwardingInformationRequest)
		if err = s.PDUSessionDataForwardingInformationRequest.Decode(r); err != nil {
			return fmt.Errorf("Decode PDUSessionDataForwardingInformationRequest failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(7) {
		s.PDUSessionInactivityTimer = new(InactivityTimer)

		{
			val, err := r.DecodeInteger(per.ConstrainedExtensible(1, 7200))
			if err != nil {
				return fmt.Errorf("decode PDUSessionInactivityTimer failed: %w", err)
			}
			s.PDUSessionInactivityTimer.Value = val
		}
	}
	if err = s.DRBToSetupModListNGRAN.Decode(r); err != nil {
		return fmt.Errorf("Decode DRBToSetupModListNGRAN failed: %w", err)
	}

	if seqDecoder.IsComponentPresent(9) {
		s.IEExtensions = new(PDUSessionResourceToSetupModItemExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
