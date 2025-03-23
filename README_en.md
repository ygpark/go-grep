# grepa

`grepa` is a simple yet powerful `grep` clone written in Go. It searches for regular expression patterns in files or from standard input, offering flexible search capabilities with various options.

## 🛠️ Features

- Regex-based pattern matching
- Case-insensitive search (`-i`)
- Show line numbers (`-n`)
- Invert match to show non-matching lines (`-v`)
- Highlight matched text in red (`--color`)
- Recursive directory search (`-r`)
- Print only the matched parts (`-o`)
- Supports standard input (pipe support)

## 📦 Installation

### From source

```bash
git clone https://github.com/your-username/grepa.git
cd grepa
go build -o grepa main.go
```

### Using go install

```bash
go install github.com/ygpark/grepa@latest
```

Make sure your `$GOBIN` is in your `PATH` to run `greppa` directly.

## 🚀 Usage

```bash
./grepa [options] "pattern" [file or directory]
```

If no file is provided, it reads from standard input (`stdin`).

### Examples

```bash
# Search for "error" in test.txt
./grepa "error" test.txt

# Case-insensitive search
./grepa -i "error" test.txt

# Show line numbers with matches
./grepa -n "error" test.txt

# Show lines that do NOT match "success"
./grepa -v "success" test.txt

# Highlight matches in red
./grepa --color "fail" test.txt

# Recursively search all files in a directory
./grepa -r "panic" ./logs

# Print only matching words
./grepa -o "[a-zA-Z]+" test.txt

# Use with pipe
cat access.log | ./grepa "404"
```

## 🔧 Options

| Option    | Description                                          |
| --------- | ---------------------------------------------------- |
| `-i`      | Case-insensitive search                              |
| `-n`      | Show line numbers                                    |
| `-v`      | Show lines that do not match                         |
| `--color` | Highlight matched text in red                        |
| `-r`      | Recursively search directories                       |
| `-o`      | Print only the matched part (incompatible with `-v`) |

## 📁 Sample

```bash
echo -e "foo bar\nbaz123\nhello world" | ./grepa -o "[a-z]+"
# Output:
# foo
# bar
# baz
# hello
# world
```

## 📜 License

MIT License
