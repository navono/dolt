package ref

import (
	"errors"
	"fmt"
	"strings"
)

// ErrUnknownRefType is the error returned when parsing a ref in the format refs/type/... where type is unknown
var ErrUnknownRefType = errors.New("unknown ref type")

const (
	refPrefix     = "refs/"
	remotesPrefix = "remotes/"
)

// IsRef returns true if the string is a reference string (meanings it starts with the prefix refs/)
func IsRef(str string) bool {
	return strings.HasPrefix(str, refPrefix)
}

// RefType is the type of the reference, and this value follows the refPrefix in a ref string.  e.g. refs/type/...
type RefType string

const (
	// BranchRefType is a reference to a local branch in the format refs/heads/...
	BranchRefType RefType = "heads"

	// RemoteRefType is a reference to a local remote tracking branch
	RemoteRefType RefType = "remotes"

	// InternalRefType is a reference to a dolt internal commit
	InternalRefType RefType = "internal"
)

// RefTypes is the set of all supported reference types.  External RefTypes can be added to this map in order to add
// RefTypes for external tooling
var RefTypes = map[RefType]struct{}{BranchRefType: {}, RemoteRefType: {}, InternalRefType: {}}

// PrefixForType returns what a reference string for a given type should start with
func PrefixForType(refType RefType) string {
	return refPrefix + string(refType) + "/"
}

// DoltRef is a reference to a commit.
type DoltRef interface {
	fmt.Stringer

	// Type is the RefType of this ref
	GetType() RefType

	// Path is the identifier for the reference
	GetPath() string
}

// Equals returns true if two DoltRefs have the same Type and Path
func Equals(dr, other DoltRef) bool {
	if dr == nil && other == nil {
		return true
	} else if dr == nil || other == nil {
		return false
	}

	return dr.GetType() == other.GetType() && dr.GetPath() == other.GetPath()
}

// EqualsStr compares a DoltRef to a reference string to see if they are refering to the same thing
func EqualsStr(dr DoltRef, str string) bool {
	other, err := Parse(str)

	if err != nil {
		return false
	}

	return Equals(dr, other)
}

// String() converts the DoltRef to a reference string in the format refs/type/path
func String(dr DoltRef) string {
	return PrefixForType(dr.GetType()) + dr.GetPath()
}

// MarshalJson implements the json Marshaler interface to json encode DoltRefs as their string representation
func MarshalJSON(dr DoltRef) ([]byte, error) {
	str := dr.String()
	data := make([]byte, len(str)+2)

	data[0] = '"'
	data[len(str)+1] = '"'

	for i, b := range str {
		data[i+1] = byte(b)
	}

	return data, nil
}

// Parse will parse ref strings and return a DoltRef or an error for refs that can't be parsed.
func Parse(str string) (DoltRef, error) {
	if !IsRef(str) {
		if strings.HasPrefix(str, remotesPrefix) {
			return NewRemoteRefFromPathStr(str)
		} else {
			return NewBranchRef(str), nil
		}
	}

	for rType := range RefTypes {
		prefix := PrefixForType(rType)
		if strings.HasPrefix(str, prefix) {
			str = str[len(prefix):]
			switch rType {
			case BranchRefType:
				return NewBranchRef(str), nil
			case RemoteRefType:
				return NewRemoteRefFromPathStr(str)
			case InternalRefType:
				return NewInternalRef(str), nil
			default:
				panic("unknown type " + rType)
			}
		}
	}

	return nil, ErrUnknownRefType
}
