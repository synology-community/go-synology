package methods

import (
	"github.com/synology-community/go-synology/pkg/api"
)

const (
	API_DockerContainer = "SYNO.Docker.Container"
	API_DockerRegistry  = "SYNO.Docker.Registry"
	API_DockerImage     = "SYNO.Docker.Image"
	API_DockerProject   = "SYNO.Docker.Project"
	API_DockerNetwork   = "SYNO.Docker.Network"
)

var (
	Create = api.Method{
		API:     API_DockerContainer,
		Version: 1,
		Method:  api.MethodCreate,
		ErrorSummaries: CommonErrors.Combine(api.ErrorSummary{
			1100: "Failed to create a folder. More information in <errors> object.",
			1101: "The number of folders to the parent folder would exceed the system limitation.",
		}),
	}
	Get = api.Method{
		API:            API_DockerContainer,
		Version:        1,
		Method:         api.MethodGet,
		ErrorSummaries: CommonErrors,
	}
	List = api.Method{
		API:            API_DockerContainer,
		Version:        1,
		Method:         api.MethodList,
		ErrorSummaries: CommonErrors,
	}
	Delete = api.Method{
		API:            API_DockerContainer,
		Version:        1,
		Method:         api.MethodDelete,
		ErrorSummaries: CommonErrors,
	}
	RegistryList = api.Method{
		API:            API_DockerRegistry,
		Version:        1,
		Method:         api.MethodGet,
		ErrorSummaries: CommonErrors,
	}
	ImagePullStart = api.Method{
		API:            API_DockerImage,
		Version:        1,
		Method:         api.MethodPullStart,
		ErrorSummaries: CommonErrors,
	}
	ImagePullStatus = api.Method{
		API:            API_DockerImage,
		Version:        1,
		Method:         api.MethodPullStatus,
		ErrorSummaries: CommonErrors,
	}
	ImageDelete = api.Method{
		API:            API_DockerImage,
		Version:        1,
		Method:         api.MethodDelete,
		ErrorSummaries: CommonErrors,
	}
	ProjectGet = api.Method{
		API:            API_DockerProject,
		Version:        1,
		Method:         api.MethodGet,
		ErrorSummaries: CommonErrors,
	}
	ProjectList = api.Method{
		API:            API_DockerProject,
		Version:        1,
		Method:         api.MethodList,
		ErrorSummaries: CommonErrors,
	}
	ProjectCreate = api.Method{
		API:            API_DockerProject,
		Version:        1,
		Method:         api.MethodCreate,
		ErrorSummaries: CommonErrors,
	}
	ProjectUpdate = api.Method{
		API:            API_DockerProject,
		Version:        1,
		Method:         api.MethodUpdate,
		ErrorSummaries: CommonErrors,
	}
	ProjectDelete = api.Method{
		API:            API_DockerProject,
		Version:        1,
		Method:         api.MethodDelete,
		ErrorSummaries: CommonErrors,
	}
	ProjectCleanStream = api.Method{
		API:            API_DockerProject,
		Version:        1,
		Method:         api.MethodCleanStream,
		ErrorSummaries: CommonErrors,
	}
	ProjectStopStream = api.Method{
		API:            API_DockerProject,
		Version:        1,
		Method:         api.MethodStopStream,
		ErrorSummaries: CommonErrors,
	}
	ProjectRestartStream = api.Method{
		API:            API_DockerProject,
		Version:        1,
		Method:         api.MethodRestartStream,
		ErrorSummaries: CommonErrors,
	}
	ProjectStartStream = api.Method{
		API:            API_DockerProject,
		Version:        1,
		Method:         api.MethodStartStream,
		ErrorSummaries: CommonErrors,
	}
	ProjectBuildStream = api.Method{
		API:            API_DockerProject,
		Version:        1,
		Method:         api.MethodBuildStream,
		ErrorSummaries: CommonErrors,
	}
	NetworkCreate = api.Method{
		API:            API_DockerNetwork,
		Version:        1,
		Method:         api.MethodCreate,
		ErrorSummaries: CommonErrors,
	}
	NetworkList = api.Method{
		API:            API_DockerNetwork,
		Version:        1,
		Method:         api.MethodList,
		ErrorSummaries: CommonErrors,
	}
	NetworkDelete = api.Method{
		API:            API_DockerNetwork,
		Version:        1,
		Method:         api.MethodRemove,
		ErrorSummaries: CommonErrors,
	}
	NetworkSet = api.Method{
		API:            API_DockerNetwork,
		Version:        1,
		Method:         api.MethodSet,
		ErrorSummaries: CommonErrors,
	}
)
