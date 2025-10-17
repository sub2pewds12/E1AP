import os
import yaml
import logging
from typing import Dict
from common_types import (
    SequenceDefinition, ListDefinition, InformationElement, ASN1Definition
)

logger = logging.getLogger(__name__)

class ASN1Patcher:
    def __init__(self, config_dir: str):
        self.config_dir = config_dir
        self.patch_definitions_path = os.path.join(config_dir, "patches.yaml")
        self.go_patches_dir = os.path.join(config_dir, "go_patches")

    def get_hardcoded_definitions(self, existing_definitions: Dict) -> Dict[str, ASN1Definition]:
        """Loads and builds patch definitions from an external YAML file."""
        definitions = {}
        try:
            with open(self.patch_definitions_path, 'r') as f:
                patch_config = yaml.safe_load(f)
        except FileNotFoundError:
            logger.warning("patches.yaml not found. No in-memory patches will be applied.")
            return {}

        for p_def in patch_config.get("definitions", []):
            name = p_def["name"]
            def_type = p_def["type"]

            if def_type == "SEQUENCE":
                item = SequenceDefinition(name=name, source_file="manual_patch")
                
                # Logic to populate members from an IE-Set
                if "from_ie_set" in p_def:
                    ie_set_name = p_def["from_ie_set"]
                    ie_set = existing_definitions.get(ie_set_name)
                    if ie_set:
                        # This part depends on how you store IE sets. Adjust if necessary.
                        # Assuming ie_set is a SequenceDefinition whose IEs we can iterate.
                        for ie_data in ie_set.ies:
                            ie_name_to_add = ie_data.ie.replace("id-", "")
                            new_ie = InformationElement(
                                ie=ie_name_to_add, type=ie_data.type, 
                                presence=ie_data.presence, criticality=ie_data.criticality
                            )
                            item.ies.append(new_ie)
                    else:
                        logger.error(f"Patcher could not find required IE-Set '{ie_set_name}' for patch '{name}'")

                # Logic to populate members from an explicit list
                if "members" in p_def:
                    for member in p_def["members"]:
                        item.ies.append(InformationElement(**member))
                
                definitions[name] = item

            elif def_type == "LIST":
                item = ListDefinition(name=name, source_file="manual_patch")
                item.of_type = p_def.get("of_type")
                item.max_val = p_def.get("max_val")
                definitions[name] = item
        
        return definitions

    def patch_generated_files(self, output_dir: str):
        """Applies file-level patches from the go_patches directory."""
        logger.info("--- Applying post-generation file patches ---")
        if not os.path.exists(self.go_patches_dir):
            logger.warning(f"Go patches directory '{self.go_patches_dir}' not found. Skipping file patches.")
            return

        # Part 1: Delete junk files (this can also be moved to config if it grows)
        files_to_delete = [
            "ue_associated_logical_e1_connection_list_res_item_ack_item.go",
            "ue_associated_logical_e1_connection_list_res_item.go",
        ]
        for filename in files_to_delete:
            filepath = os.path.join(output_dir, filename)
            if os.path.exists(filepath):
                os.remove(filepath)
                logger.info(f"PATCHER: Deleted junk file: {filename}")

        # Part 2: Overwrite files with correct content from templates
        for template_name in os.listdir(self.go_patches_dir):
            if template_name.endswith(".go.template"):
                output_filename = template_name.replace(".template", "")
                
                template_path = os.path.join(self.go_patches_dir, template_name)
                output_path = os.path.join(output_dir, output_filename)
                
                with open(template_path, 'r', encoding='utf-8') as f:
                    content = f.read()
                
                with open(output_path, 'w', encoding='utf-8') as f:
                    f.write(content)
                logger.info(f"PATCHER: Wrote correct content to: {output_filename}")
