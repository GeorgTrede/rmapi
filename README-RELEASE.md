# rmapi - reMarkable Cloud API Tool

This is a pre-built binary for macOS with Apple Silicon (M1/M2/M3).

## Installation

1. Download the `rmapi-darwin-arm64` binary
2. Make it executable: `chmod +x rmapi-darwin-arm64`
3. (Optional) Move to a directory in your PATH: `sudo mv rmapi-darwin-arm64 /usr/local/bin/rmapi`

## Usage

```bash
# Connect to your reMarkable account
./rmapi-darwin-arm64

# List documents
rmapi> ls

# Download a document with annotations
rmapi> geta document.pdf
```

## Features

- Full V6 format support (reMarkable software version 3+)
- Extracts handwritten annotations from both old (V3/V5) and new (V6) formats
- Generates PDF files with annotations overlaid

## Notes

- On first run, macOS may block the binary. Go to System Preferences > Security & Privacy to allow it
- Alternatively, remove the quarantine flag: `xattr -d com.apple.quarantine rmapi-darwin-arm64`

## More Information

See the main repository for full documentation: https://github.com/juruen/rmapi
