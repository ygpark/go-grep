# grepa

`grepa` is a simple and powerful `grep`-like tool written in Go. It searches for regular expression patterns in files or from standard input, with a range of flexible options.

## 🛠️ Features

- Regex-based search
- Case-insensitive search (`-i`)
- Show line numbers (`-n`)
- Invert match (`-v`)
- Highlight matches in red (`--color`)
- Recursive directory search (`-r`)
- Supports standard input (pipe-friendly)

## 📦 Installation

```bash
git clone https://github.com/your-username/grepa.git
cd grepa
go build -o grepa main.go
```

## 🚀 Usage

```bash
./grepa [options] "pattern" [file or directory]
```

If no file is given, it reads from `stdin`.

### Examples

```bash
# Search for "error" in test.txt
./grepa "error" test.txt

# Case-insensitive search
./grepa -i "error" test.txt

# Print line numbers
./grepa -n "error" test.txt

# Show lines that do NOT match
./grepa -v "success" test.txt

# Highlight matched text in red
./grepa --color "fail" test.txt

# Recursively search all files in a directory
./grepa -r "panic" ./logs

# Use with pipe
cat access.log | ./grepa "404"
```

## 🔧 Options

| Option    | Description                            |
| --------- | -------------------------------------- |
| `-i`      | Case-insensitive search                |
| `-n`      | Show line numbers                      |
| `-v`      | Invert match (show non-matching lines) |
| `--color` | Highlight matching text in red         |
| `-r`      | Recursive directory search             |

## 📁 Example

```bash
echo -e "Hello\nWorld\nHELLO" | ./grepa -i "hello"
# Output:
# Hello
# HELLO
```

## 📜 License

MIT License
