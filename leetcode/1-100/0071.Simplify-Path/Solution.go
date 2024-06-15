package Solution

import (
	"path/filepath"
	"strings"
)

func simplifyPath(path string) string {
	return filepath.Clean(path)
}

func simplifyPath2(path string) string {
	var stack []string
	for _, component := range strings.Split(path, "/") {
		if component == "" || component == "." {
			continue
		}
		if component == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, component)
		}
	}
	return "/" + strings.Join(stack, "/")
}

func simplifyPath3(path string) string {
	paths := strings.Split(path, "/")
	stack := make([]string, 0, len(paths))
	for _, val := range paths {
		if len(val) == 0 {
			continue
		}
		if val == "." {
			continue
		}
		if val == ".." {
			if len(stack) == 0 {
				continue
			} else {
				stack = stack[0 : len(stack)-1]
				continue
			}
		}
		stack = append(stack, val)
	}
	return "/" + strings.Join(stack, "/")
}
