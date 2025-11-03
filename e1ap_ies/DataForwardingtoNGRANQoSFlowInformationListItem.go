package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// DataForwardingtoNGRANQoSFlowInformationListItem is a generated SEQUENCE type.
type DataForwardingtoNGRANQoSFlowInformationListItem struct {
	QOSFlowIdentifier QOSFlowIdentifier           `aper:"lb:0,ub:63,mandatory,ext"`
	IEExtensions      *ProtocolExtensionContainer `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *DataForwardingtoNGRANQoSFlowInformationListItem) Encode(w *aper.AperWriter) (err error) {
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
	if err = w.WriteInteger(int64(s.QOSFlowIdentifier.Value), &aper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return fmt.Errorf("Encode QOSFlowIdentifier failed: %w", err)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *DataForwardingtoNGRANQoSFlowInformationListItem) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}

	{
		var val int64
		if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
			return fmt.Errorf("Decode QOSFlowIdentifier failed: %w", err)
		}
		s.QOSFlowIdentifier.Value = aper.Integer(val)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for DataForwardingtoNGRANQoSFlowInformationListItem")
	}
	return nil
}
