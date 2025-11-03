package e1ap

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	_ "github.com/sub2pewds12/E1AP/e1ap_ies"
)

// Encode is the master encoder for all Initiating Messages.
func (im *InitiatingMessage) Encode(w io.Writer) error {
	waper := aper.NewWriter(w)
	defer waper.Close() // Ensure all bits are flushed at the end

	// 1. Encode the PDU CHOICE index (0 for InitiatingMessage)
	if err := waper.WriteChoice(uint64(PduChoiceInitiatingMessage), 2, false); err != nil { // E1AP PDU is not extensible
		return fmt.Errorf("Encode PDU Choice failed: %w", err)
	}

	// 2. Encode the Procedure Code
	if err := im.ProcedureCode.Encode(waper); err != nil {
		return fmt.Errorf("Encode ProcedureCode failed: %w", err)
	}

	// 3. Encode the Criticality
	if err := im.Criticality.Encode(waper); err != nil {
		return fmt.Errorf("Encode Criticality failed: %w", err)
	}

	// 4. Encode the inner PDU's IEs (the Value) into a temporary buffer.
	var ieBuf bytes.Buffer
	// The Value is an interface, we need to type-assert it to the MessageEncoder interface.
	innerPDU, ok := im.Value.(MessageEncoder)
	if !ok {
		return fmt.Errorf("PDU value does not implement MessageEncoder interface")
	}
	if err := innerPDU.Encode(&ieBuf); err != nil { // This calls the PDU's generated Encode method
		return fmt.Errorf("Encode inner PDU failed: %w", err)
	}

	// 5. Write the encoded IEs as an Open Type.
	if err := waper.WriteOpenType(ieBuf.Bytes()); err != nil {
		return fmt.Errorf("Encode OpenType (IEs) failed: %w", err)
	}

	return nil
}

// Encode is the master encoder for all Successful Outcome messages.
func (so *SuccessfulOutcome) Encode(w io.Writer) error {
	waper := aper.NewWriter(w)
	defer waper.Close()

	// 1. Encode the PDU CHOICE index (1 for SuccessfulOutcome)
	if err := waper.WriteChoice(uint64(PduChoiceSuccessfulOutcome), 2, false); err != nil {
		return fmt.Errorf("Encode PDU Choice failed: %w", err)
	}

	// ... The rest of the logic is identical to InitiatingMessage ...
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

// Encode is the master encoder for all Unsuccessful Outcome messages.
func (uo *UnsuccessfulOutcome) Encode(w io.Writer) error {
	waper := aper.NewWriter(w)
	defer waper.Close()

	// 1. Encode the PDU CHOICE index (2 for UnsuccessfulOutcome)
	if err := waper.WriteChoice(uint64(PduChoiceUnsuccessfulOutcome), 2, false); err != nil {
		return fmt.Errorf("Encode PDU Choice failed: %w", err)
	}

	// ... The rest of the logic is identical ...
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
