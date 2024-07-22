go run ./tools/saveSwagger.go -app scheduler -out $PWD/cmd/scheduler/swagger.json
go run ./tools/saveSwagger.go -app cdg -out $PWD/cmd/cdg/swagger.json
go run ./tools/saveSwagger.go -app collector -out $PWD/cmd/collector/swagger.json