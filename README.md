# Claude Usage Menu Bar App

macOS 메뉴바에서 Claude 사용량을 실시간으로 확인할 수 있는 Go 애플리케이션입니다.

## 기능

- 메뉴바에 오늘의 토큰 사용량과 비용 표시
- 클릭 시 월별 상세 사용량 테이블 팝오버 표시
- 실시간 데이터 업데이트

## 요구사항

- macOS
- Go 1.21+
- ccusage CLI (설치되어 있어야 함)

## 설치

### 원클릭 설치 (권장)

```bash
curl -fsSL https://raw.githubusercontent.com/redsubmarine/clause-usage/main/install.sh | bash
```

### 수동 설치

```bash
curl -L https://github.com/redsubmarine/clause-usage/releases/download/v1.0.3/claude-usage-macos-arm64 -o claude-usage && chmod +x claude-usage && sudo mv claude-usage /usr/local/bin/
```

## 사용법

```bash
claude-usage          # 백그라운드에서 실행 (터미널 즉시 반환)
claude-usage test     # 테스트 모드 (터미널에서 직접 실행)
pkill claude-usage    # 종료
```

설치 후 터미널에서 `claude-usage` 명령으로 실행하면 백그라운드 데몬으로 실행되어 터미널이 즉시 반환됩니다.

### 개발자 설치

```bash
# 소스코드 다운로드
git clone https://github.com/redsubmarine/clause-usage.git
cd clause-usage

# 개발 모드로 실행
go run .

# 테스트 모드로 실행 (CLI)
go run . test

# 빌드 후 실행
./build.sh
./claude-usage-app
```

## 개발

```bash
# 의존성 설치
go mod tidy

# 빌드
./build.sh

# 실행
./claude-usage-app

# 테스트
./claude-usage-app test
```

## 기능

- **메뉴바 표시**: 오늘의 토큰 사용량과 비용을 메뉴바에 표시 (예: "3.3K - $17.58")
- **자동 갱신**: 5분마다 자동으로 데이터 갱신
- **수동 갱신**: 메뉴에서 "Refresh" 클릭으로 즉시 갱신
- **월별 데이터**: "Show Monthly Data" 클릭 시 터미널에 상세 테이블 출력
- **테스트 모드**: CLI에서 데이터 파싱 및 테이블 출력 테스트 