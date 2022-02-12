package solution

import (
	"github.com/kyokomi/emoji/v2"

	"fmt"
)

func GetMessage()  {
	hello_world := emoji.Sprint("hello :world_map:")
	fmt.Println(hello_world)
	
}
