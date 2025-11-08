package e1ap

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	_ "github.com/sub2pewds12/E1AP/e1ap_ies"
)

func (im *InitiatingMessage) Encode(w io.Writer) error {
	waper := aper.NewWriter(w)
	defer waper.Close()

	if err := waper.WriteChoice(uint64(PduChoiceInitiatingMessage), 2, false); err != nil {
		return fmt.Errorf("Encode PDU Choice failed: %w", err)
	}

	if err := im.ProcedureCode.Encode(waper); err != nil {
		return fmt.Errorf("Encode ProcedureCode failed: %w", err)
	}

	if err := im.Criticality.Encode(waper); err != nil {
		return fmt.Errorf("Encode Criticality failed: %w", err)
	}

	var ieBuf bytes.Buffer

	innerPDU, ok := im.Value.(MessageEncoder)
	if !ok {
		return fmt.Errorf("PDU value does not implement MessageEncoder interface")
	}
	if err := innerPDU.Encode(&ieBuf); err != nil {
		return fmt.Errorf("Encode inner PDU failed: %w", err)
	}

	if err := waper.WriteOpenType(ieBuf.Bytes()); err != nil {
		return fmt.Errorf("Encode OpenType (IEs) failed: %w", err)
	}

	return nil
}

func (so *SuccessfulOutcome) Encode(w io.Writer) error {
	waper := aper.NewWriter(w)
	defer waper.Close()

	if err := waper.WriteChoice(uint64(PduChoiceSuccessfulOutcome), 2, false); err != nil {
		return fmt.Errorf("Encode PDU Choice failed: %w", err)
	}

	if err := so.ProcedureCode.Encode(waper); err != nil {
		return err
	}
	if err := so.Criticality.Encode(waper); err != nil {
		return err
	}

	var ieBuf bytes.Buffer
	innerPDU, ok := so.Value.(MessageEncoder)
	if !ok {
		return fmt.Errorf("PDU value does not implement MessageEncoder")
	}
	if err := innerPDU.Encode(&ieBuf); err != nil {
		return err
	}

	if err := waper.WriteOpenType(ieBuf.Bytes()); err != nil {
		return err
	}

	return nil
}

func (uo *UnsuccessfulOutcome) Encode(w io.Writer) error {
	waper := aper.NewWriter(w)
	defer waper.Close()

	if err := waper.WriteChoice(uint64(PduChoiceUnsuccessfulOutcome), 2, false); err != nil {
		return fmt.Errorf("Encode PDU Choice failed: %w", err)
	}

	if err := uo.ProcedureCode.Encode(waper); err != nil {
		return err
	}
	if err := uo.Criticality.Encode(waper); err != nil {
		return err
	}

	var ieBuf bytes.Buffer
	innerPDU, ok := uo.Value.(MessageEncoder)
	if !ok {
		return fmt.Errorf("PDU value does not implement MessageEncoder")
	}
	if err := innerPDU.Encode(&ieBuf); err != nil {
		return err
	}

	if err := waper.WriteOpenType(ieBuf.Bytes()); err != nil {
		return err
	}

	return nil
}

func E1apEncode(msg MessageEncoder) (wire []byte, err error) {
	var buf bytes.Buffer

	if err = msg.Encode(&buf); err == nil {
		wire = buf.Bytes()
	}
	return
}
