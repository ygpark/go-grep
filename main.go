package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	ignoreCase    = flag.Bool("i", false, "대소문자 구분 없이 검색")
	lineNumber    = flag.Bool("n", false, "라인 번호 출력")
	invert        = flag.Bool("v", false, "매칭되지 않는 줄 출력")
	color         = flag.Bool("color", false, "일치하는 부분 색상 출력")
	recursive     = flag.Bool("r", false, "디렉토리 재귀 검색")
	onlyMatching  = flag.Bool("o", false, "일치하는 부분만 출력")
	showVersion   = flag.Bool("version", false, "버전 정보 출력")
	withFilename  = flag.Bool("H", false, "결과에 항상 파일 이름 포함")
	onlyFilenames = flag.Bool("l", false, "매칭되는 파일 이름만 출력")
	// 파일 필터 옵션 추가 (예: *.eml)
	includePattern = flag.String("include", "", "검색할 파일 패턴 (예: *.eml)")
)

const version = "v0.0.4"

func grepFile(pattern *regexp.Regexp, filename string, showFilename bool) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "파일 열기 실패: %v\n", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	lineNum := 1
	hasMatch := false

	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Fprintf(os.Stderr, "파일 읽기 에러: %v\n", err)
			break
		}

		line = strings.TrimRight(line, "\r\n")

		if *onlyMatching && *invert {
			fmt.Fprintln(os.Stderr, "-o 옵션과 -v 옵션은 함께 사용할 수 없습니다.")
			os.Exit(1)
		}

		if *onlyMatching {
			matches := pattern.FindAllString(line, -1)
			if matches != nil {
				hasMatch = true
				for _, match := range matches {
					prefix := ""
					if showFilename {
						prefix += filename + ":"
					}
					if *lineNumber {
						prefix += fmt.Sprintf("%d:", lineNum)
					}
					fmt.Println(prefix + match)
				}
			}
		} else {
			match := pattern.MatchString(line)
			if *invert {
				match = !match
			}
			if match {
				hasMatch = true
				if *onlyFilenames {
					break
				}
				prefix := ""
				if showFilename {
					prefix += filename + ":"
				}
				if *lineNumber {
					prefix += fmt.Sprintf("%d:", lineNum)
				}
				if *color && !*invert {
					line = pattern.ReplaceAllString(line, "\033[31m$0\033[0m")
				}
				fmt.Println(prefix + line)
			}
		}
		lineNum++

		if err == io.EOF {
			break
		}
	}

	if *onlyFilenames && hasMatch {
		fmt.Println(filename)
	}
}

func grepStdin(pattern *regexp.Regexp) {
	scanner := bufio.NewScanner(os.Stdin)
	lineNum := 1

	for scanner.Scan() {
		line := scanner.Text()

		if *onlyMatching && *invert {
			fmt.Fprintln(os.Stderr, "-o 옵션과 -v 옵션은 함께 사용할 수 없습니다.")
			os.Exit(1)
		}

		if *onlyMatching {
			matches := pattern.FindAllString(line, -1)
			if matches != nil {
				for _, match := range matches {
					prefix := ""
					if *lineNumber {
						prefix = fmt.Sprintf("%d:", lineNum)
					}
					fmt.Println(prefix + match)
				}
			}
		} else {
			match := pattern.MatchString(line)
			if *invert {
				match = !match
			}
			if match {
				prefix := ""
				if *lineNumber {
					prefix = fmt.Sprintf("%d:", lineNum)
				}
				if *color && !*invert {
					line = pattern.ReplaceAllString(line, "\033[31m$0\033[0m")
				}
				fmt.Println(prefix + line)
			}
		}
		lineNum++
	}
}

func processPath(pattern *regexp.Regexp, path string) {
	info, err := os.Stat(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "경로 접근 실패: %v\n", err)
		return
	}

	// 파일이 직접 지정된 경우 include 옵션 체크 후 처리
	if !info.IsDir() {
		if *includePattern != "" {
			matched, err := filepath.Match(*includePattern, filepath.Base(path))
			if err != nil || !matched {
				return
			}
		}
		// -H 옵션만 있을 때 파일명 출력
		grepFile(pattern, path, *withFilename)
		return
	}

	if info.IsDir() {
		if !*recursive {
			fmt.Fprintf(os.Stderr, "%s 는 디렉토리입니다. -r 옵션을 사용하세요\n", path)
			return
		}
		filepath.WalkDir(path, func(p string, d fs.DirEntry, err error) error {
			if err != nil {
				return nil
			}
			if !d.IsDir() {
				// --include 옵션 체크
				if *includePattern != "" {
					matched, err := filepath.Match(*includePattern, filepath.Base(p))
					if err != nil || !matched {
						return nil
					}
				}
				grepFile(pattern, p, *withFilename)
			}
			return nil
		})
	}
}

func main() {
	flag.Parse()

	if *showVersion {
		fmt.Println("grepa", version)
		return
	}

	if flag.NArg() < 1 {
		flag.Usage()
		return
	}

	rawPattern := flag.Arg(0)
	if *ignoreCase {
		rawPattern = "(?i)" + rawPattern
	}

	pattern, err := regexp.Compile(rawPattern)
	if err != nil {
		fmt.Fprintf(os.Stderr, "패턴 에러: %v\n", err)
		os.Exit(1)
	}

	if flag.NArg() == 1 {
		grepStdin(pattern)
		return
	}

	files := flag.Args()[1:]
	for _, file := range files {
		processPath(pattern, file)
	}
}
