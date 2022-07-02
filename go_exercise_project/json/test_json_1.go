package json

type RSAKey struct {
	Keypair map[string]map[string]string `yaml:"keypair"`
}

type TestStruct struct {
	Name string
	Age  int32
}

//func main() {
//	t1 := &TestStruct{
//		Name: "123",
//		Age:  12,
//	}
//	marshal, _ := json.Marshal(t1)
//	fmt.Println(string(marshal))
//
//}
