package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// CauseRadioNetwork From: 9_4_5_Information_Element_Definitions.txt:173
// ASN.1 Data Type: ENUMERATED
const (
	CauseRadioNetworkUnspecified                               aper.Enumerated = 0
	CauseRadioNetworkUnknownOrAlreadyAllocatedGNBCUCPUeE1APID  aper.Enumerated = 1
	CauseRadioNetworkUnknownOrAlreadyAllocatedGNBCUUPUeE1APID  aper.Enumerated = 2
	CauseRadioNetworkUnknownOrInconsistentPairOfUeE1APID       aper.Enumerated = 3
	CauseRadioNetworkInteractionWithOtherProcedure             aper.Enumerated = 4
	CauseRadioNetworkPPDCPCountWrapAround                      aper.Enumerated = 5
	CauseRadioNetworkNotSupportedQCIValue                      aper.Enumerated = 6
	CauseRadioNetworkNotSupported5QIValue                      aper.Enumerated = 7
	CauseRadioNetworkEncryptionAlgorithmsNotSupported          aper.Enumerated = 8
	CauseRadioNetworkIntegrityProtectionAlgorithmsNotSupported aper.Enumerated = 9
	CauseRadioNetworkUPIntegrityProtectionNotPossible          aper.Enumerated = 10
	CauseRadioNetworkUPConfidentialityProtectionNotPossible    aper.Enumerated = 11
	CauseRadioNetworkMultiplePDUSessionIDInstances             aper.Enumerated = 12
	CauseRadioNetworkUnknownPDUSessionID                       aper.Enumerated = 13
	CauseRadioNetworkMultipleQOSFlowIDInstances                aper.Enumerated = 14
	CauseRadioNetworkUnknownQOSFlowID                          aper.Enumerated = 15
	CauseRadioNetworkMultipleDRBIDInstances                    aper.Enumerated = 16
	CauseRadioNetworkUnknownDRBID                              aper.Enumerated = 17
	CauseRadioNetworkInvalidQOSCombination                     aper.Enumerated = 18
	CauseRadioNetworkProcedureCancelled                        aper.Enumerated = 19
	CauseRadioNetworkNormalRelease                             aper.Enumerated = 20
	CauseRadioNetworkNoRadioResourcesAvailable                 aper.Enumerated = 21
	CauseRadioNetworkActionEsirableForAdioEasons               aper.Enumerated = 22
	CauseRadioNetworkDRRResourcesNotAvailableForTheSlice       aper.Enumerated = 23
	CauseRadioNetworkPDCPConfigurationNotSupported             aper.Enumerated = 24
	CauseRadioNetworkUeDlMaxIPDataRateReason                   aper.Enumerated = 25
	CauseRadioNetworkUPIntegrityProtectionFailure              aper.Enumerated = 26
	CauseRadioNetworkReleaseDueToPreEmption                    aper.Enumerated = 27
	CauseRadioNetworkRsnNotAvailableForTheUP                   aper.Enumerated = 28
	CauseRadioNetworkNPNNotSupported                           aper.Enumerated = 29
)

type CauseRadioNetwork struct {
	Value aper.Enumerated
}
