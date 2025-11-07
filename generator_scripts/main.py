import os
import sys
import logging
import argparse
import re
from typing import List, Tuple
from asn1_parser import ASN1Parser
from go_generator import GoCodeGenerator
from asn1_patches import ASN1Patcher
from common_types import ASN1Definition



logger = logging.getLogger(__name__)


def setup_logging(level=logging.INFO):
    if not logging.getLogger().handlers:
        logging.basicConfig(
            level=level, format="[%(levelname)s] %(message)s", stream=sys.stdout
        )

def load_simple_list(filepath: str) -> List[str]:
    """Loads a simple list from a text file, one item per line."""
    if not os.path.exists(filepath):
        return []
    with open(filepath, 'r', encoding='utf-8') as f:
        return [line.strip() for line in f if line.strip()]

def load_lines_with_context(directory: str, files_to_skip: List[str]) -> List[Tuple[str, str, int]]:
    if not os.path.exists(directory):
        logger.error(f"Input directory '{directory}' not found.")
        return None

    section_files = sorted([f for f in os.listdir(directory) if f.endswith(".txt")])
    if not section_files:
        logger.error(f"No '.txt' files found in '{directory}'.")
        return None

    all_lines = []
    logger.info(f"Found {len(section_files)} section files to load...")

    for filename in section_files:
        if filename in set(files_to_skip):
            logger.warning(f"Skipping file '{filename}' as it is handled by manual patches.")
            continue 

        filepath = os.path.join(directory, filename)
        with open(filepath, "r", encoding="utf-8") as f:
            for i, line_text in enumerate(f, 1):
                all_lines.append((line_text, filename, i))

    logger.info(f"Successfully loaded and combined {len(all_lines)} total lines.")
    return all_lines

def main():
    parser = argparse.ArgumentParser(
        description="Parse 3GPP E1AP ASN.1 specs and generate Go code."
    )
    parser.add_argument(
        "-i",
        "--input-dir",
        required=True,
        help="Path to the directory containing the extracted ASN.1 '.txt' files.",
    )
    parser.add_argument(
        "-o",
        "--output-dir",
        required=True,
        help="Path to the directory where the generated '.go' files will be written.",
    )
    parser.add_argument(
        "-v",
        "--verbose",
        action="store_true",
        help="Enable detailed DEBUG logging for deep analysis.",
    )
    parser.add_argument(
        "--diagnose",
        action="store_true",
        help="Run a full diagnostic report on parsed and failed definitions instead of generating files.",
    )
    args = parser.parse_args()
    script_dir = os.path.dirname(os.path.abspath(__file__))
    config_dir = os.path.join(script_dir, "..", "config")
    skip_files_path = os.path.join(config_dir, "skip_files.txt")
    acronyms_path = os.path.join(config_dir, "acronyms.txt")

    log_level = logging.DEBUG if args.verbose else logging.INFO
    setup_logging(log_level)

    logger.info("--- STAGE 1: LOADING FILES ---")

    files_to_skip = load_simple_list(skip_files_path)
    acronyms = load_simple_list(acronyms_path)
    
    lines = load_lines_with_context(args.input_dir, files_to_skip)
    if not lines:
        sys.exit(1)

    logger.info("--- STAGE 2: PARSING DEFINITIONS ---")
    asn1_parser = ASN1Parser(lines)
    definitions, failures, procedures, message_map = asn1_parser.parse()
    logger.info(
        f"Parsing complete. Found {len(definitions)} definitions, {len(procedures)} procedures, and {len(failures)} failures."
    )

    logger.info("--- STAGE 2.5: APPLYING MANUAL PATCHES ---")
    patcher = ASN1Patcher(config_dir, asn1_parser)
    hardcoded_defs = patcher.get_hardcoded_definitions(definitions)

    output_dir = args.output_dir
    definitions.update(hardcoded_defs)
    
    definitions.update(hardcoded_defs)
    logger.info(
        f"Applied {len(hardcoded_defs)} hardcoded definitions. Total definitions now: {len(definitions)}."
    )

    if hardcoded_defs:
        patched_names = set(hardcoded_defs.keys())
        original_failure_count = len(failures)
    
        failures[:] = [f for f in failures if f['name'].split(' ')[0] not in patched_names]
        
        rescued_count = original_failure_count - len(failures)
        if rescued_count > 0:
            logger.info(f"Rescued {rescued_count} definitions from the failure list via manual patches.")

    if failures:
            
            logger.warning("--- The following definitions remain unparsed (after patches): ---")
            for failure in sorted(failures, key=lambda x: x['name']):
                logger.warning(f"  - {failure['name']}")

    if "ResetType" in definitions:
        reset_type_def = definitions["ResetType"]
        for ie in reset_type_def.ies:
            if ie.ie == "partOfE1-Interface":
                logger.info("Applying post-processing fix for ResetType 'partOfE1-Interface' member.")
                ie.type = "UE-associatedLogicalE1-ConnectionListRes"

    generator = GoCodeGenerator(definitions, failures, procedures, message_map, output_dir, patcher, acronyms, asn1_parser)
    if args.diagnose:
        logger.info("--- STAGE 3: RUNNING DIAGNOSTIC REPORT ---")
        generator.run_full_diagnostic()
    else:
        logger.info("--- STAGE 3: GENERATING GO CODE ---")
        generator.generate_files()
        logger.info("Code generation process complete.")

if __name__ == "__main__":
    main()
