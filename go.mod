module github.com/synology-community/go-synology

go 1.22.4

require (
	github.com/google/go-querystring v1.1.0
	github.com/hashicorp/go-multierror v1.1.1
	github.com/hashicorp/go-retryablehttp v0.7.7
	github.com/sirupsen/logrus v1.9.3
	github.com/stretchr/testify v1.9.0
	golang.org/x/net v0.26.0
	golang.org/x/text v0.16.0
	golang.org/x/tools v0.22.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/mod v0.18.0 // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/synology-community/synology-api v0.0.9 => github.com/synology-community/go-synology v0.0.9
