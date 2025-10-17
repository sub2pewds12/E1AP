package e1ap_ies

import (
	"fmt"
	"io"
)

// BearerContextSetupRequest is a generated SEQUENCE type.
type BearerContextSetupRequest struct {
	GNBCUCPUEE1APID                       GNBCUCPUEE1APID                   `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	SecurityInformation                   SecurityInformation               `aper:"mandatory,reject,ext"`
	UEDLAggregateMaximumBitRate           BitRate                           `aper:"lb:0,ub:4000000000000,mandatory,reject,ext"`
	UEDLMaximumIntegrityProtectedDataRate *BitRate                          `aper:"lb:0,ub:4000000000000,optional,reject,ext"`
	ServingPLMN                           PLMNIdentity                      `aper:"lb:3,ub:3,mandatory,ignore,ext"`
	ActivityNotificationLevel             ActivityNotificationLevel         `aper:"mandatory,reject,ext"`
	UEInactivityTimer                     *InactivityTimer                  `aper:"lb:1,ub:7200,optional,reject,ext"`
	BearerContextStatusChange             *BearerContextStatusChange        `aper:"optional,reject,ext"`
	SystemBearerContextSetupRequest       SystemBearerContextSetupRequest   `aper:"mandatory,reject,ext"`
	RANUEID                               *RANUEID                          `aper:"lb:8,ub:8,optional,ignore,ext"`
	GNBDUID                               *GNBDUID                          `aper:"lb:0,ub:68719476735,optional,ignore,ext"`
	TraceActivation                       *TraceActivation                  `aper:"optional,ignore,ext"`
	NPNContextInfo                        *NPNContextInfo                   `aper:"optional,reject,ext"`
	ManagementBasedMDTPLMNList            []PLMNIdentity                    `aper:"lb:1,ub:MaxnoofMDTPLMNs,optional,ignore,ext"`
	CHOInitiation                         *CHOInitiation                    `aper:"optional,reject,ext"`
	AdditionalHandoverInfo                *AdditionalHandoverInfo           `aper:"optional,ignore,ext"`
	DirectForwardingPathAvailability      *DirectForwardingPathAvailability `aper:"optional,ignore,ext"`
	GNBCUUPUEE1APID                       *GNBCUUPUEE1APID                  `aper:"lb:0,ub:4294967295,optional,ignore,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *BearerContextSetupRequest) Encode(w io.Writer) error {
	_ = w // Placeholder to prevent unused import warning
	return fmt.Errorf("Encode not implemented for BearerContextSetupRequest")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *BearerContextSetupRequest) Decode(r io.Reader) error {
	_ = r // Placeholder to prevent unused import warning
	return fmt.Errorf("Decode not implemented for BearerContextSetupRequest")
}
