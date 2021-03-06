package main

import ("net/http"; "encoding/json"; "github.com/gorilla/mux"; "strconv")

var (table = make([]string, 0, 0))

type Answer struct{
	Url string `json:"url"`;
}


func postKey(w http.ResponseWriter, r *http.Request) {
	var answer Answer;
	json.NewDecoder(r.Body).Decode(&answer)
	ind := answer.Url
	table = append(table, ind);
	ans := make(map[string]string)
	ans["key"] = strconv.Itoa(len(table) - 1)
	x, err := json.Marshal(ans)
	if err != nil{
		panic(err)
	}
	w.Write(x)
}

func getKey(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, err := strconv.Atoi(vars["key"])
	if err != nil {
		panic(err)
	}
	http.Redirect(w, r, table[key], 301)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", postKey).Methods("POST")
	router.HandleFunc("/{key}", getKey)
	http.ListenAndServe(":8082", router)
}


/*Сервис должен слушать порт 8082.
Метод ‘/’ должен принимать POST-запросы с json в body с единственным полем ‘url’. Этот url должен получить уникальный короткий ключ, который должен быть возвращен в body в json с единственным полем ‘key’.
Метод ‘/{key}’ (например, /u42) должен возвращать ответ со статусом Moved Permanently и заголовком Location равным тому URL, который был дабавлен POST-методом.
*/