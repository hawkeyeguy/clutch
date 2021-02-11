module github.com/lyft/clutch/backend

go 1.13

require (
	github.com/DATA-DOG/go-sqlmock v1.5.0
	github.com/Masterminds/squirrel v1.5.0
	github.com/aws/aws-sdk-go-v2 v1.2.0
	github.com/aws/aws-sdk-go-v2/config v1.1.1
	github.com/aws/aws-sdk-go-v2/service/autoscaling v1.1.1
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.1.1
	github.com/aws/aws-sdk-go-v2/service/kinesis v1.1.1
	github.com/aws/aws-sdk-go-v2/service/s3 v1.2.0
	github.com/bufbuild/buf v0.30.0
	github.com/cactus/go-statsd-client/statsd v0.0.0-20200623234511-94959e3146b2
	github.com/coreos/go-oidc/v3 v3.0.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/envoyproxy/go-control-plane v0.9.9-0.20201217023817-7fe139bd184a
	github.com/envoyproxy/protoc-gen-validate v0.4.1
	github.com/fullstorydev/grpcurl v1.8.0
	github.com/go-git/go-billy/v5 v5.0.0
	github.com/go-git/go-git/v5 v5.2.0
	github.com/gobwas/glob v0.2.3
	github.com/golang-migrate/migrate/v4 v4.14.1
	github.com/golang/protobuf v1.4.3
	github.com/google/go-github/v32 v32.1.0
	github.com/google/uuid v1.2.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.2.0
	github.com/iancoleman/strcase v0.1.3
	github.com/jhump/protoreflect v1.7.1
	github.com/lib/pq v1.9.0
	github.com/mitchellh/hashstructure/v2 v2.0.1
	github.com/shurcooL/githubv4 v0.0.0-20201206200315-234843c633fa
	github.com/shurcooL/graphql v0.0.0-20181231061246-d48a9a75455f // indirect
	github.com/shurcooL/httpfs v0.0.0-20190707220628-8d4bc4ba7749 // indirect
	github.com/shurcooL/vfsgen v0.0.0-20200824052919-0d455de96546
	github.com/slack-go/slack v0.8.0
	github.com/stretchr/testify v1.7.0
	github.com/uber-go/tally v3.3.17+incompatible
	go.uber.org/zap v1.16.0
	golang.org/x/net v0.0.0-20210119194325-5f4716e94777
	golang.org/x/oauth2 v0.0.0-20210210192628-66670185b0cd
	golang.org/x/sync v0.0.0-20201207232520-09787c993a3a
	google.golang.org/genproto v0.0.0-20210207032614-bba0dbe2a9ea
	google.golang.org/grpc v1.35.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.1.0
	google.golang.org/protobuf v1.25.0
	gopkg.in/square/go-jose.v2 v2.5.1
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776
	k8s.io/api v0.20.2
	k8s.io/apimachinery v0.20.2
	k8s.io/client-go v0.20.2
)
