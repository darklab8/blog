package logus_core

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"runtime"
)

func GetCallingFile(level int) string {
	GetTwiceParentFunctionLocation := level
	_, filename, _, _ := runtime.Caller(GetTwiceParentFunctionLocation)
	filename = filepath.Base(filename)
	return fmt.Sprintf("f:%s ", filename)
}

func StructToMap(somestruct any) map[string]any {
	var mapresult map[string]interface{}
	inrec, _ := json.Marshal(somestruct)
	json.Unmarshal(inrec, &mapresult)
	return mapresult
}
