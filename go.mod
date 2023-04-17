module github.com/t-kimura1988/japon-service

go 1.20

require github.com/t-kimura1988/japon-proto/pkg v0.0.0-20230413140722-cded3c91372c // indirectindirect

require (
	github.com/t-kimura1988/japon-domain v0.0.0-20230411162642-f13a3183fc69
	google.golang.org/grpc v1.54.0
	google.golang.org/protobuf v1.30.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.3.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/crypto v0.8.0 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1 // indirect
	gorm.io/driver/postgres v1.5.0 // indirect
	gorm.io/gorm v1.24.7-0.20230306060331-85eaf9eeda11 // indirect
)

replace github.com/t-kimura1988/japon-domain => /Users/kimuratakeshi/work/web/japon-pro/japon-domain
