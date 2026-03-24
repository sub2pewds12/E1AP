package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/asn1go/per"
)

// DRBToSetupItemNGRAN is a generated SEQUENCE type.
type DRBToSetupItemNGRAN struct {
	DRBID                               DRBID
	SDAPConfiguration                   SDAPConfiguration
	PDCPConfiguration                   PDCPConfiguration
	CellGroupInformation                CellGroupInformation
	QOSFlowInformationToBeSetup         QOSFlowQOSParameterList
	DRBDataForwardingInformationRequest *DataForwardingInformationRequest
	DRBInactivityTimer                  *InactivityTimer
	PDCPSNStatusInformation             *PDCPSNStatusInformation
	IEExtensions                        *DRBToSetupItemNGRANExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *DRBToSetupItemNGRAN) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "dRB-ID", Optional: false},
			per.ComponentInfo{Name: "sDAP-Configuration", Optional: false},
			per.ComponentInfo{Name: "pDCP-Configuration", Optional: false},
			per.ComponentInfo{Name: "cell-Group-Information", Optional: false},
			per.ComponentInfo{Name: "qos-flow-Information-To-Be-Setup", Optional: false},
			per.ComponentInfo{Name: "dRB-Data-Forwarding-Information-Request", Optional: true},
			per.ComponentInfo{Name: "dRB-Inactivity-Timer", Optional: true},
			per.ComponentInfo{Name: "pDCP-SN-Status-Information", Optional: true},
			per.ComponentInfo{Name: "iE-Extensions", Optional: true},
		},
	}
	seqEncoder := w.NewSequenceEncoder(c)
	if err := seqEncoder.EncodeExtensionBit(false); err != nil {
		return err
	}

	optionalBitmap := make([]bool, 0)

	if s.DRBDataForwardingInformationRequest != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.DRBInactivityTimer != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.PDCPSNStatusInformation != nil {
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
	if err = s.SDAPConfiguration.Encode(w); err != nil {
		return fmt.Errorf("encode SDAPConfiguration failed: %w", err)
	}
	if err = s.PDCPConfiguration.Encode(w); err != nil {
		return fmt.Errorf("encode PDCPConfiguration failed: %w", err)
	}
	if err = s.CellGroupInformation.Encode(w); err != nil {
		return fmt.Errorf("encode CellGroupInformation failed: %w", err)
	}
	if err = s.QOSFlowInformationToBeSetup.Encode(w); err != nil {
		return fmt.Errorf("encode QOSFlowInformationToBeSetup failed: %w", err)
	}

	if s.DRBDataForwardingInformationRequest != nil {
		if err = s.DRBDataForwardingInformationRequest.Encode(w); err != nil {
			return fmt.Errorf("encode DRBDataForwardingInformationRequest failed: %w", err)
		}
	}

	if s.DRBInactivityTimer != nil {
		if err = w.EncodeInteger(int64((*s.DRBInactivityTimer).Value), per.ConstrainedExtensible(1, 7200)); err != nil {
			return fmt.Errorf("encode DRBInactivityTimer failed: %w", err)
		}
	}

	if s.PDCPSNStatusInformation != nil {
		if err = s.PDCPSNStatusInformation.Encode(w); err != nil {
			return fmt.Errorf("encode PDCPSNStatusInformation failed: %w", err)
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
func (s *DRBToSetupItemNGRAN) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "dRB-ID", Optional: false},
			per.ComponentInfo{Name: "sDAP-Configuration", Optional: false},
			per.ComponentInfo{Name: "pDCP-Configuration", Optional: false},
			per.ComponentInfo{Name: "cell-Group-Information", Optional: false},
			per.ComponentInfo{Name: "qos-flow-Information-To-Be-Setup", Optional: false},
			per.ComponentInfo{Name: "dRB-Data-Forwarding-Information-Request", Optional: true},
			per.ComponentInfo{Name: "dRB-Inactivity-Timer", Optional: true},
			per.ComponentInfo{Name: "pDCP-SN-Status-Information", Optional: true},
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
	if err = s.SDAPConfiguration.Decode(r); err != nil {
		return fmt.Errorf("Decode SDAPConfiguration failed: %w", err)
	}
	if err = s.PDCPConfiguration.Decode(r); err != nil {
		return fmt.Errorf("Decode PDCPConfiguration failed: %w", err)
	}
	if err = s.CellGroupInformation.Decode(r); err != nil {
		return fmt.Errorf("Decode CellGroupInformation failed: %w", err)
	}
	if err = s.QOSFlowInformationToBeSetup.Decode(r); err != nil {
		return fmt.Errorf("Decode QOSFlowInformationToBeSetup failed: %w", err)
	}

	if seqDecoder.IsComponentPresent(5) {
		s.DRBDataForwardingInformationRequest = new(DataForwardingInformationRequest)
		if err = s.DRBDataForwardingInformationRequest.Decode(r); err != nil {
			return fmt.Errorf("Decode DRBDataForwardingInformationRequest failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(6) {
		s.DRBInactivityTimer = new(InactivityTimer)

		{
			val, err := r.DecodeInteger(per.ConstrainedExtensible(1, 7200))
			if err != nil {
				return fmt.Errorf("decode DRBInactivityTimer failed: %w", err)
			}
			s.DRBInactivityTimer.Value = val
		}
	}

	if seqDecoder.IsComponentPresent(7) {
		s.PDCPSNStatusInformation = new(PDCPSNStatusInformation)
		if err = s.PDCPSNStatusInformation.Decode(r); err != nil {
			return fmt.Errorf("Decode PDCPSNStatusInformation failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(8) {
		s.IEExtensions = new(DRBToSetupItemNGRANExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
