package mesquite

import "fmt"


func Hello(name string) string {
    // Return greeting
    message := fmt.Sprintf("Howdy, %v...", name)
    return message
}