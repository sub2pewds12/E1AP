package e1ap_ies

import (
	"fmt"

	"asn1go/per"
)

// DRBToSetupItemEUTRAN is a generated SEQUENCE type.
type DRBToSetupItemEUTRAN struct {
	DRBID                            DRBID
	PDCPConfiguration                PDCPConfiguration
	EUTRANQOS                        EUTRANQOS
	S1ULUPTNLInformation             UPTNLInformation
	DataForwardingInformationRequest *DataForwardingInformationRequest
	CellGroupInformation             CellGroupInformation
	DLUPParameters                   *UPParameters
	DRBInactivityTimer               *InactivityTimer
	ExistingAllocatedS1DLUPTNLInfo   *UPTNLInformation
	IEExtensions                     *DRBToSetupItemEUTRANExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *DRBToSetupItemEUTRAN) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "dRB-ID", Optional: false},
			per.ComponentInfo{Name: "pDCP-Configuration", Optional: false},
			per.ComponentInfo{Name: "eUTRAN-QoS", Optional: false},
			per.ComponentInfo{Name: "s1-UL-UP-TNL-Information", Optional: false},
			per.ComponentInfo{Name: "data-Forwarding-Information-Request", Optional: true},
			per.ComponentInfo{Name: "cell-Group-Information", Optional: false},
			per.ComponentInfo{Name: "dL-UP-Parameters", Optional: true},
			per.ComponentInfo{Name: "dRB-Inactivity-Timer", Optional: true},
			per.ComponentInfo{Name: "existing-Allocated-S1-DL-UP-TNL-Info", Optional: true},
			per.ComponentInfo{Name: "iE-Extensions", Optional: true},
		},
	}
	seqEncoder := w.NewSequenceEncoder(c)
	if err := seqEncoder.EncodeExtensionBit(false); err != nil {
		return err
	}

	optionalBitmap := make([]bool, 0)

	if s.DataForwardingInformationRequest != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.DLUPParameters != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.DRBInactivityTimer != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.ExistingAllocatedS1DLUPTNLInfo != nil {
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
	if err = s.PDCPConfiguration.Encode(w); err != nil {
		return fmt.Errorf("encode PDCPConfiguration failed: %w", err)
	}
	if err = s.EUTRANQOS.Encode(w); err != nil {
		return fmt.Errorf("encode EUTRANQOS failed: %w", err)
	}
	if err = s.S1ULUPTNLInformation.Encode(w); err != nil {
		return fmt.Errorf("encode S1ULUPTNLInformation failed: %w", err)
	}

	if s.DataForwardingInformationRequest != nil {
		if err = s.DataForwardingInformationRequest.Encode(w); err != nil {
			return fmt.Errorf("encode DataForwardingInformationRequest failed: %w", err)
		}
	}
	if err = s.CellGroupInformation.Encode(w); err != nil {
		return fmt.Errorf("encode CellGroupInformation failed: %w", err)
	}

	if s.DLUPParameters != nil {
		if err = s.DLUPParameters.Encode(w); err != nil {
			return fmt.Errorf("encode DLUPParameters failed: %w", err)
		}
	}

	if s.DRBInactivityTimer != nil {
		if err = w.EncodeInteger(int64((*s.DRBInactivityTimer).Value), per.ConstrainedExtensible(1, 7200)); err != nil {
			return fmt.Errorf("encode DRBInactivityTimer failed: %w", err)
		}
	}

	if s.ExistingAllocatedS1DLUPTNLInfo != nil {
		if err = s.ExistingAllocatedS1DLUPTNLInfo.Encode(w); err != nil {
			return fmt.Errorf("encode ExistingAllocatedS1DLUPTNLInfo failed: %w", err)
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
func (s *DRBToSetupItemEUTRAN) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "dRB-ID", Optional: false},
			per.ComponentInfo{Name: "pDCP-Configuration", Optional: false},
			per.ComponentInfo{Name: "eUTRAN-QoS", Optional: false},
			per.ComponentInfo{Name: "s1-UL-UP-TNL-Information", Optional: false},
			per.ComponentInfo{Name: "data-Forwarding-Information-Request", Optional: true},
			per.ComponentInfo{Name: "cell-Group-Information", Optional: false},
			per.ComponentInfo{Name: "dL-UP-Parameters", Optional: true},
			per.ComponentInfo{Name: "dRB-Inactivity-Timer", Optional: true},
			per.ComponentInfo{Name: "existing-Allocated-S1-DL-UP-TNL-Info", Optional: true},
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
	if err = s.PDCPConfiguration.Decode(r); err != nil {
		return fmt.Errorf("Decode PDCPConfiguration failed: %w", err)
	}
	if err = s.EUTRANQOS.Decode(r); err != nil {
		return fmt.Errorf("Decode EUTRANQOS failed: %w", err)
	}
	if err = s.S1ULUPTNLInformation.Decode(r); err != nil {
		return fmt.Errorf("Decode S1ULUPTNLInformation failed: %w", err)
	}

	if seqDecoder.IsComponentPresent(4) {
		s.DataForwardingInformationRequest = new(DataForwardingInformationRequest)
		if err = s.DataForwardingInformationRequest.Decode(r); err != nil {
			return fmt.Errorf("Decode DataForwardingInformationRequest failed: %w", err)
		}
	}
	if err = s.CellGroupInformation.Decode(r); err != nil {
		return fmt.Errorf("Decode CellGroupInformation failed: %w", err)
	}

	if seqDecoder.IsComponentPresent(6) {
		s.DLUPParameters = new(UPParameters)
		if err = s.DLUPParameters.Decode(r); err != nil {
			return fmt.Errorf("Decode DLUPParameters failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(7) {
		s.DRBInactivityTimer = new(InactivityTimer)

		{
			val, err := r.DecodeInteger(per.ConstrainedExtensible(1, 7200))
			if err != nil {
				return fmt.Errorf("decode DRBInactivityTimer failed: %w", err)
			}
			s.DRBInactivityTimer.Value = val
		}
	}

	if seqDecoder.IsComponentPresent(8) {
		s.ExistingAllocatedS1DLUPTNLInfo = new(UPTNLInformation)
		if err = s.ExistingAllocatedS1DLUPTNLInfo.Decode(r); err != nil {
			return fmt.Errorf("Decode ExistingAllocatedS1DLUPTNLInfo failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(9) {
		s.IEExtensions = new(DRBToSetupItemEUTRANExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
