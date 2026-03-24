package e1ap_ies

import (
	"github.com/lvdund/asn1go/per"
)

// CauseRadioNetwork is a generated ENUMERATED type.
type CauseRadioNetwork struct {
	Value int64
}

const (
	CauseRadioNetworkUnspecified                               int64 = 0
	CauseRadioNetworkUnknownOrAlreadyAllocatedGNBCUCPUeE1APID  int64 = 1
	CauseRadioNetworkUnknownOrAlreadyAllocatedGNBCUUPUeE1APID  int64 = 2
	CauseRadioNetworkUnknownOrInconsistentPairOfUeE1APID       int64 = 3
	CauseRadioNetworkInteractionWithOtherProcedure             int64 = 4
	CauseRadioNetworkPPDCPCountWrapAround                      int64 = 5
	CauseRadioNetworkNotSupportedQCIValue                      int64 = 6
	CauseRadioNetworkNotSupported5QIValue                      int64 = 7
	CauseRadioNetworkEncryptionAlgorithmsNotSupported          int64 = 8
	CauseRadioNetworkIntegrityProtectionAlgorithmsNotSupported int64 = 9
	CauseRadioNetworkUPIntegrityProtectionNotPossible          int64 = 10
	CauseRadioNetworkUPConfidentialityProtectionNotPossible    int64 = 11
	CauseRadioNetworkMultiplePDUSessionIDInstances             int64 = 12
	CauseRadioNetworkUnknownPDUSessionID                       int64 = 13
	CauseRadioNetworkMultipleQOSFlowIDInstances                int64 = 14
	CauseRadioNetworkUnknownQOSFlowID                          int64 = 15
	CauseRadioNetworkMultipleDRBIDInstances                    int64 = 16
	CauseRadioNetworkUnknownDRBID                              int64 = 17
	CauseRadioNetworkInvalidQOSCombination                     int64 = 18
	CauseRadioNetworkProcedureCancelled                        int64 = 19
	CauseRadioNetworkNormalRelease                             int64 = 20
	CauseRadioNetworkNoRadioResourcesAvailable                 int64 = 21
	CauseRadioNetworkActionDesirableForRadioReasons            int64 = 22
	CauseRadioNetworkResourcesNotAvailableForTheSlice          int64 = 23
	CauseRadioNetworkPDCPConfigurationNotSupported             int64 = 24
	CauseRadioNetworkUeDlMaxIPDataRateReason                   int64 = 25
	CauseRadioNetworkUPIntegrityProtectionFailure              int64 = 26
	CauseRadioNetworkReleaseDueToPreEmption                    int64 = 27
	CauseRadioNetworkRsnNotAvailableForTheUP                   int64 = 28
	CauseRadioNetworkNPNNotSupported                           int64 = 29
	CauseRadioNetworkReportCharacteristicEmpty                 int64 = 30
	CauseRadioNetworkExistingMeasurementID                     int64 = 31
	CauseRadioNetworkMeasurementTemporarilyNotAvailable        int64 = 32
	CauseRadioNetworkMeasurementNotSupportedForTheObject       int64 = 33
)

// Encode implements the MessageEncoder interface for CauseRadioNetwork.
func (e *CauseRadioNetwork) Encode(w *per.Encoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 34), ExtValues: nil}
	return w.EncodeEnumerated(int64(e.Value), c)
}

// Decode implements the MessageDecoder interface for CauseRadioNetwork.
func (e *CauseRadioNetwork) Decode(r *per.Decoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 34), ExtValues: nil}
	val, err := r.DecodeEnumerated(c)
	if err != nil {
		return err
	}
	e.Value = val
	return nil
}
