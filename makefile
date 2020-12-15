build-dev:
	@[ "${version}" ] || ( echo ">> please provide version=vX.Y.Z"; exit 1 )
    go build -o ~/.terraform.d/plugins/example.com/xjfan/ffmsgraph/${version}/darwin_amd64/terraform-provider-ffmsgraph .
.PHONY: build-dev

# make build-dev version=v1.0.0