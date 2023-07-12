package str

import (
	"reflect"
	"testing"
)

func TestSurround(t *testing.T) {
	type args struct {
		str   string
		left  string
		right string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Surrounds a string with two other strings",
			args: args{
				str:   "foo",
				left:  "bar",
				right: "baz",
			},
			want: "barfoobaz",
		},
		{
			name: "Surrounds a string with two \"",
			args: args{
				str:   "bar",
				left:  "\"",
				right: "\"",
			},
			want: "\"bar\"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Surround(tt.args.str, tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("Surround() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSurroundAll(t *testing.T) {
	type args struct {
		list  []string
		left  string
		right string
	}
	tests := []struct {
		name       string
		args       args
		wantResult []string
	}{
		{
			name:       "Surrounds all strings in a list with two other strings",
			args:       args{list: []string{"foo", "bar", "baz"}, left: "bar", right: "baz"},
			wantResult: []string{"barfoobaz", "barbarbaz", "barbazbaz"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := SurroundAll(tt.args.list, tt.args.left, tt.args.right); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("SurroundAll() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestTrimAll(t *testing.T) {
	type args struct {
		list   []string
		cutset string
	}
	tests := []struct {
		name       string
		args       args
		wantResult []string
	}{
		{
			name: "Trims all strings in a list with a cutset",
			args: args{list: []string{"foo", "bar", "baz", "bazba"}, cutset: "ba"},
			wantResult: []string{ "foo", "r", "z", "z" },
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := TrimAll(tt.args.list, tt.args.cutset); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("TrimAll() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
