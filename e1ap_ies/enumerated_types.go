package e1ap_ies

type ActivityNotificationLevel int32
const (
	ActivityNotificationLevel_Drb ActivityNotificationLevel = 0
	ActivityNotificationLevel_PduSession ActivityNotificationLevel = 1
	ActivityNotificationLevel_Ue ActivityNotificationLevel = 2
)

type AdditionalPDCPduplicationInformation int32
const (
	AdditionalPDCPduplicationInformation_Three AdditionalPDCPduplicationInformation = 0
	AdditionalPDCPduplicationInformation_Four AdditionalPDCPduplicationInformation = 1
)

type BearerContextStatusChange int32
const (
	BearerContextStatusChange_Suspend BearerContextStatusChange = 0
	BearerContextStatusChange_Resume BearerContextStatusChange = 1
)

type CHOInitiation int32
const (
	CHOInitiation_True CHOInitiation = 0
)

type CNSupport int32
const (
	CNSupport_CEpc CNSupport = 0
	CNSupport_C5gc CNSupport = 1
	CNSupport_Both CNSupport = 2
)

type CauseMisc int32
const (
	CauseMisc_ControlProcessingOverload CauseMisc = 0
	CauseMisc_NotEnoughUserPlaneProcessingResources CauseMisc = 1
	CauseMisc_HardwareFailure CauseMisc = 2
	CauseMisc_OmIntervention CauseMisc = 3
	CauseMisc_Unspecified CauseMisc = 4
)

type CauseProtocol int32
const (
	CauseProtocol_TransferSyntaxError CauseProtocol = 0
	CauseProtocol_AbstractSyntaxErrorReject CauseProtocol = 1
	CauseProtocol_AbstractSyntaxErrorIgnoreAndNotify CauseProtocol = 2
	CauseProtocol_MessageNotCompatibleWithReceiverState CauseProtocol = 3
	CauseProtocol_SemanticError CauseProtocol = 4
	CauseProtocol_AbstractSyntaxErrorFalselyConstructedMessage CauseProtocol = 5
	CauseProtocol_Unspecified CauseProtocol = 6
)

type CauseRadioNetwork int32
const (
	CauseRadioNetwork_Unspecified CauseRadioNetwork = 0
	CauseRadioNetwork_UnknownOrAlreadyAllocatedGnbCuCpUeE1apId CauseRadioNetwork = 1
	CauseRadioNetwork_UnknownOrAlreadyAllocatedGnbCuUpUeE1apId CauseRadioNetwork = 2
	CauseRadioNetwork_UnknownOrInconsistentPairOfUeE1apId CauseRadioNetwork = 3
	CauseRadioNetwork_InteractionWithOtherProcedure CauseRadioNetwork = 4
	CauseRadioNetwork_PPDCPCountWrapAround CauseRadioNetwork = 5
	CauseRadioNetwork_NotSupportedQciValue CauseRadioNetwork = 6
	CauseRadioNetwork_NotSupported5qiValue CauseRadioNetwork = 7
	CauseRadioNetwork_EncryptionAlgorithmsNotSupported CauseRadioNetwork = 8
	CauseRadioNetwork_IntegrityProtectionAlgorithmsNotSupported CauseRadioNetwork = 9
	CauseRadioNetwork_UPIntegrityProtectionNotPossible CauseRadioNetwork = 10
	CauseRadioNetwork_UPConfidentialityProtectionNotPossible CauseRadioNetwork = 11
	CauseRadioNetwork_MultiplePduSessionIdInstances CauseRadioNetwork = 12
	CauseRadioNetwork_UnknownPduSessionId CauseRadioNetwork = 13
	CauseRadioNetwork_MultipleQoSFlowIdInstances CauseRadioNetwork = 14
	CauseRadioNetwork_UnknownQoSFlowId CauseRadioNetwork = 15
	CauseRadioNetwork_MultipleDrbIdInstances CauseRadioNetwork = 16
	CauseRadioNetwork_UnknownDrbId CauseRadioNetwork = 17
	CauseRadioNetwork_InvalidQoSCombination CauseRadioNetwork = 18
	CauseRadioNetwork_ProcedureCancelled CauseRadioNetwork = 19
	CauseRadioNetwork_NormalRelease CauseRadioNetwork = 20
	CauseRadioNetwork_NoRadioResourcesAvailable CauseRadioNetwork = 21
	CauseRadioNetwork_ActionEsirableForAdioEasons CauseRadioNetwork = 22
	CauseRadioNetwork_DRRResourcesNotAvailableForTheSlice CauseRadioNetwork = 23
	CauseRadioNetwork_PDCPConfigurationNotSupported CauseRadioNetwork = 24
	CauseRadioNetwork_UeDlMaxIpDataRateReason CauseRadioNetwork = 25
	CauseRadioNetwork_UPIntegrityProtectionFailure CauseRadioNetwork = 26
	CauseRadioNetwork_ReleaseDueToPreEmption CauseRadioNetwork = 27
	CauseRadioNetwork_RsnNotAvailableForTheUp CauseRadioNetwork = 28
	CauseRadioNetwork_NPNNotSupported CauseRadioNetwork = 29
)

