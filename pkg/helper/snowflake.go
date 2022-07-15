package helper

import (
	"fmt"

	"github.com/bwmarrin/snowflake"
)

func CreateSnowflakeBase64() string {
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	id := node.Generate()

	return id.Base64()
}
