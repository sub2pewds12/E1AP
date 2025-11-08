package e1ap

import (
	"github.com/sub2pewds12/E1AP/e1ap_ies"
)

func createMessage(present uint8, procedureCode e1ap_ies.ProcedureCode) MessageUnmarshaller {
	switch present {
	case PduChoiceInitiatingMessage:
		switch procedureCode.Value {
		case e1ap_ies.ProcedureCodeBearerContextInactivityNotification:
			return new(e1ap_ies.BearerContextInactivityNotification)
		case e1ap_ies.ProcedureCodeBearerContextModification:
			return new(e1ap_ies.BearerContextModificationRequest)
		case e1ap_ies.ProcedureCodeBearerContextModificationRequired:
			return new(e1ap_ies.BearerContextModificationRequired)
		case e1ap_ies.ProcedureCodeBearerContextRelease:
			return new(e1ap_ies.BearerContextReleaseCommand)
		case e1ap_ies.ProcedureCodeBearerContextReleaseRequest:
			return new(e1ap_ies.BearerContextReleaseRequest)
		case e1ap_ies.ProcedureCodeBearerContextSetup:
			return new(e1ap_ies.BearerContextSetupRequest)
		case e1ap_ies.ProcedureCodeCellTrafficTrace:
			return new(e1ap_ies.CellTrafficTrace)
		case e1ap_ies.ProcedureCodeDeactivateTrace:
			return new(e1ap_ies.DeactivateTrace)
		case e1ap_ies.ProcedureCodeDLDataNotification:
			return new(e1ap_ies.DLDataNotification)
		case e1ap_ies.ProcedureCodeDataUsageReport:
			return new(e1ap_ies.DataUsageReport)
		case e1ap_ies.ProcedureCodeE1Release:
			return new(e1ap_ies.E1ReleaseRequest)
		case e1ap_ies.ProcedureCodeGNBCUCPE1Setup:
			return new(e1ap_ies.GNBCUCPE1SetupRequest)
		case e1ap_ies.ProcedureCodeGNBCUUPE1Setup:
			return new(e1ap_ies.GNBCUUPE1SetupRequest)
		case e1ap_ies.ProcedureCodeEarlyForwardingSNTransfer:
			return new(e1ap_ies.EarlyForwardingSNTransfer)
		case e1ap_ies.ProcedureCodeErrorIndication:
			return new(e1ap_ies.ErrorIndication)
		case e1ap_ies.ProcedureCodeGNBCUCPConfigurationUpdate:
			return new(e1ap_ies.GNBCUCPConfigurationUpdate)
		case e1ap_ies.ProcedureCodeGNBCUUPConfigurationUpdate:
			return new(e1ap_ies.GNBCUUPConfigurationUpdate)
		case e1ap_ies.ProcedureCodeGNBCUUPCounterCheck:
			return new(e1ap_ies.GNBCUUPCounterCheckRequest)
		case e1ap_ies.ProcedureCodeGNBCUUPStatusIndication:
			return new(e1ap_ies.GNBCUUPStatusIndication)
		case e1ap_ies.ProcedureCodeIABUPTNLAddressUpdate:
			return new(e1ap_ies.IABUPTNLAddressUpdate)
		case e1ap_ies.ProcedureCodeMRDCDataUsageReport:
			return new(e1ap_ies.MRDCDataUsageReport)
		case e1ap_ies.ProcedureCodePrivateMessage:
			return new(e1ap_ies.PrivateMessage)
		case e1ap_ies.ProcedureCodeReset:
			return new(e1ap_ies.Reset)
		case e1ap_ies.ProcedureCodeResourceStatusReporting:
			return new(e1ap_ies.ResourceStatusRequest)
		case e1ap_ies.ProcedureCodeResourceStatusReportingInitiation:
			return new(e1ap_ies.ResourceStatusUpdate)
		case e1ap_ies.ProcedureCodeTraceStart:
			return new(e1ap_ies.TraceStart)
		case e1ap_ies.ProcedureCodeULDataNotification:
			return new(e1ap_ies.ULDataNotification)
		}

	case PduChoiceSuccessfulOutcome:
		switch procedureCode.Value {
		case e1ap_ies.ProcedureCodeBearerContextModification:
			return new(e1ap_ies.BearerContextModificationResponse)
		case e1ap_ies.ProcedureCodeBearerContextModificationRequired:
			return new(e1ap_ies.BearerContextModificationConfirm)
		case e1ap_ies.ProcedureCodeBearerContextRelease:
			return new(e1ap_ies.BearerContextReleaseComplete)
		case e1ap_ies.ProcedureCodeBearerContextSetup:
			return new(e1ap_ies.BearerContextSetupResponse)
		case e1ap_ies.ProcedureCodeE1Release:
			return new(e1ap_ies.E1ReleaseResponse)
		case e1ap_ies.ProcedureCodeGNBCUCPE1Setup:
			return new(e1ap_ies.GNBCUCPE1SetupResponse)
		case e1ap_ies.ProcedureCodeGNBCUCPConfigurationUpdate:
			return new(e1ap_ies.GNBCUCPConfigurationUpdateAcknowledge)
		case e1ap_ies.ProcedureCodeGNBCUUPConfigurationUpdate:
			return new(e1ap_ies.GNBCUUPConfigurationUpdateAcknowledge)
		case e1ap_ies.ProcedureCodeIABUPTNLAddressUpdate:
			return new(e1ap_ies.IABUPTNLAddressUpdateAcknowledge)
		case e1ap_ies.ProcedureCodeReset:
			return new(e1ap_ies.ResetAcknowledge)
		case e1ap_ies.ProcedureCodeResourceStatusReporting:
			return new(e1ap_ies.ResourceStatusResponse)
		}
	case PduChoiceUnsuccessfulOutcome:
		switch procedureCode.Value {
		case e1ap_ies.ProcedureCodeBearerContextModification:
			return new(e1ap_ies.BearerContextModificationFailure)
		case e1ap_ies.ProcedureCodeBearerContextSetup:
			return new(e1ap_ies.BearerContextSetupFailure)
		case e1ap_ies.ProcedureCodeGNBCUCPE1Setup:
			return new(e1ap_ies.GNBCUCPE1SetupFailure)
		case e1ap_ies.ProcedureCodeGNBCUCPConfigurationUpdate:
			return new(e1ap_ies.GNBCUCPConfigurationUpdateFailure)
		case e1ap_ies.ProcedureCodeGNBCUUPConfigurationUpdate:
			return new(e1ap_ies.GNBCUUPConfigurationUpdateFailure)
		case e1ap_ies.ProcedureCodeIABUPTNLAddressUpdate:
			return new(e1ap_ies.IABUPTNLAddressUpdateFailure)
		case e1ap_ies.ProcedureCodeResourceStatusReporting:
			return new(e1ap_ies.ResourceStatusFailure)
		}
	}

	return nil
}
