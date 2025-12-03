package stdlib

import (
	"context"
	"time"

	"github.com/malivvan/cui/runtime"
)

var timesModule = map[string]runtime.Object{
	"format_ansic":        &runtime.String{Value: time.ANSIC},
	"format_unix_date":    &runtime.String{Value: time.UnixDate},
	"format_ruby_date":    &runtime.String{Value: time.RubyDate},
	"format_rfc822":       &runtime.String{Value: time.RFC822},
	"format_rfc822z":      &runtime.String{Value: time.RFC822Z},
	"format_rfc850":       &runtime.String{Value: time.RFC850},
	"format_rfc1123":      &runtime.String{Value: time.RFC1123},
	"format_rfc1123z":     &runtime.String{Value: time.RFC1123Z},
	"format_rfc3339":      &runtime.String{Value: time.RFC3339},
	"format_rfc3339_nano": &runtime.String{Value: time.RFC3339Nano},
	"format_kitchen":      &runtime.String{Value: time.Kitchen},
	"format_stamp":        &runtime.String{Value: time.Stamp},
	"format_stamp_milli":  &runtime.String{Value: time.StampMilli},
	"format_stamp_micro":  &runtime.String{Value: time.StampMicro},
	"format_stamp_nano":   &runtime.String{Value: time.StampNano},
	"nanosecond":          &runtime.Int{Value: int64(time.Nanosecond)},
	"microsecond":         &runtime.Int{Value: int64(time.Microsecond)},
	"millisecond":         &runtime.Int{Value: int64(time.Millisecond)},
	"second":              &runtime.Int{Value: int64(time.Second)},
	"minute":              &runtime.Int{Value: int64(time.Minute)},
	"hour":                &runtime.Int{Value: int64(time.Hour)},
	"january":             &runtime.Int{Value: int64(time.January)},
	"february":            &runtime.Int{Value: int64(time.February)},
	"march":               &runtime.Int{Value: int64(time.March)},
	"april":               &runtime.Int{Value: int64(time.April)},
	"may":                 &runtime.Int{Value: int64(time.May)},
	"june":                &runtime.Int{Value: int64(time.June)},
	"july":                &runtime.Int{Value: int64(time.July)},
	"august":              &runtime.Int{Value: int64(time.August)},
	"september":           &runtime.Int{Value: int64(time.September)},
	"october":             &runtime.Int{Value: int64(time.October)},
	"november":            &runtime.Int{Value: int64(time.November)},
	"december":            &runtime.Int{Value: int64(time.December)},
	"sleep": &runtime.BuiltinFunction{
		Name:  "sleep",
		Value: timesSleep,
	}, // sleep(int)
	"parse_duration": &runtime.BuiltinFunction{
		Name:  "parse_duration",
		Value: timesParseDuration,
	}, // parse_duration(str) => int
	"since": &runtime.BuiltinFunction{
		Name:  "since",
		Value: timesSince,
	}, // since(time) => int
	"until": &runtime.BuiltinFunction{
		Name:  "until",
		Value: timesUntil,
	}, // until(time) => int
	"duration_hours": &runtime.BuiltinFunction{
		Name:  "duration_hours",
		Value: timesDurationHours,
	}, // duration_hours(int) => float
	"duration_minutes": &runtime.BuiltinFunction{
		Name:  "duration_minutes",
		Value: timesDurationMinutes,
	}, // duration_minutes(int) => float
	"duration_nanoseconds": &runtime.BuiltinFunction{
		Name:  "duration_nanoseconds",
		Value: timesDurationNanoseconds,
	}, // duration_nanoseconds(int) => int
	"duration_seconds": &runtime.BuiltinFunction{
		Name:  "duration_seconds",
		Value: timesDurationSeconds,
	}, // duration_seconds(int) => float
	"duration_string": &runtime.BuiltinFunction{
		Name:  "duration_string",
		Value: timesDurationString,
	}, // duration_string(int) => string
	"month_string": &runtime.BuiltinFunction{
		Name:  "month_string",
		Value: timesMonthString,
	}, // month_string(int) => string
	"date": &runtime.BuiltinFunction{
		Name:  "date",
		Value: timesDate,
	}, // date(year, month, day, hour, min, sec, nsec) => time
	"now": &runtime.BuiltinFunction{
		Name:  "now",
		Value: timesNow,
	}, // now() => time
	"parse": &runtime.BuiltinFunction{
		Name:  "parse",
		Value: timesParse,
	}, // parse(format, str) => time
	"unix": &runtime.BuiltinFunction{
		Name:  "unix",
		Value: timesUnix,
	}, // unix(sec, nsec) => time
	"add": &runtime.BuiltinFunction{
		Name:  "add",
		Value: timesAdd,
	}, // add(time, int) => time
	"add_date": &runtime.BuiltinFunction{
		Name:  "add_date",
		Value: timesAddDate,
	}, // add_date(time, years, months, days) => time
	"sub": &runtime.BuiltinFunction{
		Name:  "sub",
		Value: timesSub,
	}, // sub(t time, u time) => int
	"after": &runtime.BuiltinFunction{
		Name:  "after",
		Value: timesAfter,
	}, // after(t time, u time) => bool
	"before": &runtime.BuiltinFunction{
		Name:  "before",
		Value: timesBefore,
	}, // before(t time, u time) => bool
	"time_year": &runtime.BuiltinFunction{
		Name:  "time_year",
		Value: timesTimeYear,
	}, // time_year(time) => int
	"time_month": &runtime.BuiltinFunction{
		Name:  "time_month",
		Value: timesTimeMonth,
	}, // time_month(time) => int
	"time_day": &runtime.BuiltinFunction{
		Name:  "time_day",
		Value: timesTimeDay,
	}, // time_day(time) => int
	"time_weekday": &runtime.BuiltinFunction{
		Name:  "time_weekday",
		Value: timesTimeWeekday,
	}, // time_weekday(time) => int
	"time_hour": &runtime.BuiltinFunction{
		Name:  "time_hour",
		Value: timesTimeHour,
	}, // time_hour(time) => int
	"time_minute": &runtime.BuiltinFunction{
		Name:  "time_minute",
		Value: timesTimeMinute,
	}, // time_minute(time) => int
	"time_second": &runtime.BuiltinFunction{
		Name:  "time_second",
		Value: timesTimeSecond,
	}, // time_second(time) => int
	"time_nanosecond": &runtime.BuiltinFunction{
		Name:  "time_nanosecond",
		Value: timesTimeNanosecond,
	}, // time_nanosecond(time) => int
	"time_unix": &runtime.BuiltinFunction{
		Name:  "time_unix",
		Value: timesTimeUnix,
	}, // time_unix(time) => int
	"time_unix_nano": &runtime.BuiltinFunction{
		Name:  "time_unix_nano",
		Value: timesTimeUnixNano,
	}, // time_unix_nano(time) => int
	"time_format": &runtime.BuiltinFunction{
		Name:  "time_format",
		Value: timesTimeFormat,
	}, // time_format(time, format) => string
	"time_location": &runtime.BuiltinFunction{
		Name:  "time_location",
		Value: timesTimeLocation,
	}, // time_location(time) => string
	"time_string": &runtime.BuiltinFunction{
		Name:  "time_string",
		Value: timesTimeString,
	}, // time_string(time) => string
	"is_zero": &runtime.BuiltinFunction{
		Name:  "is_zero",
		Value: timesIsZero,
	}, // is_zero(time) => bool
	"to_local": &runtime.BuiltinFunction{
		Name:  "to_local",
		Value: timesToLocal,
	}, // to_local(time) => time
	"to_utc": &runtime.BuiltinFunction{
		Name:  "to_utc",
		Value: timesToUTC,
	}, // to_utc(time) => time
}

