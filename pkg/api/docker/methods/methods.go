package methods

import (
	"github.com/synology-community/synology-api/pkg/api"

	"github.com/synology-community/synology-api/pkg/api/docker"
)

const (
	API_DockerContainer = "SYNO.Docker.Container"
	API_DockerRegistry  = "SYNO.Docker.Registry"
	API_DockerImage     = "SYNO.Docker.Image"
	API_DockerProject   = "SYNO.Docker.Project"
)

var (
	Create = api.Method{
		API:     API_DockerContainer,
		Version: 1,
		Method:  api.MethodCreate,
		ErrorSummary: docker.CommonErrors.Combine(api.ErrorSummary{
			1100: "Failed to create a folder. More information in <errors> object.",
			1101: "The number of folders to the parent folder would exceed the system limitation.",
		}),
	}
	Get = api.Method{
		API:          API_DockerContainer,
		Version:      1,
		Method:       api.MethodGet,
		ErrorSummary: docker.CommonErrors,
	}
	List = api.Method{
		API:          API_DockerContainer,
		Version:      1,
		Method:       api.MethodList,
		ErrorSummary: docker.CommonErrors,
	}
	Delete = api.Method{
		API:          API_DockerContainer,
		Version:      1,
		Method:       api.MethodDelete,
		ErrorSummary: docker.CommonErrors,
	}
	RegistryList = api.Method{
		API:          API_DockerRegistry,
		Version:      1,
		Method:       api.MethodGet,
		ErrorSummary: docker.CommonErrors,
	}
	ImagePullStart = api.Method{
		API:          API_DockerImage,
		Version:      1,
		Method:       api.MethodPullStart,
		ErrorSummary: docker.CommonErrors,
	}
	ImagePullStatus = api.Method{
		API:          API_DockerImage,
		Version:      1,
		Method:       api.MethodPullStatus,
		ErrorSummary: docker.CommonErrors,
	}
	ImageDelete = api.Method{
		API:          API_DockerImage,
		Version:      1,
		Method:       api.MethodDelete,
		ErrorSummary: docker.CommonErrors,
	}
	ProjectGet = api.Method{
		API:          API_DockerProject,
		Version:      1,
		Method:       api.MethodGet,
		ErrorSummary: docker.CommonErrors,
	}
	ProjectList = api.Method{
		API:          API_DockerProject,
		Version:      1,
		Method:       api.MethodList,
		ErrorSummary: docker.CommonErrors,
	}
	ProjectCreate = api.Method{
		API:          API_DockerProject,
		Version:      1,
		Method:       api.MethodCreate,
		ErrorSummary: docker.CommonErrors,
	}
	ProjectUpdate = api.Method{
		API:          API_DockerProject,
		Version:      1,
		Method:       api.MethodUpdate,
		ErrorSummary: docker.CommonErrors,
	}
	ProjectDelete = api.Method{
		API:          API_DockerProject,
		Version:      1,
		Method:       api.MethodDelete,
		ErrorSummary: docker.CommonErrors,
	}
	ProjectCleanStream = api.Method{
		API:          API_DockerProject,
		Version:      1,
		Method:       api.MethodCleanStream,
		ErrorSummary: docker.CommonErrors,
	}
)
