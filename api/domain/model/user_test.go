package model

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestNewUser(t *testing.T) {
	type args struct {
		id   int
		name string
	}
	tests := []struct {
		name string
		args args
		want *User
		opt  cmp.Option
	}{
		{
			name: "return User struct",
			args: args{
				id:   1,
				name: "anonymous",
			},
			want: &User{
				ID:   1,
				Name: "anonymous",
			},
			opt: cmpopts.IgnoreFields(User{}, "CreatedAt", "UpdatedAt"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUser(tt.args.id, tt.args.name); !cmp.Equal(got, tt.want, tt.opt) {
				t.Errorf("NewUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
