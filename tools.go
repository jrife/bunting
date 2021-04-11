// +build tools

package tools

import (
	_ "github.com/gogo/protobuf/gogoproto"
	_ "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/protoc-gen-gogofaster"
	_ "github.com/golang/protobuf/protoc-gen-go"
)
