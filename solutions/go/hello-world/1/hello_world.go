package greeting
import "fmt"

// HelloWorld greets the world.
func HelloWorld() string {
	return "Hello, World!"
}

func main() {
    text := HelloWorld()
    fmt.Println(text)
}
