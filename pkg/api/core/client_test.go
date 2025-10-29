package core

import (
	"context"
	"os"
	"path"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/synology-community/go-synology/pkg/api"
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

	if r, err := c.Login(context.Background(), api.LoginOptions{
		Username:  os.Getenv("SYNOLOGY_USER"),
		Password:  os.Getenv("SYNOLOGY_PASSWORD"),
		OTPSecret: os.Getenv("SYNOLOGY_OTP_SECRET"),
	}); err != nil {
		t.Error(err)
		require.NoError(t, err)
	} else {
		t.Log("Login successful")
		t.Logf("[INFO] Session: %s\nDeviceID: %s", r.SessionID, r.DeviceID)
	}

	return New(c)
}

func TestClient_PackageSettingGet(t *testing.T) {
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
		want    string
		wantErr bool
	}{
		{
			name: "Install ZSH Static Package",
			fields: fields{
				client: newClient(t),
			},
			args: args{
				ctx: context.Background(),
			},
			want: "zsh-static",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setting, err := tt.fields.client.PackageSettingGet(
				tt.args.ctx,
				PackageSettingGetRequest{},
			)
			if err != nil {
				t.Errorf("Client.PackageFind() error = %v", err)
				return
			}

			if setting.DefaultVol == "" {
				t.Errorf("Client.PackageSettingGet() error = %v", setting)
				return
			}
		})
	}
}

