package main

func main() {
	if err := CmdRoot.Execute(); err != nil {
		panic(err)
	}
}