func timesSleep(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
	if len(args) != 1 {
		err = runtime.ErrWrongNumArguments
		return
	}

	i1, ok := runtime.ToInt64(args[0])
	if !ok {
		err = runtime.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}
	ret = runtime.UndefinedValue
	if time.Duration(i1) <= time.Second {
		time.Sleep(time.Duration(i1))
		return
	}

	done := make(chan struct{})
	go func() {
		time.Sleep(time.Duration(i1))
		select {
		case <-ctx.Done():
		case done <- struct{}{}:
		}
	}()

	select {
	case <-ctx.Done():
		return nil, runtime.ErrVMAborted
	case <-done:
	}
	return
}

func timesParseDuration(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
	if len(args) != 1 {
		err = runtime.ErrWrongNumArguments
		return
	}

	s1, ok := runtime.ToString(args[0])
	if !ok {
		err = runtime.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	dur, err := time.ParseDuration(s1)
	if err != nil {
		ret = wrapError(err)
		return
	}

	ret = &runtime.Int{Value: int64(dur)}

	return
}

func timesSince(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
	if len(args) != 1 {
		err = runtime.ErrWrongNumArguments
		return
	}

	t1, ok := runtime.ToTime(args[0])
	if !ok {
		err = runtime.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &runtime.Int{Value: int64(time.Since(t1))}

	return
}

func timesUntil(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
	if len(args) != 1 {
		err = runtime.ErrWrongNumArguments
		return
	}

	t1, ok := runtime.ToTime(args[0])
	if !ok {
		err = runtime.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &runtime.Int{Value: int64(time.Until(t1))}

	return
}

func timesDurationHours(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
	if len(args) != 1 {
		err = runtime.ErrWrongNumArguments
		return
	}

	i1, ok := runtime.ToInt64(args[0])
	if !ok {
		err = runtime.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &runtime.Float{Value: time.Duration(i1).Hours()}

	return
}

func timesDurationMinutes(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
	if len(args) != 1 {
		err = runtime.ErrWrongNumArguments
		return
	}

	i1, ok := runtime.ToInt64(args[0])
	if !ok {
		err = runtime.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &runtime.Float{Value: time.Duration(i1).Minutes()}

	return
}

func timesDurationNanoseconds(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
	if len(args) != 1 {
		err = runtime.ErrWrongNumArguments
		return
	}

	i1, ok := runtime.ToInt64(args[0])
	if !ok {
		err = runtime.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &runtime.Int{Value: time.Duration(i1).Nanoseconds()}

	return
}

func timesDurationSeconds(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
	if len(args) != 1 {
		err = runtime.ErrWrongNumArguments
		return
	}

	i1, ok := runtime.ToInt64(args[0])
	if !ok {
		err = runtime.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &runtime.Float{Value: time.Duration(i1).Seconds()}

	return
}

func timesDurationString(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
	if len(args) != 1 {
		err = runtime.ErrWrongNumArguments
		return
	}

	i1, ok := runtime.ToInt64(args[0])
	if !ok {
		err = runtime.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &runtime.String{Value: time.Duration(i1).String()}

	return
}

func timesMonthString(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
	if len(args) != 1 {
		err = runtime.ErrWrongNumArguments
		return
	}

	i1, ok := runtime.ToInt64(args[0])
	if !ok {
		err = runtime.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &runtime.String{Value: time.Month(i1).String()}

	return
}

func timesDate(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
	if len(args) != 7 {
		err = runtime.ErrWrongNumArguments
		return
	}

	i1, ok := runtime.ToInt(args[0])
	if !ok {
		err = runtime.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}
	i2, ok := runtime.ToInt(args[1])
	if !ok {
		err = runtime.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}
	i3, ok := runtime.ToInt(args[2])
	if !ok {
		err = runtime.ErrInvalidArgumentType{
			Name:     "third",
			Expected: "int(compatible)",
			Found:    args[2].TypeName(),
		}
		return
	}
	i4, ok := runtime.ToInt(args[3])
	if !ok {
		err = runtime.ErrInvalidArgumentType{
			Name:     "fourth",
			Expected: "int(compatible)",
			Found:    args[3].TypeName(),
		}
		return
	}
	i5, ok := runtime.ToInt(args[4])
	if !ok {
		err = runtime.ErrInvalidArgumentType{
			Name:     "fifth",
			Expected: "int(compatible)",
			Found:    args[4].TypeName(),
		}
		return
	}
	i6, ok := runtime.ToInt(args[5])
	if !ok {
		err = runtime.ErrInvalidArgumentType{
			Name:     "sixth",
			Expected: "int(compatible)",
			Found:    args[5].TypeName(),
		}
		return
	}
	i7, ok := runtime.ToInt(args[6])
	if !ok {
		err = runtime.ErrInvalidArgumentType{
			Name:     "seventh",
			Expected: "int(compatible)",
			Found:    args[6].TypeName(),
		}
		return
	}

	ret = &runtime.Time{
		Value: time.Date(i1,
			time.Month(i2), i3, i4, i5, i6, i7, time.Now().Location()),
	}

	return
}

func timesNow(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
	if len(args) != 0 {
		err = runtime.ErrWrongNumArguments
		return
	}

	ret = &runtime.Time{Value: time.Now()}

	return
}

func timesParse(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
	if len(args) != 2 {
		err = runtime.ErrWrongNumArguments
		return
	}

	s1, ok := runtime.ToString(args[0])
	if !ok {
		err = runtime.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	s2, ok := runtime.ToString(args[1])
	if !ok {
		err = runtime.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "string(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	parsed, err := time.Parse(s1, s2)
	if err != nil {
		ret = wrapError(err)
		return
	}

	ret = &runtime.Time{Value: parsed}

	return
}

func timesUnix(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
	if len(args) != 2 {
		err = runtime.ErrWrongNumArguments
		return
	}

	i1, ok := runtime.ToInt64(args[0])
	if !ok {
		err = runtime.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	i2, ok := runtime.ToInt64(args[1])
	if !ok {
		err = runtime.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	ret = &runtime.Time{Value: time.Unix(i1, i2)}

	return
}

func timesAdd(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
	if len(args) != 2 {
		err = runtime.ErrWrongNumArguments
		return
	}

	t1, ok := runtime.ToTime(args[0])
	if !ok {
		err = runtime.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	i2, ok := runtime.ToInt64(args[1])
	if !ok {
		err = runtime.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	ret = &runtime.Time{Value: t1.Add(time.Duration(i2))}

	return
}

func timesSub(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
	if len(args) != 2 {
		err = runtime.ErrWrongNumArguments
		return
	}

	t1, ok := runtime.ToTime(args[0])
	if !ok {
		err = runtime.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	t2, ok := runtime.ToTime(args[1])
	if !ok {
		err = runtime.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "time(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	ret = &runtime.Int{Value: int64(t1.Sub(t2))}

	return
}

func timesAddDate(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
	if len(args) != 4 {
		err = runtime.ErrWrongNumArguments
		return
	}

	t1, ok := runtime.ToTime(args[0])
	if !ok {
		err = runtime.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	i2, ok := runtime.ToInt(args[1])
	if !ok {
		err = runtime.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	i3, ok := runtime.ToInt(args[2])
	if !ok {
		err = runtime.ErrInvalidArgumentType{
			Name:     "third",
			Expected: "int(compatible)",
			Found:    args[2].TypeName(),
		}
		return
	}

	i4, ok := runtime.ToInt(args[3])
	if !ok {
		err = runtime.ErrInvalidArgumentType{
			Name:     "fourth",
			Expected: "int(compatible)",
			Found:    args[3].TypeName(),
		}
		return
	}

	ret = &runtime.Time{Value: t1.AddDate(i2, i3, i4)}

	return
}

func timesAfter(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
	if len(args) != 2 {
		err = runtime.ErrWrongNumArguments
		return
	}

	t1, ok := runtime.ToTime(args[0])
	if !ok {
		err = runtime.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	t2, ok := runtime.ToTime(args[1])
	if !ok {
		err = runtime.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "time(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	if t1.After(t2) {
		ret = runtime.TrueValue
	} else {
		ret = runtime.FalseValue
	}

	return
}

func timesBefore(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
	if len(args) != 2 {
		err = runtime.ErrWrongNumArguments
		return
	}

	t1, ok := runtime.ToTime(args[0])
	if !ok {
		err = runtime.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	t2, ok := runtime.ToTime(args[1])
	if !ok {
		err = runtime.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	if t1.Before(t2) {
		ret = runtime.TrueValue
	} else {
		ret = runtime.FalseValue
	}

	return
}

func timesTimeYear(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
	if len(args) != 1 {
		err = runtime.ErrWrongNumArguments
		return
	}

	t1, ok := runtime.ToTime(args[0])
	if !ok {
		err = runtime.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &runtime.Int{Value: int64(t1.Year())}

	return
}

func timesTimeMonth(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
	if len(args) != 1 {
		err = runtime.ErrWrongNumArguments
		return
	}

	t1, ok := runtime.ToTime(args[0])
	if !ok {
		err = runtime.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &runtime.Int{Value: int64(t1.Month())}

	return
}

func timesTimeDay(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
	if len(args) != 1 {
		err = runtime.ErrWrongNumArguments
		return
	}

	t1, ok := runtime.ToTime(args[0])
	if !ok {
		err = runtime.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &runtime.Int{Value: int64(t1.Day())}

	return
}

func timesTimeWeekday(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
	if len(args) != 1 {
		err = runtime.ErrWrongNumArguments
		return
	}

	t1, ok := runtime.ToTime(args[0])
	if !ok {
		err = runtime.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &runtime.Int{Value: int64(t1.Weekday())}

	return
}

func timesTimeHour(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
	if len(args) != 1 {
		err = runtime.ErrWrongNumArguments
		return
	}

	t1, ok := runtime.ToTime(args[0])
	if !ok {
		err = runtime.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &runtime.Int{Value: int64(t1.Hour())}

	return
}

func timesTimeMinute(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
	if len(args) != 1 {
		err = runtime.ErrWrongNumArguments
		return
	}

	t1, ok := runtime.ToTime(args[0])
	if !ok {
		err = runtime.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &runtime.Int{Value: int64(t1.Minute())}

	return
}

func timesTimeSecond(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
	if len(args) != 1 {
		err = runtime.ErrWrongNumArguments
		return
	}

	t1, ok := runtime.ToTime(args[0])
	if !ok {
		err = runtime.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &runtime.Int{Value: int64(t1.Second())}

	return
}

func timesTimeNanosecond(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
	if len(args) != 1 {
		err = runtime.ErrWrongNumArguments
		return
	}

	t1, ok := runtime.ToTime(args[0])
	if !ok {
		err = runtime.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &runtime.Int{Value: int64(t1.Nanosecond())}

	return
}

func timesTimeUnix(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
	if len(args) != 1 {
		err = runtime.ErrWrongNumArguments
		return
	}

	t1, ok := runtime.ToTime(args[0])
	if !ok {
		err = runtime.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &runtime.Int{Value: t1.Unix()}

	return
}

func timesTimeUnixNano(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
	if len(args) != 1 {
		err = runtime.ErrWrongNumArguments
		return
	}

	t1, ok := runtime.ToTime(args[0])
	if !ok {
		err = runtime.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &runtime.Int{Value: t1.UnixNano()}

	return
}

func timesTimeFormat(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
	if len(args) != 2 {
		err = runtime.ErrWrongNumArguments
		return
	}

	t1, ok := runtime.ToTime(args[0])
	if !ok {
		err = runtime.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	s2, ok := runtime.ToString(args[1])
	if !ok {
		err = runtime.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "string(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	s := t1.Format(s2)
	if len(s) > runtime.MaxStringLen {

		return nil, runtime.ErrStringLimit
	}

	ret = &runtime.String{Value: s}

	return
}

func timesIsZero(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
	if len(args) != 1 {
		err = runtime.ErrWrongNumArguments
		return
	}

	t1, ok := runtime.ToTime(args[0])
	if !ok {
		err = runtime.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	if t1.IsZero() {
		ret = runtime.TrueValue
	} else {
		ret = runtime.FalseValue
	}

	return
}

func timesToLocal(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
	if len(args) != 1 {
		err = runtime.ErrWrongNumArguments
		return
	}

	t1, ok := runtime.ToTime(args[0])
	if !ok {
		err = runtime.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &runtime.Time{Value: t1.Local()}

	return
}

func timesToUTC(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
	if len(args) != 1 {
		err = runtime.ErrWrongNumArguments
		return
	}

	t1, ok := runtime.ToTime(args[0])
	if !ok {
		err = runtime.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &runtime.Time{Value: t1.UTC()}

	return
}

func timesTimeLocation(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
	if len(args) != 1 {
		err = runtime.ErrWrongNumArguments
		return
	}

	t1, ok := runtime.ToTime(args[0])
	if !ok {
		err = runtime.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &runtime.String{Value: t1.Location().String()}

	return
}

func timesTimeString(ctx context.Context, args ...runtime.Object) (ret runtime.Object, err error) {
	if len(args) != 1 {
		err = runtime.ErrWrongNumArguments
		return
	}

	t1, ok := runtime.ToTime(args[0])
	if !ok {
		err = runtime.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &runtime.String{Value: t1.String()}

	return
}
