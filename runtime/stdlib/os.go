package stdlib

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/malivvan/cui/runtime"
)

var osModule = map[string]runtime.Object{
	"o_rdonly":            &runtime.Int{Value: int64(os.O_RDONLY)},
	"o_wronly":            &runtime.Int{Value: int64(os.O_WRONLY)},
	"o_rdwr":              &runtime.Int{Value: int64(os.O_RDWR)},
	"o_append":            &runtime.Int{Value: int64(os.O_APPEND)},
	"o_create":            &runtime.Int{Value: int64(os.O_CREATE)},
	"o_excl":              &runtime.Int{Value: int64(os.O_EXCL)},
	"o_sync":              &runtime.Int{Value: int64(os.O_SYNC)},
	"o_trunc":             &runtime.Int{Value: int64(os.O_TRUNC)},
	"mode_dir":            &runtime.Int{Value: int64(os.ModeDir)},
	"mode_append":         &runtime.Int{Value: int64(os.ModeAppend)},
	"mode_exclusive":      &runtime.Int{Value: int64(os.ModeExclusive)},
	"mode_temporary":      &runtime.Int{Value: int64(os.ModeTemporary)},
	"mode_symlink":        &runtime.Int{Value: int64(os.ModeSymlink)},
	"mode_device":         &runtime.Int{Value: int64(os.ModeDevice)},
	"mode_named_pipe":     &runtime.Int{Value: int64(os.ModeNamedPipe)},
	"mode_socket":         &runtime.Int{Value: int64(os.ModeSocket)},
	"mode_setuid":         &runtime.Int{Value: int64(os.ModeSetuid)},
	"mode_setgui":         &runtime.Int{Value: int64(os.ModeSetgid)},
	"mode_char_device":    &runtime.Int{Value: int64(os.ModeCharDevice)},
	"mode_sticky":         &runtime.Int{Value: int64(os.ModeSticky)},
	"mode_type":           &runtime.Int{Value: int64(os.ModeType)},
	"mode_perm":           &runtime.Int{Value: int64(os.ModePerm)},
	"path_separator":      &runtime.Char{Value: os.PathSeparator},
	"path_list_separator": &runtime.Char{Value: os.PathListSeparator},
	"dev_null":            &runtime.String{Value: os.DevNull},
	"seek_set":            &runtime.Int{Value: int64(io.SeekStart)},
	"seek_cur":            &runtime.Int{Value: int64(io.SeekCurrent)},
	"seek_end":            &runtime.Int{Value: int64(io.SeekEnd)},
	"args": &runtime.BuiltinFunction{
		Name:  "args",
		Value: osArgs,
	}, // args() => array(string)
	"chdir": &runtime.BuiltinFunction{
		Name:  "chdir",
		Value: FuncASRE(os.Chdir),
	}, // chdir(dir string) => error
	"chmod": osFuncASFmRE("chmod", os.Chmod), // chmod(name string, mode int) => error
	"chown": &runtime.BuiltinFunction{
		Name:  "chown",
		Value: FuncASIIRE(os.Chown),
	}, // chown(name string, uid int, gid int) => error
	"clearenv": &runtime.BuiltinFunction{
		Name:  "clearenv",
		Value: FuncAR(os.Clearenv),
	}, // clearenv()
	"environ": &runtime.BuiltinFunction{
		Name:  "environ",
		Value: FuncARSs(os.Environ),
	}, // environ() => array(string)
	"exit": &runtime.BuiltinFunction{
		Name:  "exit",
		Value: FuncAIR(os.Exit),
	}, // exit(code int)
	"expand_env": &runtime.BuiltinFunction{
		Name:  "expand_env",
		Value: osExpandEnv,
	}, // expand_env(s string) => string
	"getegid": &runtime.BuiltinFunction{
		Name:  "getegid",
		Value: FuncARI(os.Getegid),
	}, // getegid() => int
	"getenv": &runtime.BuiltinFunction{
		Name:  "getenv",
		Value: FuncASRS(os.Getenv),
	}, // getenv(s string) => string
	"geteuid": &runtime.BuiltinFunction{
		Name:  "geteuid",
		Value: FuncARI(os.Geteuid),
	}, // geteuid() => int
	"getgid": &runtime.BuiltinFunction{
		Name:  "getgid",
		Value: FuncARI(os.Getgid),
	}, // getgid() => int
	"getgroups": &runtime.BuiltinFunction{
		Name:  "getgroups",
		Value: FuncARIsE(os.Getgroups),
	}, // getgroups() => array(string)/error
	"getpagesize": &runtime.BuiltinFunction{
		Name:  "getpagesize",
		Value: FuncARI(os.Getpagesize),
	}, // getpagesize() => int
	"getpid": &runtime.BuiltinFunction{
		Name:  "getpid",
		Value: FuncARI(os.Getpid),
	}, // getpid() => int
	"getppid": &runtime.BuiltinFunction{
		Name:  "getppid",
		Value: FuncARI(os.Getppid),
	}, // getppid() => int
	"getuid": &runtime.BuiltinFunction{
		Name:  "getuid",
		Value: FuncARI(os.Getuid),
	}, // getuid() => int
	"getwd": &runtime.BuiltinFunction{
		Name:  "getwd",
		Value: FuncARSE(os.Getwd),
	}, // getwd() => string/error
	"hostname": &runtime.BuiltinFunction{
		Name:  "hostname",
		Value: FuncARSE(os.Hostname),
	}, // hostname() => string/error
	"lchown": &runtime.BuiltinFunction{
		Name:  "lchown",
		Value: FuncASIIRE(os.Lchown),
	}, // lchown(name string, uid int, gid int) => error
	"link": &runtime.BuiltinFunction{
		Name:  "link",
		Value: FuncASSRE(os.Link),
	}, // link(oldname string, newname string) => error
	"lookup_env": &runtime.BuiltinFunction{
		Name:  "lookup_env",
		Value: osLookupEnv,
	}, // lookup_env(key string) => string/false
	"mkdir":     osFuncASFmRE("mkdir", os.Mkdir),        // mkdir(name string, perm int) => error
	"mkdir_all": osFuncASFmRE("mkdir_all", os.MkdirAll), // mkdir_all(name string, perm int) => error
	"readlink": &runtime.BuiltinFunction{
		Name:  "readlink",
		Value: FuncASRSE(os.Readlink),
	}, // readlink(name string) => string/error
	"remove": &runtime.BuiltinFunction{
		Name:  "remove",
		Value: FuncASRE(os.Remove),
	}, // remove(name string) => error
	"remove_all": &runtime.BuiltinFunction{
		Name:  "remove_all",
		Value: FuncASRE(os.RemoveAll),
	}, // remove_all(name string) => error
	"rename": &runtime.BuiltinFunction{
		Name:  "rename",
		Value: FuncASSRE(os.Rename),
	}, // rename(oldpath string, newpath string) => error
	"setenv": &runtime.BuiltinFunction{
		Name:  "setenv",
		Value: FuncASSRE(os.Setenv),
	}, // setenv(key string, value string) => error
	"symlink": &runtime.BuiltinFunction{
		Name:  "symlink",
		Value: FuncASSRE(os.Symlink),
	}, // symlink(oldname string newname string) => error
	"temp_dir": &runtime.BuiltinFunction{
		Name:  "temp_dir",
		Value: FuncARS(os.TempDir),
	}, // temp_dir() => string
	"truncate": &runtime.BuiltinFunction{
		Name:  "truncate",
		Value: FuncASI64RE(os.Truncate),
	}, // truncate(name string, size int) => error
	"unsetenv": &runtime.BuiltinFunction{
		Name:  "unsetenv",
		Value: FuncASRE(os.Unsetenv),
	}, // unsetenv(key string) => error
	"create": &runtime.BuiltinFunction{
		Name:  "create",
		Value: osCreate,
	}, // create(name string) => imap(file)/error
	"open": &runtime.BuiltinFunction{
		Name:  "open",
		Value: osOpen,
	}, // open(name string) => imap(file)/error
	"open_file": &runtime.BuiltinFunction{
		Name:  "open_file",
		Value: osOpenFile,
	}, // open_file(name string, flag int, perm int) => imap(file)/error
	"find_process": &runtime.BuiltinFunction{
		Name:  "find_process",
		Value: osFindProcess,
	}, // find_process(pid int) => imap(process)/error
	"start_process": &runtime.BuiltinFunction{
		Name:  "start_process",
		Value: osStartProcess,
	}, // start_process(name string, argv array(string), dir string, env array(string)) => imap(process)/error
	"exec_look_path": &runtime.BuiltinFunction{
		Name:  "exec_look_path",
		Value: FuncASRSE(exec.LookPath),
	}, // exec_look_path(file) => string/error
	"exec": &runtime.BuiltinFunction{
		Name:  "exec",
		Value: osExec,
	}, // exec(name, args...) => command
	"stat": &runtime.BuiltinFunction{
		Name:  "stat",
		Value: osStat,
	}, // stat(name) => imap(fileinfo)/error
	"read_file": &runtime.BuiltinFunction{
		Name:  "read_file",
		Value: osReadFile,
	}, // readfile(name) => array(byte)/error
}

