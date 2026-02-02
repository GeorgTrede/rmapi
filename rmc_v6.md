================================================
FILE: README.md
================================================
# rmc

Command line tool for converting to/from remarkable `.rm` version 6 (software version 3) files.

## Installation

If you want to render your documents as PDF or SVG, you will need to install cairo.

- On Windows, you need to install cairo using [msys2](https://packages.msys2.org/packages/mingw-w64-x86_64-gtk3)
- On MacOS, install cairo using homebrew: `brew install cairo libxml2 libffi`
- On Linux, the right packages should be installed through pip. If not, refer to [here](https://github.com/Kozea/CairoSVG/issues/371) for the right dependencies.

To install in your current Python environment:

    pip install rmc
    
Or use [pipx](https://pypa.github.io/pipx/) to install in an isolated environment (recommended):

    pipx install rmc

## Usage

Convert a remarkable v6 file to other formats, specified by `-t FORMAT`:

    $ rmc -t markdown file.rm
    Text in the file is printed to standard output.

Specify the filename to write the output to with `-o`:

    $ rmc -t svg -o file.svg file.rm
    
The format is guessed based on the filename if not specified:
    
    $ rmc file.rm -o file.pdf

Create a `.rm` file containing the text in `text.md`:

    $ rmc -t rm text.md -o text.rm

## SVG/PDF Conversion Status

Right now the converter works well while there are no text boxes. If you add text boxes, there are x issues:

1. if the text box contains multiple lines, the lines are actually printed in the same line, and
2. the position of the strokes gets corrupted.

# Acknowledgements

`rmc` uses [rmscene](https://github.com/ricklupton/rmscene) to read the `.rm` files, for which https://github.com/ddvk/reader helped a lot in figuring out the structure and meaning of the files.

[@chemag](https://github.com/chemag) added initial support for converting to svg and pdf.

[@Seb-sti1](https://github.com/Seb-sti1) made lots of improvements to svg export and updating to newer `rmscene` versions.

[@ChenghaoMou](https://github.com/ChenghaoMou) added support for new pen types/colours.

[@EelcovanVeldhuizen](https://github.com/EelcovanVeldhuizen) for code updates/fixes.

[@p4xel](https://github.com/p4xel) for code fixes.



================================================
FILE: convert_test_files.sh
================================================
#!/bin/bash

set -euo pipefail

# Directory containing .rm files
TEST_DIR="tests/rm"
OUTPUT_DIR="test_output"

# Check if directory exists
if [ ! -d "$TEST_DIR" ]; then
    echo "Error: Directory $TEST_DIR does not exist"
    exit 1
fi

if [ ! -d "$OUTPUT_DIR" ]; then
    mkdir -p "$OUTPUT_DIR"
    mkdir -p "$OUTPUT_DIR/markdown"
    mkdir -p "$OUTPUT_DIR/svg"
    mkdir -p "$OUTPUT_DIR/pdf"
fi

# Iterate through all .rm files in the directory
for file in "$TEST_DIR"/*.rm; do
    # Check if files exist (in case directory is empty)
    if [ -f "$file" ]; then
        echo "Testing file: $file"
        file_name=$(basename "$file")
        
        # Run first test command
        echo "Running test markdown..."
        rmc -t markdown "$file" -o "$OUTPUT_DIR/markdown/$file_name.md"
        
        # Run second test command
        echo "Running test svg..."
        rmc -t svg "$file" -o "$OUTPUT_DIR/svg/$file_name.svg"

        # Run third test command
        echo "Running test pdf..."
        rmc -t pdf "$file" -o "$OUTPUT_DIR/pdf/$file_name.pdf"

        echo "----------------------------------------"
    fi
done

echo "All tests completed"



================================================
FILE: flake.lock
================================================
{
  "nodes": {
    "flake-utils": {
      "inputs": {
        "systems": "systems"
      },
      "locked": {
        "lastModified": 1731533236,
        "narHash": "sha256-l0KFg5HjrsfsO/JpG+r7fRrqm12kzFHyUHqHCVpMMbI=",
        "owner": "numtide",
        "repo": "flake-utils",
        "rev": "11707dc2f618dd54ca8739b309ec4fc024de578b",
        "type": "github"
      },
      "original": {
        "owner": "numtide",
        "repo": "flake-utils",
        "type": "github"
      }
    },
    "flake-utils_2": {
      "inputs": {
        "systems": "systems_2"
      },
      "locked": {
        "lastModified": 1726560853,
        "narHash": "sha256-X6rJYSESBVr3hBoH0WbKE5KvhPU5bloyZ2L4K60/fPQ=",
        "owner": "numtide",
        "repo": "flake-utils",
        "rev": "c1dfcf08411b08f6b8615f7d8971a2bfa81d5e8a",
        "type": "github"
      },
      "original": {
        "owner": "numtide",
        "repo": "flake-utils",
        "type": "github"
      }
    },
    "nix-github-actions": {
      "inputs": {
        "nixpkgs": [
          "poetry2nix",
          "nixpkgs"
        ]
      },
      "locked": {
        "lastModified": 1729742964,
        "narHash": "sha256-B4mzTcQ0FZHdpeWcpDYPERtyjJd/NIuaQ9+BV1h+MpA=",
        "owner": "nix-community",
        "repo": "nix-github-actions",
        "rev": "e04df33f62cdcf93d73e9a04142464753a16db67",
        "type": "github"
      },
      "original": {
        "owner": "nix-community",
        "repo": "nix-github-actions",
        "type": "github"
      }
    },
    "nixpkgs": {
      "locked": {
        "lastModified": 1733686850,
        "narHash": "sha256-NQEO/nZWWGTGlkBWtCs/1iF1yl2lmQ1oY/8YZrumn3I=",
        "owner": "NixOS",
        "repo": "nixpkgs",
        "rev": "dd51f52372a20a93c219e8216fe528a648ffcbf4",
        "type": "github"
      },
      "original": {
        "id": "nixpkgs",
        "ref": "nixpkgs-unstable",
        "type": "indirect"
      }
    },
    "nixpkgs_2": {
      "locked": {
        "lastModified": 1730157240,
        "narHash": "sha256-P8wF4ag6Srmpb/gwskYpnIsnspbjZlRvu47iN527ABQ=",
        "owner": "NixOS",
        "repo": "nixpkgs",
        "rev": "75e28c029ef2605f9841e0baa335d70065fe7ae2",
        "type": "github"
      },
      "original": {
        "owner": "NixOS",
        "repo": "nixpkgs",
        "rev": "75e28c029ef2605f9841e0baa335d70065fe7ae2",
        "type": "github"
      }
    },
    "poetry2nix": {
      "inputs": {
        "flake-utils": "flake-utils_2",
        "nix-github-actions": "nix-github-actions",
        "nixpkgs": "nixpkgs_2",
        "systems": "systems_3",
        "treefmt-nix": "treefmt-nix"
      },
      "locked": {
        "lastModified": 1743690424,
        "narHash": "sha256-cX98bUuKuihOaRp8dNV1Mq7u6/CQZWTPth2IJPATBXc=",
        "owner": "nix-community",
        "repo": "poetry2nix",
        "rev": "ce2369db77f45688172384bbeb962bc6c2ea6f94",
        "type": "github"
      },
      "original": {
        "owner": "nix-community",
        "repo": "poetry2nix",
        "type": "github"
      }
    },
    "root": {
      "inputs": {
        "flake-utils": "flake-utils",
        "nixpkgs": "nixpkgs",
        "poetry2nix": "poetry2nix"
      }
    },
    "systems": {
      "locked": {
        "lastModified": 1681028828,
        "narHash": "sha256-Vy1rq5AaRuLzOxct8nz4T6wlgyUR7zLU309k9mBC768=",
        "owner": "nix-systems",
        "repo": "default",
        "rev": "da67096a3b9bf56a91d16901293e51ba5b49a27e",
        "type": "github"
      },
      "original": {
        "owner": "nix-systems",
        "repo": "default",
        "type": "github"
      }
    },
    "systems_2": {
      "locked": {
        "lastModified": 1681028828,
        "narHash": "sha256-Vy1rq5AaRuLzOxct8nz4T6wlgyUR7zLU309k9mBC768=",
        "owner": "nix-systems",
        "repo": "default",
        "rev": "da67096a3b9bf56a91d16901293e51ba5b49a27e",
        "type": "github"
      },
      "original": {
        "owner": "nix-systems",
        "repo": "default",
        "type": "github"
      }
    },
    "systems_3": {
      "locked": {
        "lastModified": 1681028828,
        "narHash": "sha256-Vy1rq5AaRuLzOxct8nz4T6wlgyUR7zLU309k9mBC768=",
        "owner": "nix-systems",
        "repo": "default",
        "rev": "da67096a3b9bf56a91d16901293e51ba5b49a27e",
        "type": "github"
      },
      "original": {
        "owner": "nix-systems",
        "repo": "default",
        "type": "github"
      }
    },
    "treefmt-nix": {
      "inputs": {
        "nixpkgs": [
          "poetry2nix",
          "nixpkgs"
        ]
      },
      "locked": {
        "lastModified": 1730120726,
        "narHash": "sha256-LqHYIxMrl/1p3/kvm2ir925tZ8DkI0KA10djk8wecSk=",
        "owner": "numtide",
        "repo": "treefmt-nix",
        "rev": "9ef337e492a5555d8e17a51c911ff1f02635be15",
        "type": "github"
      },
      "original": {
        "owner": "numtide",
        "repo": "treefmt-nix",
        "type": "github"
      }
    }
  },
  "root": "root",
  "version": 7
}



================================================
FILE: flake.nix
================================================
{
  description = "Nix flakes";

  inputs = {
    nixpkgs.url = "nixpkgs/nixpkgs-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    poetry2nix = { url = "github:nix-community/poetry2nix"; };
  };

  outputs = { self, nixpkgs, flake-utils, poetry2nix }:
    flake-utils.lib.eachDefaultSystem (system: 
      let
        pkgs = import nixpkgs {
          inherit system;
        };
        inherit (poetry2nix.lib.mkPoetry2Nix { inherit pkgs; })
          mkPoetryEnv mkPoetryApplication defaultPoetryOverrides;
        poetryArgs = {
          python = pkgs.python312;
          projectDir = ./.;
          preferWheels = true;
          overrides = defaultPoetryOverrides.extend (final: prev: {
            click = prev.click.overridePythonAttrs (old: {
              buildInputs = (old.buildInputs or [ ]) ++ [ prev.flit-scm ];
            });
            rmc = prev.rmc.overridePythonAttrs (old: {
              buildInputs = (old.buildInputs or [ ]) ++ [ prev.poetry-core ];
            });
          });
        };
        pythonEnv = mkPoetryEnv (poetryArgs);
        rmcBin = mkPoetryApplication (poetryArgs);
      in
      {
        packages = {
          default = rmcBin;
        };
        devShells.default = pkgs.mkShell {
          buildInputs = [
	        pkgs.poetry
	        rmcBin
	        pythonEnv
          ];
        };
      });
}



================================================
FILE: LICENSE
================================================
MIT License

Copyright (c) 2023 Rick Lupton

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.



================================================
FILE: pyproject.toml
================================================
[tool.poetry]
name = "rmc"
version = "0.3.2-dev"
description = "Convert to/from v6 .rm files from the reMarkable tablet"
authors = ["Rick Lupton <mail@ricklupton.name>"]
license = "MIT"
readme = "README.md"

[tool.poetry.dependencies]
python = "^3.10"
rmscene = { git = "https://github.com/scrybbling-together/rmscene.git", branch = "main" }
click = "^8.0"
cairosvg = "^2.8.2"

[tool.poetry.dev-dependencies]
pytest = "^7.2.0"

[tool.poetry.scripts]
rmc = 'rmc.cli:cli'

[build-system]
requires = ["poetry-core>=1.0.0"]
build-backend = "poetry.core.masonry.api"



================================================
FILE: test.py
================================================
import glob
import os

this_branch = "test_output"
main_branch = "test_output_main"

for file in glob.glob(f"{this_branch}/**/*"):
    reference = file.replace(this_branch, main_branch)
    
    # if this is a md file with actual text content, compare the text
    if (file.endswith(".md") or file.endswith(".svg")) and os.path.exists(reference):
        with open(file, "r") as f1, open(reference, "r") as f2:
            if f1.read().strip() != f2.read().strip():
                print(f"{file} and {reference} are different")

print("All tests completed")


================================================
FILE: src/rmc/__init__.py
================================================
from .exporters.svg import tree_to_svg, rm_to_svg
from .exporters.pdf import rm_to_pdf



================================================
FILE: src/rmc/cli.py
================================================
"""CLI for converting rm files."""

import os
import sys
import io
from pathlib import Path
from contextlib import contextmanager
import click
from rmscene import read_tree, read_blocks, write_blocks, simple_text_document
from .exporters.svg import tree_to_svg
from .exporters.pdf import svg_to_pdf
from .exporters.markdown import print_text

import logging


@click.command
@click.version_option()
@click.option('-v', '--verbose', count=True)
@click.option("-f", "--from", "from_", metavar="FORMAT", help="Format to convert from (default: guess from filename)")
@click.option("-t", "--to", metavar="FORMAT", help="Format to convert to (default: guess from filename)")
@click.option("-o", "--output", type=click.Path(), help="Output filename (default: write to standard out)")
@click.argument("input", nargs=-1, type=click.Path(exists=True))
def cli(verbose, from_, to, output, input):
    """Convert to/from reMarkable v6 files.

    Available FORMATs are: `rm` (reMarkable file), `markdown`, `svg`, `pdf`,
    `blocks`, `blocks-data`.

    Formats `blocks` and `blocks-data` dump the internal structure of the `rm`
    file, with and without detailed data values respectively.

    """

    if verbose >= 2:
        logging.basicConfig(level=logging.DEBUG)
    elif verbose >= 1:
        logging.basicConfig(level=logging.INFO)
    else:
        logging.basicConfig(level=logging.WARNING)

    input = [Path(p) for p in input]
    if output is not None:
        output = Path(output)

    if from_ is None:
        if not input:
            raise click.UsageError("Must specify input filename or --from")
        from_ = guess_format(input[0])
    if to is None:
        if output is None:
            raise click.UsageError("Must specify --output or --to")
        to = guess_format(output)

    if from_ == "rm":
        with open_output(to, output) as fout:
            for fn in input:
                convert_rm(Path(fn), to, fout)
    elif from_ == "markdown":
        text = "".join(
            Path(fn).read_text() for fn in input
        )
        with open_output(to, output) as fout:
            convert_text(text, fout)
    else:
        raise click.UsageError("source format %s not implemented yet" % from_)


@contextmanager
def open_output(to, output):
    to_binary = to in ("pdf", "rm")
    if output is None:
        # Write to stdout
        if to_binary:
            with os.fdopen(sys.stdout.fileno(), "wb", closefd=False) as f:
                yield f
        else:
            yield sys.stdout
    else:
        with open(output, "w" + ("b" if to_binary else "t")) as f:
            yield f


def guess_format(p: Path):
    # XXX could be neater
    if p.suffix == ".rm":
        return "rm"
    if p.suffix == ".svg":
        return "svg"
    elif p.suffix == ".pdf":
        return "pdf"
    elif p.suffix == ".md" or p.suffix == ".markdown":
        return "markdown"
    else:
        return "blocks"


from rmscene import scene_items as si
def tree_structure(item):
    if isinstance(item, si.Group):
        return (
            item.node_id,
            (
                item.label.value,
                item.visible.value,
                (
                    item.anchor_id.value if item.anchor_id else None,
                    item.anchor_type.value if item.anchor_type else None,
                    item.anchor_threshold.value if item.anchor_threshold else None,
                    item.anchor_origin_x.value if item.anchor_origin_x else None,
                )
            ),
            [tree_structure(child) for child in item.children.values() if child],
        )
    else:
        return item


def convert_rm(filename: Path, to, fout):
    with open(filename, "rb") as f:
        if to == "blocks":
            pprint_blocks(f, fout)
        elif to == "blocks-data":
            pprint_blocks(f, fout, data=False)
        elif to == "tree":
            # Experimental dumping of tree structure
            pprint_tree(f, fout, data=True)
        elif to == "tree-data":
            # Experimental dumping of tree structure
            pprint_tree(f, fout, data=False)
        elif to == "markdown":
            print_text(f, fout)
        elif to == "svg":
            tree = read_tree(f)
            tree_to_svg(tree, fout)
        elif to == "pdf":
            buf = io.StringIO()
            tree = read_tree(f)
            tree_to_svg(tree, buf)
            buf.seek(0)
            svg_to_pdf(buf, fout)
        else:
            raise click.UsageError("Unknown format %s" % to)


def pprint_blocks(f, fout, data=True) -> None:
    import pprint
    depth = None if data else 1
    result = read_blocks(f)
    for el in result:
        print(file=fout)
        pprint.pprint(el, depth=depth, stream=fout)


def pprint_tree(f, fout, data=True) -> None:
    tree = read_tree(f)

    import pprint
    import re

    def pprint_Line(self, object, stream, indent, allowance, context, level):
        min_x = min(p.x for p in object.points)
        min_y = min(p.y for p in object.points)
        max_x = max(p.x for p in object.points)
        max_y = max(p.y for p in object.points)
        rep = re.sub(r"points=\[.*\]",
                     f"points=[({min_x: 4.0f},{min_y: 4.0f})-({max_x: 4.0f},{max_y: 4.0f})]",
                     repr(object))
        stream.write(rep)

    pprint.PrettyPrinter._dispatch[si.Line.__repr__] = pprint_Line

    depth = None if data else 1
    pprint.pprint(tree_structure(tree.root), stream=fout)
    pprint.pprint(tree_structure(tree.root_text), depth=depth, stream=fout)



def convert_text(text, fout):
    write_blocks(fout, simple_text_document(text))


if __name__ == "__main__":
    cli()



================================================
FILE: src/rmc/py.typed
================================================
[Empty file]


================================================
FILE: src/rmc/exporters/markdown.py
================================================
"""Export text content of rm files as Markdown."""

from rmscene import read_tree
from rmscene import scene_items as si

from rmscene.text import TextDocument


def print_text(f, fout):
    tree = read_tree(f)

    # Find out what anchor characters are used
    anchor_ids = set(collect_anchor_ids(tree.root))

    if tree.root_text:
        print_root_text(tree.root_text, fout, anchor_ids)

    JOIN_TOLERANCE = 2
    print("\n\n# Highlights", file=fout)
    last_pos = 0
    for item in tree.walk():
        if isinstance(item, si.GlyphRange):
            if item.start > last_pos + JOIN_TOLERANCE:
                print(file=fout)
            print(">", item.text, file=fout)
            last_pos = item.start + len(item.text)
    print(file=fout)


def print_root_text(root_text: si.Text, fout, anchor_ids):
    doc = TextDocument.from_scene_item(root_text)
    for p in doc.contents:
        annotated_line = annotate_anchor_ids(anchor_ids,
                                             str(p),
                                             [char_id for s in p.contents for char_id in s.i])
        if p.style.value == si.ParagraphStyle.BULLET:
            fout.write("- " + annotated_line)
        elif p.style.value == si.ParagraphStyle.BULLET2:
            fout.write("  + " + annotated_line)
        elif p.style.value == si.ParagraphStyle.BOLD:
            fout.write("> " + annotated_line)
        elif p.style.value == si.ParagraphStyle.HEADING:
            fout.write("# " + annotated_line)
        elif p.style.value == si.ParagraphStyle.PLAIN:
            fout.write(annotated_line)
        else:
            fout.write(("[unknown format %s] " % p.style.value) + annotated_line)


def annotate_anchor_ids(anchor_ids, line, ids):
    """Annotate appearances of `anchor_ids` in `line`."""
    result = ""
    for char, char_id in zip(line, ids):
        if char_id in anchor_ids:
            result += f"<<{char_id.part1},{char_id.part2}>>"
        result += char
    return result


def collect_anchor_ids(item):
    if isinstance(item, si.Group):
        if item.anchor_id is not None:
            yield item.anchor_id.value
        for child in item.children.values():
            yield from collect_anchor_ids(child)



================================================
FILE: src/rmc/exporters/pdf.py
================================================
"""Convert blocks to pdf file.

Code originally from https://github.com/lschwetlick/maxio through
https://github.com/chemag/maxio .
"""

import logging
from tempfile import NamedTemporaryFile
from cairosvg import svg2pdf

from .svg import rm_to_svg

_logger = logging.getLogger(__name__)


def rm_to_pdf(rm_path, pdf_path, debug=0):
    """Convert `rm_path` to PDF at `pdf_path`."""
    with NamedTemporaryFile(suffix=".svg") as f_temp:
        rm_to_svg(rm_path, f_temp.name)
        svg2pdf(url=f_temp.name, write_to=pdf_path)


def svg_to_pdf(svg_file, pdf_file):
    """Read svg data from `svg_file` and write PDF data to `pdf_file`."""
    svg2pdf(bytestring=svg_file.getvalue().encode('utf-8'), write_to=pdf_file.name)




================================================
FILE: src/rmc/exporters/svg.py
================================================
"""Convert blocks to svg file.

Code originally from https://github.com/lschwetlick/maxio through
https://github.com/chemag/maxio .
"""

import logging
import string
import typing as tp
from pathlib import Path

from rmscene.scene_items import Pen as PenType
from rmscene import CrdtId, SceneTree, read_tree
from rmscene import scene_items as si
from rmscene.text import TextDocument
from xml.etree.ElementTree import _escape_attrib

from .writing_tools import Pen

_logger = logging.getLogger(__name__)

SCREEN_WIDTH = 1404
SCREEN_HEIGHT = 1872
SCREEN_DPI = 226

SCALE = 72.0 / SCREEN_DPI

PAGE_WIDTH_PT = SCREEN_WIDTH * SCALE
PAGE_HEIGHT_PT = SCREEN_HEIGHT * SCALE
X_SHIFT = PAGE_WIDTH_PT // 2


def scale(screen_unit: float) -> float:
    return screen_unit * SCALE


# For now, at least, the xx and yy function are identical to scale
xx = scale
yy = scale

TEXT_TOP_Y = -88
LINE_HEIGHTS = {
    # Based on a rm file having 4 anchors based on the line height I was able to find a value of
    # 69.5, but decided on 70 (to keep integer values)
    si.ParagraphStyle.PLAIN: 70,
    si.ParagraphStyle.BULLET: 35,
    si.ParagraphStyle.BULLET2: 35,
    si.ParagraphStyle.BOLD: 70,
    si.ParagraphStyle.HEADING: 150,
    si.ParagraphStyle.CHECKBOX: 35,
    si.ParagraphStyle.CHECKBOX_CHECKED: 35,

    # There appears to be another format code (value 0) which is used when the
    # text starts far down the page, which case it has a negative offset (line
    # height) of about -20?
    #
    # Probably, actually, the line height should be added *after* the first
    # line, but there is still something a bit odd going on here.
}

SVG_HEADER = string.Template("""<?xml version="1.0" encoding="UTF-8"?>
<svg xmlns="http://www.w3.org/2000/svg" height="$height" width="$width" viewBox="$viewbox">""")


def rm_to_svg(rm_path, svg_path):
    """Convert `rm_path` to SVG at `svg_path`."""
    with open(rm_path, "rb") as infile, open(svg_path, "wt") as outfile:
        tree = read_tree(infile)
        tree_to_svg(tree, outfile)


def read_template_svg(template_path: Path) -> str:
    lines = template_path.read_text().splitlines()
    return "\n".join(lines[2:-2])


def tree_to_svg(tree: SceneTree, output, include_template: Path | None = None):
    """Convert Blocks to SVG."""

    # find the anchor pos for further use
    anchor_pos = build_anchor_pos(tree.root_text)
    _logger.debug("anchor_pos: %s", anchor_pos)

    # find the extremum along x and y
    x_min, x_max, y_min, y_max = get_bounding_box(tree.root, anchor_pos)
    width_pt = xx(x_max - x_min + 1)
    height_pt = yy(y_max - y_min + 1)
    _logger.debug("x_min, x_max, y_min, y_max: %.1f, %.1f, %.1f, %.1f ; scalded %.1f, %.1f, %.1f, %.1f",
                  x_min, x_max, y_min, y_max, xx(x_min), xx(x_max), yy(y_min), yy(y_max))

    # add svg header
    output.write(SVG_HEADER.substitute(width=width_pt,
                                       height=height_pt,
                                       viewbox=f"{xx(x_min)} {yy(y_min)} {width_pt} {height_pt}") + "\n")

    if include_template is not None:
        output.write(read_template_svg(include_template))
        output.write(f'\n\t<rect fill="url(#template)" x="{xx(x_min)}" y="{yy(y_min)}"'
                     f' width="{width_pt}" height="{height_pt}"/>\n')

    output.write(f'\t<g id="p1" style="display:inline">\n')

    if tree.root_text is not None:
        draw_text(tree.root_text, output)

    draw_group(tree.root, output, anchor_pos)

    # Closing page group
    output.write('\t</g>\n')
    # END notebook
    output.write('</svg>\n')


def build_anchor_pos(text: tp.Optional[si.Text]) -> tp.Dict[CrdtId, int]:
    """
    Find the anchor pos

    :param text: the root text of the remarkable file
    """
    # Special anchors adjusted based on pen_size_test.strokes.rm
    anchor_pos = {
        CrdtId(0, 281474976710654): 100,
        CrdtId(0, 281474976710655): 100,
    }

    if text is not None:
        # Save anchor from text
        doc = TextDocument.from_scene_item(text)
        ypos = text.pos_y + TEXT_TOP_Y
        for i, p in enumerate(doc.contents):
            anchor_pos[p.start_id] = ypos
            for subp in p.contents:
                for k in subp.i:
                    anchor_pos[k] = ypos  # TODO check these anchor are used
            ypos += LINE_HEIGHTS.get(p.style.value, 70)

    return anchor_pos


def get_anchor(item: si.Group, anchor_pos):
    anchor_x = 0.0
    anchor_y = 0.0
    if item.anchor_id is not None:
        assert item.anchor_origin_x is not None
        anchor_x = item.anchor_origin_x.value
        if item.anchor_id.value in anchor_pos:
            anchor_y = anchor_pos[item.anchor_id.value]
            _logger.debug("Group anchor: %s -> y=%.1f (scalded y=%.1f)",
                          item.anchor_id.value,
                          anchor_y,
                          yy(anchor_y))
        else:
            _logger.warning("Group anchor: %s is unknown!", item.anchor_id.value)

    return anchor_x, anchor_y


def get_bounding_box(item: si.Group,
                     anchor_pos: tp.Dict[CrdtId, int],
                     default: tp.Tuple[int, int, int, int] = (- SCREEN_WIDTH // 2, SCREEN_WIDTH // 2, 0, SCREEN_HEIGHT)) \
        -> tp.Tuple[int, int, int, int]:
    """
    Get the bounding box of the given item.
    The minimum size is the default size of the screen.

    :return: x_min, x_max, y_min, y_max: the bounding box in screen units (need to be scalded using xx and yy functions)
    """
    x_min, x_max, y_min, y_max = default

    for child_id in item.children:
        child = item.children[child_id]
        if isinstance(child, si.Group):
            anchor_x, anchor_y = get_anchor(child, anchor_pos)
            x_min_t, x_max_t, y_min_t, y_max_t = get_bounding_box(child, anchor_pos, (0, 0, 0, 0))
            x_min = min(x_min, x_min_t + anchor_x)
            x_max = max(x_max, x_max_t + anchor_x)
            y_min = min(y_min, y_min_t + anchor_y)
            y_max = max(y_max, y_max_t + anchor_y)
        elif isinstance(child, si.Line):
            x_min = min([x_min] + [p.x for p in child.points])
            x_max = max([x_max] + [p.x for p in child.points])
            y_min = min([y_min] + [p.y for p in child.points])
            y_max = max([y_max] + [p.y for p in child.points])

    return x_min, x_max, y_min, y_max


def draw_group(item: si.Group, output, anchor_pos):
    anchor_x, anchor_y = get_anchor(item, anchor_pos)
    output.write(f'\t\t<g id="{item.node_id}" transform="translate({xx(anchor_x)}, {yy(anchor_y)})">\n')
    for child_id in item.children:
        child = item.children[child_id]
        _logger.debug("Group child: %s %s", child_id, type(child))
        if _logger.root.level == logging.DEBUG:
            output.write(f'\t\t<!-- child {child_id} {type(child)} -->\n')
        if isinstance(child, si.Group):
            draw_group(child, output, anchor_pos)
        elif isinstance(child, si.Line):
            draw_stroke(child, output)
    output.write(f'\t\t</g>\n')


def draw_stroke(item: si.Line, output):
    # print debug infos
    if _logger.root.level == logging.DEBUG:
        _logger.debug("Writing line: %s", item)
        output.write(f'\t\t\t<!-- Stroke tool: {item.tool.name} '
                     f'color: {item.color.name} thickness_scale: {item.thickness_scale} -->\n')

    # initiate the pen
    pen = Pen.create(item.tool.value, item.color.value, item.thickness_scale)

    last_xpos = -1.
    last_ypos = -1.
    last_segment_width = segment_width = 0
    # Iterate through the point to form a polyline
    for point_idx, point in enumerate(item.points):
        # align the original position
        xpos = point.x
        ypos = point.y
        if point_idx % pen.segment_length == 0:
            # if there was a previous segment, end it
            if point_idx > 0:
                output.write('"/>\n')

            segment_color = pen.get_segment_color(point.speed, point.direction, point.width, point.pressure,
                                                  last_segment_width)
            segment_width = pen.get_segment_width(point.speed, point.direction, point.width, point.pressure,
                                                  last_segment_width)
            segment_opacity = pen.get_segment_opacity(point.speed, point.direction, point.width, point.pressure,
                                                      last_segment_width)
            # create the next segment of the stroke
            output.write('\t\t\t<polyline ')
            output.write(f'style="fill:none; stroke:{segment_color}; '
                         f'stroke-width:{scale(segment_width):.3f}; opacity:{segment_opacity}" ')
            output.write(f'stroke-linecap="{pen.stroke_linecap}" ')
            output.write('points="')
            if point_idx > 0:
                # Join to previous segment
                output.write(f'{xx(last_xpos):.3f},{yy(last_ypos):.3f} ')
        # store the last position
        last_xpos = xpos
        last_ypos = ypos
        last_segment_width = segment_width

        # add current point
        output.write(f'{xx(xpos):.3f},{yy(ypos):.3f} ')

    # end stroke
    output.write('" />\n')


def draw_text(text: si.Text, output):
    output.write('\t\t<g class="root-text" style="display:inline">')

    # add some style to get readable text
    output.write('''
            <style>
                text.heading {
                    font: 14pt serif;
                }
                text.bold {
                    font: 8pt sans-serif bold;
                }
                text, text.plain {
                    font: 7pt sans-serif;
                }
            </style>
''')

    y_offset = TEXT_TOP_Y

    doc = TextDocument.from_scene_item(text)
    for p in doc.contents:
        y_offset += LINE_HEIGHTS.get(p.style.value, 70)

        xpos = text.pos_x
        ypos = text.pos_y + y_offset
        cls = p.style.value.name.lower()
        if str(p):
            # TODO: this doesn't take into account the CrdtStr.properties (font-weight/font-style)
            if _logger.root.level == logging.DEBUG:
                output.write(f'\t\t\t<!-- Text line char_id: {p.start_id} -->\n')
            output.write(f'\t\t\t<text x="{xx(xpos)}" y="{yy(ypos)}" class="{cls}">{_escape_attrib(str(p).strip())}</text>\n')
    output.write('\t\t</g>\n')



================================================
FILE: src/rmc/exporters/writing_tools.py
================================================
"""
Common code for writing tools.

Code originally from https://github.com/lschwetlick/maxio through https://github.com/chemag/maxio
"""

import logging
import math
import sys

from rmscene.scene_items import HARDCODED_COLORMAP, PenColor
from rmscene.scene_items import Pen as PenType

_logger = logging.getLogger(__name__)

# color_id to RGB conversion
# 1. we use "color_id" for a unique, proprietary ID for colors,
#   (see scene_stream.py):
RM_PALETTE = {
    # Color code can be obtained from extraMetadata in the .content file
    PenColor.BLACK: (0, 0, 0),
    PenColor.GRAY: (144, 144, 144),
    PenColor.WHITE: (255, 255, 255),
    PenColor.YELLOW: (251, 247, 25),
    PenColor.GREEN: (0, 255, 0),
    PenColor.PINK: (255, 192, 203),
    PenColor.BLUE: (78, 105, 201),
    PenColor.RED: (179, 62, 57),
    PenColor.GRAY_OVERLAP: (125, 125, 125),
    # TODO: color mechanism is broken, as HIGHLIGHT is used for all the
    # highlighter colors. It is better though to produce a highlight color
    # than to make rmc crash.
    # Note that similar single-color name issues happen for ballpoint pen and
    # paintbrusher (though color works for fineliner, caligraphy pen, and
    # marker).
    PenColor.HIGHLIGHT: (247, 232, 81),
    PenColor.GREEN_2: (161, 216, 125),
    PenColor.CYAN: (139, 208, 229),
    PenColor.MAGENTA: (183, 130, 205),
    PenColor.YELLOW_2: (247, 232, 81),
} | {v: k for k, v in HARDCODED_COLORMAP.items()}


def clamp(value):
    """
    Clamp value between 0 and 1.
    """
    return min(max(value, 0), 1)


class Pen:
    def __init__(self, name, base_width, base_color_id):
        self.base_width = base_width

        color = RM_PALETTE.get(base_color_id, (0, 0, 0))
        opacity = 1
        if len(color) == 4:
            *color, opacity = color
            opacity = opacity / 255.0

        self.base_color = color
        self.base_opacity = opacity
        self.name = name
        self.segment_length = 1000

        # initial stroke values
        self.stroke_linecap = "round"
        self.stroke_opacity = 1
        self.stroke_width = base_width
        self.stroke_color = base_color_id

    # note that the units of the points have had their units converted
    # in scene_stream.py
    # speed = d.read_float32() * 4
    # ---> replace speed with speed / 4 [input]
    # direction = 255 * d.read_float32() / (math.pi * 2)
    # ---> replace tilt with direction_to_tilt() [input]
    @classmethod
    def direction_to_tilt(cls, direction):
        return direction * (math.pi * 2) / 255

    # width = int(round(d.read_float32() * 4))
    # ---> replace width with width / 4 [input]
    # ---> replace width with 4 * width [output]
    # pressure = d.read_float32() * 255
    # ---> replace pressure with pressure / 255 [input]

    def get_segment_width(self, speed, direction, width, pressure, last_width):
        return self.base_width

    def get_segment_color(self, speed, direction, width, pressure, last_width):
        return "rgb" + str(tuple(self.base_color))

    def get_segment_opacity(self, speed, direction, width, pressure, last_width):
        return self.base_opacity

    @classmethod
    def create(cls, pen_nr, color_id, width):
        # Brush
        if pen_nr in (PenType.PAINTBRUSH_1, PenType.PAINTBRUSH_2):
            return Brush(width, color_id)
        # Calligraphy (spelling mistake in rmscene will eventually be fixed)
        elif pen_nr == PenType.CALIGRAPHY:
            return Calligraphy(width, color_id)
        # Marker
        elif pen_nr in (PenType.MARKER_1, PenType.MARKER_2):
            return Marker(width, color_id)
        # BallPoint
        elif pen_nr in (PenType.BALLPOINT_1, PenType.BALLPOINT_2):
            return Ballpoint(width, color_id)
        # Fineliner
        elif pen_nr in (PenType.FINELINER_1, PenType.FINELINER_2):
            return Fineliner(width, color_id)
        # Pencil
        elif pen_nr in (PenType.PENCIL_1, PenType.PENCIL_2):
            return Pencil(width, color_id)
        # Mechanical Pencil
        elif pen_nr in (PenType.MECHANICAL_PENCIL_1, PenType.MECHANICAL_PENCIL_2):
            return MechanicalPencil(width, color_id)
        # Highlighter
        elif pen_nr in (PenType.HIGHLIGHTER_1, PenType.HIGHLIGHTER_2):
            width = 25
            return Highlighter(width, color_id)
        elif pen_nr == PenType.SHADER:
            # TODO: check if this is correct
            width = 12
            return Shader(width, color_id)
        # Erase area
        elif pen_nr == PenType.ERASER_AREA:
            return EraseArea(width, color_id)
        # Eraser
        elif pen_nr == PenType.ERASER:
            color_id = 2
            return Eraser(width, color_id)
        raise Exception(f"Unknown pen_nr: {pen_nr}")


class Fineliner(Pen):
    def __init__(self, base_width, base_color_id):
        super().__init__("Fineliner", base_width * 1.8, base_color_id)


class Ballpoint(Pen):
    def __init__(self, base_width, base_color_id):
        super().__init__("Ballpoint", base_width, base_color_id)
        self.segment_length = 5

    def get_segment_width(self, speed, direction, width, pressure, last_width):
        segment_width = (0.5 + pressure / 255) + (width / 4) - 0.5 * ((speed / 4) / 50)
        return segment_width


class Marker(Pen):
    def __init__(self, base_width, base_color_id):
        super().__init__("Marker", base_width, base_color_id)
        self.segment_length = 3

    def get_segment_width(self, speed, direction, width, pressure, last_width):
        segment_width = 0.9 * (
            (width / 4) - 0.4 * self.direction_to_tilt(direction)
        ) + (0.1 * last_width)
        return segment_width


class Pencil(Pen):
    def __init__(self, base_width, base_color_id):
        super().__init__("Pencil", base_width, base_color_id)
        self.segment_length = 2

    def get_segment_width(self, speed, direction, width, pressure, last_width):
        segment_width = 0.7 * (
            (((0.8 * self.base_width) + (0.5 * pressure / 255)) * (width / 4))
            - (0.25 * self.direction_to_tilt(direction) ** 1.8)
            - (0.6 * (speed / 4) / 50)
        )
        # segment_width = 1.3*(((self.base_width * 0.4) * pressure) - 0.5 * ((self.direction_to_tilt(direction) ** 0.5)) + (0.5 * last_width))
        max_width = self.base_width * 10
        segment_width = segment_width if segment_width < max_width else max_width
        return segment_width

    def get_segment_opacity(self, speed, direction, width, pressure, last_width):
        segment_opacity = (0.1 * -((speed / 4) / 35)) + (1 * pressure / 255)
        segment_opacity = clamp(segment_opacity) - 0.1
        return segment_opacity


class MechanicalPencil(Pen):
    def __init__(self, base_width, base_color_id):
        super().__init__("Mechanical Pencil", base_width**2, base_color_id)
        self.base_opacity = 0.7


class Brush(Pen):
    def __init__(self, base_width, base_color_id):
        super().__init__("Brush", base_width, base_color_id)
        self.segment_length = 2
        self.stroke_linecap = "round"
        self.opacity = 1

    def get_segment_width(self, speed, direction, width, pressure, last_width):
        segment_width = 0.7 * (
            ((1 + (1.4 * pressure / 255)) * (width / 4))
            - (0.5 * self.direction_to_tilt(direction))
            - ((speed / 4) / 50)
        )  # + (0.2 * last_width)
        return segment_width


class Highlighter(Pen):
    def __init__(self, base_width, base_color_id):
        super().__init__("Highlighter", base_width, base_color_id)
        self.stroke_linecap = "square"
        self.base_opacity = 0.3
        self.stroke_opacity = 0.3


class Shader(Pen):
    def __init__(self, base_width, base_color_id):
        super().__init__("Shader", base_width, base_color_id)
        self.stroke_linecap = "round"
        # self.base_opacity = 0.1
        # self.stroke_opacity = 0.2
        self.name = "Shader"


class Eraser(Pen):
    def __init__(self, base_width, base_color_id):
        super().__init__("Eraser", base_width * 2, base_color_id)
        self.stroke_linecap = "square"


class EraseArea(Pen):
    def __init__(self, base_width, base_color_id):
        super().__init__("Erase Area", base_width, base_color_id)
        self.stroke_linecap = "square"
        self.base_opacity = 0


class Calligraphy(Pen):
    def __init__(self, base_width, base_color_id):
        super().__init__("Calligraphy", base_width, base_color_id)
        self.segment_length = 2

    def get_segment_width(self, speed, direction, width, pressure, last_width):
        segment_width = 0.9 * (
            ((1 + pressure / 255) * (width / 4))
            - 0.3 * self.direction_to_tilt(direction)
        ) + (0.1 * last_width)
        return segment_width



================================================
FILE: tests/rm/abcd.strokes.rm
================================================
[Binary file]


================================================
FILE: tests/rm/abcd.text.rm
================================================
[Binary file]


================================================
FILE: tests/rm/Bold_Heading_Bullet_Normal.rm
================================================
[Binary file]


================================================
FILE: tests/rm/dot.stroke.rm
================================================
[Binary file]


================================================
FILE: tests/rm/extended.stroke.rm
================================================
[Binary file]


================================================
FILE: tests/rm/fullpage.rm
================================================
[Binary file]


================================================
FILE: tests/rm/keyboard-checkboxes-and-bullets.rm
================================================
[Binary file]


================================================
FILE: tests/rm/layers.stroke.rm
================================================
[Binary file]


================================================
FILE: tests/rm/Lines_v2.rm
================================================
[Binary file]


================================================
FILE: tests/rm/Normal_A_stroke_2_layers.rm
================================================
[Binary file]


================================================
FILE: tests/rm/Normal_AB.rm
================================================
[Binary file]


================================================
FILE: tests/rm/page_limits.rm
================================================
[Binary file]


================================================
FILE: tests/rm/sentinel_bug.rm
================================================
[Binary file]


================================================
FILE: tests/rm/text_and_strokes.rm
================================================
[Binary file]


================================================
FILE: tests/rm/text_multiple_lines.rm
================================================
[Binary file]


================================================
FILE: .github/workflows/release.yml
================================================
# Based on https://packaging.python.org/en/latest/guides/publishing-package-distribution-releases-using-github-actions-ci-cd-workflows/

name: Publish Python distribution to PyPI

on: push

jobs:
  build:
    name: Build distribution
    runs-on: ubuntu-latest

    steps:
    - uses: actions/checkout@v4
    - name: Set up Python
      uses: actions/setup-python@v5
      with:
        python-version: "3.x"
    - name: Install pypa/build
      run: python3 -m pip install build --user
    - name: Build a binary wheel and a source tarball
      run: python3 -m build
    - name: Store the distribution packages
      uses: actions/upload-artifact@v4
      with:
        name: python-package-distributions
        path: dist/

  publish-to-pypi:
    name: Publish Python distribution to PyPI
    if: startsWith(github.ref, 'refs/tags/')  # only publish to PyPI on tag pushes
    needs:
    - build
    runs-on: ubuntu-latest
    environment:
      name: pypi
      url: https://pypi.org/p/rmc
    permissions:
      id-token: write  # IMPORTANT: mandatory for trusted publishing

    steps:
    - name: Download all the dists
      uses: actions/download-artifact@v4
      with:
        name: python-package-distributions
        path: dist/
    - name: Publish distribution 📦 to PyPI
      uses: pypa/gh-action-pypi-publish@release/v1

  github-release:
    name: Sign Python distribution and upload to Github Release
    needs:
    - publish-to-pypi
    runs-on: ubuntu-latest

    permissions:
      contents: write  # IMPORTANT: mandatory for making GitHub Releases
      id-token: write  # IMPORTANT: mandatory for sigstore

    steps:
    - name: Download all the dists
      uses: actions/download-artifact@v4
      with:
        name: python-package-distributions
        path: dist/
    - name: Sign the dists with Sigstore
      uses: sigstore/gh-action-sigstore-python@v3.0.0
      with:
        inputs: >-
          ./dist/*.tar.gz
          ./dist/*.whl
    - name: Create GitHub Release
      env:
        GITHUB_TOKEN: ${{ github.token }}
      run: >-
        gh release create
        '${{ github.ref_name }}'
        --repo '${{ github.repository }}'
        --notes ""
    - name: Upload artifact signatures to GitHub Release
      env:
        GITHUB_TOKEN: ${{ github.token }}
      # Upload to GitHub Release using the `gh` CLI.
      # `dist/` contains the built packages, and the
      # sigstore-produced signatures and certificates.
      run: >-
        gh release upload
        '${{ github.ref_name }}' dist/**
        --repo '${{ github.repository }}'

  publish-to-testpypi:
    name: Publish Python distribution to TestPyPI
    needs:
    - build
    runs-on: ubuntu-latest

    environment:
      name: testpypi
      url: https://test.pypi.org/p/rmc

    permissions:
      id-token: write  # IMPORTANT: mandatory for trusted publishing

    steps:
    - name: Download all the dists
      uses: actions/download-artifact@v4
      with:
        name: python-package-distributions
        path: dist/
    - name: Publish distribution 📦 to TestPyPI
      uses: pypa/gh-action-pypi-publish@release/v1
      with:
        repository-url: https://test.pypi.org/legacy/
        # Don't fail when this is another push to the same version. If we put
        # the git hash in the version string, this could be removed.
        skip-existing: true



================================================
FILE: .github/workflows/test.yml
================================================
# This workflow will install Python dependencies and run the tests

name: Test

on:
  push:
    branches: [ main, github_actions ]
  pull_request:
    branches: [ main ]

jobs:
  tests:
    runs-on: ${{ matrix.os }}
    strategy:
      matrix:
        python-version: ["3.10", "3.11"]
        os: [ubuntu-latest, windows-latest, macos-latest]
    steps:
    - uses: actions/checkout@v3
    - name: Install macOS dependencies
      if: matrix.os == 'macos-latest'
      run: |
        brew install cairo libxml2 libffi
        echo "PKG_CONFIG_PATH=/opt/homebrew/lib/pkgconfig:$PKG_CONFIG_PATH" >> $GITHUB_ENV
        echo "DYLD_LIBRARY_PATH=/opt/homebrew/lib:$DYLD_LIBRARY_PATH" >> $GITHUB_ENV
    - name: Install Windows dependencies
      if: matrix.os == 'windows-latest'
      run: |
        choco install msys2 --yes
        C:\tools\msys64\usr\bin\pacman -S --noconfirm mingw-w64-x86_64-cairo mingw-w64-x86_64-pkg-config
        echo "C:\tools\msys64\mingw64\bin" | Out-File -FilePath $env:GITHUB_PATH -Encoding utf8 -Append
    - name: Install poetry
      run: pipx install poetry
    - name: Setup Python ${{ matrix.python-version }}
      uses: actions/setup-python@v4
      with:
        python-version: ${{ matrix.python-version }}
        cache: 'poetry'
    - run: poetry install
    - run: poetry run bash convert_test_files.sh


