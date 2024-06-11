package virtualization

import (
	"net/url"
	"testing"

	"github.com/google/go-querystring/query"
)

func TestStorages_EncodeValues(t *testing.T) {
	type args struct {
		k string
		v *url.Values
	}
	tests := []struct {
		name    string
		s       Storages
		args    args
		wantErr bool
	}{
		{
			name: "Test Storages EncodeValues",
			s: Storages{
				{
					ID:   "test1",
					Name: "test1",
				},
				{
					ID:   "test2",
					Name: "test2",
				},
			},
			args: args{
				k: "test",
				v: &url.Values{
					"storage_ids":   []string{"['test1','test2']"},
					"storage_names": []string{"['test1','test2']"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.EncodeValues(tt.args.k, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("Storages.EncodeValues() error = %v, wantErr %v", err, tt.wantErr)
			}
		})

		s := Image{
			ID:          "test",
			Name:        "test",
			Description: "test",
			Storages: Storages{
				{ID: "test1", Name: "test1"},
				{ID: "test2", Name: "test2"},
			},
		}

		if v, err := query.Values(s); err != nil {
			t.Errorf("Storages.EncodeValues() error = %v", err)
		} else {
			if v.Get("storage_ids") != "['test1','test2']" {
				t.Errorf("Storages.EncodeValues() error = %v", v.Get("storage_ids"))
			}

			if v.Get("storage_names") != "['test1','test2']" {
				t.Errorf("Storages.EncodeValues() error = %v", v.Get("storage_names"))
			}
		}
	}
}