type CauseTransport int32
const (
	CauseTransport_Unspecified CauseTransport = 0
	CauseTransport_TransportResourceUnavailable CauseTransport = 1
	CauseTransport_UnknownTnlAddressForIab CauseTransport = 2
)

type CipheringAlgorithm int32
const (
	CipheringAlgorithm_NEA0 CipheringAlgorithm = 0
	CipheringAlgorithm_C128Nea1 CipheringAlgorithm = 1
	CipheringAlgorithm_C128Nea2 CipheringAlgorithm = 2
	CipheringAlgorithm_C128Nea3 CipheringAlgorithm = 3
)

type ConfidentialityProtectionIndication int32
const (
	ConfidentialityProtectionIndication_Required ConfidentialityProtectionIndication = 0
	ConfidentialityProtectionIndication_Preferred ConfidentialityProtectionIndication = 1
	ConfidentialityProtectionIndication_NotNeeded ConfidentialityProtectionIndication = 2
)

type ConfidentialityProtectionResult int32
const (
	ConfidentialityProtectionResult_Performed ConfidentialityProtectionResult = 0
	ConfidentialityProtectionResult_NotPerformed ConfidentialityProtectionResult = 1
)

type Criticality int32
const (
	Criticality_Reject Criticality = 0
	Criticality_Ignore Criticality = 1
	Criticality_Notify Criticality = 2
)

type DlTxStop int32
const (
	DlTxStop_Stop DlTxStop = 0
	DlTxStop_Resume DlTxStop = 1
)

type DrbActivity int32
const (
	DrbActivity_Active DrbActivity = 0
	DrbActivity_NotActive DrbActivity = 1
)

type DataForwardingRequest int32
const (
	DataForwardingRequest_UL DataForwardingRequest = 0
	DataForwardingRequest_DL DataForwardingRequest = 1
	DataForwardingRequest_Both DataForwardingRequest = 2
)

type DataDiscardRequired int32
const (
	DataDiscardRequired_Required DataDiscardRequired = 0
)

type DefaultDRB int32
const (
	DefaultDRB_True DefaultDRB = 0
	DefaultDRB_False DefaultDRB = 1
)

type DiscardTimer int32
const (
	DiscardTimer_Ms10 DiscardTimer = 0
	DiscardTimer_Ms20 DiscardTimer = 1
	DiscardTimer_Ms30 DiscardTimer = 2
	DiscardTimer_Ms40 DiscardTimer = 3
	DiscardTimer_Ms50 DiscardTimer = 4
	DiscardTimer_Ms60 DiscardTimer = 5
	DiscardTimer_Ms75 DiscardTimer = 6
	DiscardTimer_Ms100 DiscardTimer = 7
	DiscardTimer_Ms150 DiscardTimer = 8
	DiscardTimer_Ms200 DiscardTimer = 9
	DiscardTimer_Ms250 DiscardTimer = 10
	DiscardTimer_Ms300 DiscardTimer = 11
	DiscardTimer_Ms500 DiscardTimer = 12
	DiscardTimer_Ms750 DiscardTimer = 13
	DiscardTimer_Ms1500 DiscardTimer = 14
	DiscardTimer_Infinity DiscardTimer = 15
)

