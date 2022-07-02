package json

type CallVideoServerRequest struct {
	AppID             string            `json:"app_id"`
	UserID            string            `json:"user_id"`
	DeviceID          string            `json:"device_id"`
	DeviceType        int32             `json:"device_type"`
	UserModelPlatform int32             `json:"user_model_platform"`
	TrafficModelType  int32             `json:"traffic_model_type"`
	Profiles          map[string]string `json:"profiles"`
	ModelServerParams map[string]string `json:"p"`
}

//func main() {
//	// 测试用户画像平台
//	p := map[string]string{
//		"android_sdk_int":     "30",
//		"app_version":         "8.4.90.26446",
//		"app_version_build":   "26446",
//		"app_version_short":   "8.4.90",
//		"device_abis":         "'arm64-v8a','armeabi-v7a','armeabi'",
//		"device_brand":        "HONOR",
//		"device_dpi":          "480",
//		"device_manufacturer": "HONOR",
//		"device_model":        "NTH-AN00",
//	}
//	model := map[string]string{}
//	c := &CallVideoServerRequest{
//		AppID:             "5608",
//		UserID:            "f3bb55b9cd362d478c3a5672b2421d2697210010216601",
//		DeviceID:          "",
//		DeviceType:        2,
//		UserModelPlatform: 101,
//		TrafficModelType:  101,
//		Profiles:          p,
//		ModelServerParams: model,
//	}
//	m, _ := json.Marshal(c)
//	fmt.Println(string(m) + "\n")
//
//	// 测试用户模型平台
//	p1 := map[string]string{
//		"app_version_build": "24470",
//		"device_model":      "iPhone 11",
//		"ios_sys_version":   "14.7.1",
//		"app_version_short": "8.5.41",
//		"app_version":       "8.5.41.24470",
//	}
//	model1 := map[string]string{
//		"omgid": "553c38dabff83046ffc9c11a01818f5cad030010116a19",
//		"vuid":  "429370703",
//	}
//	c1 := &CallVideoServerRequest{
//		AppID:             "5608",
//		UserID:            "553c38dabff83046ffc9c11a01818f5cad030010116a19",
//		DeviceID:          "",
//		DeviceType:        101,
//		UserModelPlatform: 100,
//		TrafficModelType:  100,
//		Profiles:          p1,
//		ModelServerParams: model1,
//	}
//	m1, _ := json.Marshal(c1)
//	fmt.Println(string(m1) + "\n")
//
//	str := "{\"name\":\"faustzhao\",\"age\":\"12\"}"
//	fmt.Println(len(str))
//
//	marshal, err := json.Marshal(nil)
//	if err != nil {
//		return
//	}
//	fmt.Println(string(marshal))
//}

//type syncRequest struct {
//	ProjectID  string `json:"project_id"`   // 项目id
//	BizNodeID  string `json:"biz_node_id"`  // sre节点id
//	DstBizPath string `json:"dst_biz_path"` // 使用方业务树路径
//	DstOrgPath string `json:"dst_org_path"` // 使用方组织树路径
//}
//
//func main() {
//	s := &syncRequest{}
//	marshal, _ := json.Marshal(s)
//	fmt.Printf("get s: %v", string(marshal))
//}
