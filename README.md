# vector

向量集合模型

`import "github.com/minph/model/vector"`

- [Overview](#pkg-overview)
- [Index](#pkg-index)
- [Examples](#pkg-examples)

## <a name="pkg-overview">Overview</a>

## <a name="pkg-index">Index</a>

- [func AsFloat(e Feature) []float64](#AsFloat)
- [func AsInt(e Feature) []int](#AsInt)
- [func AsNum(e Feature) []float64](#AsNum)
- [func AsString(e Feature) []string](#AsString)
- [type AsType](#AsType)
- [type DeriveFunc](#DeriveFunc)
- [type Elem](#Elem)
  - [func New(values ...interface{}) \*Elem](#New)
  - [func (e *Elem) Add(value interface{}) *Elem](#Elem.Add)
  - [func (e \*Elem) AsFloat() []float64](#Elem.AsFloat)
  - [func (e \*Elem) AsInt() []int](#Elem.AsInt)
  - [func (e \*Elem) AsNum() []float64](#Elem.AsNum)
  - [func (e \*Elem) AsString() []string](#Elem.AsString)
  - [func (e \*Elem) Count() int](#Elem.Count)
  - [func (e *Elem) Filter(filterFunc FilterFunc) *Elem](#Elem.Filter)
  - [func (e \*Elem) First() interface{}](#Elem.First)
  - [func (e \*Elem) Foreach(rangeFunc MapFunc)](#Elem.Foreach)
  - [func (e \*Elem) Get() interface{}](#Elem.Get)
  - [func (e \*Elem) Index() int](#Elem.Index)
  - [func (e *Elem) Insert(value interface{}) *Elem](#Elem.Insert)
  - [func (e \*Elem) Last() interface{}](#Elem.Last)
  - [func (e *Elem) Map(deriveFunc DeriveFunc) *Elem](#Elem.Map)
  - [func (e \*Elem) Next() interface{}](#Elem.Next)
  - [func (e \*Elem) Prev() interface{}](#Elem.Prev)
  - [func (e *Elem) Push(value interface{}) *Elem](#Elem.Push)
  - [func (e *Elem) Remove() *Elem](#Elem.Remove)
  - [func (e *Elem) Reverse() *Elem](#Elem.Reverse)
  - [func (e \*Elem) ReverseForeach(rangeFunc MapFunc)](#Elem.ReverseForeach)
  - [func (e \*Elem) Shift() interface{}](#Elem.Shift)
  - [func (e *Elem) Turn(index int) *Elem](#Elem.Turn)
  - [func (e \*Elem) Value() []interface{}](#Elem.Value)
- [type Feature](#Feature)
- [type FilterFunc](#FilterFunc)
- [type MapFunc](#MapFunc)

## <a name="AsFloat">func</a> AsFloat

```go
func AsFloat(e Feature) []float64
```

AsInt 退化为 []float64

## <a name="AsInt">func</a> AsInt

```go
func AsInt(e Feature) []int
```

AsInt 退化为 []int

## <a name="AsNum">func</a> AsNum

```go
func AsNum(e Feature) []float64
```

AsNum 退化为 []float64，且 int 类型也会包含

## <a name="AsString">func</a> AsString

```go
func AsString(e Feature) []string
```

AsString 退化为 []string

## <a name="AsType">type</a> AsType

```go
type AsType interface {
    AsInt() []int
    AsFloat() []float64
    AsNum() []float64
    AsString() []string
}
```

## <a name="DeriveFunc">type</a> DeriveFunc

```go
type DeriveFunc func(index int, value interface{}) interface{}
```

DeriveFunc 派生函数

## <a name="Elem">type</a> Elem

```go
type Elem struct {
    // contains filtered or unexported fields
}

```

Elem 向量集合结构体

### <a name="New">func</a> New

```go
func New(values ...interface{}) *Elem
```

New 创建向量集合

### <a name="Elem.Add">func</a> (\*Elem) Add

```go
func (e *Elem) Add(value interface{}) *Elem
```

Add 在向量集合头部增加元素，当前指针位置元素不变

### <a name="Elem.AsFloat">func</a> (\*Elem) AsFloat

```go
func (e *Elem) AsFloat() []float64
```

### <a name="Elem.AsInt">func</a> (\*Elem) AsInt

```go
func (e *Elem) AsInt() []int
```

### <a name="Elem.AsNum">func</a> (\*Elem) AsNum

```go
func (e *Elem) AsNum() []float64
```

### <a name="Elem.AsString">func</a> (\*Elem) AsString

```go
func (e *Elem) AsString() []string
```

### <a name="Elem.Count">func</a> (\*Elem) Count

```go
func (e *Elem) Count() int
```

Count 获取所有元素个数

### <a name="Elem.Filter">func</a> (\*Elem) Filter

```go
func (e *Elem) Filter(filterFunc FilterFunc) *Elem
```

Filter 从原向量集合过滤产生新向量集合

### <a name="Elem.First">func</a> (\*Elem) First

```go
func (e *Elem) First() interface{}
```

First 获取第一个元素

### <a name="Elem.Foreach">func</a> (\*Elem) Foreach

```go
func (e *Elem) Foreach(rangeFunc MapFunc)
```

Foreach 遍历元素

### <a name="Elem.Get">func</a> (\*Elem) Get

```go
func (e *Elem) Get() interface{}
```

Get 获取当前元素

### <a name="Elem.Index">func</a> (\*Elem) Index

```go
func (e *Elem) Index() int
```

Index 获取元素位置

### <a name="Elem.Insert">func</a> (\*Elem) Insert

```go
func (e *Elem) Insert(value interface{}) *Elem
```

Insert 在指针位置加入元素并指向此元素

### <a name="Elem.Last">func</a> (\*Elem) Last

```go
func (e *Elem) Last() interface{}
```

Last 获取最后一个元素

### <a name="Elem.Map">func</a> (\*Elem) Map

```go
func (e *Elem) Map(deriveFunc DeriveFunc) *Elem
```

Map 从原向量集合派生新向量集合

### <a name="Elem.Next">func</a> (\*Elem) Next

```go
func (e *Elem) Next() interface{}
```

Next 获取下一个元素

### <a name="Elem.Prev">func</a> (\*Elem) Prev

```go
func (e *Elem) Prev() interface{}
```

Prev 获取上一个元素

### <a name="Elem.Push">func</a> (\*Elem) Push

```go
func (e *Elem) Push(value interface{}) *Elem
```

Push 在向量集合尾部增加元素，当前指针位置元素不变

### <a name="Elem.Remove">func</a> (\*Elem) Remove

```go
func (e *Elem) Remove() *Elem
```

Remove 移除当前指针位置元素，然后指针指向上一位

### <a name="Elem.Reverse">func</a> (\*Elem) Reverse

```go
func (e *Elem) Reverse() *Elem
```

Reverse 生成反向向量集合

### <a name="Elem.ReverseForeach">func</a> (\*Elem) ReverseForeach

```go
func (e *Elem) ReverseForeach(rangeFunc MapFunc)
```

ReverseForeach 反向遍历元素

### <a name="Elem.Shift">func</a> (\*Elem) Shift

```go
func (e *Elem) Shift() interface{}
```

Shift 删除第一个元素，并返回第一个元素的值，然后指针指向上一位

### <a name="Elem.Turn">func</a> (\*Elem) Turn

```go
func (e *Elem) Turn(index int) *Elem
```

Turn 将指针转到指定位置，越界则指向第一个元素

### <a name="Elem.Value">func</a> (\*Elem) Value

```go
func (e *Elem) Value() []interface{}
```

Value 获取所有元素

## <a name="Feature">type</a> Feature

```go
type Feature interface {
    Add(value interface{}) *Elem
    Count() int
    Filter(filterFunc FilterFunc) *Elem
    Get() interface{}
    Index() int
    Insert(value interface{}) *Elem
    First() interface{}
    Foreach(rangeFunc MapFunc)
    Last() interface{}
    Map(deriveFunc DeriveFunc) *Elem
    Next() interface{}
    Prev() interface{}
    Reverse() *Elem
    ReverseForeach(rangeFunc MapFunc)
    Push(value interface{}) *Elem
    Remove() *Elem
    Shift() interface{}
    Turn(index int) *Elem
    Value() []interface{}
}
```

Feature 向量集合接口

## <a name="FilterFunc">type</a> FilterFunc

```go
type FilterFunc func(index int, value interface{}) bool
```

FilterFunc 过滤函数

## <a name="MapFunc">type</a> MapFunc

```go
type MapFunc func(index int, value interface{})
```

MapFunc 遍历函数

# ring

数据环模型

`import "github.com/minph/model/ring"`

- [Overview](#pkg-overview)
- [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>

## <a name="pkg-index">Index</a>

- [type Elem](#Elem)
  - [func FromVector(vector *vector.Elem) *Elem](#FromVector)
  - [func New(values ...interface{}) \*Elem](#New)
  - [func (e \*Elem) Count() int](#Elem.Count)
  - [func (e \*Elem) Get() interface{}](#Elem.Get)
  - [func (e *Elem) Insert(value interface{}) *Elem](#Elem.Insert)
  - [func (e \*Elem) Next() interface{}](#Elem.Next)
  - [func (e \*Elem) Prev() interface{}](#Elem.Prev)
  - [func (e *Elem) Remove() *Elem](#Elem.Remove)
  - [func (e *Elem) Turn(index int) *Elem](#Elem.Turn)
  - [func (e *Elem) Vector() *vector.Elem](#Elem.Vector)
- [type Feature](#Feature)

## <a name="Elem">type</a> Elem

```go
type Elem struct {
    // contains filtered or unexported fields
}

```

Elem 数据环结构体

### <a name="FromVector">func</a> FromVector

```go
func FromVector(vector *vector.Elem) *Elem
```

FromVector 从向量集合创建数据环

### <a name="New">func</a> New

```go
func New(values ...interface{}) *Elem
```

New 创建数据环

### <a name="Elem.Count">func</a> (\*Elem) Count

```go
func (e *Elem) Count() int
```

Count 获取所有元素个数

### <a name="Elem.Get">func</a> (\*Elem) Get

```go
func (e *Elem) Get() interface{}
```

Get 获取当前元素

### <a name="Elem.Insert">func</a> (\*Elem) Insert

```go
func (e *Elem) Insert(value interface{}) *Elem
```

Insert 在当前位置增加元素并指向此元素

### <a name="Elem.Next">func</a> (\*Elem) Next

```go
func (e *Elem) Next() interface{}
```

Next 获取下一个元素

### <a name="Elem.Prev">func</a> (\*Elem) Prev

```go
func (e *Elem) Prev() interface{}
```

Prev 获取上一个元素

### <a name="Elem.Remove">func</a> (\*Elem) Remove

```go
func (e *Elem) Remove() *Elem
```

Remove 删去当前元素，然后指针指向上一位

### <a name="Elem.Turn">func</a> (\*Elem) Turn

```go
func (e *Elem) Turn(index int) *Elem
```

Turn 指针转到指定位置

### <a name="Elem.Vector">func</a> (\*Elem) Vector

```go
func (e *Elem) Vector() *vector.Elem
```

Vector 转为向量集合

## <a name="Feature">type</a> Feature

```go
type Feature interface {
    Count() int
    Get() interface{}
    Insert(value interface{}) *Elem
    Next() interface{}
    Prev() interface{}
    Remove() *Elem
    Turn(index int) *Elem
    Vector() *vector.Elem
}
```
