package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// DataForwardingtoEUTRANInformationListItem is a generated SEQUENCE type.
type DataForwardingtoEUTRANInformationListItem struct {
	DataForwardingTunnelInformation UPTNLInformation            `aper:"mandatory,ext"`
	QOSFlowsToBeForwardedList       QOSFlowsToBeForwardedList   `aper:"lb:1,ub:MaxnoofQoSFlows,mandatory,ext"`
	IEExtensions                    *ProtocolExtensionContainer `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *DataForwardingtoEUTRANInformationListItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(1), &aper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
		return fmt.Errorf("encode optionality bitmap failed: %w", err)
	}
	if err = s.DataForwardingTunnelInformation.Encode(w); err != nil {
		return fmt.Errorf("encode DataForwardingTunnelInformation failed: %w", err)
	}
	if err = s.QOSFlowsToBeForwardedList.Encode(w); err != nil {
		return fmt.Errorf("encode QOSFlowsToBeForwardedList failed: %w", err)
	}
	if s.IEExtensions != nil {
		if err = s.IEExtensions.Encode(w); err != nil {
			return fmt.Errorf("encode IEExtensions failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *DataForwardingtoEUTRANInformationListItem) Decode(r *aper.AperReader) (err error) {
	isExtensible, err := r.ReadBool()
	if err != nil {
		return fmt.Errorf("read extensibility bool failed: %w", err)
	}
	_ = isExtensible
	optionalityBitmap, _, err := r.ReadBitString(&aper.Constraint{Lb: 1, Ub: 1}, false)
	if err != nil {
		return fmt.Errorf("read optionality bitmap failed: %w", err)
	}
	if err = s.DataForwardingTunnelInformation.Decode(r); err != nil {
		return fmt.Errorf("decode DataForwardingTunnelInformation failed: %w", err)
	}
	if err = s.QOSFlowsToBeForwardedList.Decode(r); err != nil {
		return fmt.Errorf("decode QOSFlowsToBeForwardedList failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible { /* TODO: Implement extension skipping for DataForwardingtoEUTRANInformationListItem */
	}
	return nil
}
