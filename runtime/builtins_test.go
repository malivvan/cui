package runtime_test

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/malivvan/cui/runtime"
)

func Test_builtinDelete(t *testing.T) {
	var builtinDelete func(ctx context.Context, args ...runtime.Object) (runtime.Object, error)
	for _, f := range runtime.GetAllBuiltinFunctions() {
		if f.Name == "delete" {
			builtinDelete = f.Value
			break
		}
	}
	if builtinDelete == nil {
		t.Fatal("builtin delete not found")
	}
	type args struct {
		args []runtime.Object
	}
	tests := []struct {
		name      string
		args      args
		want      runtime.Object
		wantErr   bool
		wantedErr error
		target    interface{}
	}{
		{name: "invalid-arg", args: args{[]runtime.Object{&runtime.String{},
			&runtime.String{}}}, wantErr: true,
			wantedErr: runtime.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "map",
				Found:    "string"},
		},
		{name: "no-args",
			wantErr: true, wantedErr: runtime.ErrWrongNumArguments},
		{name: "empty-args", args: args{[]runtime.Object{}}, wantErr: true,
			wantedErr: runtime.ErrWrongNumArguments,
		},
		{name: "3-args", args: args{[]runtime.Object{
			(*runtime.Map)(nil), (*runtime.String)(nil), (*runtime.String)(nil)}},
			wantErr: true, wantedErr: runtime.ErrWrongNumArguments,
		},
		{name: "nil-map-empty-key",
			args: args{[]runtime.Object{&runtime.Map{}, &runtime.String{}}},
			want: runtime.UndefinedValue,
		},
		{name: "nil-map-nonstr-key",
			args: args{[]runtime.Object{
				&runtime.Map{}, &runtime.Int{}}}, wantErr: true,
			wantedErr: runtime.ErrInvalidArgumentType{
				Name: "second", Expected: "string", Found: "int"},
		},
		{name: "nil-map-no-key",
			args: args{[]runtime.Object{&runtime.Map{}}}, wantErr: true,
			wantedErr: runtime.ErrWrongNumArguments,
		},
		{name: "map-missing-key",
			args: args{
				[]runtime.Object{
					&runtime.Map{Value: map[string]runtime.Object{
						"key": &runtime.String{Value: "value"},
					}},
					&runtime.String{Value: "key1"}}},
			want: runtime.UndefinedValue,
			target: &runtime.Map{
				Value: map[string]runtime.Object{
					"key": &runtime.String{
						Value: "value"}}},
		},
		{name: "map-emptied",
			args: args{
				[]runtime.Object{
					&runtime.Map{Value: map[string]runtime.Object{
						"key": &runtime.String{Value: "value"},
					}},
					&runtime.String{Value: "key"}}},
			want:   runtime.UndefinedValue,
			target: &runtime.Map{Value: map[string]runtime.Object{}},
		},
		{name: "map-multi-keys",
			args: args{
				[]runtime.Object{
					&runtime.Map{Value: map[string]runtime.Object{
						"key1": &runtime.String{Value: "value1"},
						"key2": &runtime.Int{Value: 10},
					}},
					&runtime.String{Value: "key1"}}},
			want: runtime.UndefinedValue,
			target: &runtime.Map{Value: map[string]runtime.Object{
				"key2": &runtime.Int{Value: 10}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := builtinDelete(context.Background(), tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("builtinDelete() error = %v, wantErr %v",
					err, tt.wantErr)
				return
			}
			if tt.wantErr && !errors.Is(err, tt.wantedErr) {
				if err.Error() != tt.wantedErr.Error() {
					t.Errorf("builtinDelete() error = %v, wantedErr %v",
						err, tt.wantedErr)
					return
				}
			}
			if got != tt.want {
				t.Errorf("builtinDelete() = %v, want %v", got, tt.want)
				return
			}
			if !tt.wantErr && tt.target != nil {
				switch v := tt.args.args[0].(type) {
				case *runtime.Map, *runtime.Array:
					if !reflect.DeepEqual(tt.target, tt.args.args[0]) {
						t.Errorf("builtinDelete() objects are not equal "+
							"got: %+v, want: %+v", tt.args.args[0], tt.target)
					}
				default:
					t.Errorf("builtinDelete() unsuporrted arg[0] type %s",
						v.TypeName())
					return
				}
			}
		})
	}
}

