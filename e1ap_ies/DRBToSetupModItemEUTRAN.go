package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// DRBToSetupModItemEUTRAN is a generated SEQUENCE type.
type DRBToSetupModItemEUTRAN struct {
	DRBID                            DRBID                             `aper:"lb:1,ub:32,mandatory,ext"`
	PDCPConfiguration                PDCPConfiguration                 `aper:"mandatory,ext"`
	EUTRANQOS                        EUTRANQOS                         `aper:"mandatory,ext"`
	S1ULUPTNLInformation             UPTNLInformation                  `aper:"mandatory,ext"`
	DataForwardingInformationRequest *DataForwardingInformationRequest `aper:"optional,ext"`
	CellGroupInformation             []CellGroupInformationItem        `aper:"mandatory,ext"`
	DLUPParameters                   []UPParametersItem                `aper:"optional,ext"`
	DRBInactivityTimer               *InactivityTimer                  `aper:"lb:1,ub:7200,optional,ext"`
	IEExtensions                     *ProtocolExtensionContainer       `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *DRBToSetupModItemEUTRAN) Encode(w *aper.AperWriter) (err error) {
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
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 4
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(4), &aper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
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
	if err = s.CellGroupInformation.Encode(w); err != nil {
		return fmt.Errorf("Encode CellGroupInformation failed: %w", err)
	}
	if s.DLUPParameters != nil {
		if err = s.DLUPParameters.Encode(w); err != nil {
			return fmt.Errorf("Encode DLUPParameters failed: %w", err)
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
func (s *DRBToSetupModItemEUTRAN) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}

	{
		var val int64
		if val, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 32}, true); err != nil {
			return fmt.Errorf("Decode DRBID failed: %w", err)
		}
		s.DRBID.Value = aper.Integer(val)
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

		{
			var val int64
			if val, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 7200}, true); err != nil {
				return fmt.Errorf("Decode DRBInactivityTimer failed: %w", err)
			}
			s.DRBInactivityTimer = new(InactivityTimer)
			s.DRBInactivityTimer.Value = aper.Integer(val)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<4) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for DRBToSetupModItemEUTRAN")
	}
	return nil
}
