package zcnumber

import (
	"encoding/json"
	"strconv"
)

type JSONInt64 int64

func (j *JSONInt64) UnmarshalJSON(data []byte) error {
	// 尝试直接将JSON数据解析为int64类型
	var intVal int64
	if err := json.Unmarshal(data, &intVal); err == nil {
		*j = JSONInt64(intVal)
		return nil
	}

	// 如果上面的直接解析失败，尝试将JSON数据解析为字符串，然后转换为int64
	var strVal string
	if err := json.Unmarshal(data, &strVal); err != nil {
		return err
	}
	intVal, err := strconv.ParseInt(strVal, 10, 64)
	if err != nil {
		return err
	}

	*j = JSONInt64(intVal)
	return nil
}

func (j *JSONInt64) MarshalJSON() ([]byte, error) {
	return json.Marshal(int64(*j))
}
