module github.com/tnynlabs/wyrm-pipeline

go 1.16

require (
	github.com/araddon/qlbridge v0.0.2
	github.com/go-gomail/gomail v0.0.0-20160411212932-81ebce5c23df
	github.com/golang/protobuf v1.5.2
	github.com/joho/godotenv v1.3.0
	github.com/tnynlabs/wyrm v0.0.0
	google.golang.org/grpc v1.39.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
)

replace github.com/tnynlabs/wyrm v0.0.0 => ./wyrm
