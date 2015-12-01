# json map string byte[]格式互转

 这里使用的原生json库

```
import (
    "encoding/json"
    )

func main(){ 
    jsonStr :="{
        "pub": "2013-06-29 22:59",
        "name": "南宁",
        "wind": {
            "chill": 81,
            "direction": 140,
            "speed": 7
        },
        "astronomy": {
            "sunrise": "6:05",
            "sunset": "19:34"
        },
        "atmosphere": {
            "humidity": 89,
            "visibility": 6.21,
            "pressure": 29.71,
            "rising": 1
        }}"
        
    var dat map[string]interface{}
    if err := json.Unmarshal([]byte(jsonStr), &dat); err == nil {
        fmt.Println("==============json str转map=================")
        fmt.Println(dat)
    }
}
```