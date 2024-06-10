package client

import (
	"context"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/synology-community/synology-api/pkg/api/filestation"
	"github.com/synology-community/synology-api/pkg/models"
	"github.com/synology-community/synology-api/pkg/util/form"
)

func Test_fileStationClient_List(t *testing.T) {
	type fields struct {
		client *APIClient
	}
	type args struct {
		ctx        context.Context
		folderPath string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.FileList
		wantErr bool
	}{
		{
			name: "List",
			fields: fields{
				client: newClient(t),
			},
			args: args{
				ctx:        context.Background(),
				folderPath: "/data",
			},
			want:    &models.FileList{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &fileStationClient{
				client: tt.fields.client,
			}
			_, err := f.List(tt.args.ctx, tt.args.folderPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("fileStationClient.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_fileStationClient_Delete(t *testing.T) {
	type fields struct {
		client *APIClient
	}
	type args struct {
		ctx              context.Context
		paths            []string
		accurateProgress bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *filestation.DeleteStatusResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &fileStationClient{
				client: tt.fields.client,
			}
			got, err := f.Delete(tt.args.ctx, tt.args.paths, tt.args.accurateProgress)
			if (err != nil) != tt.wantErr {
				t.Errorf("fileStationClient.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fileStationClient.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fileStationClient_DeleteStart(t *testing.T) {
	type fields struct {
		client *APIClient
	}
	type args struct {
		ctx              context.Context
		paths            []string
		accurateProgress bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *filestation.DeleteStartResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &fileStationClient{
				client: tt.fields.client,
			}
			got, err := f.DeleteStart(tt.args.ctx, tt.args.paths, tt.args.accurateProgress)
			if (err != nil) != tt.wantErr {
				t.Errorf("fileStationClient.DeleteStart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fileStationClient.DeleteStart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fileStationClient_DeleteStatus(t *testing.T) {
	type fields struct {
		client *APIClient
	}
	type args struct {
		ctx    context.Context
		taskID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *filestation.DeleteStatusResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &fileStationClient{
				client: tt.fields.client,
			}
			got, err := f.DeleteStatus(tt.args.ctx, tt.args.taskID)
			if (err != nil) != tt.wantErr {
				t.Errorf("fileStationClient.DeleteStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fileStationClient.DeleteStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fileStationClient_MD5Start(t *testing.T) {
	t.Run("Upload", Test_FileStationClient_Upload)
	type fields struct {
		client *APIClient
	}
	type args struct {
		ctx  context.Context
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *filestation.MD5StartResponse
		wantErr bool
	}{
		{
			name: "MD5Start",
			fields: fields{
				client: newClient(t),
			},
			args: args{
				ctx:  context.Background(),
				path: "/data/foo/test.txt",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &fileStationClient{
				client: tt.fields.client,
			}
			got, err := f.MD5Start(tt.args.ctx, tt.args.path)
			require.NoError(t, err)
			require.NotEmpty(t, got.TaskID)
		})
	}
}

func Test_fileStationClient_MD5Status(t *testing.T) {
	type fields struct {
		client *APIClient
	}
	type args struct {
		ctx    context.Context
		taskID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *filestation.MD5StatusResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &fileStationClient{
				client: tt.fields.client,
			}
			got, err := f.MD5Status(tt.args.ctx, tt.args.taskID)
			if (err != nil) != tt.wantErr {
				t.Errorf("fileStationClient.MD5Status() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fileStationClient.MD5Status() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fileStationClient_Download(t *testing.T) {
	type fields struct {
		client *APIClient
	}
	type args struct {
		ctx  context.Context
		path string
		mode string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *form.File
		wantErr bool
	}{
		{
			name: "Download",
			fields: fields{
				client: newClient(t),
			},
			args: args{
				ctx:  context.Background(),
				path: "/data/foo/test.txt",
				mode: "download",
			},
			want: &form.File{
				Name:    "download",
				Content: "Hello, World!",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &fileStationClient{
				client: tt.fields.client,
			}
			got, err := f.Download(tt.args.ctx, tt.args.path, tt.args.mode)
			if (err != nil) != tt.wantErr {
				t.Errorf("fileStationClient.Download() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fileStationClient.Download() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fileStationClient_Rename(t *testing.T) {
	type fields struct {
		client *APIClient
	}
	type args struct {
		ctx     context.Context
		path    string
		name    string
		newName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.FileList
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &fileStationClient{
				client: tt.fields.client,
			}
			got, err := f.Rename(tt.args.ctx, tt.args.path, tt.args.name, tt.args.newName)
			if (err != nil) != tt.wantErr {
				t.Errorf("fileStationClient.Rename() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fileStationClient.Rename() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fileStationClient_CreateFolder(t *testing.T) {
	type fields struct {
		client *APIClient
	}
	type args struct {
		ctx         context.Context
		paths       []string
		names       []string
		forceParent bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.FolderList
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &fileStationClient{
				client: tt.fields.client,
			}
			got, err := f.CreateFolder(tt.args.ctx, tt.args.paths, tt.args.names, tt.args.forceParent)
			if (err != nil) != tt.wantErr {
				t.Errorf("fileStationClient.CreateFolder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fileStationClient.CreateFolder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fileStationClient_ListShares(t *testing.T) {
	type fields struct {
		client *APIClient
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.ShareList
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &fileStationClient{
				client: tt.fields.client,
			}
			got, err := f.ListShares(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("fileStationClient.ListShares() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fileStationClient.ListShares() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fileStationClient_MD5(t *testing.T) {
	t.Run("Upload", Test_FileStationClient_Upload)
	type fields struct {
		client *APIClient
	}
	type args struct {
		ctx  context.Context
		path string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "MD5",
			fields: fields{
				client: newClient(t),
			},
			args: args{
				ctx:  context.Background(),
				path: "/data/foo/test.txt",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &fileStationClient{
				client: tt.fields.client,
			}
			got, err := f.MD5(tt.args.ctx, tt.args.path)
			require.NoError(t, err)
			require.NotEmpty(t, got.MD5)
		})
	}
}

func Test_fileStationClient_Upload(t *testing.T) {
	type fields struct {
		client *APIClient
	}
	type args struct {
		ctx           context.Context
		path          string
		file          form.File
		createParents bool
		overwrite     bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *filestation.UploadResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &fileStationClient{
				client: tt.fields.client,
			}
			got, err := f.Upload(tt.args.ctx, tt.args.path, tt.args.file, tt.args.createParents, tt.args.overwrite)
			if (err != nil) != tt.wantErr {
				t.Errorf("fileStationClient.Upload() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fileStationClient.Upload() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFileStationClient(t *testing.T) {
	type args struct {
		client *APIClient
	}
	tests := []struct {
		name string
		args args
		want filestation.FileStationApi
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFileStationClient(tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFileStationClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
