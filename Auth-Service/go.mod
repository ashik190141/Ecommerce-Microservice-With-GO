module Auth-Service

go 1.25.4

require github.com/joho/godotenv v1.5.1

require github.com/gorilla/mux v1.8.1

require (
	github.com/golang-jwt/jwt/v5 v5.3.0
	github.com/jmoiron/sqlx v1.4.0
	github.com/lib/pq v1.10.9
	golang.org/x/crypto v0.45.0
	golang.org/x/net v0.47.0
	golang.org/x/sys v0.38.0
	golang.org/x/text v0.31.0
	google.golang.org/genproto/googleapis/rpc v0.0.0-20251022142026-3a174f9686a8
	google.golang.org/grpc v1.77.0
	google.golang.org/protobuf v1.36.10
)
