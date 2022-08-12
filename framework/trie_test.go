package framework

import (
	"testing"
)

func Test_filterChildNodes(t *testing.T) {
	root := &node{
		isLast: false,
		segment: "",
		handler: func( *Context) error {return nil},
		childs: []*node{
			{
				isLast: true,
				segment: "FOO",
				handler: func( *Context) error {return nil},
				childs:nil,
			},
			{
				isLast: false,
				segment: ":id",
				handler: func( *Context) error {return nil},
				childs:nil,
			},
		},
	}
	{
		nodes := root.filterChildNodes("FOO")
		if len(nodes) != 2 {
			t.Error("foo error")
		}
	}
	{
		nodes := root.filterChildNodes(":foo")
		if len(nodes) != 2 {
			t.Error(":foo error")
		}
	}
}

func Test_matchNode(t *testing.T) {
	root := &node{
		isLast: false,
		segment: "",
		handler: func( *Context) error {return nil},
		childs: []*node{
			{
				isLast: true,
				segment: "FOO",
				handler: func( *Context) error {return nil},
				childs: []*node{
					&node{
						isLast: true,
						segment: "BAR",
						handler: func( *Context) error {return nil},
						childs:[]*node{},
					},
				},

		    },
			{
							isLast:true,
							segment:":id",
							handler:nil,
							childs:nil,
			},
		},
	}
	{
		node := root.mathNode("foo/bar")
		if node == nil {
			t.Error("math normal node error")
		}
	}

	{
		node := root.mathNode("test")
		if node == nil {
			t.Error("match test")
		}
	}
}
