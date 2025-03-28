package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
)

var (
	ignoreCase   = flag.Bool("i", false, "대소문자 구분 없이 검색")
	lineNumber   = flag.Bool("n", false, "라인 번호 출력")
	invert       = flag.Bool("v", false, "매칭되지 않는 줄 출력")
	color        = flag.Bool("color", false, "일치하는 부분 색상 출력")
	recursive    = flag.Bool("r", false, "디렉토리 재귀 검색")
	onlyMatching = flag.Bool("o", false, "일치하는 부분만 출력")
	showVersion  = flag.Bool("version", false, "버전 정보 출력")
)

const version = "v0.0.4"

func grepFile(pattern *regexp.Regexp, filename string, showFilename bool) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "파일 열기 실패: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
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
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "파일 읽기 에러: %v\n", err)
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
				grepFile(pattern, p, true)
			}
			return nil
		})
	} else {
		grepFile(pattern, path, len(flag.Args()) > 2)
	}
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `사용법: grepa [옵션] 패턴 [파일...]

grepa는 grep 기능을 간단하게 구현한 경량 CLI 도구입니다.

옵션:
`)
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, `
예시:
  grepa -o "[a-zA-Z0-9._%%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}" access.log    # 이메일 주소 추출
  grepa " 2[\d]{2} " access.log                                           # 200번대 HTTP 응답 코드 추출
  grepa "POST " access.log                                                # POST 요청 추출
  grepa "wreply=" access.log                                              # wreply 파라미터 추출
  grepa -o "https?://[a-zA-Z0-9.-]+(:[0-9]+)?" access.log                 # URL 추출
  grepa -o '"https?://[^"]+"' access.log                                  # Referer 추출

`)
	}

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
