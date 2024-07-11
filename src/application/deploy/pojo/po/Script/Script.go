package Script

import (
	"path/filepath"
	"runtime"
	"script-go/src/application/util/DataUtil"
	"script-go/src/application/util/FileUtil"
	"script-go/src/application/util/GenUtil"
	"strings"
)

type Script struct {
	goName       string
	goPath       string
	yamlConfig   string
	scriptName   string
	scriptPath   string
	scriptConfig string
	scriptRun    string
	scriptImport string
	distPath     string
}

func newScript(goName string, goPath string, yamlConfig string, scriptName string, scriptPath string, scriptConfig string, scriptRun string, scriptImport string, distPath string) *Script {
	return &Script{goName: goName, goPath: goPath, yamlConfig: yamlConfig, scriptName: scriptName, scriptPath: scriptPath, scriptConfig: scriptConfig, scriptRun: scriptRun, scriptImport: scriptImport, distPath: distPath}
}

func Of(goName string, goPath string, yamlConfig string, scriptName string, scriptPath string, scriptConfig string, scriptRun string, scriptImport string, distPath string) *Script {
	return newScript(goName, goPath, yamlConfig, scriptName, scriptPath, scriptConfig, scriptRun, scriptImport, distPath)
}

func Gets() []*Script {
	path := FileUtil.GetAbsPath("src", "application", "deploy", "service")
	lstScript := GetListByDir("")
	return append(lstScript, GetListByDir(path)...)
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

		suffix := ""
		name, _ := strings.CutSuffix(goName, ".go")
		if runtime.GOOS == "windows" {
			suffix = ".exe"
		}
		scriptRun := name
		projectName := GenUtil.ToLine(name)
		scriptImport := getImportPath(goPath)
		yamlName := GenUtil.ToLine(name) + ".yaml"
		scriptName := GenUtil.ToLine(name) + suffix
		distPath := filepath.Join(distDir, scriptName)
		yamlConfig := filepath.Join(assetsDir, yamlName)
		scriptConfig := filepath.Join(scriptDir, projectName, yamlName)
		scriptPath := filepath.Join(scriptDir, projectName, scriptName)
		scriptPath = strings.Replace(scriptPath, string(filepath.Separator), "/", -1)

		lstScript[i] = Of(
			goName, goPath, yamlConfig, scriptName, scriptPath,
			scriptConfig, scriptRun, scriptImport, distPath,
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
	return DataUtil.JsonArrayToObjects(jsonArrayStr, &Script{}).([]*Script)
}

func JsonToObject(jsonStr string) *Script {
	return DataUtil.JsonToObject(jsonStr, &Script{}).(*Script)
}

func ArrayToObjects(arrayData []map[string]any) []*Script {
	return DataUtil.ArrayToObjects(arrayData, &Script{}).([]*Script)
}

func MapToObject(mapData map[string]any) *Script {
	return DataUtil.MapToObject(mapData, &Script{}).(*Script)
}

func ObjectToMap(data *Script) map[string]any {
	return DataUtil.ObjectToMap(data)
}

func ObjectsToArray(lstData []*Script) []map[string]any {
	return DataUtil.ObjectsToArray(lstData)
}

func (s *Script) GoName() string {
	return s.goName
}

func (s *Script) SetGoName(goName string) {
	s.goName = goName
}

func (s *Script) GoPath() string {
	return s.goPath
}

func (s *Script) SetGoPath(goPath string) {
	s.goPath = goPath
}

func (s *Script) YamlConfig() string {
	return s.yamlConfig
}

func (s *Script) SetYamlConfig(yamlConfig string) {
	s.yamlConfig = yamlConfig
}

func (s *Script) ScriptName() string {
	return s.scriptName
}

func (s *Script) SetScriptName(scriptName string) {
	s.scriptName = scriptName
}

func (s *Script) ScriptPath() string {
	return s.scriptPath
}

func (s *Script) SetScriptPath(scriptPath string) {
	s.scriptPath = scriptPath
}

func (s *Script) ScriptConfig() string {
	return s.scriptConfig
}

func (s *Script) SetScriptConfig(scriptConfig string) {
	s.scriptConfig = scriptConfig
}

func (s *Script) ScriptImport() string {
	return s.scriptImport
}

func (s *Script) SetScriptImport(scriptImport string) {
	s.scriptImport = scriptImport
}

func (s *Script) ScriptRun() string {
	return s.scriptRun
}

func (s *Script) SetScriptRun(scriptRun string) {
	s.scriptRun = scriptRun
}

func (s *Script) DistPath() string {
	return s.distPath
}

func (s *Script) SetDistPath(distPath string) {
	s.distPath = distPath
}