type DuplicationActivation int32
const (
	DuplicationActivation_Active DuplicationActivation = 0
	DuplicationActivation_Inactive DuplicationActivation = 1
)

type EarlyForwardingCOUNTReq int32
const (
	EarlyForwardingCOUNTReq_FirstDlCount EarlyForwardingCOUNTReq = 0
	EarlyForwardingCOUNTReq_DlDiscarding EarlyForwardingCOUNTReq = 1
)

type GnbCuUpOverloadInformation int32
const (
	GnbCuUpOverloadInformation_Overloaded GnbCuUpOverloadInformation = 0
	GnbCuUpOverloadInformation_NotOverloaded GnbCuUpOverloadInformation = 1
)

type IntegrityProtectionAlgorithm int32
const (
	IntegrityProtectionAlgorithm_NIA0 IntegrityProtectionAlgorithm = 0
	IntegrityProtectionAlgorithm_I128Nia1 IntegrityProtectionAlgorithm = 1
	IntegrityProtectionAlgorithm_I128Nia2 IntegrityProtectionAlgorithm = 2
	IntegrityProtectionAlgorithm_I128Nia3 IntegrityProtectionAlgorithm = 3
)

type IntegrityProtectionIndication int32
const (
	IntegrityProtectionIndication_Required IntegrityProtectionIndication = 0
	IntegrityProtectionIndication_Preferred IntegrityProtectionIndication = 1
	IntegrityProtectionIndication_NotNeeded IntegrityProtectionIndication = 2
)

type IntegrityProtectionResult int32
const (
	IntegrityProtectionResult_Performed IntegrityProtectionResult = 0
	IntegrityProtectionResult_NotPerformed IntegrityProtectionResult = 1
)

type LinksToLog int32
const (
	LinksToLog_Uplink LinksToLog = 0
	LinksToLog_Downlink LinksToLog = 1
	LinksToLog_BothUplinkAndDownlink LinksToLog = 2
)

type M4period int32
const (
	M4period_Ms1024 M4period = 0
	M4period_Ms2048 M4period = 1
	M4period_Ms5120 M4period = 2
	M4period_Ms10240 M4period = 3
	M4period_Min1 M4period = 4
)

type M6reportInterval int32
const (
	M6reportInterval_Ms120 M6reportInterval = 0
	M6reportInterval_Ms240 M6reportInterval = 1
	M6reportInterval_Ms480 M6reportInterval = 2
	M6reportInterval_Ms640 M6reportInterval = 3
	M6reportInterval_Ms1024 M6reportInterval = 4
	M6reportInterval_Ms2048 M6reportInterval = 5
	M6reportInterval_Ms5120 M6reportInterval = 6
	M6reportInterval_Ms10240 M6reportInterval = 7
	M6reportInterval_Ms20480 M6reportInterval = 8
	M6reportInterval_Ms40960 M6reportInterval = 9
	M6reportInterval_Min1 M6reportInterval = 10
	M6reportInterval_Min6 M6reportInterval = 11
	M6reportInterval_Min12 M6reportInterval = 12
	M6reportInterval_Min30 M6reportInterval = 13
)

type MdtActivation int32
const (
	MdtActivation_ImmediateMdtOnly MdtActivation = 0
	MdtActivation_ImmediateMdtAndTrace MdtActivation = 1
)

type MaxIPrate int32
const (
	MaxIPrate_Bitrate64kbs MaxIPrate = 0
	MaxIPrate_MaxUErate MaxIPrate = 1
)

type NewUlTnlInformationRequired int32
const (
	NewUlTnlInformationRequired_Required NewUlTnlInformationRequired = 0
)

type OutOfOrderDelivery int32
const (
	OutOfOrderDelivery_True OutOfOrderDelivery = 0
)

type PdcpDataRecovery int32
const (
	PdcpDataRecovery_True PdcpDataRecovery = 0
)

type PdcpDuplication int32
const (
	PdcpDuplication_True PdcpDuplication = 0
)

type PdcpReestablishment int32
const (
	PdcpReestablishment_True PdcpReestablishment = 0
)

type PdcpSnSize int32
const (
	PdcpSnSize_S12 PdcpSnSize = 0
	PdcpSnSize_S18 PdcpSnSize = 1
)

