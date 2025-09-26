package e1ap_ies

// Activityinformation represents the ASN.1 CHOICE type.
type Activityinformation struct {
	DrbActivityList DrbActivityList `json:"dRB-Activity-List,omitempty"`
	PduSessionResourceActivityList PduSessionResourceActivityList `json:"pDU-Session-Resource-Activity-List,omitempty"`
	UeActivity UeActivity `json:"uE-Activity,omitempty"`
	ChoiceExtension ProtocolieSinglecontainer `json:"choice-extension,omitempty"`
}

// Cause represents the ASN.1 CHOICE type.
type Cause struct {
	Radionetwork Causeradionetwork `json:"radioNetwork,omitempty"`
	Transport Causetransport `json:"transport,omitempty"`
	Protocol Causeprotocol `json:"protocol,omitempty"`
	Misc Causemisc `json:"misc,omitempty"`
	ChoiceExtension ProtocolieSinglecontainer `json:"choice-extension,omitempty"`
}

// E1apPdu represents the ASN.1 CHOICE type.
type E1apPdu struct {
	Initiatingmessage Initiatingmessage `json:"initiatingMessage,omitempty"`
	Successfuloutcome Successfuloutcome `json:"successfulOutcome,omitempty"`
	Unsuccessfuloutcome Unsuccessfuloutcome `json:"unsuccessfulOutcome,omitempty"`
}

// Earlyforwardingcountinfo represents the ASN.1 CHOICE type.
type Earlyforwardingcountinfo struct {
	Firstdlcount Firstdlcount `json:"firstDLCount,omitempty"`
	Dldiscardingcount Dldiscarding `json:"dLDiscardingCount,omitempty"`
	ChoiceExtension ProtocolieSinglecontainer `json:"choice-Extension,omitempty"`
}

// Mdtmode represents the ASN.1 CHOICE type.
type Mdtmode struct {
	Immediatemdt Immediatemdt `json:"immediateMDT,omitempty"`
	ChoiceExtension ProtocolieSinglecontainer `json:"choice-extension,omitempty"`
}

// Npncontextinfo represents the ASN.1 CHOICE type.
type Npncontextinfo struct {
	Snpn NpncontextinfoSnpn `json:"sNPN,omitempty"`
	ChoiceExtension ProtocolieSinglecontainer `json:"choice-extension,omitempty"`
}

// Npnsupportinfo represents the ASN.1 CHOICE type.
type Npnsupportinfo struct {
	Snpn NpnsupportinfoSnpn `json:"sNPN,omitempty"`
	ChoiceExtension ProtocolieSinglecontainer `json:"choice-extension,omitempty"`
}

// PrivateieId represents the ASN.1 CHOICE type.
type PrivateieId struct {
	Local PrivateieIdLocal `json:"local,omitempty"`
	Global Object `json:"global,omitempty"`
}

// QosCharacteristics represents the ASN.1 CHOICE type.
type QosCharacteristics struct {
	NonDynamic5qi NonDynamic5qidescriptor `json:"non-Dynamic-5QI,omitempty"`
	Dynamic5qi Dynamic5qidescriptor `json:"dynamic-5QI,omitempty"`
	ChoiceExtension ProtocolieSinglecontainer `json:"choice-extension,omitempty"`
}

// RohcParameters represents the ASN.1 CHOICE type.
type RohcParameters struct {
	Rohc Rohc `json:"rOHC,omitempty"`
	Uplinkonlyrohc Uplinkonlyrohc `json:"uPlinkOnlyROHC,omitempty"`
	ChoiceExtension ProtocolieSinglecontainer `json:"choice-Extension,omitempty"`
}

// Resettype represents the ASN.1 CHOICE type.
type Resettype struct {
	E1Interface Resetall `json:"e1-Interface,omitempty"`
	Partofe1Interface UeAssociatedlogicale1Connectionlistres `json:"partOfE1-Interface,omitempty"`
	ChoiceExtension ProtocolieSinglecontainer `json:"choice-extension,omitempty"`
}

// SystemBearercontextmodificationconfirm represents the ASN.1 CHOICE type.
type SystemBearercontextmodificationconfirm struct {
	EUtranBearercontextmodificationconfirm ProtocolieContainer `json:"e-UTRAN-BearerContextModificationConfirm,omitempty"`
	NgRanBearercontextmodificationconfirm ProtocolieContainer `json:"nG-RAN-BearerContextModificationConfirm,omitempty"`
	ChoiceExtension ProtocolieSinglecontainer `json:"choice-extension,omitempty"`
}

// SystemBearercontextmodificationrequest represents the ASN.1 CHOICE type.
type SystemBearercontextmodificationrequest struct {
	EUtranBearercontextmodificationrequest ProtocolieContainer `json:"e-UTRAN-BearerContextModificationRequest,omitempty"`
	NgRanBearercontextmodificationrequest ProtocolieContainer `json:"nG-RAN-BearerContextModificationRequest,omitempty"`
	ChoiceExtension ProtocolieSinglecontainer `json:"choice-extension,omitempty"`
}

// SystemBearercontextmodificationrequired represents the ASN.1 CHOICE type.
type SystemBearercontextmodificationrequired struct {
	EUtranBearercontextmodificationrequired ProtocolieContainer `json:"e-UTRAN-BearerContextModificationRequired,omitempty"`
	NgRanBearercontextmodificationrequired ProtocolieContainer `json:"nG-RAN-BearerContextModificationRequired,omitempty"`
	ChoiceExtension ProtocolieSinglecontainer `json:"choice-extension,omitempty"`
}

// SystemBearercontextmodificationresponse represents the ASN.1 CHOICE type.
type SystemBearercontextmodificationresponse struct {
	EUtranBearercontextmodificationresponse ProtocolieContainer `json:"e-UTRAN-BearerContextModificationResponse,omitempty"`
	NgRanBearercontextmodificationresponse ProtocolieContainer `json:"nG-RAN-BearerContextModificationResponse,omitempty"`
	ChoiceExtension ProtocolieSinglecontainer `json:"choice-extension,omitempty"`
}

// SystemBearercontextsetuprequest represents the ASN.1 CHOICE type.
type SystemBearercontextsetuprequest struct {
	EUtranBearercontextsetuprequest ProtocolieContainer `json:"e-UTRAN-BearerContextSetupRequest,omitempty"`
	NgRanBearercontextsetuprequest ProtocolieContainer `json:"nG-RAN-BearerContextSetupRequest,omitempty"`
	ChoiceExtension ProtocolieSinglecontainer `json:"choice-extension,omitempty"`
}

// SystemGnbCuUpCountercheckrequest represents the ASN.1 CHOICE type.
type SystemGnbCuUpCountercheckrequest struct {
	EUtranGnbCuUpCountercheckrequest ProtocolieContainer `json:"e-UTRAN-GNB-CU-UP-CounterCheckRequest,omitempty"`
	NgRanGnbCuUpCountercheckrequest ProtocolieContainer `json:"nG-RAN-GNB-CU-UP-CounterCheckRequest,omitempty"`
	ChoiceExtension ProtocolieSinglecontainer `json:"choice-extension,omitempty"`
}

