package main

func main() {
    a := App{}
    a.Initialize("root", "1telefono", "rest_api_example")

    a.Run(":8080")
}
