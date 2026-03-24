package e1ap_ies

import (
	"fmt"

	"asn1go/per"
)

// PDUSessionResourceToSetupItem is a generated SEQUENCE type.
type PDUSessionResourceToSetupItem struct {
	PDUSessionID                               PDUSessionID
	PDUSessionType                             PDUSessionType
	SNSSAI                                     SNSSAI
	SecurityIndication                         SecurityIndication
	PDUSessionResourceDLAMBR                   *BitRate
	NGULUPTNLInformation                       UPTNLInformation
	PDUSessionDataForwardingInformationRequest *DataForwardingInformationRequest
	PDUSessionInactivityTimer                  *InactivityTimer
	ExistingAllocatedNGDLUPTNLInfo             *UPTNLInformation
	NetworkInstance                            *NetworkInstance
	DRBToSetupListNGRAN                        DRBToSetupListNGRAN
	IEExtensions                               *PDUSessionResourceToSetupItemExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *PDUSessionResourceToSetupItem) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "pDU-Session-ID", Optional: false},
			per.ComponentInfo{Name: "pDU-Session-Type", Optional: false},
			per.ComponentInfo{Name: "sNSSAI", Optional: false},
			per.ComponentInfo{Name: "securityIndication", Optional: false},
			per.ComponentInfo{Name: "pDU-Session-Resource-DL-AMBR", Optional: true},
			per.ComponentInfo{Name: "nG-UL-UP-TNL-Information", Optional: false},
			per.ComponentInfo{Name: "pDU-Session-Data-Forwarding-Information-Request", Optional: true},
			per.ComponentInfo{Name: "pDU-Session-Inactivity-Timer", Optional: true},
			per.ComponentInfo{Name: "existing-Allocated-NG-DL-UP-TNL-Info", Optional: true},
			per.ComponentInfo{Name: "networkInstance", Optional: true},
			per.ComponentInfo{Name: "dRB-To-Setup-List-NG-RAN", Optional: false},
			per.ComponentInfo{Name: "iE-Extensions", Optional: true},
		},
	}
	seqEncoder := w.NewSequenceEncoder(c)
	if err := seqEncoder.EncodeExtensionBit(false); err != nil {
		return err
	}

	optionalBitmap := make([]bool, 0)

	if s.PDUSessionResourceDLAMBR != nil {
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

	if s.ExistingAllocatedNGDLUPTNLInfo != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.NetworkInstance != nil {
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

	if s.PDUSessionResourceDLAMBR != nil {
		if err = w.EncodeInteger(int64((*s.PDUSessionResourceDLAMBR).Value), per.ConstrainedExtensible(0, 4000000000000)); err != nil {
			return fmt.Errorf("encode PDUSessionResourceDLAMBR failed: %w", err)
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

	if s.ExistingAllocatedNGDLUPTNLInfo != nil {
		if err = s.ExistingAllocatedNGDLUPTNLInfo.Encode(w); err != nil {
			return fmt.Errorf("encode ExistingAllocatedNGDLUPTNLInfo failed: %w", err)
		}
	}

	if s.NetworkInstance != nil {
		if err = w.EncodeInteger(int64((*s.NetworkInstance).Value), per.ConstrainedExtensible(1, 256)); err != nil {
			return fmt.Errorf("encode NetworkInstance failed: %w", err)
		}
	}
	if err = s.DRBToSetupListNGRAN.Encode(w); err != nil {
		return fmt.Errorf("encode DRBToSetupListNGRAN failed: %w", err)
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
func (s *PDUSessionResourceToSetupItem) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "pDU-Session-ID", Optional: false},
			per.ComponentInfo{Name: "pDU-Session-Type", Optional: false},
			per.ComponentInfo{Name: "sNSSAI", Optional: false},
			per.ComponentInfo{Name: "securityIndication", Optional: false},
			per.ComponentInfo{Name: "pDU-Session-Resource-DL-AMBR", Optional: true},
			per.ComponentInfo{Name: "nG-UL-UP-TNL-Information", Optional: false},
			per.ComponentInfo{Name: "pDU-Session-Data-Forwarding-Information-Request", Optional: true},
			per.ComponentInfo{Name: "pDU-Session-Inactivity-Timer", Optional: true},
			per.ComponentInfo{Name: "existing-Allocated-NG-DL-UP-TNL-Info", Optional: true},
			per.ComponentInfo{Name: "networkInstance", Optional: true},
			per.ComponentInfo{Name: "dRB-To-Setup-List-NG-RAN", Optional: false},
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
		s.PDUSessionResourceDLAMBR = new(BitRate)

		{
			val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 4000000000000))
			if err != nil {
				return fmt.Errorf("decode PDUSessionResourceDLAMBR failed: %w", err)
			}
			s.PDUSessionResourceDLAMBR.Value = val
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

	if seqDecoder.IsComponentPresent(8) {
		s.ExistingAllocatedNGDLUPTNLInfo = new(UPTNLInformation)
		if err = s.ExistingAllocatedNGDLUPTNLInfo.Decode(r); err != nil {
			return fmt.Errorf("Decode ExistingAllocatedNGDLUPTNLInfo failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(9) {
		s.NetworkInstance = new(NetworkInstance)

		{
			val, err := r.DecodeInteger(per.ConstrainedExtensible(1, 256))
			if err != nil {
				return fmt.Errorf("decode NetworkInstance failed: %w", err)
			}
			s.NetworkInstance.Value = val
		}
	}
	if err = s.DRBToSetupListNGRAN.Decode(r); err != nil {
		return fmt.Errorf("Decode DRBToSetupListNGRAN failed: %w", err)
	}

	if seqDecoder.IsComponentPresent(11) {
		s.IEExtensions = new(PDUSessionResourceToSetupItemExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
