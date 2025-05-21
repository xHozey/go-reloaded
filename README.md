# 📝 Text Modifier Tool (Go)

## Description

This is a simple yet powerful text processing tool written in Go. It reads a plain text file, applies a set of transformation rules to the content, and writes the formatted result to a new file. The tool was built as part of a collaborative project focusing on Go programming, string manipulation, file handling, and command-line applications.

This tool can perform:

- Number base conversions (hexadecimal and binary to decimal)
- Word casing modifications (uppercase, lowercase, capitalization)
- Contextual grammar corrections (e.g., switching "a" to "an")
- Proper punctuation spacing
- Quote wrapping
- Batch word casing edits

## 🔧 Features

- `(hex)` — Converts the word before to its decimal equivalent, assuming it's a hex number.  
  **Example**: `1E (hex)` → `30`

- `(bin)` — Converts the word before to its decimal equivalent, assuming it's a binary number.  
  **Example**: `10 (bin)` → `2`

- `(up)` — Converts the previous word to uppercase.  
  **Example**: `go (up)` → `GO`

- `(low)` — Converts the previous word to lowercase.  
  **Example**: `SHOUTING (low)` → `shouting`

- `(cap)` — Capitalizes the first letter of the previous word.  
  **Example**: `bridge (cap)` → `Bridge`

- `(up, n)`, `(low, n)`, `(cap, n)` — Applies the corresponding operation to the previous `n` words.  
  **Example**: `so exciting (up, 2)` → `SO EXCITING`

- **Punctuation normalization**:
  - Ensures `.,!?;:` are correctly spaced (no space before, one space after)
  - Preserves grouped punctuation (e.g., `...`, `!?`)
  
- **Quote formatting**: Wraps `' ... '` tightly around words  
  **Example**: `' awesome '` → `'awesome'`

- **Grammar fix**: Automatically switches `a` to `an` if the next word starts with a vowel or `h`.  
  **Example**: `a amazing` → `an amazing`

## 🗂 Usage

```bash
go run . <input_file.txt> <output_file.txt>
