package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// CriticalityDiagnostics is a generated SEQUENCE type.
type CriticalityDiagnostics struct {
	ProcedureCode             *ProcedureCode                 `aper:"lb:0,ub:255,optional,ext"`
	TriggeringMessage         *TriggeringMessage             `aper:"optional,ext"`
	ProcedureCriticality      *Criticality                   `aper:"optional,ext"`
	TransactionID             TransactionID                  `aper:"lb:0,ub:255,mandatory,ext"`
	IEsCriticalityDiagnostics []CriticalityDiagnosticsIEItem `aper:"optional,ext"`
	IEExtensions              *ProtocolExtensionContainer    `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *CriticalityDiagnostics) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.ProcedureCode != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.TriggeringMessage != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if s.ProcedureCriticality != nil {
		optionalityBitmap[0] |= 1 << 5
	}
	if s.IEsCriticalityDiagnostics != nil {
		optionalityBitmap[0] |= 1 << 4
	}
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 3
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(5), &aper.Constraint{Lb: 5, Ub: 5}, false); err != nil {
		return fmt.Errorf("Encode optionality bitmap failed: %w", err)
	}
	if s.ProcedureCode != nil {
		if err = w.WriteInteger(int64(s.ProcedureCode.Value), &aper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
			return fmt.Errorf("Encode ProcedureCode failed: %w", err)
		}
	}
	if s.TriggeringMessage != nil {
		if err = w.WriteEnumerate(uint64(s.TriggeringMessage.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
			return fmt.Errorf("Encode TriggeringMessage failed: %w", err)
		}
	}
	if s.ProcedureCriticality != nil {
		if err = w.WriteEnumerate(uint64(s.ProcedureCriticality.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
			return fmt.Errorf("Encode ProcedureCriticality failed: %w", err)
		}
	}
	if err = w.WriteInteger(int64(s.TransactionID.Value), &aper.Constraint{Lb: 0, Ub: 255}, true); err != nil {
		return fmt.Errorf("Encode TransactionID failed: %w", err)
	}
	if s.IEsCriticalityDiagnostics != nil {
		if err = s.IEsCriticalityDiagnostics.Encode(w); err != nil {
			return fmt.Errorf("Encode IEsCriticalityDiagnostics failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *CriticalityDiagnostics) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 5, Ub: 5}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {

		{
			var val int64
			if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
				return fmt.Errorf("Decode ProcedureCode failed: %w", err)
			}
			s.ProcedureCode = new(ProcedureCode)
			s.ProcedureCode.Value = aper.Integer(val)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.TriggeringMessage = new(TriggeringMessage)

		{
			var val uint64
			if val, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
				return fmt.Errorf("Decode TriggeringMessage failed: %w", err)
			}
			s.TriggeringMessage.Value = aper.Enumerated(val)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<5) > 0 {
		s.ProcedureCriticality = new(Criticality)

		{
			var val uint64
			if val, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
				return fmt.Errorf("Decode ProcedureCriticality failed: %w", err)
			}
			s.ProcedureCriticality.Value = aper.Enumerated(val)
		}
	}

	{
		var val int64
		if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 255}, true); err != nil {
			return fmt.Errorf("Decode TransactionID failed: %w", err)
		}
		s.TransactionID.Value = aper.Integer(val)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<4) > 0 {
		s.IEsCriticalityDiagnostics = new(CriticalityDiagnosticsIEList)
		if err = s.IEsCriticalityDiagnostics.Decode(r); err != nil {
			return fmt.Errorf("Decode IEsCriticalityDiagnostics failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<3) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for CriticalityDiagnostics")
	}
	return nil
}
