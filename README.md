# grepa

`grepa`는 Go로 만든 간단하면서도 강력한 `grep` 클론 도구입니다. 파일 또는 표준 입력에서 정규표현식 패턴을 검색하며, 다양한 옵션을 통해 유연한 검색 기능을 제공합니다.

## 🛠️ 주요 기능

- 정규표현식 기반 검색
- 대소문자 구분 없이 검색 (`-i`)
- 라인 번호 출력 (`-n`)
- 일치하지 않는 줄 출력 (`-v`)
- 일치하는 텍스트 색상 강조 (`--color`)
- 디렉토리 재귀 검색 (`-r`)
- 일치한 부분만 출력 (`-o`)
- 표준 입력 처리 지원 (파이프 사용 가능)

## 📦 설치 방법

### 소스에서 직접 설치

```bash
git clone https://github.com/your-username/grepa.git
cd grepa
go build -o grepa main.go
```

### go install 명령어로 설치

```bash
go install github.com/ygpark/grepa@latest
```

설치 후 `$GOBIN` 경로가 `PATH`에 포함되어 있어야 `grepa` 명령을 직접 실행할 수 있습니다.

## 🚀 사용법

```bash
./grepa [옵션] "패턴" [파일 또는 디렉토리]
```

파일을 지정하지 않으면 표준 입력(`stdin`)을 통해 데이터를 읽습니다.

### 사용 예시

```bash
# test.txt 파일에서 "error" 문자열 검색
./grepa "error" test.txt

# 대소문자 무시하고 검색
./grepa -i "error" test.txt

# 라인 번호 포함하여 출력
./grepa -n "error" test.txt

# "success"와 일치하지 않는 줄 출력
./grepa -v "success" test.txt

# 일치한 문자열을 빨간색으로 강조
./grepa --color "fail" test.txt

# logs 디렉토리 내 모든 파일을 재귀적으로 검색
./grepa -r "panic" ./logs

# 일치하는 단어만 출력
./grepa -o "[a-zA-Z]+" test.txt

# 파이프 입력으로 사용
cat access.log | ./grepa "404"
```

## 🔧 지원 옵션

| 옵션      | 설명                                  |
| --------- | ------------------------------------- |
| `-i`      | 대소문자 구분 없이 검색               |
| `-n`      | 라인 번호 출력                        |
| `-v`      | 일치하지 않는 줄 출력                 |
| `--color` | 일치한 문자열을 빨간색으로 강조       |
| `-r`      | 디렉토리 재귀 검색                    |
| `-o`      | 일치한 부분만 출력 (`-v`와 병행 불가) |

## 📁 출력 예시

```bash
echo -e "foo bar\nbaz123\nhello world" | ./grepa -o "[a-z]+"
# 결과:
# foo
# bar
# baz
# hello
# world
```

## 📜 라이선스

MIT License
