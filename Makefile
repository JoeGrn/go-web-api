install:
	cd infra && npm install
	cd lambda && make install

build:
	cd lambda && make build

synth:
	cd infra && cdk deploy

deploy:
	cd infra && cdk deploy