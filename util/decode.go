package util

import (
	rds "github.com/gomodule/redigo/redis"
)

var (
	Bool = rds.Bool

	ByteSlices = rds.ByteSlices

	Bytes = rds.Bytes

	Float64 = rds.Float64

	Float64s = rds.Float64s

	Int = rds.Int

	Int64 = rds.Int64

	Int64s = rds.Int64s

	IntMap = rds.IntMap

	Ints = rds.Ints

	Scan = rds.Scan

	ScanSlice = rds.ScanSlice

	ScanStruct = rds.ScanStruct

	String = rds.String

	StringMap = rds.StringMap

	Strings = rds.Strings

	Uint64 = rds.Uint64

	Values = rds.Values
)
