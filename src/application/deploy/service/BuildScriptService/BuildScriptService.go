package BuildScriptService

import (
	"script-go/src/application/deploy/pojo/dto/BuildConfig"
	"script-go/src/application/deploy/pojo/po/Script"
	"script-go/src/application/util/FileUtil"
	"script-go/src/application/util/GenUtil"
	"script-go/src/application/util/PromptUtil"
	"script-go/src/application/util/RemoteUtil"
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
		GenUtil.Println(GenUtil.IntToString(i+1) + ". " + script.GoName)
	}
	GenUtil.Print("Please enter one or more numbers corresponding to the script: ")
	scriptNums := GenUtil.ReadParams()
	if len(scriptNums) == 0 {
		return
	}
	GenUtil.Println()

	lstOS := []string{"windows", "linux", "android", "darwin", "ios"}
	for i, os := range lstOS {
		GenUtil.Println(GenUtil.IntToString(i+1) + ". " + os)
	}
	GenUtil.Print("Please enter one number corresponding to the GOOS: ")
	osNums := GenUtil.ReadParams()
	if len(osNums) > 0 {
		b.os = lstOS[GenUtil.StrToInt(osNums[0])-1]
	}
	GenUtil.Println()

	lstArch := b.buildConfig.MapOS[b.os]
	for i, arch := range lstArch {
		GenUtil.Println(GenUtil.IntToString(i+1) + ". " + arch)
	}
	GenUtil.Print("Please enter one number corresponding to the GOARCH: ")
	archNums := GenUtil.ReadParams()
	if len(archNums) > 0 {
		b.arch = lstArch[GenUtil.StrToInt(archNums[0])-1]
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
			Script.SetDistPath(b.scripts[index], b.os, b.arch)
			b.build(b.scripts[index])
		}
	}
}

func (b *BuildScriptService) build(script *Script.Script) {
	b.changeBuildConfig(script, true)
	b.changeCrossBuild(script, true)

	bin, args := PromptUtil.PackageGoScript(b.buildConfig.CrossBuildPath)
	RemoteUtil.ChangeWorkFolder(b.buildConfig.SrcPath)
	RemoteUtil.ExecLocalCmd(bin, args...)

	b.updateScript(script)
	b.changeCrossBuild(script, false)
	b.changeBuildConfig(script, false)
}

func (b *BuildScriptService) updateScript(script *Script.Script) {
	if !FileUtil.Exist(script.ScriptProject) {
		FileUtil.Mkdir(script.ScriptProject)
	}
	if FileUtil.Exist(script.ScriptPath) &&
		FileUtil.Exist(script.DistPath) &&
		script.GoName != "BuildScriptService.go" {
		FileUtil.Delete(script.ScriptPath)
	}
	if FileUtil.Exist(script.ScriptConfig) &&
		FileUtil.Exist(script.YamlConfig) {
		FileUtil.Delete(script.ScriptConfig)
	}
	if FileUtil.Exist(script.DistPath) &&
		script.GoName != "BuildScriptService.go" {
		FileUtil.Copy(script.DistPath, script.ScriptPath)
	}
	if FileUtil.Exist(script.YamlConfig) {
		FileUtil.Copy(script.YamlConfig, script.ScriptConfig)
	}
}

func (b *BuildScriptService) changeBuildConfig(script *Script.Script, isBefore bool) {
	scriptRun := script.ScriptRun
	scriptImport := script.ScriptImport
	if !isBefore {
		scriptRun = b.buildConfig.ScriptRunOriginal
		scriptImport = b.buildConfig.PackageImportOriginal
	}
	FileUtil.ModContent(b.buildConfig.AppPath, b.buildConfig.ScriptRunPattern, false, scriptRun)
	FileUtil.ModContent(b.buildConfig.AppPath, b.buildConfig.PackageImportPattern, false, scriptImport)
}

func (b *BuildScriptService) changeCrossBuild(script *Script.Script, isBefore bool) {
	os := b.os
	cgo := b.cgo
	arch := b.arch
	distPath := script.DistPath
	if !isBefore {
		os = b.buildConfig.OsOriginal
		cgo = b.buildConfig.CgoOriginal
		arch = b.buildConfig.ArchOriginal
		distPath = b.buildConfig.DistOriginal
	}
	FileUtil.ModContent(b.buildConfig.CrossBuildPath, b.buildConfig.OsPattern, false, os)
	FileUtil.ModContent(b.buildConfig.CrossBuildPath, b.buildConfig.ArchPattern, false, arch)
	FileUtil.ModContent(b.buildConfig.CrossBuildPath, b.buildConfig.DistPattern, false, distPath)
	FileUtil.ModContent(b.buildConfig.CrossBuildPath, b.buildConfig.CgoPattern, false, GenUtil.IntToString(cgo))
}

func Run() {
	newBuildScriptService().apply()
}
