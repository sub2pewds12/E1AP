package e1ap_ies

// TSCTrafficCharacteristics From: 9_4_5_Information_Element_Definitions.txt:2246
// ASN.1 Data Type: SEQUENCE
type TSCTrafficCharacteristics struct {
	TSCTrafficCharacteristicsUL *TSCTrafficInformation `aper:"optional"`
	TSCTrafficCharacteristicsDL *TSCTrafficInformation `aper:"optional"`
}
