package filestation

import (
	"context"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/synology-community/go-synology/pkg/api"
	"github.com/synology-community/go-synology/pkg/models"
	"github.com/synology-community/go-synology/pkg/util/form"
)

func newClient(t *testing.T) Api {
	c, err := api.New(api.Options{
		Host: os.Getenv("SYNOLOGY_HOST"),
	})

	if err != nil {
		t.Error(err)
		require.NoError(t, err)
	}

	if r, err := c.Login(context.Background(), os.Getenv("SYNOLOGY_USER"), os.Getenv("SYNOLOGY_PASSWORD")); err != nil {
		t.Error(err)
		require.NoError(t, err)
	} else {
		t.Log("Login successful")
		t.Logf("[INFO] Session: %s\nDeviceID: %s", r.SessionID, r.DeviceID)
	}

	return New(c)
}

func Test_Client_List(t *testing.T) {
	type fields struct {
		client Api
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
			_, err := tt.fields.client.List(tt.args.ctx, tt.args.folderPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_Client_Delete(t *testing.T) {
	type fields struct {
		client Api
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
		want    *DeleteStatusResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.client.Delete(tt.args.ctx, tt.args.paths, tt.args.accurateProgress)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Client_DeleteStart(t *testing.T) {
	type fields struct {
		client Api
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
		want    *DeleteStartResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.client.DeleteStart(tt.args.ctx, tt.args.paths, tt.args.accurateProgress)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.DeleteStart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.DeleteStart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Client_DeleteStatus(t *testing.T) {
	type fields struct {
		client Api
	}
	type args struct {
		ctx    context.Context
		taskID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *DeleteStatusResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.client.DeleteStatus(tt.args.ctx, tt.args.taskID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.DeleteStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.DeleteStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Client_MD5Start(t *testing.T) {
	t.Run("Upload", Test_Client_Upload)
	type fields struct {
		client Api
	}
	type args struct {
		ctx  context.Context
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *MD5StartResponse
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
			got, err := tt.fields.client.MD5Start(tt.args.ctx, tt.args.path)
			require.NoError(t, err)
			require.NotEmpty(t, got.TaskID)
		})
	}
}

func Test_Client_MD5Status(t *testing.T) {
	type fields struct {
		client Api
	}
	type args struct {
		ctx    context.Context
		taskID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *MD5StatusResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.client.MD5Status(tt.args.ctx, tt.args.taskID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.MD5Status() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.MD5Status() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Client_Download(t *testing.T) {
	type fields struct {
		client Api
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
			got, err := tt.fields.client.Download(tt.args.ctx, tt.args.path, tt.args.mode)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Download() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.Download() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Client_Rename(t *testing.T) {
	type fields struct {
		client Api
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
			got, err := tt.fields.client.Rename(tt.args.ctx, tt.args.path, tt.args.name, tt.args.newName)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Rename() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.Rename() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Client_CreateFolder(t *testing.T) {
	type fields struct {
		client Api
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
			got, err := tt.fields.client.CreateFolder(tt.args.ctx, tt.args.paths, tt.args.names, tt.args.forceParent)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.CreateFolder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.CreateFolder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Client_ListShares(t *testing.T) {
	type fields struct {
		client Api
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
			got, err := tt.fields.client.ListShares(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.ListShares() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.ListShares() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Client_MD5(t *testing.T) {
	t.Run("Upload", Test_Client_Upload)
	type fields struct {
		client Api
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
			got, err := tt.fields.client.MD5(tt.args.ctx, tt.args.path)
			require.NoError(t, err)
			require.NotEmpty(t, got.MD5)
		})
	}
}

func Test_Client_Upload(t *testing.T) {
	type fields struct {
		client Api
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
		want    *UploadResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.client.Upload(tt.args.ctx, tt.args.path, tt.args.file, tt.args.createParents, tt.args.overwrite)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Upload() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.Upload() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewClient(t *testing.T) {
	type args struct {
		client api.Api
	}
	tests := []struct {
		name string
		args args
		want Api
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
