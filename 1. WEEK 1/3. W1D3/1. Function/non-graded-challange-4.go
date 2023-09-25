package main

import "fmt"

func convertToAlay(char byte) string {
    // Fungsi konversi karakter menjadi karakter alay
    switch char {
    case 'a':
        return "4"
    case 'e':
        return "3"
    case 'i':
        return "!"
    case 'l':
        return "1"
    case 'n':
        return "N"
    case 's':
        return "5"
    case 'x':
        return "*"
    default:
        return string(char)
    }
}

func main1() {
    char := 'a' // Karakter (byte) yang ingin diubah menjadi alay
    alayChar := convertToAlay(byte(char))
    
    fmt.Println(alayChar)
}
