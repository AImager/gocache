package cache

import (
	"context"
	"reflect"
	"strconv"

	"github.com/AImager/gocache/config"
	"github.com/AImager/gocache/util"
)

type Cache struct {
	Client config.ClientI
}

func (c *Cache) CacheWithFunc(ctx context.Context, cc config.CacheConfig, decoPtr, fn interface{}) error {
	decoratedFunc := reflect.ValueOf(decoPtr).Elem()
	targetFunc := reflect.ValueOf(fn)
	targetFuncType := reflect.TypeOf(fn)

	v := reflect.MakeFunc(targetFunc.Type(),
		func(in []reflect.Value) (out []reflect.Value) {
			out = []reflect.Value{}
			reply, err := c.Client.Do(ctx, util.Get, cc.Key)
			if err != nil {
				out = targetFunc.Call(in)
				var err1 error
				if cc.Expire < 1 {
					_, err1 = c.Client.Do(ctx, util.Set, cc.Key, build(out))
				} else {
					_, err1 = c.Client.Do(ctx, util.Set, cc.Key, build(out), util.EX, cc.Expire)
				}
				if err1 != nil {
					out = []reflect.Value{}
					out = append(out, reflect.Zero(targetFuncType.Out(0)))
					out = append(out, reflect.ValueOf(err1))
				}
				return
			} else {
				ret, err1 := format(targetFuncType.Out(0), reply)
				out = append(out, reflect.ValueOf(ret))
				if err1 != nil {
					out = append(out, reflect.ValueOf(err1))
				} else {
					out = append(out, reflect.ValueOf(&err1).Elem())
				}
				return
			}
		})
	decoratedFunc.Set(v)
	return nil
}

func (c Cache) CacheDelWithFunc(ctx context.Context, cc config.CacheDelConfig, decoPtr, fn interface{}) (err error) {
	var decoratedFunc, targetFunc reflect.Value
	decoratedFunc = reflect.ValueOf(decoPtr).Elem()
	targetFunc = reflect.ValueOf(fn)
	v := reflect.MakeFunc(targetFunc.Type(),
		func(in []reflect.Value) (out []reflect.Value) {
			out = targetFunc.Call(in)
			c.Client.Do(ctx, util.Del, cc.Key)
			return
		})
	decoratedFunc.Set(v)
	return
}

func build(out []reflect.Value) string {
	switch out[0].Kind() {
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(out[0].Int(), 10)
	case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(out[0].Uint(), 10)
	case reflect.Bool:
		return strconv.FormatBool(out[0].Bool())
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(out[0].Float(), 'b', 5, 64)
	case reflect.String:
		return out[0].String()
	}
	return ""
}

func format(firstField reflect.Type, reply interface{}) (ret interface{}, err error) {
	temp, err := util.String(reply, err)
	if err != nil {
		return nil, err
	}
	switch firstField.Kind() {
	case reflect.Int16:
		int16ret, err := strconv.ParseInt(temp, 10, 16)
		return int16(int16ret), err
	case reflect.Int:
		intret, err := strconv.ParseInt(temp, 10, 32)
		return int(intret), err
	case reflect.Int32:
		int32ret, err := strconv.ParseInt(temp, 10, 32)
		return int32(int32ret), err
	case reflect.Int64:
		return strconv.ParseInt(temp, 10, 64)
	case reflect.Uint16:
		uint16ret, err := strconv.ParseUint(temp, 10, 32)
		return uint16(uint16ret), err
	case reflect.Uint:
		uintret, err := strconv.ParseUint(temp, 10, 32)
		return uint(uintret), err
	case reflect.Uint32:
		uint32ret, err := strconv.ParseUint(temp, 10, 32)
		return uint32(uint32ret), err
	case reflect.Uint64:
		return strconv.ParseUint(temp, 10, 32)
	case reflect.Bool:
		return strconv.ParseBool(temp)
	case reflect.Float32:
		float32ret, err := strconv.ParseFloat(temp, 32)
		return float32(float32ret), err
	case reflect.Float64:
		return strconv.ParseFloat(temp, 64)
	case reflect.String:
		return temp, nil
	}
	return ret, err
}