type PdcpSnStatusRequest int32
const (
	PdcpSnStatusRequest_Requested PdcpSnStatusRequest = 0
)

type PdcpStatusReportIndication int32
const (
	PdcpStatusReportIndication_Downlink PdcpStatusReportIndication = 0
	PdcpStatusReportIndication_Uplink PdcpStatusReportIndication = 1
	PdcpStatusReportIndication_Both PdcpStatusReportIndication = 2
)

type PduSessionResourceActivity int32
const (
	PduSessionResourceActivity_Active PduSessionResourceActivity = 0
	PduSessionResourceActivity_NotActive PduSessionResourceActivity = 1
)

type PduSessionType int32
const (
	PduSessionType_Ipv4 PduSessionType = 0
	PduSessionType_Ipv6 PduSessionType = 1
	PduSessionType_Ipv4v6 PduSessionType = 2
	PduSessionType_Ethernet PduSessionType = 3
	PduSessionType_Unstructured PduSessionType = 4
)

type PreEmptionCapability int32
const (
	PreEmptionCapability_ShallNotTriggerPreEmption PreEmptionCapability = 0
	PreEmptionCapability_MayTriggerPreEmption PreEmptionCapability = 1
)

type PreEmptionVulnerability int32
const (
	PreEmptionVulnerability_NotPreEmptable PreEmptionVulnerability = 0
	PreEmptionVulnerability_PreEmptable PreEmptionVulnerability = 1
)

type Presence int32
const (
	Presence_Optional Presence = 0
	Presence_Conditional Presence = 1
	Presence_Mandatory Presence = 2
)

type PrivacyIndicator int32
const (
	PrivacyIndicator_ImmediateMdt PrivacyIndicator = 0
	PrivacyIndicator_LoggedMdt PrivacyIndicator = 1
)

type QoSFlowMappingIndication int32
const (
	QoSFlowMappingIndication_Ul QoSFlowMappingIndication = 0
	QoSFlowMappingIndication_Dl QoSFlowMappingIndication = 1
)

type QosMonitoringRequest int32
const (
	QosMonitoringRequest_Ul QosMonitoringRequest = 0
	QosMonitoringRequest_Dl QosMonitoringRequest = 1
	QosMonitoringRequest_Both QosMonitoringRequest = 2
)

type RatType int32
const (
	RatType_EUtra RatType = 0
	RatType_NR RatType = 1
)

type RlcMode int32
const (
	RlcMode_RlcTm RlcMode = 0
	RlcMode_RlcAm RlcMode = 1
	RlcMode_RlcUmBidirectional RlcMode = 2
	RlcMode_RlcUmUnidirectionalUl RlcMode = 3
	RlcMode_RlcUmUnidirectionalDl RlcMode = 4
)

type Rsn int32
const (
	Rsn_V1 Rsn = 0
	Rsn_V2 Rsn = 1
)

type RedundantQoSFlowIndicator int32
const (
	RedundantQoSFlowIndicator_True RedundantQoSFlowIndicator = 0
	RedundantQoSFlowIndicator_False RedundantQoSFlowIndicator = 1
)

type RegistrationRequest int32
const (
	RegistrationRequest_Start RegistrationRequest = 0
	RegistrationRequest_Stop RegistrationRequest = 1
)

type ReportingPeriodicity int32
const (
	ReportingPeriodicity_Ms500 ReportingPeriodicity = 0
	ReportingPeriodicity_Ms1000 ReportingPeriodicity = 1
	ReportingPeriodicity_Ms2000 ReportingPeriodicity = 2
	ReportingPeriodicity_Ms5000 ReportingPeriodicity = 3
	ReportingPeriodicity_Ms10000 ReportingPeriodicity = 4
	ReportingPeriodicity_Ms20000 ReportingPeriodicity = 5
	ReportingPeriodicity_Ms30000 ReportingPeriodicity = 6
	ReportingPeriodicity_Ms40000 ReportingPeriodicity = 7
	ReportingPeriodicity_Ms50000 ReportingPeriodicity = 8
	ReportingPeriodicity_Ms60000 ReportingPeriodicity = 9
	ReportingPeriodicity_Ms70000 ReportingPeriodicity = 10
	ReportingPeriodicity_Ms80000 ReportingPeriodicity = 11
	ReportingPeriodicity_Ms90000 ReportingPeriodicity = 12
	ReportingPeriodicity_Ms100000 ReportingPeriodicity = 13
	ReportingPeriodicity_Ms110000 ReportingPeriodicity = 14
	ReportingPeriodicity_Ms120000 ReportingPeriodicity = 15
)

