package virtualization

import (
	"context"
	"os"
	"reflect"
	"testing"

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

func Test_Virtualization_Image(t *testing.T) {
	c := newClient(t)

	b, err := readFile("/Users/atkini01/Downloads/nocloud_alpine-3.20.2-x86_64-bios-tiny-r0.qcow2")
	if err != nil {
		t.Error(err)
		require.NoError(t, err)
	}

	f := form.File{
		Name:    "nocloud_alpine-3.20.2-x86_64-bios-tiny-r0.qcow2",
		Content: b,
	}

	testImageUploadAndCreate(t, c, f, []string{"96914092-6022-4c21-8493-a3e1be165a1f"}, "disk")

	image := Image{
		Name: "testmantic",
		Storages: Storages{
			{Name: "default"},
		},
		Type:      "iso",
		FilePath:  "/data/cluster_storage/commoninit.iso",
		AutoClean: false,
	}

	testImageCreate(t, c, image)
	testImageDelete(t, c, image.Name)
}

func Test_Virtualization_Guest(t *testing.T) {
	v := newClient(t)

	guest := Guest{
		Name:        "test001",
		StorageName: "default",
		VcpuNum:     4,
		VramSize:    4096,
		Networks: VNICs{
			{Name: "default"},
		},
		Disks: VDisks{
			{
				CreateType: 0,
				Size:       20000,
			},
		},
	}

	g, err := v.GuestGet(context.Background(), guest)
	if err != nil {
		g, err = v.GuestCreate(context.Background(), guest)

		require.Nil(t, err)
	}

	testGuestUpdate(t, v, GuestUpdate{
		ID:        g.ID,
		Name:      guest.Name,
		IsoImages: []string{"unmounted", "unmounted"},
	})

	testGuestDelete(t, v, Guest{
		Name: guest.Name,
	})
}

func testGuestDelete(t *testing.T, v Api, guest Guest) {
	err := v.GuestDelete(context.Background(), guest)
	require.Nil(t, err)
}

func testGuestUpdate(t *testing.T, v Api, guest GuestUpdate) {
	err := v.GuestUpdate(context.Background(), guest)
	require.Nil(t, err)
}

func testImageCreate(t *testing.T, v Api, image Image) {
	got, err := v.ImageCreate(context.Background(), image)
	require.Nil(t, err)
	require.NotNil(t, got, "ImageCreate: TaskRef is nil")
}

func testImageUploadAndCreate(
	t *testing.T,
	v Api,
	file form.File,
	imageRepos []string,
	imageType string,
) {
	got, err := v.ImageUploadAndCreate(context.Background(), file, imageRepos, imageType)
	require.Nil(t, err)
	require.NotNil(t, got, "ImageUploadAndCreate: TaskRef is nil")
}

func testImageDelete(t *testing.T, v Api, imageName string) {
	err := v.ImageDelete(context.Background(), imageName)
	require.Nil(t, err)
}

func Test_Client_TaskGet(t *testing.T) {
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
		want    *Task
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.client.TaskGet(tt.args.ctx, tt.args.taskID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.TaskGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.TaskGet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Client_GuestCreate(t *testing.T) {
	type fields struct {
		client Api
	}
	type args struct {
		ctx   context.Context
		guest Guest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Guest
		wantErr bool
	}{
		{
			name: "Create guest",
			fields: fields{
				client: newClient(t),
			},
			args: args{
				ctx: context.Background(),
				guest: Guest{
					Name:        "testmantic",
					StorageName: "default",
					VcpuNum:     4,
					VramSize:    4096,
					Networks: VNICs{
						{Name: "default"},
					},
					Disks: VDisks{
						{
							CreateType: 0,
							Size:       20000,
						},
					},
				},
			},
			want: &Guest{
				ID:          "1",
				Name:        "testmantic",
				Description: "Testmantic",
				Status:      "stopped",
				StorageID:   "1",
				StorageName: "default",
				AutoRun:     0,
				VcpuNum:     1,
				VramSize:    512,
				Disks:       VDisks{},
				Networks:    VNICs{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := tt.fields.client
			got, err := v.GuestCreate(tt.args.ctx, tt.args.guest)

			require.Nil(t, err)
			require.NotNil(t, got)
		})
	}
}

func Test_Client_ListGuests(t *testing.T) {
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
		want    *GuestList
		wantErr bool
	}{
		{
			name: "List guests",
			fields: fields{
				client: newClient(t),
			},
			args: args{
				ctx: context.Background(),
			},
			want: &GuestList{
				Guests: []Guest{
					{
						ID:          "1",
						Name:        "testmantic",
						Description: "Testmantic",
						Status:      "stopped",
						StorageID:   "1",
						StorageName: "default",
						AutoRun:     0,
						VcpuNum:     1,
						VramSize:    512,
						Disks:       VDisks{},
						Networks:    VNICs{},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.client.GuestList(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.ListGuests() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.ListGuests() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewVirtualizationClient(t *testing.T) {
	type args struct {
		client Api
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
			if got := tt.args.client; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVirtualizationClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Client_StorageList(t *testing.T) {
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
		want    *StorageList
		wantErr bool
	}{
		{
			name: "List storage",
			fields: fields{
				client: newClient(t),
			},
			args: args{
				ctx: context.Background(),
			},
			want: &StorageList{
				Storages: []Storage{
					{
						ID:         "1",
						Name:       "default",
						Status:     "normal",
						HostName:   "localhost",
						HostID:     "1",
						Size:       1000000000,
						Used:       0,
						VolumePath: "/volume1",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := tt.fields.client
			got, err := v.StorageList(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.StorageList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.Equal(
				t,
				len(got.Storages),
				len(tt.want.Storages),
				"Client.StorageList() = %v, want %v",
				got,
				tt.want,
			)
		})
	}
}
