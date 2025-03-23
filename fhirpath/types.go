package fhirpath

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/DAMEDIC/fhir-toolbox-go/fhirpath/overflow"
	"github.com/DAMEDIC/fhir-toolbox-go/fhirpath/parser/gen"
	"github.com/cockroachdb/apd/v3"
	"maps"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Element interface {
	// Children returns all child nodes with given names.
	//
	// If no name is passed, all children are returned.
	Children(name ...string) Collection
	ToBoolean(explicit bool) (v Boolean, ok bool, err error)
	ToString(explicit bool) (v String, ok bool, err error)
	ToInteger(explicit bool) (v Integer, ok bool, err error)
	ToDecimal(explicit bool) (v Decimal, ok bool, err error)
	ToDate(explicit bool) (v Date, ok bool, err error)
	ToTime(explicit bool) (v Time, ok bool, err error)
	ToDateTime(explicit bool) (v DateTime, ok bool, err error)
	ToQuantity(explicit bool) (v Quantity, ok bool, err error)
	Equal(other Element, _noReverseTypeConversion ...bool) (eq bool, ok bool)
	Equivalent(other Element, _noReverseTypeConversion ...bool) bool
	TypeInfo() TypeInfo
	fmt.Stringer
}

type cmpElement interface {
	Element
	// Cmp may return nil, because attempting to operate on quantities
	// with invalid units will result in empty ({ }).
	Cmp(other Element) (cmp int, ok bool, err error)
}

type multiplyElement interface {
	Element
	Multiply(ctx context.Context, other Element) (Element, error)
}

type divideElement interface {
	Element
	Divide(ctx context.Context, other Element) (Element, error)
}

type divElement interface {
	Element
	Div(ctx context.Context, other Element) (Element, error)
}

type modElement interface {
	Element
	Mod(ctx context.Context, other Element) (Element, error)
}

type addElement interface {
	Element
	Add(ctx context.Context, other Element) (Element, error)
}

type subtractElement interface {
	Element
	Subtract(ctx context.Context, other Element) (Element, error)
}

type apdContextKey struct{}

// WithAPDContext sets the apd.Context for Decimal operations in FHIRPath evaluations.
//
// The apd.Context controls the precision and rounding behavior of decimal operations.
// By default, the precision is set to 0, which means that the precision is determined
// by the operands. To set a specific precision, use apd.BaseContext.WithPrecision(n).
//
// Example:
//
//	// Set precision to 10 digits
//	ctx := r4.Context()
//	ctx = fhirpath.WithAPDContext(ctx, apd.BaseContext.WithPrecision(10))
//
//	// Evaluate an expression with the specified precision
//	result, err := fhirpath.Evaluate(ctx, resource, expr)
func WithAPDContext(
	ctx context.Context,
	apdContext *apd.Context,
) context.Context {
	return context.WithValue(ctx, apdContextKey{}, apdContext)
}

func apdContext(ctx context.Context) *apd.Context {
	apdContext, ok := ctx.Value(apdContextKey{}).(*apd.Context)
	if !ok {
		return &apd.BaseContext
	}
	return apdContext
}

type TypeInfo interface {
	Element
	QualifiedName() (TypeSpecifier, bool)
	BaseTypeName() (TypeSpecifier, bool)
}

type SimpleTypeInfo struct {
	defaultConversionError[SimpleTypeInfo]
	Namespace string
	Name      string
	BaseType  TypeSpecifier
}

