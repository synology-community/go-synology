package methods

import (
	"github.com/synology-community/synology-api/pkg/api"

	"github.com/synology-community/synology-api/pkg/api/docker"
)

const (
	API_DockerContainer = "SYNO.Docker.Container"
	methodCreate        = "create"
	methodGet           = "get"
	methodUpload        = "upload"
	methodList          = "list"
	methodListShares    = "list_share"
	methodRename        = "rename"
	methodInfo          = "get"
	methodStart         = "start"
	methodDelete        = "delete"
)

var (
	Create = api.Method{
		API:     API_DockerContainer,
		Version: 1,
		Method:  methodCreate,
		ErrorSummary: docker.CommonErrors.Combine(api.ErrorSummary{
			1100: "Failed to create a folder. More information in <errors> object.",
			1101: "The number of folders to the parent folder would exceed the system limitation.",
		}),
	}
	Get = api.Method{
		API:          API_DockerContainer,
		Version:      1,
		Method:       methodGet,
		ErrorSummary: docker.CommonErrors,
	}
	List = api.Method{
		API:          API_DockerContainer,
		Version:      1,
		Method:       methodList,
		ErrorSummary: docker.CommonErrors,
	}
	Delete = api.Method{
		API:          API_DockerContainer,
		Version:      1,
		Method:       methodDelete,
		ErrorSummary: docker.CommonErrors,
	}
	RegistryList = api.Method{
		API:          "SYNO.Docker.Registry",
		Version:      1,
		Method:       methodGet,
		ErrorSummary: docker.CommonErrors,
	}
)
