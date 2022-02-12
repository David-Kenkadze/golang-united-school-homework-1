package solution

import (
	"github.com/kyokomi/emoji/v2"

	"fmt"
)


func GetMessage() string {
        hello_message := emoji.Sprint("hello :world_map:!")
        fmt.Println(hello_message)
        return "hello-message"
}


