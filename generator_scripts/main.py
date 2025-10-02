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
    logging.basicConfig(
        level=level, format="[%(levelname)s] %(message)s", stream=sys.stdout
    )


def load_lines_with_context(directory: str) -> List[Tuple[str, str, int]]:
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

    log_level = logging.DEBUG if args.verbose else logging.INFO
    setup_logging(log_level)

    script_dir = os.path.dirname(os.path.realpath(__file__))
    base_dir = os.path.dirname(script_dir)
    input_dir = os.path.join(base_dir, "extracted_sections")
    output_dir = os.path.join(base_dir, "e1ap_ies")

    logger.info("--- STAGE 1: LOADING FILES ---")
    lines = load_lines_with_context(input_dir)
    if not lines:
        sys.exit(1)

    logger.info("--- STAGE 2: PARSING DEFINITIONS ---")
    asn1_parser = ASN1Parser(lines)
    definitions, failures = asn1_parser.parse()
    logger.info(
        f"Parsing complete. Found {len(definitions)} definitions and {len(failures)} failures."
    )

    logger.info("--- STAGE 2.5: APPLYING MANUAL PATCHES ---")
    patcher = ASN1Patcher()
    hardcoded_defs = patcher.get_hardcoded_definitions()
    for definition in hardcoded_defs:
        definitions[definition.name] = definition
    logger.info(
        f"Applied {len(hardcoded_defs)} hardcoded definitions. Total definitions now: {len(definitions)}."
    )

    generator = GoCodeGenerator(definitions, failures, output_dir)

    if args.diagnose:
        logger.info("--- STAGE 3: RUNNING DIAGNOSTIC REPORT ---")
        generator.run_full_diagnostic()
    else:
        logger.info("--- STAGE 3: GENERATING GO CODE ---")
        generator.generate_files()
        logger.info("Code generation process complete.")


if __name__ == "__main__":
    main()
