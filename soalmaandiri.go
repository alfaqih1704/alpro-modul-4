
package main

import "fmt"

func main() {
    var a, b rune

    fmt.Scanf("%c%c", &a, &b)

    hasil1 := !((a >= 'A' && a <= 'Z') ||
        (a >= 'a' && a <= 'z')) &&
        !(a >= '0' && a <= '9')

    hasil2 := !((b >= 'A' && b <= 'Z') ||
        (b >= 'a' && b <= 'z')) &&
        !(b >= '0' && b <= '9')

    fmt.Println(hasil1, hasil2)
}