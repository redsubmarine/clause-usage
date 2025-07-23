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

## 설치 및 실행

```bash
go run main.go
```

## 개발

```bash
# 의존성 설치
go mod tidy

# 빌드
go build -o claude-usage-app

# 실행
./claude-usage-app
``` 