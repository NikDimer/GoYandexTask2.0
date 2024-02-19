package main

import(
	"net/http"
	"encoding/json"
	"math/rand"
	"time"
	"strconv"
	"io/ioutil"
	"os"
	//"sync"
	//"strings"
)
type SendResponseData struct {
	Id  string `json:"id"`
}

type Task struct {
	a int
	b int
	d string
}

type Expression struct {
	processId string
	expression string
}

func main() {
	http.HandleFunc("/send", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if (r.Method == "POST") {
			defer r.Body.Close()
			exp := make(map[string]string)
			json.NewDecoder(r.Body).Decode(&exp)
	
			rand.Seed(time.Now().UnixNano())
			newId := strconv.Itoa(rand.Intn(1000000000))
	
			responseData := SendResponseData{
				Id: "Идентификатор: " + newId,
			}
	
			responseBytes, err := json.Marshal(responseData)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
	
			/*jsonFile, err1 := os.OpenFile("./expressions.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0660)
			if err1 != nil {
				http.Error(w, err1.Error(), http.StatusInternalServerError)
			}
			defer jsonFile.Close()*/
			os.Chmod("expressions.json", 0666)
			data, err2 := ioutil.ReadFile("expressions.json")
			if err2 != nil {
				http.Error(w, err2.Error(), http.StatusInternalServerError)
			}
	
			m := make(map[string]map[string]string)
			d := make(map[string]string)
			d["id"] = newId
			d["expression"] = exp["expression"]
			d["status"] = "Waiting..."
			d["result"] = ""
			d["creationdate"] = time.Now().Format("2006-01-02 15:04:05")
			d["completedate"] = ""
			err3 := json.Unmarshal(data, &m)
			if err3 != nil {
				http.Error(w, err3.Error(), http.StatusInternalServerError)
			}
			m[newId] = d
	
			newData, err4 := json.Marshal(m)
			if err4 != nil {
				http.Error(w, err4.Error(), http.StatusInternalServerError)
			}
	
			err5 := ioutil.WriteFile("expressions.json", newData, 0)
			if err5 != nil {
				http.Error(w, err5.Error(), http.StatusInternalServerError)
			}

			//ActiveExpressions = append(ActiveExpressions, Expression{processId: id})
			
			w.Header().Set("Content-Type", "application/json")
			  w.Write(responseBytes)
		}
	})
	
		http.HandleFunc("/expressions", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			
			if (r.Method == "GET") {
				os.Chmod("expressions.json", 0644)
				data, err := ioutil.ReadFile("expressions.json")
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
		
				w.Header().Set("Content-Type", "application/json")
				w.Write(data)
			}
	})

	http.HandleFunc("/operations", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		
		if (r.Method == "GET") {
			os.Chmod("operations.json", 0644)
			data, err := ioutil.ReadFile("operations.json")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
	
			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
		}

		if (r.Method == "POST") {
			defer r.Body.Close()
			reqData := make(map[string]string)
			json.NewDecoder(r.Body).Decode(&reqData)
			answer := make(map[string]string)
			answer["ans"] = "Успешно сохранено"
	
			responseBytes, err := json.Marshal(answer)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
	
			/*jsonFile, err1 := os.OpenFile("./expressions.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0660)
			if err1 != nil {
				http.Error(w, err1.Error(), http.StatusInternalServerError)
			}
			defer jsonFile.Close()*/
			os.Chmod("operations.json", 0666)
			data, err2 := ioutil.ReadFile("operations.json")
			if err2 != nil {
				http.Error(w, err2.Error(), http.StatusInternalServerError)
			}
	
			m := make(map[string]string)
			err3 := json.Unmarshal(data, &m)
			if err3 != nil {
				http.Error(w, err3.Error(), http.StatusInternalServerError)
			}
			m["plus"] = reqData["plus"]
			m["minus"] = reqData["minus"]
			m["mn"] = reqData["mn"]
			m["del"] = reqData["del"]
	
			newData, err4 := json.Marshal(m)
			if err4 != nil {
				http.Error(w, err4.Error(), http.StatusInternalServerError)
			}
	
			err5 := ioutil.WriteFile("operations.json", newData, 0)
			if err5 != nil {
				http.Error(w, err5.Error(), http.StatusInternalServerError)
			}
			
			w.Header().Set("Content-Type", "application/json")
			  w.Write(responseBytes)
		}
	})

	http.HandleFunc("/power", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		
		if (r.Method == "GET") {
			os.Chmod("power.json", 0644)
			data, err := ioutil.ReadFile("power.json")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
	
			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
		}
	})

	http.HandleFunc("/options", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.WriteHeader(http.StatusNoContent)
		return
	})
	
	http.ListenAndServe(":8080", nil)

	//Daemon
	/*
	mu := &sync.Mutex{}

	func Worker(in chan Task, out chan int) {
		for t := <-in {
			mu.Lock()
			os.Chmod("operations.json", 0666)
			data, err2 := ioutil.ReadFile("operations.json")
			if err2 != nil {
				http.Error(w, err2.Error(), http.StatusInternalServerError)
			}
	
			m := make(map[string]string)
			err3 := json.Unmarshal(data, &m)
			if err3 != nil {
				http.Error(w, err3.Error(), http.StatusInternalServerError)
			}
			mu.Unlock()
			result := 0
			if d == "+" {
				result = a + b
				time.Sleep(m["plus"])
			} else if d == "-" {
				result = a - b
				time.Sleep(m["minus"])
			} else if d == "*" {
				result = a * b
				time.Sleep(m["mn"])
			} else {
				result = a / b
				time.Sleep(m["del"])
			}
			out <- result
			defer close(out)
		}
	}

	func parser(in chan Expression, out chan Task) {
		for s := <-in {
			mu.Lock()
			os.Chmod("expressions.json", 0666)
			data, err2 := ioutil.ReadFile("expressions.json")
			if err2 != nil {
				http.Error(w, err2.Error(), http.StatusInternalServerError)
			}
	
			m := make(map[string]map[string]string)
			err3 := json.Unmarshal(data, &m)
			if err3 != nil {
				http.Error(w, err3.Error(), http.StatusInternalServerError)
			}
			mu.Unlock()
			if (m[expression.processId]["status"] == "Waiting...") {
				m[expression.processId]["status"] = "Processing..."
			}
			var h []string
			for (k := 0; k < len(s.expression); k++) {
				if len(h) % 2 == 0 && strings.Contains("+-*/ /*", s.expression[k]) {
					h = append(h, s.expression[k])
				} else if len(h) % 2 == 1 && strings.Contains("0123456789", s.expression[k]) {
					h = append(h, s.expression[k])
				} else {
					h[len(h) - 1] += s.expression[k]
				}
			}
		}
	}

	var ActiveExpressions []Expression

	func InitActiveExpressions() {
		mu.Lock()
		os.Chmod("expressions.json", 0666)
		data, err2 := ioutil.ReadFile("expressions.json")
		if err2 != nil {
			http.Error(w, err2.Error(), http.StatusInternalServerError)
		}
	
		m := make(map[string]map[string]string)
		err3 := json.Unmarshal(data, &m)
		if err3 != nil {
			http.Error(w, err3.Error(), http.StatusInternalServerError)
		}
		mu.Unlock()
		for _, v := range m {
			if v.status == "Waiting..." || v.status == "Processing..." {
				ActiveExpressions = append(ActiveExpressions, Expression{processId: v.id, expression: v.expression})
			}
		}
	}

	go func() {
		ch := make(chan Task)
		res := make(chan int)
		for (i := 0; i < 2; i++) {
			go parser(ch, res)
		}
		for (i := 0; i < 2; i++) {
			go worker(ch, res)
		}
	}()*/
}
	