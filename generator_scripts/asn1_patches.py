from common_types import ASN1Definition
from typing import List


class ASN1Patcher:
    def get_hardcoded_definitions(self) -> List[ASN1Definition]:
        definitions = []

        cp_measurement_id = ASN1Definition(
            name="id-gNB-CU-CP-Measurement-ID",
            def_type="INTEGER",
            source_file="manual_patch",
            full_text="MANUALLY CREATED DEFINITION FOR id-gNB-CU-CP-Measurement-ID",
        )
        cp_measurement_id.min_val = "1"
        cp_measurement_id.max_val = "4095"
        definitions.append(cp_measurement_id)

        up_measurement_id = ASN1Definition(
            name="id-gNB-CU-UP-Measurement-ID",
            def_type="INTEGER",
            source_file="manual_patch",
            full_text="MANUALLY CREATED DEFINITION FOR id-gNB-CU-UP-Measurement-ID",
        )
        up_measurement_id.min_val = "1"
        up_measurement_id.max_val = "4095"
        definitions.append(up_measurement_id)

        drb_released = ASN1Definition(
            name="dRB-Released-In-Session",
            def_type="ENUMERATED",
            source_file="manual_patch",
            full_text="MANUALLY CREATED DEFINITION FOR dRB-Released-In-Session",
        )
        drb_released.enum_values = ["released-in-session", "not-released-in-session"]
        definitions.append(drb_released)

        secondary_rat = ASN1Definition(
            name="secondaryRATType",
            def_type="ENUMERATED",
            source_file="manual_patch",
            full_text="MANUALLY CREATED DEFINITION FOR secondaryRATType",
        )
        secondary_rat.enum_values = ["nr", "e-UTRA"]
        definitions.append(secondary_rat)

        diag_list = ASN1Definition(
            name="CriticalityDiagnostics-IE-List",
            def_type="LIST",
            source_file="manual_patch",
            full_text="MANUALLY CREATED DEFINITION FOR CriticalityDiagnostics-IE-List"
        )
        diag_list.of_type = "CriticalityDiagnostics-IE-Item"
        definitions.append(diag_list)

        return definitions