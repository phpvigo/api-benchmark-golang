package main

func main() {
    a := App{}
    a.Initialize("benchmark", "benchmark", "db")

    a.Run(":8080")
}
