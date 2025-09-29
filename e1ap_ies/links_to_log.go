package e1ap_ies

// LinksToLog represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1286
type LinksToLog int32

const (
	LinksToLog_Uplink                LinksToLog = 0
	LinksToLog_Downlink              LinksToLog = 1
	LinksToLog_BothUplinkAndDownlink LinksToLog = 2
)
