package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// DRBSetupItemNGRAN is a generated SEQUENCE type.
type DRBSetupItemNGRAN struct {
	DRBID                                DRBID                       `aper:"lb:1,ub:32,mandatory,ext"`
	DRBDataForwardingInformationResponse *DataForwardingInformation  `aper:"optional,ext"`
	ULUPTransportParameters              UPParameters                `aper:"lb:1,ub:MaxnoofUPParameters,mandatory,ext"`
	FlowSetupList                        QOSFlowList                 `aper:"lb:1,ub:MaxnoofQoSFlows,mandatory,ext"`
	FlowFailedList                       *QOSFlowFailedList          `aper:"lb:1,ub:MaxnoofQoSFlows,optional,ext"`
	IEExtensions                         *ProtocolExtensionContainer `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *DRBSetupItemNGRAN) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("encode extensibility bool failed: %w", err)
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
		return fmt.Errorf("encode optionality bitmap failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.DRBID.Value), &aper.Constraint{Lb: 1, Ub: 32}, true); err != nil {
		return fmt.Errorf("encode DRBID failed: %w", err)
	}
	if s.DRBDataForwardingInformationResponse != nil {
		if err = s.DRBDataForwardingInformationResponse.Encode(w); err != nil {
			return fmt.Errorf("encode DRBDataForwardingInformationResponse failed: %w", err)
		}
	}
	if err = s.ULUPTransportParameters.Encode(w); err != nil {
		return fmt.Errorf("encode ULUPTransportParameters failed: %w", err)
	}
	if err = s.FlowSetupList.Encode(w); err != nil {
		return fmt.Errorf("encode FlowSetupList failed: %w", err)
	}
	if s.FlowFailedList != nil {
		if err = s.FlowFailedList.Encode(w); err != nil {
			return fmt.Errorf("encode FlowFailedList failed: %w", err)
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
func (s *DRBSetupItemNGRAN) Decode(r *aper.AperReader) (err error) {
	isExtensible, err := r.ReadBool()
	if err != nil {
		return fmt.Errorf("read extensibility bool failed: %w", err)
	}
	_ = isExtensible
	optionalityBitmap, _, err := r.ReadBitString(&aper.Constraint{Lb: 3, Ub: 3}, false)
	if err != nil {
		return fmt.Errorf("read optionality bitmap failed: %w", err)
	}
	if err = s.DRBID.Decode(r); err != nil {
		return fmt.Errorf("decode DRBID failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.DRBDataForwardingInformationResponse = new(DataForwardingInformation)
		if err = s.DRBDataForwardingInformationResponse.Decode(r); err != nil {
			return fmt.Errorf("decode DRBDataForwardingInformationResponse failed: %w", err)
		}
	}
	if err = s.ULUPTransportParameters.Decode(r); err != nil {
		return fmt.Errorf("decode ULUPTransportParameters failed: %w", err)
	}
	if err = s.FlowSetupList.Decode(r); err != nil {
		return fmt.Errorf("decode FlowSetupList failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.FlowFailedList = new(QOSFlowFailedList)
		if err = s.FlowFailedList.Decode(r); err != nil {
			return fmt.Errorf("decode FlowFailedList failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<5) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible { /* TODO: Implement extension skipping for DRBSetupItemNGRAN */
	}
	return nil
}
