package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// DRBToSetupItemNGRAN is a generated SEQUENCE type.
type DRBToSetupItemNGRAN struct {
	DRBID                               DRBID                             `aper:"lb:1,ub:32,mandatory,ext"`
	SDAPConfiguration                   SDAPConfiguration                 `aper:"mandatory,ext"`
	PDCPConfiguration                   PDCPConfiguration                 `aper:"mandatory,ext"`
	CellGroupInformation                CellGroupInformation              `aper:"mandatory,ext"`
	QOSFlowInformationToBeSetup         QOSFlowQOSParameterList           `aper:"mandatory,ext"`
	DRBDataForwardingInformationRequest *DataForwardingInformationRequest `aper:"optional,ext"`
	DRBInactivityTimer                  *InactivityTimer                  `aper:"lb:1,ub:7200,optional,ext"`
	PDCPSNStatusInformation             *PDCPSNStatusInformation          `aper:"optional,ext"`
	IEExtensions                        *DRBToSetupItemNGRANExtensions    `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *DRBToSetupItemNGRAN) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.DRBDataForwardingInformationRequest != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.DRBInactivityTimer != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if s.PDCPSNStatusInformation != nil {
		optionalityBitmap[0] |= 1 << 5
	}
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 4
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(4), &aper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
		return fmt.Errorf("Encode optionality bitmap failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.DRBID.Value), &aper.Constraint{Lb: 1, Ub: 32}, true); err != nil {
		return fmt.Errorf("Encode DRBID failed: %w", err)
	}
	if err = s.SDAPConfiguration.Encode(w); err != nil {
		return fmt.Errorf("Encode SDAPConfiguration failed: %w", err)
	}
	if err = s.PDCPConfiguration.Encode(w); err != nil {
		return fmt.Errorf("Encode PDCPConfiguration failed: %w", err)
	}

	{
		itemPointers := make([]aper.AperMarshaller, len(s.CellGroupInformation.Value))
		for i := 0; i < len(s.CellGroupInformation.Value); i++ {
			itemPointers[i] = &(s.CellGroupInformation.Value[i])
		}
		if err = aper.WriteSequenceOf(itemPointers, w, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return fmt.Errorf("Encode CellGroupInformation failed: %w", err)
		}
	}

	{
		itemPointers := make([]aper.AperMarshaller, len(s.QOSFlowInformationToBeSetup.Value))
		for i := 0; i < len(s.QOSFlowInformationToBeSetup.Value); i++ {
			itemPointers[i] = &(s.QOSFlowInformationToBeSetup.Value[i])
		}
		if err = aper.WriteSequenceOf(itemPointers, w, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return fmt.Errorf("Encode QOSFlowInformationToBeSetup failed: %w", err)
		}
	}
	if s.DRBDataForwardingInformationRequest != nil {
		if err = s.DRBDataForwardingInformationRequest.Encode(w); err != nil {
			return fmt.Errorf("Encode DRBDataForwardingInformationRequest failed: %w", err)
		}
	}
	if s.DRBInactivityTimer != nil {
		if err = w.WriteInteger(int64(s.DRBInactivityTimer.Value), &aper.Constraint{Lb: 1, Ub: 7200}, true); err != nil {
			return fmt.Errorf("Encode DRBInactivityTimer failed: %w", err)
		}
	}
	if s.PDCPSNStatusInformation != nil {
		if err = s.PDCPSNStatusInformation.Encode(w); err != nil {
			return fmt.Errorf("Encode PDCPSNStatusInformation failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *DRBToSetupItemNGRAN) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}
	if err = s.DRBID.Decode(r); err != nil {
		return fmt.Errorf("Decode DRBID failed: %w", err)
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
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.DRBDataForwardingInformationRequest = new(DataForwardingInformationRequest)
		if err = s.DRBDataForwardingInformationRequest.Decode(r); err != nil {
			return fmt.Errorf("Decode DRBDataForwardingInformationRequest failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.DRBInactivityTimer = new(InactivityTimer)
		if err = s.DRBInactivityTimer.Decode(r); err != nil {
			return fmt.Errorf("Decode DRBInactivityTimer failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<5) > 0 {
		s.PDCPSNStatusInformation = new(PDCPSNStatusInformation)
		if err = s.PDCPSNStatusInformation.Decode(r); err != nil {
			return fmt.Errorf("Decode PDCPSNStatusInformation failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<4) > 0 {
		s.IEExtensions = new(DRBToSetupItemNGRANExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible { /* TODO: Implement extension skipping for DRBToSetupItemNGRAN */
	}
	return nil
}
