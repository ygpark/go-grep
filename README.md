# grepa

`grepa`는 Go로 구현된 간단하고 강력한 `grep` 클론 도구입니다. 파일 또는 표준 입력에서 정규표현식 패턴을 검색하며, 다양한 옵션을 통해 유연한 검색 기능을 제공합니다.

## 🛠️ 기능

- 정규표현식 기반 검색
- 대소문자 구분 없이 검색 (`-i`)
- 라인 번호 출력 (`-n`)
- 일치하지 않는 줄 출력 (`-v`)
- 일치한 텍스트 색상 강조 (`--color`)
- 디렉토리 재귀 검색 (`-r`)
- 표준 입력 처리 지원 (파이프 사용 가능)

## 📆 설치

```bash
git clone https://github.com/your-username/grepa.git
cd grepa
go build -o grepa main.go
```

## 🚀 사용법

```bash
./grepa [옵션] "패턴" [파일 또는 디렉토리]
```

파일 없이 사용하면 `stdin`(표준 입력)을 통해 검색할 수 있습니다.

### 예시

```bash
# test.txt 파일에서 "error"를 검색
./grepa "error" test.txt

# 대소문자 무시하고 검색
./grepa -i "error" test.txt

# 라인 번호 포함해서 출력
./grepa -n "error" test.txt

# 일치하지 않는 줄 출력
./grepa -v "success" test.txt

# 검색된 문자열을 빨간색으로 강조
./grepa --color "fail" test.txt

# 디렉토리 내부 모든 파일을 재귀적으로 검색
./grepa -r "panic" ./logs

# 파이프로 사용
cat access.log | ./grepa "404"
```

## 🔧 옵션

| 옵션      | 설명                              |
| --------- | --------------------------------- |
| `-i`      | 대소문자 구분 없이 검색           |
| `-n`      | 라인 번호 출력                    |
| `-v`      | 매칭되지 않는 줄 출력             |
| `--color` | 일치하는 문자열을 빨간색으로 강조 |
| `-r`      | 디렉토리 재귀 검색                |

## 📁 예제

```bash
echo -e "Hello\nWorld\nHELLO" | ./grepa -i "hello"
# 결과:
# Hello
# HELLO
```

## 📜 라이센스

MIT License
