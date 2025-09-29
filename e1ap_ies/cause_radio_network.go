package e1ap_ies

// CauseRadioNetwork represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:173
type CauseRadioNetwork int32

const (
	CauseRadioNetwork_Unspecified                               CauseRadioNetwork = 0
	CauseRadioNetwork_UnknownOrAlreadyAllocatedGNBCUCPUeE1APID  CauseRadioNetwork = 1
	CauseRadioNetwork_UnknownOrAlreadyAllocatedGNBCUUPUeE1APID  CauseRadioNetwork = 2
	CauseRadioNetwork_UnknownOrInconsistentPairOfUeE1APID       CauseRadioNetwork = 3
	CauseRadioNetwork_InteractionWithOtherProcedure             CauseRadioNetwork = 4
	CauseRadioNetwork_PPDCPCountWrapAround                      CauseRadioNetwork = 5
	CauseRadioNetwork_NotSupportedQCIValue                      CauseRadioNetwork = 6
	CauseRadioNetwork_NotSupported5QIValue                      CauseRadioNetwork = 7
	CauseRadioNetwork_EncryptionAlgorithmsNotSupported          CauseRadioNetwork = 8
	CauseRadioNetwork_IntegrityProtectionAlgorithmsNotSupported CauseRadioNetwork = 9
	CauseRadioNetwork_UPIntegrityProtectionNotPossible          CauseRadioNetwork = 10
	CauseRadioNetwork_UPConfidentialityProtectionNotPossible    CauseRadioNetwork = 11
	CauseRadioNetwork_MultiplePDUSessionIDInstances             CauseRadioNetwork = 12
	CauseRadioNetwork_UnknownPDUSessionID                       CauseRadioNetwork = 13
	CauseRadioNetwork_MultipleQOSFlowIDInstances                CauseRadioNetwork = 14
	CauseRadioNetwork_UnknownQOSFlowID                          CauseRadioNetwork = 15
	CauseRadioNetwork_MultipleDRBIDInstances                    CauseRadioNetwork = 16
	CauseRadioNetwork_UnknownDRBID                              CauseRadioNetwork = 17
	CauseRadioNetwork_InvalidQOSCombination                     CauseRadioNetwork = 18
	CauseRadioNetwork_ProcedureCancelled                        CauseRadioNetwork = 19
	CauseRadioNetwork_NormalRelease                             CauseRadioNetwork = 20
	CauseRadioNetwork_NoRadioResourcesAvailable                 CauseRadioNetwork = 21
	CauseRadioNetwork_ActionEsirableForAdioEasons               CauseRadioNetwork = 22
	CauseRadioNetwork_DRRResourcesNotAvailableForTheSlice       CauseRadioNetwork = 23
	CauseRadioNetwork_PDCPConfigurationNotSupported             CauseRadioNetwork = 24
	CauseRadioNetwork_UeDlMaxIPDataRateReason                   CauseRadioNetwork = 25
	CauseRadioNetwork_UPIntegrityProtectionFailure              CauseRadioNetwork = 26
	CauseRadioNetwork_ReleaseDueToPreEmption                    CauseRadioNetwork = 27
	CauseRadioNetwork_RsnNotAvailableForTheUP                   CauseRadioNetwork = 28
	CauseRadioNetwork_NPNNotSupported                           CauseRadioNetwork = 29
)