type ResetAll int32
const (
	ResetAll_ResetAll ResetAll = 0
)

type SdapHeaderDl int32
const (
	SdapHeaderDl_Present SdapHeaderDl = 0
	SdapHeaderDl_Absent SdapHeaderDl = 1
)

type SdapHeaderUl int32
const (
	SdapHeaderUl_Present SdapHeaderUl = 0
	SdapHeaderUl_Absent SdapHeaderUl = 1
)

type TReordering int32
const (
	TReordering_Ms0 TReordering = 0
	TReordering_Ms1 TReordering = 1
	TReordering_Ms2 TReordering = 2
	TReordering_Ms4 TReordering = 3
	TReordering_Ms5 TReordering = 4
	TReordering_Ms8 TReordering = 5
	TReordering_Ms10 TReordering = 6
	TReordering_Ms15 TReordering = 7
	TReordering_Ms20 TReordering = 8
	TReordering_Ms30 TReordering = 9
	TReordering_Ms40 TReordering = 10
	TReordering_Ms50 TReordering = 11
	TReordering_Ms60 TReordering = 12
	TReordering_Ms80 TReordering = 13
	TReordering_Ms100 TReordering = 14
	TReordering_Ms120 TReordering = 15
	TReordering_Ms140 TReordering = 16
	TReordering_Ms160 TReordering = 17
	TReordering_Ms180 TReordering = 18
	TReordering_Ms200 TReordering = 19
	TReordering_Ms220 TReordering = 20
	TReordering_Ms240 TReordering = 21
	TReordering_Ms260 TReordering = 22
	TReordering_Ms280 TReordering = 23
	TReordering_Ms300 TReordering = 24
	TReordering_Ms500 TReordering = 25
	TReordering_Ms750 TReordering = 26
	TReordering_Ms1000 TReordering = 27
	TReordering_Ms1250 TReordering = 28
	TReordering_Ms1500 TReordering = 29
	TReordering_Ms1750 TReordering = 30
	TReordering_Ms2000 TReordering = 31
	TReordering_Ms2250 TReordering = 32
	TReordering_Ms2500 TReordering = 33
	TReordering_Ms2750 TReordering = 34
	TReordering_Ms3000 TReordering = 35
)

type TNLAssociationUsage int32
const (
	TNLAssociationUsage_Ue TNLAssociationUsage = 0
	TNLAssociationUsage_NonUe TNLAssociationUsage = 1
	TNLAssociationUsage_Both TNLAssociationUsage = 2
)

type TimeToWait int32
const (
	TimeToWait_V1s TimeToWait = 0
	TimeToWait_V2s TimeToWait = 1
	TimeToWait_V5s TimeToWait = 2
	TimeToWait_V10s TimeToWait = 3
	TimeToWait_V20s TimeToWait = 4
	TimeToWait_V60s TimeToWait = 5
)

type TraceDepth int32
const (
	TraceDepth_Minimum TraceDepth = 0
	TraceDepth_Medium TraceDepth = 1
	TraceDepth_Maximum TraceDepth = 2
	TraceDepth_MinimumWithoutVendorSpecificExtension TraceDepth = 3
	TraceDepth_MediumWithoutVendorSpecificExtension TraceDepth = 4
	TraceDepth_MaximumWithoutVendorSpecificExtension TraceDepth = 5
)

type TriggeringMessage int32
const (
	TriggeringMessage_InitiatingMessage TriggeringMessage = 0
	TriggeringMessage_SuccessfulOutcome TriggeringMessage = 1
	TriggeringMessage_UnsuccessfulOutcome TriggeringMessage = 2
)

type TypeOfError int32
const (
	TypeOfError_NotUnderstood TypeOfError = 0
	TypeOfError_Missing TypeOfError = 1
)

