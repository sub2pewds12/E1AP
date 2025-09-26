package e1ap_ies

// Choinitiation represents the ASN.1 ENUMERATED type.
type Choinitiation int32

// Choinitiation Values
const (
	Choinitiation_True Choinitiation = 0
)

// Criticality represents the ASN.1 ENUMERATED type.
type Criticality int32

// Criticality Values
const (
	Criticality_Reject Criticality = 0
	Criticality_Ignore Criticality = 1
	Criticality_Notify Criticality = 2
)

// Discardtimer represents the ASN.1 ENUMERATED type.
type Discardtimer int32

// Discardtimer Values
const (
	Discardtimer_Ms10 Discardtimer = 0
	Discardtimer_Ms20 Discardtimer = 1
	Discardtimer_Ms30 Discardtimer = 2
	Discardtimer_Ms40 Discardtimer = 3
	Discardtimer_Ms50 Discardtimer = 4
	Discardtimer_Ms60 Discardtimer = 5
	Discardtimer_Ms75 Discardtimer = 6
	Discardtimer_Ms100 Discardtimer = 7
	Discardtimer_Ms150 Discardtimer = 8
	Discardtimer_Ms200 Discardtimer = 9
	Discardtimer_Ms250 Discardtimer = 10
	Discardtimer_Ms300 Discardtimer = 11
	Discardtimer_Ms500 Discardtimer = 12
	Discardtimer_Ms750 Discardtimer = 13
	Discardtimer_Ms1500 Discardtimer = 14
	Discardtimer_Infinity Discardtimer = 15
)

// Dynamic5qidescriptorDelaycritical represents the ASN.1 ENUMERATED type.
type Dynamic5qidescriptorDelaycritical int32


// Earlyforwardingcountreq represents the ASN.1 ENUMERATED type.
type Earlyforwardingcountreq int32

// Earlyforwardingcountreq Values
const (
	Earlyforwardingcountreq_FirstDlCount Earlyforwardingcountreq = 0
	Earlyforwardingcountreq_DlDiscarding Earlyforwardingcountreq = 1
)

// GnbCuUpOverloadinformation represents the ASN.1 ENUMERATED type.
type GnbCuUpOverloadinformation int32

// GnbCuUpOverloadinformation Values
const (
	GnbCuUpOverloadinformation_Overloaded GnbCuUpOverloadinformation = 0
	GnbCuUpOverloadinformation_NotOverloaded GnbCuUpOverloadinformation = 1
)

// M4period represents the ASN.1 ENUMERATED type.
type M4period int32

// M4period Values
const (
	M4period_Ms1024 M4period = 0
	M4period_Ms2048 M4period = 1
	M4period_Ms5120 M4period = 2
	M4period_Ms10240 M4period = 3
	M4period_Min1 M4period = 4
)

// M6reportInterval represents the ASN.1 ENUMERATED type.
type M6reportInterval int32

// M6reportInterval Values
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

// Presence represents the ASN.1 ENUMERATED type.
type Presence int32

// Presence Values
const (
	Presence_Optional Presence = 0
	Presence_Conditional Presence = 1
	Presence_Mandatory Presence = 2
)

// QosFlowMappingIndication represents the ASN.1 ENUMERATED type.
type QosFlowMappingIndication int32

// QosFlowMappingIndication Values
const (
	QosFlowMappingIndication_Ul QosFlowMappingIndication = 0
	QosFlowMappingIndication_Dl QosFlowMappingIndication = 1
)

// Qosmonitoringrequest represents the ASN.1 ENUMERATED type.
type Qosmonitoringrequest int32

// Qosmonitoringrequest Values
const (
	Qosmonitoringrequest_Ul Qosmonitoringrequest = 0
	Qosmonitoringrequest_Dl Qosmonitoringrequest = 1
	Qosmonitoringrequest_Both Qosmonitoringrequest = 2
)

// Rsn represents the ASN.1 ENUMERATED type.
type Rsn int32

// Rsn Values
const (
	Rsn_V1 Rsn = 0
	Rsn_V2 Rsn = 1
)

// Redundantqosflowindicator represents the ASN.1 ENUMERATED type.
type Redundantqosflowindicator int32

// Redundantqosflowindicator Values
const (
	Redundantqosflowindicator_True Redundantqosflowindicator = 0
	Redundantqosflowindicator_False Redundantqosflowindicator = 1
)

// TReordering represents the ASN.1 ENUMERATED type.
type TReordering int32


// Timetowait represents the ASN.1 ENUMERATED type.
type Timetowait int32

// Timetowait Values
const (
	Timetowait_V1s Timetowait = 0
	Timetowait_V2s Timetowait = 1
	Timetowait_V5s Timetowait = 2
	Timetowait_V10s Timetowait = 3
	Timetowait_V20s Timetowait = 4
	Timetowait_V60s Timetowait = 5
)

// Triggeringmessage represents the ASN.1 ENUMERATED type.
type Triggeringmessage int32

// Triggeringmessage Values
const (
	Triggeringmessage_InitiatingMessage Triggeringmessage = 0
	Triggeringmessage_SuccessfulOutcome Triggeringmessage = 1
	Triggeringmessage_UnsuccessfulOutcome Triggeringmessage = 2
)

// Uldatasplitthreshold represents the ASN.1 ENUMERATED type.
type Uldatasplitthreshold int32


