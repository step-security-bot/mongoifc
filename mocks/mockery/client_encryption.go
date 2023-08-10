// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	bson "go.mongodb.org/mongo-driver/bson"

	mock "github.com/stretchr/testify/mock"

	mongo "go.mongodb.org/mongo-driver/mongo"

	mongoifc "github.com/sv-tools/mongoifc"

	options "go.mongodb.org/mongo-driver/mongo/options"

	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// ClientEncryption is an autogenerated mock type for the ClientEncryption type
type ClientEncryption struct {
	mock.Mock
}

// AddKeyAltName provides a mock function with given fields: ctx, id, keyAltName
func (_m *ClientEncryption) AddKeyAltName(ctx context.Context, id primitive.Binary, keyAltName string) mongoifc.SingleResult {
	ret := _m.Called(ctx, id, keyAltName)

	var r0 mongoifc.SingleResult
	if rf, ok := ret.Get(0).(func(context.Context, primitive.Binary, string) mongoifc.SingleResult); ok {
		r0 = rf(ctx, id, keyAltName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mongoifc.SingleResult)
		}
	}

	return r0
}

// Close provides a mock function with given fields: ctx
func (_m *ClientEncryption) Close(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateDataKey provides a mock function with given fields: ctx, kmsProvider, opts
func (_m *ClientEncryption) CreateDataKey(ctx context.Context, kmsProvider string, opts ...*options.DataKeyOptions) (primitive.Binary, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, kmsProvider)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 primitive.Binary
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...*options.DataKeyOptions) (primitive.Binary, error)); ok {
		return rf(ctx, kmsProvider, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...*options.DataKeyOptions) primitive.Binary); ok {
		r0 = rf(ctx, kmsProvider, opts...)
	} else {
		r0 = ret.Get(0).(primitive.Binary)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...*options.DataKeyOptions) error); ok {
		r1 = rf(ctx, kmsProvider, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateEncryptedCollection provides a mock function with given fields: ctx, db, coll, createOpts, kmsProvider, masterKey
func (_m *ClientEncryption) CreateEncryptedCollection(ctx context.Context, db mongoifc.Database, coll string, createOpts *options.CreateCollectionOptions, kmsProvider string, masterKey interface{}) (mongoifc.Collection, primitive.M, error) {
	ret := _m.Called(ctx, db, coll, createOpts, kmsProvider, masterKey)

	var r0 mongoifc.Collection
	var r1 primitive.M
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, mongoifc.Database, string, *options.CreateCollectionOptions, string, interface{}) (mongoifc.Collection, primitive.M, error)); ok {
		return rf(ctx, db, coll, createOpts, kmsProvider, masterKey)
	}
	if rf, ok := ret.Get(0).(func(context.Context, mongoifc.Database, string, *options.CreateCollectionOptions, string, interface{}) mongoifc.Collection); ok {
		r0 = rf(ctx, db, coll, createOpts, kmsProvider, masterKey)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mongoifc.Collection)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, mongoifc.Database, string, *options.CreateCollectionOptions, string, interface{}) primitive.M); ok {
		r1 = rf(ctx, db, coll, createOpts, kmsProvider, masterKey)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(primitive.M)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, mongoifc.Database, string, *options.CreateCollectionOptions, string, interface{}) error); ok {
		r2 = rf(ctx, db, coll, createOpts, kmsProvider, masterKey)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Decrypt provides a mock function with given fields: ctx, val
func (_m *ClientEncryption) Decrypt(ctx context.Context, val primitive.Binary) (bson.RawValue, error) {
	ret := _m.Called(ctx, val)

	var r0 bson.RawValue
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, primitive.Binary) (bson.RawValue, error)); ok {
		return rf(ctx, val)
	}
	if rf, ok := ret.Get(0).(func(context.Context, primitive.Binary) bson.RawValue); ok {
		r0 = rf(ctx, val)
	} else {
		r0 = ret.Get(0).(bson.RawValue)
	}

	if rf, ok := ret.Get(1).(func(context.Context, primitive.Binary) error); ok {
		r1 = rf(ctx, val)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteKey provides a mock function with given fields: ctx, id
func (_m *ClientEncryption) DeleteKey(ctx context.Context, id primitive.Binary) (*mongo.DeleteResult, error) {
	ret := _m.Called(ctx, id)

	var r0 *mongo.DeleteResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, primitive.Binary) (*mongo.DeleteResult, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, primitive.Binary) *mongo.DeleteResult); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mongo.DeleteResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, primitive.Binary) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Encrypt provides a mock function with given fields: ctx, val, opts
