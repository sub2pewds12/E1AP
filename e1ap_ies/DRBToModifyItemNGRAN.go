package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// DRBToModifyItemNGRAN is a generated SEQUENCE type.
type DRBToModifyItemNGRAN struct {
	DRBID                        DRBID                           `aper:"lb:1,ub:32,mandatory,ext"`
	SDAPConfiguration            *SDAPConfiguration              `aper:"optional,ext"`
	PDCPConfiguration            *PDCPConfiguration              `aper:"optional,ext"`
	DRBDataForwardingInformation *DataForwardingInformation      `aper:"optional,ext"`
	PDCPSNStatusRequest          *PDCPSNStatusRequest            `aper:"optional,ext"`
	PdcpSNStatusInformation      *PDCPSNStatusInformation        `aper:"optional,ext"`
	DLUPParameters               []UPParametersItem              `aper:"optional,ext"`
	CellGroupToAdd               []CellGroupInformationItem      `aper:"optional,ext"`
	CellGroupToModify            []CellGroupInformationItem      `aper:"optional,ext"`
	CellGroupToRemove            []CellGroupInformationItem      `aper:"optional,ext"`
	FlowMappingInformation       []QOSFlowQOSParameterItem       `aper:"optional,ext"`
	DRBInactivityTimer           *InactivityTimer                `aper:"lb:1,ub:7200,optional,ext"`
	IEExtensions                 *DRBToModifyItemNGRANExtensions `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *DRBToModifyItemNGRAN) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [2]byte
	if s.SDAPConfiguration != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.PDCPConfiguration != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if s.DRBDataForwardingInformation != nil {
		optionalityBitmap[0] |= 1 << 5
	}
	if s.PDCPSNStatusRequest != nil {
		optionalityBitmap[0] |= 1 << 4
	}
	if s.PdcpSNStatusInformation != nil {
		optionalityBitmap[0] |= 1 << 3
	}
	if s.DLUPParameters != nil {
		optionalityBitmap[0] |= 1 << 2
	}
	if s.CellGroupToAdd != nil {
		optionalityBitmap[0] |= 1 << 1
	}
	if s.CellGroupToModify != nil {
		optionalityBitmap[0] |= 1 << 0
	}
	if s.CellGroupToRemove != nil {
		optionalityBitmap[1] |= 1 << 7
	}
	if s.FlowMappingInformation != nil {
		optionalityBitmap[1] |= 1 << 6
	}
	if s.DRBInactivityTimer != nil {
		optionalityBitmap[1] |= 1 << 5
	}
	if s.IEExtensions != nil {
		optionalityBitmap[1] |= 1 << 4
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(12), &aper.Constraint{Lb: 12, Ub: 12}, false); err != nil {
		return fmt.Errorf("Encode optionality bitmap failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.DRBID.Value), &aper.Constraint{Lb: 1, Ub: 32}, true); err != nil {
		return fmt.Errorf("Encode DRBID failed: %w", err)
	}
	if s.SDAPConfiguration != nil {
		if err = s.SDAPConfiguration.Encode(w); err != nil {
			return fmt.Errorf("Encode SDAPConfiguration failed: %w", err)
		}
	}
	if s.PDCPConfiguration != nil {
		if err = s.PDCPConfiguration.Encode(w); err != nil {
			return fmt.Errorf("Encode PDCPConfiguration failed: %w", err)
		}
	}
	if s.DRBDataForwardingInformation != nil {
		if err = s.DRBDataForwardingInformation.Encode(w); err != nil {
			return fmt.Errorf("Encode DRBDataForwardingInformation failed: %w", err)
		}
	}
	if s.PDCPSNStatusRequest != nil {
		if err = w.WriteEnumerate(uint64(s.PDCPSNStatusRequest.Value), aper.Constraint{Lb: 0, Ub: 0}, true); err != nil {
			return fmt.Errorf("Encode PDCPSNStatusRequest failed: %w", err)
		}
	}
	if s.PdcpSNStatusInformation != nil {
		if err = s.PdcpSNStatusInformation.Encode(w); err != nil {
			return fmt.Errorf("Encode PdcpSNStatusInformation failed: %w", err)
		}
	}
	if s.DLUPParameters != nil {
		if err = s.DLUPParameters.Encode(w); err != nil {
			return fmt.Errorf("Encode DLUPParameters failed: %w", err)
		}
	}
	if s.CellGroupToAdd != nil {
		if err = s.CellGroupToAdd.Encode(w); err != nil {
			return fmt.Errorf("Encode CellGroupToAdd failed: %w", err)
		}
	}
	if s.CellGroupToModify != nil {
		if err = s.CellGroupToModify.Encode(w); err != nil {
			return fmt.Errorf("Encode CellGroupToModify failed: %w", err)
		}
	}
	if s.CellGroupToRemove != nil {
		if err = s.CellGroupToRemove.Encode(w); err != nil {
			return fmt.Errorf("Encode CellGroupToRemove failed: %w", err)
		}
	}
	if s.FlowMappingInformation != nil {
		if err = s.FlowMappingInformation.Encode(w); err != nil {
			return fmt.Errorf("Encode FlowMappingInformation failed: %w", err)
		}
	}
	if s.DRBInactivityTimer != nil {
		if err = w.WriteInteger(int64(s.DRBInactivityTimer.Value), &aper.Constraint{Lb: 1, Ub: 7200}, true); err != nil {
			return fmt.Errorf("Encode DRBInactivityTimer failed: %w", err)
		}
	}
	if s.IEExtensions != nil {
		if err = s.IEExtensions.Encode(w); err != nil {
			return fmt.Errorf("Encode IEExtensions failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *DRBToModifyItemNGRAN) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 12, Ub: 12}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}

	{
		var val int64
		if val, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 32}, true); err != nil {
			return fmt.Errorf("Decode DRBID failed: %w", err)
		}
		s.DRBID.Value = aper.Integer(val)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.SDAPConfiguration = new(SDAPConfiguration)
		if err = s.SDAPConfiguration.Decode(r); err != nil {
			return fmt.Errorf("Decode SDAPConfiguration failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.PDCPConfiguration = new(PDCPConfiguration)
		if err = s.PDCPConfiguration.Decode(r); err != nil {
			return fmt.Errorf("Decode PDCPConfiguration failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<5) > 0 {
		s.DRBDataForwardingInformation = new(DataForwardingInformation)
		if err = s.DRBDataForwardingInformation.Decode(r); err != nil {
			return fmt.Errorf("Decode DRBDataForwardingInformation failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<4) > 0 {
		s.PDCPSNStatusRequest = new(PDCPSNStatusRequest)

		{
			var val uint64
			if val, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, true); err != nil {
				return fmt.Errorf("Decode PDCPSNStatusRequest failed: %w", err)
			}
			s.PDCPSNStatusRequest.Value = aper.Enumerated(val)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<3) > 0 {
		s.PdcpSNStatusInformation = new(PDCPSNStatusInformation)
		if err = s.PdcpSNStatusInformation.Decode(r); err != nil {
			return fmt.Errorf("Decode PdcpSNStatusInformation failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<2) > 0 {
		s.DLUPParameters = new(UPParameters)
		if err = s.DLUPParameters.Decode(r); err != nil {
			return fmt.Errorf("Decode DLUPParameters failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<1) > 0 {
		s.CellGroupToAdd = new(CellGroupInformation)
		if err = s.CellGroupToAdd.Decode(r); err != nil {
			return fmt.Errorf("Decode CellGroupToAdd failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<0) > 0 {
		s.CellGroupToModify = new(CellGroupInformation)
		if err = s.CellGroupToModify.Decode(r); err != nil {
			return fmt.Errorf("Decode CellGroupToModify failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 1 && optionalityBitmap[1]&(1<<7) > 0 {
		s.CellGroupToRemove = new(CellGroupInformation)
		if err = s.CellGroupToRemove.Decode(r); err != nil {
			return fmt.Errorf("Decode CellGroupToRemove failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 1 && optionalityBitmap[1]&(1<<6) > 0 {
		s.FlowMappingInformation = new(QOSFlowQOSParameterList)
		if err = s.FlowMappingInformation.Decode(r); err != nil {
			return fmt.Errorf("Decode FlowMappingInformation failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 1 && optionalityBitmap[1]&(1<<5) > 0 {

		{
			var val int64
			if val, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 7200}, true); err != nil {
				return fmt.Errorf("Decode DRBInactivityTimer failed: %w", err)
			}
			s.DRBInactivityTimer = new(InactivityTimer)
			s.DRBInactivityTimer.Value = aper.Integer(val)
		}
	}
	if len(optionalityBitmap) > 1 && optionalityBitmap[1]&(1<<4) > 0 {
		s.IEExtensions = new(DRBToModifyItemNGRANExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for DRBToModifyItemNGRAN")
	}
	return nil
}
