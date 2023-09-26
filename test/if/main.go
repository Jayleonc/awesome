package main

import "runtime"

func main() {

	checkOS(runtime.GOOS, runtime.GOARCH, runtime.Compiler)

	a, b := false, false
	if a && b == false {
		println("(a && b) != true")
		return
	}
	println("a && (b != true) == false")

}

func checkOS(os, arch, compiler string) {
	if os == "linux" && arch == "amd64" && compiler != "gccgo" {
		println("we are using standard go compiler on linux os for amd64")
	} else if (os == "darwin") && (arch == "arm64") && (compiler != "gccgo") {
		println("we are using standard go compiler on macOS os for arm64")
	}
}
