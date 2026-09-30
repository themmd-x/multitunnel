package minecraft

import (
	"fmt"
	"reflect"

	"github.com/TheMMD-X/multitunnel/minecraft/protocol/packet"

	v121100packet "github.com/TheMMD-X/multitunnel/libs/gophertunnel/v121100/minecraft/protocol/packet"
	v121111packet "github.com/TheMMD-X/multitunnel/libs/gophertunnel/v121111/minecraft/protocol/packet"
	v121120packet "github.com/TheMMD-X/multitunnel/libs/gophertunnel/v121120/minecraft/protocol/packet"
	v121130packet "github.com/TheMMD-X/multitunnel/libs/gophertunnel/v121130/minecraft/protocol/packet"
	v1212packet "github.com/TheMMD-X/multitunnel/libs/gophertunnel/v1212/minecraft/protocol/packet"
	v12120packet "github.com/TheMMD-X/multitunnel/libs/gophertunnel/v12120/minecraft/protocol/packet"
	v12130packet "github.com/TheMMD-X/multitunnel/libs/gophertunnel/v12130/minecraft/protocol/packet"
	v12140packet "github.com/TheMMD-X/multitunnel/libs/gophertunnel/v12140/minecraft/protocol/packet"
	v12150packet "github.com/TheMMD-X/multitunnel/libs/gophertunnel/v12150/minecraft/protocol/packet"
	v12160packet "github.com/TheMMD-X/multitunnel/libs/gophertunnel/v12160/minecraft/protocol/packet"
	v12170packet "github.com/TheMMD-X/multitunnel/libs/gophertunnel/v12170/minecraft/protocol/packet"
	v12180packet "github.com/TheMMD-X/multitunnel/libs/gophertunnel/v12180/minecraft/protocol/packet"
	v12190packet "github.com/TheMMD-X/multitunnel/libs/gophertunnel/v12190/minecraft/protocol/packet"
	v12193packet "github.com/TheMMD-X/multitunnel/libs/gophertunnel/v12193/minecraft/protocol/packet"
	v12610packet "github.com/TheMMD-X/multitunnel/libs/gophertunnel/v12610/minecraft/protocol/packet"
	v12630packet "github.com/TheMMD-X/multitunnel/libs/gophertunnel/v12630/minecraft/protocol/packet"
	v12640packet "github.com/TheMMD-X/multitunnel/libs/gophertunnel/v12640/minecraft/protocol/packet"
)

type typeRegistry map[string]reflect.Type

func buildRegistryOf[P any](pools ...map[uint32]func() P) typeRegistry {
	reg := make(typeRegistry)
	for _, pool := range pools {
		for _, factory := range pool {
			t := reflect.TypeOf(factory()).Elem()
			reg[t.Name()] = t
		}
	}
	return reg
}

var canonicalTypes = buildRegistryOf(packet.NewClientPool(), packet.NewServerPool())

var vendorTypes = map[string]typeRegistry{
	"1.21.100": buildRegistryOf(v121100packet.NewClientPool(), v121100packet.NewServerPool()),
	"1.21.111": buildRegistryOf(v121111packet.NewClientPool(), v121111packet.NewServerPool()),
	"1.21.120": buildRegistryOf(v121120packet.NewClientPool(), v121120packet.NewServerPool()),
	"1.21.130": buildRegistryOf(v121130packet.NewClientPool(), v121130packet.NewServerPool()),
	"1.21.2":   buildRegistryOf(v1212packet.NewClientPool(), v1212packet.NewServerPool()),
	"1.21.20":  buildRegistryOf(v12120packet.NewClientPool(), v12120packet.NewServerPool()),
	"1.21.30":  buildRegistryOf(v12130packet.NewClientPool(), v12130packet.NewServerPool()),
	"1.21.40":  buildRegistryOf(v12140packet.NewClientPool(), v12140packet.NewServerPool()),
	"1.21.50":  buildRegistryOf(v12150packet.NewClientPool(), v12150packet.NewServerPool()),
	"1.21.60":  buildRegistryOf(v12160packet.NewClientPool(), v12160packet.NewServerPool()),
	"1.21.70":  buildRegistryOf(v12170packet.NewClientPool(), v12170packet.NewServerPool()),
	"1.21.80":  buildRegistryOf(v12180packet.NewClientPool(), v12180packet.NewServerPool()),
	"1.21.90":  buildRegistryOf(v12190packet.NewClientPool(), v12190packet.NewServerPool()),
	"1.21.93":  buildRegistryOf(v12193packet.NewClientPool(), v12193packet.NewServerPool()),
	"1.26.10":  buildRegistryOf(v12610packet.NewClientPool(), v12610packet.NewServerPool()),
	"1.26.30":  buildRegistryOf(v12630packet.NewClientPool(), v12630packet.NewServerPool()),
	"1.26.40":  buildRegistryOf(v12640packet.NewClientPool(), v12640packet.NewServerPool()),
}