func TestClient_PackageInstall(t *testing.T) {
	type fields struct {
		client Api
	}
	type args struct {
		ctx         context.Context
		packageName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// {
		// 	name: "Install Container Manager Package",
		// 	fields: fields{
		// 		client: newClient(t),
		// 	},
		// 	args: args{
		// 		ctx:         context.Background(),
		// 		packageName: "ContainerManager",
		// 	},
		// 	want: "ContainerManager",
		// },
		{
			name: "Install ZSH Static Package",
			fields: fields{
				client: newClient(t),
			},
			args: args{
				ctx:         context.Background(),
				packageName: "zsh-static",
			},
			want: "zsh-static",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pkg, err := tt.fields.client.PackageFind(tt.args.ctx, tt.args.packageName)
			if err != nil {
				t.Errorf("Client.PackageFind() error = %v", err)
				return
			}

			size := int64(0)
			if pkg.Size != 0 {
				size = pkg.Size
			} else {
				size, err = tt.fields.client.ContentLength(context.Background(), pkg.Link)
				if err != nil {
					t.Errorf("Client.Head() error = %v", err)
					return
				}
			}

			err = tt.fields.client.PackageInstallCompound(
				tt.args.ctx,
				PackageInstallCompoundRequest{
					Name: tt.args.packageName,
					URL:  pkg.Link,
					Size: size,
				},
			)

			if (err != nil) != tt.wantErr {
				t.Errorf("Client.PackageInstall() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestClient_ContentLength(t *testing.T) {
	type fields struct {
		client Api
	}
	type args struct {
		ctx context.Context
		url string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "Test Wireguard URL",
			fields: fields{
				client: newClient(t),
			},
			args: args{
				ctx: context.Background(),
				url: "https://www.blackvoid.club/content/files/2022/08/WireGuard-geminilake-1.0.20220627.spk",
			},
			want: 1802240,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.client.ContentLength(tt.args.ctx, tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.ContentLength() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Client.ContentLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_PackageFeed(t *testing.T) {
	type fields struct {
		client Api
	}
	type args struct {
		ctx      context.Context
		feedName string
		feedUrl  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Add Package Feed",
			fields: fields{
				client: newClient(t),
			},
			args: args{
				ctx:      context.Background(),
				feedName: "Homebridge",
				feedUrl:  "https://synology.homebridge.io",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.fields.client
			err := c.PackageFeedAdd(tt.args.ctx, PackageFeedAddRequest{
				List: PackageFeedItem{
					Feed: tt.args.feedUrl,
					Name: tt.args.feedName,
				},
			})
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.PackageFeedAdd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			time.Sleep(3 * time.Second)
			p, err := c.PackageFeedList(tt.args.ctx)
			if err != nil {
				t.Errorf("Client.PackageFeedList() error = %v", err)
				return
			}
			found := false
			for _, f := range p.Items {
				if f.Name == tt.args.feedName {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Client.PackageFeedList() feed not found")
				return
			}
			time.Sleep(3 * time.Second)
			err = c.PackageFeedDelete(tt.args.ctx, PackageFeedDeleteRequest{
				List: PackageFeeds{tt.args.feedUrl},
			})
			if err != nil {
				t.Errorf("Client.PackageFeedDelete() error = %v", err)
				return
			}
		})
	}
}

func TestClient_SystemInfo(t *testing.T) {
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
		want    *SystemInfoResponse
		wantErr bool
	}{
		{
			name: "Get System Info",
			fields: fields{
				client: newClient(t),
			},
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.fields.client
			resp, err := c.SystemInfo(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.SystemInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if resp == nil {
				t.Errorf("Client.SystemInfo() = %v, want %v", resp, tt.want)
			}
		})
	}
}

func TestClient_PackageList(t *testing.T) {
	type fields struct {
		client api.Api
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *PackageListResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Client{
				client: tt.fields.client,
			}
			got, err := c.PackageList(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.PackageList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.PackageList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_PackageServerList(t *testing.T) {
	type fields struct {
		client api.Api
	}
	type args struct {
		ctx context.Context
		req PackageServerListRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *PackageServerListResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Client{
				client: tt.fields.client,
			}
			got, err := c.PackageServerList(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.PackageServerList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.PackageServerList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_PackageGet(t *testing.T) {
	type fields struct {
		client api.Api
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *PackageGetResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Client{
				client: tt.fields.client,
			}
			got, err := c.PackageGet(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.PackageGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.PackageGet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_PackageInstallStatus(t *testing.T) {
	type fields struct {
		client api.Api
	}
	type args struct {
		ctx context.Context
		req PackageInstallStatusRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *PackageInstallStatusResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Client{
				client: tt.fields.client,
			}
			got, err := c.PackageInstallStatus(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.PackageInstallStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.PackageInstallStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
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
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_PackageFind(t *testing.T) {
	type fields struct {
		client Api
	}
	type args struct {
		ctx  context.Context
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Package
		wantErr bool
	}{
		{
			name: "Find Container Manager Package",
			fields: fields{
				client: newClient(t),
			},
			args: args{
				ctx:  context.Background(),
				name: "ContainerManager",
			},
			want: &Package{
				ID:        "ContainerManager",
				Beta:      false,
				Breakpkgs: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.client.PackageFind(tt.args.ctx, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.PackageFind() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Package != tt.args.name {
				t.Errorf("Client.PackageFind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_PackageInstallUpload(t *testing.T) {
	type fields struct {
		client Api
	}
	type args struct {
		ctx      context.Context
		filePath string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		// want    *PackageInstallUploadResponse
		wantErr bool
	}{
		{
			name: "Install ZSH Static Package",
			fields: fields{
				client: newClient(t),
			},
			args: args{
				ctx:      context.Background(),
				filePath: "/Users/atkini01/src/synology-community/spksrc/packages/nomad_1.8.2_linux_amd64.spk",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := readFile(tt.args.filePath)
			if err != nil {
				t.Errorf("Client.PackageInstallUpload() error = %v", err)
				return
			}

			fileName := path.Base(tt.args.filePath)

			ctx, cancel := context.WithTimeout(tt.args.ctx, 120*time.Minute)
			defer cancel()

			c := tt.fields.client
			got, err := c.PackageInstallUpload(ctx, form.File{
				Name:    fileName,
				Content: b,
			})
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.PackageInstallUpload() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got.ID != "nomad" {
				t.Errorf("Client.PackageInstallUpload() = %v, want %v", got, "nomad")
			}
		})
	}
}

func TestClient_EventCreate(t *testing.T) {
	type fields struct {
		client Api
	}
	type args struct {
		ctx context.Context
		req EventRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Create Event",
			fields: fields{
				client: newClient(t),
			},
			args: args{
				ctx: context.Background(),
				req: EventRequest{
					Name:          "test_event_create",
					Event:         "bootup",
					Operation:     "sleep 1",
					OperationType: "script",
					Enable:        true,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.fields.client.EventCreate(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.EventCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && result == nil {
				t.Errorf("Client.EventCreate() returned nil result")
			}
		})
	}
}

func TestClient_EventUpdate(t *testing.T) {
	type fields struct {
		client Api
	}
	type args struct {
		ctx context.Context
		req EventRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Update Event",
			fields: fields{
				client: newClient(t),
			},
			args: args{
				ctx: context.Background(),
				req: EventRequest{
					Name:          "test_event_update",
					Event:         "shutdown",
					Operation:     "sleep 2",
					OperationType: "script",
					Enable:        false,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// First create the event
			_, err := tt.fields.client.EventCreate(tt.args.ctx, EventRequest{
				Name:          tt.args.req.Name,
				Event:         "bootup",
				Operation:     "sleep 1",
				OperationType: "script",
				Enable:        true,
			})
			if err != nil {
				t.Logf("Setup error creating event: %v", err)
			}

			// Then update it
			result, err := tt.fields.client.EventUpdate(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.EventUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && result == nil {
				t.Errorf("Client.EventUpdate() returned nil result")
			}

			// Cleanup
			_ = tt.fields.client.EventDelete(tt.args.ctx, EventRequest{Name: tt.args.req.Name})
		})
	}
}

func TestClient_EventGet(t *testing.T) {
	type fields struct {
		client Api
	}
	type args struct {
		ctx  context.Context
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Get Event",
			fields: fields{
				client: newClient(t),
			},
			args: args{
				ctx:  context.Background(),
				name: "test_event_get",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// First create the event
			_, err := tt.fields.client.EventCreate(tt.args.ctx, EventRequest{
				Name:          tt.args.name,
				Event:         "bootup",
				Operation:     "sleep 1",
				OperationType: "script",
				Enable:        true,
			})
			if err != nil {
				t.Logf("Setup error creating event: %v", err)
			}

			// Then get it
			result, err := tt.fields.client.EventGet(tt.args.ctx, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.EventGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if result == nil {
					t.Errorf("Client.EventGet() returned nil result")
				} else if result.Name != tt.args.name {
					t.Errorf("Client.EventGet() returned wrong name: got %v, want %v", result.Name, tt.args.name)
				}
			}

			// Cleanup
			_ = tt.fields.client.EventDelete(tt.args.ctx, EventRequest{Name: tt.args.name})
		})
	}
}

func TestClient_EventRun(t *testing.T) {
	type fields struct {
		client Api
	}
	type args struct {
		ctx  context.Context
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Run Event",
			fields: fields{
				client: newClient(t),
			},
			args: args{
				ctx:  context.Background(),
				name: "test_event_run",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// First create the event
			_, err := tt.fields.client.EventCreate(tt.args.ctx, EventRequest{
				Name:          tt.args.name,
				Event:         "bootup",
				Operation:     "echo 'test'",
				OperationType: "script",
				Enable:        true,
			})
			if err != nil {
				t.Logf("Setup error creating event: %v", err)
			}

			// Then run it
			err = tt.fields.client.EventRun(tt.args.ctx, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.EventRun() error = %v, wantErr %v", err, tt.wantErr)
			}

			// Cleanup
			_ = tt.fields.client.EventDelete(tt.args.ctx, EventRequest{Name: tt.args.name})
		})
	}
}

func TestClient_EventDelete(t *testing.T) {
	type fields struct {
		client Api
	}
	type args struct {
		ctx context.Context
		req EventRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Delete Event",
			fields: fields{
				client: newClient(t),
			},
			args: args{
				ctx: context.Background(),
				req: EventRequest{
					Name: "test_event_delete",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// First create the event
			_, err := tt.fields.client.EventCreate(tt.args.ctx, EventRequest{
				Name:          tt.args.req.Name,
				Event:         "bootup",
				Operation:     "sleep 1",
				OperationType: "script",
				Enable:        true,
			})
			if err != nil {
				t.Logf("Setup error creating event: %v", err)
			}

			// Then delete it
			err = tt.fields.client.EventDelete(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.EventDelete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_RootEventCreate(t *testing.T) {
	type fields struct {
		client Api
	}
	type args struct {
		ctx context.Context
		req EventRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Create Root Event",
			fields: fields{
				client: newClient(t),
			},
			args: args{
				ctx: context.Background(),
				req: EventRequest{
					Name:          "test_root_event_create",
					Event:         "bootup",
					Operation:     "echo \"Root event created\"; sleep 1",
					OperationType: "script",
					Enable:        true,
					Owner: map[string]string{
						"0": "root",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.fields.client.RootEventCreate(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.RootEventCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && result == nil {
				t.Errorf("Client.RootEventCreate() returned nil result")
			}

			// Cleanup
			_ = tt.fields.client.RootEventDelete(tt.args.ctx, EventRequest{Name: tt.args.req.Name})
		})
	}
}

func TestClient_RootEventUpdate(t *testing.T) {
	type fields struct {
		client Api
	}
	type args struct {
		ctx context.Context
		req EventRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Update Root Event",
			fields: fields{
				client: newClient(t),
			},
			args: args{
				ctx: context.Background(),
				req: EventRequest{
					Name:          "test_root_event_update",
					Event:         "shutdown",
					Operation:     "sleep 2",
					OperationType: "script",
					Enable:        false,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// First create the root event
			_, err := tt.fields.client.RootEventCreate(tt.args.ctx, EventRequest{
				Name:          tt.args.req.Name,
				Event:         "bootup",
				Operation:     "sleep 1",
				OperationType: "script",
				Enable:        true,
			})
			if err != nil {
				t.Logf("Setup error creating root event: %v", err)
			}

			// Then update it
			result, err := tt.fields.client.RootEventUpdate(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.RootEventUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && result == nil {
				t.Errorf("Client.RootEventUpdate() returned nil result")
			}

			// Cleanup
			_ = tt.fields.client.RootEventDelete(tt.args.ctx, EventRequest{Name: tt.args.req.Name})
		})
	}
}

func TestClient_RootEventDelete(t *testing.T) {
	type fields struct {
		client Api
	}
	type args struct {
		ctx context.Context
		req EventRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Delete Root Event",
			fields: fields{
				client: newClient(t),
			},
			args: args{
				ctx: context.Background(),
				req: EventRequest{
					Name: "test_root_event_delete",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// First create the root event
			_, err := tt.fields.client.RootEventCreate(tt.args.ctx, EventRequest{
				Name:          tt.args.req.Name,
				Event:         "bootup",
				Operation:     "sleep 1",
				OperationType: "script",
				Enable:        true,
			})
			if err != nil {
				t.Logf("Setup error creating root event: %v", err)
			}

			// Then delete it
			err = tt.fields.client.RootEventDelete(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.RootEventDelete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_UserCreate(t *testing.T) {
	type fields struct {
		client Api
	}
	type args struct {
		ctx context.Context
		req UserCreateRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "Create user",
			fields: fields{client: newClient(t)},
			args: args{
				ctx: context.Background(),
				req: UserCreateRequest{
					Name:        "test_api_user",
					Password:    "TestPassword123!",
					Email:       "testuser@example.com",
					Description: "Test user for creation test",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createResp, err := tt.fields.client.UserCreate(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Fatalf("UserCreate error = %v, wantErr %v", err, tt.wantErr)
			}
			require.Equal(t, tt.args.req.Name, createResp.User.Name)

			// Cleanup
			delReq := UserDeleteRequest{Name: tt.args.req.Name}
			_, _ = tt.fields.client.UserDelete(tt.args.ctx, delReq)
		})
	}
}

func TestClient_UserModify(t *testing.T) {
	type fields struct {
		client Api
	}
	type args struct {
		ctx       context.Context
		createReq UserCreateRequest
		req       UserModifyRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "Modify user",
			fields: fields{client: newClient(t)},
			args: args{
				ctx: context.Background(),
				createReq: UserCreateRequest{
					Name:        "test_api_user_mod",
					Password:    "TestPassword123!",
					Email:       "testuser@example.com",
					Description: "Test user for modify test",
				},
				req: UserModifyRequest{
					Name:        "test_api_user_mod",
					Description: "Updated description",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup: create user
			_, err := tt.fields.client.UserCreate(tt.args.ctx, tt.args.createReq)
			if err != nil {
				t.Fatalf("Setup UserCreate failed: %v", err)
			}

			modResp, err := tt.fields.client.UserModify(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Fatalf("UserModify error = %v, wantErr %v", err, tt.wantErr)
			}
			require.Equal(t, "Updated description", modResp.User.Description)

			// Cleanup
			delReq := UserDeleteRequest{Name: tt.args.createReq.Name}
			_, _ = tt.fields.client.UserDelete(tt.args.ctx, delReq)
		})
	}
}

func TestClient_UserDelete(t *testing.T) {
	type fields struct {
		client Api
	}
	type args struct {
		ctx       context.Context
		createReq UserCreateRequest
		req       UserDeleteRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "Delete user",
			fields: fields{client: newClient(t)},
			args: args{
				ctx: context.Background(),
				createReq: UserCreateRequest{
					Name:        "test_api_user_del",
					Password:    "TestPassword123!",
					Email:       "testuser@example.com",
					Description: "Test user for delete test",
				},
				req: UserDeleteRequest{Name: "test_api_user_del"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup: create user
			_, err := tt.fields.client.UserCreate(tt.args.ctx, tt.args.createReq)
			if err != nil {
				t.Fatalf("Setup UserCreate failed: %v", err)
			}

			delResp, err := tt.fields.client.UserDelete(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Fatalf("UserDelete error = %v, wantErr %v", err, tt.wantErr)
			}
			require.True(t, delResp.Success, "UserDelete did not succeed")
		})
	}
}

func TestClient_GroupCreate(t *testing.T) {
	type fields struct {
		client Api
	}
	type args struct {
		ctx context.Context
		req GroupCreateRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "Create group",
			fields: fields{client: newClient(t)},
			args: args{
				ctx: context.Background(),
				req: GroupCreateRequest{
					Name:        "test_api_group",
					Description: "Test group for creation test",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createResp, err := tt.fields.client.GroupCreate(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Fatalf("GroupCreate error = %v, wantErr %v", err, tt.wantErr)
			}
			require.Equal(t, tt.args.req.Name, createResp.Group.Name)

			// Cleanup
			delReq := GroupDeleteRequest{Name: tt.args.req.Name}
			_, _ = tt.fields.client.GroupDelete(tt.args.ctx, delReq)
		})
	}
}

func TestClient_GroupModify(t *testing.T) {
	type fields struct {
		client Api
	}
	type args struct {
		ctx       context.Context
		createReq GroupCreateRequest
		req       GroupModifyRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "Modify group",
			fields: fields{client: newClient(t)},
			args: args{
				ctx: context.Background(),
				createReq: GroupCreateRequest{
					Name:        "test_api_group_mod",
					Description: "Test group for modify test",
				},
				req: GroupModifyRequest{
					Name:        "test_api_group_mod",
					Description: "Updated group description",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup: create group
			_, err := tt.fields.client.GroupCreate(tt.args.ctx, tt.args.createReq)
			if err != nil {
				t.Fatalf("Setup GroupCreate failed: %v", err)
			}

			modResp, err := tt.fields.client.GroupModify(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Fatalf("GroupModify error = %v, wantErr %v", err, tt.wantErr)
			}
			require.Equal(t, "Updated group description", modResp.Group.Description)

			// Cleanup
			delReq := GroupDeleteRequest{Name: tt.args.createReq.Name}
			_, _ = tt.fields.client.GroupDelete(tt.args.ctx, delReq)
		})
	}
}

func TestClient_GroupDelete(t *testing.T) {
	type fields struct {
		client Api
	}
	type args struct {
		ctx       context.Context
		createReq GroupCreateRequest
		req       GroupDeleteRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "Delete group",
			fields: fields{client: newClient(t)},
			args: args{
				ctx: context.Background(),
				createReq: GroupCreateRequest{
					Name:        "test_api_group_del",
					Description: "Test group for delete test",
				},
				req: GroupDeleteRequest{Name: "test_api_group_del"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup: create group
			_, err := tt.fields.client.GroupCreate(tt.args.ctx, tt.args.createReq)
			if err != nil {
				t.Fatalf("Setup GroupCreate failed: %v", err)
			}

			delResp, err := tt.fields.client.GroupDelete(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Fatalf("GroupDelete error = %v, wantErr %v", err, tt.wantErr)
			}
			require.True(t, delResp.Success, "GroupDelete did not succeed")
		})
	}
}
