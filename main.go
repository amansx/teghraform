package main

func main() {
	BuildIndex()
	contents := ReadFeature(CWD() + "/examples/s3.feature")
	LoadFeature(contents)
}
