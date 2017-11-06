package main

import (
	"reflect"
	"testing"
)

func TestDecode(t *testing.T) {
	tests := []struct {
		name    string
		ep      EncodedPath
		want    []string
		wantErr bool
	}{{"testData1", EncodedPath("a(b(c,d),e)"), []string{"a", "a:b", "a:e", "a:b:c", "a:b:d"}, false}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ep.Decode()
			if (err != nil) != tt.wantErr {
				t.Errorf("EncodedPath.Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EncodedPath.Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncode(t *testing.T) {
	tests := []struct {
		name    string
		paths   []string
		want    EncodedPath
		wantErr bool
	}{{"testData1", []string{"a", "a:b", "a:e", "a:b:c", "a:b:d"}, EncodedPath("a(b(c,d),e)"), false}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Encode(tt.paths)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNodeStruct_Path(t *testing.T) {

	paths, _ := EncodedPath("a(b(c,d),e)").Decode()
	_, _, root, _ := pathArrayToNodeArray(paths)

	var testData = root.Children[0].Children[0]

	tests := []struct {
		name string
		node *NodeStruct
		want string
	}{{"testData1", testData, "a:b:c"}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.node.Path(); got != tt.want {
				t.Errorf("NodeStruct.Path() = %v, want %v", got, tt.want)
			}
		})
	}
}
