package e1ap_ies

import (
	"fmt"

	"asn1go/per"
)

// CriticalityDiagnostics is a generated SEQUENCE type.
type CriticalityDiagnostics struct {
	ProcedureCode             *ProcedureCode
	TriggeringMessage         *TriggeringMessage
	ProcedureCriticality      *Criticality
	TransactionID             TransactionID
	IEsCriticalityDiagnostics *CriticalityDiagnosticsIEList
	IEExtensions              *CriticalityDiagnosticsExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *CriticalityDiagnostics) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "procedureCode", Optional: true},
			per.ComponentInfo{Name: "triggeringMessage", Optional: true},
			per.ComponentInfo{Name: "procedureCriticality", Optional: true},
			per.ComponentInfo{Name: "transactionID", Optional: false},
			per.ComponentInfo{Name: "iEsCriticalityDiagnostics", Optional: true},
			per.ComponentInfo{Name: "iE-Extensions", Optional: true},
		},
	}
	seqEncoder := w.NewSequenceEncoder(c)
	if err := seqEncoder.EncodeExtensionBit(false); err != nil {
		return err
	}

	optionalBitmap := make([]bool, 0)

	if s.ProcedureCode != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.TriggeringMessage != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.ProcedureCriticality != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.IEsCriticalityDiagnostics != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.IEExtensions != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if err := seqEncoder.EncodePreamble(optionalBitmap); err != nil {
		return err
	}

	if s.ProcedureCode != nil {
		if err = w.EncodeInteger(int64((*s.ProcedureCode).Value), per.Constrained(0, 255)); err != nil {
			return fmt.Errorf("encode ProcedureCode failed: %w", err)
		}
	}

	if s.TriggeringMessage != nil {

		{
			enumC := per.EnumeratedConstraints{Extensible: false, RootValues: make([]int64, 3), ExtValues: nil}
			if err = w.EncodeEnumerated(int64((*s.TriggeringMessage).Value), enumC); err != nil {
				return fmt.Errorf("encode TriggeringMessage failed: %w", err)
			}
		}
	}

	if s.ProcedureCriticality != nil {

		{
			enumC := per.EnumeratedConstraints{Extensible: false, RootValues: make([]int64, 3), ExtValues: nil}
			if err = w.EncodeEnumerated(int64((*s.ProcedureCriticality).Value), enumC); err != nil {
				return fmt.Errorf("encode ProcedureCriticality failed: %w", err)
			}
		}
	}
	if err = w.EncodeInteger(int64(s.TransactionID.Value), per.ConstrainedExtensible(0, 255)); err != nil {
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

	if err := seqEncoder.EncodeExtensionAdditions([]bool{}, [][]byte{}); err != nil {
		return err
	}

	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *CriticalityDiagnostics) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "procedureCode", Optional: true},
			per.ComponentInfo{Name: "triggeringMessage", Optional: true},
			per.ComponentInfo{Name: "procedureCriticality", Optional: true},
			per.ComponentInfo{Name: "transactionID", Optional: false},
			per.ComponentInfo{Name: "iEsCriticalityDiagnostics", Optional: true},
			per.ComponentInfo{Name: "iE-Extensions", Optional: true},
		},
	}
	seqDecoder := r.NewSequenceDecoder(c)
	if err := seqDecoder.DecodeExtensionBit(); err != nil {
		return err
	}

	if err := seqDecoder.DecodePreamble(); err != nil {
		return err
	}

	if seqDecoder.IsComponentPresent(0) {
		s.ProcedureCode = new(ProcedureCode)

		{
			val, err := r.DecodeInteger(per.Constrained(0, 255))
			if err != nil {
				return fmt.Errorf("decode ProcedureCode failed: %w", err)
			}
			s.ProcedureCode.Value = val
		}
	}

	if seqDecoder.IsComponentPresent(1) {
		s.TriggeringMessage = new(TriggeringMessage)

		{
			enumC := per.EnumeratedConstraints{Extensible: false, RootValues: make([]int64, 3), ExtValues: nil}
			val, err := r.DecodeEnumerated(enumC)
			if err != nil {
				return fmt.Errorf("decode TriggeringMessage failed: %w", err)
			}
			s.TriggeringMessage.Value = val
		}
	}

	if seqDecoder.IsComponentPresent(2) {
		s.ProcedureCriticality = new(Criticality)

		{
			enumC := per.EnumeratedConstraints{Extensible: false, RootValues: make([]int64, 3), ExtValues: nil}
			val, err := r.DecodeEnumerated(enumC)
			if err != nil {
				return fmt.Errorf("decode ProcedureCriticality failed: %w", err)
			}
			s.ProcedureCriticality.Value = val
		}
	}

	{
		val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 255))
		if err != nil {
			return fmt.Errorf("decode TransactionID failed: %w", err)
		}
		s.TransactionID.Value = val
	}

	if seqDecoder.IsComponentPresent(4) {
		s.IEsCriticalityDiagnostics = new(CriticalityDiagnosticsIEList)
		if err = s.IEsCriticalityDiagnostics.Decode(r); err != nil {
			return fmt.Errorf("Decode IEsCriticalityDiagnostics failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(5) {
		s.IEExtensions = new(CriticalityDiagnosticsExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
