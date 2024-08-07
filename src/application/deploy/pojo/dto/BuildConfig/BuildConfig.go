package BuildConfig

import (
	"script-go/src/application/util/DataUtil"
	"script-go/src/application/util/FileUtil"
)

type BuildConfig struct {
	SrcPath               string              `json:"srcPath"`
	AppPath               string              `json:"appPath"`
	AppTestPath           string              `json:"appTestPath"`
	CrossBuildPath        string              `json:"crossBuildPath"`
	MapOS                 map[string][]string `json:"mapOS"`
	CgoPattern            string              `json:"cgoPattern"`
	CgoOriginal           int                 `json:"cgoOriginal"`
	OsPattern             string              `json:"osPattern"`
	OsOriginal            string              `json:"osOriginal"`
	ArchPattern           string              `json:"archPattern"`
	ArchOriginal          string              `json:"archOriginal"`
	DistPattern           string              `json:"distPattern"`
	DistOriginal          string              `json:"distOriginal"`
	ScriptRunPattern      string              `json:"scriptRunPattern"`
	ScriptRunOriginal     string              `json:"scriptRunOriginal"`
	PackageImportPattern  string              `json:"packageImportPattern"`
	PackageImportOriginal string              `json:"packageImportOriginal"`
}

func newBuildConfig(srcPath string, appPath string, appTestPath string, crossBuildPath string, mapOS map[string][]string, cgoPattern string, cgoOriginal int, osPattern string, osOriginal string, archPattern string, archOriginal string, distPattern string, distOriginal string, scriptRunPattern string, scriptRunOriginal string, packageImportPattern string, packageImportOriginal string) *BuildConfig {
	return &BuildConfig{SrcPath: srcPath, AppPath: appPath, AppTestPath: appTestPath, CrossBuildPath: crossBuildPath, MapOS: mapOS, CgoPattern: cgoPattern, CgoOriginal: cgoOriginal, OsPattern: osPattern, OsOriginal: osOriginal, ArchPattern: archPattern, ArchOriginal: archOriginal, DistPattern: distPattern, DistOriginal: distOriginal, ScriptRunPattern: scriptRunPattern, ScriptRunOriginal: scriptRunOriginal, PackageImportPattern: packageImportPattern, PackageImportOriginal: packageImportOriginal}
}

func Of(srcPath string, appPath string, appTestPath string, crossBuildPath string, mapOS map[string][]string, cgoPattern string, cgoOriginal int, osPattern string, osOriginal string, archPattern string, archOriginal string, distPattern string, distOriginal string, scriptRunPattern string, scriptRunOriginal string, packageImportPattern string, packageImportOriginal string) *BuildConfig {
	return newBuildConfig(srcPath, appPath, appTestPath, crossBuildPath, mapOS, cgoPattern, cgoOriginal, osPattern, osOriginal, archPattern, archOriginal, distPattern, distOriginal, scriptRunPattern, scriptRunOriginal, packageImportPattern, packageImportOriginal)
}

func Get() *BuildConfig {
	mapOS := map[string][]string{
		"windows": {"amd64", "arm64", "arm", "386"},
		"linux":   {"amd64", "arm64", "arm", "386"},
		"android": {"amd64", "arm64", "arm", "386"},
		"darwin":  {"amd64", "arm64"},
		"ios":     {"amd64", "arm64"},
	}
	srcPath := FileUtil.GetAbsPath("src")
	crossBuildPath := FileUtil.GetAbsPath("cross_build.sh")
	appPath := FileUtil.GetAbsPath("src", "application", "Application.go")
	appTestPath := FileUtil.GetAbsPath("src", "application", "ApplicationTest.go")
	return Of(
		srcPath, appPath,
		appTestPath, crossBuildPath, mapOS,
		".+ENABLED=(.+)", 1,
		".+OS=(.+)", "windows",
		".+ARCH=(.+)", "amd64",
		".+-o\\s(.+)", "../dist/script-go.exe",
		"\\s+(\\S+)\\.Run\\(\\)", "Demo",
		".+\"(.+)\"", "script-go/src/application/applet/Demo",
	)
}

func GetMapOSKeys(mapData map[string][]string) []string {
	var lstKey []string
	for key := range mapData {
		lstKey = append(lstKey, key)
	}
	return lstKey
}

func JsonArrayToObjects(jsonArrayStr string) []*BuildConfig {
	var lstData []*BuildConfig
	return *(DataUtil.JsonArrayToObjects(jsonArrayStr, &lstData).(*[]*BuildConfig))
}

func JsonArrayToMaps(jsonArrayStr string) []map[string]any {
	return DataUtil.JsonArrayToMaps(jsonArrayStr)
}

func JsonToObject(jsonStr string) *BuildConfig {
	return DataUtil.JsonToObject(jsonStr, &BuildConfig{}).(*BuildConfig)
}

func JsonToMap(jsonStr string) map[string]any {
	return DataUtil.JsonToMap(jsonStr)
}

func ObjectsToJsonArray(lstData []*BuildConfig) string {
	return DataUtil.ObjectsToJsonArray(lstData)
}

func ObjectsToMaps(lstData []*BuildConfig) []map[string]any {
	return DataUtil.ObjectsToMaps(lstData)
}

func ObjectToJson(data *BuildConfig) string {
	return DataUtil.ObjectToJson(data)
}

func ObjectToMap(data *BuildConfig) map[string]any {
	return DataUtil.ObjectToMap(data)
}

func MapsToJsonArray(arrayData []map[string]any) string {
	return DataUtil.MapsToJsonArray(arrayData)
}

func MapsToObjects(arrayData []map[string]any) []*BuildConfig {
	var lstData []*BuildConfig
	return *(DataUtil.MapsToObjects(arrayData, &lstData).(*[]*BuildConfig))
}

func MapToJson(mapData map[string]any) string {
	return DataUtil.MapToJson(mapData)
}

func MapToObject(mapData map[string]any) *BuildConfig {
	return DataUtil.MapToObject(mapData, &BuildConfig{}).(*BuildConfig)
}
