package main

import (
    "fmt"
)

type author struct {
    name, branch string
    books, salary int
}

func (a author) show() {
    fmt.Println("Author's Name: ", a.name)
    fmt.Println("Branch Name: ", a.branch)
    fmt.Println("Published articles: ", a.books)
    fmt.Println("Salary ", a.salary)
}

func (a *author) bookup() {
    a.books++
}

func main() {
    res := author{
        name: "RZFeeser",
        branch: "Pennsylvania",
        books: 14,
        salary: 34000,
    }

    res.show()
    res.bookup()
    res.show()
}
