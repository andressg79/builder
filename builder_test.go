package builder

import (
	"reflect"
	"testing"
)

func TestFill(t *testing.T) {
	type args struct {
		t    any
		opts []fillerOpt
	}
	tests := []struct {
		name    string
		args    args
		want    any
		wantErr bool
	}{
		{
			name: "want error on scalar",
			args: args{
				t:    1,
				opts: []fillerOpt{},
			},
			want:    1,
			wantErr: true,
		},
		{
			name: "want error on pointer",
			args: args{
				t:    &([]int{1, 2, 3}),
				opts: []fillerOpt{},
			},
			want:    &([]int{1, 2, 3}),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Fill(tt.args.t, tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Fill() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Fill() = %v, want %v", got, tt.want)
			}
		})
	}
}
