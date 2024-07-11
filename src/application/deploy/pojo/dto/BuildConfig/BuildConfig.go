package BuildConfig

import (
	"script-go/src/application/util/DataUtil"
	"script-go/src/application/util/FileUtil"
)

type BuildConfig struct {
	srcPath               string
	appPath               string
	appTestPath           string
	crossBuildPath        string
	mapOS                 map[string][]string
	cgoPattern            string
	cgoOriginal           int
	osPattern             string
	osOriginal            string
	archPattern           string
	archOriginal          string
	distPattern           string
	distOriginal          string
	scriptRunPattern      string
	scriptRunOriginal     string
	packageImportPattern  string
	packageImportOriginal string
}

func newBuildConfig(srcPath string, appPath string, appTestPath string, crossBuildPath string, mapOS map[string][]string, cgoPattern string, cgoOriginal int, osPattern string, osOriginal string, archPattern string, archOriginal string, distPattern string, distOriginal string, scriptRunPattern string, scriptRunOriginal string, packageImportPattern string, packageImportOriginal string) *BuildConfig {
	return &BuildConfig{srcPath: srcPath, appPath: appPath, appTestPath: appTestPath, crossBuildPath: crossBuildPath, mapOS: mapOS, cgoPattern: cgoPattern, cgoOriginal: cgoOriginal, osPattern: osPattern, osOriginal: osOriginal, archPattern: archPattern, archOriginal: archOriginal, distPattern: distPattern, distOriginal: distOriginal, scriptRunPattern: scriptRunPattern, scriptRunOriginal: scriptRunOriginal, packageImportPattern: packageImportPattern, packageImportOriginal: packageImportOriginal}
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
		".+-o\\s(.+)", "./dist/script-go.exe",
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

func (b *BuildConfig) SrcPath() string {
	return b.srcPath
}

func (b *BuildConfig) SetSrcPath(srcPath string) {
	b.srcPath = srcPath
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

func (b *BuildConfig) MapOS() map[string][]string {
	return b.mapOS
}

func (b *BuildConfig) SetMapOS(mapOS map[string][]string) {
	b.mapOS = mapOS
}

func (b *BuildConfig) CrossBuildPath() string {
	return b.crossBuildPath
}

func (b *BuildConfig) SetCrossBuildPath(crossBuildPath string) {
	b.crossBuildPath = crossBuildPath
}

func (b *BuildConfig) CgoPattern() string {
	return b.cgoPattern
}

func (b *BuildConfig) SetCgoPattern(cgoPattern string) {
	b.cgoPattern = cgoPattern
}

func (b *BuildConfig) CgoOriginal() int {
	return b.cgoOriginal
}

func (b *BuildConfig) SetCgoOriginal(cgoOriginal int) {
	b.cgoOriginal = cgoOriginal
}

func (b *BuildConfig) OsPattern() string {
	return b.osPattern
}

func (b *BuildConfig) SetOsPattern(osPattern string) {
	b.osPattern = osPattern
}

func (b *BuildConfig) OsOriginal() string {
	return b.osOriginal
}

func (b *BuildConfig) SetOsOriginal(osOriginal string) {
	b.osOriginal = osOriginal
}

func (b *BuildConfig) ArchPattern() string {
	return b.archPattern
}

func (b *BuildConfig) SetArchPattern(archPattern string) {
	b.archPattern = archPattern
}

func (b *BuildConfig) ArchOriginal() string {
	return b.archOriginal
}

func (b *BuildConfig) SetArchOriginal(archOriginal string) {
	b.archOriginal = archOriginal
}

func (b *BuildConfig) DistPattern() string {
	return b.distPattern
}

func (b *BuildConfig) SetDistPattern(distPattern string) {
	b.distPattern = distPattern
}

func (b *BuildConfig) DistOriginal() string {
	return b.distOriginal
}

func (b *BuildConfig) SetDistOriginal(distOriginal string) {
	b.distOriginal = distOriginal
}
