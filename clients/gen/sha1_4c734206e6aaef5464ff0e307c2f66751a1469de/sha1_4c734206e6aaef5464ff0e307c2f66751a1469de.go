// This file was generated by nomdl/codegen.

package sha1_4c734206e6aaef5464ff0e307c2f66751a1469de

import (
	"github.com/attic-labs/noms/clients/gen/sha1_52bbaa7c5bcb39759981ccb12ee457f21fa7517d"
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

var __sha1_4c734206e6aaef5464ff0e307c2f66751a1469dePackageInFile_sha1_4c734206e6aaef5464ff0e307c2f66751a1469de_CachedRef = __sha1_4c734206e6aaef5464ff0e307c2f66751a1469dePackageInFile_sha1_4c734206e6aaef5464ff0e307c2f66751a1469de_Ref()

// This function builds up a Noms value that describes the type
// package implemented by this file and registers it with the global
// type package definition cache.
func __sha1_4c734206e6aaef5464ff0e307c2f66751a1469dePackageInFile_sha1_4c734206e6aaef5464ff0e307c2f66751a1469de_Ref() ref.Ref {
	p := types.NewPackage([]types.TypeRef{

		types.MakeStructTypeRef("RemotePhoto",
			[]types.Field{
				types.Field{"Id", types.MakePrimitiveTypeRef(types.StringKind), false},
				types.Field{"Title", types.MakePrimitiveTypeRef(types.StringKind), false},
				types.Field{"Url", types.MakePrimitiveTypeRef(types.StringKind), false},
				types.Field{"Geoposition", types.MakeTypeRef(ref.Parse("sha1-52bbaa7c5bcb39759981ccb12ee457f21fa7517d"), 0), false},
				types.Field{"Sizes", types.MakeCompoundTypeRef("", types.MapKind, types.MakeTypeRef(ref.Ref{}, 1), types.MakePrimitiveTypeRef(types.StringKind)), false},
				types.Field{"Tags", types.MakeCompoundTypeRef("", types.SetKind, types.MakePrimitiveTypeRef(types.StringKind)), false},
			},
			types.Choices{},
		),
		types.MakeStructTypeRef("Size",
			[]types.Field{
				types.Field{"Width", types.MakePrimitiveTypeRef(types.UInt32Kind), false},
				types.Field{"Height", types.MakePrimitiveTypeRef(types.UInt32Kind), false},
			},
			types.Choices{},
		),
	}, []ref.Ref{})
	return types.RegisterPackage(&p)
}

// RemotePhoto

type RemotePhoto struct {
	m types.Map
}

func NewRemotePhoto() RemotePhoto {
	return RemotePhoto{types.NewMap(
		types.NewString("$type"), types.MakeTypeRef(__sha1_4c734206e6aaef5464ff0e307c2f66751a1469dePackageInFile_sha1_4c734206e6aaef5464ff0e307c2f66751a1469de_CachedRef, 0),
		types.NewString("Id"), types.NewString(""),
		types.NewString("Title"), types.NewString(""),
		types.NewString("Url"), types.NewString(""),
		types.NewString("Geoposition"), sha1_52bbaa7c5bcb39759981ccb12ee457f21fa7517d.NewGeoposition().NomsValue(),
		types.NewString("Sizes"), types.NewMap(),
		types.NewString("Tags"), types.NewSet(),
	)}
}

type RemotePhotoDef struct {
	Id          string
	Title       string
	Url         string
	Geoposition sha1_52bbaa7c5bcb39759981ccb12ee457f21fa7517d.GeopositionDef
	Sizes       MapOfSizeToStringDef
	Tags        SetOfStringDef
}

func (def RemotePhotoDef) New() RemotePhoto {
	return RemotePhoto{
		types.NewMap(
			types.NewString("$type"), types.MakeTypeRef(__sha1_4c734206e6aaef5464ff0e307c2f66751a1469dePackageInFile_sha1_4c734206e6aaef5464ff0e307c2f66751a1469de_CachedRef, 0),
			types.NewString("Id"), types.NewString(def.Id),
			types.NewString("Title"), types.NewString(def.Title),
			types.NewString("Url"), types.NewString(def.Url),
			types.NewString("Geoposition"), def.Geoposition.New().NomsValue(),
			types.NewString("Sizes"), def.Sizes.New().NomsValue(),
			types.NewString("Tags"), def.Tags.New().NomsValue(),
		)}
}

func (s RemotePhoto) Def() (d RemotePhotoDef) {
	d.Id = s.m.Get(types.NewString("Id")).(types.String).String()
	d.Title = s.m.Get(types.NewString("Title")).(types.String).String()
	d.Url = s.m.Get(types.NewString("Url")).(types.String).String()
	d.Geoposition = sha1_52bbaa7c5bcb39759981ccb12ee457f21fa7517d.GeopositionFromVal(s.m.Get(types.NewString("Geoposition"))).Def()
	d.Sizes = MapOfSizeToStringFromVal(s.m.Get(types.NewString("Sizes"))).Def()
	d.Tags = SetOfStringFromVal(s.m.Get(types.NewString("Tags"))).Def()
	return
}

var __typeRefForRemotePhoto = types.MakeTypeRef(__sha1_4c734206e6aaef5464ff0e307c2f66751a1469dePackageInFile_sha1_4c734206e6aaef5464ff0e307c2f66751a1469de_CachedRef, 0)

func (m RemotePhoto) TypeRef() types.TypeRef {
	return __typeRefForRemotePhoto
}

func init() {
	types.RegisterFromValFunction(__typeRefForRemotePhoto, func(v types.Value) types.NomsValue {
		return RemotePhotoFromVal(v)
	})
}

func RemotePhotoFromVal(val types.Value) RemotePhoto {
	// TODO: Validate here
	return RemotePhoto{val.(types.Map)}
}

func (s RemotePhoto) NomsValue() types.Value {
	return s.m
}

func (s RemotePhoto) Equals(other types.Value) bool {
	if other, ok := other.(RemotePhoto); ok {
		return s.m.Equals(other.m)
	}
	return false
}

func (s RemotePhoto) Ref() ref.Ref {
	return s.m.Ref()
}

func (s RemotePhoto) Chunks() (futures []types.Future) {
	futures = append(futures, s.TypeRef().Chunks()...)
	futures = append(futures, s.m.Chunks()...)
	return
}

func (s RemotePhoto) Id() string {
	return s.m.Get(types.NewString("Id")).(types.String).String()
}

func (s RemotePhoto) SetId(val string) RemotePhoto {
	return RemotePhoto{s.m.Set(types.NewString("Id"), types.NewString(val))}
}

func (s RemotePhoto) Title() string {
	return s.m.Get(types.NewString("Title")).(types.String).String()
}

func (s RemotePhoto) SetTitle(val string) RemotePhoto {
	return RemotePhoto{s.m.Set(types.NewString("Title"), types.NewString(val))}
}

func (s RemotePhoto) Url() string {
	return s.m.Get(types.NewString("Url")).(types.String).String()
}

func (s RemotePhoto) SetUrl(val string) RemotePhoto {
	return RemotePhoto{s.m.Set(types.NewString("Url"), types.NewString(val))}
}

func (s RemotePhoto) Geoposition() sha1_52bbaa7c5bcb39759981ccb12ee457f21fa7517d.Geoposition {
	return sha1_52bbaa7c5bcb39759981ccb12ee457f21fa7517d.GeopositionFromVal(s.m.Get(types.NewString("Geoposition")))
}

func (s RemotePhoto) SetGeoposition(val sha1_52bbaa7c5bcb39759981ccb12ee457f21fa7517d.Geoposition) RemotePhoto {
	return RemotePhoto{s.m.Set(types.NewString("Geoposition"), val.NomsValue())}
}

func (s RemotePhoto) Sizes() MapOfSizeToString {
	return MapOfSizeToStringFromVal(s.m.Get(types.NewString("Sizes")))
}

func (s RemotePhoto) SetSizes(val MapOfSizeToString) RemotePhoto {
	return RemotePhoto{s.m.Set(types.NewString("Sizes"), val.NomsValue())}
}

func (s RemotePhoto) Tags() SetOfString {
	return SetOfStringFromVal(s.m.Get(types.NewString("Tags")))
}

func (s RemotePhoto) SetTags(val SetOfString) RemotePhoto {
	return RemotePhoto{s.m.Set(types.NewString("Tags"), val.NomsValue())}
}

// Size

type Size struct {
	m types.Map
}

func NewSize() Size {
	return Size{types.NewMap(
		types.NewString("$type"), types.MakeTypeRef(__sha1_4c734206e6aaef5464ff0e307c2f66751a1469dePackageInFile_sha1_4c734206e6aaef5464ff0e307c2f66751a1469de_CachedRef, 1),
		types.NewString("Width"), types.UInt32(0),
		types.NewString("Height"), types.UInt32(0),
	)}
}

type SizeDef struct {
	Width  uint32
	Height uint32
}

func (def SizeDef) New() Size {
	return Size{
		types.NewMap(
			types.NewString("$type"), types.MakeTypeRef(__sha1_4c734206e6aaef5464ff0e307c2f66751a1469dePackageInFile_sha1_4c734206e6aaef5464ff0e307c2f66751a1469de_CachedRef, 1),
			types.NewString("Width"), types.UInt32(def.Width),
			types.NewString("Height"), types.UInt32(def.Height),
		)}
}

func (s Size) Def() (d SizeDef) {
	d.Width = uint32(s.m.Get(types.NewString("Width")).(types.UInt32))
	d.Height = uint32(s.m.Get(types.NewString("Height")).(types.UInt32))
	return
}

var __typeRefForSize = types.MakeTypeRef(__sha1_4c734206e6aaef5464ff0e307c2f66751a1469dePackageInFile_sha1_4c734206e6aaef5464ff0e307c2f66751a1469de_CachedRef, 1)

func (m Size) TypeRef() types.TypeRef {
	return __typeRefForSize
}

func init() {
	types.RegisterFromValFunction(__typeRefForSize, func(v types.Value) types.NomsValue {
		return SizeFromVal(v)
	})
}

func SizeFromVal(val types.Value) Size {
	// TODO: Validate here
	return Size{val.(types.Map)}
}

func (s Size) NomsValue() types.Value {
	return s.m
}

func (s Size) Equals(other types.Value) bool {
	if other, ok := other.(Size); ok {
		return s.m.Equals(other.m)
	}
	return false
}

func (s Size) Ref() ref.Ref {
	return s.m.Ref()
}

func (s Size) Chunks() (futures []types.Future) {
	futures = append(futures, s.TypeRef().Chunks()...)
	futures = append(futures, s.m.Chunks()...)
	return
}

func (s Size) Width() uint32 {
	return uint32(s.m.Get(types.NewString("Width")).(types.UInt32))
}

func (s Size) SetWidth(val uint32) Size {
	return Size{s.m.Set(types.NewString("Width"), types.UInt32(val))}
}

func (s Size) Height() uint32 {
	return uint32(s.m.Get(types.NewString("Height")).(types.UInt32))
}

func (s Size) SetHeight(val uint32) Size {
	return Size{s.m.Set(types.NewString("Height"), types.UInt32(val))}
}

// MapOfSizeToString

type MapOfSizeToString struct {
	m types.Map
}

func NewMapOfSizeToString() MapOfSizeToString {
	return MapOfSizeToString{types.NewMap()}
}

type MapOfSizeToStringDef map[SizeDef]string

func (def MapOfSizeToStringDef) New() MapOfSizeToString {
	kv := make([]types.Value, 0, len(def)*2)
	for k, v := range def {
		kv = append(kv, k.New().NomsValue(), types.NewString(v))
	}
	return MapOfSizeToString{types.NewMap(kv...)}
}

func (m MapOfSizeToString) Def() MapOfSizeToStringDef {
	def := make(map[SizeDef]string)
	m.m.Iter(func(k, v types.Value) bool {
		def[SizeFromVal(k).Def()] = v.(types.String).String()
		return false
	})
	return def
}

func MapOfSizeToStringFromVal(p types.Value) MapOfSizeToString {
	// TODO: Validate here
	return MapOfSizeToString{p.(types.Map)}
}

func (m MapOfSizeToString) NomsValue() types.Value {
	return m.m
}

func (m MapOfSizeToString) Equals(other types.Value) bool {
	if other, ok := other.(MapOfSizeToString); ok {
		return m.m.Equals(other.m)
	}
	return false
}

func (m MapOfSizeToString) Ref() ref.Ref {
	return m.m.Ref()
}

func (m MapOfSizeToString) Chunks() (futures []types.Future) {
	futures = append(futures, m.TypeRef().Chunks()...)
	futures = append(futures, m.m.Chunks()...)
	return
}

// A Noms Value that describes MapOfSizeToString.
var __typeRefForMapOfSizeToString types.TypeRef

func (m MapOfSizeToString) TypeRef() types.TypeRef {
	return __typeRefForMapOfSizeToString
}

func init() {
	__typeRefForMapOfSizeToString = types.MakeCompoundTypeRef("", types.MapKind, types.MakeTypeRef(__sha1_4c734206e6aaef5464ff0e307c2f66751a1469dePackageInFile_sha1_4c734206e6aaef5464ff0e307c2f66751a1469de_CachedRef, 1), types.MakePrimitiveTypeRef(types.StringKind))
	types.RegisterFromValFunction(__typeRefForMapOfSizeToString, func(v types.Value) types.NomsValue {
		return MapOfSizeToStringFromVal(v)
	})
}

func (m MapOfSizeToString) Empty() bool {
	return m.m.Empty()
}

func (m MapOfSizeToString) Len() uint64 {
	return m.m.Len()
}

func (m MapOfSizeToString) Has(p Size) bool {
	return m.m.Has(p.NomsValue())
}

func (m MapOfSizeToString) Get(p Size) string {
	return m.m.Get(p.NomsValue()).(types.String).String()
}

func (m MapOfSizeToString) MaybeGet(p Size) (string, bool) {
	v, ok := m.m.MaybeGet(p.NomsValue())
	if !ok {
		return "", false
	}
	return v.(types.String).String(), ok
}

func (m MapOfSizeToString) Set(k Size, v string) MapOfSizeToString {
	return MapOfSizeToString{m.m.Set(k.NomsValue(), types.NewString(v))}
}

// TODO: Implement SetM?

func (m MapOfSizeToString) Remove(p Size) MapOfSizeToString {
	return MapOfSizeToString{m.m.Remove(p.NomsValue())}
}

type MapOfSizeToStringIterCallback func(k Size, v string) (stop bool)

func (m MapOfSizeToString) Iter(cb MapOfSizeToStringIterCallback) {
	m.m.Iter(func(k, v types.Value) bool {
		return cb(SizeFromVal(k), v.(types.String).String())
	})
}

type MapOfSizeToStringIterAllCallback func(k Size, v string)

func (m MapOfSizeToString) IterAll(cb MapOfSizeToStringIterAllCallback) {
	m.m.IterAll(func(k, v types.Value) {
		cb(SizeFromVal(k), v.(types.String).String())
	})
}

type MapOfSizeToStringFilterCallback func(k Size, v string) (keep bool)

func (m MapOfSizeToString) Filter(cb MapOfSizeToStringFilterCallback) MapOfSizeToString {
	nm := NewMapOfSizeToString()
	m.IterAll(func(k Size, v string) {
		if cb(k, v) {
			nm = nm.Set(k, v)
		}
	})
	return nm
}

// SetOfString

type SetOfString struct {
	s types.Set
}

func NewSetOfString() SetOfString {
	return SetOfString{types.NewSet()}
}

type SetOfStringDef map[string]bool

func (def SetOfStringDef) New() SetOfString {
	l := make([]types.Value, len(def))
	i := 0
	for d, _ := range def {
		l[i] = types.NewString(d)
		i++
	}
	return SetOfString{types.NewSet(l...)}
}

func (s SetOfString) Def() SetOfStringDef {
	def := make(map[string]bool, s.Len())
	s.s.Iter(func(v types.Value) bool {
		def[v.(types.String).String()] = true
		return false
	})
	return def
}

func SetOfStringFromVal(p types.Value) SetOfString {
	return SetOfString{p.(types.Set)}
}

func (s SetOfString) NomsValue() types.Value {
	return s.s
}

func (s SetOfString) Equals(other types.Value) bool {
	if other, ok := other.(SetOfString); ok {
		return s.s.Equals(other.s)
	}
	return false
}

func (s SetOfString) Ref() ref.Ref {
	return s.s.Ref()
}

func (s SetOfString) Chunks() (futures []types.Future) {
	futures = append(futures, s.TypeRef().Chunks()...)
	futures = append(futures, s.s.Chunks()...)
	return
}

// A Noms Value that describes SetOfString.
var __typeRefForSetOfString types.TypeRef

func (m SetOfString) TypeRef() types.TypeRef {
	return __typeRefForSetOfString
}

func init() {
	__typeRefForSetOfString = types.MakeCompoundTypeRef("", types.SetKind, types.MakePrimitiveTypeRef(types.StringKind))
	types.RegisterFromValFunction(__typeRefForSetOfString, func(v types.Value) types.NomsValue {
		return SetOfStringFromVal(v)
	})
}

func (s SetOfString) Empty() bool {
	return s.s.Empty()
}

func (s SetOfString) Len() uint64 {
	return s.s.Len()
}

func (s SetOfString) Has(p string) bool {
	return s.s.Has(types.NewString(p))
}

type SetOfStringIterCallback func(p string) (stop bool)

func (s SetOfString) Iter(cb SetOfStringIterCallback) {
	s.s.Iter(func(v types.Value) bool {
		return cb(v.(types.String).String())
	})
}

type SetOfStringIterAllCallback func(p string)

func (s SetOfString) IterAll(cb SetOfStringIterAllCallback) {
	s.s.IterAll(func(v types.Value) {
		cb(v.(types.String).String())
	})
}

type SetOfStringFilterCallback func(p string) (keep bool)

func (s SetOfString) Filter(cb SetOfStringFilterCallback) SetOfString {
	ns := NewSetOfString()
	s.IterAll(func(v string) {
		if cb(v) {
			ns = ns.Insert(v)
		}
	})
	return ns
}

func (s SetOfString) Insert(p ...string) SetOfString {
	return SetOfString{s.s.Insert(s.fromElemSlice(p)...)}
}

func (s SetOfString) Remove(p ...string) SetOfString {
	return SetOfString{s.s.Remove(s.fromElemSlice(p)...)}
}

func (s SetOfString) Union(others ...SetOfString) SetOfString {
	return SetOfString{s.s.Union(s.fromStructSlice(others)...)}
}

func (s SetOfString) Subtract(others ...SetOfString) SetOfString {
	return SetOfString{s.s.Subtract(s.fromStructSlice(others)...)}
}

func (s SetOfString) Any() string {
	return s.s.Any().(types.String).String()
}

func (s SetOfString) fromStructSlice(p []SetOfString) []types.Set {
	r := make([]types.Set, len(p))
	for i, v := range p {
		r[i] = v.s
	}
	return r
}

func (s SetOfString) fromElemSlice(p []string) []types.Value {
	r := make([]types.Value, len(p))
	for i, v := range p {
		r[i] = types.NewString(v)
	}
	return r
}
