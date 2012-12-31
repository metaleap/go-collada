package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	ugo "github.com/metaleap/go-util"
	uio "github.com/metaleap/go-util/io"
)

func main() {
	//	pkgreflect -novars=true -norecurs=true -nofuncs=false -gofile=-pkgreflect.go $dir
	runtime.LockOSThread()
	srcDirPath := os.Args[1]
	impPath := strings.Trim(strings.Replace(strings.Replace(strings.ToLower(srcDirPath), strings.ToLower(ugo.GopathSrc()), "", -1), "\\", "/", -1), "/")
	srcFilePath := filepath.Join(srcDirPath, "-pkgreflect.go")
	outDirPath := filepath.Join(srcDirPath, "pkgreflect")
	outFilePath := filepath.Join(outDirPath, "-gen-pkgreflect.go")
	if raw, err := exec.Command("pkgreflect", "-novars=true", "-norecurs=true", "-nofuncs=false", "-gofile=-pkgreflect.go", srcDirPath).CombinedOutput(); err != nil {
		print(err)
		panic(err)
	} else {
		if len(raw) > 0 {
			println(string(raw))
		}
		src := uio.ReadTextFile(srcFilePath, true, "")
		os.Remove(srcFilePath)
		src = strings.Replace(src, "package cdom", "package pkgreflect", -1)
		src = strings.Replace(src, "import \"reflect\"", "import \"reflect\"\nimport cdom \""+impPath+"\"", -1)
		src = strings.Replace(src, "((*", "((*cdom.", -1)
		src = strings.Replace(src, "ValueOf(", "ValueOf(cdom.", -1)
		uio.WriteTextFile(outFilePath, src)
	}
}
