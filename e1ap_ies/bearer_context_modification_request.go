package e1ap_ies

import (
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

// BearerContextModificationRequest is a generated SEQUENCE type.
type BearerContextModificationRequest struct {
	GNBCUCPUEE1APID                        aper.Integer                            `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	GNBCUUPUEE1APID                        aper.Integer                            `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	SecurityInformation                    *SecurityInformation                    `aper:"optional,reject,ext"`
	UEDLAggregateMaximumBitRate            *aper.Integer                           `aper:"lb:0,ub:4000000000000,optional,reject,ext"`
	UEDLMaximumIntegrityProtectedDataRate  *aper.Integer                           `aper:"lb:0,ub:4000000000000,optional,reject,ext"`
	BearerContextStatusChange              *BearerContextStatusChange              `aper:"optional,reject,ext"`
	NewULTNLInformationRequired            *NewULTNLInformationRequired            `aper:"optional,reject,ext"`
	UEInactivityTimer                      *aper.Integer                           `aper:"lb:1,ub:7200,optional,reject,ext"`
	DataDiscardRequired                    *DataDiscardRequired                    `aper:"optional,ignore,ext"`
	SystemBearerContextModificationRequest *SystemBearerContextModificationRequest `aper:"optional,reject,ext"`
	RANUEID                                *aper.OctetString                       `aper:"lb:8,ub:8,optional,ignore,ext"`
	GNBDUID                                *aper.Integer                           `aper:"lb:0,ub:68719476735,optional,ignore,ext"`
	ActivityNotificationLevel              *ActivityNotificationLevel              `aper:"optional,ignore,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *BearerContextModificationRequest) Encode(w io.Writer) error {
	_ = w // Placeholder to prevent unused import warning
	return fmt.Errorf("Encode not implemented for BearerContextModificationRequest")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *BearerContextModificationRequest) Decode(r io.Reader) error {
	_ = r // Placeholder to prevent unused import warning
	return fmt.Errorf("Decode not implemented for BearerContextModificationRequest")
}
