package briantest

import (
	"testing"

	"github.com/briansorahan/death"
	"github.com/gogo/protobuf/jsonpb"
)

func TestMarshal(t *testing.T) {
	t.Run("with no gogoproto.jsontag", func(t *testing.T) {
		var (
			die       = death.By(t)
			marshaler = &jsonpb.Marshaler{}
		)
		data, err := marshaler.MarshalToString(&Msg{
			AField: "foo",
			BField: "bar",
		})
		die(err)

		if expected, got := `{"aField":"foo","bField":"bar"}`, data; expected != got {
			t.Fatalf("expected %s, got %s", expected, got)
		}
	})

	t.Run("using gogoproto.jsontag", func(t *testing.T) {
		var (
			die       = death.By(t)
			marshaler = &jsonpb.Marshaler{}
		)
		data, err := marshaler.MarshalToString(&MsgTags{
			AField: "foo",
			BField: "bar",
		})
		die(err)

		if expected, got := `{"a":"foo","b":"bar"}`, data; expected != got {
			t.Fatalf("expected %s, got %s", expected, got)
		}
	})
}
