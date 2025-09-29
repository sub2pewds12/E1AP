package e1ap_ies

// TSCTrafficCharacteristics represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:2246
type TSCTrafficCharacteristics struct {
	TSCTrafficCharacteristicsUL *TSCTrafficInformation `asn1:"optional"`
	TSCTrafficCharacteristicsDL *TSCTrafficInformation `asn1:"optional"`
}
