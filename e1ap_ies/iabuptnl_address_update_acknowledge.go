package e1ap_ies

import (
	"fmt"
	"io"
)

// IABUPTNLAddressUpdateAcknowledge is a generated SEQUENCE type.
type IABUPTNLAddressUpdateAcknowledge struct {
	TransactionID              TransactionID               `aper:"lb:0,ub:255,mandatory,reject,ext"`
	CriticalityDiagnostics     *CriticalityDiagnostics     `aper:"optional,ignore,ext"`
	ULUPTNLAddressToUpdateList *ULUPTNLAddressToUpdateList `aper:"optional,ignore,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *IABUPTNLAddressUpdateAcknowledge) Encode(w io.Writer) error {
	_ = w // Placeholder to prevent unused import warning
	return fmt.Errorf("Encode not implemented for IABUPTNLAddressUpdateAcknowledge")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *IABUPTNLAddressUpdateAcknowledge) Decode(r io.Reader) error {
	_ = r // Placeholder to prevent unused import warning
	return fmt.Errorf("Decode not implemented for IABUPTNLAddressUpdateAcknowledge")
}
