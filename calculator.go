package main

    import (
        "flag"
        "fmt"
        "os"
    )

    func main() {
        add := flag.Bool("add", false, "Add two numbers")
        subs := flag.Bool("subtract", false, "Subtract two numbers")
        mult := flag.Bool("multiply", false, "Multiply two numbers")
        div := flag.Bool("divide", false, "Divide two numbers")

        flag.Parse()

        var a, b int
        fmt.Println("Enter 1st Number: ")
        fmt.Scan(&a)
        fmt.Println("Enter 2nd Number: ")
        fmt.Scan(&b)

        switch {
        case *add:
            fmt.Printf("Additon: %d \n", addition(a, b))
        case *subs:
            fmt.Printf("Difference: %d \n", subtract(a, b))
        case *mult:
            fmt.Printf("Product: %d \n", multiply(a, b))
        case *div:
            fmt.Printf("Division: %d \n", division(a, b))
        default:
            fmt.Fprintln(os.Stderr, "Wrong option try with add, subtract, div and multply")
            os.Exit(1)
        }

    }

    func addition(a int, b int) int {
        return a + b
    }

    func subtract(a int, b int) int {
        return a - b
    }

    func multiply(a int, b int) int {
        return a * b
    }

    func division(a int, b int) int {
        return a / b
    }