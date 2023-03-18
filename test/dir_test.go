package test

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"
)

type Node struct {
	Name     string
	Path     string
	IsDir    bool
	Children []*Node
}

func TestName(t *testing.T) {

	rootPath := "/Users/zylzyl/go/src/fsm_client" // 指定根目录
	root := &Node{
		Name:  rootPath,
		Path:  rootPath,
		IsDir: true,
	}
	buildTree(root)
	printTree(root, 0)
}

func buildTree(node *Node) {
	fileInfos, err := ioutil.ReadDir(node.Path)
	if err != nil {
		fmt.Printf("读取目录失败：%s\n", node.Path)
		return
	}
	for _, fileInfo := range fileInfos {
		child := &Node{
			Name:  fileInfo.Name(),
			Path:  filepath.Join(node.Path, fileInfo.Name()),
			IsDir: fileInfo.IsDir(),
		}
		node.Children = append(node.Children, child)
		if child.IsDir {
			buildTree(child)
		}
	}
}

func printTree(node *Node, depth int) {
	prefix := ""
	for i := 0; i < depth; i++ {
		prefix += "  "
	}
	fmt.Printf("%s%s\n", prefix, node.Name)
	for _, child := range node.Children {
		printTree(child, depth+1)
	}
}
