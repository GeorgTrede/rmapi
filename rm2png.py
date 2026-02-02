import os
import sys
import zipfile
import tempfile
import argparse
import logging
import re
import io
from pathlib import Path

import cairosvg
from PIL import Image

from rmscene import read_tree
from rmc.exporters.svg import tree_to_svg

# Configure logging to suppress warnings from rmscene
logging.basicConfig(level=logging.ERROR)


def fix_svg_content(svg_content: str) -> str:
    """Fix common XML issues in the generated SVG (keep original behavior)."""
    # Escape ampersands that are not already escaped
    # This regex finds ampersands that are not followed by 'amp;', 'lt;', 'gt;', 'quot;', 'apos;'
    return re.sub(r'&(?!(amp|lt|gt|quot|apos);)', '&amp;', svg_content)


def convert_rmdoc_to_png(rmdoc_path: str, output_dir: str, scale_factor: float = 2.0, margin_px: int = 50) -> None:
    """
    Convert a .rmdoc into per-page PNGs.

    This stays faithful to the original script:
    * generate SVG via tree_to_svg
    * apply only the existing ampersand fix
    * rasterize with CairoSVG

    Enhancements (without modifying SVG geometry):
    * higher PNG resolution via CairoSVG `scale`
    * margin added AFTER rasterization (padding in pixel space)
    """
    if not os.path.exists(output_dir):
        os.makedirs(output_dir)

    with tempfile.TemporaryDirectory() as tmpdir:
        print(f"Extracting {rmdoc_path}...")
        try:
            with zipfile.ZipFile(rmdoc_path, "r") as zip_ref:
                zip_ref.extractall(tmpdir)
        except zipfile.BadZipFile:
            print(f"Error: {rmdoc_path} is not a valid zip file.")
            return

        rm_files = []
        for root, dirs, files in os.walk(tmpdir):
            for file in files:
                if file.endswith(".rm"):
                    rm_files.append(os.path.join(root, file))

        if not rm_files:
            print("No .rm files found in the document.")
            return

        print(f"Found {len(rm_files)} pages. Converting all layers...")

        for rm_file in rm_files:
            page_id = os.path.splitext(os.path.basename(rm_file))[0]
            svg_path = os.path.join(tmpdir, f"{page_id}.svg")
            png_path = os.path.join(output_dir, f"{page_id}.png")

            try:
                # Read reMarkable page
                with open(rm_file, "rb") as infile:
                    tree = read_tree(infile)

                # Export SVG (original behavior)
                with open(svg_path, "wt", encoding="utf-8") as outfile:
                    tree_to_svg(tree, outfile)

                # Apply only the original XML text fix (original behavior)
                with open(svg_path, "r", encoding="utf-8") as f:
                    svg_content = f.read()

                fixed_content = fix_svg_content(svg_content)

                with open(svg_path, "w", encoding="utf-8") as f:
                    f.write(fixed_content)
                
                # COPY SVG TO OUTPUT DIR FOR ANALYSIS
                import shutil
                shutil.copy(svg_path, os.path.join(output_dir, f"{page_id}.svg"))

                # Rasterize to PNG at higher resolution (no SVG geometry changes)
                png_bytes = cairosvg.svg2png(
                    url=svg_path,
                    background_color="white",
                    scale=scale_factor,
                )

                # Add margin in pixel space (prevents any viewBox/shift issues)
                if margin_px > 0:
                    img = Image.open(io.BytesIO(png_bytes)).convert("RGB")
                    w, h = img.size
                    out = Image.new("RGB", (w + 2 * margin_px, h + 2 * margin_px), "white")
                    out.paste(img, (margin_px, margin_px))
                    out.save(png_path)
                else:
                    # Write bytes directly if no margin requested
                    with open(png_path, "wb") as f:
                        f.write(png_bytes)

                print(f"Converted page {page_id} to PNG.")

            except Exception as e:
                print(f"Error converting page {page_id}: {e}")


def main() -> None:
    parser = argparse.ArgumentParser(description="Convert reMarkable .rmdoc files to PNG images.")
    parser.add_argument("input", help="Path to the .rmdoc file")
    parser.add_argument("-o", "--output", help="Output directory for PNG files", default="output_png")

    # Enhancements that do not alter the SVG geometry
    parser.add_argument(
        "-s",
        "--scale",
        type=float,
        default=2.0,
        help="Resolution scale factor for PNG rasterization (default: 2.0).",
    )
    parser.add_argument(
        "-m",
        "--margin",
        type=int,
        default=50,
        help="Margin around the drawing in output pixels (default: 50).",
    )

    args = parser.parse_args()

    if not os.path.exists(args.input):
        print(f"Error: File {args.input} not found.")
        sys.exit(1)

    convert_rmdoc_to_png(args.input, args.output, args.scale, args.margin)
    print(f"Conversion complete. Results are in '{args.output}'.")


if __name__ == "__main__":
    main()

