module go-learning-server

go 1.12

require (
	github.com/bilibili/kratos v0.0.0-20190417043059-a08e0f771592 // indirect
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.7.1
	github.com/go-sql-driver/mysql v1.4.1
	github.com/google/wire v0.5.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.4.0
	golang.org/x/net v0.0.0-20210510120150-4163338589ed // indirect
	golang.org/x/sys v0.0.0-20210514084401-e8d321eab015 // indirect
	golang.org/x/text v0.3.6 // indirect
	google.golang.org/genproto v0.0.0-20210513213006-bf773b8c8384 // indirect
	google.golang.org/grpc v1.37.1 // indirect
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.0.0
	google.golang.org/protobuf v1.26.0
	gopkg.in/yaml.v2 v2.3.0
)

replace (
	cloud.google.com/go => github.com/googleapis/google-cloud-go v0.26.0
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190123085648-057139ce5d2b
	golang.org/x/lint => github.com/golang/lint v0.0.0-20181026193005-c67002cb31c3
	golang.org/x/net => github.com/golang/net v0.0.0-20190420063019-afa5a82059c6
	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20180821212333-d2e6202438be
	golang.org/x/sync => github.com/golang/sync v0.0.0-20181108010431-42b317875d0f
	golang.org/x/sys => github.com/golang/sys v0.0.0-20180905080454-ebe1bf3edb33
	golang.org/x/text => github.com/golang/text v0.3.0
	golang.org/x/time => github.com/golang/time v0.0.0-20190308202827-9d24e82272b4
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190328211700-ab21143f2384
	google.golang.org/appengine => github.com/golang/appengine v1.1.0
	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20180817151627-c66870c02cf8
)
