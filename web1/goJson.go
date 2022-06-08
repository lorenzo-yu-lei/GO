package main

import (
	"encoding/json"
	"log"
	"net/http"
)

//类型映射
//go bool 	: JSON boolean
//go float64 	: JSON 数值
//go string 	: JSON strings
//go nil 		: JSON null

//对于未知结构的JSON
//map[string]interface{} 		可以存储任意JSON对象
//[]interface{}				可以存储任意的JSON数组

//读取JSON
//需要一个解码器： 	dec:= json.NewDecoder(r.Body)
//在解码器上进行解码：	dec.Decode(&query)

//写入JSON
//需要一个编码器：		enc := json.NewEncoder(w)
//编码：				enc.Encode(results)

func companies() {
	http.HandleFunc("/companies", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			dec := json.NewDecoder(r.Body)
			company := Company{}
			err := dec.Decode(&company)
			if err != nil {
				log.Panicln(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			enc := json.NewEncoder(w)
			err = enc.Encode(company)
			if err != nil {
				log.Println(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

}

func main() {
	companies()
}
