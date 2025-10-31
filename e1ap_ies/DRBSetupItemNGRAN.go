package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// DRBSetupItemNGRAN is a generated SEQUENCE type.
type DRBSetupItemNGRAN struct {
	DRBID                                DRBID                       `aper:"lb:1,ub:32,mandatory,ext"`
	DRBDataForwardingInformationResponse *DataForwardingInformation  `aper:"optional,ext"`
	ULUPTransportParameters              []UPParametersItem          `aper:"mandatory,ext"`
	FlowSetupList                        []QOSFlowItem               `aper:"mandatory,ext"`
	FlowFailedList                       []QOSFlowFailedItem         `aper:"optional,ext"`
	IEExtensions                         *ProtocolExtensionContainer `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *DRBSetupItemNGRAN) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.DRBDataForwardingInformationResponse != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.FlowFailedList != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 5
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(3), &aper.Constraint{Lb: 3, Ub: 3}, false); err != nil {
		return fmt.Errorf("Encode optionality bitmap failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.DRBID), &aper.Constraint{Lb: 1, Ub: 32}, true); err != nil {
		return fmt.Errorf("Encode DRBID failed: %w", err)
	}
	if s.DRBDataForwardingInformationResponse != nil {
		if err = s.DRBDataForwardingInformationResponse.Encode(w); err != nil {
			return fmt.Errorf("Encode DRBDataForwardingInformationResponse failed: %w", err)
		}
	}
	if err = s.ULUPTransportParameters.Encode(w); err != nil {
		return fmt.Errorf("Encode ULUPTransportParameters failed: %w", err)
	}
	if err = s.FlowSetupList.Encode(w); err != nil {
		return fmt.Errorf("Encode FlowSetupList failed: %w", err)
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
func (s *DRBSetupItemNGRAN) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 3, Ub: 3}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}

	{
		var val int64
		if val, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 32}, true); err != nil {
			return fmt.Errorf("Decode DRBID failed: %w", err)
		}
		s.DRBID = DRBID(val)
	}

	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.DRBDataForwardingInformationResponse = new(DataForwardingInformation)
		if err = s.DRBDataForwardingInformationResponse.Decode(r); err != nil {
			return fmt.Errorf("Decode DRBDataForwardingInformationResponse failed: %w", err)
		}
	}
	if err = s.ULUPTransportParameters.Decode(r); err != nil {
		return fmt.Errorf("Decode ULUPTransportParameters failed: %w", err)
	}
	if err = s.FlowSetupList.Decode(r); err != nil {
		return fmt.Errorf("Decode FlowSetupList failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.FlowFailedList = new(QOSFlowFailedList)
		if err = s.FlowFailedList.Decode(r); err != nil {
			return fmt.Errorf("Decode FlowFailedList failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<5) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for DRBSetupItemNGRAN")
	}
	return nil
}
