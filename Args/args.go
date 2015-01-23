func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ", os.Args[0], "host seconds")
		os.Exit(1) 
	}
	url := os.Args[1]
	seconds, err := strconv.Atoi(os.Args[2])
    if err != nil {
        fmt.Println(err)
        os.Exit(2)
    }
}
