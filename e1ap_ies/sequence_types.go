package e1ap_ies

// Alternativeqosparasetitem represents the ASN.1 SEQUENCE type.
type Alternativeqosparasetitem struct {
	Guaranteedflowbitratedl *Bitrate `json:"guaranteedFlowBitRateDL,omitempty"`
	Guaranteedflowbitrateul *Bitrate `json:"guaranteedFlowBitRateUL,omitempty"`
	Packetdelaybudget *Packetdelaybudget `json:"packetDelayBudget,omitempty"`
	Packeterrorrate *Packeterrorrate `json:"packetErrorRate,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// Bearercontextinactivitynotification represents the ASN.1 SEQUENCE type.
type Bearercontextinactivitynotification struct {
	Protocolies ProtocolieContainer `json:"protocolIEs,omitempty"`
}

// Bearercontextmodificationconfirm represents the ASN.1 SEQUENCE type.
type Bearercontextmodificationconfirm struct {
	Protocolies ProtocolieContainer `json:"protocolIEs,omitempty"`
}

// Bearercontextmodificationfailure represents the ASN.1 SEQUENCE type.
type Bearercontextmodificationfailure struct {
	Protocolies ProtocolieContainer `json:"protocolIEs,omitempty"`
}

// Bearercontextmodificationrequest represents the ASN.1 SEQUENCE type.
type Bearercontextmodificationrequest struct {
	Protocolies ProtocolieContainer `json:"protocolIEs,omitempty"`
}

// Bearercontextmodificationrequired represents the ASN.1 SEQUENCE type.
type Bearercontextmodificationrequired struct {
	Protocolies ProtocolieContainer `json:"protocolIEs,omitempty"`
}

// Bearercontextmodificationresponse represents the ASN.1 SEQUENCE type.
type Bearercontextmodificationresponse struct {
	Protocolies ProtocolieContainer `json:"protocolIEs,omitempty"`
}

// Bearercontextreleasecommand represents the ASN.1 SEQUENCE type.
type Bearercontextreleasecommand struct {
	Protocolies ProtocolieContainer `json:"protocolIEs,omitempty"`
}

// Bearercontextreleasecomplete represents the ASN.1 SEQUENCE type.
type Bearercontextreleasecomplete struct {
	Protocolies ProtocolieContainer `json:"protocolIEs,omitempty"`
}

// Bearercontextreleaserequest represents the ASN.1 SEQUENCE type.
type Bearercontextreleaserequest struct {
	Protocolies ProtocolieContainer `json:"protocolIEs,omitempty"`
}

// Bearercontextsetupfailure represents the ASN.1 SEQUENCE type.
type Bearercontextsetupfailure struct {
	Protocolies ProtocolieContainer `json:"protocolIEs,omitempty"`
}

// Bearercontextsetuprequest represents the ASN.1 SEQUENCE type.
type Bearercontextsetuprequest struct {
	Protocolies ProtocolieContainer `json:"protocolIEs,omitempty"`
}

// Bearercontextsetupresponse represents the ASN.1 SEQUENCE type.
type Bearercontextsetupresponse struct {
	Protocolies ProtocolieContainer `json:"protocolIEs,omitempty"`
}

// CellGroupInformationItem represents the ASN.1 SEQUENCE type.
type CellGroupInformationItem struct {
	CellGroupId CellGroupId `json:"cell-Group-ID,omitempty"`
	UlConfiguration *UlConfiguration `json:"uL-Configuration,omitempty"`
	DlTxStop *DlTxStop `json:"dL-TX-Stop,omitempty"`
	RatType *RatType `json:"rAT-Type,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// Celltraffictrace represents the ASN.1 SEQUENCE type.
type Celltraffictrace struct {
	Protocolies ProtocolieContainer `json:"protocolIEs,omitempty"`
}

// Criticalitydiagnostics represents the ASN.1 SEQUENCE type.
type Criticalitydiagnostics struct {
	Procedurecode *Procedurecode `json:"procedureCode,omitempty"`
	Triggeringmessage *Triggeringmessage `json:"triggeringMessage,omitempty"`
	Procedurecriticality *Criticality `json:"procedureCriticality,omitempty"`
	Transactionid *Transactionid `json:"transactionID,omitempty"`
	Iescriticalitydiagnostics *CriticalitydiagnosticsIeList `json:"iEsCriticalityDiagnostics,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// Dapsrequestinfo represents the ASN.1 SEQUENCE type.
type Dapsrequestinfo struct {
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// Dldatanotification represents the ASN.1 SEQUENCE type.
type Dldatanotification struct {
	Protocolies ProtocolieContainer `json:"protocolIEs,omitempty"`
}

// Dldiscarding represents the ASN.1 SEQUENCE type.
type Dldiscarding struct {
	Dldiscardingcountval PdcpCount `json:"dLDiscardingCountVal,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// Dluptnladdresstoupdateitem represents the ASN.1 SEQUENCE type.
type Dluptnladdresstoupdateitem struct {
	Oldtnladress Transportlayeraddress `json:"oldTNLAdress,omitempty"`
	Newtnladress Transportlayeraddress `json:"newTNLAdress,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// DrbActivityItem represents the ASN.1 SEQUENCE type.
type DrbActivityItem struct {
	DrbId DrbId `json:"dRB-ID,omitempty"`
	DrbActivity DrbActivity `json:"dRB-Activity,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// DrbConfirmModifiedItemEutran represents the ASN.1 SEQUENCE type.
type DrbConfirmModifiedItemEutran struct {
	DrbId DrbId `json:"dRB-ID,omitempty"`
	CellGroupInformation *CellGroupInformation `json:"cell-Group-Information,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// DrbConfirmModifiedItemNgRan represents the ASN.1 SEQUENCE type.
type DrbConfirmModifiedItemNgRan struct {
	DrbId DrbId `json:"dRB-ID,omitempty"`
	CellGroupInformation *CellGroupInformation `json:"cell-Group-Information,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// DrbFailedItemEutran represents the ASN.1 SEQUENCE type.
type DrbFailedItemEutran struct {
	DrbId DrbId `json:"dRB-ID,omitempty"`
	Cause Cause `json:"cause,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// DrbFailedItemNgRan represents the ASN.1 SEQUENCE type.
type DrbFailedItemNgRan struct {
	DrbId DrbId `json:"dRB-ID,omitempty"`
	Cause Cause `json:"cause,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// DrbFailedModItemEutran represents the ASN.1 SEQUENCE type.
type DrbFailedModItemEutran struct {
	DrbId DrbId `json:"dRB-ID,omitempty"`
	Cause Cause `json:"cause,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// DrbFailedModItemNgRan represents the ASN.1 SEQUENCE type.
type DrbFailedModItemNgRan struct {
	DrbId DrbId `json:"dRB-ID,omitempty"`
	Cause Cause `json:"cause,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// DrbFailedToModifyItemEutran represents the ASN.1 SEQUENCE type.
type DrbFailedToModifyItemEutran struct {
	DrbId DrbId `json:"dRB-ID,omitempty"`
	Cause Cause `json:"cause,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// DrbFailedToModifyItemNgRan represents the ASN.1 SEQUENCE type.
type DrbFailedToModifyItemNgRan struct {
	DrbId DrbId `json:"dRB-ID,omitempty"`
	Cause Cause `json:"cause,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// DrbModifiedItemEutran represents the ASN.1 SEQUENCE type.
type DrbModifiedItemEutran struct {
	DrbId DrbId `json:"dRB-ID,omitempty"`
	S1DlUpTnlInformation *UpTnlInformation `json:"s1-DL-UP-TNL-Information,omitempty"`
	PdcpSnStatusInformation *PdcpSnStatusInformation `json:"pDCP-SN-Status-Information,omitempty"`
	UlUpTransportParameters *UpParameters `json:"uL-UP-Transport-Parameters,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// DrbModifiedItemNgRan represents the ASN.1 SEQUENCE type.
type DrbModifiedItemNgRan struct {
	DrbId DrbId `json:"dRB-ID,omitempty"`
	UlUpTransportParameters *UpParameters `json:"uL-UP-Transport-Parameters,omitempty"`
	PdcpSnStatusInformation *PdcpSnStatusInformation `json:"pDCP-SN-Status-Information,omitempty"`
	FlowSetupList *QosFlowList `json:"flow-Setup-List,omitempty"`
	FlowFailedList *QosFlowFailedList `json:"flow-Failed-List,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// DrbRemovedItem represents the ASN.1 SEQUENCE type.
type DrbRemovedItem struct {
	DrbId DrbId `json:"dRB-ID,omitempty"`
	DrbAccumulatedSessionTime *DrbRemovedItemDrbAccumulatedSessionTime `json:"dRB-Accumulated-Session-Time,omitempty"`
	QosFlowRemovedList *Sequence `json:"qoS-Flow-Removed-List,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// DrbRequiredToModifyItemEutran represents the ASN.1 SEQUENCE type.
type DrbRequiredToModifyItemEutran struct {
	DrbId DrbId `json:"dRB-ID,omitempty"`
	S1DlUpTnlInformation *UpTnlInformation `json:"s1-DL-UP-TNL-Information,omitempty"`
	GnbCuUpCellgrouprelatedconfiguration *GnbCuUpCellgrouprelatedconfiguration `json:"gNB-CU-UP-CellGroupRelatedConfiguration,omitempty"`
	Cause *Cause `json:"cause,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// DrbRequiredToModifyItemNgRan represents the ASN.1 SEQUENCE type.
type DrbRequiredToModifyItemNgRan struct {
	DrbId DrbId `json:"dRB-ID,omitempty"`
	GnbCuUpCellgrouprelatedconfiguration *GnbCuUpCellgrouprelatedconfiguration `json:"gNB-CU-UP-CellGroupRelatedConfiguration,omitempty"`
	FlowToRemove *QosFlowList `json:"flow-To-Remove,omitempty"`
	Cause *Cause `json:"cause,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// DrbRequiredToRemoveItemEutran represents the ASN.1 SEQUENCE type.
type DrbRequiredToRemoveItemEutran struct {
	DrbId DrbId `json:"dRB-ID,omitempty"`
	Cause Cause `json:"cause,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// DrbRequiredToRemoveItemNgRan represents the ASN.1 SEQUENCE type.
type DrbRequiredToRemoveItemNgRan struct {
	DrbId DrbId `json:"dRB-ID,omitempty"`
	Cause Cause `json:"cause,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// DrbSetupItemEutran represents the ASN.1 SEQUENCE type.
type DrbSetupItemEutran struct {
	DrbId DrbId `json:"dRB-ID,omitempty"`
	S1DlUpTnlInformation UpTnlInformation `json:"s1-DL-UP-TNL-Information,omitempty"`
	DataForwardingInformationResponse *DataForwardingInformation `json:"data-Forwarding-Information-Response,omitempty"`
	UlUpTransportParameters UpParameters `json:"uL-UP-Transport-Parameters,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// DrbSetupItemNgRan represents the ASN.1 SEQUENCE type.
type DrbSetupItemNgRan struct {
	DrbId DrbId `json:"dRB-ID,omitempty"`
	DrbDataForwardingInformationResponse *DataForwardingInformation `json:"dRB-data-Forwarding-Information-Response,omitempty"`
	UlUpTransportParameters UpParameters `json:"uL-UP-Transport-Parameters,omitempty"`
	FlowSetupList QosFlowList `json:"flow-Setup-List,omitempty"`
	FlowFailedList *QosFlowFailedList `json:"flow-Failed-List,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// DrbSetupModItemEutran represents the ASN.1 SEQUENCE type.
type DrbSetupModItemEutran struct {
	DrbId DrbId `json:"dRB-ID,omitempty"`
	S1DlUpTnlInformation UpTnlInformation `json:"s1-DL-UP-TNL-Information,omitempty"`
	DataForwardingInformationResponse *DataForwardingInformation `json:"data-Forwarding-Information-Response,omitempty"`
	UlUpTransportParameters UpParameters `json:"uL-UP-Transport-Parameters,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// DrbSetupModItemNgRan represents the ASN.1 SEQUENCE type.
type DrbSetupModItemNgRan struct {
	DrbId DrbId `json:"dRB-ID,omitempty"`
	DrbDataForwardingInformationResponse *DataForwardingInformation `json:"dRB-data-Forwarding-Information-Response,omitempty"`
	UlUpTransportParameters UpParameters `json:"uL-UP-Transport-Parameters,omitempty"`
	FlowSetupList QosFlowList `json:"flow-Setup-List,omitempty"`
	FlowFailedList *QosFlowFailedList `json:"flow-Failed-List,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// DrbStatusItem represents the ASN.1 SEQUENCE type.
type DrbStatusItem struct {
	DrbId DrbId `json:"dRB-ID,omitempty"`
	PdcpDlCount *PdcpCount `json:"pDCP-DL-Count,omitempty"`
	PdcpUlCount *PdcpCount `json:"pDCP-UL-Count,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// DrbToModifyItemEutran represents the ASN.1 SEQUENCE type.
type DrbToModifyItemEutran struct {
	DrbId DrbId `json:"dRB-ID,omitempty"`
	PdcpConfiguration *PdcpConfiguration `json:"pDCP-Configuration,omitempty"`
	EutranQos *EutranQos `json:"eUTRAN-QoS,omitempty"`
	S1UlUpTnlInformation *UpTnlInformation `json:"s1-UL-UP-TNL-Information,omitempty"`
	DataForwardingInformation *DataForwardingInformation `json:"data-Forwarding-Information,omitempty"`
	PdcpSnStatusRequest *PdcpSnStatusRequest `json:"pDCP-SN-Status-Request,omitempty"`
	PdcpSnStatusInformation *PdcpSnStatusInformation `json:"pDCP-SN-Status-Information,omitempty"`
	DlUpParameters *UpParameters `json:"dL-UP-Parameters,omitempty"`
	CellGroupToAdd *CellGroupInformation `json:"cell-Group-To-Add,omitempty"`
	CellGroupToModify *CellGroupInformation `json:"cell-Group-To-Modify,omitempty"`
	CellGroupToRemove *CellGroupInformation `json:"cell-Group-To-Remove,omitempty"`
	DrbInactivityTimer *InactivityTimer `json:"dRB-Inactivity-Timer,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// DrbToModifyItemNgRan represents the ASN.1 SEQUENCE type.
type DrbToModifyItemNgRan struct {
	DrbId DrbId `json:"dRB-ID,omitempty"`
	SdapConfiguration *SdapConfiguration `json:"sDAP-Configuration,omitempty"`
	PdcpConfiguration *PdcpConfiguration `json:"pDCP-Configuration,omitempty"`
	DrbDataForwardingInformation *DataForwardingInformation `json:"dRB-Data-Forwarding-Information,omitempty"`
	PdcpSnStatusRequest *PdcpSnStatusRequest `json:"pDCP-SN-Status-Request,omitempty"`
	PdcpSnStatusInformation *PdcpSnStatusInformation `json:"pdcp-SN-Status-Information,omitempty"`
	DlUpParameters *UpParameters `json:"dL-UP-Parameters,omitempty"`
	CellGroupToAdd *CellGroupInformation `json:"cell-Group-To-Add,omitempty"`
	CellGroupToModify *CellGroupInformation `json:"cell-Group-To-Modify,omitempty"`
	CellGroupToRemove *CellGroupInformation `json:"cell-Group-To-Remove,omitempty"`
	FlowMappingInformation *QosFlowQosParameterList `json:"flow-Mapping-Information,omitempty"`
	DrbInactivityTimer *InactivityTimer `json:"dRB-Inactivity-Timer,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// DrbToRemoveItemEutran represents the ASN.1 SEQUENCE type.
type DrbToRemoveItemEutran struct {
	DrbId DrbId `json:"dRB-ID,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// DrbToRemoveItemNgRan represents the ASN.1 SEQUENCE type.
type DrbToRemoveItemNgRan struct {
	DrbId DrbId `json:"dRB-ID,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// DrbToSetupItemEutran represents the ASN.1 SEQUENCE type.
type DrbToSetupItemEutran struct {
	DrbId DrbId `json:"dRB-ID,omitempty"`
	PdcpConfiguration PdcpConfiguration `json:"pDCP-Configuration,omitempty"`
	EutranQos EutranQos `json:"eUTRAN-QoS,omitempty"`
	S1UlUpTnlInformation UpTnlInformation `json:"s1-UL-UP-TNL-Information,omitempty"`
	DataForwardingInformationRequest *DataForwardingInformationRequest `json:"data-Forwarding-Information-Request,omitempty"`
	CellGroupInformation CellGroupInformation `json:"cell-Group-Information,omitempty"`
	DlUpParameters *UpParameters `json:"dL-UP-Parameters,omitempty"`
	DrbInactivityTimer *InactivityTimer `json:"dRB-Inactivity-Timer,omitempty"`
	ExistingAllocatedS1DlUpTnlInfo *UpTnlInformation `json:"existing-Allocated-S1-DL-UP-TNL-Info,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// DrbToSetupItemNgRan represents the ASN.1 SEQUENCE type.
type DrbToSetupItemNgRan struct {
	DrbId DrbId `json:"dRB-ID,omitempty"`
	SdapConfiguration SdapConfiguration `json:"sDAP-Configuration,omitempty"`
	PdcpConfiguration PdcpConfiguration `json:"pDCP-Configuration,omitempty"`
	CellGroupInformation CellGroupInformation `json:"cell-Group-Information,omitempty"`
	QosFlowInformationToBeSetup QosFlowQosParameterList `json:"qos-flow-Information-To-Be-Setup,omitempty"`
	DrbDataForwardingInformationRequest *DataForwardingInformationRequest `json:"dRB-Data-Forwarding-Information-Request,omitempty"`
	DrbInactivityTimer *InactivityTimer `json:"dRB-Inactivity-Timer,omitempty"`
	PdcpSnStatusInformation *PdcpSnStatusInformation `json:"pDCP-SN-Status-Information,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// DrbToSetupModItemEutran represents the ASN.1 SEQUENCE type.
type DrbToSetupModItemEutran struct {
	DrbId DrbId `json:"dRB-ID,omitempty"`
	PdcpConfiguration PdcpConfiguration `json:"pDCP-Configuration,omitempty"`
	EutranQos EutranQos `json:"eUTRAN-QoS,omitempty"`
	S1UlUpTnlInformation UpTnlInformation `json:"s1-UL-UP-TNL-Information,omitempty"`
	DataForwardingInformationRequest *DataForwardingInformationRequest `json:"data-Forwarding-Information-Request,omitempty"`
	CellGroupInformation CellGroupInformation `json:"cell-Group-Information,omitempty"`
	DlUpParameters *UpParameters `json:"dL-UP-Parameters,omitempty"`
	DrbInactivityTimer *InactivityTimer `json:"dRB-Inactivity-Timer,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// DrbToSetupModItemNgRan represents the ASN.1 SEQUENCE type.
type DrbToSetupModItemNgRan struct {
	DrbId DrbId `json:"dRB-ID,omitempty"`
	SdapConfiguration SdapConfiguration `json:"sDAP-Configuration,omitempty"`
	PdcpConfiguration PdcpConfiguration `json:"pDCP-Configuration,omitempty"`
	CellGroupInformation CellGroupInformation `json:"cell-Group-Information,omitempty"`
	FlowMappingInformation QosFlowQosParameterList `json:"flow-Mapping-Information,omitempty"`
	DrbDataForwardingInformationRequest *DataForwardingInformationRequest `json:"dRB-Data-Forwarding-Information-Request,omitempty"`
	DrbInactivityTimer *InactivityTimer `json:"dRB-Inactivity-Timer,omitempty"`
	PdcpSnStatusInformation *PdcpSnStatusInformation `json:"pDCP-SN-Status-Information,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// DrbUsageReportItem represents the ASN.1 SEQUENCE type.
type DrbUsageReportItem struct {
	Starttimestamp DrbUsageReportItemStarttimestamp `json:"startTimeStamp,omitempty"`
	Endtimestamp DrbUsageReportItemEndtimestamp `json:"endTimeStamp,omitempty"`
	Usagecountul DrbUsageReportItemUsagecountul `json:"usageCountUL,omitempty"`
	Usagecountdl DrbUsageReportItemUsagecountdl `json:"usageCountDL,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// Drbbstatustransfer represents the ASN.1 SEQUENCE type.
type Drbbstatustransfer struct {
	Receivestatusofpdcpsdu *DrbbstatustransferReceivestatusofpdcpsdu `json:"receiveStatusofPDCPSDU,omitempty"`
	Countvalue PdcpCount `json:"countValue,omitempty"`
	IeExtension *Protocolextensioncontainer `json:"iE-Extension,omitempty"`
}

// DrbsSubjectToCounterCheckItemEutran represents the ASN.1 SEQUENCE type.
type DrbsSubjectToCounterCheckItemEutran struct {
	DrbId DrbId `json:"dRB-ID,omitempty"`
	PdcpUlCount PdcpCount `json:"pDCP-UL-Count,omitempty"`
	PdcpDlCount PdcpCount `json:"pDCP-DL-Count,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// DrbsSubjectToCounterCheckItemNgRan represents the ASN.1 SEQUENCE type.
type DrbsSubjectToCounterCheckItemNgRan struct {
	PduSessionId PduSessionId `json:"pDU-Session-ID,omitempty"`
	DrbId DrbId `json:"dRB-ID,omitempty"`
	PdcpUlCount PdcpCount `json:"pDCP-UL-Count,omitempty"`
	PdcpDlCount PdcpCount `json:"pDCP-DL-Count,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// DrbsSubjectToEarlyForwardingItem represents the ASN.1 SEQUENCE type.
type DrbsSubjectToEarlyForwardingItem struct {
	DrbId DrbId `json:"dRB-ID,omitempty"`
	Dlcountvalue PdcpCount `json:"dLCountValue,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// DataForwardingInformation represents the ASN.1 SEQUENCE type.
type DataForwardingInformation struct {
	UlDataForwarding *UpTnlInformation `json:"uL-Data-Forwarding,omitempty"`
	DlDataForwarding *UpTnlInformation `json:"dL-Data-Forwarding,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// DataForwardingInformationRequest represents the ASN.1 SEQUENCE type.
type DataForwardingInformationRequest struct {
	DataForwardingRequest DataForwardingRequest `json:"data-Forwarding-Request,omitempty"`
	QosFlowsForwardedOnFwdTunnels *QosFlowMappingList `json:"qoS-Flows-Forwarded-On-Fwd-Tunnels,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// DataUsageReportItem represents the ASN.1 SEQUENCE type.
type DataUsageReportItem struct {
	DrbId DrbId `json:"dRB-ID,omitempty"`
	RatType RatType `json:"rAT-Type,omitempty"`
	DrbUsageReportList DrbUsageReportList `json:"dRB-Usage-Report-List,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// DataUsagePerPduSessionReport represents the ASN.1 SEQUENCE type.
type DataUsagePerPduSessionReport struct {
	PduSessionTimedReportList Sequence `json:"pDU-session-Timed-Report-List,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// DataUsagePerQosFlowItem represents the ASN.1 SEQUENCE type.
type DataUsagePerQosFlowItem struct {
	QosFlowIdentifier QosFlowIdentifier `json:"qoS-Flow-Identifier,omitempty"`
	QosFlowTimedReportList Sequence `json:"qoS-Flow-Timed-Report-List,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// Datausagereport represents the ASN.1 SEQUENCE type.
type Datausagereport struct {
	Protocolies ProtocolieContainer `json:"protocolIEs,omitempty"`
}

// Deactivatetrace represents the ASN.1 SEQUENCE type.
type Deactivatetrace struct {
	Protocolies ProtocolieContainer `json:"protocolIEs,omitempty"`
}

// Dynamic5qidescriptor represents the ASN.1 SEQUENCE type.
type Dynamic5qidescriptor struct {
	Qosprioritylevel Qosprioritylevel `json:"qoSPriorityLevel,omitempty"`
	Packetdelaybudget Packetdelaybudget `json:"packetDelayBudget,omitempty"`
	Packeterrorrate Packeterrorrate `json:"packetErrorRate,omitempty"`
	Delaycritical *Dynamic5qidescriptorDelaycritical `json:"delayCritical,omitempty"`
	Averagingwindow *Averagingwindow `json:"averagingWindow,omitempty"`
	Maxdataburstvolume *Maxdataburstvolume `json:"maxDataBurstVolume,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// E1releaserequest represents the ASN.1 SEQUENCE type.
type E1releaserequest struct {
	Protocolies ProtocolieContainer `json:"protocolIEs,omitempty"`
}

// E1releaseresponse represents the ASN.1 SEQUENCE type.
type E1releaseresponse struct {
	Protocolies ProtocolieContainer `json:"protocolIEs,omitempty"`
}

// EhcCommonParameters represents the ASN.1 SEQUENCE type.
type EhcCommonParameters struct {
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// EhcDownlinkParameters represents the ASN.1 SEQUENCE type.
type EhcDownlinkParameters struct {
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// EhcParameters represents the ASN.1 SEQUENCE type.
type EhcParameters struct {
	EhcCommon EhcCommonParameters `json:"ehc-Common,omitempty"`
	EhcDownlink *EhcDownlinkParameters `json:"ehc-Downlink,omitempty"`
	EhcUplink *EhcUplinkParameters `json:"ehc-Uplink,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// EhcUplinkParameters represents the ASN.1 SEQUENCE type.
type EhcUplinkParameters struct {
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// EutranQos represents the ASN.1 SEQUENCE type.
type EutranQos struct {
	Qci Qci `json:"qCI,omitempty"`
	Eutranallocationandretentionpriority Eutranallocationandretentionpriority `json:"eUTRANallocationAndRetentionPriority,omitempty"`
	Gbrqosinformation *GbrQosinformation `json:"gbrQosInformation,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// EutranQosSupportItem represents the ASN.1 SEQUENCE type.
type EutranQosSupportItem struct {
	EutranQos EutranQos `json:"eUTRAN-QoS,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// Eutranallocationandretentionpriority represents the ASN.1 SEQUENCE type.
type Eutranallocationandretentionpriority struct {
	Prioritylevel Prioritylevel `json:"priorityLevel,omitempty"`
	PreEmptioncapability PreEmptioncapability `json:"pre-emptionCapability,omitempty"`
	PreEmptionvulnerability PreEmptionvulnerability `json:"pre-emptionVulnerability,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// Earlyforwardingsntransfer represents the ASN.1 SEQUENCE type.
type Earlyforwardingsntransfer struct {
	Protocolies ProtocolieContainer `json:"protocolIEs,omitempty"`
}

// EndpointIpAddressAndPort represents the ASN.1 SEQUENCE type.
type EndpointIpAddressAndPort struct {
	EndpointIpAddress Transportlayeraddress `json:"endpoint-IP-Address,omitempty"`
	Portnumber Portnumber `json:"portNumber,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// Errorindication represents the ASN.1 SEQUENCE type.
type Errorindication struct {
	Protocolies ProtocolieContainer `json:"protocolIEs,omitempty"`
}

// Firstdlcount represents the ASN.1 SEQUENCE type.
type Firstdlcount struct {
	Firstdlcountval PdcpCount `json:"firstDLCountVal,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// GbrQosflowinformation represents the ASN.1 SEQUENCE type.
type GbrQosflowinformation struct {
	Maxflowbitratedownlink Bitrate `json:"maxFlowBitRateDownlink,omitempty"`
	Maxflowbitrateuplink Bitrate `json:"maxFlowBitRateUplink,omitempty"`
	Guaranteedflowbitratedownlink Bitrate `json:"guaranteedFlowBitRateDownlink,omitempty"`
	Guaranteedflowbitrateuplink Bitrate `json:"guaranteedFlowBitRateUplink,omitempty"`
	Maxpacketlossratedownlink *Maxpacketlossrate `json:"maxPacketLossRateDownlink,omitempty"`
	Maxpacketlossrateuplink *Maxpacketlossrate `json:"maxPacketLossRateUplink,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// GbrQosinformation represents the ASN.1 SEQUENCE type.
type GbrQosinformation struct {
	ERabMaximumbitratedl Bitrate `json:"e-RAB-MaximumBitrateDL,omitempty"`
	ERabMaximumbitrateul Bitrate `json:"e-RAB-MaximumBitrateUL,omitempty"`
	ERabGuaranteedbitratedl Bitrate `json:"e-RAB-GuaranteedBitrateDL,omitempty"`
	ERabGuaranteedbitrateul Bitrate `json:"e-RAB-GuaranteedBitrateUL,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// GnbCuCpConfigurationupdate represents the ASN.1 SEQUENCE type.
type GnbCuCpConfigurationupdate struct {
	Protocolies ProtocolieContainer `json:"protocolIEs,omitempty"`
}

// GnbCuCpConfigurationupdateacknowledge represents the ASN.1 SEQUENCE type.
type GnbCuCpConfigurationupdateacknowledge struct {
	Protocolies ProtocolieContainer `json:"protocolIEs,omitempty"`
}

// GnbCuCpConfigurationupdatefailure represents the ASN.1 SEQUENCE type.
type GnbCuCpConfigurationupdatefailure struct {
	Protocolies ProtocolieContainer `json:"protocolIEs,omitempty"`
}

// GnbCuCpE1setupfailure represents the ASN.1 SEQUENCE type.
type GnbCuCpE1setupfailure struct {
	Protocolies ProtocolieContainer `json:"protocolIEs,omitempty"`
}

// GnbCuCpE1setuprequest represents the ASN.1 SEQUENCE type.
type GnbCuCpE1setuprequest struct {
	Protocolies ProtocolieContainer `json:"protocolIEs,omitempty"`
}

// GnbCuCpE1setupresponse represents the ASN.1 SEQUENCE type.
type GnbCuCpE1setupresponse struct {
	Protocolies ProtocolieContainer `json:"protocolIEs,omitempty"`
}

// GnbCuCpTnlaFailedToSetupItem represents the ASN.1 SEQUENCE type.
type GnbCuCpTnlaFailedToSetupItem struct {
	Tnlassociationtransportlayeraddress CpTnlInformation `json:"tNLAssociationTransportLayerAddress,omitempty"`
	Cause Cause `json:"cause,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// GnbCuCpTnlaSetupItem represents the ASN.1 SEQUENCE type.
type GnbCuCpTnlaSetupItem struct {
	Tnlassociationtransportlayeraddress CpTnlInformation `json:"tNLAssociationTransportLayerAddress,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// GnbCuCpTnlaToAddItem represents the ASN.1 SEQUENCE type.
type GnbCuCpTnlaToAddItem struct {
	Tnlassociationtransportlayeraddress CpTnlInformation `json:"tNLAssociationTransportLayerAddress,omitempty"`
	Tnlassociationusage Tnlassociationusage `json:"tNLAssociationUsage,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// GnbCuCpTnlaToRemoveItem represents the ASN.1 SEQUENCE type.
type GnbCuCpTnlaToRemoveItem struct {
	Tnlassociationtransportlayeraddress CpTnlInformation `json:"tNLAssociationTransportLayerAddress,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// GnbCuCpTnlaToUpdateItem represents the ASN.1 SEQUENCE type.
type GnbCuCpTnlaToUpdateItem struct {
	Tnlassociationtransportlayeraddress CpTnlInformation `json:"tNLAssociationTransportLayerAddress,omitempty"`
	Tnlassociationusage *Tnlassociationusage `json:"tNLAssociationUsage,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// GnbCuUpCellgrouprelatedconfigurationItem represents the ASN.1 SEQUENCE type.
type GnbCuUpCellgrouprelatedconfigurationItem struct {
	CellGroupId CellGroupId `json:"cell-Group-ID,omitempty"`
	UpTnlInformation UpTnlInformation `json:"uP-TNL-Information,omitempty"`
	UlConfiguration *UlConfiguration `json:"uL-Configuration,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// GnbCuUpConfigurationupdate represents the ASN.1 SEQUENCE type.
type GnbCuUpConfigurationupdate struct {
	Protocolies ProtocolieContainer `json:"protocolIEs,omitempty"`
}

// GnbCuUpConfigurationupdateacknowledge represents the ASN.1 SEQUENCE type.
type GnbCuUpConfigurationupdateacknowledge struct {
	Protocolies ProtocolieContainer `json:"protocolIEs,omitempty"`
}

// GnbCuUpConfigurationupdatefailure represents the ASN.1 SEQUENCE type.
type GnbCuUpConfigurationupdatefailure struct {
	Protocolies ProtocolieContainer `json:"protocolIEs,omitempty"`
}

// GnbCuUpCountercheckrequest represents the ASN.1 SEQUENCE type.
type GnbCuUpCountercheckrequest struct {
	Protocolies ProtocolieContainer `json:"protocolIEs,omitempty"`
}

// GnbCuUpE1setupfailure represents the ASN.1 SEQUENCE type.
type GnbCuUpE1setupfailure struct {
	Protocolies ProtocolieContainer `json:"protocolIEs,omitempty"`
}

// GnbCuUpE1setuprequest represents the ASN.1 SEQUENCE type.
type GnbCuUpE1setuprequest struct {
	Protocolies ProtocolieContainer `json:"protocolIEs,omitempty"`
}

// GnbCuUpE1setupresponse represents the ASN.1 SEQUENCE type.
type GnbCuUpE1setupresponse struct {
	Protocolies ProtocolieContainer `json:"protocolIEs,omitempty"`
}

// GnbCuUpStatusindication represents the ASN.1 SEQUENCE type.
type GnbCuUpStatusindication struct {
	Protocolies ProtocolieContainer `json:"protocolIEs,omitempty"`
}

// GnbCuUpTnlaToRemoveItem represents the ASN.1 SEQUENCE type.
type GnbCuUpTnlaToRemoveItem struct {
	Tnlassociationtransportlayeraddress CpTnlInformation `json:"tNLAssociationTransportLayerAddress,omitempty"`
	Tnlassociationtransportlayeraddressgnbcucp *CpTnlInformation `json:"tNLAssociationTransportLayerAddressgNBCUCP,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// GtptlaItem represents the ASN.1 SEQUENCE type.
type GtptlaItem struct {
	Gtptransportlayeraddresses Transportlayeraddress `json:"gTPTransportLayerAddresses,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// Gtptunnel represents the ASN.1 SEQUENCE type.
type Gtptunnel struct {
	Transportlayeraddress Transportlayeraddress `json:"transportLayerAddress,omitempty"`
	GtpTeid GtpTeid `json:"gTP-TEID,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// HwCapacityindicator represents the ASN.1 SEQUENCE type.
type HwCapacityindicator struct {
	IeExtensions Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// IabUptnladdressupdate represents the ASN.1 SEQUENCE type.
type IabUptnladdressupdate struct {
	Protocolies ProtocolieContainer `json:"protocolIEs,omitempty"`
}

// IabUptnladdressupdateacknowledge represents the ASN.1 SEQUENCE type.
type IabUptnladdressupdateacknowledge struct {
	Protocolies ProtocolieContainer `json:"protocolIEs,omitempty"`
}

// IabUptnladdressupdatefailure represents the ASN.1 SEQUENCE type.
type IabUptnladdressupdatefailure struct {
	Protocolies ProtocolieContainer `json:"protocolIEs,omitempty"`
}

// Immediatemdt represents the ASN.1 SEQUENCE type.
type Immediatemdt struct {
	Measurementstoactivate Measurementstoactivate `json:"measurementsToActivate,omitempty"`
	Measurementfour *M4configuration `json:"measurementFour,omitempty"`
	Measurementsix *M6configuration `json:"measurementSix,omitempty"`
	Measurementseven *M7configuration `json:"measurementSeven,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// Initiatingmessage represents the ASN.1 SEQUENCE type.
type Initiatingmessage struct {
	Procedurecode E1apElementaryProcedureProcedurecode `json:"procedureCode,omitempty"`
	Criticality E1apElementaryProcedureCriticality `json:"criticality,omitempty"`
	Value E1apElementaryProcedureInitiatingmessage `json:"value,omitempty"`
}

// M4configuration represents the ASN.1 SEQUENCE type.
type M4configuration struct {
	M4period M4period `json:"m4period,omitempty"`
	M4LinksToLog LinksToLog `json:"m4-links-to-log,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// M6configuration represents the ASN.1 SEQUENCE type.
type M6configuration struct {
	M6reportInterval M6reportInterval `json:"m6report-Interval,omitempty"`
	M6LinksToLog LinksToLog `json:"m6-links-to-log,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// M7configuration represents the ASN.1 SEQUENCE type.
type M7configuration struct {
	M7period M7period `json:"m7period,omitempty"`
	M7LinksToLog LinksToLog `json:"m7-links-to-log,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// MdtConfiguration represents the ASN.1 SEQUENCE type.
type MdtConfiguration struct {
	MdtActivation MdtActivation `json:"mdt-Activation,omitempty"`
	Mdtmode Mdtmode `json:"mDTMode,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// MrdcDataUsageReportItem represents the ASN.1 SEQUENCE type.
type MrdcDataUsageReportItem struct {
	Starttimestamp MrdcDataUsageReportItemStarttimestamp `json:"startTimeStamp,omitempty"`
	Endtimestamp MrdcDataUsageReportItemEndtimestamp `json:"endTimeStamp,omitempty"`
	Usagecountul MrdcDataUsageReportItemUsagecountul `json:"usageCountUL,omitempty"`
	Usagecountdl MrdcDataUsageReportItemUsagecountdl `json:"usageCountDL,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// MrdcDatausagereport represents the ASN.1 SEQUENCE type.
type MrdcDatausagereport struct {
	Protocolies ProtocolieContainer `json:"protocolIEs,omitempty"`
}

// MrdcUsageInformation represents the ASN.1 SEQUENCE type.
type MrdcUsageInformation struct {
	DataUsagePerPduSessionReport *DataUsagePerPduSessionReport `json:"data-Usage-per-PDU-Session-Report,omitempty"`
	DataUsagePerQosFlowList *DataUsagePerQosFlowList `json:"data-Usage-per-QoS-Flow-List,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// Maximumipdatarate represents the ASN.1 SEQUENCE type.
type Maximumipdatarate struct {
	Maxiprate Maxiprate `json:"maxIPrate,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// NgRanQosSupportItem represents the ASN.1 SEQUENCE type.
type NgRanQosSupportItem struct {
	NonDynamic5qidescriptor NonDynamic5qidescriptor `json:"non-Dynamic5QIDescriptor,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// Ngranallocationandretentionpriority represents the ASN.1 SEQUENCE type.
type Ngranallocationandretentionpriority struct {
	Prioritylevel Prioritylevel `json:"priorityLevel,omitempty"`
	PreEmptioncapability PreEmptioncapability `json:"pre-emptionCapability,omitempty"`
	PreEmptionvulnerability PreEmptionvulnerability `json:"pre-emptionVulnerability,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// NpncontextinfoSnpn represents the ASN.1 SEQUENCE type.
type NpncontextinfoSnpn struct {
	Nid Nid `json:"nID,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// NpnsupportinfoSnpn represents the ASN.1 SEQUENCE type.
type NpnsupportinfoSnpn struct {
	Nid Nid `json:"nID,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// NrCgi represents the ASN.1 SEQUENCE type.
type NrCgi struct {
	PlmnIdentity PlmnIdentity `json:"pLMN-Identity,omitempty"`
	NrCellIdentity NrCellIdentity `json:"nR-Cell-Identity,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// NrCgiSupportItem represents the ASN.1 SEQUENCE type.
type NrCgiSupportItem struct {
	NrCgi NrCgi `json:"nR-CGI,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// NonDynamic5qidescriptor represents the ASN.1 SEQUENCE type.
type NonDynamic5qidescriptor struct {
	Qosprioritylevel *Qosprioritylevel `json:"qoSPriorityLevel,omitempty"`
	Averagingwindow *Averagingwindow `json:"averagingWindow,omitempty"`
	Maxdataburstvolume *Maxdataburstvolume `json:"maxDataBurstVolume,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// PdcpConfiguration represents the ASN.1 SEQUENCE type.
type PdcpConfiguration struct {
	PdcpSnSizeUl PdcpSnSize `json:"pDCP-SN-Size-UL,omitempty"`
	PdcpSnSizeDl PdcpSnSize `json:"pDCP-SN-Size-DL,omitempty"`
	RlcMode RlcMode `json:"rLC-Mode,omitempty"`
	RohcParameters *RohcParameters `json:"rOHC-Parameters,omitempty"`
	TReorderingtimer *TReorderingtimer `json:"t-ReorderingTimer,omitempty"`
	Discardtimer *Discardtimer `json:"discardTimer,omitempty"`
	Uldatasplitthreshold *Uldatasplitthreshold `json:"uLDataSplitThreshold,omitempty"`
	PdcpDuplication *PdcpDuplication `json:"pDCP-Duplication,omitempty"`
	PdcpReestablishment *PdcpReestablishment `json:"pDCP-Reestablishment,omitempty"`
	PdcpDatarecovery *PdcpDatarecovery `json:"pDCP-DataRecovery,omitempty"`
	DuplicationActivation *DuplicationActivation `json:"duplication-Activation,omitempty"`
	Outoforderdelivery *Outoforderdelivery `json:"outOfOrderDelivery,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// PdcpCount represents the ASN.1 SEQUENCE type.
type PdcpCount struct {
	PdcpSn PdcpSn `json:"pDCP-SN,omitempty"`
	Hfn Hfn `json:"hFN,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// PdcpSnStatusInformation represents the ASN.1 SEQUENCE type.
type PdcpSnStatusInformation struct {
	PdcpstatustransferUl Drbbstatustransfer `json:"pdcpStatusTransfer-UL,omitempty"`
	PdcpstatustransferDl PdcpCount `json:"pdcpStatusTransfer-DL,omitempty"`
	IeExtension *Protocolextensioncontainer `json:"iE-Extension,omitempty"`
}

// PduSessionResourceActivityItem represents the ASN.1 SEQUENCE type.
type PduSessionResourceActivityItem struct {
	PduSessionId PduSessionId `json:"pDU-Session-ID,omitempty"`
	PduSessionResourceActivity PduSessionResourceActivity `json:"pDU-Session-Resource-Activity,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// PduSessionResourceConfirmModifiedItem represents the ASN.1 SEQUENCE type.
type PduSessionResourceConfirmModifiedItem struct {
	PduSessionId PduSessionId `json:"pDU-Session-ID,omitempty"`
	DrbConfirmModifiedListNgRan *DrbConfirmModifiedListNgRan `json:"dRB-Confirm-Modified-List-NG-RAN,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// PduSessionResourceDataUsageItem represents the ASN.1 SEQUENCE type.
type PduSessionResourceDataUsageItem struct {
	PduSessionId PduSessionId `json:"pDU-Session-ID,omitempty"`
	MrdcUsageInformation MrdcUsageInformation `json:"mRDC-Usage-Information,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// PduSessionResourceFailedItem represents the ASN.1 SEQUENCE type.
type PduSessionResourceFailedItem struct {
	PduSessionId PduSessionId `json:"pDU-Session-ID,omitempty"`
	Cause Cause `json:"cause,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// PduSessionResourceFailedModItem represents the ASN.1 SEQUENCE type.
type PduSessionResourceFailedModItem struct {
	PduSessionId PduSessionId `json:"pDU-Session-ID,omitempty"`
	Cause Cause `json:"cause,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// PduSessionResourceFailedToModifyItem represents the ASN.1 SEQUENCE type.
type PduSessionResourceFailedToModifyItem struct {
	PduSessionId PduSessionId `json:"pDU-Session-ID,omitempty"`
	Cause Cause `json:"cause,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// PduSessionResourceModifiedItem represents the ASN.1 SEQUENCE type.
type PduSessionResourceModifiedItem struct {
	PduSessionId PduSessionId `json:"pDU-Session-ID,omitempty"`
	NgDlUpTnlInformation *UpTnlInformation `json:"nG-DL-UP-TNL-Information,omitempty"`
	Securityresult *Securityresult `json:"securityResult,omitempty"`
	PduSessionDataForwardingInformationResponse *DataForwardingInformation `json:"pDU-Session-Data-Forwarding-Information-Response,omitempty"`
	DrbSetupListNgRan *DrbSetupListNgRan `json:"dRB-Setup-List-NG-RAN,omitempty"`
	DrbFailedListNgRan *DrbFailedListNgRan `json:"dRB-Failed-List-NG-RAN,omitempty"`
	DrbModifiedListNgRan *DrbModifiedListNgRan `json:"dRB-Modified-List-NG-RAN,omitempty"`
	DrbFailedToModifyListNgRan *DrbFailedToModifyListNgRan `json:"dRB-Failed-To-Modify-List-NG-RAN,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// PduSessionResourceRequiredToModifyItem represents the ASN.1 SEQUENCE type.
type PduSessionResourceRequiredToModifyItem struct {
	PduSessionId PduSessionId `json:"pDU-Session-ID,omitempty"`
	NgDlUpTnlInformation *UpTnlInformation `json:"nG-DL-UP-TNL-Information,omitempty"`
	DrbRequiredToModifyListNgRan *DrbRequiredToModifyListNgRan `json:"dRB-Required-To-Modify-List-NG-RAN,omitempty"`
	DrbRequiredToRemoveListNgRan *DrbRequiredToRemoveListNgRan `json:"dRB-Required-To-Remove-List-NG-RAN,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// PduSessionResourceSetupItem represents the ASN.1 SEQUENCE type.
type PduSessionResourceSetupItem struct {
	PduSessionId PduSessionId `json:"pDU-Session-ID,omitempty"`
	Securityresult *Securityresult `json:"securityResult,omitempty"`
	NgDlUpTnlInformation UpTnlInformation `json:"nG-DL-UP-TNL-Information,omitempty"`
	PduSessionDataForwardingInformationResponse *DataForwardingInformation `json:"pDU-Session-Data-Forwarding-Information-Response,omitempty"`
	DrbSetupListNgRan DrbSetupListNgRan `json:"dRB-Setup-List-NG-RAN,omitempty"`
	DrbFailedListNgRan *DrbFailedListNgRan `json:"dRB-Failed-List-NG-RAN,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// PduSessionResourceSetupModItem represents the ASN.1 SEQUENCE type.
type PduSessionResourceSetupModItem struct {
	PduSessionId PduSessionId `json:"pDU-Session-ID,omitempty"`
	Securityresult *Securityresult `json:"securityResult,omitempty"`
	NgDlUpTnlInformation UpTnlInformation `json:"nG-DL-UP-TNL-Information,omitempty"`
	PduSessionDataForwardingInformationResponse *DataForwardingInformation `json:"pDU-Session-Data-Forwarding-Information-Response,omitempty"`
	DrbSetupModListNgRan DrbSetupModListNgRan `json:"dRB-Setup-Mod-List-NG-RAN,omitempty"`
	DrbFailedModListNgRan *DrbFailedModListNgRan `json:"dRB-Failed-Mod-List-NG-RAN,omitempty"`
	IeExtensions Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// PduSessionResourceToModifyItem represents the ASN.1 SEQUENCE type.
type PduSessionResourceToModifyItem struct {
	PduSessionId PduSessionId `json:"pDU-Session-ID,omitempty"`
	Securityindication *Securityindication `json:"securityIndication,omitempty"`
	PduSessionResourceDlAmbr *Bitrate `json:"pDU-Session-Resource-DL-AMBR,omitempty"`
	NgUlUpTnlInformation *UpTnlInformation `json:"nG-UL-UP-TNL-Information,omitempty"`
	PduSessionDataForwardingInformationRequest *DataForwardingInformationRequest `json:"pDU-Session-Data-Forwarding-Information-Request,omitempty"`
	PduSessionDataForwardingInformation *DataForwardingInformation `json:"pDU-Session-Data-Forwarding-Information,omitempty"`
	PduSessionInactivityTimer *InactivityTimer `json:"pDU-Session-Inactivity-Timer,omitempty"`
	Networkinstance *Networkinstance `json:"networkInstance,omitempty"`
	DrbToSetupListNgRan *DrbToSetupListNgRan `json:"dRB-To-Setup-List-NG-RAN,omitempty"`
	DrbToModifyListNgRan *DrbToModifyListNgRan `json:"dRB-To-Modify-List-NG-RAN,omitempty"`
	DrbToRemoveListNgRan *DrbToRemoveListNgRan `json:"dRB-To-Remove-List-NG-RAN,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// PduSessionResourceToRemoveItem represents the ASN.1 SEQUENCE type.
type PduSessionResourceToRemoveItem struct {
	PduSessionId PduSessionId `json:"pDU-Session-ID,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// PduSessionResourceToSetupItem represents the ASN.1 SEQUENCE type.
type PduSessionResourceToSetupItem struct {
	PduSessionId PduSessionId `json:"pDU-Session-ID,omitempty"`
	PduSessionType PduSessionType `json:"pDU-Session-Type,omitempty"`
	Snssai Snssai `json:"sNSSAI,omitempty"`
	Securityindication Securityindication `json:"securityIndication,omitempty"`
	PduSessionResourceDlAmbr *Bitrate `json:"pDU-Session-Resource-DL-AMBR,omitempty"`
	NgUlUpTnlInformation UpTnlInformation `json:"nG-UL-UP-TNL-Information,omitempty"`
	PduSessionDataForwardingInformationRequest *DataForwardingInformationRequest `json:"pDU-Session-Data-Forwarding-Information-Request,omitempty"`
	PduSessionInactivityTimer *InactivityTimer `json:"pDU-Session-Inactivity-Timer,omitempty"`
	ExistingAllocatedNgDlUpTnlInfo *UpTnlInformation `json:"existing-Allocated-NG-DL-UP-TNL-Info,omitempty"`
	Networkinstance *Networkinstance `json:"networkInstance,omitempty"`
	DrbToSetupListNgRan DrbToSetupListNgRan `json:"dRB-To-Setup-List-NG-RAN,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// PduSessionResourceToSetupModItem represents the ASN.1 SEQUENCE type.
type PduSessionResourceToSetupModItem struct {
	PduSessionId PduSessionId `json:"pDU-Session-ID,omitempty"`
	PduSessionType PduSessionType `json:"pDU-Session-Type,omitempty"`
	Snssai Snssai `json:"sNSSAI,omitempty"`
	Securityindication Securityindication `json:"securityIndication,omitempty"`
	PduSessionResourceAmbr *Bitrate `json:"pDU-Session-Resource-AMBR,omitempty"`
	NgUlUpTnlInformation UpTnlInformation `json:"nG-UL-UP-TNL-Information,omitempty"`
	PduSessionDataForwardingInformationRequest *DataForwardingInformationRequest `json:"pDU-Session-Data-Forwarding-Information-Request,omitempty"`
	PduSessionInactivityTimer *InactivityTimer `json:"pDU-Session-Inactivity-Timer,omitempty"`
	DrbToSetupModListNgRan DrbToSetupModListNgRan `json:"dRB-To-Setup-Mod-List-NG-RAN,omitempty"`
	IeExtensions Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// PduSessionToNotifyItem represents the ASN.1 SEQUENCE type.
type PduSessionToNotifyItem struct {
	PduSessionId PduSessionId `json:"pDU-Session-ID,omitempty"`
	QosFlowList QosFlowList `json:"qoS-Flow-List,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// Packeterrorrate represents the ASN.1 SEQUENCE type.
type Packeterrorrate struct {
	PerScalar PerScalar `json:"pER-Scalar,omitempty"`
	PerExponent PerExponent `json:"pER-Exponent,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// PrivateieFieldE1apPrivateIesIessetparam represents the ASN.1 SEQUENCE type.
type PrivateieFieldE1apPrivateIesIessetparam struct {
	Id E1apPrivateIesId `json:"id,omitempty"`
	Criticality E1apPrivateIesCriticality `json:"criticality,omitempty"`
	Value E1apPrivateIesValue `json:"value,omitempty"`
}

// Privatemessage represents the ASN.1 SEQUENCE type.
type Privatemessage struct {
	Privateies PrivateieContainer `json:"privateIEs,omitempty"`
}

// ProtocolextensionfieldE1apProtocolExtensionExtensionsetparam represents the ASN.1 SEQUENCE type.
type ProtocolextensionfieldE1apProtocolExtensionExtensionsetparam struct {
	Id E1apProtocolExtensionId `json:"id,omitempty"`
	Criticality E1apProtocolExtensionCriticality `json:"criticality,omitempty"`
	Extensionvalue E1apProtocolExtensionExtension `json:"extensionValue,omitempty"`
}

// ProtocolieFieldE1apProtocolIesIessetparam represents the ASN.1 SEQUENCE type.
type ProtocolieFieldE1apProtocolIesIessetparam struct {
	Id E1apProtocolIesId `json:"id,omitempty"`
	Criticality E1apProtocolIesCriticality `json:"criticality,omitempty"`
	Value E1apProtocolIesValue `json:"value,omitempty"`
}

// QosFlowFailedItem represents the ASN.1 SEQUENCE type.
type QosFlowFailedItem struct {
	QosFlowIdentifier QosFlowIdentifier `json:"qoS-Flow-Identifier,omitempty"`
	Cause Cause `json:"cause,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// QosFlowItem represents the ASN.1 SEQUENCE type.
type QosFlowItem struct {
	QosFlowIdentifier QosFlowIdentifier `json:"qoS-Flow-Identifier,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// QosFlowMappingItem represents the ASN.1 SEQUENCE type.
type QosFlowMappingItem struct {
	QosFlowIdentifier QosFlowIdentifier `json:"qoS-Flow-Identifier,omitempty"`
	Qosflowmappingindication *QosFlowMappingIndication `json:"qoSFlowMappingIndication,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// QosFlowQosParameterItem represents the ASN.1 SEQUENCE type.
type QosFlowQosParameterItem struct {
	QosFlowIdentifier QosFlowIdentifier `json:"qoS-Flow-Identifier,omitempty"`
	Qosflowlevelqosparameters Qosflowlevelqosparameters `json:"qoSFlowLevelQoSParameters,omitempty"`
	Qosflowmappingindication *QosFlowMappingIndication `json:"qoSFlowMappingIndication,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// QosFlowRemovedItem represents the ASN.1 SEQUENCE type.
type QosFlowRemovedItem struct {
	QosFlowIdentifier QosFlowIdentifier `json:"qoS-Flow-Identifier,omitempty"`
	QosFlowAccumulatedSessionTime *QosFlowRemovedItemQosFlowAccumulatedSessionTime `json:"qoS-Flow-Accumulated-Session-Time,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// QosMappingInformation represents the ASN.1 SEQUENCE type.
type QosMappingInformation struct {
	Dscp *QosMappingInformationDscp `json:"dscp,omitempty"`
	FlowLabel *QosMappingInformationFlowLabel `json:"flow-label,omitempty"`
}

// QosParametersSupportList represents the ASN.1 SEQUENCE type.
type QosParametersSupportList struct {
	EutranQosSupportList *EutranQosSupportList `json:"eUTRAN-QoS-Support-List,omitempty"`
	NgRanQosSupportList *NgRanQosSupportList `json:"nG-RAN-QoS-Support-List,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// Qosflowlevelqosparameters represents the ASN.1 SEQUENCE type.
type Qosflowlevelqosparameters struct {
	QosCharacteristics QosCharacteristics `json:"qoS-Characteristics,omitempty"`
	Ngranallocationretentionpriority Ngranallocationandretentionpriority `json:"nGRANallocationRetentionPriority,omitempty"`
	GbrQosFlowInformation *GbrQosflowinformation `json:"gBR-QoS-Flow-Information,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// Rohc represents the ASN.1 SEQUENCE type.
type Rohc struct {
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// Redundantpdusessioninformation represents the ASN.1 SEQUENCE type.
type Redundantpdusessioninformation struct {
	Rsn Rsn `json:"rSN,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// Reset represents the ASN.1 SEQUENCE type.
type Reset struct {
	Protocolies ProtocolieContainer `json:"protocolIEs,omitempty"`
}

// Resetacknowledge represents the ASN.1 SEQUENCE type.
type Resetacknowledge struct {
	Protocolies ProtocolieContainer `json:"protocolIEs,omitempty"`
}

// Resourcestatusfailure represents the ASN.1 SEQUENCE type.
type Resourcestatusfailure struct {
	Protocolies ProtocolieContainer `json:"protocolIEs,omitempty"`
}

// Resourcestatusrequest represents the ASN.1 SEQUENCE type.
type Resourcestatusrequest struct {
	Protocolies ProtocolieContainer `json:"protocolIEs,omitempty"`
}

// Resourcestatusresponse represents the ASN.1 SEQUENCE type.
type Resourcestatusresponse struct {
	Protocolies ProtocolieContainer `json:"protocolIEs,omitempty"`
}

// Resourcestatusupdate represents the ASN.1 SEQUENCE type.
type Resourcestatusupdate struct {
	Protocolies ProtocolieContainer `json:"protocolIEs,omitempty"`
}

// SdapConfiguration represents the ASN.1 SEQUENCE type.
type SdapConfiguration struct {
	Defaultdrb Defaultdrb `json:"defaultDRB,omitempty"`
	SdapHeaderUl SdapHeaderUl `json:"sDAP-Header-UL,omitempty"`
	SdapHeaderDl SdapHeaderDl `json:"sDAP-Header-DL,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// Snssai represents the ASN.1 SEQUENCE type.
type Snssai struct {
	Sst SnssaiSst `json:"sST,omitempty"`
	Sd *SnssaiSd `json:"sD,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// Securityalgorithm represents the ASN.1 SEQUENCE type.
type Securityalgorithm struct {
	Cipheringalgorithm Cipheringalgorithm `json:"cipheringAlgorithm,omitempty"`
	Integrityprotectionalgorithm *Integrityprotectionalgorithm `json:"integrityProtectionAlgorithm,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// Securityindication represents the ASN.1 SEQUENCE type.
type Securityindication struct {
	Integrityprotectionindication Integrityprotectionindication `json:"integrityProtectionIndication,omitempty"`
	Confidentialityprotectionindication Confidentialityprotectionindication `json:"confidentialityProtectionIndication,omitempty"`
	Maximumipdatarate *Maximumipdatarate `json:"maximumIPdatarate,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// Securityinformation represents the ASN.1 SEQUENCE type.
type Securityinformation struct {
	Securityalgorithm Securityalgorithm `json:"securityAlgorithm,omitempty"`
	Upsecuritykey Upsecuritykey `json:"uPSecuritykey,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// Securityresult represents the ASN.1 SEQUENCE type.
type Securityresult struct {
	Integrityprotectionresult Integrityprotectionresult `json:"integrityProtectionResult,omitempty"`
	Confidentialityprotectionresult Confidentialityprotectionresult `json:"confidentialityProtectionResult,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// SliceSupportItem represents the ASN.1 SEQUENCE type.
type SliceSupportItem struct {
	Snssai Snssai `json:"sNSSAI,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// Successfuloutcome represents the ASN.1 SEQUENCE type.
type Successfuloutcome struct {
	Procedurecode E1apElementaryProcedureProcedurecode `json:"procedureCode,omitempty"`
	Criticality E1apElementaryProcedureCriticality `json:"criticality,omitempty"`
	Value E1apElementaryProcedureSuccessfuloutcome `json:"value,omitempty"`
}

// SupportedplmnsItem represents the ASN.1 SEQUENCE type.
type SupportedplmnsItem struct {
	PlmnIdentity PlmnIdentity `json:"pLMN-Identity,omitempty"`
	SliceSupportList *SliceSupportList `json:"slice-Support-List,omitempty"`
	NrCgiSupportList *NrCgiSupportList `json:"nR-CGI-Support-List,omitempty"`
	QosParametersSupportList *QosParametersSupportList `json:"qoS-Parameters-Support-List,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// TReorderingtimer represents the ASN.1 SEQUENCE type.
type TReorderingtimer struct {
	TReordering TReordering `json:"t-Reordering,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// TnlAvailablecapacityindicator represents the ASN.1 SEQUENCE type.
type TnlAvailablecapacityindicator struct {
	IeExtensions Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// Tsctrafficcharacteristics represents the ASN.1 SEQUENCE type.
type Tsctrafficcharacteristics struct {
	Tsctrafficcharacteristicsul *Tsctrafficinformation `json:"tSCTrafficCharacteristicsUL,omitempty"`
	Tsctrafficcharacteristicsdl *Tsctrafficinformation `json:"tSCTrafficCharacteristicsDL,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// Tsctrafficinformation represents the ASN.1 SEQUENCE type.
type Tsctrafficinformation struct {
	Periodicity Periodicity `json:"periodicity,omitempty"`
	Burstarrivaltime *Burstarrivaltime `json:"burstArrivalTime,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// Traceactivation represents the ASN.1 SEQUENCE type.
type Traceactivation struct {
	Traceid Traceid `json:"traceID,omitempty"`
	Interfacestotrace Interfacestotrace `json:"interfacesToTrace,omitempty"`
	Tracedepth Tracedepth `json:"traceDepth,omitempty"`
	Tracecollectionentityipaddress Transportlayeraddress `json:"traceCollectionEntityIPAddress,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// Tracestart represents the ASN.1 SEQUENCE type.
type Tracestart struct {
	Protocolies ProtocolieContainer `json:"protocolIEs,omitempty"`
}

// TransportLayerAddressInfo represents the ASN.1 SEQUENCE type.
type TransportLayerAddressInfo struct {
	TransportUpLayerAddressesInfoToAddList *TransportUpLayerAddressesInfoToAddList `json:"transport-UP-Layer-Addresses-Info-To-Add-List,omitempty"`
	TransportUpLayerAddressesInfoToRemoveList *TransportUpLayerAddressesInfoToRemoveList `json:"transport-UP-Layer-Addresses-Info-To-Remove-List,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// TransportUpLayerAddressesInfoToAddItem represents the ASN.1 SEQUENCE type.
type TransportUpLayerAddressesInfoToAddItem struct {
	IpSectransportlayeraddress Transportlayeraddress `json:"iP-SecTransportLayerAddress,omitempty"`
	Gtptransportlayeraddressestoadd *Gtptlas `json:"gTPTransportLayerAddressesToAdd,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// TransportUpLayerAddressesInfoToRemoveItem represents the ASN.1 SEQUENCE type.
type TransportUpLayerAddressesInfoToRemoveItem struct {
	IpSectransportlayeraddress Transportlayeraddress `json:"iP-SecTransportLayerAddress,omitempty"`
	Gtptransportlayeraddressestoremove *Gtptlas `json:"gTPTransportLayerAddressesToRemove,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// UeAssociatedlogicale1Connectionitem represents the ASN.1 SEQUENCE type.
type UeAssociatedlogicale1Connectionitem struct {
	GnbCuCpUeE1apId *GnbCuCpUeE1apId `json:"gNB-CU-CP-UE-E1AP-ID,omitempty"`
	GnbCuUpUeE1apId *GnbCuUpUeE1apId `json:"gNB-CU-UP-UE-E1AP-ID,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// Uldatanotification represents the ASN.1 SEQUENCE type.
type Uldatanotification struct {
	Protocolies ProtocolieContainer `json:"protocolIEs,omitempty"`
}

// Uluptnladdresstoupdateitem represents the ASN.1 SEQUENCE type.
type Uluptnladdresstoupdateitem struct {
	Oldtnladress Transportlayeraddress `json:"oldTNLAdress,omitempty"`
	Newtnladress Transportlayeraddress `json:"newTNLAdress,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// UpParametersItem represents the ASN.1 SEQUENCE type.
type UpParametersItem struct {
	UpTnlInformation UpTnlInformation `json:"uP-TNL-Information,omitempty"`
	CellGroupId CellGroupId `json:"cell-Group-ID,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// Upsecuritykey represents the ASN.1 SEQUENCE type.
type Upsecuritykey struct {
	Encryptionkey Encryptionkey `json:"encryptionKey,omitempty"`
	Integrityprotectionkey *Integrityprotectionkey `json:"integrityProtectionKey,omitempty"`
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

// Unsuccessfuloutcome represents the ASN.1 SEQUENCE type.
type Unsuccessfuloutcome struct {
	Procedurecode E1apElementaryProcedureProcedurecode `json:"procedureCode,omitempty"`
	Criticality E1apElementaryProcedureCriticality `json:"criticality,omitempty"`
	Value E1apElementaryProcedureUnsuccessfuloutcome `json:"value,omitempty"`
}

// Uplinkonlyrohc represents the ASN.1 SEQUENCE type.
type Uplinkonlyrohc struct {
	IeExtensions *Protocolextensioncontainer `json:"iE-Extensions,omitempty"`
}

