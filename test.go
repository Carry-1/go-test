func TestRegex(t *testing.T) {
	format := "redis[(\\-sentinel\\-master)|(\\-sentinel\\-replica\\-1)|(\\-sentinel\\-replica\\-2)]*:[0-9]+"
	ex := "redis-sentinel-master:17000"
	re := regexp.MustCompile(format)
	output := re.MatchString(ex)
	assert.Equal(t, true, output)
}

func TestPrintAsPrettyString(t *testing.T) {
	data, _ := os.ReadFile("/home/xiaqingsheng/PRA/Links.json")
	var parsed map[string]interface{}
	json.Unmarshal(data, &parsed)
	jsonValue, _ := json.MarshalIndent(parsed, "", "  ")

	f, _ := os.Create("../../linksPrettyPrint.json")

	defer f.Close()
	f.WriteString(string(jsonValue))
}


// get directory
func TestGetDir() {
    ex, err := os.Executable()
    if err != nil {
        panic(err)
    }
    exPath := filepath.Dir(ex)
    fmt.Println(exPath)
}