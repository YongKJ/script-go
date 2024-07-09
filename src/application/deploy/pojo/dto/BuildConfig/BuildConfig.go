package BuildConfig

import (
	"script-go/src/application/util/DataUtil"
	"script-go/src/application/util/FileUtil"
)

type BuildConfig struct {
	appPath               string
	appTestPath           string
	scriptRunPattern      string
	scriptRunOriginal     string
	packageImportPattern  string
	packageImportOriginal string
	mapData               map[string][]string
}

func newBuildConfig(mapData map[string][]string, appPath string, appTestPath string, scriptRunPattern string, scriptRunOriginal string, packageImportPattern string, packageImportOriginal string) *BuildConfig {
	return &BuildConfig{appPath: appPath, appTestPath: appTestPath, scriptRunPattern: scriptRunPattern, scriptRunOriginal: scriptRunOriginal, packageImportPattern: packageImportPattern, packageImportOriginal: packageImportOriginal, mapData: mapData}
}

func Of(mapData map[string][]string, appPath string, appTestPath string, scriptRunPattern string, scriptRunOriginal string, packageImportPattern string, packageImportOriginal string) *BuildConfig {
	return newBuildConfig(mapData, appPath, appTestPath, scriptRunPattern, scriptRunOriginal, packageImportPattern, packageImportOriginal)
}

func Get() *BuildConfig {
	mapData := map[string][]string{
		"windows": {"amd64", "arm64", "arm", "386"},
		"linux":   {"amd64", "arm64", "arm", "386"},
		"android": {"amd64", "arm64", "arm", "386"},
		"darwin":  {"amd64", "arm64"},
		"ios":     {"amd64", "arm64"},
	}
	appPath := FileUtil.GetAbsPath("src", "application", "Application.go")
	appTestPath := FileUtil.GetAbsPath("src", "application", "ApplicationTest.go")
	return Of(
		mapData, appPath, appTestPath,
		"\\s+(\\S+)\\.Run\\(\\)", "Demo",
		".+\"(.+)\"", "script-go/src/application/applet/Demo",
	)
}

func JsonArrayToObjects(jsonArrayStr string) []*BuildConfig {
	return DataUtil.JsonArrayToObjects(jsonArrayStr, &BuildConfig{}).([]*BuildConfig)
}

func JsonToObject(jsonStr string) *BuildConfig {
	return DataUtil.JsonToObject(jsonStr, &BuildConfig{}).(*BuildConfig)
}

func ArrayToObjects(arrayData []map[string]any) []*BuildConfig {
	return DataUtil.ArrayToObjects(arrayData, &BuildConfig{}).([]*BuildConfig)
}

func MapToObject(mapData map[string]any) *BuildConfig {
	return DataUtil.MapToObject(mapData, &BuildConfig{}).(*BuildConfig)
}

func ObjectToMap(data *BuildConfig) map[string]any {
	return DataUtil.ObjectToMap(data)
}

func ObjectsToArray(lstData []*BuildConfig) []map[string]any {
	return DataUtil.ObjectsToArray(lstData)
}

func (b *BuildConfig) AppPath() string {
	return b.appPath
}

func (b *BuildConfig) SetAppPath(appPath string) {
	b.appPath = appPath
}

func (b *BuildConfig) AppTestPath() string {
	return b.appTestPath
}

func (b *BuildConfig) SetAppTestPath(appTestPath string) {
	b.appTestPath = appTestPath
}

func (b *BuildConfig) ScriptRunPattern() string {
	return b.scriptRunPattern
}

func (b *BuildConfig) SetScriptRunPattern(scriptRunPattern string) {
	b.scriptRunPattern = scriptRunPattern
}

func (b *BuildConfig) ScriptRunOriginal() string {
	return b.scriptRunOriginal
}

func (b *BuildConfig) SetScriptRunOriginal(scriptRunOriginal string) {
	b.scriptRunOriginal = scriptRunOriginal
}

func (b *BuildConfig) PackageImportPattern() string {
	return b.packageImportPattern
}

func (b *BuildConfig) SetPackageImportPattern(packageImportPattern string) {
	b.packageImportPattern = packageImportPattern
}

func (b *BuildConfig) PackageImportOriginal() string {
	return b.packageImportOriginal
}

func (b *BuildConfig) SetPackageImportOriginal(packageImportOriginal string) {
	b.packageImportOriginal = packageImportOriginal
}

func (b *BuildConfig) MapData() map[string][]string {
	return b.mapData
}

func (b *BuildConfig) SetMapData(mapData map[string][]string) {
	b.mapData = mapData
}