func registryFor(version string) (typeRegistry, error) {
	reg, ok := vendorTypes[version]
	if !ok {
		return nil, fmt.Errorf("minecraft: no vendor registered for version %q", version)
	}
	return reg, nil
}

func fromVendor(vendorPk any) (packet.Packet, error) {
	src := reflect.ValueOf(vendorPk)
	if src.Kind() != reflect.Ptr || src.IsNil() {
		return nil, fmt.Errorf("minecraft: fromVendor: expected a non-nil pointer, got %T", vendorPk)
	}
	name := src.Elem().Type().Name()
	dstType, ok := canonicalTypes[name]
	if !ok {
		return nil, fmt.Errorf("minecraft: fromVendor: no canonical packet named %q", name)
	}
	dst := reflect.New(dstType)
	copyStruct(src.Elem(), dst.Elem())
	pk, ok := dst.Interface().(packet.Packet)
	if !ok {
		return nil, fmt.Errorf("minecraft: fromVendor: %q does not implement Packet", name)
	}
	return pk, nil
}

func toVendor(pk packet.Packet, version string) (any, error) {
	reg, err := registryFor(version)
	if err != nil {
		return nil, err
	}
	src := reflect.ValueOf(pk)
	if src.Kind() != reflect.Ptr || src.IsNil() {
		return nil, fmt.Errorf("minecraft: toVendor: expected a non-nil pointer, got %T", pk)
	}
	name := src.Elem().Type().Name()
	dstType, ok := reg[name]
	if !ok {
		return nil, fmt.Errorf("minecraft: toVendor: version %q has no packet named %q", version, name)
	}
	dst := reflect.New(dstType)
	copyStruct(src.Elem(), dst.Elem())
	return dst.Interface(), nil
}

func copyStruct(src, dst reflect.Value) {
	srcType := src.Type()
	for i := 0; i < srcType.NumField(); i++ {
		sf := srcType.Field(i)
		if sf.PkgPath != "" {
			continue
		}
		df := dst.FieldByName(sf.Name)
		if !df.IsValid() || !df.CanSet() {
			continue
		}
		copyValue(src.Field(i), df)
	}
}

func copyValue(src, dst reflect.Value) {
	if !src.IsValid() || !dst.CanSet() {
		return
	}
	st, dt := src.Type(), dst.Type()

	if st.AssignableTo(dt) {
		dst.Set(src)
		return
	}
	if st.ConvertibleTo(dt) && st.Kind() == dt.Kind() && st.Kind() != reflect.Struct &&
		st.Kind() != reflect.Slice && st.Kind() != reflect.Map && st.Kind() != reflect.Ptr {
		dst.Set(src.Convert(dt))
		return
	}

	switch {
	case st.Kind() == reflect.Ptr && dt.Kind() == reflect.Ptr:
		if src.IsNil() {
			return
		}
		np := reflect.New(dt.Elem())
		copyValue(src.Elem(), np.Elem())
		dst.Set(np)

	case st.Kind() == reflect.Struct && dt.Kind() == reflect.Struct:
		copyStruct(src, dst)

	case st.Kind() == reflect.Slice && dt.Kind() == reflect.Slice:
		if src.IsNil() {
			return
		}
		n := src.Len()
		ns := reflect.MakeSlice(dt, n, n)
		for i := 0; i < n; i++ {
			copyValue(src.Index(i), ns.Index(i))
		}
		dst.Set(ns)

	case st.Kind() == reflect.Array && dt.Kind() == reflect.Array && st.Len() == dt.Len():
		for i := 0; i < st.Len(); i++ {
			copyValue(src.Index(i), dst.Index(i))
		}

	case st.Kind() == reflect.Map && dt.Kind() == reflect.Map:
		if src.IsNil() {
			return
		}
		nm := reflect.MakeMapWithSize(dt, src.Len())
		iter := src.MapRange()
		for iter.Next() {
			k := iter.Key()
			if !k.Type().AssignableTo(dt.Key()) {
				continue
			}
			nv := reflect.New(dt.Elem()).Elem()
			copyValue(iter.Value(), nv)
			nm.SetMapIndex(k, nv)
		}
		dst.Set(nm)

	default:
	}
}
