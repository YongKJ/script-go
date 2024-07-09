package BuildConfig

import (
	"script-go/src/application/util/DataUtil"
)

type BuildConfig struct {
	appPath               string
	lstGoOS               []string
	lstGoArch             []string
	scriptRunPattern      string
	scriptRunOriginal     string
	packageImportPattern  string
	packageImportOriginal string
}

func newBuildConfig(appPath string, lstGoOS []string, lstGoArch []string, scriptRunPattern string, scriptRunOriginal string, packageImportPattern string, packageImportOriginal string) *BuildConfig {
	return &BuildConfig{appPath: appPath, lstGoOS: lstGoOS, lstGoArch: lstGoArch, scriptRunPattern: scriptRunPattern, scriptRunOriginal: scriptRunOriginal, packageImportPattern: packageImportPattern, packageImportOriginal: packageImportOriginal}
}

func Of(appPath string, lstGoOS []string, lstGoArch []string, scriptRunPattern string, scriptRunOriginal string, packageImportPattern string, packageImportOriginal string) *BuildConfig {
	return newBuildConfig(appPath, lstGoOS, lstGoArch, scriptRunPattern, scriptRunOriginal, packageImportPattern, packageImportOriginal)
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

func (b *BuildConfig) LstGoOS() []string {
	return b.lstGoOS
}

func (b *BuildConfig) SetLstGoOS(lstGoOS []string) {
	b.lstGoOS = lstGoOS
}

func (b *BuildConfig) LstGoArch() []string {
	return b.lstGoArch
}

func (b *BuildConfig) SetLstGoArch(lstGoArch []string) {
	b.lstGoArch = lstGoArch
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
