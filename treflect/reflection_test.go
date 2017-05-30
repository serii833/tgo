// Copyright 2015-2016 trivago GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package treflect

import (
	"github.com/trivago/tgo/ttesting"
	"testing"
)

type renamedUint32 uint32

type reflectTestData struct {
	b    bool
	ui   uint
	ui8  uint8
	ui16 uint16
	ui32 uint32
	ui64 uint64
	i    int
	i8   int8
	i16  int16
	i32  int32
	i64  int64
	f32  float32
	f64  float64
	s    string
	by   byte
	r    renamedUint32
	sa   []string
	m    map[string]bool
}

func TestSetValue(t *testing.T) {
	expect := ttesting.NewExpect(t)
	data := &reflectTestData{}
	arrayData := []string{"foo", "bar"}
	mapData := map[string]bool{"foo": true, "bar": true}

	SetMemberByName(data, "b", true)
	SetMemberByName(data, "by", byte(10))
	SetMemberByName(data, "ui", uint(1))
	SetMemberByName(data, "ui8", uint8(2))
	SetMemberByName(data, "ui16", uint16(3))
	SetMemberByName(data, "ui32", uint32(4))
	SetMemberByName(data, "ui64", uint64(5))
	SetMemberByName(data, "i", int(1))
	SetMemberByName(data, "i8", int8(2))
	SetMemberByName(data, "i16", int16(3))
	SetMemberByName(data, "i32", int32(4))
	SetMemberByName(data, "i64", int64(5))
	SetMemberByName(data, "f32", float32(1.0))
	SetMemberByName(data, "f64", float64(2.0))
	SetMemberByName(data, "s", string("test"))
	SetMemberByName(data, "sa", arrayData)
	SetMemberByName(data, "r", renamedUint32(11))
	SetMemberByName(data, "m", mapData)

	expect.Equal(true, data.b)
	expect.Equal(byte(10), data.by)
	expect.Equal(uint(1), data.ui)
	expect.Equal(uint8(2), data.ui8)
	expect.Equal(uint16(3), data.ui16)
	expect.Equal(uint32(4), data.ui32)
	expect.Equal(uint64(5), data.ui64)
	expect.Equal(int(1), data.i)
	expect.Equal(int8(2), data.i8)
	expect.Equal(int16(3), data.i16)
	expect.Equal(int32(4), data.i32)
	expect.Equal(int64(5), data.i64)
	expect.Equal(float32(1.0), data.f32)
	expect.Equal(float64(2.0), data.f64)
	expect.Equal("test", data.s)
	expect.Equal(arrayData, data.sa)
	expect.Equal(renamedUint32(11), data.r)
	expect.Equal(mapData, data.m)
}

func TestSetArray(t *testing.T) {
	expect := ttesting.NewExpect(t)
	data := &reflectTestData{}
	arrayData := []string{"foo", "bar"}

	SetMemberByName(data, "sa", arrayData)

	arrayData = append([]string{"some", "more", "data"}, arrayData...)

	expect.Equal([]string{"foo", "bar"}, data.sa)
}

func TestSetMap(t *testing.T) {
	expect := ttesting.NewExpect(t)
	data := &reflectTestData{}
	mapData := map[string]bool{"foo": true, "bar": true}

	SetMemberByName(data, "m", mapData)
	mapData["a"] = true

	expect.Equal(map[string]bool{"foo": true, "bar": true, "a": true}, data.m)
}