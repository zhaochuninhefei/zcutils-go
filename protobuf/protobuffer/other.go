package protobuffer

import (
	"errors"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/protoadapt"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoiface"
	"google.golang.org/protobuf/runtime/protoimpl"
)

var (
	ErrNil = errors.New("proto: Marshal called with nil")
)

func marshalAppend(buf []byte, m protoadapt.MessageV1, deterministic bool) ([]byte, error) {
	if m == nil {
		return nil, ErrNil
	}
	mi := protoadapt.MessageV2Of(m)
	nbuf, err := proto.MarshalOptions{
		Deterministic: deterministic,
		AllowPartial:  true,
	}.MarshalAppend(buf, mi)
	if err != nil {
		return buf, err
	}
	if len(buf) == len(nbuf) {
		if !mi.ProtoReflect().IsValid() {
			return buf, ErrNil
		}
	}
	return nbuf, checkRequiredNotSet(mi)
}

// RequiredNotSetError is an error type returned when
// marshaling or unmarshaling a message with missing required fields.
type RequiredNotSetError struct {
	err error
}

func (e *RequiredNotSetError) Error() string {
	if e.err != nil {
		return e.err.Error()
	}
	return "proto: required field not set"
}
func (e *RequiredNotSetError) RequiredNotSet() bool {
	return true
}

func checkRequiredNotSet(m proto.Message) error {
	if err := proto.CheckInitialized(m); err != nil {
		return &RequiredNotSetError{err: err}
	}
	return nil
}

// unmarshalMerge parses a wire-format message in b and places the decoded results in m.
func unmarshalMerge(b []byte, m protoadapt.MessageV1) error {
	mi := protoadapt.MessageV2Of(m)
	out, err := proto.UnmarshalOptions{
		AllowPartial: true,
		Merge:        true,
	}.UnmarshalState(protoiface.UnmarshalInput{
		Buf:     b,
		Message: mi.ProtoReflect(),
	})
	if err != nil {
		return err
	}
	if out.Flags&protoiface.UnmarshalInitialized > 0 {
		return nil
	}
	return checkRequiredNotSet(mi)
}

// messageReflect returns a reflective view for a message.
// It returns nil if m is nil.
func messageReflect(m protoadapt.MessageV1) protoreflect.Message {
	return protoimpl.X.MessageOf(m)
}

// size returns the size in bytes of the wire-format encoding of m.
func size(m protoadapt.MessageV1) int {
	if m == nil {
		return 0
	}
	mi := protoadapt.MessageV2Of(m)
	return proto.Size(mi)
}
