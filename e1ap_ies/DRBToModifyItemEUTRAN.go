package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/asn1go/per"
)

// DRBToModifyItemEUTRAN is a generated SEQUENCE type.
type DRBToModifyItemEUTRAN struct {
	DRBID                     DRBID
	PDCPConfiguration         *PDCPConfiguration
	EUTRANQOS                 *EUTRANQOS
	S1ULUPTNLInformation      *UPTNLInformation
	DataForwardingInformation *DataForwardingInformation
	PDCPSNStatusRequest       *PDCPSNStatusRequest
	PDCPSNStatusInformation   *PDCPSNStatusInformation
	DLUPParameters            *UPParameters
	CellGroupToAdd            *CellGroupInformation
	CellGroupToModify         *CellGroupInformation
	CellGroupToRemove         *CellGroupInformation
	DRBInactivityTimer        *InactivityTimer
	IEExtensions              *DRBToModifyItemEUTRANExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *DRBToModifyItemEUTRAN) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "dRB-ID", Optional: false},
			per.ComponentInfo{Name: "pDCP-Configuration", Optional: true},
			per.ComponentInfo{Name: "eUTRAN-QoS", Optional: true},
			per.ComponentInfo{Name: "s1-UL-UP-TNL-Information", Optional: true},
			per.ComponentInfo{Name: "data-Forwarding-Information", Optional: true},
			per.ComponentInfo{Name: "pDCP-SN-Status-Request", Optional: true},
			per.ComponentInfo{Name: "pDCP-SN-Status-Information", Optional: true},
			per.ComponentInfo{Name: "dL-UP-Parameters", Optional: true},
			per.ComponentInfo{Name: "cell-Group-To-Add", Optional: true},
			per.ComponentInfo{Name: "cell-Group-To-Modify", Optional: true},
			per.ComponentInfo{Name: "cell-Group-To-Remove", Optional: true},
			per.ComponentInfo{Name: "dRB-Inactivity-Timer", Optional: true},
			per.ComponentInfo{Name: "iE-Extensions", Optional: true},
		},
	}
	seqEncoder := w.NewSequenceEncoder(c)
	if err := seqEncoder.EncodeExtensionBit(false); err != nil {
		return err
	}

	optionalBitmap := make([]bool, 0)

	if s.PDCPConfiguration != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.EUTRANQOS != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.S1ULUPTNLInformation != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.DataForwardingInformation != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.PDCPSNStatusRequest != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.PDCPSNStatusInformation != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.DLUPParameters != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.CellGroupToAdd != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.CellGroupToModify != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.CellGroupToRemove != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.DRBInactivityTimer != nil {
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

	if s.PDCPConfiguration != nil {
		if err = s.PDCPConfiguration.Encode(w); err != nil {
			return fmt.Errorf("encode PDCPConfiguration failed: %w", err)
		}
	}

	if s.EUTRANQOS != nil {
		if err = s.EUTRANQOS.Encode(w); err != nil {
			return fmt.Errorf("encode EUTRANQOS failed: %w", err)
		}
	}

	if s.S1ULUPTNLInformation != nil {
		if err = s.S1ULUPTNLInformation.Encode(w); err != nil {
			return fmt.Errorf("encode S1ULUPTNLInformation failed: %w", err)
		}
	}

	if s.DataForwardingInformation != nil {
		if err = s.DataForwardingInformation.Encode(w); err != nil {
			return fmt.Errorf("encode DataForwardingInformation failed: %w", err)
		}
	}

	if s.PDCPSNStatusRequest != nil {

		{
			enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 1), ExtValues: nil}
			if err = w.EncodeEnumerated(int64((*s.PDCPSNStatusRequest).Value), enumC); err != nil {
				return fmt.Errorf("encode PDCPSNStatusRequest failed: %w", err)
			}
		}
	}

	if s.PDCPSNStatusInformation != nil {
		if err = s.PDCPSNStatusInformation.Encode(w); err != nil {
			return fmt.Errorf("encode PDCPSNStatusInformation failed: %w", err)
		}
	}

	if s.DLUPParameters != nil {
		if err = s.DLUPParameters.Encode(w); err != nil {
			return fmt.Errorf("encode DLUPParameters failed: %w", err)
		}
	}

	if s.CellGroupToAdd != nil {
		if err = s.CellGroupToAdd.Encode(w); err != nil {
			return fmt.Errorf("encode CellGroupToAdd failed: %w", err)
		}
	}

	if s.CellGroupToModify != nil {
		if err = s.CellGroupToModify.Encode(w); err != nil {
			return fmt.Errorf("encode CellGroupToModify failed: %w", err)
		}
	}

	if s.CellGroupToRemove != nil {
		if err = s.CellGroupToRemove.Encode(w); err != nil {
			return fmt.Errorf("encode CellGroupToRemove failed: %w", err)
		}
	}

	if s.DRBInactivityTimer != nil {
		if err = w.EncodeInteger(int64((*s.DRBInactivityTimer).Value), per.ConstrainedExtensible(1, 7200)); err != nil {
			return fmt.Errorf("encode DRBInactivityTimer failed: %w", err)
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
func (s *DRBToModifyItemEUTRAN) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "dRB-ID", Optional: false},
			per.ComponentInfo{Name: "pDCP-Configuration", Optional: true},
			per.ComponentInfo{Name: "eUTRAN-QoS", Optional: true},
			per.ComponentInfo{Name: "s1-UL-UP-TNL-Information", Optional: true},
			per.ComponentInfo{Name: "data-Forwarding-Information", Optional: true},
			per.ComponentInfo{Name: "pDCP-SN-Status-Request", Optional: true},
			per.ComponentInfo{Name: "pDCP-SN-Status-Information", Optional: true},
			per.ComponentInfo{Name: "dL-UP-Parameters", Optional: true},
			per.ComponentInfo{Name: "cell-Group-To-Add", Optional: true},
			per.ComponentInfo{Name: "cell-Group-To-Modify", Optional: true},
			per.ComponentInfo{Name: "cell-Group-To-Remove", Optional: true},
			per.ComponentInfo{Name: "dRB-Inactivity-Timer", Optional: true},
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
		s.PDCPConfiguration = new(PDCPConfiguration)
		if err = s.PDCPConfiguration.Decode(r); err != nil {
			return fmt.Errorf("Decode PDCPConfiguration failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(2) {
		s.EUTRANQOS = new(EUTRANQOS)
		if err = s.EUTRANQOS.Decode(r); err != nil {
			return fmt.Errorf("Decode EUTRANQOS failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(3) {
		s.S1ULUPTNLInformation = new(UPTNLInformation)
		if err = s.S1ULUPTNLInformation.Decode(r); err != nil {
			return fmt.Errorf("Decode S1ULUPTNLInformation failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(4) {
		s.DataForwardingInformation = new(DataForwardingInformation)
		if err = s.DataForwardingInformation.Decode(r); err != nil {
			return fmt.Errorf("Decode DataForwardingInformation failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(5) {
		s.PDCPSNStatusRequest = new(PDCPSNStatusRequest)

		{
			enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 1), ExtValues: nil}
			val, err := r.DecodeEnumerated(enumC)
			if err != nil {
				return fmt.Errorf("decode PDCPSNStatusRequest failed: %w", err)
			}
			s.PDCPSNStatusRequest.Value = val
		}
	}

	if seqDecoder.IsComponentPresent(6) {
		s.PDCPSNStatusInformation = new(PDCPSNStatusInformation)
		if err = s.PDCPSNStatusInformation.Decode(r); err != nil {
			return fmt.Errorf("Decode PDCPSNStatusInformation failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(7) {
		s.DLUPParameters = new(UPParameters)
		if err = s.DLUPParameters.Decode(r); err != nil {
			return fmt.Errorf("Decode DLUPParameters failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(8) {
		s.CellGroupToAdd = new(CellGroupInformation)
		if err = s.CellGroupToAdd.Decode(r); err != nil {
			return fmt.Errorf("Decode CellGroupToAdd failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(9) {
		s.CellGroupToModify = new(CellGroupInformation)
		if err = s.CellGroupToModify.Decode(r); err != nil {
			return fmt.Errorf("Decode CellGroupToModify failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(10) {
		s.CellGroupToRemove = new(CellGroupInformation)
		if err = s.CellGroupToRemove.Decode(r); err != nil {
			return fmt.Errorf("Decode CellGroupToRemove failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(11) {
		s.DRBInactivityTimer = new(InactivityTimer)

		{
			val, err := r.DecodeInteger(per.ConstrainedExtensible(1, 7200))
			if err != nil {
				return fmt.Errorf("decode DRBInactivityTimer failed: %w", err)
			}
			s.DRBInactivityTimer.Value = val
		}
	}

	if seqDecoder.IsComponentPresent(12) {
		s.IEExtensions = new(DRBToModifyItemEUTRANExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
