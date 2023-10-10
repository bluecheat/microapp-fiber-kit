# Microapp Fiber Kit

## 소개

Microapp Fiber Kit은 Go 언어로 작성된 마이크로서비스 애플리케이션 템플릿입니다. 이 프로젝트는 [Fiber](https://github.com/gofiber/fiber) 웹 프레임워크를 기반으로
하며, 다양한 기능과 유틸리티를 제공합니다.

## 주요 특징

- **Clean Architecture & DDD 구조**: 클린 아키텍처 및 DDD 관점으로 설계된 패키지 구조
- **손쉬운 환경 설정**: `env.yaml` 파일을 통한 환경 변수 관리
- **데이터베이스 연동**: GORM을 이용한 데이터베이스 관리
- **API 문서화**: Swagger를 통한 API 문서 자동 생성
- **비즈니스로직 개발 편의성**: fx 라이브러리를 사용한 DI 주입 및 템플릿화

## 설치 방법

1. 저장소를 클론합니다.
   ```bash
   git clone https://github.com/bluecheat/microapp-fiber-kit.git
   ```
2. 의존성을 설치합니다.
   ```bash
   go mod download
   ```
3. 애플리케이션을 실행합니다.
   ```bash
   go run cmd/microapp/main.go
   ```

## 사용 방법

### API

API 문서는 /docs 경로에서 확인할 수 있습니다.

### 데이터베이스

데이터베이스 설정은 .env.yaml에서 할 수 있습니다.