func (_m *ClientEncryption) Encrypt(ctx context.Context, val bson.RawValue, opts ...*options.EncryptOptions) (primitive.Binary, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, val)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 primitive.Binary
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, bson.RawValue, ...*options.EncryptOptions) (primitive.Binary, error)); ok {
		return rf(ctx, val, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, bson.RawValue, ...*options.EncryptOptions) primitive.Binary); ok {
		r0 = rf(ctx, val, opts...)
	} else {
		r0 = ret.Get(0).(primitive.Binary)
	}

	if rf, ok := ret.Get(1).(func(context.Context, bson.RawValue, ...*options.EncryptOptions) error); ok {
		r1 = rf(ctx, val, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EncryptExpression provides a mock function with given fields: ctx, expr, result, opts
func (_m *ClientEncryption) EncryptExpression(ctx context.Context, expr interface{}, result interface{}, opts ...*options.EncryptOptions) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, expr, result)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, interface{}, ...*options.EncryptOptions) error); ok {
		r0 = rf(ctx, expr, result, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetKey provides a mock function with given fields: ctx, id
func (_m *ClientEncryption) GetKey(ctx context.Context, id primitive.Binary) mongoifc.SingleResult {
	ret := _m.Called(ctx, id)

	var r0 mongoifc.SingleResult
	if rf, ok := ret.Get(0).(func(context.Context, primitive.Binary) mongoifc.SingleResult); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mongoifc.SingleResult)
		}
	}

	return r0
}

// GetKeyByAltName provides a mock function with given fields: ctx, keyAltName
func (_m *ClientEncryption) GetKeyByAltName(ctx context.Context, keyAltName string) mongoifc.SingleResult {
	ret := _m.Called(ctx, keyAltName)

	var r0 mongoifc.SingleResult
	if rf, ok := ret.Get(0).(func(context.Context, string) mongoifc.SingleResult); ok {
		r0 = rf(ctx, keyAltName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mongoifc.SingleResult)
		}
	}

	return r0
}

// GetKeys provides a mock function with given fields: ctx
func (_m *ClientEncryption) GetKeys(ctx context.Context) (mongoifc.Cursor, error) {
	ret := _m.Called(ctx)

	var r0 mongoifc.Cursor
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (mongoifc.Cursor, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) mongoifc.Cursor); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mongoifc.Cursor)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveKeyAltName provides a mock function with given fields: ctx, id, keyAltName
func (_m *ClientEncryption) RemoveKeyAltName(ctx context.Context, id primitive.Binary, keyAltName string) mongoifc.SingleResult {
	ret := _m.Called(ctx, id, keyAltName)

	var r0 mongoifc.SingleResult
	if rf, ok := ret.Get(0).(func(context.Context, primitive.Binary, string) mongoifc.SingleResult); ok {
		r0 = rf(ctx, id, keyAltName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mongoifc.SingleResult)
		}
	}

	return r0
}

// RewrapManyDataKey provides a mock function with given fields: ctx, filter, opts
func (_m *ClientEncryption) RewrapManyDataKey(ctx context.Context, filter interface{}, opts ...*options.RewrapManyDataKeyOptions) (*mongo.RewrapManyDataKeyResult, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, filter)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *mongo.RewrapManyDataKeyResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...*options.RewrapManyDataKeyOptions) (*mongo.RewrapManyDataKeyResult, error)); ok {
		return rf(ctx, filter, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...*options.RewrapManyDataKeyOptions) *mongo.RewrapManyDataKeyResult); ok {
		r0 = rf(ctx, filter, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mongo.RewrapManyDataKeyResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, interface{}, ...*options.RewrapManyDataKeyOptions) error); ok {
		r1 = rf(ctx, filter, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewClientEncryption creates a new instance of ClientEncryption. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClientEncryption(t interface {
	mock.TestingT
	Cleanup(func())
}) *ClientEncryption {
	mock := &ClientEncryption{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
