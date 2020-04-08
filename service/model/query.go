package model

import (
	"reflect"
	"strings"
	"sync"
)

var (
	valPaging        = reflect.ValueOf(Paging{})
	pagingTypeString = valPaging.Type().String()
)

//NewQuery generate condition for query
func NewQueryCond(cond interface{}) *QueryCond {
	where := newCond(structToMap(cond))

	var paging Paging
	if val, ok := where.Result()[pagingTypeString]; ok {
		where.Delete(pagingTypeString)
		if tmp, ok := val.(Paging); ok {
			paging = tmp
		}
	}

	return &QueryCond{
		where:  where,
		paging: &paging,
	}
}

type IQueryCond interface {
	Where() map[string]interface{}
	Paging() *Paging
}

type QueryCond struct {
	where  *cond
	paging *Paging
}

func (q *QueryCond) Where() map[string]interface{} {
	return q.where.Result()
}

func (q *QueryCond) Paging() *Paging {
	return q.paging
}

//NewUpdate generate condition for update
func NewUpdateCond(update interface{}, where interface{}) *UpdateCond {
	return &UpdateCond{
		update: newCond(structToMap(update)),
		where:  newCond(structToMap(where)),
	}
}

type IUpdateCond interface {
	Update() map[string]interface{}
	Where() map[string]interface{}
}

type UpdateCond struct {
	update *cond
	where  *cond
}

func (q *UpdateCond) Update() map[string]interface{} {
	return q.update.Result()
}

func (q *UpdateCond) Where() map[string]interface{} {
	return q.where.Result()
}

func newCond(m _map) *cond {
	if m == nil {
		m = make(map[string]interface{})
	}

	return &cond{
		_map: m,
	}
}

type _map map[string]interface{}

type cond struct {
	mx sync.RWMutex
	_map
}

func (c *cond) Set(key string, val interface{}) *cond {
	if !c.isNil(val) {
		c.mx.Lock()
		c._map[key] = val
		c.mx.Unlock()
	}
	return c
}

func (c *cond) NullSet(key string, val interface{}) *cond {
	c.mx.Lock()
	c._map[key] = val
	c.mx.Unlock()
	return c
}

func (c *cond) Delete(key string) *cond {
	c.mx.Lock()
	delete(c._map, key)
	c.mx.Unlock()
	return c
}

func (c *cond) Result() map[string]interface{} {
	return c._map
}

func (c *cond) isNil(v interface{}) bool {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		return rv.IsNil()
	}
	return false
}

func structToMap(item interface{}) map[string]interface{} {
	res := map[string]interface{}{}
	if item == nil {
		return res
	}

	v := reflect.TypeOf(item)
	reflectValue := reflect.ValueOf(item)
	reflectValue = reflect.Indirect(reflectValue)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		tag := v.Field(i).Tag.Get("json")
		fieldValue := reflectValue.Field(i)

		if fieldValue.Kind() == reflect.Ptr && fieldValue.IsNil() {
			continue
		}

		field := fieldValue.Interface()
		if tag != "" && tag != "-" {
			tags := strings.Split(tag, ",")
			tag = strings.TrimSpace(tags[0])
			if v.Field(i).Type.Kind() == reflect.Struct {
				res[tag] = structToMap(field)
			} else {
				res[tag] = field
			}
		} else if fieldValue.Type() == valPaging.Type() {
			res[pagingTypeString] = field
		}
	}

	return res
}
