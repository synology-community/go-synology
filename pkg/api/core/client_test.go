package core

import (
	"context"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/synology-community/go-synology/pkg/api"
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

			setting, err := tt.fields.client.PackageSettingGet(tt.args.ctx, PackageSettingGetRequest{})

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

			err = tt.fields.client.PackageInstallCompound(tt.args.ctx, PackageInstallCompoundRequest{
				Name: tt.args.packageName,
				URL:  pkg.Link,
				Size: size,
			})

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
