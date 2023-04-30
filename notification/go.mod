module notification

go 1.19

require (
	github.com/aws/aws-lambda-go v1.36.1
	github.com/aws/aws-sdk-go v1.44.246
	github.com/mitchellh/mapstructure v1.5.0
)

require github.com/jmespath/go-jmespath v0.4.0 // indirect

replace gopkg.in/yaml.v2 => gopkg.in/yaml.v2 v2.2.8
