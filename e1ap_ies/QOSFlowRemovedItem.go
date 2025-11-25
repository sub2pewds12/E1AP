package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// QOSFlowRemovedItem is a generated SEQUENCE type.
type QOSFlowRemovedItem struct {
	QOSFlowIdentifier             QOSFlowIdentifier                                `aper:"lb:0,ub:63,mandatory,ext"`
	QOSFlowReleasedInSession      *QOSFlowRemovedItemQOSFlowReleasedInSession      `aper:"optional,ext"`
	QOSFlowAccumulatedSessionTime *QOSFlowRemovedItemQOSFlowAccumulatedSessionTime `aper:"lb:5,ub:5,optional,ext"`
	IEExtensions                  *ProtocolExtensionContainer                      `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *QOSFlowRemovedItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.QOSFlowReleasedInSession != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.QOSFlowAccumulatedSessionTime != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 5
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(3), &aper.Constraint{Lb: 3, Ub: 3}, false); err != nil {
		return fmt.Errorf("Encode optionality bitmap failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.QOSFlowIdentifier.Value), &aper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return fmt.Errorf("Encode QOSFlowIdentifier failed: %w", err)
	}
	if s.QOSFlowReleasedInSession != nil {
		if err = s.QOSFlowReleasedInSession.Encode(w); err != nil {
			return fmt.Errorf("Encode QOSFlowReleasedInSession failed: %w", err)
		}
	}
	if s.QOSFlowAccumulatedSessionTime != nil {
		if err = w.WriteOctetString([]byte(s.QOSFlowAccumulatedSessionTime.Value), &aper.Constraint{Lb: 5, Ub: 5}, false); err != nil {
			return fmt.Errorf("Encode QOSFlowAccumulatedSessionTime failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *QOSFlowRemovedItem) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 3, Ub: 3}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}
	if err = s.QOSFlowIdentifier.Decode(r); err != nil {
		return fmt.Errorf("Decode QOSFlowIdentifier failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.QOSFlowReleasedInSession = new(QOSFlowRemovedItemQOSFlowReleasedInSession)
		if err = s.QOSFlowReleasedInSession.Decode(r); err != nil {
			return fmt.Errorf("Decode QOSFlowReleasedInSession failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.QOSFlowAccumulatedSessionTime = new(QOSFlowRemovedItemQOSFlowAccumulatedSessionTime)
		if err = s.QOSFlowAccumulatedSessionTime.Decode(r); err != nil {
			return fmt.Errorf("Decode QOSFlowAccumulatedSessionTime failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<5) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible { /* TODO: Implement extension skipping for QOSFlowRemovedItem */
	}
	return nil
}