type UeActivity int32
const (
	UeActivity_Active UeActivity = 0
	UeActivity_NotActive UeActivity = 1
)

type UlConfiguration int32
const (
	UlConfiguration_NoData UlConfiguration = 0
	UlConfiguration_Shared UlConfiguration = 1
	UlConfiguration_Only UlConfiguration = 2
)

type ULDataSplitThreshold int32
const (
	ULDataSplitThreshold_B0 ULDataSplitThreshold = 0
	ULDataSplitThreshold_B100 ULDataSplitThreshold = 1
	ULDataSplitThreshold_B200 ULDataSplitThreshold = 2
	ULDataSplitThreshold_B400 ULDataSplitThreshold = 3
	ULDataSplitThreshold_B800 ULDataSplitThreshold = 4
	ULDataSplitThreshold_B1600 ULDataSplitThreshold = 5
	ULDataSplitThreshold_B3200 ULDataSplitThreshold = 6
	ULDataSplitThreshold_B6400 ULDataSplitThreshold = 7
	ULDataSplitThreshold_B12800 ULDataSplitThreshold = 8
	ULDataSplitThreshold_B25600 ULDataSplitThreshold = 9
	ULDataSplitThreshold_B51200 ULDataSplitThreshold = 10
	ULDataSplitThreshold_B102400 ULDataSplitThreshold = 11
	ULDataSplitThreshold_B204800 ULDataSplitThreshold = 12
	ULDataSplitThreshold_B409600 ULDataSplitThreshold = 13
	ULDataSplitThreshold_B819200 ULDataSplitThreshold = 14
	ULDataSplitThreshold_B1228800 ULDataSplitThreshold = 15
	ULDataSplitThreshold_B1638400 ULDataSplitThreshold = 16
	ULDataSplitThreshold_B2457600 ULDataSplitThreshold = 17
	ULDataSplitThreshold_B3276800 ULDataSplitThreshold = 18
	ULDataSplitThreshold_B4096000 ULDataSplitThreshold = 19
	ULDataSplitThreshold_B4915200 ULDataSplitThreshold = 20
	ULDataSplitThreshold_B5734400 ULDataSplitThreshold = 21
	ULDataSplitThreshold_B6553600 ULDataSplitThreshold = 22
	ULDataSplitThreshold_Infinity ULDataSplitThreshold = 23
)

type AdditionalQoSInformation int32
const (
	AdditionalQoSInformation_Present AdditionalQoSInformation = 0
)

type ContinueROHC int32
const (
	ContinueROHC_Present ContinueROHC = 0
)

type DRBReleasedInSession int32
const (
	DRBReleasedInSession_ReleasedInSession DRBReleasedInSession = 0
	DRBReleasedInSession_NotReleasedInSession DRBReleasedInSession = 1
)

type DapsIndicator int32
const (
	DapsIndicator_Present DapsIndicator = 0
)

type DelayCritical int32
const (
	DelayCritical_Present DelayCritical = 0
)

type DrbContinueEHCDl int32
const (
	DrbContinueEHCDl_Present DrbContinueEHCDl = 0
)

type DrbContinueEHCUl int32
const (
	DrbContinueEHCUl_Present DrbContinueEHCUl = 0
)

type EhcCidLength int32
const (
	EhcCidLength_Present EhcCidLength = 0
)

type NGDlUpUnchanged int32
const (
	NGDlUpUnchanged_Present NGDlUpUnchanged = 0
)

type QoSFlowReleasedInSession int32
const (
	QoSFlowReleasedInSession_Present QoSFlowReleasedInSession = 0
)

type ReflectiveQoSAttribute int32
const (
	ReflectiveQoSAttribute_Present ReflectiveQoSAttribute = 0
)

type ReflectiveQoSIndicator int32
const (
	ReflectiveQoSIndicator_Present ReflectiveQoSIndicator = 0
)

type S1DlUpUnchanged int32
const (
	S1DlUpUnchanged_Present S1DlUpUnchanged = 0
)

type SecondaryRATType int32
const (
	SecondaryRATType_Nr SecondaryRATType = 0
	SecondaryRATType_EUtra SecondaryRATType = 1
)

