package client

import (
	"context"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/synology-community/synology-api/pkg/api/virtualization"
)

func Test_Virtualization_Image(t *testing.T) {
	c := newClient(t)
	v := c.Virtualization
	image := virtualization.Image{
		Name: "testmantic",
		Storages: virtualization.Storages{
			{Name: "default"},
		},
		Type:      "iso",
		FilePath:  "/data/cluster_storage/commoninit.iso",
		AutoClean: false,
	}

	testImageCreate(t, v, image)
	testImageDelete(t, v, image.Name)
}

func Test_Virtualization_Guest(t *testing.T) {
	c := newClient(t)
	v := c.Virtualization

	guest := virtualization.Guest{
		Name:        "test001",
		StorageName: "default",
		VcpuNum:     4,
		VramSize:    4096,
		Networks: virtualization.VNICs{
			{Name: "default"},
		},
		Disks: virtualization.VDisks{
			{
				CreateType: 0,
				Size:       20000,
			},
		},
	}

	g, err := v.GuestGet(context.Background(), guest)

	if err != nil {
		g, err = v.GuestCreate(context.Background(), guest)

		assert.Nil(t, err)
	}

	testGuestUpdate(t, v, virtualization.GuestUpdate{
		ID:        g.ID,
		Name:      guest.Name,
		IsoImages: []string{"unmounted", "unmounted"},
	})

	testGuestDelete(t, v, virtualization.Guest{
		Name: guest.Name,
	})
}

func testGuestGet(t *testing.T, v *virtualizationClient, guest virtualization.Guest) *virtualization.Guest {
	g, err := v.GuestGet(context.Background(), guest)
	assert.Nil(t, err)
	return g
}

func testGuestCreate(t *testing.T, v *virtualizationClient, guest virtualization.Guest) *virtualization.Guest {
	g, err := v.GuestCreate(context.Background(), guest)
	assert.Nil(t, err)
	return g
}

func testGuestDelete(t *testing.T, v *virtualizationClient, guest virtualization.Guest) {
	err := v.GuestDelete(context.Background(), guest)
	assert.Nil(t, err)
}

func testGuestUpdate(t *testing.T, v *virtualizationClient, guest virtualization.GuestUpdate) {
	err := v.GuestUpdate(context.Background(), guest)
	assert.Nil(t, err)
}

func testImageCreate(t *testing.T, v *virtualizationClient, image virtualization.Image) {
	got, err := v.ImageCreate(context.Background(), image)
	assert.Nil(t, err)
	assert.NotNil(t, got, "TaskRef is nil")
}

func testImageDelete(t *testing.T, v *virtualizationClient, imageName string) {
	err := v.ImageDelete(context.Background(), imageName)
	assert.Nil(t, err)
}

func Test_virtualizationClient_TaskGet(t *testing.T) {
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
		want    *virtualization.Task
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &virtualizationClient{
				client: tt.fields.client,
			}
			got, err := v.TaskGet(tt.args.ctx, tt.args.taskID)
			if (err != nil) != tt.wantErr {
				t.Errorf("virtualizationClient.TaskGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("virtualizationClient.TaskGet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_virtualizationClient_GuestCreate(t *testing.T) {
	type fields struct {
		client *APIClient
	}
	type args struct {
		ctx   context.Context
		guest virtualization.Guest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *virtualization.Guest
		wantErr bool
	}{
		{
			name: "Create guest",
			fields: fields{
				client: newClient(t),
			},
			args: args{
				ctx: context.Background(),
				guest: virtualization.Guest{
					Name:        "testmantic",
					StorageName: "default",
					VcpuNum:     4,
					VramSize:    4096,
					Networks: virtualization.VNICs{
						{Name: "default"},
					},
					Disks: virtualization.VDisks{
						{
							CreateType: 0,
							Size:       20000,
						},
					},
				},
			},
			want: &virtualization.Guest{
				ID:          "1",
				Name:        "testmantic",
				Description: "Testmantic",
				Status:      "stopped",
				StorageID:   "1",
				StorageName: "default",
				AutoRun:     0,
				VcpuNum:     1,
				VramSize:    512,
				Disks:       virtualization.VDisks{},
				Networks:    virtualization.VNICs{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &virtualizationClient{
				client: tt.fields.client,
			}
			got, err := v.GuestCreate(tt.args.ctx, tt.args.guest)

			assert.Nil(t, err)
			assert.NotNil(t, got)
		})
	}
}

func Test_virtualizationClient_ListGuests(t *testing.T) {
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
		want    *virtualization.GuestList
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
			want: &virtualization.GuestList{
				Guests: []virtualization.Guest{
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
						Disks:       virtualization.VDisks{},
						Networks:    virtualization.VNICs{},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.client.Virtualization.GuestList(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("virtualizationClient.ListGuests() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("virtualizationClient.ListGuests() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewVirtualizationClient(t *testing.T) {
	type args struct {
		client *APIClient
	}
	tests := []struct {
		name string
		args args
		want virtualization.VirtualizationAPI
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewVirtualizationClient(tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVirtualizationClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_virtualizationClient_StorageList(t *testing.T) {
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
		want    *virtualization.StorageList
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
			want: &virtualization.StorageList{
				Storages: []virtualization.Storage{
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
			v := tt.fields.client.VirtualizationAPI()
			got, err := v.StorageList(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("virtualizationClient.StorageList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !assert.Equal(t, len(got.Storages), len(tt.want.Storages)) {
				t.Errorf("virtualizationClient.StorageList() = %v, want %v", got, tt.want)
			}
		})
	}
}
