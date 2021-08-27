package iso

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

/**
 * @Author: mf
 * @email: 18539271635@163.com
 * @Date: 2021/8/27 2:21 下午
 * @Desc:
 */

//[{"name":"Afghanistan","alpha-2":"AF","country-code":"004"}
type Slim2 struct {
	Name        string `json:"name"`
	Alpha2      string `json:"alpha-2"`
	CountryCode string `json:"country-code"`
}

func GetISOMapFromFile() {
	byteVal, err := readFile()
	if err != nil {
		return
	}

	var slim []Slim2
	if err = json.Unmarshal(byteVal, &slim); err != nil {
		fmt.Printf("json格式化失败 [Err:%s]", err.Error())
		return
	}

	byteSlim, err := json.Marshal(formatSlimToMap(slim))
	if err != nil {
		fmt.Printf("json格式化失败 [Err:%s]", err.Error())
		return
	}

	fmt.Printf("slimp=%s", string(byteSlim))
}

func formatSlimToMap(slim []Slim2) map[string]string {
	slimMap :=make(map[string]string, len(slim))
	for _, v := range slim {
		slimMap[v.Alpha2] = v.Name
	}
	return slimMap
}

func readFile() ([]byte, error) {
	filePtr, err := os.Open("/Users/mengfei/workspace/goutil/GolangUtilCode/iso/slim-2/slim-2.json")
	if err != nil {
		fmt.Printf("文件打开失败 [Err:%s]", err.Error())
		return nil, err
	}
	defer filePtr.Close()
	byteVal, err := ioutil.ReadAll(filePtr)
	if err != nil {
		fmt.Printf("读取文件失败 [Err:%s]", err.Error())
		return nil, err
	}
	return byteVal, nil
}
