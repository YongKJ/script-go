package BuildConfig

import (
	"encoding/json"
	"log"
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
	var arrayData []map[string]any
	err := json.Unmarshal([]byte(jsonArrayStr), &arrayData)
	if err != nil {
		log.Println(err)
	}
	return ArrayToObjects(arrayData)
}

func JsonToObject(jsonStr string) *BuildConfig {
	var mapData map[string]any
	err := json.Unmarshal([]byte(jsonStr), &mapData)
	if err != nil {
		log.Println(err)
	}
	return MapToObject(mapData)
}

func ArrayToObjects(arrayData []map[string]any) []*BuildConfig {
	length := len(arrayData)
	lstData := make([]*BuildConfig, length)
	for i := 0; i < length; i++ {
		lstData[i] = MapToObject(arrayData[i])
	}
	return lstData
}

func MapToObject(mapData map[string]any) *BuildConfig {
	return DataUtil.MapToObject(mapData, &BuildConfig{}).(*BuildConfig)
}

func ObjectToMap(buildConfig *BuildConfig) map[string]any {
	return DataUtil.ObjectToMap(buildConfig)
}

func ObjectsToArray(buildConfigs []*BuildConfig) []map[string]any {
	length := len(buildConfigs)
	lstData := make([]map[string]any, length)
	for i := 0; i < length; i++ {
		lstData[i] = ObjectToMap(buildConfigs[i])
	}
	return lstData
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
