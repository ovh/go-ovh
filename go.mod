module github.com/ovh/go-ovh

go 1.18

require (
	github.com/jarcoal/httpmock v1.3.0
	github.com/maxatome/go-testdeep v1.12.0
	gopkg.in/ini.v1 v1.67.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/stretchr/testify v1.8.2 // indirect
)

retract (
	v1.4.1 // Configuration fetch from wrong folder
	v1.4.0 // Configuration fetch from wrong folder
)
