package builder

import (
	"reflect"
	"testing"
)

func TestFill(t *testing.T) {
	type args[T any] struct {
		t    T
		opts []fillerOpt[T]
	}
	tests := []struct {
		name    string
		args    args[any]
		want    any
		wantErr bool
	}{
		{
			name: "want error on scalar",
			args: args[any]{
				t:    1,
				opts: []fillerOpt[any]{},
			},
			want:    1,
			wantErr: true,
		},
		{
			name: "want error on pointer",
			args: args[any]{
				t:    &([]int{1, 2, 3}),
				opts: []fillerOpt[any]{},
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
