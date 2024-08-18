package helpers

func ArrayMap[T any, R any](raws []T, cb func(r T) R) []R {
	items := make([]R, 0)

	for _, r := range raws {
		items = append(items, cb(r))
	}

	return items
}

// func ReadFromFileMaps(data []byte, options *core.ICSVOptions) ([]map[string]interface{}, error) {
// 	reader := bytes.NewReader(data)
// 	csvRead := csvs.NewReader(reader)
// 	csvRead.LazyQuotes = true
// 	csvData, err := csvRead.ReadAll()
// 	if err != nil {
// 		return nil, err
// 	}

// 	dataList := []map[string]interface{}{}
// 	var dataLineOne []string

// 	for i, line := range csvData {
// 		if i == 0 {
// 			dataLineOne = line
// 		} else {
// 			data := map[string]interface{}{}
// 			for i2, line2 := range dataLineOne {
// 				clean := zerowidth.RemoveZeroWidthCharacters(line2)
// 				data[strings.ReplaceAll(clean, "\"", "")] = line[i2]
// 			}
// 			dataList = append(dataList, data)
// 		}
// 	}

// 	return dataList, nil
// }
