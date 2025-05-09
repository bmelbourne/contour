// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	cache "sigs.k8s.io/controller-runtime/pkg/cache"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	context "context"

	mock "github.com/stretchr/testify/mock"

	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

// Cache is an autogenerated mock type for the Cache type
type Cache struct {
	mock.Mock
}

// Get provides a mock function with given fields: ctx, key, obj, opts
func (_m *Cache) Get(ctx context.Context, key client.ObjectKey, obj client.Object, opts ...client.GetOption) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, key, obj)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, client.ObjectKey, client.Object, ...client.GetOption) error); ok {
		r0 = rf(ctx, key, obj, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetInformer provides a mock function with given fields: ctx, obj, opts
func (_m *Cache) GetInformer(ctx context.Context, obj client.Object, opts ...cache.InformerGetOption) (cache.Informer, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, obj)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetInformer")
	}

	var r0 cache.Informer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, client.Object, ...cache.InformerGetOption) (cache.Informer, error)); ok {
		return rf(ctx, obj, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, client.Object, ...cache.InformerGetOption) cache.Informer); ok {
		r0 = rf(ctx, obj, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cache.Informer)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, client.Object, ...cache.InformerGetOption) error); ok {
		r1 = rf(ctx, obj, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetInformerForKind provides a mock function with given fields: ctx, gvk, opts
func (_m *Cache) GetInformerForKind(ctx context.Context, gvk schema.GroupVersionKind, opts ...cache.InformerGetOption) (cache.Informer, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, gvk)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetInformerForKind")
	}

	var r0 cache.Informer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, schema.GroupVersionKind, ...cache.InformerGetOption) (cache.Informer, error)); ok {
		return rf(ctx, gvk, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, schema.GroupVersionKind, ...cache.InformerGetOption) cache.Informer); ok {
		r0 = rf(ctx, gvk, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cache.Informer)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, schema.GroupVersionKind, ...cache.InformerGetOption) error); ok {
		r1 = rf(ctx, gvk, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IndexField provides a mock function with given fields: ctx, obj, field, extractValue
func (_m *Cache) IndexField(ctx context.Context, obj client.Object, field string, extractValue client.IndexerFunc) error {
	ret := _m.Called(ctx, obj, field, extractValue)

	if len(ret) == 0 {
		panic("no return value specified for IndexField")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, client.Object, string, client.IndexerFunc) error); ok {
		r0 = rf(ctx, obj, field, extractValue)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// List provides a mock function with given fields: ctx, list, opts
func (_m *Cache) List(ctx context.Context, list client.ObjectList, opts ...client.ListOption) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, list)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, client.ObjectList, ...client.ListOption) error); ok {
		r0 = rf(ctx, list, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveInformer provides a mock function with given fields: ctx, obj
func (_m *Cache) RemoveInformer(ctx context.Context, obj client.Object) error {
	ret := _m.Called(ctx, obj)

	if len(ret) == 0 {
		panic("no return value specified for RemoveInformer")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, client.Object) error); ok {
		r0 = rf(ctx, obj)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Start provides a mock function with given fields: ctx
func (_m *Cache) Start(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Start")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitForCacheSync provides a mock function with given fields: ctx
func (_m *Cache) WaitForCacheSync(ctx context.Context) bool {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for WaitForCacheSync")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context) bool); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// NewCache creates a new instance of Cache. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCache(t interface {
	mock.TestingT
	Cleanup(func())
}) *Cache {
	mock := &Cache{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
