echo Update swagger json with following command:
echo go run ./tools/saveSwagger.go -app scheduler -out %CD%\swagger.json
docker run --rm -v %CD%:/src swaggerapi/swagger-codegen-cli:2.4.41 generate -l go -i /src/swagger.json -o /src/go_client -c /src/config.json