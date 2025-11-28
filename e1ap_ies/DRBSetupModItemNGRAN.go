package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// DRBSetupModItemNGRAN is a generated SEQUENCE type.
type DRBSetupModItemNGRAN struct {
	DRBID                                DRBID                       `aper:"lb:1,ub:32,mandatory,ext"`
	DRBDataForwardingInformationResponse *DataForwardingInformation  `aper:"optional,ext"`
	ULUPTransportParameters              UPParameters                `aper:"lb:1,ub:MaxnoofUPParameters,mandatory,ext"`
	FlowSetupList                        QOSFlowList                 `aper:"lb:1,ub:MaxnoofQoSFlows,mandatory,ext"`
	FlowFailedList                       *QOSFlowFailedList          `aper:"lb:1,ub:MaxnoofQoSFlows,optional,ext"`
	IEExtensions                         *ProtocolExtensionContainer `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *DRBSetupModItemNGRAN) Encode(w *aper.AperWriter) (err error) {
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
	if err = w.WriteInteger(int64(s.DRBID.Value), &aper.Constraint{Lb: 1, Ub: 32}, true); err != nil {
		return fmt.Errorf("Encode DRBID failed: %w", err)
	}
	if s.DRBDataForwardingInformationResponse != nil {
		if err = s.DRBDataForwardingInformationResponse.Encode(w); err != nil {
			return fmt.Errorf("Encode DRBDataForwardingInformationResponse failed: %w", err)
		}
	}

	{
		itemPointers := make([]aper.AperMarshaller, len(s.ULUPTransportParameters.Value))
		for i := 0; i < len(s.ULUPTransportParameters.Value); i++ {
			itemPointers[i] = &(s.ULUPTransportParameters.Value[i])
		}
		if err = aper.WriteSequenceOf(itemPointers, w, &aper.Constraint{Lb: 1, Ub: MaxnoofUPParameters}, false); err != nil {
			return fmt.Errorf("Encode ULUPTransportParameters failed: %w", err)
		}
	}

	{
		itemPointers := make([]aper.AperMarshaller, len(s.FlowSetupList.Value))
		for i := 0; i < len(s.FlowSetupList.Value); i++ {
			itemPointers[i] = &(s.FlowSetupList.Value[i])
		}
		if err = aper.WriteSequenceOf(itemPointers, w, &aper.Constraint{Lb: 1, Ub: MaxnoofQoSFlows}, false); err != nil {
			return fmt.Errorf("Encode FlowSetupList failed: %w", err)
		}
	}
	if s.FlowFailedList != nil {
		{
			itemPointers := make([]aper.AperMarshaller, len(s.FlowFailedList.Value))
			for i := 0; i < len(s.FlowFailedList.Value); i++ {
				itemPointers[i] = &(s.FlowFailedList.Value[i])
			}
			if err = aper.WriteSequenceOf(itemPointers, w, &aper.Constraint{Lb: 1, Ub: MaxnoofQoSFlows}, false); err != nil {
				return fmt.Errorf("Encode FlowFailedList failed: %w", err)
			}
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *DRBSetupModItemNGRAN) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 3, Ub: 3}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}
	if err = s.DRBID.Decode(r); err != nil {
		return fmt.Errorf("Decode DRBID failed: %w", err)
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
	if isExtensible { /* TODO: Implement extension skipping for DRBSetupModItemNGRAN */
	}
	return nil
}
