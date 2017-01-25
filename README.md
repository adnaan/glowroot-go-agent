protoc -I /Users/adnaan/code/web/glowroot/wire-api/src/main/proto/ \
/Users/adnaan/code/web/glowroot/wire-api/src/main/proto/*.proto \
--go_out=plugins=grpc:org_glowroot_wire_api_model

