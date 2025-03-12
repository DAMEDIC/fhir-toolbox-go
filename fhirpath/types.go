package fhirpath

import (
	"context"
	"errors"
	"fmt"
	"github.com/DAMEDIC/fhir-toolbox-go/fhirpath/overflow"
	"github.com/DAMEDIC/fhir-toolbox-go/utils"
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
	ToBoolean(explicit bool) (*Boolean, error)
	ToString(explicit bool) (*String, error)
	ToInteger(explicit bool) (*Integer, error)
	ToDecimal(explicit bool) (*Decimal, error)
	ToDate(explicit bool) (*Date, error)
	ToTime(explicit bool) (*Time, error)
	ToDateTime(explicit bool) (*DateTime, error)
	ToQuantity(explicit bool) (*Quantity, error)
	Equal(other Element, _noReverseTypeConversion ...bool) bool
	Equivalent(other Element, _noReverseTypeConversion ...bool) bool
	TypeInfo() TypeInfo
	fmt.Stringer
}

type cmpElement interface {
	Element
	// Cmp may return nil, because attempting to operate on quantities
	// with invalid units will result in empty ({ }).
	Cmp(other Element) (*int, error)
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

// WithAPDContext sets the apd.Context for Decimal operations.
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
	QualifiedName() (TypeSpecifier, bool)
	BaseTypeName() (TypeSpecifier, bool)
}

type SimpleTypeInfo struct {
	Namespace string
	Name      string
	BaseType  TypeSpecifier
}

func (i SimpleTypeInfo) QualifiedName() (TypeSpecifier, bool) {
	return TypeSpecifier{i.Namespace, i.Name}, true
}
func (i SimpleTypeInfo) BaseTypeName() (TypeSpecifier, bool) {
	return i.BaseType, true
}

type ClassInfo struct {
	SimpleTypeInfo
	Element []ClassInfoElement
}

func (i ClassInfo) QualifiedName() (TypeSpecifier, bool) {
	return TypeSpecifier{i.Namespace, i.Name}, true
}
func (i ClassInfo) BaseTypeName() (TypeSpecifier, bool) {
	return i.BaseType, true
}

type ClassInfoElement struct {
	Name       string
	Type       string
	IsOneBased bool
}

type ListTypeInfo struct {
	ElementType string
}

func (i ListTypeInfo) QualifiedName() (TypeSpecifier, bool) {
	return TypeSpecifier{}, false
}
func (i ListTypeInfo) BaseTypeName() (TypeSpecifier, bool) {
	return TypeSpecifier{}, false
}

type TupleTypeInfo struct {
	Element []TupleTypeInfoElement
}

func (i TupleTypeInfo) QualifiedName() (TypeSpecifier, bool) {
	return TypeSpecifier{}, false
}
func (i TupleTypeInfo) BaseTypeName() (TypeSpecifier, bool) {
	return TypeSpecifier{}, false
}

type TupleTypeInfoElement struct {
	Name       string
	Type       string
	IsOneBased bool
}

type TypeSpecifier struct {
	Namespace string
	Name      string
}

