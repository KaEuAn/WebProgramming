package main

import ("net/http"; "json"; "github.com/gorilla/mux")

table = make([]string, 0, 0)


func postKey(w http.ResponseWriter, r *http.Request) {
	json.NewDecoder(r).Decode(answer)
	ind := answer[url]
	table = append(table, ind);
	w.Write(len(table))
}

func getKey(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.MovedPermanently)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", postKey).Methods("POST")
	router.HandleFunc("/{key}", getKey)
	router.ListenAndServe(":8082")
}


/*Сервис должен слушать порт 8082.
Метод ‘/’ должен принимать POST-запросы с json в body с единственным полем ‘url’. Этот url должен получить уникальный короткий ключ, который должен быть возвращен в body в json с единственным полем ‘key’.
Метод ‘/{key}’ (например, /u42) должен возвращать ответ со статусом Moved Permanently и заголовком Location равным тому URL, который был дабавлен POST-методом.
*/