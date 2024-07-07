package pointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrapListPointer(t *testing.T) {
	type args struct {
		list []string
	}
	s := "a"
	s2 := "b"
	s3 := "c"
	tests := []struct {
		name string
		args args
		want []*string
	}{
		{
			name: "",
			args: args{[]string{s, s2, s3}},
			want: []*string{&s, &s2, &s3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, WrapListPointer(tt.args.list), "WrapListPointer(%v)", tt.args.list)
		})
	}
}

func TestUnwrapListPointer(t *testing.T) {
	type args struct {
		list []*string
	}
	s := "a"
	s2 := "b"
	s3 := "c"
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "",
			args: args{[]*string{&s, &s2, &s3}},
			want: []string{"a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, UnwrapListPointer(tt.args.list), "UnwrapListPointer(%v)", tt.args.list)
		})
	}
}

type A struct {
}

func TestWrap(t *testing.T) {
	a := "a"
	assert.Equal(t, &a, Wrap(a))
	num := 123
	assert.Equal(t, &num, Wrap(num))
	s := []string{"1"}
	assert.Equal(t, &s, Wrap(s))
	obj := A{}
	assert.Equal(t, &obj, Wrap(obj))
}
