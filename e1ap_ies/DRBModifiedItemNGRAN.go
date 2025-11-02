package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// DRBModifiedItemNGRAN is a generated SEQUENCE type.
type DRBModifiedItemNGRAN struct {
	DRBID                   DRBID                           `aper:"lb:1,ub:32,mandatory,ext"`
	ULUPTransportParameters []UPParametersItem              `aper:"optional,ext"`
	PDCPSNStatusInformation *PDCPSNStatusInformation        `aper:"optional,ext"`
	FlowSetupList           []QOSFlowItem                   `aper:"optional,ext"`
	FlowFailedList          []QOSFlowFailedItem             `aper:"optional,ext"`
	IEExtensions            *DRBModifiedItemNGRANExtensions `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *DRBModifiedItemNGRAN) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.ULUPTransportParameters != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.PDCPSNStatusInformation != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if s.FlowSetupList != nil {
		optionalityBitmap[0] |= 1 << 5
	}
	if s.FlowFailedList != nil {
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
	if s.ULUPTransportParameters != nil {
		if err = s.ULUPTransportParameters.Encode(w); err != nil {
			return fmt.Errorf("Encode ULUPTransportParameters failed: %w", err)
		}
	}
	if s.PDCPSNStatusInformation != nil {
		if err = s.PDCPSNStatusInformation.Encode(w); err != nil {
			return fmt.Errorf("Encode PDCPSNStatusInformation failed: %w", err)
		}
	}
	if s.FlowSetupList != nil {
		if err = s.FlowSetupList.Encode(w); err != nil {
			return fmt.Errorf("Encode FlowSetupList failed: %w", err)
		}
	}
	if s.FlowFailedList != nil {
		if err = s.FlowFailedList.Encode(w); err != nil {
			return fmt.Errorf("Encode FlowFailedList failed: %w", err)
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
func (s *DRBModifiedItemNGRAN) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 5, Ub: 5}, false); err != nil {
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
		s.ULUPTransportParameters = new(UPParameters)
		if err = s.ULUPTransportParameters.Decode(r); err != nil {
			return fmt.Errorf("Decode ULUPTransportParameters failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.PDCPSNStatusInformation = new(PDCPSNStatusInformation)
		if err = s.PDCPSNStatusInformation.Decode(r); err != nil {
			return fmt.Errorf("Decode PDCPSNStatusInformation failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<5) > 0 {
		s.FlowSetupList = new(QOSFlowList)
		if err = s.FlowSetupList.Decode(r); err != nil {
			return fmt.Errorf("Decode FlowSetupList failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<4) > 0 {
		s.FlowFailedList = new(QOSFlowFailedList)
		if err = s.FlowFailedList.Decode(r); err != nil {
			return fmt.Errorf("Decode FlowFailedList failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<3) > 0 {
		s.IEExtensions = new(DRBModifiedItemNGRANExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for DRBModifiedItemNGRAN")
	}
	return nil
}
