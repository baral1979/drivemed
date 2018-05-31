package main

func main() {
	a := App{}
	a.Initialize("localhost", "drive", "drive", "DriveMedical", 1433)
	a.Run(":8080")
}
