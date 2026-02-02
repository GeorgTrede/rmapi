================================================
FILE: README.md
================================================
# rmscene

Python library to read v6 files from reMarkable tables (software version 3).

In particular, this version introduces the ability to include text as well as drawn lines. Extracting this text is the original motivation to develop this library, but it also can read much of the other types of data in the reMarkable files.

To convert rm files to other formats, you can use [rmc](https://github.com/ricklupton/rmc), which combines this library with code for converting lines to SVG, PDF, and simple Markdown.

## Changelog

### Unreleased

### v0.7.0

Change in block properties:

- Some SceneInfo values are now optional ([#40](https://github.com/ricklupton/rmscene/issues/40))

New feature:

- Add support for `paper_size` field on some SceneInfo blocks

### v0.6.1

Fixes:

- Fix AssertionError when some ids are missing in a `CrdtSequence` ([#36](https://github.com/ricklupton/rmscene/pull/36))
- Fix ValueError when the node_id is missing in a `SceneGroupItemBlock` ([#16](https://github.com/ricklupton/rmscene/issues/16)) 
- Store any unparsed data in blocks as raw bytes to allow for round-trip saving of files written in a newer format than the parsing code knows about.

### v0.6.0

New features:

- Add support for new blocks: `0x0D` (SceneInfo) and `0x08` (SceneTombstoneItemBlock) ([#24](https://github.com/ricklupton/rmscene/pull/24/))
- Add support for `move_id` field on some SceneLineItems ([#24](https://github.com/ricklupton/rmscene/pull/24/))
- Add support for new pen types and colours ([#31](https://github.com/ricklupton/rmscene/pull/31))

### v0.5.0

Breaking changes:

- The `start` property of `GlyphRange` items is now optional
  ([#15](https://github.com/ricklupton/rmscene/pull/15/)).
- The representation of formatted text spans has changed. Rather than
  using nested structures like `BoldSpan` and `ItalicSpan`, the
  `CrdtStr` objects now have optional text properties like
  `font-weight` and `font-style`. This simplifies the parsing code and
  the resulting data structure.

New features:

- Improved error recovery. An error during parsing, or an unknown block type,
  results in an `UnreadableBlock` containing the data that could not be read, so
  that parsing of other blocks can continue.
- Compatible with new reMarkable software version 3.6 format for
  highlighted text
  ([#15](https://github.com/ricklupton/rmscene/pull/15/)).
- New methods `read_bool_optional` and similar of `TaggedBlockReader`
  which return a default value if no matching tagged value is present
  in the block.
  
Other changes and fixes:

- The `value` attribute of scene item blocks, which was not being used, has been
  removed.
- Check more carefully for sub-blocks
  ([#17](https://github.com/ricklupton/rmscene/issues/17#issuecomment-1701071477)).
- Type hints fixed for `expand_text_items`.

### v0.4.0

Breaking changes:

- Rename `scene_items.TextFormat` to `ParagraphStyle` to better describe its
  meaning, now that we have inline bold/italic text styles.
- Remove methods from `scene_items.Text` object; use `text.TextDocument`
  instead.
- Writer: experimental change to emulate different reMarkable software versions
  by passing `{"version": "3.2.2"}` options to `write_blocks`. This allows us to
  continue to test round-trip reading and writing of old test files as new data
  values are added. Replaces `"line_version"` option.
  
New features:

- Parse text formatting information (bold and italic) introduced in reMarkable
  software version 3.3.

Other changes:

- Allow empty text items and unknown text formats without throwing exceptions.
- When extra data is present in the file, log the unrecognised bytes at DEBUG
  logging level along with the call stack, to make it easier to figure out where
  the code needs to be modified to read new data.
- Parse new data values (with unknown meaning) in PageInfoBlock and
  MigrationInfoBlock.

### v0.3.0

- Introduce `CrdtSequence` type to handle the different places that CRDT
  sequences are used, not just for text.
- Introduce `scene_items` module with data structures representing the data,
  independently from the `Block`s used to serialize them to `.rm` files.
- Introduce a `SceneTree` structure which holds the `SceneItem`s in
  groups/layers.
- Move Text data from `RootTextBlock` to `scene_items.Text` class, which
  includes methods for extracting lines of text and formatting.
- Text lines now include the trailing newline character.
- Read `GlyphRange` scene items, representing highlighted text in PDFs.

### v0.2.0

- Try to be more robust to unexpected data introduced by newer reMarkable software versions.
- Only warn once if unknown data is present, rather than for every block.
- Small API change to return type of `read_block` and `read_subblock` methods.

### v0.1.0

- Initial release

## Acknowledgements

https://github.com/ddvk/reader helped a lot in figuring out the structure and meaning of the files.  [@adq](https://github.com/adq) discovered a means to get debug output (see [issue 25](https://github.com/ricklupton/rmscene/issues/25)) which is very helpful for understanding the format.

Contributors:
- [@Azeirah](https://github.com/Azeirah) -- code and reporting issues
- [@adq](https://github.com/adq) -- code and reporting issues
- [@dotlambda](https://github.com/dotlambda) -- packaging
- [@ChenghaoMou](https://github.com/ChenghaoMou)



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
name = "rmscene"
version = "0.7.0"
description = "Read v6 .rm files from the reMarkable tablet"
authors = ["Rick Lupton <mail@ricklupton.name>"]
license = "MIT"
repository = "https://github.com/ricklupton/rmscene"
readme = "README.md"

[tool.poetry.dependencies]
python = "^3.10"
packaging = "^23.0"

[tool.poetry.dev-dependencies]
pytest = "^7.2.0"
hypothesis = "^6.68.2"

[build-system]
requires = ["poetry-core>=1.0.0"]
build-backend = "poetry.core.masonry.api"



================================================
FILE: src/rmscene/__init__.py
================================================
from .tagged_block_common import *
from .tagged_block_reader import *
from .tagged_block_writer import *
from .scene_stream import *



================================================
FILE: src/rmscene/__main__.py
================================================
"""Experimental cli helpers."""

import sys
import argparse
from . import read_blocks


def parse_args(args):
    parser = argparse.ArgumentParser(prog="rmscene")
    parser.add_argument("file", type=argparse.FileType("rb"), help="filename to read")
    return parser.parse_args(args)


def pprint_file(args) -> None:
    import pprint

    result = read_blocks(args.file)
    for el in result:
        print()
        pprint.pprint(el)


if __name__ == "__main__":
    args = parse_args(sys.argv[1:])
    pprint_file(args)



================================================
FILE: src/rmscene/crdt_sequence.py
================================================
"""Data structure representing CRDT sequence.

"""

import logging
import typing as tp
from typing import Iterable
from collections import defaultdict
from dataclasses import dataclass

from .tagged_block_common import CrdtId

_logger = logging.getLogger(__name__)


# If the type constraint is for a CrdtSequenceItem[Superclass], then a
# CrdtSequenceItem[Subclass] would do, so it is covariant.

_T = tp.TypeVar("_T", covariant=True)


@dataclass
class CrdtSequenceItem(tp.Generic[_T]):
    item_id: CrdtId
    left_id: CrdtId
    right_id: CrdtId
    deleted_length: int
    value: _T


# As a mutable container, CrdtSequence is invariant.
_Ti = tp.TypeVar("_Ti", covariant=False)


class CrdtSequence(tp.Generic[_Ti]):
    """Ordered CRDT Sequence container.

    The Sequence contains `CrdtSequenceItem`s, each of which has an ID and
    left/right IDs establishing a partial order.

    Iterating through the `CrdtSequence` yields IDs following this order.

    """

    def __init__(self, items=None):
        if items is None:
            items = []
        self._items = {item.item_id: item for item in items}

    def __eq__(self, other):
        if isinstance(other, CrdtSequence):
            return self._items == other._items
        if isinstance(other, (list, tuple)):
            return self == CrdtSequence(other)
        raise NotImplemented

    def __repr__(self):
        return "CrdtSequence(%s)" % (", ".join(str(i) for i in self._items.values()))

    ## Access values, in order

    def __iter__(self) -> tp.Iterator[CrdtId]:
        """Return ids in order"""
        yield from toposort_items(self._items.values())

    def keys(self) -> list[CrdtId]:
        """Return CrdtIds in order."""
        return list(self)

    def values(self) -> list[_Ti]:
        """Return list of sorted values."""
        return [self[item_id] for item_id in self]

    def items(self) -> Iterable[tuple[CrdtId, _Ti]]:
        """Return list of sorted key, value pairs."""
        return [(item_id, self[item_id]) for item_id in self]

    def __getitem__(self, key: CrdtId) -> _Ti:
        """Return item with key"""
        return self._items[key].value

    ## Access SequenceItems

    def sequence_items(self) -> list[CrdtSequenceItem[_Ti]]:
        """Iterate through CrdtSequenceItems."""
        return list(self._items.values())

    ## Modify sequence

    def add(self, item: CrdtSequenceItem[_Ti]):
        if item.item_id in self._items:
            raise ValueError("Already have item %s" % item.item_id)
        self._items[item.item_id] = item


END_MARKER = CrdtId(0, 0)


def toposort_items(items: Iterable[CrdtSequenceItem]) -> Iterable[CrdtId]:
    """Sort SequenceItems based on left and right ids.

    Returns `CrdtId`s in the sorted order.

    """

    item_dict = {}
    for item in items:
        item_dict[item.item_id] = item
    if not item_dict:
        return  # nothing to do

    def _side_id(item, side):
        side_id = getattr(item, f"{side}_id")
        if side_id == END_MARKER or side_id not in item_dict:
            if side_id != END_MARKER:
                _logger.debug("Ignoring unknown %s_id %s of %s", side, side_id, item)
            return "__start" if side == "left" else "__end"
        else:
            return side_id

    # build dictionary: key "comes after" values
    data = defaultdict(set)
    for item in item_dict.values():
        left_id = _side_id(item, "left")
        right_id = _side_id(item, "right")
        data[item.item_id].add(left_id)
        data[right_id].add(item.item_id)

    # fill in sources not explicitly included
    sources_not_in_data = {dep for deps in data.values() for dep in deps} - {
        k for k in data.keys()
    }
    data.update({k: set() for k in sources_not_in_data})

    while True:
        next_items = {item for item, deps in data.items() if not deps}
        if next_items == {"__end"}:
            break
        assert next_items
        yield from sorted(k for k in next_items if k in item_dict)
        data = {
            item: (deps - next_items)
            for item, deps in data.items()
            if item not in next_items
        }

    if data != {"__end": set()}:
        raise ValueError("cyclic dependency")



================================================
FILE: src/rmscene/py.typed
================================================
[Empty file]


================================================
FILE: src/rmscene/scene_items.py
================================================
"""Data structures for the contents of a scene."""

import enum
import logging
import typing as tp
from dataclasses import dataclass, field

from .crdt_sequence import CrdtSequence
from .tagged_block_common import CrdtId, LwwValue
from .text import expand_text_items

_logger = logging.getLogger(__name__)


## Base class


@dataclass
class SceneItem:
    """Base class for items stored in scene tree."""


## Group


@dataclass
class Group(SceneItem):
    """A Group represents a group of nested items.

    Groups are used to represent layers.

    node_id is the id that this sub-tree is stored as a "SceneTreeBlock".

    children is a sequence of other SceneItems.

    `anchor_id` refers to a text character which provides the anchor y-position
    for this group. There are two values that seem to be special:
    - `0xfffffffffffe` seems to be used for lines right at the top of the page?
    - `0xffffffffffff` seems to be used for lines right at the bottom of the page?

    """

    node_id: CrdtId
    children: CrdtSequence[SceneItem] = field(default_factory=CrdtSequence)
    label: LwwValue[str] = LwwValue(CrdtId(0, 0), "")
    visible: LwwValue[bool] = LwwValue(CrdtId(0, 0), True)

    anchor_id: tp.Optional[LwwValue[CrdtId]] = None
    anchor_type: tp.Optional[LwwValue[int]] = None
    anchor_threshold: tp.Optional[LwwValue[float]] = None
    anchor_origin_x: tp.Optional[LwwValue[float]] = None


## Strokes


@enum.unique
class PenColor(enum.IntEnum):
    """
    Color index value.
    """

    # XXX list from remt pre-v6

    BLACK = 0
    GRAY = 1
    WHITE = 2

    YELLOW = 3
    GREEN = 4
    PINK = 5

    BLUE = 6
    RED = 7

    GRAY_OVERLAP = 8

    # All highlight colors share the same value.
    # This is a placeholder, see the colormap below for details.
    HIGHLIGHT = 9

    GREEN_2 = 10
    CYAN = 11
    MAGENTA = 12
    
    YELLOW_2 = 13

    # HIGHLIGHT enumerated
    HIGHLIGHT_YELLOW = 14
    HIGHLIGHT_BLUE = 15
    HIGHLIGHT_PINK = 16
    HIGHLIGHT_ORANGE = 17
    HIGHLIGHT_GREEN = 18
    HIGHLIGHT_GRAY = 19

    # SHADER enumerated
    SHADER_GRAY = 20
    SHADER_ORANGE = 21
    SHADER_MAGENTA = 22
    SHADER_BLUE = 23
    SHADER_RED = 24
    SHADER_GREEN = 25
    SHADER_YELLOW = 26
    SHADER_CYAN = 27

# colors hardcoded in rm files for highlight and shader
HARDCODED_COLORMAP = {
    (255, 237, 117, 255): PenColor.HIGHLIGHT_YELLOW,
    (190, 234, 254, 255): PenColor.HIGHLIGHT_BLUE,
    (242, 158, 255, 255): PenColor.HIGHLIGHT_PINK,
    (255, 195, 140, 255): PenColor.HIGHLIGHT_ORANGE,
    (172, 255, 133, 255): PenColor.HIGHLIGHT_GREEN,
    (199, 199, 198, 255): PenColor.HIGHLIGHT_GRAY,
    (33, 30, 28, 64): PenColor.SHADER_GRAY,
    (254, 178, 0, 115): PenColor.SHADER_ORANGE,
    (192, 127, 210, 128): PenColor.SHADER_MAGENTA,
    (48, 74, 224, 77): PenColor.SHADER_BLUE,
    (194, 49, 50, 102): PenColor.SHADER_RED,
    (145, 218, 113, 128): PenColor.SHADER_GREEN,
    (250, 231, 25, 115): PenColor.SHADER_YELLOW,
    (116, 210, 232, 102): PenColor.SHADER_CYAN,
}


@enum.unique
class Pen(enum.IntEnum):
    """
    Stroke pen id representing reMarkable tablet tools.

    Tool examples: ballpoint, fineliner, highlighter or eraser.
    """

    # XXX this list is from remt pre-v6

    BALLPOINT_1 = 2
    BALLPOINT_2 = 15
    CALIGRAPHY = 21
    ERASER = 6
    ERASER_AREA = 8
    FINELINER_1 = 4
    FINELINER_2 = 17
    HIGHLIGHTER_1 = 5
    HIGHLIGHTER_2 = 18
    MARKER_1 = 3
    MARKER_2 = 16
    MECHANICAL_PENCIL_1 = 7
    MECHANICAL_PENCIL_2 = 13
    PAINTBRUSH_1 = 0
    PAINTBRUSH_2 = 12
    PENCIL_1 = 1
    PENCIL_2 = 14
    SHADER = 23

    @classmethod
    def is_highlighter(cls, value: int) -> bool:
        return value in (cls.HIGHLIGHTER_1, cls.HIGHLIGHTER_2)


@dataclass
class Point:
    x: float
    y: float
    speed: int
    direction: int
    width: int
    pressure: int


@dataclass
class Line(SceneItem):
    color: PenColor
    tool: Pen
    points: list[Point]
    thickness_scale: float
    starting_length: float
    move_id: tp.Optional[CrdtId] = None


## Text


@enum.unique
class ParagraphStyle(enum.IntEnum):
    """
    Text paragraph style.
    """

    BASIC = 0
    PLAIN = 1
    HEADING = 2
    BOLD = 3
    BULLET = 4
    BULLET2 = 5
    CHECKBOX = 6
    CHECKBOX_CHECKED = 7


END_MARKER = CrdtId(0, 0)


@dataclass
class Text(SceneItem):
    """Block of text.

    `items` are a CRDT sequence of strings. The `item_id` for each string refers
    to its first character; subsequent characters implicitly have sequential
    ids.

    When formatting is present, some of `items` have a value of an integer
    formatting code instead of a string.

    `styles` are LWW values representing a mapping of character IDs to
    `ParagraphStyle` values. These formats apply to each line of text (until the
    next newline).

    `pos_x`, `pos_y` and `width` are dimensions for the text block.

    """

    items: CrdtSequence[str | int]
    styles: dict[CrdtId, LwwValue[ParagraphStyle]]
    pos_x: float
    pos_y: float
    width: float


## Glyph range


@dataclass
class Rectangle:
    x: float
    y: float
    w: float
    h: float


@dataclass
class GlyphRange(SceneItem):
    """Highlighted text

    `start` is only available in SceneGlyphItemBlock version=0, prior to ReMarkable v3.6

    `length` is the length of the text

    `text` is the highlighted text itself

    `color` represents the highlight color

    `rectangles` represent the locations of the highlight.
    """
    start: tp.Optional[int]
    length: int
    text: str
    color: PenColor
    rectangles: list[Rectangle]



================================================
FILE: src/rmscene/scene_stream.py
================================================
"""Read structure of reMarkable tablet lines format v6

With help from ddvk's v6 reader, and enum values from remt.

"""

from __future__ import annotations

import io
import logging
import math
import typing as tp
from abc import ABC, abstractmethod
from collections.abc import Iterable, Iterator
from dataclasses import KW_ONLY, dataclass, replace
from typing import Optional
from uuid import UUID, uuid4

from packaging.version import Version

from . import scene_items as si
from .crdt_sequence import CrdtSequence, CrdtSequenceItem
from .scene_tree import SceneTree
from .tagged_block_common import CrdtId, LwwValue, UnexpectedBlockError
from .tagged_block_reader import MainBlockInfo, TaggedBlockReader
from .tagged_block_writer import TaggedBlockWriter

_logger = logging.getLogger(__name__)


############################################################
# Top-level block types
############################################################


@dataclass
class Block(ABC):
    BLOCK_TYPE: tp.ClassVar

    # Store any unrecognised data we can't understand
    _: KW_ONLY
    extra_data: bytes = b""

    def version_info(self, writer: TaggedBlockWriter) -> tuple[int, int]:
        """Return (min_version, current_version) to use when writing."""
        return (1, 1)

    def get_block_type(self) -> int:
        """Return block type for this block.

        By default, returns the block's BLOCK_TYPE attribute, but this method
        can be overriden if a single block subclass can handle multiple block
        types.

        """
        return self.BLOCK_TYPE

    @classmethod
    def lookup(cls, block_type: int) -> tp.Optional[tp.Type[Block]]:
        if getattr(cls, "BLOCK_TYPE", None) == block_type:
            return cls
        for subclass in cls.__subclasses__():
            if match := subclass.lookup(block_type):
                return match
        return None

    @classmethod
    def read(self, reader: TaggedBlockReader) -> Optional[Block]:
        """
        Maybe parse a block from the reader stream.
        """
        with reader.read_block() as block_info:
            if block_info is None:
                return

            block_type = Block.lookup(block_info.block_type)
            if block_type:
                try:
                    block = block_type.from_stream(reader)
                except Exception as e:
                    _logger.warning("Error reading block: %s", e)
                    reader.data.data.seek(block_info.offset)
                    data = reader.data.read_bytes(block_info.size)
                    block = UnreadableBlock(str(e), data, block_info)
            else:
                msg = (
                    f"Unknown block type {block_info.block_type}. "
                    f"Skipping {block_info.size} bytes."
                )
                _logger.warning(msg)
                data = reader.data.read_bytes(block_info.size)
                block = UnreadableBlock(msg, data, block_info)

        # Keep any unparsed extra data
        block.extra_data = block_info.extra_data
        return block

    def write(self, writer: TaggedBlockWriter):
        """Write the block header and content to the stream."""
        min_version, current_version = self.version_info(writer)
        with writer.write_block(self.get_block_type(), min_version, current_version):
            self.to_stream(writer)
            # Write any leftover extra data that wasn't parsed
            writer.data.write_bytes(self.extra_data)

    @classmethod
    @abstractmethod
    def from_stream(cls, reader: TaggedBlockReader) -> Block:
        """Read content of block from stream."""
        raise NotImplementedError()

    @abstractmethod
    def to_stream(self, writer: TaggedBlockWriter):
        """Write content of block to stream."""
        raise NotImplementedError()


@dataclass
class UnreadableBlock(Block):
    """Represent a block which could not be read for some reason."""

    error: str
    data: bytes
    info: MainBlockInfo

    def get_block_type(self) -> int:
        return self.info.block_type

    @classmethod
    def from_stream(cls, reader: TaggedBlockReader) -> Block:
        raise NotImplementedError()

    def to_stream(self, writer: TaggedBlockWriter):
        writer.data.write_bytes(self.data)


@dataclass
class SceneInfo(Block):
    BLOCK_TYPE: tp.ClassVar = 0x0D

    def version_info(self, _) -> tuple[int, int]:
        """Return (min_version, current_version) to use when writing."""
        return (0, 1)

    current_layer: LwwValue[CrdtId]
    background_visible: tp.Optional[LwwValue[bool]]
    root_document_visible: tp.Optional[LwwValue[bool]]
    paper_size: tp.Optional[tuple[int, int]]

    @classmethod
    def from_stream(cls, stream: TaggedBlockReader) -> SceneInfo:
        current_layer = stream.read_lww_id(1)
        background_visible = (
            stream.read_lww_bool(2) if stream.bytes_remaining_in_block() > 0 else None
        )
        root_document_visible = (
            stream.read_lww_bool(3) if stream.bytes_remaining_in_block() > 0 else None
        )
        paper_size = (
            stream.read_int_pair(5) if stream.bytes_remaining_in_block() > 0 else None
        )

        return SceneInfo(
            current_layer=current_layer,
            background_visible=background_visible,
            root_document_visible=root_document_visible,
            paper_size=paper_size,
        )

    def to_stream(self, writer: TaggedBlockWriter):
        writer.write_lww_id(1, self.current_layer)
        if self.background_visible:
            writer.write_lww_bool(2, self.background_visible)
        if self.root_document_visible:
            writer.write_lww_bool(3, self.root_document_visible)
        if self.paper_size:
            writer.write_int_pair(5, self.paper_size)


@dataclass
class AuthorIdsBlock(Block):
    BLOCK_TYPE: tp.ClassVar = 0x09

    author_uuids: dict[int, UUID]

    @classmethod
    def from_stream(cls, stream: TaggedBlockReader) -> AuthorIdsBlock:
        _logger.debug("Reading %s", cls.__name__)
        num_subblocks = stream.data.read_varuint()
        author_ids = {}
        for _ in range(num_subblocks):
            with stream.read_subblock(0):
                uuid_length = stream.data.read_varuint()
                if uuid_length != 16:
                    raise ValueError("Expected UUID length to be 16 bytes")
                uuid = UUID(bytes_le=stream.data.read_bytes(uuid_length))
                author_id = stream.data.read_uint16()
                author_ids[author_id] = uuid
        return AuthorIdsBlock(author_ids)

    def to_stream(self, writer: TaggedBlockWriter):
        _logger.debug("Writing %s", type(self).__name__)
        num_subblocks = len(self.author_uuids)
        writer.data.write_varuint(num_subblocks)
        for author_id, uuid in self.author_uuids.items():
            with writer.write_subblock(0):
                writer.data.write_varuint(len(uuid.bytes_le))
                writer.data.write_bytes(uuid.bytes_le)
                writer.data.write_uint16(author_id)


@dataclass
class MigrationInfoBlock(Block):
    BLOCK_TYPE: tp.ClassVar = 0x00

    migration_id: CrdtId
    is_device: bool
    _unknown: bool = False

    @classmethod
    def from_stream(cls, stream: TaggedBlockReader) -> MigrationInfoBlock:
        "Parse migration info"
        _logger.debug("Reading %s", cls.__name__)
        migration_id = stream.read_id(1)
        is_device = stream.read_bool(2)
        if stream.bytes_remaining_in_block():
            unknown = stream.read_bool(3)
        else:
            unknown = False
        return MigrationInfoBlock(migration_id, is_device, unknown)

    def to_stream(self, writer: TaggedBlockWriter):
        _logger.debug("Writing %s", type(self).__name__)
        version = writer.options.get("version", Version("9.9.9"))
        writer.write_id(1, self.migration_id)
        writer.write_bool(2, self.is_device)
        if version >= Version("3.2.2"):
            writer.write_bool(3, self._unknown)


@dataclass
class TreeNodeBlock(Block):
    BLOCK_TYPE: tp.ClassVar = 0x02

    def version_info(self, writer: TaggedBlockWriter) -> tuple[int, int]:
        """Return (min_version, current_version) to use when writing."""
        version = writer.options.get("version", Version("9999"))
        # XXX this is a guess about which version this changed in
        return (1, 2) if (version >= Version("3.4")) else (1, 1)

    group: si.Group

    @classmethod
    def from_stream(cls, stream: TaggedBlockReader) -> TreeNodeBlock:
        "Parse tree node block."
        _logger.debug("Reading %s", cls.__name__)

        group = si.Group(
            node_id=stream.read_id(1),
            label=stream.read_lww_string(2),
            visible=stream.read_lww_bool(3),
        )

        # XXX this may need to be generalised for other examples
        if stream.bytes_remaining_in_block() > 0:
            group.anchor_id = stream.read_lww_id(7)
            group.anchor_type = stream.read_lww_byte(8)
            group.anchor_threshold = stream.read_lww_float(9)
            group.anchor_origin_x = stream.read_lww_float(10)

        return cls(group)

    def to_stream(self, writer: TaggedBlockWriter):
        _logger.debug("Writing %s", type(self).__name__)
        group = self.group
        writer.write_id(1, group.node_id)
        writer.write_lww_string(2, group.label)
        writer.write_lww_bool(3, group.visible)
        if group.anchor_id is not None:
            # FIXME group together in an anchor type?
            assert (
                group.anchor_type is not None
                and group.anchor_threshold is not None
                and group.anchor_origin_x is not None
            )
            writer.write_lww_id(7, group.anchor_id)
            writer.write_lww_byte(8, group.anchor_type)
            writer.write_lww_float(9, group.anchor_threshold)
            writer.write_lww_float(10, group.anchor_origin_x)


@dataclass
class PageInfoBlock(Block):
    BLOCK_TYPE: tp.ClassVar = 0x0A

    def version_info(self, _) -> tuple[int, int]:
        """Return (min_version, current_version) to use when writing."""
        return (0, 1)

    loads_count: int
    merges_count: int
    text_chars_count: int
    text_lines_count: int
    type_folio_use_count: int = 0

    @classmethod
    def from_stream(cls, stream: TaggedBlockReader) -> PageInfoBlock:
        "Parsehttps://github.com/ChenghaoMou/rmscene/tree/feature/highlight_color page info block"
        _logger.debug("Reading %s", cls.__name__)
        info = PageInfoBlock(
            loads_count=stream.read_int(1),
            merges_count=stream.read_int(2),
            text_chars_count=stream.read_int(3),
            text_lines_count=stream.read_int(4),
        )
        if stream.bytes_remaining_in_block():
            info.type_folio_use_count = stream.read_int(5)
        return info

    def to_stream(self, writer: TaggedBlockWriter):
        _logger.debug("Writing %s", type(self).__name__)
        writer.write_int(1, self.loads_count)
        writer.write_int(2, self.merges_count)
        writer.write_int(3, self.text_chars_count)
        writer.write_int(4, self.text_lines_count)
        version = writer.options.get("version", Version("9999"))
        if version >= Version("3.2.2"):
            writer.write_int(5, self.type_folio_use_count)


@dataclass
class SceneTreeBlock(Block):
    BLOCK_TYPE: tp.ClassVar = 0x01

    # XXX not sure what the difference is
    tree_id: CrdtId
    node_id: CrdtId
    is_update: bool
    parent_id: CrdtId

    @classmethod
    def from_stream(cls, stream: TaggedBlockReader) -> SceneTreeBlock:
        "Parse scene tree block"
        _logger.debug("Reading %s", cls.__name__)

        # XXX not sure what the difference is. This "tree_id" is used as the
        # plain "Id" in the SceneTree.NodeMap in ddvk's reader. If the parent_id
        # is equal to the root_id (1, 1), this node represents a layer.
        tree_id = stream.read_id(1)
        node_id = stream.read_id(2)
        is_update = stream.read_bool(3)
        with stream.read_subblock(4):
            parent_id = stream.read_id(1)
            # XXX can there sometimes be something else here?

        return SceneTreeBlock(tree_id, node_id, is_update, parent_id)

    def to_stream(self, writer: TaggedBlockWriter):
        _logger.debug("Writing %s", type(self).__name__)
        writer.write_id(1, self.tree_id)
        writer.write_id(2, self.node_id)
        writer.write_bool(3, self.is_update)
        with writer.write_subblock(4):
            writer.write_id(1, self.parent_id)


def point_from_stream(stream: TaggedBlockReader, version: int = 2) -> si.Point:
    if version not in (1, 2):
        raise ValueError("Unknown version %s" % version)
    d = stream.data
    x = d.read_float32()
    y = d.read_float32()
    if version == 1:
        # calculation based on ddvk's reader
        # XXX removed rounding so that can round-trip correctly?
        speed = d.read_float32() * 4
        # speed = int(round(d.read_float32() * 4))
        direction = 255 * d.read_float32() / (math.pi * 2)
        # direction = int(round(255 * d.read_float32() / (math.pi * 2)))
        width = int(round(d.read_float32() * 4))
        pressure = d.read_float32() * 255
        # pressure = int(round(d.read_float32() * 255))
    else:
        speed = d.read_uint16()
        width = d.read_uint16()
        direction = d.read_uint8()
        pressure = d.read_uint8()
    return si.Point(x, y, speed, direction, width, pressure)


def point_serialized_size(version: int = 2) -> int:
    if version == 1:
        return 0x18
    elif version == 2:
        return 0x0E
    else:
        raise ValueError("Unknown version %s" % version)


def point_to_stream(point: si.Point, writer: TaggedBlockWriter, version: int = 2):
    if version not in (1, 2):
        raise ValueError("Unknown version %s" % version)
    d = writer.data
    d.write_float32(point.x)
    d.write_float32(point.y)
    _logger.debug("Writing Point v%d: %s", version, point)
    if version == 1:
        # calculation based on ddvk's reader
        d.write_float32(point.speed / 4)
        d.write_float32(point.direction * (2 * math.pi) / 255)
        d.write_float32(point.width / 4)
        d.write_float32(point.pressure / 255)
    else:
        d.write_uint16(point.speed)
        d.write_uint16(point.width)
        d.write_uint8(point.direction)
        d.write_uint8(point.pressure)


def line_from_stream(stream: TaggedBlockReader, version: int = 2) -> si.Line:
    _logger.debug("Reading Line version %d", version)
    tool_id = stream.read_int(1)
    tool = si.Pen(tool_id)
    color_id = stream.read_int(2)
    color = si.PenColor(color_id)
    thickness_scale = stream.read_double(3)
    starting_length = stream.read_float(4)
    with stream.read_subblock(5) as block_info:
        data_length = block_info.size
        point_size = point_serialized_size(version)
        if data_length % point_size != 0:
            raise ValueError(
                "Point data size mismatch: %d is not multiple of point_size"
                % data_length
            )
        num_points = data_length // point_size
        points = [point_from_stream(stream, version=version) for _ in range(num_points)]

    # XXX unused
    timestamp = stream.read_id(6)

    if stream.bytes_remaining_in_block() >= 3:
        try:
            move_id = stream.read_id(7)
        except UnexpectedBlockError as _:
            move_id = None
    else:
        move_id = None

    # This is for the color information (highlight & shader)
    if stream.bytes_remaining_in_block() >= 6:
        # not sure what this is, seems fixed x84x01
        unk = stream.data.read_bytes(2)
        b = stream.data.read_uint8()
        g = stream.data.read_uint8()
        r = stream.data.read_uint8()
        a = stream.data.read_uint8()
        rgba = (r, g, b, a)

        if unk != b"\x84\x01" or rgba not in si.HARDCODED_COLORMAP:
            _logger.warning(f"Unhandled color {rgba} with prefix {unk}")
            stream.data.data.seek(-6, io.SEEK_CUR)
        else:
            color = si.HARDCODED_COLORMAP[rgba]

    return si.Line(color, tool, points, thickness_scale, starting_length, move_id)


def line_to_stream(line: si.Line, writer: TaggedBlockWriter, version: int = 2):
    _logger.debug("Writing Line version %d", version)
    writer.write_int(1, line.tool)
    is_highlight_color = line.color in si.HARDCODED_COLORMAP.values()
    if is_highlight_color:
        writer.write_int(2, si.PenColor.HIGHLIGHT)
    else:
        writer.write_int(2, line.color)
    writer.write_double(3, line.thickness_scale)
    writer.write_float(4, line.starting_length)
    with writer.write_subblock(5):
        for point in line.points:
            point_to_stream(point, writer, version)

    # XXX didn't save
    timestamp = CrdtId(0, 1)
    writer.write_id(6, timestamp)
    if line.move_id is not None:
        writer.write_id(7, line.move_id)

    if is_highlight_color:
        rgba = [
            key for key, value in si.HARDCODED_COLORMAP.items() if value == line.color
        ][0]
        writer.data.write_bytes(b"\x84\x01")
        writer.data.write_uint8(rgba[2])
        writer.data.write_uint8(rgba[1])
        writer.data.write_uint8(rgba[0])
        writer.data.write_uint8(rgba[3])


@dataclass
class SceneItemBlock(Block):
    parent_id: CrdtId
    item: CrdtSequenceItem
    extra_value_data: bytes = b""

    ITEM_TYPE: tp.ClassVar[int] = 0

    @classmethod
    def from_stream(cls, stream: TaggedBlockReader) -> SceneItemBlock:
        "Group item block?"
        _logger.debug("Reading %s", cls.__name__)

        assert stream.current_block
        block_type = stream.current_block.block_type
        if block_type == SceneGlyphItemBlock.BLOCK_TYPE:
            subclass = SceneGlyphItemBlock
        elif block_type == SceneGroupItemBlock.BLOCK_TYPE:
            subclass = SceneGroupItemBlock
        elif block_type == SceneLineItemBlock.BLOCK_TYPE:
            subclass = SceneLineItemBlock
        elif block_type == SceneTextItemBlock.BLOCK_TYPE:
            subclass = SceneTextItemBlock
        elif block_type == SceneTombstoneItemBlock.BLOCK_TYPE:
            subclass = SceneTombstoneItemBlock
        else:
            raise ValueError(
                "unknown scene type %d in %s" % (block_type, stream.current_block)
            )

        parent_id = stream.read_id(1)
        item_id = stream.read_id(2)
        left_id = stream.read_id(3)
        right_id = stream.read_id(4)
        deleted_length = stream.read_int(5)

        if stream.has_subblock(6):
            with stream.read_subblock(6) as block_info:
                item_type = stream.data.read_uint8()
                assert item_type == subclass.ITEM_TYPE
                value = subclass.value_from_stream(stream)
            # Keep known extra data from within the value subblock

            extra_value_data = block_info.extra_data
        else:
            value = None
            extra_value_data = b""

        return subclass(
            parent_id,
            CrdtSequenceItem(item_id, left_id, right_id, deleted_length, value),
            extra_value_data=extra_value_data,
        )

    def to_stream(self, writer: TaggedBlockWriter):
        _logger.debug("Writing %s", type(self).__name__)
        writer.write_id(1, self.parent_id)
        writer.write_id(2, self.item.item_id)
        writer.write_id(3, self.item.left_id)
        writer.write_id(4, self.item.right_id)
        writer.write_int(5, self.item.deleted_length)

        if self.item.value is not None:
            with writer.write_subblock(6):
                writer.data.write_uint8(self.ITEM_TYPE)
                self.value_to_stream(writer, self.item.value)

                writer.data.write_bytes(self.extra_value_data)

    @classmethod
    @abstractmethod
    def value_from_stream(cls, reader: TaggedBlockReader) -> tp.Any:
        """Read the specific content of this block"""
        raise NotImplementedError()

    @abstractmethod
    def value_to_stream(self, writer: TaggedBlockWriter, value: tp.Any):
        """Write the specific content of this block"""
        raise NotImplementedError()


# These share the same structure so can share the same implementation?


def glyph_range_from_stream(stream: TaggedBlockReader) -> si.GlyphRange:
    # Since reMarkable version 3.6, the start and length are optional
    start = stream.read_int_optional(2)
    length = stream.read_int_optional(3)

    color_id = stream.read_int(4)
    color = si.PenColor(color_id)
    text = stream.read_string(5)

    if length is None:
        length = len(text)

    # Note: the decoded text length is not always the same as the length in the
    # glyph range...
    if len(text) != length:
        _logger.debug(
            "GlyphRange text length %d != length value %d: %r",
            len(text),
            length,
            text,
        )

    with stream.read_subblock(6):
        num_rects = stream.data.read_varuint()
        rectangles = [
            si.Rectangle(*[stream.data.read_float64() for _ in range(4)])
            for _ in range(num_rects)
        ]

    return si.GlyphRange(start, length, text, color, rectangles)


def glyph_range_to_stream(stream: TaggedBlockWriter, item: si.GlyphRange):
    if item.start is not None:
        stream.write_int(2, item.start)
        stream.write_int(3, item.length)
    stream.write_int(4, item.color)
    stream.write_string(5, item.text)
    with stream.write_subblock(6):
        stream.data.write_varuint(len(item.rectangles))
        for rect in item.rectangles:
            stream.data.write_float64(rect.x)
            stream.data.write_float64(rect.y)
            stream.data.write_float64(rect.w)
            stream.data.write_float64(rect.h)


class SceneTombstoneItemBlock(SceneItemBlock):
    BLOCK_TYPE: tp.ClassVar = 0x08

    @classmethod
    def value_from_stream(cls, reader: TaggedBlockReader):
        pass

    def value_to_stream(self, writer: TaggedBlockWriter, value):
        pass


class SceneGlyphItemBlock(SceneItemBlock):
    BLOCK_TYPE: tp.ClassVar = 0x03
    ITEM_TYPE: tp.ClassVar = 0x01

    @classmethod
    def value_from_stream(cls, reader: TaggedBlockReader) -> si.GlyphRange:
        value = glyph_range_from_stream(reader)
        return value

    def value_to_stream(self, writer: TaggedBlockWriter, value):
        glyph_range_to_stream(writer, value)


class SceneGroupItemBlock(SceneItemBlock):
    BLOCK_TYPE: tp.ClassVar = 0x04
    ITEM_TYPE: tp.ClassVar = 0x02

    @classmethod
    def value_from_stream(cls, reader: TaggedBlockReader) -> CrdtId:
        # XXX don't know what this means
        value = reader.read_id(2)
        return value

    def value_to_stream(self, writer: TaggedBlockWriter, value: CrdtId):
        writer.write_id(2, value)


class SceneLineItemBlock(SceneItemBlock):
    BLOCK_TYPE: tp.ClassVar = 0x05
    ITEM_TYPE: tp.ClassVar = 0x03

    def version_info(self, writer: TaggedBlockWriter) -> tuple[int, int]:
        """Return (min_version, current_version) to use when writing."""
        version = writer.options.get("version", Version("9999"))
        return (2, 2) if (version > Version("3.0")) else (1, 1)

    @classmethod
    def value_from_stream(cls, reader: TaggedBlockReader) -> si.Line:
        assert reader.current_block is not None
        version = reader.current_block.current_version
        value = line_from_stream(reader, version)
        return value

    def value_to_stream(self, writer: TaggedBlockWriter, value: si.Line):
        # XXX make sure this version ends up in block header
        version = writer.options.get("version", Version("9999"))
        line_version = 2 if (version > Version("3.0")) else 1
        line_to_stream(value, writer, version=line_version)


# XXX missing "PathItemBlock"? with ITEM_TYPE 0x04


class SceneTextItemBlock(SceneItemBlock):
    BLOCK_TYPE: tp.ClassVar = 0x06
    ITEM_TYPE: tp.ClassVar = 0x05

    @classmethod
    def value_from_stream(cls, reader: TaggedBlockReader) -> tp.Any:
        return None

    def value_to_stream(self, writer: TaggedBlockWriter, value):
        pass


def text_item_from_stream(stream: TaggedBlockReader) -> CrdtSequenceItem[str | int]:
    with stream.read_subblock(0):
        item_id = stream.read_id(2)
        left_id = stream.read_id(3)
        right_id = stream.read_id(4)
        deleted_length = stream.read_int(5)

        if stream.has_subblock(6):
            text, fmt = stream.read_string_with_format(6)
            # It seems that formats are stored on empty strings, so it's one or the other
            if fmt is not None:
                if text:
                    _logger.error(
                        "Unhandled combined text and format: %s, %s", text, fmt
                    )
                value = fmt
            else:
                value = text
        else:
            value = ""

    return CrdtSequenceItem(item_id, left_id, right_id, deleted_length, value)


def text_item_to_stream(item: CrdtSequenceItem[str | int], writer: TaggedBlockWriter):
    with writer.write_subblock(0):
        writer.write_id(2, item.item_id)
        writer.write_id(3, item.left_id)
        writer.write_id(4, item.right_id)
        writer.write_int(5, item.deleted_length)

        if item.value:
            if isinstance(item.value, str):
                writer.write_string(6, item.value)
            elif isinstance(item.value, int):
                writer.write_string_with_format(6, "", item.value)


def text_format_from_stream(
    stream: TaggedBlockReader,
) -> tuple[CrdtId, LwwValue[si.ParagraphStyle]]:
    # These are character ids, but not with an initial tag like other ids have.
    char_id = stream.data.read_crdt_id()

    # This seems to be the item ID for this format data? It doesn't appear
    # elsewhere in the file. Sometimes coincides with a character id but I don't
    # think it is referring to it.
    timestamp = stream.read_id(1)

    with stream.read_subblock(2):
        # XXX not sure what this is format?
        c = stream.data.read_uint8()
        assert c == 17
        format_code = stream.data.read_uint8()
        try:
            format_type = si.ParagraphStyle(format_code)
        except ValueError:
            _logger.warning("Unrecognised text format code %d.", format_code)
            _logger.debug(
                "Unrecognised text format code %d at position %d.",
                format_code,
                stream.data.tell(),
            )
            format_type = si.ParagraphStyle.PLAIN  # fallback

    return (char_id, LwwValue(timestamp, format_type))


def text_format_to_stream(
    char_id: CrdtId, value: LwwValue[si.ParagraphStyle], writer: TaggedBlockWriter
):
    format_type = value.value

    writer.data.write_crdt_id(char_id)
    writer.write_id(1, value.timestamp)
    with writer.write_subblock(2):
        # XXX not sure what this is format?
        c = 17
        writer.data.write_uint8(c)
        writer.data.write_uint8(format_type)


@dataclass
class RootTextBlock(Block):
    BLOCK_TYPE: tp.ClassVar = 0x07

    block_id: CrdtId
    value: si.Text

    @classmethod
    def from_stream(cls, stream: TaggedBlockReader) -> RootTextBlock:
        "Parse root text block."
        _logger.debug("Reading %s", cls.__name__)

        block_id = stream.read_id(1)
        assert block_id == CrdtId(0, 0)

        with stream.read_subblock(2):
            # Text items
            with stream.read_subblock(1):
                with stream.read_subblock(1):
                    num_subblocks = stream.data.read_varuint()
                    text_items = [
                        text_item_from_stream(stream) for _ in range(num_subblocks)
                    ]

            # Formatting
            with stream.read_subblock(2):
                with stream.read_subblock(1):
                    num_subblocks = stream.data.read_varuint()
                    text_formats = dict(
                        text_format_from_stream(stream) for _ in range(num_subblocks)
                    )

        # Last section
        with stream.read_subblock(3):
            # "pos_x" and "pos_y" from ddvk? Gives negative number -- possibly could
            # be bounding box?
            pos_x = stream.data.read_float64()
            pos_y = stream.data.read_float64()

        # "width" from ddvk
        width = stream.read_float(4)

        value = si.Text(
            items=CrdtSequence(text_items),
            styles=text_formats,
            pos_x=pos_x,
            pos_y=pos_y,
            width=width,
        )
        return RootTextBlock(block_id, value)

    def to_stream(self, writer: TaggedBlockWriter):
        _logger.debug("Writing %s", type(self).__name__)
        writer.write_id(1, self.block_id)

        with writer.write_subblock(2):
            # Text items
            text_items = self.value.items.sequence_items()
            with writer.write_subblock(1):
                with writer.write_subblock(1):
                    writer.data.write_varuint(len(text_items))
                    for item in text_items:
                        text_item_to_stream(item, writer)

            # Formatting
            text_formats = self.value.styles
            with writer.write_subblock(2):
                with writer.write_subblock(1):
                    writer.data.write_varuint(len(text_formats))
                    for key, item in text_formats.items():
                        text_format_to_stream(key, item, writer)

        # Last section
        with writer.write_subblock(3):
            writer.data.write_float64(self.value.pos_x)
            writer.data.write_float64(self.value.pos_y)

        # "width" from ddvk
        writer.write_float(4, self.value.width)


## Functions to read and write streams of blocks


def _read_blocks(stream: TaggedBlockReader) -> Iterator[Block]:
    """
    Parse blocks from reMarkable v6 file.
    """
    while True:
        maybe_block = Block.read(stream)
        if maybe_block:
            yield maybe_block
        else:
            # no more blocks
            return


def read_blocks(data: tp.BinaryIO) -> Iterator[Block]:
    """
    Parse reMarkable file and return iterator of document items.

    :param data: reMarkable file data.
    """
    stream = TaggedBlockReader(data)
    stream.read_header()
    yield from _read_blocks(stream)


def write_blocks(
    data: tp.BinaryIO, blocks: Iterable[Block], options: tp.Optional[dict] = None
):
    """
    Write blocks to file.
    """
    if options is not None and "version" in options:
        options["version"] = Version(options["version"])
    stream = TaggedBlockWriter(data, options=options)
    stream.write_header()
    for block in blocks:
        block.write(stream)


def build_tree(tree: SceneTree, blocks: Iterable[Block]):
    """Read `blocks` and add contents to `tree`."""
    for b in blocks:
        if isinstance(b, SceneTreeBlock):
            # XXX check node_id and is_update
            # pending_tree_nodes[b.tree_id] = b
            tree.add_node(b.tree_id, parent_id=b.parent_id)
        elif isinstance(b, TreeNodeBlock):
            # Expect this node to already exist; adding information
            # if b.node_id not in pending_tree_nodes:
            if b.group.node_id not in tree:
                raise ValueError(
                    "Node does not exist for TreeNodeBlock: %s" % b.group.node_id
                )
            node = tree[b.group.node_id]
            node.label = b.group.label
            node.visible = b.group.visible
            node.anchor_id = b.group.anchor_id
            node.anchor_type = b.group.anchor_type
            node.anchor_threshold = b.group.anchor_threshold
            node.anchor_origin_x = b.group.anchor_origin_x
        elif isinstance(b, SceneGroupItemBlock):
            # Add this entry to children of parent_id
            node_id = b.item.value
            if node_id is None:
                continue
            if node_id not in tree:
                raise ValueError(
                    "Node does not exist for SceneGroupItemBlock: %s" % node_id
                )
            item = replace(b.item, value=tree[node_id])
            tree.add_item(item, b.parent_id)
        elif isinstance(b, (SceneLineItemBlock, SceneGlyphItemBlock)):
            # Add this entry to children of parent_id
            tree.add_item(b.item, b.parent_id)
        elif isinstance(b, RootTextBlock):
            if tree.root_text is not None:
                _logger.error(
                    "Overwriting root text\n  Old: %s\n  New: %s",
                    tree.root_text,
                    b.value,
                )
            tree.root_text = b.value

    pass


def read_tree(data: tp.BinaryIO) -> SceneTree:
    """
    Parse reMarkable file and return `SceneTree`.

    :param data: reMarkable file data.
    """
    tree = SceneTree()
    build_tree(tree, read_blocks(data))
    return tree


def simple_text_document(text: str, author_uuid=None) -> Iterator[Block]:
    """Return the basic blocks to represent `text` as plain text.

    TODO: replace this with a way to generate the tree with given text, and a
    function to write a tree to blocks.

    """

    if author_uuid is None:
        author_uuid = uuid4()

    yield AuthorIdsBlock(author_uuids={1: author_uuid})

    yield MigrationInfoBlock(migration_id=CrdtId(1, 1), is_device=True)

    yield PageInfoBlock(
        loads_count=1,
        merges_count=0,
        text_chars_count=len(text) + 1,
        text_lines_count=text.count("\n") + 1,
    )

    yield SceneTreeBlock(
        tree_id=CrdtId(0, 11),
        node_id=CrdtId(0, 0),
        is_update=True,
        parent_id=CrdtId(0, 1),
    )

    yield RootTextBlock(
        block_id=CrdtId(0, 0),
        value=si.Text(
            items=CrdtSequence(
                [
                    CrdtSequenceItem(
                        item_id=CrdtId(1, 16),
                        left_id=CrdtId(0, 0),
                        right_id=CrdtId(0, 0),
                        deleted_length=0,
                        value=text,
                    )
                ]
            ),
            styles={
                CrdtId(0, 0): LwwValue(
                    timestamp=CrdtId(1, 15), value=si.ParagraphStyle.PLAIN
                ),
            },
            pos_x=-468.0,
            pos_y=234.0,
            width=936.0,
        ),
    )

    yield TreeNodeBlock(
        si.Group(
            node_id=CrdtId(0, 1),
        )
    )

    yield TreeNodeBlock(
        si.Group(
            node_id=CrdtId(0, 11),
            label=LwwValue(timestamp=CrdtId(0, 12), value="Layer 1"),
        )
    )

    yield SceneGroupItemBlock(
        parent_id=CrdtId(0, 1),
        item=CrdtSequenceItem(
            item_id=CrdtId(0, 13),
            left_id=CrdtId(0, 0),
            right_id=CrdtId(0, 0),
            deleted_length=0,
            value=CrdtId(0, 11),
        ),
    )



================================================
FILE: src/rmscene/scene_tree.py
================================================
"""Build scene tree structure from block data.

"""

from __future__ import annotations

import logging
import typing as tp

from .tagged_block_common import CrdtId
from .crdt_sequence import CrdtSequenceItem
from . import scene_items as si

_logger = logging.getLogger(__name__)


ROOT_ID = CrdtId(0, 1)


class SceneTree:
    def __init__(self):
        self.root = si.Group(ROOT_ID)
        self._node_ids = {self.root.node_id: self.root}
        self.root_text: tp.Optional[si.Text] = None

    def __contains__(self, node_id: CrdtId):
        return node_id in self._node_ids

    def __getitem__(self, node_id: CrdtId):
        return self._node_ids[node_id]

    def add_node(self, node_id: CrdtId, parent_id: CrdtId):
        if node_id in self._node_ids:
            raise ValueError("Node %s already in tree" % node_id)
        node = si.Group(node_id)
        self._node_ids[node_id] = node
        # parent = self._node_ids[parent_id]
        # parent.children.add(item)

    def add_item(self, item: CrdtSequenceItem[si.SceneItem], parent_id: CrdtId):
        if parent_id not in self._node_ids:
            raise ValueError("Parent id not known: %s" % parent_id)
        parent = self._node_ids[parent_id]
        parent.children.add(item)

    def walk(self) -> tp.Iterator[si.SceneItem]:
        """Iterate through all leaf items (not groups)."""
        yield from _walk_items(self.root)


def _walk_items(item):
    if isinstance(item, si.Group):
        for child in item.children.values():
            yield from _walk_items(child)
    else:
        yield item



================================================
FILE: src/rmscene/tagged_block_common.py
================================================
"""Helpers for reading/writing tagged block files.

"""

from __future__ import annotations

from collections.abc import Iterator
from contextlib import contextmanager
from dataclasses import dataclass
from io import BytesIO
import enum
import logging
import struct
import typing as tp


_logger = logging.getLogger(__name__)


HEADER_V6 = b"reMarkable .lines file, version=6          "


class TagType(enum.IntEnum):
    "Tag type representing the type of following data."
    ID = 0xF
    Length4 = 0xC
    Byte8 = 0x8
    Byte4 = 0x4
    Byte1 = 0x1


class UnexpectedBlockError(Exception):
    """Unexpected tag or index in block stream."""


@dataclass(eq=True, order=True, frozen=True)
class CrdtId:
    "An identifier or timestamp."
    part1: int
    part2: int

    def __repr__(self) -> str:
        return f"CrdtId({self.part1}, {self.part2})"


class DataStream:
    """Read basic values from a remarkable v6 file stream."""

    def __init__(self, data: tp.BinaryIO):
        self.data = data

    def tell(self) -> int:
        return self.data.tell()

    def read_header(self) -> None:
        """Read the file header.

        This should be the first call when starting to read a new file.

        """
        header = self.read_bytes(len(HEADER_V6))
        if header != HEADER_V6:
            raise ValueError("Wrong header: %r" % header)

    def write_header(self) -> None:
        """Write the file header.

        This should be the first call when starting to read a new file.

        """
        self.write_bytes(HEADER_V6)

    def check_tag(self, expected_index: int, expected_type: TagType) -> bool:
        """Check that INDEX and TAG_TYPE are next.

        Returns True if the expected index and tag type are found. Does not
        advance the stream.

        """
        pos = self.data.tell()
        try:
            index, tag_type = self._read_tag_values()
            return (index == expected_index) and (tag_type == expected_type)
        except (ValueError, EOFError):
            return False
        finally:
            self.data.seek(pos)  # Go back

    def read_tag(
        self, expected_index: int, expected_type: TagType
    ) -> tuple[int, TagType]:
        """Read a tag from the stream.

        Raise an error if the expected index and tag type are not found, and
        rewind the stream.

        """
        pos = self.data.tell()
        index, tag_type = self._read_tag_values()

        if index != expected_index:
            self.data.seek(pos)  # Go back
            raise UnexpectedBlockError(
                "Expected index %d, got %d, at position %d"
                % (expected_index, index, self.data.tell())
            )

        if tag_type != expected_type:
            self.data.seek(pos)  # Go back
            raise UnexpectedBlockError(
                "Expected tag type %s (0x%X), got 0x%X at position %d"
                % (
                    expected_type.name,
                    expected_type.value,
                    tag_type,
                    self.data.tell(),
                )
            )

        return index, tag_type

    def _read_tag_values(self) -> tuple[int, TagType]:
        """Read tag values from the stream."""

        x = self.read_varuint()

        # First part is an index number that identifies if this is the right
        # data we're expecting
        index = x >> 4

        # Second part is a tag type that identifies what kind of data it is
        tag_type = x & 0xF
        try:
            tag_type = TagType(tag_type)
        except ValueError as e:
            raise ValueError(
                "Bad tag type 0x%X at position %d" % (tag_type, self.data.tell())
            )

        return index, tag_type

    def write_tag(self, index: int, tag_type: TagType):
        """Write a tag to the stream."""
        x = index << 4 | int(tag_type)
        self.write_varuint(x)

    def read_bytes(self, n: int) -> bytes:
        "Read `n` bytes, raising `EOFError` if there are not enough."
        result = self.data.read(n)
        if len(result) != n:
            raise EOFError()
        return result

    def write_bytes(self, b: bytes):
        "Write bytes to underlying stream."
        self.data.write(b)

    def _read_struct(self, pattern: str):
        pattern = "<" + pattern
        n = struct.calcsize(pattern)
        return struct.unpack(pattern, self.read_bytes(n))[0]

    def _write_struct(self, pattern: str, value):
        pattern = "<" + pattern
        self.data.write(struct.pack(pattern, value))

    def read_bool(self) -> bool:
        """Read a bool from the data stream."""
        return self._read_struct("?")

    def read_uint8(self) -> int:
        """Read a uint8 from the data stream."""
        return self._read_struct("B")

    def read_uint16(self) -> int:
        """Read a uint16 from the data stream."""
        return self._read_struct("H")

    def read_uint32(self) -> int:
        """Read a uint32 from the data stream."""
        return self._read_struct("I")

    def read_float32(self) -> float:
        """Read a float32 from the data stream."""
        return self._read_struct("f")

    def read_float64(self) -> float:
        """Read a float64 (double) from the data stream."""
        return self._read_struct("d")

    def read_varuint(self) -> int:
        """Read a varuint from the data stream."""
        shift = 0
        result = 0
        while True:
            i = ord(self.read_bytes(1))
            result |= (i & 0x7F) << shift
            shift += 7
            if not (i & 0x80):
                break
        return result

    def read_crdt_id(self) -> CrdtId:
        # Based on ddvk's reader.go
        # TODO: should be var unit?
        part1 = self.read_uint8()
        part2 = self.read_varuint()
        # result = (part1 << 48) | part2
        return CrdtId(part1, part2)

    def write_bool(self, value: bool):
        """Write a bool to the data stream."""
        self._write_struct("?", value)

    def write_uint8(self, value: int):
        """Write a uint8 to the data stream."""
        return self._write_struct("B", value)

    def write_uint16(self, value: int):
        """Write a uint16 to the data stream."""
        return self._write_struct("H", value)

    def write_uint32(self, value: int):
        """Write a uint32 to the data stream."""
        return self._write_struct("I", value)

    def write_float32(self, value: float):
        """Write a float32 to the data stream."""
        return self._write_struct("f", value)

    def write_float64(self, value: float):
        """Write a float64 (double) to the data stream."""
        return self._write_struct("d", value)

    def write_varuint(self, value: int):
        """Write a varuint to the data stream."""
        if value < 0:
            raise ValueError("value is negative")
        b = bytearray()
        while True:
            to_write = value & 0x7F
            value >>= 7
            if value:
                b.append(to_write | 0x80)
            else:
                b.append(to_write)
                break
        self.data.write(b)

    def write_crdt_id(self, value: CrdtId):
        """Write a `CrdtId` to the data stream."""
        # Based on ddvk's reader.go
        # TODO: should be var unit?
        if value.part1 >= 2**8 or value.part2 >= 2**64:
            raise ValueError("CrdtId too large: %s" % value)
        self.write_uint8(value.part1)
        self.write_varuint(value.part2)
        # result = (part1 << 48) | part2


_T = tp.TypeVar("_T")

# This makes sense to be frozen, since the value should not be changed without
# updating the timestamp.
@dataclass(eq=True, frozen=True)
class LwwValue(tp.Generic[_T]):
    "Container for a last-write-wins value."
    timestamp: CrdtId
    value: _T



================================================
FILE: src/rmscene/tagged_block_reader.py
================================================
"""Read structure of remarkable .rm files version 6.

Based on my investigation of the format with lots of help from ddvk's v6 reader
code.

"""

from __future__ import annotations

from collections.abc import Iterator
from contextlib import contextmanager
from dataclasses import dataclass, KW_ONLY
import logging
import typing as tp

from .tagged_block_common import (
    DataStream,
    TagType,
    CrdtId,
    UnexpectedBlockError,
    LwwValue,
)


_logger = logging.getLogger(__name__)


@dataclass
class BlockInfo:
    "Base class for block/subblock info."
    offset: int
    size: int

    _: KW_ONLY
    extra_data: bytes = b""


@dataclass
class MainBlockInfo(BlockInfo):
    "Top-level block info."
    block_type: int
    min_version: int
    current_version: int


@dataclass
class SubBlockInfo(BlockInfo):
    "Sub-block info."


class BlockOverflowError(Exception):
    """Read past end of block."""


class TaggedBlockReader:
    """Read blocks and values from a remarkable v6 file stream."""

    def __init__(self, data: tp.BinaryIO):
        rm_data = DataStream(data)
        self.data = rm_data
        self.current_block: tp.Optional[MainBlockInfo] = None
        self._warned_about_extra_data = False

    def read_header(self) -> None:
        """Read the file header.

        This should be the first call when starting to read a new file.

        """
        self.data.read_header()

    ## Read simple values

    def read_id(self, index: int) -> CrdtId:
        """Read a tagged CRDT ID."""
        self.data.read_tag(index, TagType.ID)
        result = self.data.read_crdt_id()
        return result

    def read_bool(self, index: int) -> bool:
        """Read a tagged bool."""
        self.data.read_tag(index, TagType.Byte1)
        result = self.data.read_bool()
        return result

    def read_byte(self, index: int) -> int:
        """Read a tagged byte as an unsigned integer."""
        self.data.read_tag(index, TagType.Byte1)
        result = self.data.read_uint8()
        return result

    def read_int(self, index: int) -> int:
        """Read a tagged 4-byte unsigned integer."""
        self.data.read_tag(index, TagType.Byte4)
        # TODO: is this supposed to be signed or unsigned?
        result = self.data.read_uint32()
        return result

    def read_float(self, index: int) -> float:
        """Read a tagged 4-byte float."""
        self.data.read_tag(index, TagType.Byte4)
        result = self.data.read_float32()
        return result

    def read_double(self, index: int) -> float:
        """Read a tagged 8-byte double."""
        self.data.read_tag(index, TagType.Byte8)
        result = self.data.read_float64()
        return result

    ## Read simple values -- optional variants

    def _read_optional(self, func, index, default):
        try:
            return func(index)
        except (UnexpectedBlockError, EOFError):
            return default

    def read_id_optional(
        self, index: int, default: tp.Optional[CrdtId] = None
    ) -> tp.Optional[CrdtId]:
        """Read a tagged CRDT ID, return `default` if not present."""
        return self._read_optional(self.read_id, index, default)

    def read_bool_optional(
        self, index: int, default: tp.Optional[bool] = None
    ) -> tp.Optional[bool]:
        """Read a tagged bool, return `default` if not present."""
        return self._read_optional(self.read_bool, index, default)

    def read_byte_optional(
        self, index: int, default: tp.Optional[int] = None
    ) -> tp.Optional[int]:
        """Read a tagged byte as an unsigned integer, return `default` if not present."""
        return self._read_optional(self.read_byte, index, default)

    def read_int_optional(
        self, index: int, default: tp.Optional[int] = None
    ) -> tp.Optional[int]:
        """Read a tagged 4-byte unsigned integer, return `default` if not present."""
        return self._read_optional(self.read_int, index, default)

    def read_float_optional(
        self, index: int, default: tp.Optional[float] = None
    ) -> tp.Optional[float]:
        """Read a tagged 4-byte float, return `default` if not present."""
        return self._read_optional(self.read_float, index, default)

    def read_double_optional(
        self, index: int, default: tp.Optional[float] = None
    ) -> tp.Optional[float]:
        """Read a tagged 8-byte double, return `default` if not present."""
        return self._read_optional(self.read_double, index, default)

    ## Blocks

    @contextmanager
    def read_block(self) -> Iterator[tp.Optional[MainBlockInfo]]:
        """Read a top-level block header.

        This acts as a context manager. Upon exiting the with-block, the amount
        of data read is checked and an error raised if it has not reached the
        end of the block.

        Returns the `BlockInfo` if successfully read. If no block can be read,
        None is returned.

        """
        if self.current_block is not None:
            raise UnexpectedBlockError("Already in a block")

        try:
            block_length = self.data.read_uint32()
        except EOFError:
            yield None  # no more blocks to read
            return

        unknown = self.data.read_uint8()
        min_version = self.data.read_uint8()
        current_version = self.data.read_uint8()
        block_type = self.data.read_uint8()
        _logger.debug(
            "Block header: %d %d %d", min_version, current_version, block_type
        )
        assert unknown == 0
        assert current_version >= 0
        assert min_version >= 0
        assert min_version <= current_version

        i0 = self.data.tell()
        self.current_block = MainBlockInfo(
            offset=i0,
            size=block_length,
            block_type=block_type,
            min_version=min_version,
            current_version=current_version,
        )

        yield self.current_block

        assert self.current_block is not None
        self._check_position(self.current_block)
        self.current_block = None

    def bytes_remaining_in_block(self) -> int:
        """Return the number of bytes remaining in the current block."""
        block_info = self.current_block
        if block_info is None:
            raise ValueError("Not in a block")
        return block_info.offset + block_info.size - self.data.tell()

    @contextmanager
    def read_subblock(self, index: int) -> Iterator[SubBlockInfo]:
        """Read a subblock length and return `SubBlockInfo` as context object.

        Checks that the correct length has been read at the end of the with
        block.
        """
        self.data.read_tag(index, TagType.Length4)
        subblock_length = self.data.read_uint32()
        i0 = self.data.tell()

        subblock = SubBlockInfo(i0, subblock_length)
        yield subblock

        self._check_position(subblock)

    def has_subblock(self, index: int) -> bool:
        """Check if a subblock with the given index is next."""
        # It's possible that if we are at the end of the block, the next bytes
        # (the size of the next block) could happen to match the tag and index
        # of what we are looking for -- so check explicitly for end of block.
        if self.current_block:
            if self.bytes_remaining_in_block() <= 0:
                return False
        return self.data.check_tag(index, TagType.Length4)

    def _check_position(self, block_info: BlockInfo):
        length = block_info.size
        i0 = block_info.offset
        i1 = self.data.tell()
        if i1 > i0 + length:
            raise BlockOverflowError(
                "%s starting at %d, length %d, read up to %d (overflow by %d)"
                % (type(block_info), i0, length, i1, i1 - (i0 + length))
            )
        if i1 < i0 + length:
            if not self._warned_about_extra_data:
                _logger.warning(
                    "Some data has not been read. The data may have been written using "
                    "a newer format than this reader supports."
                )
                self._warned_about_extra_data = True
            _logger.info("In %s only read %d bytes", block_info, i1 - i0)
            # Discard the rest
            remaining = i0 + length - i1
            excess = self.data.read_bytes(remaining)
            block_info.extra_data = excess
            _logger.debug(
                "Excess bytes:\n %s",
                "\n".join(excess[i : i + 32].hex() for i in range(0, len(excess), 32)),
                stack_info=True,
                stacklevel=4,
            )

    ## Higher level constructs

    def read_lww_bool(self, index: int) -> LwwValue[bool]:
        "Read a LWW bool."
        with self.read_subblock(index):
            timestamp = self.read_id(1)
            value = self.read_bool(2)
        return LwwValue(timestamp, value)

    def read_lww_byte(self, index: int) -> LwwValue[int]:
        "Read a LWW byte."
        with self.read_subblock(index):
            timestamp = self.read_id(1)
            value = self.read_byte(2)
        return LwwValue(timestamp, value)

    def read_lww_float(self, index: int) -> LwwValue[float]:
        "Read a LWW float."
        with self.read_subblock(index):
            timestamp = self.read_id(1)
            value = self.read_float(2)
        return LwwValue(timestamp, value)

    def read_lww_id(self, index: int) -> LwwValue[CrdtId]:
        "Read a LWW ID."
        with self.read_subblock(index):
            # XXX ddvk has these the other way round?
            timestamp = self.read_id(1)
            value = self.read_id(2)
        return LwwValue(timestamp, value)

    def read_lww_string(self, index: int) -> LwwValue[str]:
        "Read a LWW string."
        with self.read_subblock(index):
            timestamp = self.read_id(1)
            string = self.read_string(2)
        return LwwValue(timestamp, string)

    def read_string(self, index: int) -> str:
        """Read a standard string block."""
        with self.read_subblock(index) as block_info:
            string_length = self.data.read_varuint()
            # XXX not sure if this is right meaning?
            is_ascii = self.data.read_bool()
            assert is_ascii == 1
            assert string_length + 2 <= block_info.size
            b = self.data.read_bytes(string_length)
            string = b.decode()
            if len(b) != len(string):
                _logger.debug(
                    "read_string: decoded %r (%d) to %r (%d)",
                    b,
                    len(b),
                    string,
                    len(string),
                )
            return string

    def read_string_with_format(self, index: int) -> tuple[str, tp.Optional[int]]:
        """Read a string block with formatting."""
        with self.read_subblock(index) as block_info:
            string_length = self.data.read_varuint()
            # XXX not sure if this is right meaning?
            is_ascii = self.data.read_bool()
            assert is_ascii == 1
            assert string_length + 2 <= block_info.size
            b = self.data.read_bytes(string_length)
            string = b.decode()
            if len(b) != len(string):
                _logger.debug(
                    "read_string: decoded %r (%d) to %r (%d)",
                    b,
                    len(b),
                    string,
                    len(string),
                )

            if self.data.check_tag(2, TagType.Byte4):
                # We have a format code
                fmt = self.read_int(2)
            else:
                fmt = None

            return string, fmt

    def read_int_pair(self, index: int) -> tp.Optional[tuple[int, int]]:
        """Read a sub block containing two uint32"""
        with self.read_subblock(index):
            first = self.data.read_uint32()
            second = self.data.read_uint32()
            return first, second



================================================
FILE: src/rmscene/tagged_block_writer.py
================================================
"""Read structure of remarkable .rm files version 6.

Based on my investigation of the format with lots of help from ddvk's v6 reader
code.

"""

from __future__ import annotations

from collections.abc import Iterator
from contextlib import contextmanager
from io import BytesIO
import logging
import typing as tp

from .tagged_block_common import (
    TagType,
    DataStream,
    CrdtId,
    LwwValue,
    UnexpectedBlockError,
)


_logger = logging.getLogger(__name__)


class TaggedBlockWriter:
    """Write blocks and values to a remarkable v6 file stream."""

    def __init__(self, data: tp.BinaryIO, options: tp.Optional[dict] = None):
        if options is None:
            options = {}
        self.options = options
        rm_data = DataStream(data)
        self.data = rm_data
        self._in_block: bool = False

    def write_header(self) -> None:
        """Write the file header.

        This should be the first call when starting to write a new file.

        """
        self.data.write_header()

    ## Write simple values

    def write_id(self, index: int, value: CrdtId):
        """Write a tagged CRDT ID."""
        self.data.write_tag(index, TagType.ID)
        self.data.write_crdt_id(value)

    def write_bool(self, index: int, value: bool):
        """Write a tagged bool."""
        self.data.write_tag(index, TagType.Byte1)
        self.data.write_bool(value)

    def write_byte(self, index: int, value: int):
        """Write a tagged byte as an unsigned integer."""
        self.data.write_tag(index, TagType.Byte1)
        self.data.write_uint8(value)

    def write_int(self, index: int, value: int):
        """Write a tagged 4-byte unsigned integer."""
        self.data.write_tag(index, TagType.Byte4)
        # TODO: is this supposed to be signed or unsigned?
        self.data.write_uint32(value)

    def write_float(self, index: int, value: float):
        """Write a tagged 4-byte float."""
        self.data.write_tag(index, TagType.Byte4)
        self.data.write_float32(value)

    def write_double(self, index: int, value: float):
        """Write a tagged 8-byte double."""
        self.data.write_tag(index, TagType.Byte8)
        self.data.write_float64(value)

    ## Blocks

    @contextmanager
    def write_block(
        self, block_type: int, min_version: int, current_version: int
    ) -> Iterator[None]:
        """Write a top-level block header.

        Within this block, other writes are accumulated, so that the
        whole block can be written out with its length at the end.

        """
        if self._in_block:
            raise UnexpectedBlockError("Already in a block")

        previous_data = self.data
        block_buf = BytesIO()
        block_data = DataStream(block_buf)
        try:
            self.data = block_data
            self._in_block = True
            yield
        finally:
            self.data = previous_data

        assert self._in_block
        self._in_block = False

        self.data.write_uint32(len(block_buf.getbuffer()))
        self.data.write_uint8(0)
        self.data.write_uint8(min_version)
        self.data.write_uint8(current_version)
        self.data.write_uint8(block_type)
        self.data.write_bytes(block_buf.getbuffer())

    @contextmanager
    def write_subblock(self, index: int) -> Iterator[None]:
        """Write a subblock tag and length once the with-block has exited.

        Within this block, other writes are accumulated, so that the
        whole block can be written out with its length at the end.
        """
        previous_data = self.data
        subblock_buf = BytesIO()
        subblock_data = DataStream(subblock_buf)
        try:
            self.data = subblock_data
            yield
        finally:
            self.data = previous_data

        self.data.write_tag(index, TagType.Length4)
        self.data.write_uint32(len(subblock_buf.getbuffer()))
        self.data.write_bytes(subblock_buf.getbuffer())
        _logger.debug("Wrote subblock %d: %s", index, subblock_buf.getvalue().hex())

    ## Higher level constructs

    def write_lww_bool(self, index: int, value: LwwValue[bool]):
        "Write a LWW bool."
        with self.write_subblock(index):
            self.write_id(1, value.timestamp)
            self.write_bool(2, value.value)

    def write_lww_byte(self, index: int, value: LwwValue[int]):
        "Write a LWW byte."
        with self.write_subblock(index):
            self.write_id(1, value.timestamp)
            self.write_byte(2, value.value)

    def write_lww_float(self, index: int, value: LwwValue[float]):
        "Write a LWW float."
        with self.write_subblock(index):
            self.write_id(1, value.timestamp)
            self.write_float(2, value.value)

    def write_lww_id(self, index: int, value: LwwValue[CrdtId]):
        "Write a LWW ID."
        with self.write_subblock(index):
            # XXX ddvk has these the other way round?
            self.write_id(1, value.timestamp)
            self.write_id(2, value.value)

    def write_lww_string(self, index: int, value: LwwValue[str]):
        "Write a LWW string."
        with self.write_subblock(index):
            self.write_id(1, value.timestamp)
            self.write_string(2, value.value)

    def write_string(self, index: int, value: str):
        """Write a standard string block."""
        with self.write_subblock(index):
            b = value.encode()
            bytes_length = len(b)
            is_ascii = True  # XXX not sure if this is right meaning?
            self.data.write_varuint(bytes_length)
            self.data.write_bool(is_ascii)
            self.data.write_bytes(b)

    def write_string_with_format(self, index: int, text: str, fmt: int):
        """Write a string block with formatting."""
        with self.write_subblock(index):
            b = text.encode()
            bytes_length = len(b)
            is_ascii = True  # XXX not sure if this is right meaning?
            self.data.write_varuint(bytes_length)
            self.data.write_bool(is_ascii)
            self.data.write_bytes(b)
            self.write_int(2, fmt)

    def write_int_pair(self, index: int, value: tuple[int, int]):
        """Read a sub block containing two uint32"""
        with self.write_subblock(index):
            self.data.write_uint32(value[0])
            self.data.write_uint32(value[1])



================================================
FILE: src/rmscene/text.py
================================================
"""Process text from remarkable scene files.

"""

from __future__ import annotations

from collections.abc import Iterable
from collections import defaultdict
from dataclasses import dataclass, field
import logging
import typing as tp

from . import scene_items as si
from .tagged_block_common import CrdtId, LwwValue
from .crdt_sequence import CrdtSequence, CrdtSequenceItem

_logger = logging.getLogger(__name__)


def expand_text_item(
    item: CrdtSequenceItem[str | int],
) -> tp.Iterator[CrdtSequenceItem[str | int]]:
    """Expand TextItem into single-character TextItems.

    Text is stored as strings in TextItems, each with an associated ID for the
    block. This ID identifies the character at the start of the block. The
    subsequent characters' IDs are implicit.

    This function expands a TextItem into multiple single-character TextItems,
    so that each character has an explicit ID.

    """

    if item.deleted_length > 0:
        assert item.value == ""
        chars = [""] * item.deleted_length
        deleted_length = 1
    elif isinstance(item.value, int):
        yield item
        return
    else:
        # Actually the value can be empty
        # assert len(item.value) > 0
        chars = item.value
        deleted_length = 0

    if not chars:
        _logger.warning("Unexpected empty text item: %s", item)
        return

    item_id = item.item_id
    left_id = item.left_id
    for c in chars[:-1]:
        right_id = CrdtId(item_id.part1, item_id.part2 + 1)
        yield CrdtSequenceItem(item_id, left_id, right_id, deleted_length, c)
        left_id = item_id
        item_id = right_id
    yield CrdtSequenceItem(item_id, left_id, item.right_id, deleted_length, chars[-1])


def expand_text_items(
    items: Iterable[CrdtSequenceItem[str | int]],
) -> tp.Iterator[CrdtSequenceItem[str | int]]:
    """Expand a sequence of TextItems into single-character TextItems."""
    for item in items:
        yield from expand_text_item(item)


@dataclass
class CrdtStr:
    """String with CrdtIds for chars and optional properties.

    The properties apply to the whole `CrdtStr`. Use a list of
    `CrdtStr`s to represent a sequence of spans of text with different
    properties.

    """

    s: str = ""
    i: list[CrdtId] = field(default_factory=list)
    properties: dict = field(default_factory=dict)

    def __str__(self):
        return self.s


@dataclass
class Paragraph:
    """Paragraph of text."""

    contents: list[CrdtStr]
    start_id: CrdtId
    style: LwwValue[si.ParagraphStyle] = field(
        default_factory=lambda: LwwValue(CrdtId(0, 0), si.ParagraphStyle.PLAIN)
    )

    def __str__(self):
        return "".join(str(s) for s in self.contents)


@dataclass
class TextDocument:
    contents: list[Paragraph]

    @classmethod
    def from_scene_item(cls, text: si.Text):
        """Extract spans of text with associated formatting and char ids.

        This uses the inline formatting introduced in v3.3.2.
        """

        char_formats = {k: lww.value for k, lww in text.styles.items()}
        if si.END_MARKER not in char_formats:
            char_formats[si.END_MARKER] = si.ParagraphStyle.PLAIN

        # Expand from strings to characters
        char_items = CrdtSequence(expand_text_items(text.items.sequence_items()))
        keys = list(char_items)
        properties = {"font-weight": "normal", "font-style": "normal"}

        def handle_formatting_code(code):
            if code == 1:
                properties["font-weight"] = "bold"
            elif code == 2:
                properties["font-weight"] = "normal"
            if code == 3:
                properties["font-style"] = "italic"
            elif code == 4:
                properties["font-style"] = "normal"
            else:
                _logger.warning("Unknown formatting code in text: %d", code)
            return properties

        def parse_paragraph_contents():
            if keys and char_items[keys[0]] == "\n":
                start_id = keys.pop(0)
            else:
                start_id = si.END_MARKER
            contents = []
            while keys:
                char = char_items[keys[0]]
                if isinstance(char, int):
                    handle_formatting_code(char)
                elif char == "\n":
                    # End of paragraph
                    break
                else:
                    assert len(char) <= 1
                    # Start a new string if text properties have changed
                    if not contents or contents[-1].properties != properties:
                        contents += [CrdtStr(properties=properties.copy())]
                    contents[-1].s += char
                    contents[-1].i += [keys[0]]
                keys.pop(0)

            return start_id, contents

        paragraphs = []
        while keys:
            start_id, contents = parse_paragraph_contents()
            if start_id in text.styles:
                p = Paragraph(contents, start_id, text.styles[start_id])
            else:
                p = Paragraph(contents, start_id)
            paragraphs += [p]

        doc = cls(paragraphs)
        return doc



================================================
FILE: tests/test_basic_data_stream.py
================================================
import pytest
from io import BytesIO
from rmscene.tagged_block_common import (
    DataStream
)


@pytest.mark.parametrize(
    "value,hexstr",
    [
        (0, "00"),
        (3, "03"),
    ],
)
def test_write_uint8(value, hexstr):
    buf = BytesIO()
    s = DataStream(buf)
    s.write_uint8(value)
    assert buf.getvalue().hex() == hexstr


@pytest.mark.parametrize(
    "value,hexstr",
    [
        (0x00, "00"),
        (0x03, "03"),
        (0x7f, "7f"),
        (0x8c, "8c01"),
        (0x9c, "9c01"),
        (0x3fff, "ff7f"),
    ],
)
def test_write_varuint(value, hexstr):
    buf = BytesIO()
    s = DataStream(buf)
    s.write_varuint(value)
    assert buf.getvalue().hex() == hexstr
    buf.seek(0)
    assert s.read_varuint() == value



================================================
FILE: tests/test_color_tool.py
================================================
from pathlib import Path

import pytest

from rmscene import SceneGlyphItemBlock, SceneLineItemBlock, read_blocks
from rmscene.scene_items import Pen, PenColor

DATA_PATH = Path(__file__).parent / "data"


@pytest.mark.parametrize(
    "block_type,colors,tools",
    [
        (SceneGlyphItemBlock, {PenColor.HIGHLIGHT}, None),
        (SceneLineItemBlock, {PenColor.HIGHLIGHT}, {Pen.SHADER}),
        (
            SceneLineItemBlock,
            {PenColor.GREEN_2, PenColor.CYAN, PenColor.MAGENTA},
            {Pen.BALLPOINT_2},
        ),
    ],
)
def test_color_tool_parsing(block_type, colors, tools):
    FILE_NAME = DATA_PATH / "Color_and_tool_v3.14.4.rm"
    with open(FILE_NAME, "rb") as f:
        result = read_blocks(f)
        for el in result:
            if not getattr(el, "item", None):
                continue
            if not getattr(el.item, "value", None):
                continue
            if isinstance(el, block_type) and el.item.value.color in colors:
                if tools is None:
                    continue
                assert el.item.value.tool in tools, "Tool and colors don't match"


def test_highlight_shader_colors():
    FILE_NAME = DATA_PATH / "More_color_highlight_shader_v3.15.4.2.rm"
    with open(FILE_NAME, "rb") as f:
        result = read_blocks(f)
        expected_colors = [
            PenColor.HIGHLIGHT_YELLOW,
            PenColor.HIGHLIGHT_BLUE,
            PenColor.HIGHLIGHT_PINK,
            PenColor.HIGHLIGHT_ORANGE,
            PenColor.HIGHLIGHT_GREEN,
            PenColor.HIGHLIGHT_GRAY,
            PenColor.SHADER_GRAY,
            PenColor.SHADER_ORANGE,
            PenColor.SHADER_MAGENTA,
            PenColor.SHADER_BLUE,
            PenColor.SHADER_RED,
            PenColor.SHADER_GREEN,
            PenColor.SHADER_YELLOW,
            PenColor.SHADER_CYAN,
            PenColor.BLACK,
            PenColor.GRAY,
            PenColor.WHITE,
            PenColor.BLUE,
            PenColor.RED,
            PenColor.GREEN_2,
            PenColor.YELLOW_2,
            PenColor.CYAN,
            PenColor.MAGENTA,
        ]
        start = 0
        for block in result:
            if isinstance(block, SceneLineItemBlock) and block.item.value:
                assert (
                    block.item.value.color == expected_colors[start]
                ), f"Unexpected color {block.item.value.color} at index {start}"
                start += 1
        assert start == len(expected_colors)



================================================
FILE: tests/test_crdt_sequence.py
================================================
from rmscene.crdt_sequence import CrdtSequenceItem, CrdtSequence
from rmscene import CrdtId


def cid(k):
    "Shorthand for making end-markers or IDs with author-id = 1."
    return CrdtId(0 if k == 0 else 1, k)


def make_item(item_id, left_id, right_id, *args):
    "Shorthand for creating ids."
    return CrdtSequenceItem(cid(item_id), cid(left_id), cid(right_id), *args)


def test_empty():
    assert list(CrdtSequence([])) == []


def test_just_one():
    items = [
        make_item(1, 0, 0, 0, "A"),
    ]
    result = list(CrdtSequence(items))
    assert result == [cid(1)]


def test_two():
    items = [
        make_item(1, 0, 0, 0, "A"),
        make_item(2, 1, 0, 0, "B"),
    ]
    result1 = list(CrdtSequence(items))
    result2 = list(CrdtSequence(reversed(items)))
    assert result1 == result2


def test_overlapping():
    items = [
        make_item(1, 0, 0, 0, "A"),
        make_item(2, 1, 0, 0, "B"),
        make_item(3, 0, 0, 0, "C"),
    ]
    result1 = list(CrdtSequence(items))
    result2 = list(CrdtSequence(reversed(items)))
    assert result1 == result2
    assert result1 == [cid(1), cid(3), cid(2)]


def test_overlapping_sorted_by_id():
    items = [
        make_item(8, 0, 0, 0, "A"),
        make_item(9, 8, 0, 0, "B"),
        make_item(3, 0, 0, 0, "C"),
    ]
    result1 = list(CrdtSequence(items))
    result2 = list(CrdtSequence(reversed(items)))
    assert result1 == result2
    assert result1 == [cid(3), cid(8), cid(9)]


def test_unknown_id():
    items = [
        make_item(28, 0, 15, 0, "A"),
        make_item(31, 30, 15, 2, ""),
        make_item(33, 32, 15, 0, "B"),
        make_item(15, 0, 0, 0, "C"),
    ]
    result1 = list(CrdtSequence(items))
    result2 = list(CrdtSequence(reversed(items)))
    assert result1 == result2
    assert result1 == [cid(28), cid(31), cid(33), cid(15)]


def test_unknown_id_at_right():
    items = [
        make_item(14, 0, 0, 0, "A"),
        make_item(19, 14, 15, 0, "V"),
        # item 15 is not defined -- this shouldn't happen, as there should be a
        # tombstone, but don't want to raise errors in this case
    ]
    result = list(CrdtSequence(items))
    assert result == [cid(14), cid(19)]


def test_iterates_in_order():
    # Order should be "AB"
    items = [
        make_item(1, 0, 0, 0, "A"),
        make_item(2, 1, 0, 0, "B"),
    ]

    for test_items in (items, reversed(items)):
        seq = CrdtSequence(test_items)
        assert list(seq.keys()) == [cid(1), cid(2)]
        assert list(seq.values()) == ["A", "B"]
        assert list(seq.items()) == list(zip(seq.keys(), seq.values()))


from hypothesis import given, note, strategies as st
from hypothesis.stateful import Bundle, RuleBasedStateMachine, rule, precondition, invariant


class CrdtSequenceComparison(RuleBasedStateMachine):
    """State machine that adds and deletes text in the sequence"""
    def __init__(self):
        super().__init__()
        # All items, even deleted ones
        self.items: dict[CrdtId, CrdtSequenceItem] = {}
        # List of items corresponding to string
        self.string_items: list[CrdtSequenceItem] = []
        self.string = ""
        self.last_id = 1

    # keys = Bundle("keys")
    # values = Bundle("values")

    # @rule(target=keys, k=st.binary())
    # def add_key(self, k):
    #     return k

    @rule(data=st.data(), c=st.characters())
    def add_char(self, data, c):
        i = data.draw(st.integers(min_value=0, max_value=len(self.string)))
        new_item = CrdtSequenceItem(
            item_id=cid(self.last_id),
            left_id=self.string_items[i-1].item_id if i >= 1 else cid(0),
            right_id=self.string_items[i].item_id if i < len(self.string_items) else cid(0),
            deleted_length=0,
            value=c
        )
        note(f"new_item: {new_item}")
        self.last_id += 1
        self.string = self.string[:i] + c + self.string[i:]
        self.string_items = self.string_items[:i] + [new_item] + self.string_items[i:]
        self.items[new_item.item_id] = new_item

    @rule(data=st.data())
    def add_empty_item(self, data):
        i = data.draw(st.integers(min_value=0, max_value=len(self.string)))
        new_item = CrdtSequenceItem(
            item_id=cid(self.last_id),
            left_id=self.string_items[i-1].item_id if i >= 1 else cid(0),
            right_id=self.string_items[i].item_id if i < len(self.string_items) else cid(0),
            deleted_length=0,
            value=""
        )
        note(f"new_item: {new_item}")
        self.last_id += 1
        self.items[new_item.item_id] = new_item

    @precondition(lambda self: len(self.string) > 0)
    @rule(data=st.data())
    def delete_char(self, data):
        i = data.draw(st.integers(min_value=0, max_value=len(self.string) - 1))
        item_id = self.string_items[i].item_id
        note(f"deleting_item: {item_id}")
        self.string = self.string[:i] + self.string[i+1:]
        self.string_items = self.string_items[:i] + self.string_items[i+1:]
        self.items[item_id].value = ""
        self.items[item_id].deleted_length = 1

    @invariant()
    def values_agree(self):
        seq = CrdtSequence(self.items.values())
        assert "".join(seq.values()) == self.string


TestCrdtSequenceComparison = CrdtSequenceComparison.TestCase



================================================
FILE: tests/test_scene_stream.py
================================================
import pytest
from io import BytesIO
from pathlib import Path
from uuid import UUID
from rmscene import (
    read_blocks,
    write_blocks,
    LwwValue,
    TaggedBlockWriter,
    TaggedBlockReader,
)
from rmscene.scene_stream import *
from rmscene.tagged_block_common import HEADER_V6
from rmscene.tagged_block_reader import MainBlockInfo
from rmscene.crdt_sequence import CrdtSequenceItem
from rmscene.scene_items import ParagraphStyle

import logging

logger = logging.getLogger(__name__)


DATA_PATH = Path(__file__).parent / "data"


def _hex_lines(b, n=32):
    return [b[i * n : (i + 1) * n].hex() for i in range(len(b) // n + 1)]


LINES_V2_FILES = [
    "Lines_v2.rm",
    "Wikipedia_highlighted_p2.rm",
]


TEST_FILES_AND_VERSIONS = [
    ("Normal_AB.rm", "3.0"),
    ("Normal_A_stroke_2_layers.rm", "3.0"),
    ("Normal_A_stroke_2_layers_v3.2.2.rm", "3.2.2"),
    ("Normal_A_stroke_2_layers_v3.3.2.rm", "3.3.2"),
    ("Bold_Heading_Bullet_Normal.rm", "3.0"),
    ("Lines_v2.rm", "3.1"),
    ("Lines_v2_updated.rm", "3.2"),  # extra 7fXXXX part of Line data was added
    ("Wikipedia_highlighted_p1.rm", "3.1"),
    ("Wikipedia_highlighted_p2.rm", "3.1"),
    ("With_SceneInfo_Block.rm", "3.4"),  # XXX version?
    ("Color_and_tool_v3.14.4.rm", "3.14"),
    ("More_color_highlight_shader_v3.15.4.2.rm", "3.15"),
]


@pytest.mark.parametrize("test_file,version", TEST_FILES_AND_VERSIONS)
def test_full_roundtrip(test_file, version):
    with open(DATA_PATH / test_file, "rb") as f:
        data = f.read()

    # XXX not sure why this is a problem -- reMarkable seems to be inconsistent
    # about the min version written to line block headers?
    if version in ("3.2.2", "3.3.2"):
        # This is not a very good way of doing it...
        data = data.replace(bytes.fromhex("010205"), bytes.fromhex("020205"))

    input_buf = BytesIO(data)
    output_buf = BytesIO()
    options = {"version": version}

    write_blocks(output_buf, read_blocks(input_buf), options)

    assert _hex_lines(input_buf.getvalue()) == _hex_lines(output_buf.getvalue())


# FIXME: remove xfail when parsing updated

TEST_FILES_FOR_FULL_PARSING = [
    pytest.param(
        filename,
        marks=pytest.mark.xfail if filename == "Color_and_tool_v3.14.4.rm" else [],
    )
    for filename, _ in TEST_FILES_AND_VERSIONS
]


@pytest.mark.parametrize("test_file", TEST_FILES_FOR_FULL_PARSING)
def test_files_fully_parsed(test_file):
    with open(DATA_PATH / test_file, "rb") as f:
        result = list(read_blocks(f))

    # Check none of the blocks were unreadable and do not have extra data
    for block in result:
        assert not isinstance(block, UnreadableBlock)
        assert not block.extra_data


def test_normal_ab():
    with open(DATA_PATH / "Normal_AB.rm", "rb") as f:
        result = list(read_blocks(f))

    assert result == [
        AuthorIdsBlock(author_uuids={1: UUID("495ba59f-c943-2b5c-b455-3682f6948906")}),
        MigrationInfoBlock(migration_id=CrdtId(1, 1), is_device=True),
        PageInfoBlock(
            loads_count=1, merges_count=0, text_chars_count=3, text_lines_count=1
        ),
        SceneTreeBlock(
            tree_id=CrdtId(0, 11),
            node_id=CrdtId(0, 0),
            is_update=True,
            parent_id=CrdtId(0, 1),
        ),
        RootTextBlock(
            block_id=CrdtId(0, 0),
            value=si.Text(
                items=CrdtSequence(
                    [
                        CrdtSequenceItem(
                            item_id=CrdtId(1, 16),
                            left_id=CrdtId(0, 0),
                            right_id=CrdtId(0, 0),
                            deleted_length=0,
                            value="AB",
                        )
                    ]
                ),
                styles={
                    CrdtId(0, 0): LwwValue(
                        timestamp=CrdtId(1, 15), value=si.ParagraphStyle.PLAIN
                    ),
                },
                pos_x=-468.0,
                pos_y=234.0,
                width=936.0,
            ),
        ),
        TreeNodeBlock(
            group=si.Group(node_id=CrdtId(0, 1)),
        ),
        TreeNodeBlock(
            group=si.Group(
                node_id=CrdtId(0, 11),
                label=LwwValue(CrdtId(0, 12), "Layer 1"),
            ),
        ),
        SceneGroupItemBlock(
            parent_id=CrdtId(0, 1),
            item=CrdtSequenceItem(
                item_id=CrdtId(0, 13),
                left_id=CrdtId(0, 0),
                right_id=CrdtId(0, 0),
                deleted_length=0,
                value=CrdtId(0, 11),
            ),
        ),
    ]


def test_read_glyph_range():
    with open(DATA_PATH / "Wikipedia_highlighted_p1.rm", "rb") as f:
        result = [
            block for block in read_blocks(f) if isinstance(block, SceneGlyphItemBlock)
        ]

    assert result[0].item.value.text == "The reMarkable uses electronic paper"


@pytest.mark.parametrize(
    "block",
    [
        AuthorIdsBlock(author_uuids={1: UUID("495ba59f-c943-2b5c-b455-3682f6948906")}),
        AuthorIdsBlock(
            author_uuids={
                1: UUID("495ba59f-c943-2b5c-b455-3682f6948906"),
                2: UUID("cd83324a-917f-11ed-bb7b-3c0754484e34"),
            }
        ),
        MigrationInfoBlock(migration_id=CrdtId(1, 1), is_device=True),
        PageInfoBlock(
            loads_count=3, merges_count=2, text_chars_count=3, text_lines_count=1
        ),
        SceneTreeBlock(
            tree_id=CrdtId(0, 11),
            node_id=CrdtId(0, 0),
            is_update=True,
            parent_id=CrdtId(0, 1),
        ),
        RootTextBlock(
            block_id=CrdtId(0, 0),
            value=si.Text(
                items=CrdtSequence(
                    [
                        CrdtSequenceItem(
                            item_id=CrdtId(1, 16),
                            left_id=CrdtId(0, 0),
                            right_id=CrdtId(0, 0),
                            deleted_length=0,
                            value="AB",
                        )
                    ]
                ),
                styles={
                    CrdtId(0, 0): LwwValue(
                        timestamp=CrdtId(1, 15), value=si.ParagraphStyle.PLAIN
                    ),
                },
                pos_x=-468.0,
                pos_y=234.0,
                width=936.0,
            ),
        ),
        TreeNodeBlock(
            group=si.Group(
                node_id=CrdtId(0, 11),
                label=LwwValue(CrdtId(0, 12), "Layer 1"),
            ),
        ),
        SceneGroupItemBlock(
            parent_id=CrdtId(0, 1),
            item=CrdtSequenceItem(
                item_id=CrdtId(0, 13),
                left_id=CrdtId(0, 0),
                right_id=CrdtId(0, 0),
                deleted_length=0,
                value=CrdtId(0, 11),
            ),
        ),
        SceneGlyphItemBlock(
            parent_id=CrdtId(0, 11),
            item=CrdtSequenceItem(
                item_id=CrdtId(1, 17),
                left_id=CrdtId(1, 16),
                right_id=CrdtId(0, 0),
                deleted_length=0,
                value=si.GlyphRange(
                    start=1536,
                    length=23,
                    text="display technology.[13]",
                    color=si.PenColor.YELLOW,
                    rectangles=[
                        si.Rectangle(
                            x=-809.061564750815,
                            y=1724.1146737357485,
                            w=333.5427440226558,
                            h=56.30432956921868,
                        ),
                        si.Rectangle(
                            x=-485.51105154941456,
                            y=1730.4364894378523,
                            w=58.22011730763188,
                            h=33.42225280328421,
                        ),
                    ],
                ),
            ),
        ),
    ],
)
def test_blocks_roundtrip(block):
    buf = BytesIO()
    writer = TaggedBlockWriter(buf)
    reader = TaggedBlockReader(buf)

    block.write(writer)
    buf.seek(0)
    logger.info("After writing block %s", type(block))
    logger.info("Buffer: %s", buf.getvalue().hex())

    block2 = Block.read(reader)

    assert block2 == block


def test_write_blocks():
    blocks = [
        MigrationInfoBlock(migration_id=CrdtId(1, 1), is_device=True),
    ]

    buf = BytesIO()
    write_blocks(buf, blocks, options={"version": "3.1"})

    assert buf.getvalue()[:43] == b"reMarkable .lines file, version=6          "
    assert buf.getvalue()[43:].hex() == "05000000000101001f01012101"


def test_blocks_keep_unknown_data_in_main_block():
    # The "E1 FF" represents new, unknown data -- note that this might need
    # to be changed in future if the next id starts to actually be used in a
    # future update!
    data_hex = """
    2E000000 0000010D
    1C 06000000
       1F 0000
       2F 0000
    2C 05000000
       1F 0000 21 01
    3C 05000000
       1F 0000 21 01
    5C 08000000
       7C050000 50070000
    E1 FF
    """
    buf = BytesIO(HEADER_V6 + bytes.fromhex(data_hex))
    block = next(read_blocks(buf))
    assert isinstance(block, SceneInfo)
    assert block.extra_data == bytes.fromhex("E1 FF")


def test_blocks_keep_unknown_data_in_value_subblock():
    # The "8f 010f" is represents new, unknown data -- note that this might need
    # to be changed in future if the next id starts to actually be used in a
    # future update!
    data_hex = """
    59000000 00020205
    1f 0219
    2f 021e
    3f 0000
    4f 0000
    54 0000 0000
    6c 4300 0000
       03
       14 0f000000
       24 00000000
       38 00000000 0000f03f
       44 00000000
       5c 1c000000
          f8fe82c2 f42a30c3 03000800 0000b869
          83c2622d 30c30000 08000000
       6f 0001
       7f 010f
       8f 0101
    """
    buf = BytesIO(HEADER_V6 + bytes.fromhex(data_hex))
    block = next(read_blocks(buf))
    assert isinstance(block, SceneLineItemBlock)
    assert block.extra_value_data == bytes.fromhex("8f 0101")


def test_error_in_block_contained():
    # First block will cause a parsing error at `0xff`. Second block should
    # still be parsed.
    data_hex = """
    06000000 00010103
    1f 0219
    aa bbcc
    05000000 00010100
    1f 0219
    21 01
    """
    buf = BytesIO(HEADER_V6 + bytes.fromhex(data_hex))
    blocks = list(read_blocks(buf))
    assert blocks == [
        UnreadableBlock(
            error="Bad tag type 0xA at position 58",
            data=bytes.fromhex("1f0219aabbcc"),
            info=MainBlockInfo(
                offset=51,
                size=6,
                extra_data=b"",
                block_type=3,
                min_version=1,
                current_version=1,
            ),
        ),
        MigrationInfoBlock(migration_id=CrdtId(0x02, 0x19), is_device=True),
    ]

    # Check all data is preserved
    buf2 = BytesIO()
    # Version 3 to be consistent with the input data used here
    write_blocks(buf2, blocks, options={"version": "3.0"})

    assert buf2.getvalue() == buf.getvalue()


from hypothesis import given, strategies as st


crdt_id_strategy = st.builds(
    CrdtId,
    st.integers(min_value=0, max_value=2**8 - 1),
    st.integers(min_value=0, max_value=2**64 - 1),
)
st.register_type_strategy(CrdtId, crdt_id_strategy)

author_ids_block_strategy = st.builds(
    AuthorIdsBlock,
    st.dictionaries(st.integers(min_value=0, max_value=65535), st.uuids()),
    extra_data=st.binary(),
)

block_strategy = st.one_of(
    [
        author_ids_block_strategy,
        st.builds(MigrationInfoBlock),
    ]
)


@given(block_strategy)
def test_blocks_roundtrip_2(block):
    buf = BytesIO()
    writer = TaggedBlockWriter(buf)
    reader = TaggedBlockReader(buf)

    block.write(writer)
    buf.seek(0)
    logger.info("After writing block %s", type(block))
    logger.info("Buffer: %s", buf.getvalue().hex())
    block2 = Block.read(reader)
    assert block2 == block


@given(...)
def test_write_id(crdt_id: CrdtId):
    buf = BytesIO()
    s = TaggedBlockWriter(buf)
    s.write_id(3, crdt_id)



================================================
FILE: tests/test_scene_tree.py
================================================
import pytest
from io import BytesIO
from pathlib import Path
from uuid import UUID
from rmscene import (
    read_blocks,
    write_blocks,
    LwwValue,
    TaggedBlockWriter,
    TaggedBlockReader,
)
from rmscene.scene_tree import *
from rmscene.scene_stream import *
from rmscene import scene_items as si
from rmscene.crdt_sequence import CrdtSequenceItem

import logging

logger = logging.getLogger(__name__)


DATA_PATH = Path(__file__).parent / "data"


# @pytest.mark.parametrize(
#     "test_file",
#     [
#         "Normal_AB.rm",
#         "Normal_A_stroke_2_layers.rm",
#         "Bold_Heading_Bullet_Normal.rm",
#         "Lines_v2.rm"
#     ]
# )
# def test_full_roundtrip(test_file):
#     with open(DATA_PATH / test_file, "rb") as f:
#         data = f.read()

#     input_buf = BytesIO(data)
#     output_buf = BytesIO()
#     options = {
#         "line_version": (2 if test_file == "Lines_v2.rm" else 1)
#     }

#     write_blocks(output_buf, read_blocks(input_buf), options)

#     assert _hex_lines(input_buf.getvalue()) == _hex_lines(output_buf.getvalue())


def tree_structure(item):
    if isinstance(item, si.Group):
        return (
            item.node_id,
            item.label.value,
            [tree_structure(child) for child in item.children.values()],
        )
    else:
        return item


def test_basic_tree_structure():
    # this is the bare minimum structure
    blocks = [
        SceneTreeBlock(
            tree_id=CrdtId(0, 11),
            node_id=CrdtId(0, 0),
            is_update=True,
            parent_id=CrdtId(0, 1),
        ),
        TreeNodeBlock(
            si.Group(CrdtId(0, 1)),
        ),
        TreeNodeBlock(
            si.Group(
                node_id=CrdtId(0, 11),
                label=LwwValue(CrdtId(0, 12), "Layer 1"),
            )
        ),
        SceneGroupItemBlock(
            parent_id=CrdtId(0, 1),
            item=CrdtSequenceItem(
                item_id=CrdtId(0, 13),
                left_id=CrdtId(0, 0),
                right_id=CrdtId(0, 0),
                deleted_length=0,
                value=CrdtId(0, 11),
            ),
        ),
    ]

    tree = SceneTree()
    build_tree(tree, blocks)

    assert tree_structure(tree.root) == (
        CrdtId(0, 1),
        "",
        [(CrdtId(0, 11), "Layer 1", [])],
    )

    assert list(tree.root.children.values()) == [
        si.Group(
            node_id=CrdtId(0, 11), children=[], label=LwwValue(CrdtId(0, 12), "Layer 1")
        )
    ]


# Example file layers.stroke.rm
#
# SceneTreeBlocks -- all have parent (0, 1)
#
# (0, 11) update=True
# (1, 27) update=True
# (1, 30) update=True
# (1, 40) update=False --> "node_id" (1, 27), parent (0, 0)
# (1, 41) update=True
#
# In ddvk's reader this is described as a "node move" info?
#
# TreeNodeBlock
#
# (0, 1)
# (0, 11) Layer 1
# (1, 27) Layer 2
# (1, 30) Layer 3
# (1, 41) Layer 4
#
# These don't have any structure (parents etc), just information -- they are
# giving data about a node in the tree? ddvk's reader says it's an error if the
# node here has not already been seen, and if it's not marked as a layer (i.e.
# its parent is the root node). It doesn't add anything new to the tree, just
# attaches the info.
#
# All of the following are "items", so are expected to be added with a parent
# which is an already defined node. Items have left/right IDs, so are ordered,
# and leave tombstones.
#
# SceneGroupItemBlock
#
# Parent: (0, 1)
#         (0, 13)       -> (0, 11)
# (0, 13) (1, 29)       -> (1, 27)
# (0, 29) (1, 32)       -> (1, 30)
# (0, 32) (1, 43)       -> (1, 41)
#
# The values refer to the "tree blocks" above, with new IDs -- like the Groups
# have their own identity and order (amongst what?), but are referring to nodes
# (layers).
#
# Is this because this defines the order of the layers? Otherwise it's not yet
# specified. So a "GroupItem" defines the order of further nested nodes.
#
# Line, parent: (0, 11)
#
#         (1, 14)       -> [line data]
# (1, 14) (1, 15)       -> DELETED
# ...
# (1, 19) (1, 20)       -> DELETED
# (1, 20) (1, 21)       -> [line data]
# ...
# (1, 25) (1, 26)       -> [line data]      at this point, the next layer must have been added
# (1, 26) (1, 46)       -> [line data]
# ...
# (1, 51) (1, 52)       -> [line data]
#
# Line, parent: (1, 30)
#
#         (1, 33)       -> [line data]
# (1, 33) (1, 34)       -> [line data]
# ...
# (1, 38) (1, 39)       -> [line data]
#
#
# So the actual line data has a parent which is a layer. (0, 11) or (1, 30)


def test_text_and_strokes():
    line1 = si.Line(
        color=si.PenColor.RED,
        tool=si.Pen.PENCIL_2,
        points=[],
        thickness_scale=2.0,
        starting_length=0.0,
    )
    line2 = si.Line(
        color=si.PenColor.BLACK,
        tool=si.Pen.FINELINER_2,
        points=[],
        thickness_scale=2.0,
        starting_length=0.0,
    )
    blocks = [
        SceneTreeBlock(
            tree_id=CrdtId(0, 13),
            node_id=CrdtId(0, 0),
            is_update=True,
            parent_id=CrdtId(0, 1),
        ),
        SceneTreeBlock(
            tree_id=CrdtId(1, 17),
            node_id=CrdtId(0, 0),
            is_update=True,
            parent_id=CrdtId(0, 1),
        ),
        SceneTreeBlock(
            tree_id=CrdtId(1, 20),
            node_id=CrdtId(0, 0),
            is_update=True,
            parent_id=CrdtId(0, 13),
        ),
        SceneTreeBlock(
            tree_id=CrdtId(1, 26),
            node_id=CrdtId(0, 0),
            is_update=True,
            parent_id=CrdtId(1, 17),
        ),
        RootTextBlock(
            block_id=CrdtId(0, 0),
            value=si.Text(
                items=CrdtSequence([
                    CrdtSequenceItem(
                        item_id=CrdtId(1, 14),
                        left_id=CrdtId(0, 0),
                        right_id=CrdtId(0, 0),
                        deleted_length=0,
                        value="A",
                    )
                ]),
                styles={},
                pos_x=-468.0,
                pos_y=234.0,
                width=936.0,
            )
        ),
        TreeNodeBlock(
            si.Group(node_id=CrdtId(0, 1)),
        ),
        TreeNodeBlock(
            si.Group(
                node_id=CrdtId(0, 13),
                label=LwwValue(timestamp=CrdtId(0, 15), value="Layer 1"),
                visible=LwwValue(timestamp=CrdtId(0, 16), value=True),
            )
        ),
        TreeNodeBlock(
            si.Group(
                node_id=CrdtId(1, 17),
                label=LwwValue(timestamp=CrdtId(1, 18), value="Layer 2"),
                visible=LwwValue(timestamp=CrdtId(0, 0), value=True),
            )
        ),
        TreeNodeBlock(
            si.Group(
                node_id=CrdtId(1, 20),
                label=LwwValue(timestamp=CrdtId(0, 0), value=""),
                visible=LwwValue(timestamp=CrdtId(0, 0), value=True),
                anchor_id=LwwValue(timestamp=CrdtId(1, 22), value=CrdtId(1, 14)),
                anchor_type=LwwValue(timestamp=CrdtId(1, 23), value=2),
                anchor_threshold=LwwValue(timestamp=CrdtId(1, 24), value=67.02755737304688),
                anchor_origin_x=LwwValue(timestamp=CrdtId(1, 20), value=-464.0),
            )
        ),
        TreeNodeBlock(
            si.Group(
                node_id=CrdtId(1, 26),
                label=LwwValue(timestamp=CrdtId(0, 0), value=""),
                visible=LwwValue(timestamp=CrdtId(0, 0), value=True),
                anchor_id=LwwValue(timestamp=CrdtId(1, 28), value=CrdtId(1, 14)),
                anchor_type=LwwValue(timestamp=CrdtId(1, 29), value=2),
                anchor_threshold=LwwValue(timestamp=CrdtId(1, 30), value=67.02755737304688),
                anchor_origin_x=LwwValue(timestamp=CrdtId(1, 26), value=-464.0),
            )
        ),
        SceneGroupItemBlock(
            parent_id=CrdtId(0, 1),
            item=CrdtSequenceItem(
                item_id=CrdtId(0, 14),
                left_id=CrdtId(0, 0),
                right_id=CrdtId(0, 0),
                deleted_length=0,
                value=CrdtId(0, 13),
            ),
        ),
        SceneGroupItemBlock(
            parent_id=CrdtId(0, 1),
            item=CrdtSequenceItem(
                item_id=CrdtId(1, 19),
                left_id=CrdtId(0, 14),
                right_id=CrdtId(0, 0),
                deleted_length=0,
                value=CrdtId(1, 17),
            ),
        ),
        SceneGroupItemBlock(
            parent_id=CrdtId(0, 13),
            item=CrdtSequenceItem(
                item_id=CrdtId(1, 21),
                left_id=CrdtId(0, 0),
                right_id=CrdtId(0, 0),
                deleted_length=0,
                value=CrdtId(1, 20),
            ),
        ),
        SceneGroupItemBlock(
            parent_id=CrdtId(1, 17),
            item=CrdtSequenceItem(
                item_id=CrdtId(1, 27),
                left_id=CrdtId(0, 0),
                right_id=CrdtId(0, 0),
                deleted_length=0,
                value=CrdtId(1, 26),
            ),
        ),
        SceneLineItemBlock(
            parent_id=CrdtId(1, 20),
            item=CrdtSequenceItem(
                item_id=CrdtId(1, 25),
                left_id=CrdtId(0, 0),
                right_id=CrdtId(0, 0),
                deleted_length=0,
                value=line1,
            ),
        ),
        SceneLineItemBlock(
            parent_id=CrdtId(1, 26),
            item=CrdtSequenceItem(
                item_id=CrdtId(1, 31),
                left_id=CrdtId(0, 0),
                right_id=CrdtId(0, 0),
                deleted_length=0,
                value=line2,
            ),
        ),
    ]

    tree = SceneTree()
    build_tree(tree, blocks)

    assert tree_structure(tree.root) == (
        CrdtId(0, 1),
        "",
        [
            (
                CrdtId(0, 13),
                "Layer 1",
                [
                    (CrdtId(1, 20), "", [line1]),
                ],
            ),
            (
                CrdtId(1, 17),
                "Layer 2",
                [
                    (CrdtId(1, 26), "", [line2]),
                ],
            ),
        ],
    )

    # TODO should also include the text items, and anchor reference
    #
    # Does this mean the tree should also be able to look up items by ID? And
    # maybe actually children list should just contain item ids?
    #
    # Otherwise how do you make use of the "anchor id" value? Could set the
    # explicit reference.
    #
    # When writing back to blocks, need to be able to flatten the structure
    # again (i.e. find out the item id for the anchor, if referenced directly).
    #
    # Currently the item id is held in the CrdtSequenceItem, not in its value.



================================================
FILE: tests/test_tagged_block_reader.py
================================================
import pytest
from io import BytesIO
from rmscene import TaggedBlockReader, UnexpectedBlockError, BlockOverflowError, CrdtId


def stream(hex_string: str) -> TaggedBlockReader:
    data = bytes.fromhex(hex_string)
    return TaggedBlockReader(BytesIO(data))


class TestBlock:
    TEST_DATA = (
        "04000000"  # length
        "00010205"  # header
        "ff000000"  # data
        "00000000"  # more data to test overflow
    )

    def test_read_block(self):
        s = stream(self.TEST_DATA)
        with s.read_block() as block_info:
            assert block_info is not None
            assert block_info.size == 4
            assert block_info.block_type == 5
            assert block_info.min_version == 1
            assert block_info.current_version == 2
            assert block_info.offset == 8

            value = s.data.read_uint32()
            assert value == 0xFF

    def test_bytes_remaining(self):
        s = stream(self.TEST_DATA)
        with s.read_block():
            assert s.bytes_remaining_in_block() == 4
            s.data.read_uint32()
            assert s.bytes_remaining_in_block() == 0

    def test_error_on_overflow(self):
        s = stream(self.TEST_DATA)
        with pytest.raises(BlockOverflowError):
            with s.read_block():
                s.data.read_uint32()
                s.data.read_uint32()

    def test_warns_if_incomplete(self, caplog):
        s = stream(self.TEST_DATA)
        with s.read_block():
            pass  # not reading anything
        assert "not been read" in caplog.records[0].message

    def test_skips_to_end_of_block_if_not_all_read(self):
        s = stream(self.TEST_DATA)
        with s.read_block():
            assert s.data.tell() == 8
            # not reading anything
        assert s.data.tell() == 12

    def test_error_if_already_in_block(self):
        s = stream(self.TEST_DATA)
        with s.read_block():
            with pytest.raises(UnexpectedBlockError):
                with s.read_block():
                    pass


class TestSubblock:
    TEST_DATA = (
        "5c" "04000000" "ff000000" "00000000"  # tag  # length  # data
    )  # more data to test overflow

    def test_read_subblock(self):
        s = stream(self.TEST_DATA)
        with s.read_subblock(5) as block_info:
            assert block_info.size == 4
            value = s.data.read_uint32()
            assert value == 0xFF

    def test_has_subblock(self):
        s = stream(self.TEST_DATA)
        # Rewinds so can be called repeatedly
        assert s.has_subblock(5) == True
        assert s.has_subblock(5) == True
        assert s.has_subblock(4) == False
        assert s.has_subblock(5) == True

    def test_error_on_wrong_index(self):
        s = stream(self.TEST_DATA)
        with pytest.raises(UnexpectedBlockError):
            with s.read_subblock(3):
                pass

    def test_error_on_wrong_tag(self):
        s = stream(self.TEST_DATA)
        with pytest.raises(UnexpectedBlockError):
            s.read_int(5)

    def test_error_on_wrong_index_does_not_consume_tag(self):
        s = stream(self.TEST_DATA)
        with pytest.raises(UnexpectedBlockError):
            with s.read_subblock(3):
                pass

        # We can still read the next block after catching the error.
        with s.read_subblock(5) as block_info:
            assert block_info.size == 4

    def test_error_on_overflow(self):
        s = stream(self.TEST_DATA)
        with pytest.raises(BlockOverflowError):
            with s.read_subblock(5):
                s.data.read_uint32()
                s.data.read_uint32()

    def test_warns_if_incomplete(self, caplog):
        s = stream(self.TEST_DATA)
        with s.read_subblock(5):
            pass  # not reading anything
        assert "not been read" in caplog.records[0].message


def test_has_subblock_returns_False_with_bad_data():
    s = stream("1d000000")
    assert s.has_subblock(1) == False
    assert s.has_subblock(6) == False
    assert s.data.tell() == 0


def test_has_subblock_returns_False_at_end_of_file():
    s = stream("")
    assert s.has_subblock(2) == False


def test_has_subblock_checks_for_end_of_block():
    # See https://github.com/ricklupton/rmscene/issues/17#issuecomment-1701071477
    #
    # Construct some potentially confusing data -- the 0x2c is the start of the
    # next block, but if we don't take care, `has_subblock(2)` could see it as a
    # subblock instead.
    data_hex = """
    03000000 00010103
    1f 0219
    2c000000 00010100
    """

    s = stream(data_hex)
    with s.read_block():
        assert s.read_id(1)
        assert s.has_subblock(2) == False


def test_read_int():
    s = stream("34abcd0000")
    assert s.read_int(3) == 0xCDAB


def test_read_int_wrong_index():
    s = stream("34abcd0000")
    with pytest.raises(UnexpectedBlockError):
        s.read_int(2)


def test_read_int_default_value():
    s = stream("34abcd0000")
    assert s.read_int_optional(2, -1) == -1
    assert s.read_int_optional(3) == 0xCDAB
    assert s.read_int_optional(4) == None


def test_read_lww_string():
    s = stream("1c0d000000" "1f0101" "2c05000000" "0301616263")
    lww = s.read_lww_string(1)
    assert lww.timestamp == CrdtId(1, 1)
    assert lww.value == "abc"


def test_read_string_ascii():
    s = stream("1c05000000" "0301616263")
    result = s.read_string(1)
    assert result == "abc"


def test_read_string_utf():
    s = stream("1c05000000" "030161c397")
    result = s.read_string(1)
    assert result == "a×"



================================================
FILE: tests/test_tagged_block_writer.py
================================================
import pytest
from io import BytesIO
from rmscene import (
    TaggedBlockReader,
    TaggedBlockWriter,
    CrdtId,
    LwwValue,
    UnexpectedBlockError,
)


def test_write_id_zero():
    buf = BytesIO()
    s = TaggedBlockWriter(buf)
    s.write_id(3, CrdtId(0, 0))
    assert buf.getvalue().hex() == "3f0000"


def test_write_int():
    buf = BytesIO()
    s = TaggedBlockWriter(buf)
    s.write_int(3, 0xCDAB)
    assert buf.getvalue().hex() == "34abcd0000"


def test_write_block():
    buf = BytesIO()
    s = TaggedBlockWriter(buf)
    with s.write_block(5, 1, 2):
        s.write_int(3, 0x1234)
    assert buf.getvalue().hex() == "05000000000102053434120000"


def test_write_block_error_if_nested():
    buf = BytesIO()
    s = TaggedBlockWriter(buf)
    with pytest.raises(UnexpectedBlockError):
        with s.write_block(5, 1, 2):
            with s.write_block(4, 1, 1):
                s.write_int(3, 0x1234)


def test_write_block_error_recovery():
    buf = BytesIO()
    s = TaggedBlockWriter(buf)
    try:
        with s.write_block(5, 1, 2):
            s.write_int(3, 0x1234)
            raise Exception
    except:
        pass

    # The subblock redirection should have recovered
    s.write_bool(7, True)
    assert buf.getvalue().hex() == "7101"


def test_write_subblock():
    buf = BytesIO()
    s = TaggedBlockWriter(buf)
    with s.write_subblock(2):
        s.write_int(3, 0x1234)
    assert buf.getvalue().hex() == "2c050000003434120000"


def test_write_subblock_nested():
    buf = BytesIO()
    s = TaggedBlockWriter(buf)
    with s.write_subblock(1):
        with s.write_subblock(2):
            s.write_int(3, 0x1234)
    assert buf.getvalue().hex() == "1c0a0000002c050000003434120000"


def test_write_subblock_error_recovery():
    buf = BytesIO()
    s = TaggedBlockWriter(buf)
    try:
        with s.write_subblock(2):
            s.write_int(3, 0x1234)
            raise Exception
    except:
        pass

    # The subblock redirection should have recovered
    s.write_bool(7, True)
    assert buf.getvalue().hex() == "7101"


@pytest.mark.parametrize(
    "data_type,value",
    [
        ("id", CrdtId(1, 5)),
        ("id", CrdtId(0, 0)),
        ("bool", True),
        ("bool", False),
        ("byte", 7),
        ("int", 45),
        ("float", 8.0),
        ("double", 5.4),
        ("lww_bool", LwwValue(CrdtId(1, 4), True)),
        ("lww_byte", LwwValue(CrdtId(1, 4), 7)),
        ("lww_float", LwwValue(CrdtId(1, 4), 8.0)),
        ("lww_id", LwwValue(CrdtId(1, 4), CrdtId(1, 5))),
        ("lww_string", LwwValue(CrdtId(1, 4), "hello")),
        ("string", "abc"),
        ("string", "a×"),
    ],
)
@pytest.mark.parametrize("index", [0, 3, 20])
def test_values_roundtrip(data_type, value, index):
    buf = BytesIO()
    writer = TaggedBlockWriter(buf)
    reader = TaggedBlockReader(buf)
    write = getattr(writer, f"write_{data_type}")
    read = getattr(reader, f"read_{data_type}")
    write(index, value)
    buf.seek(0)
    assert read(index) == value



================================================
FILE: tests/test_text.py
================================================
import pytest
from rmscene.text import (
    expand_text_item,
    expand_text_items,
    TextDocument,
    CrdtStr,
    Paragraph,
)
from rmscene import scene_items as si
from rmscene import CrdtId, CrdtSequenceItem, CrdtSequence


def cid(k):
    "Shorthand for making end-markers or IDs with author-id = 1."
    return CrdtId(0 if k == 0 else 1, k)


def make_item(item_id, left_id, right_id, *args):
    "Shorthand for creating ids."
    return CrdtSequenceItem(cid(item_id), cid(left_id), cid(right_id), *args)


def test_expand_text_1():
    result = expand_text_item(make_item(17, 0, 0, 0, "AAAA"))
    assert list(result) == [
        make_item(17, 0, 18, 0, "A"),
        make_item(18, 17, 19, 0, "A"),
        make_item(19, 18, 20, 0, "A"),
        make_item(20, 19, 0, 0, "A"),
    ]


def test_expand_text_2():
    result = expand_text_item(make_item(34, 20, 21, 0, "x"))
    assert list(result) == [
        make_item(34, 20, 21, 0, "x"),
    ]


def test_expand_text_3():
    result = expand_text_item(make_item(21, 20, 0, 0, "A\nB"))
    assert list(result) == [
        make_item(21, 20, 22, 0, "A"),
        make_item(22, 21, 23, 0, "\n"),
        make_item(23, 22, 0, 0, "B"),
    ]


def test_expand_text_empty():
    result = expand_text_item(make_item(21, 20, 0, 2, ""))
    assert list(result) == [
        make_item(21, 20, 22, 1, ""),
        make_item(22, 21, 0, 1, ""),
    ]


START_BOLD = 1
END_BOLD = 2
START_ITALIC = 3
END_ITALIC = 4


def doc_from_items(items):
    root_text = si.Text(
        items=CrdtSequence(items),
        styles={},
        pos_x=-468.0,
        pos_y=234.0,
        width=936.0,
    )
    doc = TextDocument.from_scene_item(root_text)
    return doc


def test_inline_formatting_italic_over_paragraphs():
    doc = doc_from_items(
        [
            make_item(20, 0, 0, 0, "A"),
            make_item(21, 20, 0, 0, "B\nC"),
            make_item(24, 23, 0, 0, "D"),
            # Start italic between A and B
            make_item(30, 20, 21, 0, START_ITALIC),
            # End italic between C and D
            make_item(31, 23, 24, 0, END_ITALIC),
        ]
    )

    assert doc.contents == [
        Paragraph(
            contents=[
                CrdtStr(
                    "A",
                    [CrdtId(1, 20)],
                    {"font-weight": "normal", "font-style": "normal"},
                ),
                CrdtStr(
                    "B",
                    [CrdtId(1, 21)],
                    {"font-weight": "normal", "font-style": "italic"},
                ),
            ],
            start_id=CrdtId(0, 0),
        ),
        Paragraph(
            contents=[
                CrdtStr(
                    "C",
                    [CrdtId(1, 23)],
                    {"font-weight": "normal", "font-style": "italic"},
                ),
                CrdtStr(
                    "D",
                    [CrdtId(1, 24)],
                    {"font-weight": "normal", "font-style": "normal"},
                ),
            ],
            start_id=CrdtId(1, 22),
        ),
    ]


def test_inline_formatting_italic_over_paragraphs():
    doc = doc_from_items(
        [
            make_item(20, 0, 0, 0, "A"),
            make_item(21, 20, 0, 0, "B\nC"),
            make_item(24, 23, 0, 0, "D"),
            # Start italic between A and B
            make_item(30, 20, 21, 0, START_ITALIC),
            # End italic between C and D
            make_item(31, 23, 24, 0, END_ITALIC),
        ]
    )

    assert doc.contents == [
        Paragraph(
            contents=[
                CrdtStr(
                    "A",
                    [CrdtId(1, 20)],
                    {"font-weight": "normal", "font-style": "normal"},
                ),
                CrdtStr(
                    "B",
                    [CrdtId(1, 21)],
                    {"font-weight": "normal", "font-style": "italic"},
                ),
            ],
            start_id=CrdtId(0, 0),
        ),
        Paragraph(
            contents=[
                CrdtStr(
                    "C",
                    [CrdtId(1, 23)],
                    {"font-weight": "normal", "font-style": "italic"},
                ),
                CrdtStr(
                    "D",
                    [CrdtId(1, 24)],
                    {"font-weight": "normal", "font-style": "normal"},
                ),
            ],
            start_id=CrdtId(1, 22),
        ),
    ]


def test_inline_formatting_bold_italic_interleaved_over_paragraphs():
    doc = doc_from_items(
        [
            make_item(20, 0, 0, 0, "ABC\nDEF"),
            # Start italic between A and B
            make_item(30, 20, 21, 0, START_ITALIC),
            # Start bold between B and C
            make_item(31, 21, 22, 0, START_BOLD),
            # End italic between D and E
            make_item(32, 24, 25, 0, END_ITALIC),
            # End bold between E and F
            make_item(33, 25, 26, 0, END_BOLD),
        ]
    )

    assert doc.contents == [
        Paragraph(
            contents=[
                CrdtStr(
                    "A",
                    [CrdtId(1, 20)],
                    {"font-weight": "normal", "font-style": "normal"},
                ),
                CrdtStr(
                    "B",
                    [CrdtId(1, 21)],
                    {"font-weight": "normal", "font-style": "italic"},
                ),
                CrdtStr(
                    "C",
                    [CrdtId(1, 22)],
                    {"font-weight": "bold", "font-style": "italic"},
                ),
            ],
            start_id=CrdtId(0, 0),
        ),
        Paragraph(
            contents=[
                CrdtStr(
                    "D",
                    [CrdtId(1, 24)],
                    {"font-weight": "bold", "font-style": "italic"},
                ),
                CrdtStr(
                    "E",
                    [CrdtId(1, 25)],
                    {"font-weight": "bold", "font-style": "normal"},
                ),
                CrdtStr(
                    "F",
                    [CrdtId(1, 26)],
                    {"font-weight": "normal", "font-style": "normal"},
                ),
            ],
            start_id=CrdtId(1, 23),
        ),
    ]



================================================
FILE: tests/test_text_files.py
================================================
from uuid import UUID
from io import BytesIO
from pathlib import Path
from rmscene.scene_stream import *
from rmscene.scene_items import Text, ParagraphStyle
from rmscene.text import TextDocument, CrdtStr

DATA_PATH = Path(__file__).parent / "data"


def _hex_lines(b, n=32):
    return [b[i * n : (i + 1) * n].hex() for i in range(len(b) // n + 1)]


def extract_doc(filename):
    with open(filename, "rb") as f:
        tree = read_tree(f)
        assert tree.root_text
        doc = TextDocument.from_scene_item(tree.root_text)
        return doc


def show_str_formatting(x):
    # Basic logic -- could be smarter about removing adjacent
    # unnecessary opening/closing tags
    s = x.s
    if x.properties.get("font-weight") == "bold":
        s = f"<b>{s}</b>"
    if x.properties.get("font-style") == "italic":
        s = f"<i>{s}</i>"
    return s


def formatted_lines(doc):
    return [
        (p.style.value, "".join(show_str_formatting(s) for s in p.contents))
        for p in doc.contents
    ]


def extract_paragraphs(filename):
    with open(filename, "rb") as f:
        tree = read_tree(f)
        assert tree.root_text


def test_normal_ab():
    lines = formatted_lines(extract_doc(DATA_PATH / "Normal_AB.rm"))
    assert lines == [(ParagraphStyle.PLAIN, "AB")]


def test_list():
    lines = formatted_lines(extract_doc(DATA_PATH / "Bold_Heading_Bullet_Normal.rm"))
    assert lines == [
        (ParagraphStyle.BOLD, "A"),
        (ParagraphStyle.HEADING, "new line"),
        (ParagraphStyle.BULLET, "B is a letter of the alphabet"),
        (ParagraphStyle.PLAIN, "C"),
    ]


def test_inline_formats():
    doc = extract_doc(DATA_PATH / "Normal_A_stroke_2_layers_v3.3.2.rm")
    lines = formatted_lines(doc)
    assert lines == [
        (ParagraphStyle.PLAIN, "A"),
        (ParagraphStyle.PLAIN, "v3.2.2"),
        (
            ParagraphStyle.PLAIN,
            "Normal <b>bold</b> <i>italic</i>",
        ),
        (
            ParagraphStyle.PLAIN,
            "<b>Bold</b> <i>italic</i> normal",
        ),
        (ParagraphStyle.BOLD, "Bold line"),
        (ParagraphStyle.PLAIN, "Normal line"),
        (ParagraphStyle.HEADING, "Heading line"),
    ]


def test_simple_text_document():
    test_file = "Normal_AB.rm"
    with open(DATA_PATH / test_file, "rb") as f:
        expected = f.read()

    output_buf = BytesIO()
    author_id = UUID("495ba59f-c943-2b5c-b455-3682f6948906")
    write_blocks(
        output_buf, simple_text_document("AB", author_id), options={"version": "3.0"}
    )

    assert _hex_lines(output_buf.getvalue()) == _hex_lines(expected)



================================================
FILE: tests/data/Bold_Heading_Bullet_Normal.rm
================================================
[Binary file]


================================================
FILE: tests/data/Color_and_tool_v3.14.4.rm
================================================
[Binary file]


================================================
FILE: tests/data/Lines_v2.rm
================================================
[Binary file]


================================================
FILE: tests/data/Lines_v2_updated.rm
================================================
[Binary file]


================================================
FILE: tests/data/More_color_highlight_shader_v3.15.4.2.rm
================================================
[Binary file]


================================================
FILE: tests/data/Normal_A_stroke_2_layers.rm
================================================
[Binary file]


================================================
FILE: tests/data/Normal_A_stroke_2_layers_v3.2.2.rm
================================================
[Binary file]


================================================
FILE: tests/data/Normal_A_stroke_2_layers_v3.3.2.rm
================================================
[Binary file]


================================================
FILE: tests/data/Normal_AB.rm
================================================
[Binary file]


================================================
FILE: tests/data/Wikipedia_highlighted_p1.rm
================================================
[Binary file]


================================================
FILE: tests/data/Wikipedia_highlighted_p2.rm
================================================
[Binary file]


================================================
FILE: tests/data/With_SceneInfo_Block.rm
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
    if: startsWith(github.ref, 'refs/tags/') || github.ref == 'refs/heads/main'
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
      url: https://pypi.org/p/rmscene
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
    if: github.repository == 'ricklupton/rmscene' # publish to testpypi only on the original repo
    needs:
    - build
    runs-on: ubuntu-latest

    environment:
      name: testpypi
      url: https://test.pypi.org/p/rmscene

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
    - name: Install poetry
      run: pipx install poetry
    - name: Setup Python ${{ matrix.python-version }}
      uses: actions/setup-python@v4
      with:
        python-version: ${{ matrix.python-version }}
        cache: 'poetry'
    - run: poetry install
    - run: poetry run pytest tests