func Test_builtinSplice(t *testing.T) {
	var builtinSplice func(ctx context.Context, args ...runtime.Object) (runtime.Object, error)
	for _, f := range runtime.GetAllBuiltinFunctions() {
		if f.Name == "splice" {
			builtinSplice = f.Value
			break
		}
	}
	if builtinSplice == nil {
		t.Fatal("builtin splice not found")
	}
	tests := []struct {
		name      string
		args      []runtime.Object
		deleted   runtime.Object
		Array     *runtime.Array
		wantErr   bool
		wantedErr error
	}{
		{name: "no args", args: []runtime.Object{}, wantErr: true,
			wantedErr: runtime.ErrWrongNumArguments,
		},
		{name: "invalid args", args: []runtime.Object{&runtime.Map{}},
			wantErr: true,
			wantedErr: runtime.ErrInvalidArgumentType{
				Name: "first", Expected: "array", Found: "map"},
		},
		{name: "invalid args",
			args:    []runtime.Object{&runtime.Array{}, &runtime.String{}},
			wantErr: true,
			wantedErr: runtime.ErrInvalidArgumentType{
				Name: "second", Expected: "int", Found: "string"},
		},
		{name: "negative index",
			args:      []runtime.Object{&runtime.Array{}, &runtime.Int{Value: -1}},
			wantErr:   true,
			wantedErr: runtime.ErrIndexOutOfBounds},
		{name: "non int count",
			args: []runtime.Object{
				&runtime.Array{}, &runtime.Int{Value: 0},
				&runtime.String{Value: ""}},
			wantErr: true,
			wantedErr: runtime.ErrInvalidArgumentType{
				Name: "third", Expected: "int", Found: "string"},
		},
		{name: "negative count",
			args: []runtime.Object{
				&runtime.Array{Value: []runtime.Object{
					&runtime.Int{Value: 0},
					&runtime.Int{Value: 1},
					&runtime.Int{Value: 2}}},
				&runtime.Int{Value: 0},
				&runtime.Int{Value: -1}},
			wantErr:   true,
			wantedErr: runtime.ErrIndexOutOfBounds,
		},
		{name: "insert with zero count",
			args: []runtime.Object{
				&runtime.Array{Value: []runtime.Object{
					&runtime.Int{Value: 0},
					&runtime.Int{Value: 1},
					&runtime.Int{Value: 2}}},
				&runtime.Int{Value: 0},
				&runtime.Int{Value: 0},
				&runtime.String{Value: "b"}},
			deleted: &runtime.Array{Value: []runtime.Object{}},
			Array: &runtime.Array{Value: []runtime.Object{
				&runtime.String{Value: "b"},
				&runtime.Int{Value: 0},
				&runtime.Int{Value: 1},
				&runtime.Int{Value: 2}}},
		},
		{name: "insert",
			args: []runtime.Object{
				&runtime.Array{Value: []runtime.Object{
					&runtime.Int{Value: 0},
					&runtime.Int{Value: 1},
					&runtime.Int{Value: 2}}},
				&runtime.Int{Value: 1},
				&runtime.Int{Value: 0},
				&runtime.String{Value: "c"},
				&runtime.String{Value: "d"}},
			deleted: &runtime.Array{Value: []runtime.Object{}},
			Array: &runtime.Array{Value: []runtime.Object{
				&runtime.Int{Value: 0},
				&runtime.String{Value: "c"},
				&runtime.String{Value: "d"},
				&runtime.Int{Value: 1},
				&runtime.Int{Value: 2}}},
		},
		{name: "insert with zero count",
			args: []runtime.Object{
				&runtime.Array{Value: []runtime.Object{
					&runtime.Int{Value: 0},
					&runtime.Int{Value: 1},
					&runtime.Int{Value: 2}}},
				&runtime.Int{Value: 1},
				&runtime.Int{Value: 0},
				&runtime.String{Value: "c"},
				&runtime.String{Value: "d"}},
			deleted: &runtime.Array{Value: []runtime.Object{}},
			Array: &runtime.Array{Value: []runtime.Object{
				&runtime.Int{Value: 0},
				&runtime.String{Value: "c"},
				&runtime.String{Value: "d"},
				&runtime.Int{Value: 1},
				&runtime.Int{Value: 2}}},
		},
		{name: "insert with delete",
			args: []runtime.Object{
				&runtime.Array{Value: []runtime.Object{
					&runtime.Int{Value: 0},
					&runtime.Int{Value: 1},
					&runtime.Int{Value: 2}}},
				&runtime.Int{Value: 1},
				&runtime.Int{Value: 1},
				&runtime.String{Value: "c"},
				&runtime.String{Value: "d"}},
			deleted: &runtime.Array{
				Value: []runtime.Object{&runtime.Int{Value: 1}}},
			Array: &runtime.Array{Value: []runtime.Object{
				&runtime.Int{Value: 0},
				&runtime.String{Value: "c"},
				&runtime.String{Value: "d"},
				&runtime.Int{Value: 2}}},
		},
		{name: "insert with delete multi",
			args: []runtime.Object{
				&runtime.Array{Value: []runtime.Object{
					&runtime.Int{Value: 0},
					&runtime.Int{Value: 1},
					&runtime.Int{Value: 2}}},
				&runtime.Int{Value: 1},
				&runtime.Int{Value: 2},
				&runtime.String{Value: "c"},
				&runtime.String{Value: "d"}},
			deleted: &runtime.Array{Value: []runtime.Object{
				&runtime.Int{Value: 1},
				&runtime.Int{Value: 2}}},
			Array: &runtime.Array{
				Value: []runtime.Object{
					&runtime.Int{Value: 0},
					&runtime.String{Value: "c"},
					&runtime.String{Value: "d"}}},
		},
		{name: "delete all with positive count",
			args: []runtime.Object{
				&runtime.Array{Value: []runtime.Object{
					&runtime.Int{Value: 0},
					&runtime.Int{Value: 1},
					&runtime.Int{Value: 2}}},
				&runtime.Int{Value: 0},
				&runtime.Int{Value: 3}},
			deleted: &runtime.Array{Value: []runtime.Object{
				&runtime.Int{Value: 0},
				&runtime.Int{Value: 1},
				&runtime.Int{Value: 2}}},
			Array: &runtime.Array{Value: []runtime.Object{}},
		},
		{name: "delete all with big count",
			args: []runtime.Object{
				&runtime.Array{Value: []runtime.Object{
					&runtime.Int{Value: 0},
					&runtime.Int{Value: 1},
					&runtime.Int{Value: 2}}},
				&runtime.Int{Value: 0},
				&runtime.Int{Value: 5}},
			deleted: &runtime.Array{Value: []runtime.Object{
				&runtime.Int{Value: 0},
				&runtime.Int{Value: 1},
				&runtime.Int{Value: 2}}},
			Array: &runtime.Array{Value: []runtime.Object{}},
		},
		{name: "nothing2",
			args: []runtime.Object{
				&runtime.Array{Value: []runtime.Object{
					&runtime.Int{Value: 0},
					&runtime.Int{Value: 1},
					&runtime.Int{Value: 2}}}},
			Array: &runtime.Array{Value: []runtime.Object{}},
			deleted: &runtime.Array{Value: []runtime.Object{
				&runtime.Int{Value: 0},
				&runtime.Int{Value: 1},
				&runtime.Int{Value: 2}}},
		},
		{name: "pop without count",
			args: []runtime.Object{
				&runtime.Array{Value: []runtime.Object{
					&runtime.Int{Value: 0},
					&runtime.Int{Value: 1},
					&runtime.Int{Value: 2}}},
				&runtime.Int{Value: 2}},
			deleted: &runtime.Array{Value: []runtime.Object{&runtime.Int{Value: 2}}},
			Array: &runtime.Array{Value: []runtime.Object{
				&runtime.Int{Value: 0}, &runtime.Int{Value: 1}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := builtinSplice(context.Background(), tt.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("builtinSplice() error = %v, wantErr %v",
					err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.deleted) {
				t.Errorf("builtinSplice() = %v, want %v", got, tt.deleted)
			}
			if tt.wantErr && tt.wantedErr.Error() != err.Error() {
				t.Errorf("builtinSplice() error = %v, wantedErr %v",
					err, tt.wantedErr)
			}
			if tt.Array != nil && !reflect.DeepEqual(tt.Array, tt.args[0]) {
				t.Errorf("builtinSplice() arrays are not equal expected"+
					" %s, got %s", tt.Array, tt.args[0].(*runtime.Array))
			}
		})
	}
}

func Test_builtinRange(t *testing.T) {
	var builtinRange func(ctx context.Context, args ...runtime.Object) (runtime.Object, error)
	for _, f := range runtime.GetAllBuiltinFunctions() {
		if f.Name == "range" {
			builtinRange = f.Value
			break
		}
	}
	if builtinRange == nil {
		t.Fatal("builtin range not found")
	}
	tests := []struct {
		name      string
		args      []runtime.Object
		result    *runtime.Array
		wantErr   bool
		wantedErr error
	}{
		{name: "no args", args: []runtime.Object{}, wantErr: true,
			wantedErr: runtime.ErrWrongNumArguments,
		},
		{name: "single args", args: []runtime.Object{&runtime.Map{}},
			wantErr:   true,
			wantedErr: runtime.ErrWrongNumArguments,
		},
		{name: "4 args", args: []runtime.Object{&runtime.Map{}, &runtime.String{}, &runtime.String{}, &runtime.String{}},
			wantErr:   true,
			wantedErr: runtime.ErrWrongNumArguments,
		},
		{name: "invalid start",
			args:    []runtime.Object{&runtime.String{}, &runtime.String{}},
			wantErr: true,
			wantedErr: runtime.ErrInvalidArgumentType{
				Name: "start", Expected: "int", Found: "string"},
		},
		{name: "invalid stop",
			args:    []runtime.Object{&runtime.Int{}, &runtime.String{}},
			wantErr: true,
			wantedErr: runtime.ErrInvalidArgumentType{
				Name: "stop", Expected: "int", Found: "string"},
		},
		{name: "invalid step",
			args:    []runtime.Object{&runtime.Int{}, &runtime.Int{}, &runtime.String{}},
			wantErr: true,
			wantedErr: runtime.ErrInvalidArgumentType{
				Name: "step", Expected: "int", Found: "string"},
		},
		{name: "zero step",
			args:      []runtime.Object{&runtime.Int{}, &runtime.Int{}, &runtime.Int{}}, //must greate than 0
			wantErr:   true,
			wantedErr: runtime.ErrInvalidRangeStep,
		},
		{name: "negative step",
			args:      []runtime.Object{&runtime.Int{}, &runtime.Int{}, intObject(-2)}, //must greate than 0
			wantErr:   true,
			wantedErr: runtime.ErrInvalidRangeStep,
		},
		{name: "same bound",
			args:    []runtime.Object{&runtime.Int{}, &runtime.Int{}},
			wantErr: false,
			result: &runtime.Array{
				Value: nil,
			},
		},
		{name: "positive range",
			args:    []runtime.Object{&runtime.Int{}, &runtime.Int{Value: 5}},
			wantErr: false,
			result: &runtime.Array{
				Value: []runtime.Object{
					intObject(0),
					intObject(1),
					intObject(2),
					intObject(3),
					intObject(4),
				},
			},
		},
		{name: "negative range",
			args:    []runtime.Object{&runtime.Int{}, &runtime.Int{Value: -5}},
			wantErr: false,
			result: &runtime.Array{
				Value: []runtime.Object{
					intObject(0),
					intObject(-1),
					intObject(-2),
					intObject(-3),
					intObject(-4),
				},
			},
		},

		{name: "positive with step",
			args:    []runtime.Object{&runtime.Int{}, &runtime.Int{Value: 5}, &runtime.Int{Value: 2}},
			wantErr: false,
			result: &runtime.Array{
				Value: []runtime.Object{
					intObject(0),
					intObject(2),
					intObject(4),
				},
			},
		},

		{name: "negative with step",
			args:    []runtime.Object{&runtime.Int{}, &runtime.Int{Value: -10}, &runtime.Int{Value: 2}},
			wantErr: false,
			result: &runtime.Array{
				Value: []runtime.Object{
					intObject(0),
					intObject(-2),
					intObject(-4),
					intObject(-6),
					intObject(-8),
				},
			},
		},

		{name: "large range",
			args:    []runtime.Object{intObject(-10), intObject(10), &runtime.Int{Value: 3}},
			wantErr: false,
			result: &runtime.Array{
				Value: []runtime.Object{
					intObject(-10),
					intObject(-7),
					intObject(-4),
					intObject(-1),
					intObject(2),
					intObject(5),
					intObject(8),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := builtinRange(context.Background(), tt.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("builtinRange() error = %v, wantErr %v",
					err, tt.wantErr)
				return
			}
			if tt.wantErr && tt.wantedErr.Error() != err.Error() {
				t.Errorf("builtinRange() error = %v, wantedErr %v",
					err, tt.wantedErr)
			}
			if tt.result != nil && !reflect.DeepEqual(tt.result, got) {
				t.Errorf("builtinRange() arrays are not equal expected"+
					" %s, got %s", tt.result, got.(*runtime.Array))
			}
		})
	}
}
