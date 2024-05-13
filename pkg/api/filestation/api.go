package filestation

import (
	"github.com/synology-community/synology-api/pkg/api"
	"github.com/synology-community/synology-api/pkg/models"
	"github.com/synology-community/synology-api/pkg/util/form"
)

const UploadKey = "Upload"

type FileStationApi interface {
	CreateFolder(paths []string, names []string, forceParent bool) (*models.FolderList, error)
	ListShares() (*models.ShareList, error)
	Upload(path string, file *form.File, createParents bool, overwrite bool) (*UploadResponse, error)
	Rename(path string, name string, newName string) (*models.FileList, error)
	Download(path string, mode string) (*DownloadResponse, error)
	DeleteStart(paths []string, accurateProgress bool) (*DeleteStartResponse, error)
	DeleteStatus(taskID string) (*DeleteStatusResponse, error)
	MD5Start(path string) (*MD5StartResponse, error)
	MD5Status(taskID string) (*MD5StatusResponse, error)
}

var API_METHODS = api.APIMethodLookup{
	UploadKey: {
		API:     "SYNO.FileStation.Upload",
		Version: 3,
		Method:  "upload",
		ErrorSummary: CommonErrors.Combine(api.ErrorSummary{
			1100: "Failed to create a folder. More information in <errors> object.",
			1101: "The number of folders to the parent folder would exceed the system limitation.",
		}),
	},
	"ListShares": {
		API:          "SYNO.FileStation.List",
		Version:      2,
		Method:       "list_share",
		ErrorSummary: CommonErrors,
	},
	"Rename": {
		API:     "SYNO.FileStation.Rename",
		Version: 2,
		Method:  "rename",
		ErrorSummary: CommonErrors.Combine(api.ErrorSummary{
			1200: "Failed to rename it.",
		}),
	},
	"Info": {
		API:          "SYNO.FileStation.Info",
		Version:      1,
		Method:       "get",
		ErrorSummary: CommonErrors,
	},
	"CreateFolder": {
		API:     "SYNO.FileStation.CreateFolder",
		Version: 2,
		Method:  "create",
		ErrorSummary: api.ErrorSummary{
			1100: "Failed to create a folder. More information in <errors> object.",
			1101: "The number of folders to the parent folder would exceed the system limitation.",
		},
	},
	"DeleteStart": {
		API:          "SYNO.FileStation.Delete",
		Version:      2,
		Method:       "start",
		ErrorSummary: CommonErrors,
	},
	"DeleteStatus": {
		API:          "SYNO.FileStation.Delete",
		Version:      1,
		Method:       "status",
		ErrorSummary: CommonErrors,
	},
	"Download": {
		API:          "SYNO.FileStation.Download",
		Version:      2,
		Method:       "download",
		ErrorSummary: CommonErrors,
	},
	"MD5Start": {
		API:          "SYNO.FileStation.MD5",
		Version:      2,
		Method:       "start",
		ErrorSummary: CommonErrors,
	},
	"MD5Status": {
		API:          "SYNO.FileStation.MD5",
		Version:      2,
		Method:       "status",
		ErrorSummary: CommonErrors,
	},
	"FileStation.BackgroundTask":        {API: "SYNO.FileStation.BackgroundTask", Version: 3},
	"FileStation.CheckExist":            {API: "SYNO.FileStation.CheckExist", Version: 2},
	"FileStation.CheckPermission":       {API: "SYNO.FileStation.CheckPermission", Version: 3},
	"FileStation.Compress":              {API: "SYNO.FileStation.Compress", Version: 3},
	"FileStation.CopyMove":              {API: "SYNO.FileStation.CopyMove", Version: 3},
	"FileStation.CreateFolder":          {API: "SYNO.FileStation.CreateFolder", Version: 2},
	"FileStation.Delete":                {API: "SYNO.FileStation.Delete", Version: 2},
	"FileStation.DirSize":               {API: "SYNO.FileStation.DirSize", Version: 2},
	"FileStation.External.GoogleDrive":  {API: "SYNO.FileStation.External.GoogleDrive", Version: 2},
	"FileStation.Extract":               {API: "SYNO.FileStation.Extract", Version: 2},
	"FileStation.Favorite":              {API: "SYNO.FileStation.Favorite", Version: 2},
	"FileStation.FormUpload":            {API: "SYNO.FileStation.FormUpload", Version: 2},
	"FileStation.Info":                  {API: "SYNO.FileStation.Info", Version: 2},
	"FileStation.List":                  {API: "SYNO.FileStation.List", Version: 2},
	"FileStation.MD5":                   {API: "SYNO.FileStation.MD5", Version: 2},
	"FileStation.Mount":                 {API: "SYNO.FileStation.Mount", Version: 1},
	"FileStation.Mount.List":            {API: "SYNO.FileStation.Mount.List", Version: 1},
	"FileStation.Notify":                {API: "SYNO.FileStation.Notify", Version: 1},
	"FileStation.PhotoUpload":           {API: "SYNO.FileStation.PhotoUpload", Version: 3},
	"FileStation.Property":              {API: "SYNO.FileStation.Property", Version: 1},
	"FileStation.Property.ACLOwner":     {API: "SYNO.FileStation.Property.ACLOwner", Version: 1},
	"FileStation.Property.CompressSize": {API: "SYNO.FileStation.Property.CompressSize", Version: 1},
	"FileStation.Property.Mtime":        {API: "SYNO.FileStation.Property.Mtime", Version: 1},
	"FileStation.Rename":                {API: "SYNO.FileStation.Rename", Version: 2},
	"FileStation.Search":                {API: "SYNO.FileStation.Search", Version: 2},
	"FileStation.Search.History":        {API: "SYNO.FileStation.Search.History", Version: 1},
	"FileStation.Settings":              {API: "SYNO.FileStation.Settings", Version: 1},
	"FileStation.Sharing":               {API: "SYNO.FileStation.Sharing", Version: 3},
	"FileStation.Sharing.Download":      {API: "SYNO.FileStation.Sharing.Download", Version: 1},
	"FileStation.Snapshot":              {API: "SYNO.FileStation.Snapshot", Version: 2},
	"FileStation.Thumb":                 {API: "SYNO.FileStation.Thumb", Version: 3},
	"FileStation.Timeout":               {API: "SYNO.FileStation.Timeout", Version: 1},
	"FileStation.UIString":              {API: "SYNO.FileStation.UIString", Version: 1},
	"FileStation.Upload":                {API: "SYNO.FileStation.Upload", Version: 3},
	"FileStation.UserGrp":               {API: "SYNO.FileStation.UserGrp", Version: 1},
	"FileStation.VFS.Connection":        {API: "SYNO.FileStation.VFS.Connection", Version: 1},
	"FileStation.VFS.File":              {API: "SYNO.FileStation.VFS.File", Version: 1},
	"FileStation.VFS.GDrive":            {API: "SYNO.FileStation.VFS.GDrive", Version: 1},
	"FileStation.VFS.Profile":           {API: "SYNO.FileStation.VFS.Profile", Version: 1},
	"FileStation.VFS.Protocol":          {API: "SYNO.FileStation.VFS.Protocol", Version: 1},
	"FileStation.VFS.User":              {API: "SYNO.FileStation.VFS.User", Version: 1},
	"FileStation.VirtualFolder":         {API: "SYNO.FileStation.VirtualFolder", Version: 2},
	"FileStation.Worm":                  {API: "SYNO.FileStation.Worm", Version: 2},
	"FileStation.Worm.Lock":             {API: "SYNO.FileStation.Worm.Lock", Version: 2},
}
