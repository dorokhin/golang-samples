package main

import (
    "fmt"
    "os"
    "math/big"
)


func main() {
    if len(os.Args) == 1 {
        fmt.Println("🤨 Oh no. Usage: factorial <factorial number>")
        os.Exit(1)
    }
    arguments := os.Args
    factNum := new(big.Int)
    factNum, _ = factNum.SetString(arguments[1], 10)
    fmt.Println("Factorial is:", factorial(factNum), "😏")
}

func Mul(x, y *big.Int) *big.Int {
    return big.NewInt(0).Mul(x, y)
}

func Sub(x, y *big.Int) *big.Int {
    return big.NewInt(0).Sub(x, y)
}

func factorial(num *big.Int) *big.Int {
    if num.Cmp(big.NewInt(0)) ==  0{
        return big.NewInt(1)
    }
    newFact := big.NewInt(0)
    newFact = factorial(Sub(num, big.NewInt(1)))
    return Mul(num, newFact)
}
