package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 设置跨域请求头
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method Not Allowed")
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Bad Request")
		return
	}
	defer file.Close()

	// minio文件上传
	client, err := NewClient()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Internal Server Error")
		return
	}

	bucketName := "test"
	filePath := Upload(client, r.Context(), bucketName, handler.Filename, file, handler.Size)

	json, _ := json.Marshal(map[string]string{"filePath": filePath})
	w.Write(json)

	//fmt.Fprintf(w, "File uploaded successfully: %s", handler.Filename)
}

func main() {
	// 创建一个路由器
	mux := http.NewServeMux()

	// 使用中间件处理跨域
	mux.Handle("/upload", withCORS(http.HandlerFunc(uploadHandler)))

	port := 10000

	// 设置监听端口
	fmt.Printf("Server listening on port %d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
}
