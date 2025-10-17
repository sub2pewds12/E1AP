package e1ap_ies

import (
	"fmt"
	"io"
)

// IABUPTNLAddressUpdateFailure is a generated SEQUENCE type.
type IABUPTNLAddressUpdateFailure struct {
	TransactionID          TransactionID           `aper:"lb:0,ub:255,mandatory,reject,ext"`
	Cause                  Cause                   `aper:"mandatory,ignore,ext"`
	TimeToWait             *TimeToWait             `aper:"optional,ignore,ext"`
	CriticalityDiagnostics *CriticalityDiagnostics `aper:"optional,ignore,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *IABUPTNLAddressUpdateFailure) Encode(w io.Writer) error {
	_ = w // Placeholder to prevent unused import warning
	return fmt.Errorf("Encode not implemented for IABUPTNLAddressUpdateFailure")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *IABUPTNLAddressUpdateFailure) Decode(r io.Reader) error {
	_ = r // Placeholder to prevent unused import warning
	return fmt.Errorf("Decode not implemented for IABUPTNLAddressUpdateFailure")
}
