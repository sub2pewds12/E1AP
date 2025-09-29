package e1ap_ies

// TimeToWait represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:2224
type TimeToWait int32

const (
	TimeToWait_V1s  TimeToWait = 0
	TimeToWait_V2s  TimeToWait = 1
	TimeToWait_V5s  TimeToWait = 2
	TimeToWait_V10s TimeToWait = 3
	TimeToWait_V20s TimeToWait = 4
	TimeToWait_V60s TimeToWait = 5
)
