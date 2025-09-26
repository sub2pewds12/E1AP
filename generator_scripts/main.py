# In main.py

import os
import sys
import logging
import argparse # New import for command-line arguments
from asn1_parser import ASN1Parser
from go_generator import GoCodeGenerator

# --- NEW: Set up a logger ---
logger = logging.getLogger(__name__)

# --- NEW: Logging configuration function ---
def setup_logging(level=logging.INFO):
    """Configures the root logger."""
    logging.basicConfig(
        level=level,
        format='[%(levelname)s] %(message)s',
        stream=sys.stdout,
    )

# --- Updated: Paths are now inside the main function ---

def load_and_combine_specs(directory: str) -> str:
    """Loads all .txt files from a directory and combines them into a single string."""
    if not os.path.exists(directory):
        # Use logger.error instead of raising an exception directly in this function
        logger.error(f"Input directory '{directory}' not found.")
        return None
        
    section_files = sorted([f for f in os.listdir(directory) if f.endswith('.txt')])
    if not section_files:
        logger.error(f"No '.txt' files found in '{directory}'.")
        return None

    all_content = []
    logger.info(f"Found {len(section_files)} section files to load...")
    for filename in section_files:
        filepath = os.path.join(directory, filename)
        logger.debug(f"Loading file: {filepath}")
        with open(filepath, 'r', encoding='utf-8') as f:
            all_content.append(f.read())
            
    logger.info(f"Successfully loaded and combined {len(section_files)} files.")
    return "\n".join(all_content)

def main():
    """Main execution function."""
    # --- NEW: Command-line argument parsing ---
    parser = argparse.ArgumentParser(description="Parse ASN.1 specs and generate Go code.")
    parser.add_argument(
        '-v', '--verbose',
        action='store_true',
        help='Enable detailed DEBUG logging.'
    )
    args = parser.parse_args()

    # --- NEW: Setup logging based on arguments ---
    log_level = logging.DEBUG if args.verbose else logging.INFO
    setup_logging(log_level)

    # Make paths relative to the project's base directory
    script_dir = os.path.dirname(os.path.realpath(__file__))
    base_dir = os.path.dirname(script_dir)
    input_sections_dir = os.path.join(base_dir, "extracted_sections")
    output_dir = os.path.join(base_dir, "e1ap_ies")

    logger.info("--- STAGE 1: LOADING FILES ---")
    e1ap_asn1_spec = load_and_combine_specs(input_sections_dir)
    if not e1ap_asn1_spec:
        logger.critical("Halting execution due to file loading errors.")
        sys.exit(1)

    logger.info("--- STAGE 2: PARSING DEFINITIONS ---")
    asn1_parser = ASN1Parser(e1ap_asn1_spec)
    definitions = asn1_parser.parse()
    logger.info(f"Parsing complete. Found {len(definitions)} total definitions.")
    
    logger.info("--- STAGE 3: RESOLVING TYPES ---")
    resolved_count, unresolved_set = asn1_parser.resolve_type_references()
    logger.info(f"Type resolution complete. Linked {resolved_count} custom type references.")
    if unresolved_set:
        logger.warning(f"Found {len(unresolved_set)} unresolved types: {unresolved_set}")


    logger.info("--- STAGE 4: GENERATING GO CODE ---")
    generator = GoCodeGenerator(definitions, output_dir)
    generator.generate_files()
    logger.info("Code generation process complete.")

if __name__ == "__main__":
    main()