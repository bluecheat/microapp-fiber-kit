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
3. `.env.yaml` 을 `env.yaml`로 변경합니다.
    ```bash
   mv .env.yaml env.yaml
   ```
4. 애플리케이션을 실행합니다.
   ```bash
   go run cmd/microapp/main.go
   ```

## 패키지 구조

패키지 구조는 기본적으로 클린아키텍처를 기반으로 도메인 영역과 인프라스트럭처 영역을 분리하였고,
내부 비즈니스 로직의 경우 internal 내에서 `aggregateRoot` > `service`, `repository`, `message` 순으로 작성하였습니다.

( 예시로 board와 user 디렉토리가 템플릿에 포함되어 있습니다. )

```markdown
ㄴ cmd
    ㄴ microapp
        ㄴ app.go # 서비스 의존성 주입 관리 (*중요)
        ㄴ args.go
        ㄴ main.go # 실행 파일 존재
ㄴ config # 환경변수 관리
ㄴ database # 데이터베이스 관리
ㄴ docs # swagger 자동생성파일 (수정 X)
ㄴ internal # 실제 비즈니스,도메인 로직
   ㄴ domains 엔티티
    ㄴ ... 도메인 애그리게이트
ㄴ server
    ㄴ middleware # 세션 혹은 token 미들웨어
    ㄴ router # api 라우팅 테이블
ㄴ utils # 공통 사용 함수
```

## 사용 방법

### 환경변수 설정

프로젝트 루트에 `env.yaml` 를 통해 아래와 같은 환경변수를 설정 할 수 있습니다.

- 서버의 host, port 변경
- cors 설정
- log 설정
- database 설정

### API 추가방법

1. internal 디렉토리 하위에 도메인 디렉토리 생성 및 `message`, `service`, `repository` 로직 작성
2. `cmd / microapp / app` 디렉토리 하위에 의존성 주입
   ```go
   // 의존성 주입
   func providers() []interface{} {
         return []interface{}{
            database.NewDatabase,
            board.NewBoardRepository,
            board.NewBoardService,
            user.NewUserRepository,
            user.NewUserService,
   
             // 아래 부분에 생성자 함수 추가
            other.NewOtherService,
            ...
       }
   }
   ```

3. `server / api.go` 에 앞서 주입한 서비스를 파라미터로 전달 받습니다.
4. `server / router / handler.go` 에서 api 설정 및 service 를 매핑합니다.
   ```go
    func Route(
         router fiber.Router,
         boardSrv *board.BoardService,
         userSrv *user.UserService,

          // 아래 부분에 매핑할 서비스 추가
         other.OtherService,
         ...
      ) {
         v1 := router.Group(V1)
      
         v1.Post("/board", wrapHandler[board.CreateBoardRequest, board.BoardMsg](boardSrv.CreateBoard))
         v1.Get("/board/:id", wrapHandler[board.GetBoardRequest, board.BoardMsg](boardSrv.GetBoard))
         v1.Get("/board", wrapHandler[board.GetBoardsRequest, board.BoardsMsg](boardSrv.GetBoards))
      
         v1.Post("/login", wrapHandler[user.LoginRequest, user.UserMsg](userSrv.Login))
         v1.Post("/join", wrapHandler[user.JoinRequest, user.UserMsg](userSrv.Join))
        
         // 아래 부분에 매핑할 API 작성 
        ...
      }
   ```

5. `go run cmd/microapp/main.go` 실행

### 문서 작성 방법

1. swag 기반으로 swagger 주석을 사용하고 있기때문에 https://github.com/swaggo/swag 참고하여 각 `service` 함수 상단에 작성합니다.
   ```go
   // Service godoc
   // @Summary		Other API
   // @Accept		json
   // @Produce		json
   // @Param 		Other body Other true "Other"
   // @Success		200		{object}	OtherMessage
   // @Router		/v1/other [post]
   func Service(..) ... {
   ```
2. 루트에 make_apidocs.sh 실행
3. api 실행 후 `/swagger/index.html` 접속