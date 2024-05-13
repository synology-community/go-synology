package virtualization

import "github.com/synology-community/synology-api/pkg/api"

type VirtualizationAPI interface {
	GetGuest(name string) (*Guest, error)
	ListGuests() (*GuestList, error)
}

var API_METHODS = api.APIMethodLookup{
	"GetGuest": {
		API:          "SYNO.Virtualization.API.Guest",
		Version:      1,
		Method:       "get",
		ErrorSummary: api.GlobalErrors,
	},
	"ListGuests": {
		API:          "SYNO.Virtualization.API.Guest",
		Version:      1,
		Method:       "list",
		ErrorSummary: api.GlobalErrors,
	},
}
