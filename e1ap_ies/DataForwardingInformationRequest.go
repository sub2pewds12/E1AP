package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// DataForwardingInformationRequest is a generated SEQUENCE type.
type DataForwardingInformationRequest struct {
	DataForwardingRequest         DataForwardingRequest       `aper:"mandatory,ext"`
	QOSFlowsForwardedOnFwdTunnels *QOSFlowMappingList         `aper:"lb:1,ub:MaxnoofQoSFlows,optional,ext"`
	IEExtensions                  *ProtocolExtensionContainer `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *DataForwardingInformationRequest) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.QOSFlowsForwardedOnFwdTunnels != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(2), &aper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return fmt.Errorf("Encode optionality bitmap failed: %w", err)
	}
	if err = w.WriteEnumerate(uint64(s.DataForwardingRequest.Value), aper.Constraint{Lb: 0, Ub: 2}, true); err != nil {
		return fmt.Errorf("Encode DataForwardingRequest failed: %w", err)
	}
	if s.QOSFlowsForwardedOnFwdTunnels != nil {
		if err = s.QOSFlowsForwardedOnFwdTunnels.Encode(w); err != nil {
			return fmt.Errorf("Encode QOSFlowsForwardedOnFwdTunnels failed: %w", err)
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
func (s *DataForwardingInformationRequest) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}
	if err = s.DataForwardingRequest.Decode(r); err != nil {
		return fmt.Errorf("Decode DataForwardingRequest failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.QOSFlowsForwardedOnFwdTunnels = new(QOSFlowMappingList)
		if err = s.QOSFlowsForwardedOnFwdTunnels.Decode(r); err != nil {
			return fmt.Errorf("Decode QOSFlowsForwardedOnFwdTunnels failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible { /* TODO: Implement extension skipping for DataForwardingInformationRequest */
	}
	return nil
}
