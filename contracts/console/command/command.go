package command

import (
	"fmt"
	"strings"

	"github.com/spf13/cast"
)

const (
	ArgTypeBool         = "bool"
	ArgTypeFloat64      = "float64"
	ArgTypeFloat64Slice = "float64_slice"
	ArgTypeInt          = "int"
	ArgTypeIntSlice     = "int_slice"
	ArgTypeInt64        = "int64"
	ArgTypeInt64Slice   = "int64_slice"
	ArgTypeString       = "string"
	ArgTypeStringSlice  = "string_slice"
)

type Extend struct {
	Category string
	Flags    []Flag
	Args     []Arg
}

func (e Extend) ArgsUsage() string {
	args := e.Args
	var usage []string
	for i, v := range args {
		u := ""
		switch v.(type) {
		case *BoolArg:
			t := v.(*BoolArg)
			u = fmt.Sprintf("   Option%d:【%s】%s, 【Required】%s , param will be casted as %s finally.", i+1, t.Name, t.Usage, cast.ToString(t.Required), ArgTypeIntSlice)
		case *Float64Arg:
			t := v.(*Float64Arg)
			u = fmt.Sprintf("   Option%d:【%s】%s, 【Required】%s , param will be casted as %s finally.", i+1, t.Name, t.Usage, cast.ToString(t.Required), ArgTypeFloat64)
		case *Float64SliceArg:
			t := v.(*Float64SliceArg)
			u = fmt.Sprintf("   Option%d:【%s】%s, 【Required】%s , param will be casted as %s finally.", i+1, t.Name, t.Usage, cast.ToString(t.Required), ArgTypeFloat64Slice)
		case *IntArg:
			t := v.(*IntArg)
			u = fmt.Sprintf("   Option%d:【%s】%s, 【Required】%s , param will be casted as %s finally.", i+1, t.Name, t.Usage, cast.ToString(t.Required), ArgTypeInt)
		case *IntSliceArg:
			t := v.(*IntSliceArg)
			u = fmt.Sprintf("   Option%d:【%s】%s, 【Required】%s , param will be casted as %s finally.", i+1, t.Name, t.Usage, cast.ToString(t.Required), ArgTypeIntSlice)
		case *Int64Arg:
			t := v.(*Int64Arg)
			u = fmt.Sprintf("   Option%d:【%s】%s, 【Required】%s , param will be casted as %s finally.", i+1, t.Name, t.Usage, cast.ToString(t.Required), ArgTypeInt64)
		case *Int64SliceArg:
			t := v.(*Int64SliceArg)
			u = fmt.Sprintf("   Option%d:【%s】%s, 【Required】%s , param will be casted as %s finally.", i+1, t.Name, t.Usage, cast.ToString(t.Required), ArgTypeInt64)
		case *StringArg:
			t := v.(*StringArg)
			u = fmt.Sprintf("   Option%d:【%s】%s, 【Required】%s , param will be casted as %s finally.", i+1, t.Name, t.Usage, cast.ToString(t.Required), ArgTypeString)
		case *StringSliceArg:
			t := v.(*StringSliceArg)
			u = fmt.Sprintf("   Option%d:【%s】%s, 【Required】%s , param will be casted as %s finally.", i+1, t.Name, t.Usage, cast.ToString(t.Required), ArgTypeStringSlice)
		}
		usage = append(usage, u)
	}

	return "\n" + strings.Join(usage, "\n")
}

type Flag interface {
	Type() string
}

type BoolFlag struct {
	Name     string
	Aliases  []string
	Usage    string
	Required bool
	Value    bool
}

func (receiver *BoolFlag) Type() string {
	return ArgTypeBool
}

type Float64Flag struct {
	Name     string
	Aliases  []string
	Usage    string
	Required bool
	Value    float64
}

func (receiver *Float64Flag) Type() string {
	return ArgTypeFloat64
}

type Float64SliceFlag struct {
	Name     string
	Aliases  []string
	Usage    string
	Required bool
	Value    []float64
}

func (receiver *Float64SliceFlag) Type() string {
	return ArgTypeFloat64Slice
}

type IntFlag struct {
	Name     string
	Aliases  []string
	Usage    string
	Required bool
	Value    int
}

func (receiver *IntFlag) Type() string {
	return ArgTypeInt
}

type IntSliceFlag struct {
	Name     string
	Aliases  []string
	Usage    string
	Required bool
	Value    []int
}

func (receiver *IntSliceFlag) Type() string {
	return ArgTypeIntSlice
}

type Int64Flag struct {
	Name     string
	Aliases  []string
	Usage    string
	Required bool
	Value    int64
}

func (receiver *Int64Flag) Type() string {
	return ArgTypeInt64
}

type Int64SliceFlag struct {
	Name     string
	Aliases  []string
	Usage    string
	Required bool
	Value    []int64
}

func (receiver *Int64SliceFlag) Type() string {
	return ArgTypeInt64Slice
}

type StringFlag struct {
	Name     string
	Aliases  []string
	Usage    string
	Required bool
	Value    string
}

func (receiver *StringFlag) Type() string {
	return ArgTypeString
}

type StringSliceFlag struct {
	Name     string
	Aliases  []string
	Usage    string
	Required bool
	Value    []string
}

func (receiver *StringSliceFlag) Type() string {
	return ArgTypeStringSlice
}

type Arg interface {
	Type() string
}

type BoolArg struct {
	Name     string
	Aliases  []string
	Usage    string
	Required bool
	Value    bool
}

func (receiver *BoolArg) Type() string {
	return ArgTypeBool
}

type Float64Arg struct {
	Name     string
	Aliases  []string
	Usage    string
	Required bool
	Value    float64
}

func (receiver *Float64Arg) Type() string {
	return ArgTypeFloat64
}

type Float64SliceArg struct {
	Name     string
	Aliases  []string
	Usage    string
	Required bool
	Value    []float64
}

func (receiver *Float64SliceArg) Type() string {
	return ArgTypeFloat64Slice
}

type IntArg struct {
	Name     string
	Aliases  []string
	Usage    string
	Required bool
	Value    int
}

func (receiver *IntArg) Type() string {
	return ArgTypeInt
}

type IntSliceArg struct {
	Name     string
	Aliases  []string
	Usage    string
	Required bool
	Value    []int
}

func (receiver *IntSliceArg) Type() string {
	return ArgTypeIntSlice
}

type Int64Arg struct {
	Name     string
	Aliases  []string
	Usage    string
	Required bool
	Value    int64
}

func (receiver *Int64Arg) Type() string {
	return ArgTypeInt64
}

type Int64SliceArg struct {
	Name     string
	Aliases  []string
	Usage    string
	Required bool
	Value    []int64
}

func (receiver *Int64SliceArg) Type() string {
	return ArgTypeInt64Slice
}

type StringArg struct {
	Name     string
	Aliases  []string
	Usage    string
	Required bool
	Value    string
}

func (receiver *StringArg) Type() string {
	return ArgTypeString
}

type StringSliceArg struct {
	Name     string
	Aliases  []string
	Usage    string
	Required bool
	Value    []string
}

func (receiver *StringSliceArg) Type() string {
	return ArgTypeStringSlice
}
