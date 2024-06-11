package core

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/synology-community/synology-api/pkg/api"
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
		{
			name: "Install Container Manager Package",
			fields: fields{
				client: newClient(t),
			},
			args: args{
				ctx:         context.Background(),
				packageName: "ContainerManager",
			},
			want: "ContainerManager",
		},
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
			packageLookup := map[string]Package{}

			_, err := tt.fields.client.PackageGet(tt.args.ctx, tt.args.packageName)
			if err == nil {
				t.Skip("Package already installed")
			}

			packages, err := tt.fields.client.PackageServerList(tt.args.ctx, PackageServerListRequest{
				ForceReload: false,
				LoadOthers:  false,
			})

			if err != nil {
				t.Errorf("Client.PackageServerList() error = %v", err)
				return
			}

			for _, p := range packages.Packages {
				packageLookup[p.Package] = p
			}

			otherPackages, err := tt.fields.client.PackageServerList(tt.args.ctx, PackageServerListRequest{
				ForceReload: false,
				LoadOthers:  true,
			})

			if err != nil {
				t.Errorf("Client.PackageServerList() error = %v", err)
				return
			}

			for _, p := range otherPackages.Packages {
				if _, ok := packageLookup[p.Package]; !ok {
					packageLookup[p.Package] = p
				}
			}

			if pkg, ok := packageLookup[tt.args.packageName]; ok {
				size, err := tt.fields.client.ContentLength(context.Background(), pkg.Link)
				if err != nil {
					t.Errorf("Client.Head() error = %v", err)
					return
				}

				dlRes, err := tt.fields.client.PackageInstall(tt.args.ctx, PackageInstallRequest{
					Name:       pkg.Package,
					URL:        pkg.Link,
					Type:       0,
					BigInstall: false,
					FileSize:   size,
				})
				if err != nil {
					t.Errorf("Client.PackageInstall() error = %v", err)
					return
				}

				if dlRes.TaskID == "" {
					t.Errorf("Client.PackageInstall() success = %v", dlRes.TaskID)
					return
				}

				status := new(PackageInstallStatusResponse)

				for retry := 0; !status.Finished; retry++ {
					status, err = tt.fields.client.PackageInstallStatus(tt.args.ctx, PackageInstallStatusRequest{
						TaskID: dlRes.TaskID,
					})

					if err != nil {
						t.Errorf("Client.PackageInstallStatus() error = %v", err)
						return
					}

					if status.Finished {
						t.Logf("Package installed: %s", status.Name)
						break
					}

					if retry > 10 {
						t.Errorf("Client.PackageInstallStatus() retry = %d", retry)
						return
					}

					if !status.Finished {
						time.Sleep(2 * time.Second)
					}
				}

				path := fmt.Sprintf("%s/%s", status.TmpFolder, status.Taskid)

				instRes, err := tt.fields.client.PackageInstall(tt.args.ctx, PackageInstallRequest{
					Name:              status.Name,
					Path:              path,
					InstallRunPackage: true,
					Force:             true,
					CheckCodesign:     false,
					Type:              0,
					ExtraValues:       "{}",
				})

				if err != nil {
					t.Errorf("Client.PackageInstall() error = %v", err)
					return
				}

				if instRes.TaskID == "" {
					t.Errorf("Client.PackageInstall() success = %v", instRes.TaskID)
					return
				}

				status, err = tt.fields.client.PackageInstallStatus(tt.args.ctx, PackageInstallStatusRequest{
					TaskID: instRes.TaskID,
				})

				if err != nil {
					t.Errorf("Client.PackageInstallStatus() error = %v", err)
					return
				}

				if status.Finished {
					t.Logf("Package installed: %s", status.Name)
				} else {
					t.Errorf("Client.PackageInstallStatus() status = %s", status.Status)
				}

				require.Equal(t, tt.want, status.Name)
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

func TestClient_SystemInfo(t *testing.T) {
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
		want    *SystemInfoResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Client{
				client: tt.fields.client,
			}
			got, err := c.SystemInfo(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.SystemInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.SystemInfo() = %v, want %v", got, tt.want)
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