func ParseTypeSpecifier(s string) TypeSpecifier {
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

func (t TypeSpecifier) String() string {
	if t.Namespace != "" {
		return fmt.Sprintf("%s.%s", t.Namespace, t.Name)
	} else {
		return t.Name
	}
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

func ElementTo[T Element](e Element, explicit bool) (*T, error) {
	var v *T
	switch any(v).(type) {
	case Boolean:
		v, err := e.ToBoolean(explicit)
		return any(v).(*T), err
	case String:
		v, err := e.ToString(explicit)
		return any(v).(*T), err
	case Integer:
		v, err := e.ToInteger(explicit)
		return any(v).(*T), err
	case Decimal:
		v, err := e.ToDecimal(explicit)
		return any(v).(*T), err
	case Date:
		v, err := e.ToDate(explicit)
		return any(v).(*T), err
	case Time:
		v, err := e.ToTime(explicit)
		return any(v).(*T), err
	case DateTime:
		v, err := e.ToDateTime(explicit)
		return any(v).(*T), err
	case Quantity:
		v, err := e.ToQuantity(explicit)
		return any(v).(*T), err
	default:
		return v, fmt.Errorf("can not convert to type %T", v)
	}
}

type Collection []Element

func (c Collection) Equal(other Collection) *bool {
	if len(c) == 0 || len(other) == 0 {
		return nil
	}
	if len(c) != len(other) {
		return utils.Ptr(false)
	}
	for i, e := range c {
		if !e.Equal(other[i]) {
			return utils.Ptr(false)
		}
	}
	return utils.Ptr(true)
}

func (c Collection) Equivalent(other Collection) *bool {
	if len(c) == 0 && len(other) == 0 {
		return utils.Ptr(true)
	}
	if len(c) != len(other) {
		return utils.Ptr(false)
	}

outer:
	for _, e := range c {
		for _, o := range other {
			if e.Equivalent(o) {
				continue outer
			}
		}
		return utils.Ptr(false)
	}
	return utils.Ptr(true)
}

func (c Collection) Cmp(other Collection) (*int, error) {
	if len(c) == 0 || len(other) == 0 {
		return nil, nil
	}
	if len(c) != 1 || len(other) != 1 {
		return nil, fmt.Errorf("can not compare collections with len != 1: %v and %v", c, other)
	}

	left, ok := c[0].(cmpElement)
	if !ok {
		return nil, errors.New("only strings, integers, decimals, quantities, dates, datetimes and times can be compared")
	}
	right := other[0]

	return left.Cmp(right)
}
func (c Collection) Union(other Collection) Collection {
	union := slices.Clone(c)
	for _, o := range other {
		for _, e := range union {
			if !o.Equal(e) {
				union = append(union, e)
			}
		}
	}
	return union
}
func (c Collection) Contains(element Element) bool {
	for _, e := range c {
		if e.Equal(element) {
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

func (b Boolean) ToBoolean(explicit bool) (*Boolean, error) {
	return &b, nil
}
func (b Boolean) ToString(explicit bool) (*String, error) {
	if explicit {
		return utils.Ptr(String(b.String())), nil
	}
	return nil, implicitConversionError[Boolean, String]()
}
func (b Boolean) ToInteger(explicit bool) (*Integer, error) {
	if explicit {
		if b {
			return utils.Ptr[Integer](1), nil
		} else {
			return utils.Ptr[Integer](0), nil
		}
	}
	return nil, implicitConversionError[Boolean, Integer]()
}
func (b Boolean) ToDecimal(explicit bool) (*Decimal, error) {
	if explicit {
		if b {
			return &Decimal{apd.New(10, -1)}, nil
		} else {
			return &Decimal{apd.New(00, -1)}, nil
		}
	}
	return nil, implicitConversionError[Boolean, Decimal]()
}
func (b Boolean) ToDate(explicit bool) (*Date, error) {
	return nil, conversionError[Boolean, Date]()
}
func (b Boolean) ToTime(explicit bool) (*Time, error) {
	return nil, conversionError[Boolean, Time]()
}
func (b Boolean) ToDateTime(explicit bool) (*DateTime, error) {
	return nil, conversionError[Boolean, DateTime]()
}
func (b Boolean) ToQuantity(explicit bool) (*Quantity, error) {
	if explicit {
		if b {
			return &Quantity{Value: Decimal{Value: apd.New(10, -1)}, Unit: "1"}, nil
		} else {
			return &Quantity{Value: Decimal{Value: apd.New(00, -1)}, Unit: "1"}, nil
		}
	}
	return nil, conversionError[Boolean, Quantity]()
}
func (b Boolean) Equal(other Element, _noReverseTypeConversion ...bool) bool {
	o, err := other.ToBoolean(false)
	if err == nil && o != nil {
		return b == *o
	}
	if len(_noReverseTypeConversion) > 0 && _noReverseTypeConversion[0] {
		return other.Equal(b, true)
	} else {
		return false
	}
}
func (b Boolean) Equivalent(other Element, _noReverseTypeConversion ...bool) bool {
	return b.Equal(other, _noReverseTypeConversion...)
}
func (b Boolean) TypeInfo() TypeInfo {
	return SimpleTypeInfo{
		Namespace: "System",
		Name:      "Boolean",
		BaseType:  TypeSpecifier{"System", "Any"},
	}
}
func (b Boolean) String() string {
	return strconv.FormatBool(bool(b))
}

type String string

func (s String) Children(name ...string) Collection {
	return nil
}

func (s String) ToBoolean(explicit bool) (*Boolean, error) {
	if explicit {
		if slices.Contains([]string{"true", "t", "yes", "y", "1", "1.0"}, string(s)) {
			return utils.Ptr[Boolean](true), nil
		} else if slices.Contains([]string{"false", "f", "no", "n", "0", "0.0"}, string(s)) {
			return utils.Ptr[Boolean](false), nil
		} else {
			return nil, nil
		}
	}
	return nil, implicitConversionError[String, Boolean]()
}
func (s String) ToString(explicit bool) (*String, error) {
	return &s, nil
}
func (s String) ToInteger(explicit bool) (*Integer, error) {
	if explicit {
		i, err := strconv.Atoi(string(s))
		if err != nil {
			return nil, nil
		}
		return utils.Ptr(Integer(i)), nil
	}
	return nil, conversionError[String, Integer]()
}
func (s String) ToDecimal(explicit bool) (*Decimal, error) {
	if explicit {
		d, _, err := apd.NewFromString(string(s))
		if err != nil {
			return nil, nil
		}
		return &Decimal{Value: d}, nil
	}
	return nil, conversionError[String, Decimal]()
}
func (s String) ToDate(explicit bool) (*Date, error) {
	if explicit {
		d, err := ParseDate(string(s))
		if err != nil {
			return nil, nil
		}
		return utils.Ptr[Date](d), nil
	}
	return nil, conversionError[String, Date]()
}
func (s String) ToTime(explicit bool) (*Time, error) {
	if explicit {
		d, err := parseTime(string(s), false)
		if err != nil {
			return nil, nil
		}
		return utils.Ptr[Time](d), nil
	}
	return nil, conversionError[String, Time]()
}
func (s String) ToDateTime(explicit bool) (*DateTime, error) {
	if explicit {
		d, err := ParseDateTime(string(s))
		if err != nil {
			return nil, nil
		}
		return utils.Ptr[DateTime](d), nil
	}
	return nil, conversionError[String, DateTime]()
}
func (s String) ToQuantity(explicit bool) (*Quantity, error) {
	q, err := ParseQuantity(string(s))
	if err != nil {
		return nil, nil
	}
	return &q, nil
}
func (s String) Equal(other Element, _noReverseTypeConversion ...bool) bool {
	o, err := other.ToString(false)
	if err == nil && o != nil {
		return s == *o
	}
	if len(_noReverseTypeConversion) > 0 && _noReverseTypeConversion[0] {
		return other.Equal(s, true)
	} else {
		return false
	}
}

var whitespaceReplaceRegex = regexp.MustCompile("[\t\r\n]")

func (s String) Equivalent(other Element, _noReverseTypeConversion ...bool) bool {
	o, err := other.ToString(false)
	if err == nil && o != nil {
		return whitespaceReplaceRegex.ReplaceAllString(strings.ToLower(string(s)), " ") ==
			whitespaceReplaceRegex.ReplaceAllString(strings.ToLower(string(*o)), " ")
	}
	if len(_noReverseTypeConversion) > 0 && _noReverseTypeConversion[0] {
		return other.Equal(s, true)
	} else {
		return false
	}
}
func (s String) Cmp(other Element) (*int, error) {
	o, err := other.ToString(false)
	if err != nil || o == nil {
		return nil, fmt.Errorf("can not compare String to %T, left: %v right: %v", other, s, other)
	}
	return utils.Ptr(strings.Compare(string(s), string(*o))), nil
}
func (s String) Add(ctx context.Context, other Element) (Element, error) {
	o, err := other.ToString(false)
	if err != nil {
		return nil, fmt.Errorf("can not add %T to String, %v + %v", other, s, other)
	}
	if o == nil {
		return nil, nil
	}
	return s + *o, nil
}
func (s String) TypeInfo() TypeInfo {
	return SimpleTypeInfo{
		Namespace: "System",
		Name:      "String",
		BaseType:  TypeSpecifier{"System", "Any"},
	}
}
func (s String) String() string {
	return string(s)
}

var (
	unescapeReplacer = strings.NewReplacer(
		`\'`, `'`,
		`\"`, `""`,
		`\'`, `'`,
		`\'`, `'`,
		`\'`, `'`,
	)
	escapeUnicodeRegex = regexp.MustCompile(`\\u([[:xdigit:]]){4}`)
)

func unescape(s string) (string, error) {
	unescaped := unescapeReplacer.Replace(s)

	var errs []error
	return escapeUnicodeRegex.ReplaceAllStringFunc(unescaped, func(s string) string {
		unquoted, err := strconv.Unquote(fmt.Sprintf("'%s'", s))
		if err != nil {
			errs = append(errs, err)
			return "�"
		}
		return unquoted
	}), errors.Join(errs...)
}

type Integer int32

func (i Integer) Children(name ...string) Collection {
	return nil
}

func (i Integer) ToBoolean(explicit bool) (*Boolean, error) {
	if explicit {
		switch i {
		case 0:
			return utils.Ptr[Boolean](false), nil
		case 1:
			return utils.Ptr[Boolean](true), nil
		default:
			return nil, nil
		}
	}
	return nil, conversionError[Integer, Boolean]()
}
func (i Integer) ToString(explicit bool) (*String, error) {
	return utils.Ptr(String(i.String())), nil
}
func (i Integer) ToInteger(explicit bool) (*Integer, error) {
	return &i, nil
}
func (i Integer) ToDecimal(explicit bool) (*Decimal, error) {
	return &Decimal{Value: apd.New(int64(i), 0)}, nil
}
func (i Integer) ToDate(explicit bool) (*Date, error) {
	return nil, conversionError[Integer, Date]()
}
func (i Integer) ToTime(explicit bool) (*Time, error) {
	return nil, conversionError[Integer, Time]()
}
func (i Integer) ToDateTime(explicit bool) (*DateTime, error) {
	return nil, conversionError[Integer, DateTime]()
}
func (i Integer) ToQuantity(explicit bool) (*Quantity, error) {
	return &Quantity{
		Value: Decimal{
			Value: apd.New(int64(i), 0),
		},
		Unit: "1",
	}, nil
}
func (i Integer) Equal(other Element, _noReverseTypeConversion ...bool) bool {
	o, err := other.ToInteger(false)
	if err == nil && o != nil {
		return i == *o
	}
	if len(_noReverseTypeConversion) > 0 && _noReverseTypeConversion[0] {
		return other.Equal(i, true)
	} else {
		return false
	}
}
func (i Integer) Equivalent(other Element, _noReverseTypeConversion ...bool) bool {
	return i.Equal(other, _noReverseTypeConversion...)
}
func (i Integer) Cmp(other cmpElement) (*int, error) {
	d, _ := i.ToDecimal(false)
	cmp, err := d.Cmp(other)
	if err != nil {
		return nil, fmt.Errorf("can not compare Integer to %T, left: %v right: %v", other, i, other)
	}
	return cmp, nil
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
		d, _ := i.ToDecimal(false)
		return d.Multiply(ctx, o)
	}
	return nil, fmt.Errorf("can not multiply Integer with %T: %v * %v", other, i, other)
}
func (i Integer) Divide(ctx context.Context, other Element) (Element, error) {
	d, err := i.ToDecimal(false)
	if err != nil {
		return nil, fmt.Errorf("can not divide Integer with %T: %v * %v", other, i, other)
	}
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
		d, _ := i.ToDecimal(false)
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
		d, _ := i.ToDecimal(false)
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
		d, _ := i.ToDecimal(false)
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
		d, _ := i.ToDecimal(false)
		return d.Subtract(ctx, o)
	}
	return nil, fmt.Errorf("can not subtract %T from Integer: %v - %v", other, i, other)
}
func (i Integer) TypeInfo() TypeInfo {
	return SimpleTypeInfo{
		Namespace: "System",
		Name:      "Integer",
		BaseType:  TypeSpecifier{"System", "Any"},
	}
}
func (i Integer) String() string {
	return strconv.Itoa(int(i))
}

type Decimal struct {
	Value *apd.Decimal
}

func (d Decimal) Children(name ...string) Collection {
	return nil
}

func (d Decimal) ToBoolean(explicit bool) (*Boolean, error) {
	if explicit {
		if d.Value.Cmp(apd.New(1, 0)) == 0 {
			return utils.Ptr[Boolean](true), nil
		} else if d.Value.Cmp(apd.New(0, 0)) == 0 {
			return utils.Ptr[Boolean](false), nil
		} else {
			return nil, nil
		}
	}
	return nil, conversionError[Decimal, Boolean]()
}
func (d Decimal) ToString(explicit bool) (*String, error) {
	return utils.Ptr(String(d.String())), nil
}
func (d Decimal) ToInteger(explicit bool) (*Integer, error) {
	return nil, conversionError[Decimal, Integer]()
}
func (d Decimal) ToDecimal(explicit bool) (*Decimal, error) {
	return &d, nil
}
func (d Decimal) ToDate(explicit bool) (*Date, error) {
	return nil, conversionError[Decimal, Date]()
}
func (d Decimal) ToTime(explicit bool) (*Time, error) {
	return nil, conversionError[Decimal, Time]()
}
func (d Decimal) ToDateTime(explicit bool) (*DateTime, error) {
	return nil, conversionError[Decimal, DateTime]()
}
func (d Decimal) ToQuantity(explicit bool) (*Quantity, error) {
	return &Quantity{
		Value: d,
		Unit:  "1",
	}, nil
}
func (d Decimal) Equal(other Element, _noReverseTypeConversion ...bool) bool {
	o, err := other.ToDecimal(false)
	if err == nil && o != nil {
		return d == *o
	}
	if len(_noReverseTypeConversion) > 0 && _noReverseTypeConversion[0] {
		return other.Equal(d, true)
	} else {
		return false
	}
}
func (d Decimal) Equivalent(other Element, _noReverseTypeConversion ...bool) bool {
	// TODO: values must be equal, comparison is done on values rounded to the precision of the least precise operand.
	//       Trailing zeroes after the decimal are ignored in determining precision.
	return d.Equal(other, _noReverseTypeConversion...)
}
func (d Decimal) Cmp(other Element) (*int, error) {
	o, err := other.ToDecimal(false)
	if err != nil || o == nil {
		return nil, fmt.Errorf("can not compare Decimal to %T, left: %v right: %v", other, d, other)
	}
	return utils.Ptr(d.Value.Cmp(o.Value)), nil
}
func (d Decimal) Multiply(ctx context.Context, other Element) (Element, error) {
	o, err := other.ToDecimal(false)
	if err != nil || o == nil {
		return nil, fmt.Errorf("can not multiply Decimal with %T: %v * %v", other, d, other)
	}
	var res *apd.Decimal
	_, err = apdContext(ctx).Mul(res, d.Value, o.Value)
	if err != nil {
		return nil, err
	}
	return Decimal{res}, nil
}
func (d Decimal) Divide(ctx context.Context, other Element) (Element, error) {
	o, err := other.ToDecimal(false)
	if err != nil || o == nil {
		return nil, fmt.Errorf("can not divide Decimal with %T: %v / %v", other, d, other)
	}
	if o.Value.IsZero() {
		return nil, nil
	}
	var res *apd.Decimal
	_, err = apdContext(ctx).Quo(res, d.Value, o.Value)
	if err != nil {
		return nil, err
	}
	return Decimal{res}, nil
}
func (d Decimal) Div(ctx context.Context, other Element) (Element, error) {
	o, err := other.ToDecimal(false)
	if err != nil || o == nil {
		return nil, fmt.Errorf("can not div Decimal with %T: %v div %v", other, d, other)
	}
	if o.Value.IsZero() {
		return nil, nil
	}
	var res *apd.Decimal
	_, err = apdContext(ctx).QuoInteger(res, d.Value, o.Value)
	if err != nil {
		return nil, err
	}
	return Decimal{res}, nil
}
func (d Decimal) Mod(ctx context.Context, other Element) (Element, error) {
	o, err := other.ToDecimal(false)
	if err != nil || o == nil {
		return nil, fmt.Errorf("can not mod Decimal with %T: %v mod %v", other, d, other)
	}
	if o.Value.IsZero() {
		return nil, nil
	}
	var res *apd.Decimal
	_, err = apdContext(ctx).Rem(res, d.Value, o.Value)
	if err != nil {
		return nil, err
	}
	return Decimal{res}, nil
}
func (d Decimal) Add(ctx context.Context, other Element) (Element, error) {
	o, err := other.ToDecimal(false)
	if err != nil || o == nil {
		return nil, fmt.Errorf("can not add Decimal and %T: %v + %v", other, d, other)
	}
	if o.Value.IsZero() {
		return nil, nil
	}
	var res *apd.Decimal
	_, err = apdContext(ctx).Add(res, d.Value, o.Value)
	if err != nil {
		return nil, err
	}
	return Decimal{res}, nil
}
func (d Decimal) Subtract(ctx context.Context, other Element) (Element, error) {
	o, err := other.ToDecimal(false)
	if err != nil || o == nil {
		return nil, fmt.Errorf("can not subtract %T from Decimal: %v - %v", other, d, other)
	}
	if o.Value.IsZero() {
		return nil, nil
	}
	var res *apd.Decimal
	_, err = apdContext(ctx).Sub(res, d.Value, o.Value)
	if err != nil {
		return nil, err
	}
	return Decimal{res}, nil
}
func (d Decimal) TypeInfo() TypeInfo {
	return SimpleTypeInfo{
		Namespace: "System",
		Name:      "Decimal",
		BaseType:  TypeSpecifier{"System", "Any"},
	}
}
func (d Decimal) String() string {
	return d.Value.Text('f')
}

type Date struct {
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

func (d Date) ToBoolean(explicit bool) (*Boolean, error) {
	return nil, conversionError[Date, Boolean]()
}
func (d Date) ToString(explicit bool) (*String, error) {
	return utils.Ptr(String(d.String())), nil
}
func (d Date) ToInteger(explicit bool) (*Integer, error) {
	return nil, conversionError[Date, Integer]()
}
func (d Date) ToDecimal(explicit bool) (*Decimal, error) {
	return nil, conversionError[Date, Decimal]()
}
func (d Date) ToDate(explicit bool) (*Date, error) {
	return &d, nil
}
func (d Date) ToTime(explicit bool) (*Time, error) {
	return nil, conversionError[Date, Time]()
}
func (d Date) ToDateTime(explicit bool) (*DateTime, error) {
	return &DateTime{
		Value:     d.Value,
		Precision: DateTimePrecisionDay,
	}, nil
}
func (d Date) ToQuantity(explicit bool) (*Quantity, error) {
	return nil, conversionError[Date, Quantity]()
}
func (d Date) Equal(other Element, _noReverseTypeConversion ...bool) bool {
	o, err := other.ToDate(false)
	if err == nil && o != nil {
		cmp, err := d.Cmp(o)
		if err == nil && cmp != nil {
			return *cmp == 0
		}
	}
	if len(_noReverseTypeConversion) > 0 && _noReverseTypeConversion[0] {
		return other.Equal(d, true)
	} else {
		return false
	}
}
func (d Date) Equivalent(other Element, _noReverseTypeConversion ...bool) bool {
	o, err := other.ToDate(false)
	if err == nil && o != nil && d.Precision == o.Precision {
		return d.Equal(other, _noReverseTypeConversion...)
	}
	return false
}
func (d Date) Cmp(other Element) (*int, error) {
	o, err := other.ToDate(false)
	if err != nil || o == nil {
		return nil, fmt.Errorf("can not compare Date to %T, left: %v right: %v", other, d, other)
	}
	switch d.Precision {
	case DatePrecisionYear:
		if o.Precision != DatePrecisionYear {
			return nil, nil
		}
		if d.Value.Year() < o.Value.In(d.Value.Location()).Year() {
			return utils.Ptr(-1), nil
		} else if d.Value.Year() > o.Value.In(d.Value.Location()).Year() {
			return utils.Ptr(1), nil
		} else {
			return utils.Ptr(0), nil
		}
	case DatePrecisionMonth:
		if o.Precision != DatePrecisionMonth {
			return nil, nil
		}
		if d.Value.Year() < o.Value.In(d.Value.Location()).Year() {
			return utils.Ptr(-1), nil
		} else if d.Value.Year() > o.Value.In(d.Value.Location()).Year() {
			return utils.Ptr(1), nil
		} else if d.Value.Month() < o.Value.In(d.Value.Location()).Month() {
			return utils.Ptr(-1), nil
		} else if d.Value.Month() > o.Value.In(d.Value.Location()).Month() {
			return utils.Ptr(1), nil
		} else {
			return utils.Ptr(0), nil
		}
	default:
		if d.Value.Year() < o.Value.In(d.Value.Location()).Year() {
			return utils.Ptr(-1), nil
		} else if d.Value.Year() > o.Value.In(d.Value.Location()).Year() {
			return utils.Ptr(1), nil
		} else if d.Value.Month() < o.Value.In(d.Value.Location()).Month() {
			return utils.Ptr(-1), nil
		} else if d.Value.Month() > o.Value.In(d.Value.Location()).Month() {
			return utils.Ptr(1), nil
		} else if d.Value.Day() < o.Value.In(d.Value.Location()).Day() {
			return utils.Ptr(-1), nil
		} else if d.Value.Day() > o.Value.In(d.Value.Location()).Day() {
			return utils.Ptr(1), nil
		} else {
			return utils.Ptr(0), nil
		}
	}
}
func (d Date) TypeInfo() TypeInfo {
	return SimpleTypeInfo{
		Namespace: "System",
		Name:      "Date",
		BaseType:  TypeSpecifier{"System", "Any"},
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
	return fmt.Sprintf("@%s", ds)
}

type Time struct {
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

func (t Time) ToBoolean(explicit bool) (*Boolean, error) {
	return nil, conversionError[Time, Boolean]()
}
func (t Time) ToString(explicit bool) (*String, error) {
	return utils.Ptr(String(t.String())), nil
}
func (t Time) ToInteger(explicit bool) (*Integer, error) {
	return nil, conversionError[Time, Integer]()
}
func (t Time) ToDecimal(explicit bool) (*Decimal, error) {
	return nil, conversionError[Time, Decimal]()
}
func (t Time) ToDate(explicit bool) (*Date, error) {
	return nil, conversionError[Time, Date]()
}
func (t Time) ToTime(explicit bool) (*Time, error) {
	return &t, nil
}
func (t Time) ToDateTime(explicit bool) (*DateTime, error) {
	return nil, conversionError[Time, DateTime]()
}
func (t Time) ToQuantity(explicit bool) (*Quantity, error) {
	return nil, conversionError[Time, Quantity]()
}
func (t Time) Equal(other Element, _noReverseTypeConversion ...bool) bool {
	o, err := other.ToTime(false)
	if err == nil && o != nil {
		cmp, err := t.Cmp(o)
		if err == nil && cmp != nil {
			return *cmp == 0
		}
	}
	if len(_noReverseTypeConversion) > 0 && _noReverseTypeConversion[0] {
		return other.Equal(t, true)
	} else {
		return false
	}
}
func (t Time) Equivalent(other Element, _noReverseTypeConversion ...bool) bool {
	o, err := other.ToTime(false)
	if err == nil && o != nil && t.Precision == o.Precision {
		return t.Equal(other, _noReverseTypeConversion...)
	}
	return false
}
func (t Time) Cmp(other Element) (*int, error) {
	o, err := other.ToTime(false)
	if err != nil || o == nil {
		return nil, fmt.Errorf("can not compare Time to %T, left: %v right: %v", other, t, other)
	}
	switch t.Precision {
	case TimePrecisionHour:
		if o.Precision != TimePrecisionHour {
			return nil, nil
		}
		return utils.Ptr(t.Value.Truncate(time.Hour).Compare(
			o.Value.In(t.Value.Location()).Truncate(time.Hour))), nil
	case TimePrecisionMinute:
		if o.Precision != TimePrecisionMinute {
			return nil, nil
		}
		return utils.Ptr(t.Value.Truncate(time.Minute).Compare(
			o.Value.In(t.Value.Location()).Truncate(time.Minute))), nil
	default:
		return utils.Ptr(t.Value.Compare(o.Value.In(t.Value.Location()))), nil
	}
}
func (t Time) TypeInfo() TypeInfo {
	return SimpleTypeInfo{
		Namespace: "System",
		Name:      "Time",
		BaseType:  TypeSpecifier{"System", "Any"},
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
	return fmt.Sprintf("@T%s", ts)
}

type DateTime struct {
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

func (dt DateTime) ToBoolean(explicit bool) (*Boolean, error) {
	return nil, conversionError[DateTime, Boolean]()
}
func (dt DateTime) ToString(explicit bool) (*String, error) {
	return utils.Ptr(String(dt.String())), nil
}
func (dt DateTime) ToInteger(explicit bool) (*Integer, error) {
	return nil, conversionError[DateTime, Integer]()
}
func (dt DateTime) ToDecimal(explicit bool) (*Decimal, error) {
	return nil, conversionError[DateTime, Decimal]()
}
func (dt DateTime) ToDate(explicit bool) (*Date, error) {
	var precision DatePrecision
	switch dt.Precision {
	case DateTimePrecisionYear, DateTimePrecisionMonth:
		precision = DatePrecision(dt.Precision)
	default:
		precision = DatePrecisionFull
	}
	return &Date{
		Value:     dt.Value,
		Precision: precision,
	}, nil
}
func (dt DateTime) ToTime(explicit bool) (*Time, error) {
	return nil, conversionError[DateTime, Time]()
}
func (dt DateTime) ToDateTime(explicit bool) (*DateTime, error) {
	return &dt, nil
}
func (dt DateTime) ToQuantity(explicit bool) (*Quantity, error) {
	return nil, conversionError[DateTime, Quantity]()
}
func (dt DateTime) Equal(other Element, _noReverseTypeConversion ...bool) bool {
	o, err := other.ToDateTime(false)
	if err == nil && o != nil {
		cmp, err := dt.Cmp(o)
		if err == nil && cmp != nil {
			return *cmp == 0
		}
	}
	if len(_noReverseTypeConversion) > 0 && _noReverseTypeConversion[0] {
		return other.Equal(dt, true)
	} else {
		return false
	}
}
func (dt DateTime) Equivalent(other Element, _noReverseTypeConversion ...bool) bool {
	o, err := other.ToDateTime(false)
	if err == nil && o != nil && dt.Precision == o.Precision {
		return dt.Equal(other, _noReverseTypeConversion...)
	}
	return false
}
func (dt DateTime) Cmp(other Element) (*int, error) {
	o, err := other.ToTime(false)
	if err != nil || o == nil {
		return nil, fmt.Errorf("can not compare Time to %T, left: %v right: %v", other, dt, other)
	}
	switch dt.Precision {
	case DateTimePrecisionYear:
		if o.Precision != DateTimePrecisionYear {
			return nil, nil
		}
		if dt.Value.Year() < o.Value.In(dt.Value.Location()).Year() {
			return utils.Ptr(-1), nil
		} else if dt.Value.Year() > o.Value.In(dt.Value.Location()).Year() {
			return utils.Ptr(1), nil
		} else {
			return utils.Ptr(0), nil
		}
	case DateTimePrecisionMonth:
		if o.Precision != DateTimePrecisionMonth {
			return nil, nil
		}
		if dt.Value.Year() < o.Value.In(dt.Value.Location()).Year() {
			return utils.Ptr(-1), nil
		} else if dt.Value.Year() > o.Value.In(dt.Value.Location()).Year() {
			return utils.Ptr(1), nil
		} else if dt.Value.Month() < o.Value.In(dt.Value.Location()).Month() {
			return utils.Ptr(-1), nil
		} else if dt.Value.Month() > o.Value.In(dt.Value.Location()).Month() {
			return utils.Ptr(1), nil
		} else {
			return utils.Ptr(0), nil
		}
	case DateTimePrecisionDay:
		if o.Precision != DateTimePrecisionDay {
			return nil, nil
		}
		if dt.Value.Year() < o.Value.In(dt.Value.Location()).Year() {
			return utils.Ptr(-1), nil
		} else if dt.Value.Year() > o.Value.In(dt.Value.Location()).Year() {
			return utils.Ptr(1), nil
		} else if dt.Value.Month() < o.Value.In(dt.Value.Location()).Month() {
			return utils.Ptr(-1), nil
		} else if dt.Value.Month() > o.Value.In(dt.Value.Location()).Month() {
			return utils.Ptr(1), nil
		} else if dt.Value.Day() < o.Value.In(dt.Value.Location()).Day() {
			return utils.Ptr(-1), nil
		} else if dt.Value.Day() > o.Value.In(dt.Value.Location()).Day() {
			return utils.Ptr(1), nil
		} else {
			return utils.Ptr(0), nil
		}
	case DateTimePrecisionHour:
		if o.Precision != DateTimePrecisionHour {
			return nil, nil
		}
		return utils.Ptr(dt.Value.Truncate(time.Hour).Compare(
			o.Value.In(dt.Value.Location()).Truncate(time.Hour))), nil
	case DateTimePrecisionMinute:
		if o.Precision != DateTimePrecisionMinute {
			return nil, nil
		}
		return utils.Ptr(dt.Value.Truncate(time.Minute).Compare(
			o.Value.In(dt.Value.Location()).Truncate(time.Minute))), nil
	default:
		return utils.Ptr(dt.Value.Compare(
			o.Value.In(dt.Value.Location()))), nil
	}
}
func (dt DateTime) TypeInfo() TypeInfo {
	return SimpleTypeInfo{
		Namespace: "System",
		Name:      "DateTime",
		BaseType:  TypeSpecifier{"System", "Any"},
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

	if ts == "" {
		return fmt.Sprintf("@%sT", ds)
	}

	return fmt.Sprintf("@%sT%s", ds, ts)
}

const (
	DateFormatOnlyYear     = "2006"
	DateFormatUpToMonth    = "2006-01"
	DateFormatFull         = "2006-01-02"
	TimeFormatOnlyHour     = "15"
	TimeFormatOnlyHourTZ   = "15Z07:00"
	TimeFormatUpToMinute   = "15:04Z07:00"
	TimeFormatUpToMinuteTZ = "15:04Z07:00"
	TimeFormatFull         = "15:04:05.999999999"
	TimeFormatFullTZ       = "15:04:05.999999999Z07:00"
)

func ParseDate(s string) (Date, error) {
	ds := strings.TrimLeft(s, "@")

	d, err := time.Parse(DateFormatOnlyYear, ds)
	if err == nil {
		return Date{d, DatePrecisionYear}, nil
	}
	d, err = time.Parse(DateFormatUpToMonth, ds)
	if err == nil {
		return Date{d, DatePrecisionMonth}, nil
	}
	d, err = time.Parse(DateFormatFull, ds)
	if err == nil {
		return Date{d, DatePrecisionFull}, nil
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
		return Time{t, TimePrecisionHour}, nil
	}
	if withTZ {
		t, err = time.Parse(TimeFormatOnlyHourTZ, ts)
		if err == nil {
			return Time{t, TimePrecisionHour}, nil
		}
	}
	t, err = time.Parse(TimeFormatUpToMinute, ts)
	if err == nil {
		return Time{t, TimePrecisionMinute}, nil
	}
	if withTZ {
		t, err = time.Parse(TimeFormatUpToMinuteTZ, ts)
		if err == nil {
			return Time{t, TimePrecisionMinute}, nil
		}
	}
	t, err = time.Parse(TimeFormatFull, ts)
	if err == nil {
		return Time{t, TimePrecisionFull}, nil
	}
	if withTZ {
		t, err = time.Parse(TimeFormatFullTZ, ts)
		if err == nil {
			return Time{t, TimePrecisionFull}, nil
		}
	}

	return Time{}, fmt.Errorf("invalid Date format: %s", s)
}

func ParseDateTime(s string) (DateTime, error) {
	splits := strings.Split(s, "T")
	ds := splits[0]
	ts := splits[1]

	d, err := ParseDate(ds)
	if err != nil {
		return DateTime{}, fmt.Errorf("invalid DateTime format (date part): %s", s)
	}

	if len(ts) == 0 {
		if d.Precision == DateTimePrecisionFull {
			return DateTime{d.Value, DateTimePrecisionDay}, nil
		}
		return DateTime{d.Value, DateTimePrecision(d.Precision)}, nil
	}

	t, err := parseTime(ts, true)
	if err != nil {
		return DateTime{}, fmt.Errorf("invalid DateTime format (time part): %s", s)
	}

	dt := d.Value.Add(
		time.Hour*time.Duration(t.Value.Hour()) +
			time.Minute*time.Duration(t.Value.Minute()) +
			time.Second*time.Duration(t.Value.Second()) +
			time.Nanosecond*time.Duration(t.Value.Nanosecond()),
	)
	return DateTime{dt, DateTimePrecision(t.Precision)}, nil
}

type Quantity struct {
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

func (q Quantity) ToBoolean(explicit bool) (*Boolean, error) {
	return nil, conversionError[Quantity, Boolean]()
}
func (q Quantity) ToString(explicit bool) (*String, error) {
	return utils.Ptr(String(q.String())), nil
}
func (q Quantity) ToInteger(explicit bool) (*Integer, error) {
	return nil, conversionError[Quantity, Integer]()
}
func (q Quantity) ToDecimal(explicit bool) (*Decimal, error) {
	return nil, conversionError[Quantity, Decimal]()
}
func (q Quantity) ToDate(explicit bool) (*Date, error) {
	return nil, conversionError[Quantity, Date]()
}
func (q Quantity) ToTime(explicit bool) (*Time, error) {
	return nil, conversionError[Quantity, Time]()
}
func (q Quantity) ToDateTime(explicit bool) (*DateTime, error) {
	return nil, conversionError[Quantity, DateTime]()
}
func (q Quantity) ToQuantity(explicit bool) (*Quantity, error) {
	return &q, nil
}
func (q Quantity) Equal(other Element, _noReverseTypeConversion ...bool) bool {
	o, err := other.ToQuantity(false)
	if err == nil && o != nil {
		return q == *o
	}
	if len(_noReverseTypeConversion) > 0 && _noReverseTypeConversion[0] {
		return other.Equal(q, true)
	} else {
		return false
	}
}
func (q Quantity) Equivalent(other Element, _noReverseTypeConversion ...bool) bool {
	o, err := other.ToQuantity(false)
	if err == nil && o != nil {
		return q.dateAsUCUM().Equal(o.dateAsUCUM(), true)
	}
	return false
}
func (q Quantity) Cmp(other Element) (*int, error) {
	o, err := other.ToQuantity(false)
	if err != nil || o == nil {
		return utils.Ptr(0), fmt.Errorf("can not compare Quantity to %T, left: %v right: %v", other, q, other)
	}
	left := q.dateAsUCUM()
	right := o.dateAsUCUM()
	if left.Unit != right.Unit {
		return utils.Ptr(0), fmt.Errorf("quantity units do not match, left: %v right: %v", left.Unit, right.Unit)
	}
	return left.Value.Cmp(right.Value)
}
func (q Quantity) Multiply(ctx context.Context, other Element) (Element, error) {
	o, err := other.ToQuantity(false)
	if err != nil || o == nil {
		return nil, fmt.Errorf("can not multiply Quantity with %T: %v * %v", other, q, other)
	}
	left := q.dateAsUCUM()
	right := o.dateAsUCUM()

	if left.Unit != right.Unit {
		return Quantity{}, fmt.Errorf("quantity units do not match, left: %v right: %v", left, right)
	}

	value, err := left.Value.Multiply(ctx, right.Value)
	if err != nil {
		return Quantity{}, err
	}
	unit := fmt.Sprintf("(%s).(%s)", left.Unit, right.Unit)
	return Quantity{Value: value.(Decimal), Unit: String(unit)}, nil
}
func (q Quantity) Divide(ctx context.Context, other Element) (Element, error) {
	o, err := other.ToQuantity(false)
	if err != nil || o == nil {
		return nil, fmt.Errorf("can not divide Quantity with %T: %v / %v", other, q, other)
	}
	left := q.dateAsUCUM()
	right := o.dateAsUCUM()

	if left.Unit != right.Unit {
		return Quantity{}, fmt.Errorf("quantity units do not match, left: %v right: %v", left, right)
	}

	value, err := left.Value.Divide(ctx, right.Value)
	if err != nil {
		return Quantity{}, err
	}
	unit := fmt.Sprintf("(%s).(%s)", left.Unit, right.Unit)
	return Quantity{Value: value.(Decimal), Unit: String(unit)}, nil
}
func (q Quantity) Add(ctx context.Context, other Element) (Element, error) {
	o, err := other.ToQuantity(false)
	if err != nil || o == nil {
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
	o, err := other.ToQuantity(false)
	if err != nil || o == nil {
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
			BaseType:  TypeSpecifier{"System", "Any"},
		},
		Element: []ClassInfoElement{
			{Name: "Value", Type: "System.Decimal"},
			{Name: "Unit", Type: "System.String"},
		},
	}
}
func (q Quantity) String() string {
	return fmt.Sprintf("%s '%s'", q.Value.String(), q.Unit)
}

func ParseQuantity(s string) (Quantity, error) {
	parser, err := parse(s)
	if err != nil {
		return Quantity{}, err
	}

	q := parser.Quantity()
	v, _, err := apd.NewFromString(q.NUMBER().GetText())
	u := q.Unit().GetText()
	if err != nil {
		return Quantity{}, err
	}

	return Quantity{Value: Decimal{v}, Unit: String(u)}, nil
}

func conversionError[F Element, T Element]() error {
	var (
		f F
		t T
	)
	return fmt.Errorf("primitive %v of type %T can not be converted to type %T", f, f, t)
}

func implicitConversionError[F Element, T Element]() error {
	var (
		f F
		t T
	)
	return fmt.Errorf("primitive %v of type %T can not be implicitly converted to type %T", f, f, t)
}
