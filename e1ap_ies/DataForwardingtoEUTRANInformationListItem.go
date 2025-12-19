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
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(1), &aper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
		return fmt.Errorf("Encode optionality bitmap failed: %w", err)
	}
	if err = s.DataForwardingTunnelInformation.Encode(w); err != nil {
		return fmt.Errorf("Encode DataForwardingTunnelInformation failed: %w", err)
	}
	if err = s.QOSFlowsToBeForwardedList.Encode(w); err != nil {
		return fmt.Errorf("Encode QOSFlowsToBeForwardedList failed: %w", err)
	}
	if s.IEExtensions != nil {
		if err = s.IEExtensions.Encode(w); err != nil {
			return fmt.Errorf("Encode IEExtensions failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *DataForwardingtoEUTRANInformationListItem) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}
	if err = s.DataForwardingTunnelInformation.Decode(r); err != nil {
		return fmt.Errorf("Decode DataForwardingTunnelInformation failed: %w", err)
	}
	if err = s.QOSFlowsToBeForwardedList.Decode(r); err != nil {
		return fmt.Errorf("Decode QOSFlowsToBeForwardedList failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible { /* TODO: Implement extension skipping for DataForwardingtoEUTRANInformationListItem */
	}
	return nil
}
