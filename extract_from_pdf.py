import pdfplumber
import os
import re
import textwrap

PDF_FILENAME = "3GPP_TS_38_463.pdf"  # Make sure this PDF is in the same folder
OUTPUT_DIR = "extracted_sections"

def is_in_bounds(element, header_box, footer_box):
    """Check if a pdfplumber element is outside the header/footer boxes."""
    x0, top, x1, bottom = element['x0'], element['top'], element['x1'], element['bottom']
    
    # Check if element is inside the header
    if top < header_box['bottom']:
        return False
        
    # Check if element is inside the footer
    if bottom > footer_box['top']:
        return False
        
    return True

def reconstruct_page_layout(page, header_bbox, footer_bbox):
    """
    Reconstructs the text layout of a single page by analyzing character positions,
    skipping anything in the header or footer. Guarantees 1:1 formatting.
    """
    # 1. Filter out characters in the header/footer zones
    content_chars = [char for char in page.chars if is_in_bounds(char, header_bbox, footer_bbox)]
    
    # 2. Sort characters for reading order: top-to-bottom, then left-to-right
    # The 'top' value is the primary sort key, 'x0' is the secondary.
    sorted_chars = sorted(content_chars, key=lambda c: (c['top'], c['x0']))

    if not sorted_chars:
        return ""

    # 3. Reconstruct the text with precise spacing
    reconstructed_text = ""
    current_line_top = -1
    last_char_x1 = 0
    
    # Estimate the width of a single space character
    # This helps calculate how many spaces to insert for large gaps.
    avg_space_width = page.width / 120 # A reasonable heuristic for standard fonts

    for char in sorted_chars:
        # Check for a new line (a significant jump in vertical position)
        if char['top'] > current_line_top + 2:  # Using a small tolerance
            reconstructed_text += '\n'
            current_line_top = char['top']
            last_char_x1 = 0 # Reset horizontal position for the new line

        # Calculate the horizontal gap since the last character
        gap = char['x0'] - last_char_x1
        
        # If there's a noticeable gap, insert the appropriate number of spaces
        if gap > 1: # A gap of 1 point or less is negligible
            num_spaces = round(gap / avg_space_width)
            reconstructed_text += ' ' * int(num_spaces)
        
        # Append the actual character
        reconstructed_text += char['text']
        last_char_x1 = char['x1']
        
    return reconstructed_text.lstrip('\n')

def refine_formatting(section_text):
    """
    Refines a block of text by:
    1. Removing any uniform leading indentation from the entire block.
    2. Removing any blank lines that appear at the very start of the content.
    """
    # 1. Remove the uniform over-indentation from the entire block.
    # textwrap.dedent() is specifically designed for this.
    dedented_text = textwrap.dedent(section_text)

    # 2. Remove leading blank lines after the title.
    lines = dedented_text.splitlines()
    if len(lines) <= 1:
        return dedented_text # Nothing to do if there's only a title

    title_line = lines[0]
    body_lines = lines[1:]
    
    first_content_index = 0
    # Find the first line in the body that isn't just whitespace
    for i, line in enumerate(body_lines):
        if line.strip():
            first_content_index = i
            break
            
    # Join the title with the body, starting from the first real content line
    final_content = title_line + '\n' + '\n'.join(body_lines[first_content_index:])
    
    return final_content
    

def extract_and_split_sections():
    """
    (Final Version)
    Reads the PDF, reconstructs content, refines formatting, and splits into section files.
    """
    if not os.path.exists(PDF_FILENAME):
        print(f"Error: PDF file '{PDF_FILENAME}' not found. Please place it in the same directory.")
        return

    full_content = []
    
    with pdfplumber.open(PDF_FILENAME) as pdf:
        print(f"Processing {len(pdf.pages)} pages from '{PDF_FILENAME}'...")
        
        header_bbox = {'bottom': 50}
        footer_bbox = {'top': pdf.pages[0].height - 40}

        for i, page in enumerate(pdf.pages):
            page_text = reconstruct_page_layout(page, header_bbox, footer_bbox)
            full_content.append(page_text)
    
    cleaned_text = "\n".join(full_content)
    
    section_pattern = re.compile(r'^\s*9\.4\.\d+ .*$', re.MULTILINE)
    matches = list(section_pattern.finditer(cleaned_text))

    if not matches:
        print("\nCould not find any section headings after processing the PDF.")
        return

    if not os.path.exists(OUTPUT_DIR):
        os.makedirs(OUTPUT_DIR)

    print(f"\nFound {len(matches)} sections. Refining and writing files...")

    for i, match in enumerate(matches):
        title = match.group(0).strip()
        start_index = match.start()
        end_index = matches[i + 1].start() if i + 1 < len(matches) else len(cleaned_text)
        
        # Extract the raw section slice
        section_content_slice = cleaned_text[start_index:end_index]
        
        # **NEW: Apply the final formatting refinements**
        final_content = refine_formatting(section_content_slice)

        # Sanitize title for the filename
        temp_title = title.replace('.', '_')
        sanitized_title = re.sub(r'[^\w-]+', '_', temp_title)
        filename = f"{sanitized_title}.txt"
        full_path = os.path.join(OUTPUT_DIR, filename)

        try:
            with open(full_path, 'w', encoding='utf-8') as f:
                # Write the fully refined content
                f.write(final_content)
            print(f"  - Successfully wrote: {filename}")
        except IOError as e:
            print(f"  - Failed to write {filename}: {e}")

    print("\nExtraction complete. All refinements have been applied.")
# --- Main execution ---
if __name__ == "__main__":
    extract_and_split_sections()