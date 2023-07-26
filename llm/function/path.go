// Package function
// @Description: 用于猜解类型文件应该被保存的位置和其文件名
package function

import "github.com/mewway/go-laravel/support"

func Resolve(pathType string) string {
	path := ""
	switch pathType {
	case support.DirService:
	case support.DirLogic:
	case support.DirDatabase:
	case support.DirRoute:
	case support.DirListener:
	case support.DirJob:
	case support.DirEvent:
	case support.DirCommand:
	}

	return path
}
