package e1ap_ies

import (
	"bytes"
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// GNBCUCPMeasurementResultsInformation is a generated SEQUENCE type.
type GNBCUCPMeasurementResultsInformation struct {
	GNBCUCPUEE1APID                      GNBCUCPUEE1APID                      `aper:"lb:0,ub:4294967295,mandatory,ext"`
	GNBCUUPUEE1APID                      GNBCUUPUEE1APID                      `aper:"lb:0,ub:4294967295,mandatory,ext"`
	DRBMeasurementResultsInformationList DRBMeasurementResultsInformationList `aper:"mandatory,ext"`
}

// toIes transforms the GNBCUCPMeasurementResultsInformation struct into a slice of E1APMessageIEs.
func (msg *GNBCUCPMeasurementResultsInformation) toIes() ([]E1APMessageIE, error) {
	ies := make([]E1APMessageIE, 0)
	{

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEID{Value: ProtocolIEIDGNBCUCPUEE1APID},
				Criticality: Criticality{Value: CriticalityReject},
				Value: &INTEGER{
					c:     aper.Constraint{Lb: 0, Ub: 4294967295},
					ext:   false,
					Value: msg.GNBCUCPUEE1APID.Value,
				},
			})
		}
	}
	{

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEID{Value: ProtocolIEIDGNBCUUPUEE1APID},
				Criticality: Criticality{Value: CriticalityReject},
				Value: &INTEGER{
					c:     aper.Constraint{Lb: 0, Ub: 4294967295},
					ext:   false,
					Value: msg.GNBCUUPUEE1APID.Value,
				},
			})
		}
	}
	{

		tmp_DRBMeasurementResultsInformationList := Sequence[aper.IE]{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}

		for i := 0; i < len(msg.DRBMeasurementResultsInformationList.Value); i++ {
			tmp_DRBMeasurementResultsInformationList.Value = append(tmp_DRBMeasurementResultsInformationList.Value, &msg.DRBMeasurementResultsInformationList.Value[i])
		}

		{

			tmp_DRBMeasurementResultsInformationList := Sequence[aper.IE]{
				c:   aper.Constraint{Lb: 0, Ub: 0},
				ext: false,
			}

			for i := 0; i < len(msg.DRBMeasurementResultsInformationList.Value); i++ {
				tmp_DRBMeasurementResultsInformationList.Value = append(tmp_DRBMeasurementResultsInformationList.Value, &msg.DRBMeasurementResultsInformationList.Value[i])
			}

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEID{Value: ProtocolIEIDDRBMeasurementResultsInformationList},
				Criticality: Criticality{Value: CriticalityIgnore},
				Value:       &tmp_DRBMeasurementResultsInformationList,
			})
		}
	}
	var err error
	return ies, err
}

// Encode for GNBCUCPMeasurementResultsInformation: Could not find associated procedure.

// Decode implements the aper.AperUnmarshaller interface for GNBCUCPMeasurementResultsInformation.
func (msg *GNBCUCPMeasurementResultsInformation) Decode(buf []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("GNBCUCPMeasurementResultsInformation: %w", err)
		}
	}()

	r := aper.NewReader(bytes.NewReader(buf))

	decoder := GNBCUCPMeasurementResultsInformationDecoder{
		msg:  msg,
		list: make(map[ProtocolIEID]*E1APMessageIE),
	}

	// aper.ReadSequenceOf will decode the IEs and call the callback for each one.
	if _, err = aper.ReadSequenceOf[E1APMessageIE](decoder.decodeIE, r, &aper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
		return
	}

	// After decoding all present IEs, validate that mandatory ones were found.

	if _, ok := decoder.list[ProtocolIEIDGNBCUCPUEE1APID]; !ok {
		err = fmt.Errorf("mandatory field GNBCUCPUEE1APID is missing")
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject},
			IEID:          ProtocolIEIDGNBCUCPUEE1APID,
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
	}

	if _, ok := decoder.list[ProtocolIEIDGNBCUUPUEE1APID]; !ok {
		err = fmt.Errorf("mandatory field GNBCUUPUEE1APID is missing")
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject},
			IEID:          ProtocolIEIDGNBCUUPUEE1APID,
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
	}

	if _, ok := decoder.list[ProtocolIEIDDRBMeasurementResultsInformationList]; !ok {
		err = fmt.Errorf("mandatory field DRBMeasurementResultsInformationList is missing")
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject},
			IEID:          ProtocolIEIDDRBMeasurementResultsInformationList,
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
	}
	if err != nil {
		return
	}

	return
}

type GNBCUCPMeasurementResultsInformationDecoder struct {
	msg      *GNBCUCPMeasurementResultsInformation
	diagList []CriticalityDiagnosticsIEItem
	list     map[ProtocolIEID]*E1APMessageIE
}

func (decoder *GNBCUCPMeasurementResultsInformationDecoder) decodeIE(r *aper.AperReader) (msgIe *E1APMessageIE, err error) {
	var id int64
	var c uint64
	var buf []byte
	if id, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
		return nil, err
	}
	msgIe = new(E1APMessageIE)
	msgIe.Id = ProtocolIEID{Value: aper.Integer(id)}
	if c, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return nil, err
	}
	msgIe.Criticality = Criticality{Value: aper.Enumerated(c)}

	if buf, err = r.ReadOpenType(); err != nil {
		return nil, err
	}

	ieId := msgIe.Id
	if _, ok := decoder.list[ieId]; ok {
		return nil, fmt.Errorf("duplicated protocol IE ID %%d", ieId)
	}
	decoder.list[ieId] = msgIe

	ieR := aper.NewReader(bytes.NewReader(buf))
	msg := decoder.msg

	switch msgIe.Id {
	case ProtocolIEIDGNBCUCPUEE1APID:

		{
			var val int64
			if val, err = ieR.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4294967295}, false); err != nil {
				return nil, fmt.Errorf("Decode GNBCUCPUEE1APID failed: %w", err)
			}
			msg.GNBCUCPUEE1APID.Value = aper.Integer(val)
		}
	case ProtocolIEIDGNBCUUPUEE1APID:

		{
			var val int64
			if val, err = ieR.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4294967295}, false); err != nil {
				return nil, fmt.Errorf("Decode GNBCUUPUEE1APID failed: %w", err)
			}
			msg.GNBCUUPUEE1APID.Value = aper.Integer(val)
		}
	case ProtocolIEIDDRBMeasurementResultsInformationList:

		{
			itemDecoder := func(r *aper.AperReader) (*DRBMeasurementResultsInformationItem, error) {

				item := new(DRBMeasurementResultsInformationItem)
				if err := item.Decode(r); err != nil {
					return nil, err
				}
				return item, nil
			}
			var decodedItems []DRBMeasurementResultsInformationItem
			if decodedItems, err = aper.ReadSequenceOf(itemDecoder, ieR, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
				return nil, fmt.Errorf("Decode DRBMeasurementResultsInformationList failed: %w", err)
			}
			msg.DRBMeasurementResultsInformationList.Value = decodedItems
		}
	default:
		// Handle unknown IEs based on criticality here, if needed.
		// For now, we'll just ignore them.

	}
	return msgIe, nil // Return the populated msgIe and a nil error
}
