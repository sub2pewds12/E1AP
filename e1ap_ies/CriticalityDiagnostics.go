package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// CriticalityDiagnostics is a generated SEQUENCE type.
type CriticalityDiagnostics struct {
	ProcedureCode             *ProcedureCode                `aper:"lb:0,ub:255,optional,ext"`
	TriggeringMessage         *TriggeringMessage            `aper:"optional,ext"`
	ProcedureCriticality      *Criticality                  `aper:"optional,ext"`
	TransactionID             TransactionID                 `aper:"lb:0,ub:255,mandatory,ext"`
	IEsCriticalityDiagnostics *CriticalityDiagnosticsIEList `aper:"optional,ext"`
	IEExtensions              *ProtocolExtensionContainer   `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *CriticalityDiagnostics) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("encode extensibility bool failed: %w", err)
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
		return fmt.Errorf("encode optionality bitmap failed: %w", err)
	}
	if s.ProcedureCode != nil {
		if err = w.WriteInteger(int64((*s.ProcedureCode).Value), &aper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
			return fmt.Errorf("encode ProcedureCode failed: %w", err)
		}
	}
	if s.TriggeringMessage != nil {
		if err = w.WriteEnumerate(uint64((*s.TriggeringMessage).Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
			return fmt.Errorf("encode TriggeringMessage failed: %w", err)
		}
	}
	if s.ProcedureCriticality != nil {
		if err = w.WriteEnumerate(uint64((*s.ProcedureCriticality).Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
			return fmt.Errorf("encode ProcedureCriticality failed: %w", err)
		}
	}
	if err = w.WriteInteger(int64(s.TransactionID.Value), &aper.Constraint{Lb: 0, Ub: 255}, true); err != nil {
		return fmt.Errorf("encode TransactionID failed: %w", err)
	}
	if s.IEsCriticalityDiagnostics != nil {
		if err = s.IEsCriticalityDiagnostics.Encode(w); err != nil {
			return fmt.Errorf("encode IEsCriticalityDiagnostics failed: %w", err)
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
func (s *CriticalityDiagnostics) Decode(r *aper.AperReader) (err error) {
	isExtensible, err := r.ReadBool()
	if err != nil {
		return fmt.Errorf("read extensibility bool failed: %w", err)
	}
	_ = isExtensible
	optionalityBitmap, _, err := r.ReadBitString(&aper.Constraint{Lb: 5, Ub: 5}, false)
	if err != nil {
		return fmt.Errorf("read optionality bitmap failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.ProcedureCode = new(ProcedureCode)
		if err = s.ProcedureCode.Decode(r); err != nil {
			return fmt.Errorf("decode ProcedureCode failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.TriggeringMessage = new(TriggeringMessage)
		if err = s.TriggeringMessage.Decode(r); err != nil {
			return fmt.Errorf("decode TriggeringMessage failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<5) > 0 {
		s.ProcedureCriticality = new(Criticality)
		if err = s.ProcedureCriticality.Decode(r); err != nil {
			return fmt.Errorf("decode ProcedureCriticality failed: %w", err)
		}
	}
	if err = s.TransactionID.Decode(r); err != nil {
		return fmt.Errorf("decode TransactionID failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<4) > 0 {
		s.IEsCriticalityDiagnostics = new(CriticalityDiagnosticsIEList)
		if err = s.IEsCriticalityDiagnostics.Decode(r); err != nil {
			return fmt.Errorf("decode IEsCriticalityDiagnostics failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<3) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible { /* TODO: Implement extension skipping for CriticalityDiagnostics */
	}
	return nil
}
