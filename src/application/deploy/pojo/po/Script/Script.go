package Script

import (
	"path/filepath"
	"script-go/src/application/util/DataUtil"
	"script-go/src/application/util/FileUtil"
	"script-go/src/application/util/GenUtil"
	"strings"
)

type Script struct {
	GoName        string `json:"goName"`
	GoPath        string `json:"goPath"`
	YamlConfig    string `json:"yamlConfig"`
	ScriptName    string `json:"scriptName"`
	ScriptPath    string `json:"scriptPath"`
	ScriptConfig  string `json:"scriptConfig"`
	ScriptRun     string `json:"scriptRun"`
	ScriptImport  string `json:"scriptImport"`
	ScriptProject string `json:"scriptProject"`
	DistPath      string `json:"distPath"`
}

func newScript(goName string, goPath string, yamlConfig string, scriptName string, scriptPath string, scriptConfig string, scriptRun string, scriptImport string, scriptProject string, distPath string) *Script {
	return &Script{GoName: goName, GoPath: goPath, YamlConfig: yamlConfig, ScriptName: scriptName, ScriptPath: scriptPath, ScriptConfig: scriptConfig, ScriptRun: scriptRun, ScriptImport: scriptImport, ScriptProject: scriptProject, DistPath: distPath}
}

func Of(goName string, goPath string, yamlConfig string, scriptName string, scriptPath string, scriptConfig string, scriptRun string, scriptImport string, scriptProject string, distPath string) *Script {
	return newScript(goName, goPath, yamlConfig, scriptName, scriptPath, scriptConfig, scriptRun, scriptImport, scriptProject, distPath)
}

func Gets() []*Script {
	path := FileUtil.GetAbsPath("src", "application", "deploy", "service")
	lstScript := GetListByDir("")
	return append(lstScript, GetListByDir(path)...)
}

func SetDistPath(script *Script, os string, arch string) {
	distPath := script.DistPath
	scriptPath := script.ScriptPath
	if !(os == "windows" && arch == "amd64") {
		distPath = strings.Join([]string{distPath, os, arch}, "-")
		scriptPath = strings.Join([]string{scriptPath, os, arch}, "-")
	}
	if os == "windows" {
		distPath = distPath + ".exe"
		scriptPath = scriptPath + ".exe"
	}
	script.DistPath = distPath
	script.ScriptPath = scriptPath
}

func GetListByDir(appletDir string) []*Script {
	if len(appletDir) == 0 {
		appletDir = FileUtil.GetAbsPath("src", "application", "applet")
	}
	assetsDir := FileUtil.GetAbsPath("src", "assets")
	scriptDir := FileUtil.GetAbsPath("script")
	distDir := FileUtil.GetAbsPath("dist")
	lstFile := FileUtil.List(appletDir)

	lstScript := make([]*Script, len(lstFile))
	for i, file := range lstFile {
		goPath := filepath.Join(appletDir, file)
		if FileUtil.IsFolder(goPath) {
			goPath = getScript(goPath)
		}
		index := strings.LastIndex(goPath, string(filepath.Separator))
		goName := goPath[index+1:]
		name, _ := strings.CutSuffix(goName, ".go")

		scriptRun := name
		scriptName := GenUtil.ToLine(name)
		projectName := GenUtil.ToLine(name)
		scriptImport := getImportPath(goPath)
		yamlName := GenUtil.ToLine(name) + ".yaml"
		distPath := filepath.Join(distDir, scriptName)
		yamlConfig := filepath.Join(assetsDir, yamlName)
		scriptProject := filepath.Join(scriptDir, projectName)
		scriptConfig := filepath.Join(scriptDir, projectName, yamlName)
		scriptPath := filepath.Join(scriptDir, projectName, scriptName)
		distPath = strings.Replace(distPath, string(filepath.Separator), "/", -1)

		lstScript[i] = Of(
			goName, goPath, yamlConfig, scriptName, scriptPath,
			scriptConfig, scriptRun, scriptImport, scriptProject, distPath,
		)
	}
	return lstScript
}

func getImportPath(path string) string {
	path = filepath.Dir(path)
	path = strings.Split(path, "script-go")[1]
	return "script-go" + strings.Replace(path, string(filepath.Separator), "/", -1)
}

func getScript(folder string) string {
	lstFile := FileUtil.List(folder)
	for _, file := range lstFile {
		if strings.HasSuffix(file, ".go") {
			return filepath.Join(folder, file)
		}
	}
	return ""
}

func JsonArrayToObjects(jsonArrayStr string) []*Script {
	var lstData []*Script
	return *(DataUtil.JsonArrayToObjects(jsonArrayStr, &lstData).(*[]*Script))
}

func JsonArrayToMaps(jsonArrayStr string) []map[string]any {
	return DataUtil.JsonArrayToMaps(jsonArrayStr)
}

func JsonToObject(jsonStr string) *Script {
	return DataUtil.JsonToObject(jsonStr, &Script{}).(*Script)
}

func JsonToMap(jsonStr string) map[string]any {
	return DataUtil.JsonToMap(jsonStr)
}

func ObjectsToJsonArray(lstData []*Script) string {
	return DataUtil.ObjectsToJsonArray(lstData)
}

func ObjectsToMaps(lstData []*Script) []map[string]any {
	return DataUtil.ObjectsToMaps(lstData)
}

func ObjectToJson(data *Script) string {
	return DataUtil.ObjectToJson(data)
}

func ObjectToMap(data *Script) map[string]any {
	return DataUtil.ObjectToMap(data)
}

func MapsToJsonArray(arrayData []map[string]any) string {
	return DataUtil.MapsToJsonArray(arrayData)
}

func MapsToObjects(arrayData []map[string]any) []*Script {
	var lstData []*Script
	return *(DataUtil.MapsToObjects(arrayData, &lstData).(*[]*Script))
}

func MapToJson(mapData map[string]any) string {
	return DataUtil.MapToJson(mapData)
}

func MapToObject(mapData map[string]any) *Script {
	return DataUtil.MapToObject(mapData, &Script{}).(*Script)
}
