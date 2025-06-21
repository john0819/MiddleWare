/*
	json
	序列化: 从结构体类型转换成json字符串 - 用户传输
	反序列化: 将json字符串转换成结构体类型 - 程序内部处理
*/

package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// mock user db
var userDB = map[int]User{
	1: {
		Name: "John",
		Age:  20,
	},
	2: {
		Name: "wong",
		Age:  21,
	},
	3: {
		Name: "lily",
		Age:  22,
	},
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id, err := strconv.Atoi(query.Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	user, ok := userDB[id]
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userDB[len(userDB)+1] = user
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func main() {

	// http demo
	http.HandleFunc("/user", getUserHandler)
	http.HandleFunc("/create", createUserHandler)
	http.ListenAndServe(":8080", nil)

	// 简单的demo
	/*
		// 序列化: 将结构体转换成json字符串
		user := User{
			Name: "John",
			Age:  20,
		}
		jsonUser, err := json.Marshal(user)
		if err != nil {
			fmt.Println("序列化失败:", err)
			return
		}
		fmt.Println("序列化成功:", string(jsonUser))

		// 反序列化: 将json转换成结构体
		jsonStr := `{"name": "wong","age":20}`
		var user2 User
		err = json.Unmarshal([]byte(jsonStr), &user2)
		if err != nil {
			fmt.Println("反序列化失败:", err)
			return
		}
		fmt.Println("反序列化成功:", user2)
	*/
}
