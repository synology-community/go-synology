package methods

import (
	"github.com/synology-community/synology-api/pkg/api"
)

const (
	API_Core_Package = "SYNO.Core.Package"
)

var (
	PackageList = api.Method{
		API:          API_Core_Package,
		Version:      2,
		Method:       api.MethodGet,
		ErrorSummary: api.GlobalErrors,
	}
)
