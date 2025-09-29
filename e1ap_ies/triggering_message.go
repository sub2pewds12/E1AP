package e1ap_ies

// TriggeringMessage represents the ASN.1 definition from 9_4_6_Common_Definitions.txt:49
type TriggeringMessage int32

const (
	TriggeringMessage_InitiatingMessage   TriggeringMessage = 0
	TriggeringMessage_SuccessfulOutcome   TriggeringMessage = 1
	TriggeringMessage_UnsuccessfulOutcome TriggeringMessage = 2
)
