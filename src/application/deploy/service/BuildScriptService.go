package service

import (
	"script-go/src/application/deploy/pojo/dto/BuildConfig"
	"script-go/src/application/deploy/pojo/po/Script"
	"script-go/src/application/util/GenUtil"
)

type BuildScriptService struct {
	cgo         int
	os          string
	arch        string
	scripts     []*Script.Script
	buildConfig *BuildConfig.BuildConfig
}

func newBuildScriptService() *BuildScriptService {
	return &BuildScriptService{
		cgo:         1,
		os:          "windows",
		arch:        "amd64",
		scripts:     Script.Gets(),
		buildConfig: BuildConfig.Get(),
	}
}

func (b *BuildScriptService) apply() {
	GenUtil.Println()
	for i, script := range b.scripts {
		GenUtil.Println(GenUtil.IntToString(i+1) + ". " + script.GoName())
	}
	GenUtil.Print("Please enter one or more numbers corresponding to the script: ")
	scriptNums := GenUtil.ReadParams()
	if len(scriptNums) == 0 {
		return
	}
	GenUtil.Println()

	lstOS := BuildConfig.GetMapOSKeys(b.buildConfig.MapOS())
	for i, os := range lstOS {
		GenUtil.Println(GenUtil.IntToString(i+1) + ". " + os)
	}
	GenUtil.Print("Please enter one number corresponding to the GOOS: ")
	osNums := GenUtil.ReadParams()
	if len(osNums) > 0 {
		b.os = lstOS[GenUtil.StrToInt(osNums[0])-1]
	}
	GenUtil.Println()

	lstArch := b.buildConfig.MapOS()[b.os]
	for i, arch := range lstArch {
		GenUtil.Println(GenUtil.IntToString(i+1) + ". " + arch)
	}
	GenUtil.Print("Please enter one number corresponding to the GOARCH: ")
	archNums := GenUtil.ReadParams()
	if len(archNums) > 0 {
		b.arch = lstOS[GenUtil.StrToInt(archNums[0])-1]
	}
	GenUtil.Println()

	GenUtil.Println("1. cgo enable")
	GenUtil.Println("2. cgo disable")
	GenUtil.Print("Please enter one number corresponding to the CGO_ENABLED: ")
	cgoNums := GenUtil.ReadParams()
	if len(cgoNums) > 0 {
		b.cgo = GenUtil.StrToInt(cgoNums[0])
		if b.cgo == 2 {
			b.cgo = 0
		}
	}
	GenUtil.Println()

	for _, scriptNum := range scriptNums {
		index := GenUtil.StrToInt(scriptNum) - 1
		if 0 <= index && index < len(b.scripts) {
			b.build(b.scripts[index])
		}
		if index == len(b.scripts) {

		}
	}
}

func (b *BuildScriptService) build(script *Script.Script) {

}

func Run() {
	newBuildScriptService().apply()
}
