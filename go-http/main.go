// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// );

// /*___MODELS___*/
// type User struct {
// 	Name  string `json:"name"`
// 	Email string `json:"email"`
// };

// type Product struct {
// 	Name string `json:"name"`
// 	Price int16 `json:"price"`
// }

// type APIResponse struct {
// 	Message string `json:"message"`
// 	Success bool   `json:"status"`
// }

// /*___HANDLERS___*/
// func hello(response http.ResponseWriter, request *http.Request){
// 	response.WriteHeader(http.StatusOK);
// 	response.Write([]byte("Hello, World!"))
// }

// func productsHandler(response http.ResponseWriter, request *http.Request) {
// 	switch request.Method {
// 	case http.MethodGet:
// 		query := request.URL.Query()

// 		id := query.Get("id");
// 		name := query.Get("name");

// 		fmt.Println(id);
// 		fmt.Println(name);

// 		response.WriteHeader(http.StatusOK)
// 		response.Write([]byte("GET products"))

// 	case http.MethodPost:
// 		var product Product

// 		err := json.NewDecoder(request.Body).Decode(&product);

// 		if err != nil {
// 			http.Error(response, err.Error(), http.StatusBadRequest)
// 			return
// 		}

// 		fmt.Println(product.Name);
// 		fmt.Println(product.Price);

// 		response.WriteHeader(http.StatusCreated);
// 		response.Header().Set("Content-Type", "application/json");
// 		json.NewEncoder(response).Encode(APIResponse{Message: "POST product", Success: true});
// 	case http.MethodPut:
// 		response.WriteHeader(http.StatusOK)
// 		response.Write([]byte("PUT product"))

// 	case http.MethodDelete:
// 		response.WriteHeader(http.StatusOK)
// 		response.Write([]byte("DELETE product"))

// 	default:
// 		response.WriteHeader(http.StatusMethodNotAllowed)
// 		response.Write([]byte("Method not allowed"))
// 	}
// }

// func usersHandler(response http.ResponseWriter, request *http.Request){
// 	switch request.Method {
// 	case http.MethodGet:
// 		query := request.URL.Query()

// 		id := query.Get("id");
// 		name := query.Get("name");

// 		fmt.Println(id);
// 		fmt.Println(name);

// 		response.WriteHeader(http.StatusOK)
// 		response.Write([]byte("GET users"))

// 	case http.MethodPost:
// 		var user User

// 		err := json.NewDecoder(request.Body).Decode(&user);

// 		if err != nil {
// 			http.Error(response, err.Error(), http.StatusBadRequest)
// 			return
// 		}

// 		fmt.Println(user.Name);
// 		fmt.Println(user.Email);

// 		response.WriteHeader(http.StatusCreated);
// 		response.Header().Set("Content-Type", "application/json");
// 		json.NewEncoder(response).Encode(APIResponse{Message: "POST user", Success: true});
// 	case http.MethodPut:
// 		response.WriteHeader(http.StatusOK)
// 		response.Write([]byte("PUT user"))

// 	case http.MethodDelete:
// 		response.WriteHeader(http.StatusOK)
// 		response.Write([]byte("DELETE user"))

// 	default:
// 		response.WriteHeader(http.StatusMethodNotAllowed)
// 		response.Write([]byte("Method not allowed"))
// 	}
// }

// func main() {
// 	/* Mapping route paths with their handler functions */
// 	http.HandleFunc("/api/hello", hello);
// 	http.HandleFunc("/api/products", productsHandler);
// 	http.HandleFunc("/api/users", usersHandler);

// 	log.Println("Server running on port 8000");

// 	err := http.ListenAndServe(":8000", nil)
// 	if err != nil {
// 		log.Fatal("Server failed to start:", err)
// 	}
// };

/*
___IDIOMATIC GO___
*/
package main

import (
	"log"
	"net/http"
	"os"

	"go-http/config"
	"go-http/handlers"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ .env not found, using system env")
	}

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080" // default fallback
	}

	config.ConnectDB();
	
	/* Mapping route paths with their handler functions */
	http.HandleFunc("/api/products", handlers.ProductsHandler);
	http.HandleFunc("/api/products/", handlers.ProductsHandler);

	log.Println("🚀 Server running on port", port)
	
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
};