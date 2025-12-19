package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// DRBToSetupModItemNGRAN is a generated SEQUENCE type.
type DRBToSetupModItemNGRAN struct {
	DRBID                               DRBID                             `aper:"lb:1,ub:32,mandatory,ext"`
	SDAPConfiguration                   SDAPConfiguration                 `aper:"mandatory,ext"`
	PDCPConfiguration                   PDCPConfiguration                 `aper:"mandatory,ext"`
	CellGroupInformation                CellGroupInformation              `aper:"lb:1,ub:MaxnoofCellGroups,mandatory,ext"`
	FlowMappingInformation              QOSFlowQOSParameterList           `aper:"lb:1,ub:MaxnoofQoSFlows,mandatory,ext"`
	DRBDataForwardingInformationRequest *DataForwardingInformationRequest `aper:"optional,ext"`
	DRBInactivityTimer                  *InactivityTimer                  `aper:"lb:1,ub:7200,optional,ext"`
	PDCPSNStatusInformation             *PDCPSNStatusInformation          `aper:"optional,ext"`
	IEExtensions                        *DRBToSetupModItemNGRANExtensions `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *DRBToSetupModItemNGRAN) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("encode extensibility bool failed: %w", err)
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
		return fmt.Errorf("encode optionality bitmap failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.DRBID.Value), &aper.Constraint{Lb: 1, Ub: 32}, true); err != nil {
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
	if err = s.FlowMappingInformation.Encode(w); err != nil {
		return fmt.Errorf("encode FlowMappingInformation failed: %w", err)
	}
	if s.DRBDataForwardingInformationRequest != nil {
		if err = s.DRBDataForwardingInformationRequest.Encode(w); err != nil {
			return fmt.Errorf("encode DRBDataForwardingInformationRequest failed: %w", err)
		}
	}
	if s.DRBInactivityTimer != nil {
		if err = w.WriteInteger(int64((*s.DRBInactivityTimer).Value), &aper.Constraint{Lb: 1, Ub: 7200}, true); err != nil {
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
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *DRBToSetupModItemNGRAN) Decode(r *aper.AperReader) (err error) {
	isExtensible, err := r.ReadBool()
	if err != nil {
		return fmt.Errorf("read extensibility bool failed: %w", err)
	}
	_ = isExtensible
	optionalityBitmap, _, err := r.ReadBitString(&aper.Constraint{Lb: 4, Ub: 4}, false)
	if err != nil {
		return fmt.Errorf("read optionality bitmap failed: %w", err)
	}
	if err = s.DRBID.Decode(r); err != nil {
		return fmt.Errorf("decode DRBID failed: %w", err)
	}
	if err = s.SDAPConfiguration.Decode(r); err != nil {
		return fmt.Errorf("decode SDAPConfiguration failed: %w", err)
	}
	if err = s.PDCPConfiguration.Decode(r); err != nil {
		return fmt.Errorf("decode PDCPConfiguration failed: %w", err)
	}
	if err = s.CellGroupInformation.Decode(r); err != nil {
		return fmt.Errorf("decode CellGroupInformation failed: %w", err)
	}
	if err = s.FlowMappingInformation.Decode(r); err != nil {
		return fmt.Errorf("decode FlowMappingInformation failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.DRBDataForwardingInformationRequest = new(DataForwardingInformationRequest)
		if err = s.DRBDataForwardingInformationRequest.Decode(r); err != nil {
			return fmt.Errorf("decode DRBDataForwardingInformationRequest failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.DRBInactivityTimer = new(InactivityTimer)
		if err = s.DRBInactivityTimer.Decode(r); err != nil {
			return fmt.Errorf("decode DRBInactivityTimer failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<5) > 0 {
		s.PDCPSNStatusInformation = new(PDCPSNStatusInformation)
		if err = s.PDCPSNStatusInformation.Decode(r); err != nil {
			return fmt.Errorf("decode PDCPSNStatusInformation failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<4) > 0 {
		s.IEExtensions = new(DRBToSetupModItemNGRANExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible { /* TODO: Implement extension skipping for DRBToSetupModItemNGRAN */
	}
	return nil
}
