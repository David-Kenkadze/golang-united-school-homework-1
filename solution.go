package solution

import (
	"github.com/kyokomi/emoji/v2"

	"fmt"
)

func GetMessage() string {
	GetMessage := emoji.Sprint("hello world :world_map:!!")
	fmt.Println(GetMessage)
	return GetMessage
}