func (i SimpleTypeInfo) QualifiedName() (TypeSpecifier, bool) {
	return TypeSpecifier{Namespace: i.Namespace, Name: i.Name}, true
}
func (i SimpleTypeInfo) BaseTypeName() (TypeSpecifier, bool) {
	return i.BaseType, true
}
func (i SimpleTypeInfo) Children(name ...string) Collection {
	var children Collection
	if len(name) == 0 || slices.Contains(name, "namespace") {
		children = append(children, String(i.Namespace))
	}
	if len(name) == 0 || slices.Contains(name, "name") {
		children = append(children, String(i.Name))
	}
	if len(name) == 0 || slices.Contains(name, "baseType") {
		children = append(children, i.BaseType)
	}
	return children
}
func (i SimpleTypeInfo) Equal(other Element, _noReverseTypeConversion ...bool) (eq bool, ok bool) {
	return i == other, true
}
func (i SimpleTypeInfo) Equivalent(other Element, _noReverseTypeConversion ...bool) bool {
	eq, _ := i.Equal(other)
	return eq
}
func (i SimpleTypeInfo) TypeInfo() TypeInfo {
	return ClassInfo{
		SimpleTypeInfo: SimpleTypeInfo{
			Namespace: "System",
			Name:      "SimpleTypeInfo",
			BaseType:  TypeSpecifier{Namespace: "System", Name: "Any"},
		},
		Element: []ClassInfoElement{
			{Name: "namespace", Type: TypeSpecifier{Namespace: "System", Name: "String"}},
			{Name: "name", Type: TypeSpecifier{Namespace: "System", Name: "String"}},
			{Name: "baseType", Type: TypeSpecifier{Namespace: "System", Name: "TypeSpecifier"}},
		},
	}
}
func (i SimpleTypeInfo) String() string {
	buf, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

type ClassInfo struct {
	defaultConversionError[ClassInfo]
	SimpleTypeInfo
	Element []ClassInfoElement
}

func (i ClassInfo) QualifiedName() (TypeSpecifier, bool) {
	return TypeSpecifier{Namespace: i.Namespace, Name: i.Name}, true
}
func (i ClassInfo) BaseTypeName() (TypeSpecifier, bool) {
	return i.BaseType, true
}
func (i ClassInfo) Children(name ...string) Collection {
	var children Collection
	if len(name) == 0 || slices.Contains(name, "namespace") {
		children = append(children, String(i.Namespace))
	}
	if len(name) == 0 || slices.Contains(name, "name") {
		children = append(children, String(i.Name))
	}
	if len(name) == 0 || slices.Contains(name, "baseType") {
		children = append(children, i.BaseType)
	}
	if len(name) == 0 || slices.Contains(name, "element") {
		for _, e := range i.Element {
			children = append(children, e)
		}
	}
	return children
}
func (i ClassInfo) Equal(other Element, _noReverseTypeConversion ...bool) (eq bool, ok bool) {
	o, ok := other.(ClassInfo)
	if !ok {
		return false, true
	}
	if i.SimpleTypeInfo != o.SimpleTypeInfo {
		return false, true
	}
	if len(i.Element) != len(o.Element) {
		return false, true
	}
	for i, e := range i.Element {
		if e != o.Element[i] {
			return false, true
		}
	}
	return true, true
}
func (i ClassInfo) Equivalent(other Element, _noReverseTypeConversion ...bool) bool {
	eq, _ := i.Equal(other)
	return eq
}
func (i ClassInfo) TypeInfo() TypeInfo {
	return ClassInfo{
		SimpleTypeInfo: SimpleTypeInfo{
			Namespace: "System",
			Name:      "ClassInfo",
			BaseType:  TypeSpecifier{Namespace: "System", Name: "Any"},
		},
		Element: []ClassInfoElement{
			{Name: "namespace", Type: TypeSpecifier{Namespace: "System", Name: "String"}},
			{Name: "name", Type: TypeSpecifier{Namespace: "System", Name: "String"}},
			{Name: "baseType", Type: TypeSpecifier{Namespace: "System", Name: "TypeSpecifier"}},
			{Name: "element", Type: TypeSpecifier{Namespace: "System", Name: "ClassInfoElement"}},
		},
	}
}
func (i ClassInfo) String() string {
	buf, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

type ClassInfoElement struct {
	defaultConversionError[ClassInfoElement]
	Name       string
	Type       TypeSpecifier
	IsOneBased bool
}

func (i ClassInfoElement) Children(name ...string) Collection {
	var children Collection
	if len(name) == 0 || slices.Contains(name, "name") {
		children = append(children, String(i.Name))
	}
	if len(name) == 0 || slices.Contains(name, "type") {
		children = append(children, i.Type)
	}
	if len(name) == 0 || slices.Contains(name, "isOneBased") {
		children = append(children, Boolean(i.IsOneBased))
	}
	return children
}
func (i ClassInfoElement) Equal(other Element, _noReverseTypeConversion ...bool) (eq bool, ok bool) {
	return i == other, true
}
func (i ClassInfoElement) Equivalent(other Element, _noReverseTypeConversion ...bool) bool {
	eq, _ := i.Equal(other)
	return eq
}
func (i ClassInfoElement) TypeInfo() TypeInfo {
	return ClassInfo{
		SimpleTypeInfo: SimpleTypeInfo{
			Namespace: "System",
			Name:      "ClassInfoElement",
			BaseType:  TypeSpecifier{Namespace: "System", Name: "Any"},
		},
		Element: []ClassInfoElement{
			{Name: "name", Type: TypeSpecifier{Namespace: "System", Name: "String"}},
			{Name: "type", Type: TypeSpecifier{Namespace: "System", Name: "TypeSpecifier"}},
			{Name: "isOneBased", Type: TypeSpecifier{Namespace: "System", Name: "Boolean"}},
		},
	}
}
func (i ClassInfoElement) String() string {
	buf, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

type ListTypeInfo struct {
	defaultConversionError[ListTypeInfo]
	ElementType TypeSpecifier
}

func (i ListTypeInfo) QualifiedName() (TypeSpecifier, bool) {
	return TypeSpecifier{}, false
}
func (i ListTypeInfo) BaseTypeName() (TypeSpecifier, bool) {
	return TypeSpecifier{}, false
}
func (i ListTypeInfo) Children(name ...string) Collection {
	var children Collection
	if len(name) == 0 || slices.Contains(name, "elementType") {
		children = append(children, i.ElementType)
	}
	return children
}
func (i ListTypeInfo) Equal(other Element, _noReverseTypeConversion ...bool) (eq bool, ok bool) {
	return i == other, true
}
func (i ListTypeInfo) Equivalent(other Element, _noReverseTypeConversion ...bool) bool {
	eq, _ := i.Equal(other)
	return eq
}
func (i ListTypeInfo) TypeInfo() TypeInfo {
	return ClassInfo{
		SimpleTypeInfo: SimpleTypeInfo{
			Namespace: "System",
			Name:      "ListTypeInfo",
			BaseType:  TypeSpecifier{Namespace: "System", Name: "Any"},
		},
		Element: []ClassInfoElement{
			{Name: "elementType", Type: TypeSpecifier{Namespace: "System", Name: "TypeSpecifier"}},
		},
	}
}
func (i ListTypeInfo) String() string {
	buf, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

type TupleTypeInfo struct {
	defaultConversionError[TupleTypeInfo]
	Element []TupleTypeInfoElement
}

func (i TupleTypeInfo) QualifiedName() (TypeSpecifier, bool) {
	return TypeSpecifier{}, false
}
func (i TupleTypeInfo) BaseTypeName() (TypeSpecifier, bool) {
	return TypeSpecifier{}, false
}
func (i TupleTypeInfo) Children(name ...string) Collection {
	var children Collection
	if len(name) == 0 || slices.Contains(name, "element") {
		for _, e := range i.Element {
			children = append(children, e)
		}
	}
	return children
}
func (i TupleTypeInfo) Equal(other Element, _noReverseTypeConversion ...bool) (eq bool, ok bool) {
	o, ok := other.(TupleTypeInfo)
	if !ok {
		return false, true
	}
	if len(i.Element) != len(o.Element) {
		return false, true
	}
	for i, e := range i.Element {
		if e != o.Element[i] {
			return false, true
		}
	}
	return true, true
}
func (i TupleTypeInfo) Equivalent(other Element, _noReverseTypeConversion ...bool) bool {
	eq, _ := i.Equal(other)
	return eq
}
func (i TupleTypeInfo) TypeInfo() TypeInfo {
	return ClassInfo{
		SimpleTypeInfo: SimpleTypeInfo{
			Namespace: "System",
			Name:      "TupleTypeInfo",
			BaseType:  TypeSpecifier{Namespace: "System", Name: "Any"},
		},
		Element: []ClassInfoElement{
			{Name: "element", Type: TypeSpecifier{Namespace: "System", Name: "TupleTypeInfoElement"}},
		},
	}
}
func (i TupleTypeInfo) String() string {
	buf, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

type TupleTypeInfoElement struct {
	defaultConversionError[TupleTypeInfoElement]
	Name       string
	Type       TypeSpecifier
	IsOneBased bool
}

func (i TupleTypeInfoElement) Children(name ...string) Collection {
	var children Collection
	if len(name) == 0 || slices.Contains(name, "name") {
		children = append(children, String(i.Name))
	}
	if len(name) == 0 || slices.Contains(name, "type") {
		children = append(children, i.Type)
	}
	if len(name) == 0 || slices.Contains(name, "isOneBased") {
		children = append(children, Boolean(i.IsOneBased))
	}
	return children
}
func (i TupleTypeInfoElement) Equal(other Element, _noReverseTypeConversion ...bool) (eq bool, ok bool) {
	return i == other, true
}
func (i TupleTypeInfoElement) Equivalent(other Element, _noReverseTypeConversion ...bool) bool {
	eq, _ := i.Equal(other)
	return eq
}
func (i TupleTypeInfoElement) TypeInfo() TypeInfo {
	return ClassInfo{
		SimpleTypeInfo: SimpleTypeInfo{
			Namespace: "System",
			Name:      "TupleTypeInfoElement",
			BaseType:  TypeSpecifier{Namespace: "System", Name: "Any"},
		},
		Element: []ClassInfoElement{
			{Name: "name", Type: TypeSpecifier{Namespace: "System", Name: "String"}},
			{Name: "type", Type: TypeSpecifier{Namespace: "System", Name: "TypeSpecifier"}},
			{Name: "isOneBased", Type: TypeSpecifier{Namespace: "System", Name: "Boolean"}},
		},
	}
}
func (i TupleTypeInfoElement) String() string {
	buf, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

type TypeSpecifier struct {
	defaultConversionError[TypeSpecifier]
	Namespace string
	Name      string
	List      bool
}

func ParseTypeSpecifier(s string) TypeSpecifier {
	if strings.HasPrefix(s, "List<") {
		s = strings.TrimPrefix(s, "List<")
		s = strings.TrimSuffix(s, ">")
	}
	split := strings.SplitN(s, ".", 2)
	if len(split) == 1 {
		return TypeSpecifier{
			Name: split[0],
		}
	} else {
		return TypeSpecifier{
			Namespace: split[0],
			Name:      split[1],
		}
	}
}

func (t TypeSpecifier) Children(name ...string) Collection {
	return nil
}
func (t TypeSpecifier) Equal(other Element, _noReverseTypeConversion ...bool) (eq bool, ok bool) {
	return t == other, true
}
func (t TypeSpecifier) Equivalent(other Element, _noReverseTypeConversion ...bool) bool {
	eq, _ := t.Equal(other)
	return eq
}
func (t TypeSpecifier) TypeInfo() TypeInfo {
	return SimpleTypeInfo{
		Namespace: "System",
		Name:      "TypeSpecifier",
		BaseType:  TypeSpecifier{Namespace: "System", Name: "Any"},
	}
}

func (t TypeSpecifier) String() string {
	var s string
	if t.Namespace != "" {
		s = fmt.Sprintf("%s.%s", t.Namespace, t.Name)
	} else {
		s = t.Name
	}
	if t.List {
		return fmt.Sprintf("List<%s>", s)
	}
	return s
}

type namespaceKey struct{}

// WithNamespace installs the default namespace into the context.
func WithNamespace(
	ctx context.Context,
	namespace string,
) context.Context {
	return context.WithValue(ctx, namespaceKey{}, namespace)
}

func contextNamespace(ctx context.Context) string {
	ns, ok := ctx.Value(namespaceKey{}).(string)
	if !ok {
		return "System"
	}
	return ns
}

type knownTypesKey struct{}

// WithTypes installs the known types into the context.
func WithTypes(
	ctx context.Context,
	types []TypeInfo,
) context.Context {
	typeMap := knownTypes(ctx)
	for _, t := range types {
		qual, ok := t.QualifiedName()
		if !ok {
			continue
		}
		typeMap[qual] = t
	}
	return context.WithValue(ctx, knownTypesKey{}, typeMap)
}

func knownTypes(ctx context.Context) map[TypeSpecifier]TypeInfo {
	types, ok := ctx.Value(knownTypesKey{}).(map[TypeSpecifier]TypeInfo)
	if !ok {
		types = maps.Clone(systemTypesMap())
	}
	return types
}

var (
	systemTypes = []TypeInfo{
		Boolean(false).TypeInfo(),
		String("").TypeInfo(),
		Integer(0).TypeInfo(),
		Decimal{}.TypeInfo(),
		Date{}.TypeInfo(),
		Time{}.TypeInfo(),
		DateTime{}.TypeInfo(),
		Quantity{}.TypeInfo(),
	}
	systemTypesMap = sync.OnceValue(func() map[TypeSpecifier]TypeInfo {
		m := map[TypeSpecifier]TypeInfo{}
		for _, t := range systemTypes {
			q, ok := t.QualifiedName()
			if !ok {
				continue
			}
			m[q] = t
		}
		return m
	})
)

func resolveType(ctx context.Context, spec TypeSpecifier) (TypeInfo, bool) {
	if spec.Namespace == "" {
		// search context-specific namespace first
		info, ok := resolveType(ctx, TypeSpecifier{
			Namespace: contextNamespace(ctx),
			Name:      spec.Name,
		})
		if !ok {
			info, ok = resolveType(ctx, TypeSpecifier{
				Namespace: "System",
				Name:      spec.Name,
			})
		}
		return info, ok
	}

	typeMap := knownTypes(ctx)

	t, ok := typeMap[spec]
	return t, ok
}

func subTypeOf(ctx context.Context, target, isOf TypeInfo) bool {
	isOfQual, ok := isOf.QualifiedName()
	if !ok {
		// has no type
		return false
	}

	typQual, ok := target.QualifiedName()
	if ok && typQual == isOfQual {
		return true
	}
	baseQual, ok := target.BaseTypeName()
	if ok && baseQual == isOfQual {
		return true
	}

	baseType, ok := resolveType(ctx, baseQual)
	if !ok {
		return false
	}

	return subTypeOf(ctx, baseType, isOf)
}

func isType(ctx context.Context, target Element, isOf TypeSpecifier) (Element, error) {
	typ, ok := resolveType(ctx, isOf)
	if !ok {
		return nil, fmt.Errorf("can not resolve type `%s`", isOf)
	}
	return Boolean(subTypeOf(ctx, target.TypeInfo(), typ)), nil
}

func asType(ctx context.Context, target Element, asOf TypeSpecifier) (Collection, error) {
	typ, ok := resolveType(ctx, asOf)
	if !ok {
		return nil, fmt.Errorf("can not resolve type `%s`", asOf)
	}
	if subTypeOf(ctx, target.TypeInfo(), typ) {
		return Collection{target}, nil
	} else {
		return nil, nil
	}
}

func elementTo[T Element](e Element, explicit bool) (v T, ok bool, err error) {
	switch any(v).(type) {
	case Boolean:
		v, ok, err := e.ToBoolean(explicit)
		return any(v).(T), ok, err
	case String:
		v, ok, err := e.ToString(explicit)
		return any(v).(T), ok, err
	case Integer:
		v, ok, err := e.ToInteger(explicit)
		return any(v).(T), ok, err
	case Decimal:
		v, ok, err := e.ToDecimal(explicit)
		return any(v).(T), ok, err
	case Date:
		v, ok, err := e.ToDate(explicit)
		return any(v).(T), ok, err
	case Time:
		v, ok, err := e.ToTime(explicit)
		return any(v).(T), ok, err
	case DateTime:
		v, ok, err := e.ToDateTime(explicit)
		return any(v).(T), ok, err
	case Quantity:
		v, ok, err := e.ToQuantity(explicit)
		return any(v).(T), ok, err
	default:
		return v, false, fmt.Errorf("can not convert to type %T", v)
	}
}

func toPrimitive(e Element) (Element, bool) {
	if p, ok, err := e.ToBoolean(false); err == nil && ok {
		return p, true
	}
	if p, ok, err := e.ToString(false); err == nil && ok {
		return p, true
	}
	if p, ok, err := e.ToInteger(false); err == nil && ok {
		return p, true
	}
	if p, ok, err := e.ToDecimal(false); err == nil && ok {
		return p, true
	}
	if p, ok, err := e.ToDateTime(false); err == nil && ok {
		return p, true
	}
	if p, ok, err := e.ToDate(false); err == nil && ok {
		return p, true
	}
	if p, ok, err := e.ToTime(false); err == nil && ok {
		return p, true
	}
	if p, ok, err := e.ToQuantity(false); err == nil && ok {
		return p, true
	}
	return nil, false
}

type Collection []Element

func (c Collection) Equal(other Collection) (eq bool, ok bool) {
	if len(c) == 0 || len(other) == 0 {
		return false, false
	}
	if len(c) != len(other) {
		return false, true
	}
	for i, e := range c {
		eq, ok := e.Equal(other[i])
		if !ok || !eq {
			return false, ok
		}
	}
	return true, true
}

func (c Collection) Equivalent(other Collection) bool {
	if len(c) == 0 && len(other) == 0 {
		return true
	}
	if len(c) != len(other) {
		return false
	}

outer:
	for _, e := range c {
		for _, o := range other {
			if e.Equivalent(o) {
				continue outer
			}
		}
		return false
	}
	return true
}

func (c Collection) Cmp(other Collection) (cmp int, ok bool, err error) {
	if len(c) == 0 || len(other) == 0 {
		return 0, false, nil
	}
	if len(c) != 1 || len(other) != 1 {
		return 0, false, fmt.Errorf("can not compare collections with len != 1: %v and %v", c, other)
	}

	left, ok := c[0].(cmpElement)
	if !ok {
		primitive, _ := toPrimitive(c[0])
		left, ok = primitive.(cmpElement)
	}
	if !ok {
		return 0, false, errors.New("only strings, integers, decimals, quantities, dates, datetimes and times can be compared")
	}
	right := other[0]

	return left.Cmp(right)
}
func (c Collection) Union(other Collection) Collection {
	// If the input collection is empty, return the other collection
	if len(c) == 0 {
		return slices.Clone(other)
	}

	// If the other collection is empty, return the input collection
	if len(other) == 0 {
		return slices.Clone(c)
	}

	var union Collection

	// add elements from the first collection
outer1:
	for _, e := range c {
		for _, u := range union {
			eq, ok := e.Equal(u)
			if ok && eq {
				break outer1
			}
		}
		union = append(union, e)
	}

	// add elements from the second collection
outer2:
	for _, e := range other {
		for _, u := range union {
			eq, ok := e.Equal(u)
			if ok && eq {
				break outer2
			}
		}
		union = append(union, e)
	}

	return union
}

func (c Collection) Combine(other Collection) Collection {
	// If the input collection is empty, return the other collection
	if len(c) == 0 {
		return slices.Clone(other)
	}

	// If the other collection is empty, return the input collection
	if len(other) == 0 {
		return slices.Clone(c)
	}

	// Combine the two collections without eliminating duplicates
	combined := slices.Clone(c)
	combined = append(combined, other...)

	return combined
}
func (c Collection) Contains(element Element) bool {
	for _, e := range c {
		eq, ok := e.Equal(element)
		if ok && eq {
			return true
		}
	}
	return false
}

func (c Collection) Multiply(ctx context.Context, other Collection) (Collection, error) {
	if len(c) == 0 || len(other) == 0 {
		return nil, nil
	}
	if len(c) != 1 {
		return nil, fmt.Errorf("left value for multiplication has len != 1: %v", c)
	}
	if len(other) != 1 {
		return nil, fmt.Errorf("right value for multiplication has len != 1: %v", other)
	}

	left, ok := c[0].(multiplyElement)
	if !ok {
		primitive, _ := toPrimitive(c[0])
		left, ok = primitive.(multiplyElement)
	}
	if !ok {
		return nil, errors.New("can only multiply Integer, Decimal or Quantity")
	}
	right := other[0]

	res, err := left.Multiply(ctx, right)
	if err != nil {
		return nil, err
	}
	return Collection{res}, nil
}

func (c Collection) Divide(ctx context.Context, other Collection) (Collection, error) {
	if len(c) == 0 || len(other) == 0 {
		return nil, nil
	}
	if len(c) != 1 {
		return nil, fmt.Errorf("left value for division has len != 1: %v", c)
	}
	if len(other) != 1 {
		return nil, fmt.Errorf("right value for division has len != 1: %v", other)
	}

	left, ok := c[0].(divideElement)
	if !ok {
		primitive, _ := toPrimitive(c[0])
		left, ok = primitive.(divideElement)
	}
	if !ok {
		return nil, errors.New("can only divide Integer, Decimal or Quantity")
	}
	right := other[0]

	res, err := left.Divide(ctx, right)
	if err != nil {
		return nil, err
	}
	return Collection{res}, nil
}

func (c Collection) Div(ctx context.Context, other Collection) (Collection, error) {
	if len(c) == 0 || len(other) == 0 {
		return nil, nil
	}
	if len(c) != 1 {
		return nil, fmt.Errorf("left value for div has len != 1: %v", c)
	}
	if len(other) != 1 {
		return nil, fmt.Errorf("right value for div has len != 1: %v", other)
	}

	left, ok := c[0].(divElement)
	if !ok {
		primitive, _ := toPrimitive(c[0])
		left, ok = primitive.(divElement)
	}
	if !ok {
		return nil, errors.New("can only div Integer, Decimal")
	}
	right := other[0]

	res, err := left.Div(ctx, right)
	if err != nil {
		return nil, err
	}
	return Collection{res}, nil
}

func (c Collection) Mod(ctx context.Context, other Collection) (Collection, error) {
	if len(c) == 0 || len(other) == 0 {
		return nil, nil
	}
	if len(c) != 1 {
		return nil, fmt.Errorf("left value for div has len != 1: %v", c)
	}
	if len(other) != 1 {
		return nil, fmt.Errorf("right value for div has len != 1: %v", other)
	}

	left, ok := c[0].(modElement)
	if !ok {
		primitive, _ := toPrimitive(c[0])
		left, ok = primitive.(modElement)
	}
	if !ok {
		return nil, errors.New("can only div Integer, Decimal")
	}
	right := other[0]

	res, err := left.Mod(ctx, right)
	if err != nil {
		return nil, err
	}
	return Collection{res}, nil
}

func (c Collection) Add(ctx context.Context, other Collection) (Collection, error) {
	if len(c) == 0 || len(other) == 0 {
		return nil, nil
	}
	if len(c) != 1 {
		return nil, fmt.Errorf("left value for addition has len != 1: %v", c)
	}
	if len(other) != 1 {
		return nil, fmt.Errorf("right value for addition has len != 1: %v", other)
	}

	left, ok := c[0].(addElement)
	if !ok {
		primitive, _ := toPrimitive(c[0])
		left, ok = primitive.(addElement)
	}
	if !ok {
		return nil, errors.New("can only div Integer, Decimal, Quantity and String")
	}
	right := other[0]

	res, err := left.Add(ctx, right)
	if err != nil {
		return nil, err
	}
	return Collection{res}, nil
}

func (c Collection) Subtract(ctx context.Context, other Collection) (Collection, error) {
	if len(c) == 0 || len(other) == 0 {
		return nil, nil
	}
	if len(c) != 1 {
		return nil, fmt.Errorf("left value for subtract has len != 1: %v", c)
	}
	if len(other) != 1 {
		return nil, fmt.Errorf("right value for subtract has len != 1: %v", other)
	}

	left, ok := c[0].(subtractElement)
	if !ok {
		primitive, _ := toPrimitive(c[0])
		left, ok = primitive.(subtractElement)
	}
	if !ok {
		return nil, errors.New("can only div Integer, Decimal, Quantity")
	}
	right := other[0]

	res, err := left.Subtract(ctx, right)
	if err != nil {
		return nil, err
	}
	return Collection{res}, nil
}

func (c Collection) Concat(ctx context.Context, other Collection) (Collection, error) {
	if len(c) > 1 {
		return nil, fmt.Errorf("left value for concat has len > 1: %v", c)
	}
	if len(other) > 1 {
		return nil, fmt.Errorf("right value for concat has len > 1: %v", other)
	}
	if len(c) == 0 && len(other) == 0 {
		return Collection{String("")}, nil
	}

	var left, right String
	if len(c) == 1 {
		s, ok := c[0].(String)
		if !ok {
			return nil, fmt.Errorf("can only concat String, got left %T: %v", c[0], c[0])
		}
		left = s
	}
	if len(other) == 1 {
		s, ok := other[0].(String)
		if !ok {
			return nil, fmt.Errorf("can only concat String, got right %T: %v", other[0], other[0])
		}
		right = s
	}
	return Collection{left + right}, nil
}

func (c Collection) String() string {
	if len(c) == 0 {
		return "{ }"
	}

	var b strings.Builder
	b.WriteString("{ ")

	for _, e := range c[:len(c)-1] {
		// strings.Builder Write implementation does not return error
		_, _ = fmt.Fprint(&b, e, ", ")
	}
	_, _ = fmt.Fprint(&b, c[len(c)-1])

	b.WriteString(" }")
	return b.String()
}

type Boolean bool

func (b Boolean) Children(name ...string) Collection {
	return nil
}

func (b Boolean) ToBoolean(explicit bool) (v Boolean, ok bool, err error) {
	return b, true, nil
}
func (b Boolean) ToString(explicit bool) (v String, ok bool, err error) {
	if explicit {
		return String(b.String()), true, nil
	}
	return "", false, implicitConversionError[Boolean, String](b)
}
func (b Boolean) ToInteger(explicit bool) (v Integer, ok bool, err error) {
	if explicit {
		if b {
			return 1, true, nil
		} else {
			return 0, true, nil
		}
	}
	return 0, false, implicitConversionError[Boolean, Integer](b)
}
func (b Boolean) ToDecimal(explicit bool) (v Decimal, ok bool, err error) {
	if explicit {
		if b {
			return Decimal{Value: apd.New(10, -1)}, true, nil
		} else {
			return Decimal{Value: apd.New(00, -1)}, true, nil
		}
	}
	return Decimal{}, false, implicitConversionError[Boolean, Decimal](b)
}

func (b Boolean) ToDate(explicit bool) (v Date, ok bool, err error) {
	return Date{}, false, conversionError[Boolean, Date]()
}
func (b Boolean) ToTime(explicit bool) (v Time, ok bool, err error) {
	return Time{}, false, conversionError[Boolean, Time]()
}
func (b Boolean) ToDateTime(explicit bool) (v DateTime, ok bool, err error) {
	return DateTime{}, false, conversionError[Boolean, DateTime]()
}
func (b Boolean) ToQuantity(explicit bool) (v Quantity, ok bool, err error) {
	if explicit {
		if b {
			return Quantity{Value: Decimal{Value: apd.New(10, -1)}, Unit: "1"}, true, nil
		} else {
			return Quantity{Value: Decimal{Value: apd.New(00, -1)}, Unit: "1"}, true, nil
		}
	}
	return Quantity{}, false, conversionError[Boolean, Quantity]()
}
func (b Boolean) Equal(other Element, _noReverseTypeConversion ...bool) (eq bool, ok bool) {
	o, ok, err := other.ToBoolean(false)
	if err == nil && ok {
		return b == o, true
	}
	if len(_noReverseTypeConversion) == 0 || !_noReverseTypeConversion[0] {
		return other.Equal(b, true)
	} else {
		return false, true
	}
}
func (b Boolean) Equivalent(other Element, _noReverseTypeConversion ...bool) bool {
	eq, ok := b.Equal(other, _noReverseTypeConversion...)
	return ok && eq == true
}
func (b Boolean) TypeInfo() TypeInfo {
	return SimpleTypeInfo{
		Namespace: "System",
		Name:      "Boolean",
		BaseType:  TypeSpecifier{Namespace: "System", Name: "Any"},
	}
}
func (b Boolean) String() string {
	return strconv.FormatBool(bool(b))
}

type String string

func (s String) Children(name ...string) Collection {
	return nil
}

func (s String) ToBoolean(explicit bool) (v Boolean, ok bool, err error) {
	if explicit {
		if slices.Contains([]string{"true", "t", "yes", "y", "1", "1.0"}, strings.ToLower(string(s))) {
			return true, true, nil
		} else if slices.Contains([]string{"false", "f", "no", "n", "0", "0.0"}, strings.ToLower(string(s))) {
			return false, true, nil
		} else {
			return false, false, nil
		}
	}
	return false, false, implicitConversionError[String, Boolean](s)
}
func (s String) ToString(explicit bool) (v String, ok bool, err error) {
	return s, true, nil
}
func (s String) ToInteger(explicit bool) (v Integer, ok bool, err error) {
	if explicit {
		i, err := strconv.Atoi(string(s))
		if err != nil {
			return 0, false, nil
		}
		return Integer(i), true, nil
	}
	return 0, false, implicitConversionError[String, Integer](s)
}
func (s String) ToDecimal(explicit bool) (v Decimal, ok bool, err error) {
	if explicit {
		d, _, err := apd.NewFromString(string(s))
		if err != nil {
			return Decimal{}, false, nil
		}
		return Decimal{Value: d}, true, nil
	}
	return Decimal{}, false, implicitConversionError[String, Decimal](s)
}
func (s String) ToDate(explicit bool) (v Date, ok bool, err error) {
	if explicit {
		d, err := ParseDate(string(s))
		if err != nil {
			return Date{}, false, nil
		}
		return d, true, nil
	}
	return Date{}, false, implicitConversionError[String, Date](s)
}
func (s String) ToTime(explicit bool) (v Time, ok bool, err error) {
	if explicit {
		d, err := parseTime(string(s), false)
		if err != nil {
			return Time{}, false, nil
		}
		return d, true, nil
	}
	return Time{}, false, implicitConversionError[String, Time](s)
}
func (s String) ToDateTime(explicit bool) (v DateTime, ok bool, err error) {
	if explicit {
		d, err := ParseDateTime(string(s))
		if err != nil {
			return DateTime{}, false, nil
		}
		return d, true, nil
	}
	return DateTime{}, false, implicitConversionError[String, DateTime](s)
}
func (s String) ToQuantity(explicit bool) (v Quantity, ok bool, err error) {
	q, err := ParseQuantity(string(s))
	if err != nil {
		return Quantity{}, false, nil
	}
	return q, true, nil
}
func (s String) Equal(other Element, _noReverseTypeConversion ...bool) (eq bool, ok bool) {
	o, ok, err := other.ToString(false)
	if err == nil && ok {
		return s == o, true
	}
	if len(_noReverseTypeConversion) == 0 || !_noReverseTypeConversion[0] {
		return other.Equal(s, true)
	} else {
		return false, true
	}
}

var whitespaceReplaceRegex = regexp.MustCompile("[\t\r\n]")

func (s String) Equivalent(other Element, _noReverseTypeConversion ...bool) bool {
	o, ok, err := other.ToString(false)
	if err == nil && ok {
		return whitespaceReplaceRegex.ReplaceAllString(strings.ToLower(string(s)), " ") ==
			whitespaceReplaceRegex.ReplaceAllString(strings.ToLower(string(o)), " ")
	}
	if len(_noReverseTypeConversion) == 0 || !_noReverseTypeConversion[0] {
		return other.Equivalent(s, true)
	} else {
		return false
	}
}
func (s String) Cmp(other Element) (cmp int, ok bool, err error) {
	o, ok, err := other.ToString(false)
	if err != nil || !ok {
		return 0, false, fmt.Errorf("can not compare String to %T, left: %v right: %v", other, s, other)
	}
	return strings.Compare(string(s), string(o)), true, nil
}
func (s String) Add(ctx context.Context, other Element) (Element, error) {
	o, ok, err := other.ToString(false)
	if err != nil {
		return nil, fmt.Errorf("can not add %T to String, %v + %v", other, s, other)
	}
	if !ok {
		return nil, nil
	}
	return s + o, nil
}
func (s String) TypeInfo() TypeInfo {
	return SimpleTypeInfo{
		Namespace: "System",
		Name:      "String",
		BaseType:  TypeSpecifier{Namespace: "System", Name: "Any"},
	}
}
func (s String) String() string {
	return fmt.Sprintf("'%s'", string(s))
}

var (
	// escapes not handled by strconv.Unquote
	unescapeReplacer = strings.NewReplacer(
		`\'`, `'`,
		"\\`", "`",
		`\/`, `/`,
	)
)

func unescape(s string) (string, error) {
	unescaped := unescapeReplacer.Replace(s)
	// handles \", \r, \n, \t, \f, \\, \uXXXX
	return strconv.Unquote(`"` + unescaped + `"`)
}

type Integer int32

func (i Integer) Children(name ...string) Collection {
	return nil
}

func (i Integer) ToBoolean(explicit bool) (v Boolean, ok bool, err error) {
	if explicit {
		switch i {
		case 0:
			return false, true, nil
		case 1:
			return true, true, nil
		default:
			return false, false, nil
		}
	}
	return false, false, implicitConversionError[Integer, Boolean](i)
}
func (i Integer) ToString(explicit bool) (v String, ok bool, err error) {
	return String(i.String()), true, nil
}
func (i Integer) ToInteger(explicit bool) (v Integer, ok bool, err error) {
	return i, true, nil
}
func (i Integer) ToDecimal(explicit bool) (v Decimal, ok bool, err error) {
	return Decimal{Value: apd.New(int64(i), 0)}, true, nil
}
func (i Integer) ToDate(explicit bool) (v Date, ok bool, err error) {
	return Date{}, false, conversionError[Integer, Date]()
}
func (i Integer) ToTime(explicit bool) (v Time, ok bool, err error) {
	return Time{}, false, conversionError[Integer, Time]()
}
func (i Integer) ToDateTime(explicit bool) (v DateTime, ok bool, err error) {
	return DateTime{}, false, conversionError[Integer, DateTime]()
}
func (i Integer) ToQuantity(explicit bool) (v Quantity, ok bool, err error) {
	return Quantity{
		Value: Decimal{
			Value: apd.New(int64(i), 0),
		},
		Unit: "1",
	}, true, nil
}
func (i Integer) Equal(other Element, _noReverseTypeConversion ...bool) (eq bool, ok bool) {
	o, ok, err := other.ToInteger(false)
	if err == nil && ok {
		return i == o, true
	}
	if len(_noReverseTypeConversion) == 0 || !_noReverseTypeConversion[0] {
		return other.Equal(i, true)
	} else {
		return false, false
	}
}
func (i Integer) Equivalent(other Element, _noReverseTypeConversion ...bool) bool {
	eq, ok := i.Equal(other, _noReverseTypeConversion...)
	return ok && eq == true
}
func (i Integer) Cmp(other Element) (cmp int, ok bool, err error) {
	d, _, _ := i.ToDecimal(false)
	cmp, ok, err = d.Cmp(other)
	if err != nil || !ok {
		return 0, false, fmt.Errorf("can not compare Integer to %T, left: %v right: %v", other, i, other)
	}
	return cmp, true, nil
}
func (i Integer) Multiply(ctx context.Context, other Element) (Element, error) {
	switch o := other.(type) {
	case Integer:
		result, ok := overflow.Mul32(int32(i), int32(o))
		if !ok {
			return nil, nil
		}
		return Integer(result), nil
	case Decimal:
		d, _, _ := i.ToDecimal(false)
		return d.Multiply(ctx, o)
	}
	return nil, fmt.Errorf("can not multiply Integer with %T: %v * %v", other, i, other)
}
func (i Integer) Divide(ctx context.Context, other Element) (Element, error) {
	d, _, _ := i.ToDecimal(false)
	return d.Divide(ctx, other)
}
func (i Integer) Div(ctx context.Context, other Element) (Element, error) {
	switch o := other.(type) {
	case Integer:
		result, ok := overflow.Div32(int32(i), int32(o))
		if !ok {
			return nil, nil
		}
		return Integer(result), nil
	case Decimal:
		d, _, _ := i.ToDecimal(false)
		return d.Div(ctx, o)
	}
	return nil, fmt.Errorf("can not div Integer with %T: %v div %v", other, i, other)
}
func (i Integer) Mod(ctx context.Context, other Element) (Element, error) {
	switch o := other.(type) {
	case Integer:
		result, ok := overflow.Mod32(int32(i), int32(o))
		if !ok {
			return nil, nil
		}
		return Integer(result), nil
	case Decimal:
		d, _, _ := i.ToDecimal(false)
		return d.Mod(ctx, o)
	}
	return nil, fmt.Errorf("can not mod Integer with %T: %v mod %v", other, i, other)
}
func (i Integer) Add(ctx context.Context, other Element) (Element, error) {
	switch o := other.(type) {
	case Integer:
		result, ok := overflow.Add32(int32(i), int32(o))
		if !ok {
			return nil, nil
		}
		return Integer(result), nil
	case Decimal:
		d, _, _ := i.ToDecimal(false)
		return d.Add(ctx, o)
	}
	return nil, fmt.Errorf("can not add Integer and %T: %v + %v", other, i, other)
}
func (i Integer) Subtract(ctx context.Context, other Element) (Element, error) {
	switch o := other.(type) {
	case Integer:
		result, ok := overflow.Sub32(int32(i), int32(o))
		if !ok {
			return nil, nil
		}
		return Integer(result), nil
	case Decimal:
		d, _, _ := i.ToDecimal(false)
		return d.Subtract(ctx, o)
	}
	return nil, fmt.Errorf("can not subtract %T from Integer: %v - %v", other, i, other)
}
func (i Integer) TypeInfo() TypeInfo {
	return SimpleTypeInfo{
		Namespace: "System",
		Name:      "Integer",
		BaseType:  TypeSpecifier{Namespace: "System", Name: "Any"},
	}
}
func (i Integer) String() string {
	return strconv.Itoa(int(i))
}

type Decimal struct {
	defaultConversionError[Decimal]
	Value *apd.Decimal
}

func (d Decimal) Children(name ...string) Collection {
	return nil
}

func (d Decimal) ToBoolean(explicit bool) (v Boolean, ok bool, err error) {
	if explicit {
		if d.Value.Cmp(apd.New(1, 0)) == 0 {
			return true, true, nil
		} else if d.Value.Cmp(apd.New(0, 0)) == 0 {
			return false, true, nil
		} else {
			return false, false, nil
		}
	}
	return false, false, implicitConversionError[Decimal, Boolean](d)
}
func (d Decimal) ToString(explicit bool) (v String, ok bool, err error) {
	return String(d.String()), true, nil
}
func (d Decimal) ToDecimal(explicit bool) (v Decimal, ok bool, err error) {
	return d, true, nil
}
func (d Decimal) ToQuantity(explicit bool) (v Quantity, ok bool, err error) {
	return Quantity{
		Value: d,
		Unit:  "1",
	}, true, nil
}
func (d Decimal) Equal(other Element, _noReverseTypeConversion ...bool) (eq bool, ok bool) {
	o, ok, err := other.ToDecimal(false)
	if err == nil && ok {
		return d.Value.Cmp(o.Value) == 0, true
	}
	if len(_noReverseTypeConversion) == 0 || !_noReverseTypeConversion[0] {
		return other.Equal(d, true)
	} else {
		return false, true
	}
}
func (d Decimal) Equivalent(other Element, _noReverseTypeConversion ...bool) bool {
	o, ok, err := other.ToDecimal(false)
	if err == nil && ok {
		prec := uint32(min(d.Value.NumDigits(), o.Value.NumDigits()))
		ctx := apd.BaseContext.WithPrecision(prec)
		var a, b apd.Decimal
		_, err = ctx.Round(&a, d.Value)
		if err != nil {
			return false
		}
		_, err = ctx.Round(&b, o.Value)
		if err != nil {
			return false
		}
		return a.Cmp(&b) == 0
	}
	if len(_noReverseTypeConversion) == 0 || !_noReverseTypeConversion[0] {
		eq, ok := other.Equal(d, true)
		return ok && eq
	} else {
		return false
	}
}
func (d Decimal) Cmp(other Element) (cmp int, ok bool, err error) {
	o, ok, err := other.ToDecimal(false)
	if err != nil || !ok {
		return 0, false, fmt.Errorf("can not compare Decimal to %T, left: %v right: %v", other, d, other)
	}
	return d.Value.Cmp(o.Value), true, nil
}
func (d Decimal) Multiply(ctx context.Context, other Element) (Element, error) {
	o, ok, err := other.ToDecimal(false)
	if err != nil || !ok {
		return nil, fmt.Errorf("can not multiply Decimal with %T: %v * %v", other, d, other)
	}
	var res apd.Decimal
	_, err = apdContext(ctx).Mul(&res, d.Value, o.Value)
	if err != nil {
		return nil, err
	}
	return Decimal{Value: &res}, nil
}
func (d Decimal) Divide(ctx context.Context, other Element) (Element, error) {
	o, ok, err := other.ToDecimal(false)
	if err != nil || !ok {
		return nil, fmt.Errorf("can not divide Decimal with %T: %v / %v", other, d, other)
	}
	if o.Value.IsZero() {
		return nil, nil
	}
	var res apd.Decimal
	_, err = apdContext(ctx).Quo(&res, d.Value, o.Value)
	if err != nil {
		return nil, err
	}
	return Decimal{Value: &res}, nil
}
func (d Decimal) Div(ctx context.Context, other Element) (Element, error) {
	o, ok, err := other.ToDecimal(false)
	if err != nil || !ok {
		return nil, fmt.Errorf("can not div Decimal with %T: %v div %v", other, d, other)
	}
	if o.Value.IsZero() {
		return nil, nil
	}
	var res apd.Decimal
	_, err = apdContext(ctx).QuoInteger(&res, d.Value, o.Value)
	if err != nil {
		return nil, err
	}
	return Decimal{Value: &res}, nil
}
func (d Decimal) Mod(ctx context.Context, other Element) (Element, error) {
	o, ok, err := other.ToDecimal(false)
	if err != nil || !ok {
		return nil, fmt.Errorf("can not mod Decimal with %T: %v mod %v", other, d, other)
	}
	if o.Value.IsZero() {
		return nil, nil
	}
	var res apd.Decimal
	_, err = apdContext(ctx).Rem(&res, d.Value, o.Value)
	if err != nil {
		return nil, err
	}
	return Decimal{Value: &res}, nil
}
func (d Decimal) Add(ctx context.Context, other Element) (Element, error) {
	o, ok, err := other.ToDecimal(false)
	if err != nil || !ok {
		return nil, fmt.Errorf("can not add Decimal and %T: %v + %v", other, d, other)
	}
	if o.Value.IsZero() {
		return nil, nil
	}
	var res apd.Decimal
	_, err = apdContext(ctx).Add(&res, d.Value, o.Value)
	if err != nil {
		return nil, err
	}
	return Decimal{Value: &res}, nil
}
func (d Decimal) Subtract(ctx context.Context, other Element) (Element, error) {
	o, ok, err := other.ToDecimal(false)
	if err != nil || !ok {
		return nil, fmt.Errorf("can not subtract %T from Decimal: %v - %v", other, d, other)
	}
	if o.Value.IsZero() {
		return nil, nil
	}
	var res apd.Decimal
	_, err = apdContext(ctx).Sub(&res, d.Value, o.Value)
	if err != nil {
		return nil, err
	}
	return Decimal{Value: &res}, nil
}
func (d Decimal) TypeInfo() TypeInfo {
	return SimpleTypeInfo{
		Namespace: "System",
		Name:      "Decimal",
		BaseType:  TypeSpecifier{Namespace: "System", Name: "Any"},
	}
}
func (d Decimal) String() string {
	return d.Value.Text('f')
}

type Date struct {
	defaultConversionError[Date]
	Value     time.Time
	Precision DatePrecision
}

type DatePrecision string

const (
	DatePrecisionYear  = "year"
	DatePrecisionMonth = "month"
	DatePrecisionFull  = "full"
)

func (d Date) Children(name ...string) Collection {
	return nil
}
func (d Date) ToString(explicit bool) (v String, ok bool, err error) {
	return String(d.String()), true, nil
}
func (d Date) ToDate(explicit bool) (v Date, ok bool, err error) {
	return d, true, nil
}
func (d Date) ToDateTime(explicit bool) (v DateTime, ok bool, err error) {
	return DateTime{
		Value:     d.Value,
		Precision: DateTimePrecisionDay,
	}, true, nil
}
func (d Date) Equal(other Element, _noReverseTypeConversion ...bool) (eq bool, ok bool) {
	o, ok, err := other.ToDate(false)
	if err == nil && ok {
		cmp, ok, err := d.Cmp(o)
		if err == nil {
			return cmp == 0, ok
		}
	}
	if len(_noReverseTypeConversion) == 0 || !_noReverseTypeConversion[0] {
		return other.Equal(d, true)
	} else {
		return false, true
	}
}
func (d Date) Equivalent(other Element, _noReverseTypeConversion ...bool) bool {
	o, ok, err := other.ToDate(false)
	if err == nil && ok && d.Precision == o.Precision {
		eq, ok := d.Equal(other, _noReverseTypeConversion...)
		return ok && eq == true
	}
	return false
}
func (d Date) Cmp(other Element) (cmp int, ok bool, err error) {
	o, ok, err := other.ToDate(false)
	if err != nil || !ok {
		return 0, false, fmt.Errorf("can not compare Date to %T, left: %v right: %v", other, d, other)
	}
	switch d.Precision {
	case DatePrecisionYear:
		if o.Precision != DatePrecisionYear {
			return 0, false, nil
		}
		if d.Value.Year() < o.Value.In(d.Value.Location()).Year() {
			return -1, true, nil
		} else if d.Value.Year() > o.Value.In(d.Value.Location()).Year() {
			return 1, true, nil
		} else {
			return 0, true, nil
		}
	case DatePrecisionMonth:
		if o.Precision != DatePrecisionMonth {
			return 0, false, nil
		}
		if d.Value.Year() < o.Value.In(d.Value.Location()).Year() {
			return -1, true, nil
		} else if d.Value.Year() > o.Value.In(d.Value.Location()).Year() {
			return 1, true, nil
		} else if d.Value.Month() < o.Value.In(d.Value.Location()).Month() {
			return -1, true, nil
		} else if d.Value.Month() > o.Value.In(d.Value.Location()).Month() {
			return 1, true, nil
		} else {
			return 0, true, nil
		}
	default:
		if o.Precision != d.Precision {
			return 0, false, nil
		}
		if d.Value.Year() < o.Value.In(d.Value.Location()).Year() {
			return -1, true, nil
		} else if d.Value.Year() > o.Value.In(d.Value.Location()).Year() {
			return 1, true, nil
		} else if d.Value.Month() < o.Value.In(d.Value.Location()).Month() {
			return -1, true, nil
		} else if d.Value.Month() > o.Value.In(d.Value.Location()).Month() {
			return 1, true, nil
		} else if d.Value.Day() < o.Value.In(d.Value.Location()).Day() {
			return -1, true, nil
		} else if d.Value.Day() > o.Value.In(d.Value.Location()).Day() {
			return 1, true, nil
		} else {
			return 0, true, nil
		}
	}
}

// Add implements date arithmetic for Date values
func (d Date) Add(ctx context.Context, other Element) (Element, error) {
	// Check for empty date
	if d.Value.IsZero() {
		return nil, fmt.Errorf("cannot perform arithmetic on empty date")
	}

	q, ok, err := other.ToQuantity(false)
	if err != nil || !ok {
		return nil, fmt.Errorf("can not add Date with %T: %v + %v", other, d, other)
	}

	unit := normalizeTimeUnit(string(q.Unit))
	if !isTimeUnit(unit) {
		return nil, fmt.Errorf("invalid time unit: %v", q.Unit)
	}

	// Get the value for the quantity, ignoring decimal portion for calendar durations
	var integ, frac apd.Decimal
	q.Value.Value.Modf(&integ, &frac)
	value, err := integ.Int64()
	if err != nil {
		return nil, fmt.Errorf("invalid quantity value for date arithmetic: %v", err)
	}

	// Perform calendar-based arithmetic (negate the value for subtraction)
	var result time.Time
	switch unit {
	case UnitYear:
		result = d.Value.AddDate(int(value), 0, 0)
		// If the month and day of the date or time value is not a valid date in the resulting year,
		// the last day of the calendar month is used.
		if result.Day() < d.Value.Day() {
			result = result.AddDate(0, 0, -result.Day())
		}
	case UnitMonth:
		years, months := value/12, value%12
		result = d.Value.AddDate(int(years), int(months), 0)
		// If the resulting date is not a valid date in the resulting year,
		// the last day of the resulting calendar month is used.
		if result.Day() < d.Value.Day() {
			result = result.AddDate(0, 0, -result.Day())
		}
	case UnitWeek:
		result = d.Value.AddDate(0, 0, int(value)*7)
	case UnitDay:
		result = d.Value.AddDate(0, 0, int(value))
	default:
		return nil, fmt.Errorf("invalid time unit for Date: %v", q.Unit)
	}

	return Date{Value: result, Precision: d.Precision}, nil
}

// Subtract implements date arithmetic for Date values
func (d Date) Subtract(ctx context.Context, other Element) (Element, error) {
	// Check for empty date
	if d.Value.IsZero() {
		return nil, fmt.Errorf("cannot perform arithmetic on empty date")
	}

	q, ok, err := other.ToQuantity(false)
	if err != nil || !ok {
		return nil, fmt.Errorf("can not subtract from Date with %T: %v - %v", other, d, other)
	}

	unit := normalizeTimeUnit(string(q.Unit))
	if !isTimeUnit(unit) {
		return nil, fmt.Errorf("invalid time unit: %v", q.Unit)
	}

	// Get the value for the quantity, ignoring decimal portion for calendar durations
	var integ, frac apd.Decimal
	q.Value.Value.Modf(&integ, &frac)
	value, err := integ.Int64()
	if err != nil {
		return nil, fmt.Errorf("invalid quantity value for date arithmetic: %v", err)
	}

	// Perform calendar-based arithmetic (negate the value for subtraction)
	var result time.Time
	switch unit {
	case UnitYear:
		result = d.Value.AddDate(int(-value), 0, 0)
		// If the month and day of the date or time value is not a valid date in the resulting year,
		// the last day of the calendar month is used.
		if result.Day() < d.Value.Day() {
			result = result.AddDate(0, 0, -result.Day())
		}
	case UnitMonth:
		years, months := value/12, value%12
		result = d.Value.AddDate(int(-years), int(-months), 0)
		// If the resulting date is not a valid date in the resulting year,
		// the last day of the resulting calendar month is used.
		if result.Day() < d.Value.Day() {
			result = result.AddDate(0, 0, -result.Day())
		}
	case UnitWeek:
		result = d.Value.AddDate(0, 0, int(-value)*7)
	case UnitDay:
		result = d.Value.AddDate(0, 0, int(-value))
	default:
		return nil, fmt.Errorf("invalid time unit for Date: %v", q.Unit)
	}

	return Date{Value: result, Precision: d.Precision}, nil
}

func (d Date) TypeInfo() TypeInfo {
	return SimpleTypeInfo{
		Namespace: "System",
		Name:      "Date",
		BaseType:  TypeSpecifier{Namespace: "System", Name: "Any"},
	}
}
func (d Date) String() string {
	var ds string
	switch d.Precision {
	case DatePrecisionYear:
		ds = d.Value.Format(DateFormatOnlyYear)
	case DatePrecisionMonth:
		ds = d.Value.Format(DateFormatUpToMonth)
	default:
		ds = d.Value.Format(DateFormatFull)
	}
	return fmt.Sprintf("%s", ds)
}

type Time struct {
	defaultConversionError[Time]
	Value     time.Time
	Precision TimePrecision
}

type TimePrecision string

const (
	TimePrecisionHour   = "hour"
	TimePrecisionMinute = "minute"
	TimePrecisionFull   = "full"
)

func (t Time) Children(name ...string) Collection {
	return nil
}

func (t Time) ToString(explicit bool) (v String, ok bool, err error) {
	return String(t.String()), true, nil
}
func (t Time) ToTime(explicit bool) (v Time, ok bool, err error) {
	return t, true, nil
}
func (t Time) Equal(other Element, _noReverseTypeConversion ...bool) (eq bool, ok bool) {
	o, ok, err := other.ToTime(false)
	if err == nil && ok {
		cmp, ok, err := t.Cmp(o)
		if err == nil {
			return cmp == 0, ok
		}
	}
	if len(_noReverseTypeConversion) == 0 || !_noReverseTypeConversion[0] {
		return other.Equal(t, true)
	} else {
		return false, true
	}
}
func (t Time) Equivalent(other Element, _noReverseTypeConversion ...bool) bool {
	o, ok, err := other.ToTime(false)
	if err == nil && ok && t.Precision == o.Precision {
		eq, ok := t.Equal(other, _noReverseTypeConversion...)
		return ok && eq
	}
	return false
}
func (t Time) Cmp(other Element) (cmp int, ok bool, err error) {
	o, ok, err := other.ToTime(false)
	if err != nil || !ok {
		return 0, false, fmt.Errorf("can not compare Time to %T, left: %v right: %v", other, t, other)
	}
	switch t.Precision {
	case TimePrecisionHour:
		if o.Precision != TimePrecisionHour {
			return 0, false, nil
		}
		return t.Value.Truncate(time.Hour).Compare(
			o.Value.In(t.Value.Location()).Truncate(time.Hour)), true, nil
	case TimePrecisionMinute:
		if o.Precision != TimePrecisionMinute {
			return 0, false, nil
		}
		return t.Value.Truncate(time.Minute).Compare(
			o.Value.In(t.Value.Location()).Truncate(time.Minute)), true, nil
	default:
		if o.Precision != t.Precision {
			return 0, false, nil
		}
		return t.Value.Compare(o.Value.In(t.Value.Location())), true, nil
	}
}

// Add implements time arithmetic for Time values
func (t Time) Add(ctx context.Context, other Element) (Element, error) {
	// Check for empty time
	if t.Value.IsZero() {
		return nil, fmt.Errorf("cannot perform arithmetic on empty time")
	}

	q, ok, err := other.ToQuantity(false)
	if err != nil || !ok {
		return nil, fmt.Errorf("can not add Time with %T: %v + %v", other, t, other)
	}

	unit := normalizeTimeUnit(string(q.Unit))
	if !isTimeUnit(unit) {
		return nil, fmt.Errorf("invalid time unit: %v", q.Unit)
	}

	// For calendar durations, truncate decimal values
	var integ, frac apd.Decimal
	q.Value.Value.Modf(&integ, &frac)
	value, err := integ.Int64()
	if err != nil {
		return nil, fmt.Errorf("invalid quantity value for date arithmetic: %v", err)
	}

	var result time.Time
	switch unit {
	case UnitHour:
		result = t.Value.Add(time.Duration(value * int64(time.Hour)))
	case UnitMinute:
		result = t.Value.Add(time.Duration(value * int64(time.Minute)))
	case UnitSecond:
		seconds, err := q.Value.Value.Float64()
		if err != nil {
			return nil, fmt.Errorf("invalid quantity value for datetime arithmetic: %v", err)
		}
		result = t.Value.Add(time.Duration(seconds * float64(time.Second)))
	case UnitMillisecond:
		milliseconds, err := q.Value.Value.Float64()
		if err != nil {
			return nil, fmt.Errorf("invalid quantity value for datetime arithmetic: %v", err)
		}
		result = t.Value.Add(time.Duration(milliseconds * float64(time.Millisecond)))
	default:
		return nil, fmt.Errorf("invalid time unit for Time: %v", q.Unit)
	}

	return Time{Value: result, Precision: t.Precision}, nil
}

// Subtract implements time arithmetic for Time values
func (t Time) Subtract(ctx context.Context, other Element) (Element, error) {
	// Check for empty time
	if t.Value.IsZero() {
		return nil, fmt.Errorf("cannot perform arithmetic on empty time")
	}

	q, ok, err := other.ToQuantity(false)
	if err != nil || !ok {
		return nil, fmt.Errorf("can not subtract from Time with %T: %v - %v", other, t, other)
	}

	unit := normalizeTimeUnit(string(q.Unit))
	if !isTimeUnit(unit) {
		return nil, fmt.Errorf("invalid time unit: %v", q.Unit)
	}

	// For calendar durations, truncate decimal values
	var integ, frac apd.Decimal
	q.Value.Value.Modf(&integ, &frac)
	value, err := integ.Int64()
	if err != nil {
		return nil, fmt.Errorf("invalid quantity value for date arithmetic: %v", err)
	}

	var result time.Time
	switch unit {
	case UnitHour:
		result = t.Value.Add(time.Duration(-value * int64(time.Hour)))
	case UnitMinute:
		result = t.Value.Add(time.Duration(-value * int64(time.Minute)))
	case UnitSecond:
		seconds, err := q.Value.Value.Float64()
		if err != nil {
			return nil, fmt.Errorf("invalid quantity value for datetime arithmetic: %v", err)
		}
		result = t.Value.Add(time.Duration(-seconds * float64(time.Second)))
	case UnitMillisecond:
		milliseconds, err := q.Value.Value.Float64()
		if err != nil {
			return nil, fmt.Errorf("invalid quantity value for datetime arithmetic: %v", err)
		}
		result = t.Value.Add(time.Duration(-milliseconds * float64(time.Millisecond)))
	default:
		return nil, fmt.Errorf("invalid time unit for Time: %v", q.Unit)
	}

	return Time{Value: result, Precision: t.Precision}, nil
}

func (t Time) TypeInfo() TypeInfo {
	return SimpleTypeInfo{
		Namespace: "System",
		Name:      "Time",
		BaseType:  TypeSpecifier{Namespace: "System", Name: "Any"},
	}
}
func (t Time) String() string {
	var ts string
	switch t.Precision {
	case TimePrecisionHour:
		ts = t.Value.Format(TimeFormatOnlyHour)
	case TimePrecisionMinute:
		ts = t.Value.Format(TimeFormatUpToMinute)
	default:
		ts = t.Value.Format(TimeFormatFull)
	}
	return fmt.Sprintf("%s", ts)
}

type DateTime struct {
	defaultConversionError[DateTime]
	Value     time.Time
	Precision DateTimePrecision
}

type DateTimePrecision string

const (
	DateTimePrecisionYear   = "year"
	DateTimePrecisionMonth  = "month"
	DateTimePrecisionDay    = "day"
	DateTimePrecisionHour   = "hour"
	DateTimePrecisionMinute = "minute"
	DateTimePrecisionFull   = "full"
)

func (dt DateTime) Children(name ...string) Collection {
	return nil
}
func (dt DateTime) ToString(explicit bool) (v String, ok bool, err error) {
	return String(dt.String()), true, nil
}
func (dt DateTime) ToDate(explicit bool) (v Date, ok bool, err error) {
	if explicit {
		var precision DatePrecision
		switch dt.Precision {
		case DateTimePrecisionYear, DateTimePrecisionMonth:
			precision = DatePrecision(dt.Precision)
		default:
			precision = DatePrecisionFull
		}
		return Date{
			Value:     dt.Value,
			Precision: precision,
		}, true, nil
	}
	return Date{}, false, implicitConversionError[DateTime, Date](dt)
}
func (dt DateTime) ToDateTime(explicit bool) (v DateTime, ok bool, err error) {
	return dt, true, nil
}
func (dt DateTime) Equal(other Element, _noReverseTypeConversion ...bool) (eq bool, ok bool) {
	o, ok, err := other.ToDateTime(false)
	if err == nil && ok {
		cmp, ok, err := dt.Cmp(o)
		if err == nil {
			return cmp == 0, ok
		}
	}
	if len(_noReverseTypeConversion) == 0 || !_noReverseTypeConversion[0] {
		return other.Equal(dt, true)
	} else {
		return false, true
	}
}
func (dt DateTime) Equivalent(other Element, _noReverseTypeConversion ...bool) bool {
	o, ok, err := other.ToDateTime(false)
	if err == nil && ok && dt.Precision == o.Precision {
		eq, ok := dt.Equal(other, _noReverseTypeConversion...)
		return ok && eq
	}
	return false
}
func (dt DateTime) Cmp(other Element) (cmp int, ok bool, err error) {
	o, ok, err := other.ToDateTime(false)
	if err != nil || !ok {
		return 0, false, fmt.Errorf("can not compare Time to %T, left: %v right: %v", other, dt, other)
	}
	switch dt.Precision {
	case DateTimePrecisionYear:
		if o.Precision != DateTimePrecisionYear {
			return 0, false, nil
		}
		if dt.Value.Year() < o.Value.In(dt.Value.Location()).Year() {
			return -1, true, nil
		} else if dt.Value.Year() > o.Value.In(dt.Value.Location()).Year() {
			return 1, true, nil
		} else {
			return 0, true, nil
		}
	case DateTimePrecisionMonth:
		if o.Precision != DateTimePrecisionMonth {
			return 0, false, nil
		}
		if dt.Value.Year() < o.Value.In(dt.Value.Location()).Year() {
			return -1, true, nil
		} else if dt.Value.Year() > o.Value.In(dt.Value.Location()).Year() {
			return 1, true, nil
		} else if dt.Value.Month() < o.Value.In(dt.Value.Location()).Month() {
			return -1, true, nil
		} else if dt.Value.Month() > o.Value.In(dt.Value.Location()).Month() {
			return 1, true, nil
		} else {
			return 0, true, nil
		}
	case DateTimePrecisionDay:
		if o.Precision != DateTimePrecisionDay {
			return 0, false, nil
		}
		if dt.Value.Year() < o.Value.In(dt.Value.Location()).Year() {
			return -1, true, nil
		} else if dt.Value.Year() > o.Value.In(dt.Value.Location()).Year() {
			return 1, true, nil
		} else if dt.Value.Month() < o.Value.In(dt.Value.Location()).Month() {
			return -1, true, nil
		} else if dt.Value.Month() > o.Value.In(dt.Value.Location()).Month() {
			return 1, true, nil
		} else if dt.Value.Day() < o.Value.In(dt.Value.Location()).Day() {
			return -1, true, nil
		} else if dt.Value.Day() > o.Value.In(dt.Value.Location()).Day() {
			return 1, true, nil
		} else {
			return 0, true, nil
		}
	case DateTimePrecisionHour:
		if o.Precision != DateTimePrecisionHour {
			return 0, false, nil
		}
		return dt.Value.Truncate(time.Hour).Compare(
			o.Value.In(dt.Value.Location()).Truncate(time.Hour)), true, nil
	case DateTimePrecisionMinute:
		if o.Precision != DateTimePrecisionMinute {
			return 0, false, nil
		}
		return dt.Value.Truncate(time.Minute).Compare(
			o.Value.In(dt.Value.Location()).Truncate(time.Minute)), true, nil
	default:
		if o.Precision != dt.Precision {
			return 0, false, nil
		}
		return dt.Value.Compare(
			o.Value.In(dt.Value.Location())), true, nil
	}
}

// Add implements date/time arithmetic for DateTime values
func (dt DateTime) Add(ctx context.Context, other Element) (Element, error) {
	// Check for empty datetime
	if dt.Value.IsZero() {
		return nil, fmt.Errorf("cannot perform arithmetic on empty datetime")
	}

	q, ok, err := other.ToQuantity(false)
	if err != nil || !ok {
		return nil, fmt.Errorf("can not add DateTime with %T: %v + %v", other, dt, other)
	}

	unit := normalizeTimeUnit(string(q.Unit))
	if !isTimeUnit(unit) {
		return nil, fmt.Errorf("invalid time unit: %v", q.Unit)
	}

	// For calendar durations, truncate decimal values
	var integ, frac apd.Decimal
	q.Value.Value.Modf(&integ, &frac)
	value, err := integ.Int64()
	if err != nil {
		return nil, fmt.Errorf("invalid quantity value for date arithmetic: %v", err)
	}

	var result time.Time
	switch unit {
	case UnitYear:
		result = dt.Value.AddDate(int(value), 0, 0)
		// If the month and day of the date or time value is not a valid date in the resulting year,
		// the last day of the calendar month is used.
		if result.Day() < dt.Value.Day() {
			result = result.AddDate(0, 0, -result.Day())
		}
	case UnitMonth:
		years, months := value/12, value%12
		result = dt.Value.AddDate(int(years), int(months), 0)
		// If the resulting date is not a valid date in the resulting year,
		// the last day of the resulting calendar month is used.
		if result.Day() < dt.Value.Day() {
			result = result.AddDate(0, 0, -result.Day())
		}
	case UnitWeek:
		result = dt.Value.AddDate(0, 0, int(value)*7)
	case UnitDay:
		result = dt.Value.AddDate(0, 0, int(value))
	case UnitHour:
		result = dt.Value.Add(time.Duration(value * int64(time.Hour)))
	case UnitMinute:
		result = dt.Value.Add(time.Duration(value * int64(time.Minute)))
	case UnitSecond:
		seconds, err := q.Value.Value.Float64()
		if err != nil {
			return nil, fmt.Errorf("invalid quantity value for datetime arithmetic: %v", err)
		}
		result = dt.Value.Add(time.Duration(seconds * float64(time.Second)))
	case UnitMillisecond:
		milliseconds, err := q.Value.Value.Float64()
		if err != nil {
			return nil, fmt.Errorf("invalid quantity value for datetime arithmetic: %v", err)
		}
		result = dt.Value.Add(time.Duration(milliseconds * float64(time.Millisecond)))
	}

	return DateTime{Value: result, Precision: dt.Precision}, nil
}

// Subtract implements date/time arithmetic for DateTime values
func (dt DateTime) Subtract(ctx context.Context, other Element) (Element, error) {
	// Check for empty datetime
	if dt.Value.IsZero() {
		return nil, fmt.Errorf("cannot perform arithmetic on empty datetime")
	}

	q, ok, err := other.ToQuantity(false)
	if err != nil || !ok {
		return nil, fmt.Errorf("can not subtract from DateTime with %T: %v - %v", other, dt, other)
	}

	unit := normalizeTimeUnit(string(q.Unit))
	if !isTimeUnit(unit) {
		return nil, fmt.Errorf("invalid time unit: %v", q.Unit)
	}

	// For calendar durations, truncate decimal values
	var integ, frac apd.Decimal
	q.Value.Value.Modf(&integ, &frac)
	value, err := integ.Int64()
	if err != nil {
		return nil, fmt.Errorf("invalid quantity value for date arithmetic: %v", err)
	}

	var result time.Time
	switch unit {
	case UnitYear:
		result = dt.Value.AddDate(int(-value), 0, 0)
		// If the month and day of the date or time value is not a valid date in the resulting year,
		// the last day of the calendar month is used.
		if result.Day() < dt.Value.Day() {
			result = result.AddDate(0, 0, -result.Day())
		}
	case UnitMonth:
		years, months := value/12, value%12
		result = dt.Value.AddDate(int(-years), int(-months), 0)
		// If the resulting date is not a valid date in the resulting year,
		// the last day of the resulting calendar month is used.
		if result.Day() < dt.Value.Day() {
			result = result.AddDate(0, 0, -result.Day())
		}
	case UnitWeek:
		result = dt.Value.AddDate(0, 0, int(-value)*7)
	case UnitDay:
		result = dt.Value.Add(time.Duration(-value * int64(time.Hour)))
	case UnitHour:
		result = dt.Value.Add(time.Duration(-value * int64(time.Hour)))
	case UnitMinute:
		result = dt.Value.Add(time.Duration(-value * int64(time.Minute)))
	case UnitSecond:
		seconds, err := q.Value.Value.Float64()
		if err != nil {
			return nil, fmt.Errorf("invalid quantity value for datetime arithmetic: %v", err)
		}
		result = dt.Value.Add(time.Duration(-seconds * float64(time.Second)))
	case UnitMillisecond:
		milliseconds, err := q.Value.Value.Float64()
		if err != nil {
			return nil, fmt.Errorf("invalid quantity value for datetime arithmetic: %v", err)
		}
		result = dt.Value.Add(time.Duration(-milliseconds * float64(time.Millisecond)))
	}

	return DateTime{Value: result, Precision: dt.Precision}, nil
}

func (dt DateTime) TypeInfo() TypeInfo {
	return SimpleTypeInfo{
		Namespace: "System",
		Name:      "DateTime",
		BaseType:  TypeSpecifier{Namespace: "System", Name: "Any"},
	}
}
func (dt DateTime) String() string {
	var ds, ts string
	switch dt.Precision {
	case DateTimePrecisionYear:
		ds = dt.Value.Format(DateFormatOnlyYear)
	case DateTimePrecisionMonth:
		ds = dt.Value.Format(DateFormatUpToMonth)
	case DateTimePrecisionDay:
		ds = dt.Value.Format(DateFormatFull)
	case DateTimePrecisionHour:
		ds = dt.Value.Format(DateFormatFull)
		ts = dt.Value.Format(TimeFormatOnlyHourTZ)
	case DateTimePrecisionMinute:
		ds = dt.Value.Format(DateFormatFull)
		ts = dt.Value.Format(TimeFormatUpToMinuteTZ)
	default:
		ds = dt.Value.Format(DateFormatFull)
		ts = dt.Value.Format(TimeFormatFullTZ)
	}

	return fmt.Sprintf("%sT%s", ds, ts)
}

const (
	DateFormatOnlyYear     = "2006"
	DateFormatUpToMonth    = "2006-01"
	DateFormatFull         = "2006-01-02"
	TimeFormatOnlyHour     = "15"
	TimeFormatOnlyHourTZ   = "15Z07:00"
	TimeFormatUpToMinute   = "15:04"
	TimeFormatUpToMinuteTZ = "15:04Z07:00"
	TimeFormatFull         = "15:04:05.999999999"
	TimeFormatFullTZ       = "15:04:05.999999999Z07:00"
)

func ParseDate(s string) (Date, error) {
	ds := strings.TrimLeft(s, "@")

	d, err := time.Parse(DateFormatOnlyYear, ds)
	if err == nil {
		return Date{Value: d, Precision: DatePrecisionYear}, nil
	}
	d, err = time.Parse(DateFormatUpToMonth, ds)
	if err == nil {
		return Date{Value: d, Precision: DatePrecisionMonth}, nil
	}
	d, err = time.Parse(DateFormatFull, ds)
	if err == nil {
		return Date{Value: d, Precision: DatePrecisionFull}, nil
	}

	return Date{}, fmt.Errorf("invalid Date format: %s", s)
}

func ParseTime(s string) (Time, error) {
	return parseTime(s, false)
}

func parseTime(s string, withTZ bool) (Time, error) {
	ts := strings.TrimLeft(s, "@T")

	t, err := time.Parse(TimeFormatOnlyHour, ts)
	if err == nil {
		return Time{Value: t, Precision: TimePrecisionHour}, nil
	}
	if withTZ {
		t, err = time.Parse(TimeFormatOnlyHourTZ, ts)
		if err == nil {
			return Time{Value: t, Precision: TimePrecisionHour}, nil
		}
	}
	t, err = time.Parse(TimeFormatUpToMinute, ts)
	if err == nil {
		return Time{Value: t, Precision: TimePrecisionMinute}, nil
	}
	if withTZ {
		t, err = time.Parse(TimeFormatUpToMinuteTZ, ts)
		if err == nil {
			return Time{Value: t, Precision: TimePrecisionMinute}, nil
		}
	}
	t, err = time.Parse(TimeFormatFull, ts)
	if err == nil {
		return Time{Value: t, Precision: TimePrecisionFull}, nil
	}
	if withTZ {
		t, err = time.Parse(TimeFormatFullTZ, ts)
		if err == nil {
			return Time{Value: t, Precision: TimePrecisionFull}, nil
		}
	}

	return Time{}, fmt.Errorf("invalid Date format: %s", s)
}

func ParseDateTime(s string) (DateTime, error) {
	splits := strings.Split(s, "T")

	ds := splits[0]
	d, err := ParseDate(ds)
	if err != nil {
		return DateTime{}, fmt.Errorf("invalid DateTime format (date part): %s", s)
	}

	if len(splits) == 1 || splits[1] == "" {
		if d.Precision == DateTimePrecisionFull {
			return DateTime{Value: d.Value, Precision: DateTimePrecisionDay}, nil
		}
		return DateTime{Value: d.Value, Precision: DateTimePrecision(d.Precision)}, nil
	}

	ts := splits[1]
	t, err := parseTime(ts, true)
	if err != nil {
		return DateTime{}, fmt.Errorf("invalid DateTime format (time part): %s", s)
	}

	tv := t.Value.In(d.Value.Location())

	dt := d.Value.Add(
		time.Hour*time.Duration(tv.Hour()) +
			time.Minute*time.Duration(tv.Minute()) +
			time.Second*time.Duration(tv.Second()) +
			time.Nanosecond*time.Duration(tv.Nanosecond()),
	)
	return DateTime{Value: dt, Precision: DateTimePrecision(t.Precision)}, nil
}

// Time units for date/time arithmetic
const (
	UnitYear         = "year"
	UnitYears        = "years"
	UnitMonth        = "month"
	UnitMonths       = "months"
	UnitWeek         = "week"
	UnitWeeks        = "weeks"
	UnitDay          = "day"
	UnitDays         = "days"
	UnitHour         = "hour"
	UnitHours        = "hours"
	UnitMinute       = "minute"
	UnitMinutes      = "minutes"
	UnitSecond       = "second"
	UnitSeconds      = "seconds"
	UnitS            = "s"
	UnitMillisecond  = "millisecond"
	UnitMilliseconds = "milliseconds"
	UnitMs           = "ms"
)

// isTimeUnit returns true if the unit is a valid time unit
func isTimeUnit(unit string) bool {
	switch unit {
	case UnitYear, UnitYears,
		UnitMonth, UnitMonths,
		UnitWeek, UnitWeeks,
		UnitDay, UnitDays,
		UnitHour, UnitHours,
		UnitMinute, UnitMinutes,
		UnitSecond, UnitSeconds, UnitS,
		UnitMillisecond, UnitMilliseconds, UnitMs:
		return true
	}
	return false
}

// normalizeTimeUnit returns the canonical form of a time unit
func normalizeTimeUnit(unit string) string {
	switch unit {
	case UnitYear, UnitYears:
		return UnitYear
	case UnitMonth, UnitMonths:
		return UnitMonth
	case UnitWeek, UnitWeeks:
		return UnitWeek
	case UnitDay, UnitDays:
		return UnitDay
	case UnitHour, UnitHours:
		return UnitHour
	case UnitMinute, UnitMinutes:
		return UnitMinute
	case UnitSecond, UnitSeconds, UnitS:
		return UnitSecond
	case UnitMillisecond, UnitMilliseconds, UnitMs:
		return UnitMillisecond
	}
	return unit
}

type Quantity struct {
	defaultConversionError[Quantity]
	Value Decimal
	Unit  String
}

func (q Quantity) Children(name ...string) Collection {
	var children Collection
	if len(name) == 0 || slices.Contains(name, "value") {
		children = append(children, q.Value)
	}
	if len(name) == 0 || slices.Contains(name, "unit") {
		children = append(children, q.Unit)
	}
	return children
}
func (q Quantity) ToString(explicit bool) (v String, ok bool, err error) {
	return String(q.String()), true, nil
}
func (q Quantity) ToQuantity(explicit bool) (v Quantity, ok bool, err error) {
	return q, true, nil
}
func (q Quantity) Equal(other Element, _noReverseTypeConversion ...bool) (eq bool, ok bool) {
	o, ok, err := other.ToQuantity(false)
	if err == nil && ok {
		left := q.dateAsUCUM()
		right := o.dateAsUCUM()
		eq, ok := left.Value.Equal(right.Value)
		return ok && eq, left.Unit == right.Unit
	}
	if len(_noReverseTypeConversion) == 0 || !_noReverseTypeConversion[0] {
		return other.Equal(q, true)
	} else {
		return false, true
	}
}
func (q Quantity) Equivalent(other Element, _noReverseTypeConversion ...bool) bool {
	eq, ok := q.Equal(other, _noReverseTypeConversion...)
	return ok && eq == true
}
func (q Quantity) Cmp(other Element) (cmp int, ok bool, err error) {
	o, ok, err := other.ToQuantity(false)
	if err != nil || !ok {
		return 0, false, fmt.Errorf("can not compare Quantity to %T, left: %v right: %v", other, q, other)
	}
	left := q.dateAsUCUM()
	right := o.dateAsUCUM()
	if left.Unit != right.Unit {
		return 0, false, fmt.Errorf("quantity units do not match, left: %v right: %v", left.Unit, right.Unit)
	}
	return left.Value.Cmp(right.Value)
}
func (q Quantity) Multiply(ctx context.Context, other Element) (Element, error) {
	o, ok, err := other.ToQuantity(false)
	if err != nil || !ok {
		return nil, fmt.Errorf("can not multiply Quantity with %T: %v * %v", other, q, other)
	}
	left := q.dateAsUCUM()
	right := o.dateAsUCUM()

	value, err := left.Value.Multiply(ctx, right.Value)
	if err != nil {
		return Quantity{}, err
	}
	unit := fmt.Sprintf("(%s).(%s)", left.Unit, right.Unit)
	return Quantity{Value: value.(Decimal), Unit: String(unit)}, nil
}
func (q Quantity) Divide(ctx context.Context, other Element) (Element, error) {
	o, ok, err := other.ToQuantity(false)
	if err != nil || !ok {
		return nil, fmt.Errorf("can not divide Quantity with %T: %v / %v", other, q, other)
	}
	left := q.dateAsUCUM()
	right := o.dateAsUCUM()

	value, err := left.Value.Divide(ctx, right.Value)
	if err != nil {
		return Quantity{}, err
	}
	unit := fmt.Sprintf("(%s)/(%s)", left.Unit, right.Unit)
	return Quantity{Value: value.(Decimal), Unit: String(unit)}, nil
}
func (q Quantity) Add(ctx context.Context, other Element) (Element, error) {
	o, ok, err := other.ToQuantity(false)
	if err != nil || !ok {
		return nil, fmt.Errorf("can not add Quantity and %T: %v + %v", other, q, other)
	}
	left := q.dateAsUCUM()
	right := o.dateAsUCUM()

	if left.Unit != right.Unit {
		return Quantity{}, fmt.Errorf("quantity units do not match, left: %v right: %v", left, right)
	}

	value, err := left.Value.Add(ctx, right.Value)
	if err != nil {
		return Quantity{}, err
	}
	unit := fmt.Sprintf("(%s).(%s)", left.Unit, right.Unit)
	return Quantity{Value: value.(Decimal), Unit: String(unit)}, nil
}
func (q Quantity) Subtract(ctx context.Context, other Element) (Element, error) {
	o, ok, err := other.ToQuantity(false)
	if err != nil || !ok {
		return nil, fmt.Errorf("can not subtract %T from Quantity: %v - %v", other, q, other)
	}
	left := q.dateAsUCUM()
	right := o.dateAsUCUM()

	if left.Unit != right.Unit {
		return Quantity{}, fmt.Errorf("quantity units do not match, left: %v right: %v", left, right)
	}

	value, err := left.Value.Subtract(ctx, right.Value)
	if err != nil {
		return Quantity{}, err
	}
	unit := fmt.Sprintf("(%s).(%s)", left.Unit, right.Unit)
	return Quantity{Value: value.(Decimal), Unit: String(unit)}, nil
}
func (q Quantity) dateAsUCUM() Quantity {
	switch q.Unit {
	case "year":
		q.Unit = "a"
	case "month":
		q.Unit = "mo"
	case "week":
		q.Unit = "wk"
	case "day":
		q.Unit = "d"
	case "hour":
		q.Unit = "h"
	case "minute":
		q.Unit = "m"
	case "second":
		q.Unit = "s"
	case "millisecond":
		q.Unit = "ms"
	}
	return q
}
func (q Quantity) TypeInfo() TypeInfo {
	return ClassInfo{
		SimpleTypeInfo: SimpleTypeInfo{
			Namespace: "System",
			Name:      "Quantity",
			BaseType:  TypeSpecifier{Namespace: "System", Name: "Any"},
		},
		Element: []ClassInfoElement{
			{Name: "Value", Type: TypeSpecifier{Namespace: "System", Name: "Decimal"}},
			{Name: "Unit", Type: TypeSpecifier{Namespace: "System", Name: "String"}},
		},
	}
}
func (q Quantity) String() string {
	return fmt.Sprintf("%s %s", q.Value.String(), q.Unit)
}

func ParseQuantity(s string) (Quantity, error) {
	expr, err := Parse(s)
	if err != nil {
		return Quantity{}, fmt.Errorf("cannot parse quantity '%s': %v", s, err)
	}
	termCtx, ok := expr.tree.(*parser.TermExpressionContext)
	if !ok {
		return Quantity{}, fmt.Errorf("cannot parse quantity '%s'", s)
	}
	literalCtx, ok := termCtx.Term().(*parser.LiteralTermContext)
	if !ok {
		return Quantity{}, fmt.Errorf("cannot parse quantity '%s'", s)
	}

	switch lit := literalCtx.Literal().(type) {
	case *parser.QuantityLiteralContext:
		quantityCtx, ok := lit.Quantity().(*parser.QuantityContext)
		if !ok {
			return Quantity{}, fmt.Errorf("cannot parse quantity '%s'", s)
		}
		v, _, err := apd.NewFromString(quantityCtx.NUMBER().GetText())
		if err != nil {
			return Quantity{}, err
		}

		t := quantityCtx.Unit().GetText()
		u := strings.Trim(t, "'")

		return Quantity{Value: Decimal{Value: v}, Unit: String(u)}, nil
	case *parser.NumberLiteralContext:
		v, _, err := apd.NewFromString(lit.NUMBER().GetText())
		if err != nil {
			return Quantity{}, err
		}

		return Quantity{Value: Decimal{Value: v}, Unit: "1"}, nil
	default:
		return Quantity{}, fmt.Errorf("cannot parse quantity '%s'", s)
	}
}

type defaultConversionError[F any] struct {
}

func (_ defaultConversionError[F]) ToBoolean(explicit bool) (v Boolean, ok bool, err error) {
	return false, false, conversionError[F, Boolean]()
}
func (_ defaultConversionError[F]) ToString(explicit bool) (v String, ok bool, err error) {
	return "", false, conversionError[F, Boolean]()
}
func (_ defaultConversionError[F]) ToInteger(explicit bool) (v Integer, ok bool, err error) {
	return 0, false, conversionError[F, Integer]()
}
func (_ defaultConversionError[F]) ToDecimal(explicit bool) (v Decimal, ok bool, err error) {
	return Decimal{}, false, conversionError[F, Decimal]()
}
func (_ defaultConversionError[F]) ToDate(explicit bool) (v Date, ok bool, err error) {
	return Date{}, false, conversionError[F, Date]()
}
func (_ defaultConversionError[F]) ToTime(explicit bool) (v Time, ok bool, err error) {
	return Time{}, false, conversionError[F, Time]()
}
func (_ defaultConversionError[F]) ToDateTime(explicit bool) (v DateTime, ok bool, err error) {
	return DateTime{}, false, conversionError[F, DateTime]()
}
func (_ defaultConversionError[F]) ToQuantity(explicit bool) (v Quantity, ok bool, err error) {
	return Quantity{}, false, conversionError[F, Quantity]()
}

func conversionError[F any, T Element]() error {
	var (
		f F
		t T
	)
	return fmt.Errorf("primitive %v of type %T can not be converted to type %T", f, f, t)
}

func implicitConversionError[F Element, T Element](f F) error {
	var t T
	return fmt.Errorf("primitive %T %v can not be implicitly converted to %T", f, f, t)
}