func osReadFile(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
	if len(args) != 1 {
		return nil, runtime.ErrWrongNumArguments
	}
	fname, ok := runtime.ToString(args[0])
	if !ok {
		return nil, runtime.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	bytes, err := os.ReadFile(fname)
	if err != nil {
		return wrapError(err), nil
	}
	if len(bytes) > runtime.MaxBytesLen {
		return nil, runtime.ErrBytesLimit
	}
	return &runtime.Bytes{Value: bytes}, nil
}

func osStat(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
	if len(args) != 1 {
		return nil, runtime.ErrWrongNumArguments
	}
	fname, ok := runtime.ToString(args[0])
	if !ok {
		return nil, runtime.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	stat, err := os.Stat(fname)
	if err != nil {
		return wrapError(err), nil
	}
	fstat := &runtime.ImmutableMap{
		Value: map[string]runtime.Object{
			"name":  &runtime.String{Value: stat.Name()},
			"mtime": &runtime.Time{Value: stat.ModTime()},
			"size":  &runtime.Int{Value: stat.Size()},
			"mode":  &runtime.Int{Value: int64(stat.Mode())},
		},
	}
	if stat.IsDir() {
		fstat.Value["directory"] = runtime.TrueValue
	} else {
		fstat.Value["directory"] = runtime.FalseValue
	}
	return fstat, nil
}

func osCreate(ctx context.Context, args ...runtime.Object) (runtime.Object, error) {
	if len(args) != 1 {
		return nil, runtime.ErrWrongNumArguments
	}
	s1, ok := runtime.ToString(args[0])
	if !ok {
		return nil, runtime.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	res, err := os.Create(s1)
	if err != nil {
		return wrapError(err), nil
	}
	return makeOSFile(res), nil
}

func osOpen(ctx context.Context, args ...runtime.Object) (runtime.Object, error) {
	if len(args) != 1 {
		return nil, runtime.ErrWrongNumArguments
	}
	s1, ok := runtime.ToString(args[0])
	if !ok {
		return nil, runtime.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	res, err := os.Open(s1)
	if err != nil {
		return wrapError(err), nil
	}
	return makeOSFile(res), nil
}

func osOpenFile(ctx context.Context, args ...runtime.Object) (runtime.Object, error) {
	if len(args) != 3 {
		return nil, runtime.ErrWrongNumArguments
	}
	s1, ok := runtime.ToString(args[0])
	if !ok {
		return nil, runtime.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	i2, ok := runtime.ToInt(args[1])
	if !ok {
		return nil, runtime.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
	}
	i3, ok := runtime.ToInt(args[2])
	if !ok {
		return nil, runtime.ErrInvalidArgumentType{
			Name:     "third",
			Expected: "int(compatible)",
			Found:    args[2].TypeName(),
		}
	}
	res, err := os.OpenFile(s1, i2, os.FileMode(i3))
	if err != nil {
		return wrapError(err), nil
	}
	return makeOSFile(res), nil
}

func osArgs(ctx context.Context, args ...runtime.Object) (runtime.Object, error) {
	vm := ctx.Value(runtime.ContextKey("vm")).(*runtime.VM)
	if len(args) != 0 {
		return nil, runtime.ErrWrongNumArguments
	}
	arr := &runtime.Array{}
	for _, osArg := range vm.Args {
		if len(osArg) > runtime.MaxStringLen {
			return nil, runtime.ErrStringLimit
		}
		arr.Value = append(arr.Value, &runtime.String{Value: osArg})
	}
	return arr, nil
}

func osFuncASFmRE(
	name string,
	fn func(string, os.FileMode) error,
) *runtime.BuiltinFunction {
	return &runtime.BuiltinFunction{
		Name: name,
		Value: func(ctx context.Context, args ...runtime.Object) (runtime.Object, error) {
			if len(args) != 2 {
				return nil, runtime.ErrWrongNumArguments
			}
			s1, ok := runtime.ToString(args[0])
			if !ok {
				return nil, runtime.ErrInvalidArgumentType{
					Name:     "first",
					Expected: "string(compatible)",
					Found:    args[0].TypeName(),
				}
			}
			i2, ok := runtime.ToInt64(args[1])
			if !ok {
				return nil, runtime.ErrInvalidArgumentType{
					Name:     "second",
					Expected: "int(compatible)",
					Found:    args[1].TypeName(),
				}
			}
			return wrapError(fn(s1, os.FileMode(i2))), nil
		},
	}
}

func osLookupEnv(ctx context.Context, args ...runtime.Object) (runtime.Object, error) {
	if len(args) != 1 {
		return nil, runtime.ErrWrongNumArguments
	}
	s1, ok := runtime.ToString(args[0])
	if !ok {
		return nil, runtime.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	res, ok := os.LookupEnv(s1)
	if !ok {
		return runtime.FalseValue, nil
	}
	if len(res) > runtime.MaxStringLen {
		return nil, runtime.ErrStringLimit
	}
	return &runtime.String{Value: res}, nil
}

func osExpandEnv(ctx context.Context, args ...runtime.Object) (runtime.Object, error) {
	if len(args) != 1 {
		return nil, runtime.ErrWrongNumArguments
	}
	s1, ok := runtime.ToString(args[0])
	if !ok {
		return nil, runtime.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	var vlen int
	var failed bool
	s := os.Expand(s1, func(k string) string {
		if failed {
			return ""
		}
		v := os.Getenv(k)

		// this does not count the other texts that are not being replaced
		// but the code checks the final length at the end
		vlen += len(v)
		if vlen > runtime.MaxStringLen {
			failed = true
			return ""
		}
		return v
	})
	if failed || len(s) > runtime.MaxStringLen {
		return nil, runtime.ErrStringLimit
	}
	return &runtime.String{Value: s}, nil
}

func osExec(ctx context.Context, args ...runtime.Object) (runtime.Object, error) {
	if len(args) == 0 {
		return nil, runtime.ErrWrongNumArguments
	}
	name, ok := runtime.ToString(args[0])
	if !ok {
		return nil, runtime.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	var execArgs []string
	for idx, arg := range args[1:] {
		execArg, ok := runtime.ToString(arg)
		if !ok {
			return nil, runtime.ErrInvalidArgumentType{
				Name:     fmt.Sprintf("args[%d]", idx),
				Expected: "string(compatible)",
				Found:    args[1+idx].TypeName(),
			}
		}
		execArgs = append(execArgs, execArg)
	}
	return makeOSExecCommand(exec.Command(name, execArgs...)), nil
}

func osFindProcess(ctx context.Context, args ...runtime.Object) (runtime.Object, error) {
	if len(args) != 1 {
		return nil, runtime.ErrWrongNumArguments
	}
	i1, ok := runtime.ToInt(args[0])
	if !ok {
		return nil, runtime.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	proc, err := os.FindProcess(i1)
	if err != nil {
		return wrapError(err), nil
	}
	return makeOSProcess(proc), nil
}

func osStartProcess(ctx context.Context, args ...runtime.Object) (runtime.Object, error) {
	if len(args) != 4 {
		return nil, runtime.ErrWrongNumArguments
	}
	name, ok := runtime.ToString(args[0])
	if !ok {
		return nil, runtime.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	var argv []string
	var err error
	switch arg1 := args[1].(type) {
	case *runtime.Array:
		argv, err = stringArray(arg1.Value, "second")
		if err != nil {
			return nil, err
		}
	case *runtime.ImmutableArray:
		argv, err = stringArray(arg1.Value, "second")
		if err != nil {
			return nil, err
		}
	default:
		return nil, runtime.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "array",
			Found:    arg1.TypeName(),
		}
	}

	dir, ok := runtime.ToString(args[2])
	if !ok {
		return nil, runtime.ErrInvalidArgumentType{
			Name:     "third",
			Expected: "string(compatible)",
			Found:    args[2].TypeName(),
		}
	}

	var env []string
	switch arg3 := args[3].(type) {
	case *runtime.Array:
		env, err = stringArray(arg3.Value, "fourth")
		if err != nil {
			return nil, err
		}
	case *runtime.ImmutableArray:
		env, err = stringArray(arg3.Value, "fourth")
		if err != nil {
			return nil, err
		}
	default:
		return nil, runtime.ErrInvalidArgumentType{
			Name:     "fourth",
			Expected: "array",
			Found:    arg3.TypeName(),
		}
	}

	proc, err := os.StartProcess(name, argv, &os.ProcAttr{
		Dir: dir,
		Env: env,
	})
	if err != nil {
		return wrapError(err), nil
	}
	return makeOSProcess(proc), nil
}

func stringArray(arr []runtime.Object, argName string) ([]string, error) {
	var sarr []string
	for idx, elem := range arr {
		str, ok := elem.(*runtime.String)
		if !ok {
			return nil, runtime.ErrInvalidArgumentType{
				Name:     fmt.Sprintf("%s[%d]", argName, idx),
				Expected: "string",
				Found:    elem.TypeName(),
			}
		}
		sarr = append(sarr, str.Value)
	}
	return sarr, nil
}
