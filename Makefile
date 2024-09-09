.PHONY: codegen

gen-model:
	sqlboiler mysql -c .sqlboiler.toml -o generated/models -p models --no-tests --wipe

gen-oapi:
	oapi-codegen -generate types -package opapp -config generated/openapi/config.yaml -o ./generated/openapi/app/zz_generated.types.go  ./openapi/api.yaml
	oapi-codegen -generate server -package opapp  -o ./generated/openapi/app/zz_generated.server.go  ./openapi/api.yaml


migrate:
	atlas schema apply -u "mysql://root:pass@localhost:13306/test" --to file://schema.hcl

schema-inspect:
	atlas schema inspect -u "mysql://root:pass@localhost:13306/test" --web