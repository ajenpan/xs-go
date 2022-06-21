package transmit

import (
	"fmt"
	"reflect"
	"strings"
	"sync"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type MethodStyle int

// Todo:
const (
	StyleMicro = iota // func (context.Context, proto.Message, proto.Message) ( error)
	StyleGRpc  = iota // func (context.Context, proto.Message) (proto.Message, error)
)

type Method struct {
	Imp reflect.Method

	H interface{}

	RequestType  reflect.Type
	ResponseType reflect.Type
}

func (m *Method) Call(args ...interface{}) []reflect.Value {
	values := make([]reflect.Value, 0, len(args)+1)
	values = append(values, reflect.ValueOf(m.H))
	for _, v := range args {
		values = append(values, reflect.ValueOf(v))
	}
	return m.Imp.Func.Call(values)
}

func (m *Method) NewRequest() interface{} {
	return reflect.New(m.RequestType).Interface()
}

func (m *Method) NewResponse() interface{} {
	return reflect.New(m.ResponseType).Interface()
}

type CallTable struct {
	sync.RWMutex
	list map[string]*Method
}

func (m *CallTable) Len() int {
	m.RLock()
	defer m.RUnlock()
	return len(m.list)
}

func (m *CallTable) Has(name string) bool {
	m.RLock()
	defer m.RUnlock()
	_, has := m.list[name]
	return has
}

func (m *CallTable) Get(name string) *Method {
	m.RLock()
	defer m.RUnlock()

	ret, has := m.list[name]
	if has {
		return ret
	}
	return nil
}

func (m *CallTable) Range(f func(key string, value *Method) bool) {
	m.Lock()
	defer m.Unlock()
	for k, v := range m.list {
		if !f(k, v) {
			return
		}
	}
}

func (m *CallTable) Merge(other *CallTable, overWrite bool) int {
	ret := 0
	other.RWMutex.RLock()
	defer other.RWMutex.RUnlock()

	m.Lock()
	defer m.Unlock()

	for k, v := range other.list {
		_, has := m.list[k]
		if has && !overWrite {
			continue
		}
		m.list[k] = v
		ret++
	}
	return ret
}

func ParseProtoMessageWithSuffix(suffix string, ms protoreflect.MessageDescriptors, handler interface{}) *CallTable {
	ret := &CallTable{
		list: make(map[string]*Method),
	}

	refHandler := reflect.TypeOf(handler)

	for i := 0; i < ms.Len(); i++ {
		msg := ms.Get(i)
		requestName := string(msg.Name())
		if !strings.HasSuffix(requestName, suffix) {
			continue
		}
		method, has := refHandler.MethodByName(requestName)
		if !has {
			continue
		}
		ret.list[requestName] = &Method{
			Imp:         method,
			RequestType: method.Type.In(2).Elem(),
		}
	}
	return ret
}

//ParseRpcMethod
func ParseRpcMethod(ms protoreflect.ServiceDescriptors, h interface{}) (*CallTable, error) {
	rv := reflect.ValueOf(h)

	if rv.Kind() != reflect.Ptr {
		return nil, fmt.Errorf("%T is not a pointer", h)
	}

	ret := &CallTable{
		list: make(map[string]*Method),
	}

	refh := reflect.TypeOf(h)
	for i := 0; i < ms.Len(); i++ {
		rpcName := string(ms.Get(i).Name())
		rpcMethods := ms.Get(i).Methods()
		for j := 0; j < rpcMethods.Len(); j++ {
			rpcMethod := rpcMethods.Get(j)
			rpcMethodName := string(rpcMethod.Name())

			method, has := refh.MethodByName(rpcMethodName)
			if !has {
				continue
			}
			epn := strings.Join([]string{rpcName, rpcMethodName}, "/")

			ret.list[epn] = &Method{
				Imp:          method,
				H:            h,
				RequestType:  method.Type.In(2).Elem(),
				ResponseType: method.Type.In(3).Elem(),
			}
		}
	}
	return ret, nil
}
