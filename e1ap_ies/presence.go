package e1ap_ies

// Presence represents the ASN.1 definition from 9_4_6_Common_Definitions.txt:36
type Presence int32

const (
	Presence_Optional    Presence = 0
	Presence_Conditional Presence = 1
	Presence_Mandatory   Presence = 2
)
