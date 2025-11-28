package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// DRBToSetupItemEUTRAN is a generated SEQUENCE type.
type DRBToSetupItemEUTRAN struct {
	DRBID                            DRBID                             `aper:"lb:1,ub:32,mandatory,ext"`
	PDCPConfiguration                PDCPConfiguration                 `aper:"mandatory,ext"`
	EUTRANQOS                        EUTRANQOS                         `aper:"mandatory,ext"`
	S1ULUPTNLInformation             UPTNLInformation                  `aper:"mandatory,ext"`
	DataForwardingInformationRequest *DataForwardingInformationRequest `aper:"optional,ext"`
	CellGroupInformation             CellGroupInformation              `aper:"lb:1,ub:MaxnoofCellGroups,mandatory,ext"`
	DLUPParameters                   *UPParameters                     `aper:"lb:1,ub:MaxnoofUPParameters,optional,ext"`
	DRBInactivityTimer               *InactivityTimer                  `aper:"lb:1,ub:7200,optional,ext"`
	ExistingAllocatedS1DLUPTNLInfo   *UPTNLInformation                 `aper:"optional,ext"`
	IEExtensions                     *ProtocolExtensionContainer       `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *DRBToSetupItemEUTRAN) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.DataForwardingInformationRequest != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.DLUPParameters != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if s.DRBInactivityTimer != nil {
		optionalityBitmap[0] |= 1 << 5
	}
	if s.ExistingAllocatedS1DLUPTNLInfo != nil {
		optionalityBitmap[0] |= 1 << 4
	}
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 3
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(5), &aper.Constraint{Lb: 5, Ub: 5}, false); err != nil {
		return fmt.Errorf("Encode optionality bitmap failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.DRBID.Value), &aper.Constraint{Lb: 1, Ub: 32}, true); err != nil {
		return fmt.Errorf("Encode DRBID failed: %w", err)
	}
	if err = s.PDCPConfiguration.Encode(w); err != nil {
		return fmt.Errorf("Encode PDCPConfiguration failed: %w", err)
	}
	if err = s.EUTRANQOS.Encode(w); err != nil {
		return fmt.Errorf("Encode EUTRANQOS failed: %w", err)
	}
	if err = s.S1ULUPTNLInformation.Encode(w); err != nil {
		return fmt.Errorf("Encode S1ULUPTNLInformation failed: %w", err)
	}
	if s.DataForwardingInformationRequest != nil {
		if err = s.DataForwardingInformationRequest.Encode(w); err != nil {
			return fmt.Errorf("Encode DataForwardingInformationRequest failed: %w", err)
		}
	}

	{
		itemPointers := make([]aper.AperMarshaller, len(s.CellGroupInformation.Value))
		for i := 0; i < len(s.CellGroupInformation.Value); i++ {
			itemPointers[i] = &(s.CellGroupInformation.Value[i])
		}
		if err = aper.WriteSequenceOf(itemPointers, w, &aper.Constraint{Lb: 1, Ub: MaxnoofCellGroups}, false); err != nil {
			return fmt.Errorf("Encode CellGroupInformation failed: %w", err)
		}
	}
	if s.DLUPParameters != nil {
		{
			itemPointers := make([]aper.AperMarshaller, len(s.DLUPParameters.Value))
			for i := 0; i < len(s.DLUPParameters.Value); i++ {
				itemPointers[i] = &(s.DLUPParameters.Value[i])
			}
			if err = aper.WriteSequenceOf(itemPointers, w, &aper.Constraint{Lb: 1, Ub: MaxnoofUPParameters}, false); err != nil {
				return fmt.Errorf("Encode DLUPParameters failed: %w", err)
			}
		}
	}
	if s.DRBInactivityTimer != nil {
		if err = w.WriteInteger(int64(s.DRBInactivityTimer.Value), &aper.Constraint{Lb: 1, Ub: 7200}, true); err != nil {
			return fmt.Errorf("Encode DRBInactivityTimer failed: %w", err)
		}
	}
	if s.ExistingAllocatedS1DLUPTNLInfo != nil {
		if err = s.ExistingAllocatedS1DLUPTNLInfo.Encode(w); err != nil {
			return fmt.Errorf("Encode ExistingAllocatedS1DLUPTNLInfo failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *DRBToSetupItemEUTRAN) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 5, Ub: 5}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}
	if err = s.DRBID.Decode(r); err != nil {
		return fmt.Errorf("Decode DRBID failed: %w", err)
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
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.DataForwardingInformationRequest = new(DataForwardingInformationRequest)
		if err = s.DataForwardingInformationRequest.Decode(r); err != nil {
			return fmt.Errorf("Decode DataForwardingInformationRequest failed: %w", err)
		}
	}
	if err = s.CellGroupInformation.Decode(r); err != nil {
		return fmt.Errorf("Decode CellGroupInformation failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.DLUPParameters = new(UPParameters)
		if err = s.DLUPParameters.Decode(r); err != nil {
			return fmt.Errorf("Decode DLUPParameters failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<5) > 0 {
		s.DRBInactivityTimer = new(InactivityTimer)
		if err = s.DRBInactivityTimer.Decode(r); err != nil {
			return fmt.Errorf("Decode DRBInactivityTimer failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<4) > 0 {
		s.ExistingAllocatedS1DLUPTNLInfo = new(UPTNLInformation)
		if err = s.ExistingAllocatedS1DLUPTNLInfo.Decode(r); err != nil {
			return fmt.Errorf("Decode ExistingAllocatedS1DLUPTNLInfo failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<3) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible { /* TODO: Implement extension skipping for DRBToSetupItemEUTRAN */
	}
	return nil
}
