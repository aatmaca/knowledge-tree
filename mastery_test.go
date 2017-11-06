package main

import (
	"reflect"
	"testing"
)

func TestMasteryTree_EncodedPath(t *testing.T) {

	var ep = EncodedPath("a(b(c,d),e)")
	paths, _ := ep.Decode()
	nodes, nodeStructs, root, _ := pathArrayToNodeArray(paths)
	var mt = MasteryTree{ep, root, nodes, nodeStructs}

	tests := []struct {
		name string
		mt   MasteryTree
		want []string
	}{{"testData1", mt, []string{"a", "a:b", "a:e", "a:b:c", "a:b:d"}}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mt.EncodedPath(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MasteryTree.EncodedPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMasteryTree_FindNode(t *testing.T) {

	var ep = EncodedPath("a(b(c,d),e)")
	paths, _ := ep.Decode()
	nodes, nodeStructs, root, _ := pathArrayToNodeArray(paths)
	var mt = MasteryTree{ep, root, nodes, nodeStructs}

	var testData Node = root.Children[0].Children[0]

	tests := []struct {
		name    string
		mt      MasteryTree
		path    string
		want    Node
		wantErr bool
	}{{"testData1", mt, "a:b:c", testData, false},
		{"testData2", mt, "a:kb:c", nil, true}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.mt.FindNode(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("MasteryTree.FindNode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MasteryTree.FindNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMasteryTree_SortTree(t *testing.T) {

	var ep = EncodedPath("a(b(c,d),e)")
	paths, _ := ep.Decode()
	nodes, nodeStructs, root, _ := pathArrayToNodeArray(paths)
	var mt = MasteryTree{ep, root, nodes, nodeStructs}

	var testData = root.Children[0].Children[0].Value

	type args struct {
		sortParam string
		path      string
	}
	tests := []struct {
		name string
		mt   MasteryTree
		args args
		want string
	}{{"testData1", mt, args{"+", "a:b:c"}, testData}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mt.Update(tt.args.path, nil)
			tt.mt.SortTree(tt.args.sortParam)
			var got = root.Children[1].Children[1].Value
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MasteryTree.SortTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMasteryTree_Update(t *testing.T) {

	var ep = EncodedPath("a(b(c,d),e)")
	paths, _ := ep.Decode()
	nodes, nodeStructs, root, _ := pathArrayToNodeArray(paths)
	var mt = MasteryTree{ep, root, nodes, nodeStructs}

	type args struct {
		path string
		val  interface{}
	}
	tests := []struct {
		name    string
		mt      MasteryTree
		args    args
		want    float64
		wantErr bool
	}{{"testData1", mt, args{"a:b:d", nil}, 1, false}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.mt.Update(tt.args.path, tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("MasteryTree.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
			var got = root.Children[0].Children[1].scoreValue
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MasteryTree.Update() = %v, want %v", got, tt.want)
			}

		})
	}
}

func TestMasteryTree_Nodes(t *testing.T) {

	var ep = EncodedPath("a(b(c,d),e)")
	paths, _ := ep.Decode()
	nodes, nodeStructs, root, _ := pathArrayToNodeArray(paths)
	var mt = MasteryTree{ep, root, nodes, nodeStructs}

	tests := []struct {
		name string
		mt   MasteryTree
		want int
	}{{"testData1", mt, 5}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.mt.Nodes()
			if !reflect.DeepEqual(len(got), tt.want) {
				t.Errorf("MasteryTree.Nodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMastery(t *testing.T) {

	var ep = EncodedPath("a(b(c,d),e)")
	type args struct {
		ep EncodedPath
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{{"testData1", args{ep}, "a", false}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewMastery(tt.args.ep)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMastery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Nodes()[0].Path(), tt.want) {
				t.Errorf("NewMastery() = %v, want %v", got, tt.want)
			}
		})
	}
}
