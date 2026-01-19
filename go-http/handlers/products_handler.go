package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go-http/helpers"
	"go-http/models"
)


func ProductsHandler(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		handleGetProduct(response, request);

	case http.MethodPost:
		handleCreateProduct(response, request);
	case http.MethodPut:
		handleUpdateProduct(response, request);

	case http.MethodDelete:
		handleDeleteProduct(response, request);

	default:
		helpers.JSONResponse(response, http.StatusMethodNotAllowed, models.APIResponse{
			Message: "Method not allowed",
			Success: false,
		});
	}
};

func handleGetProduct(response http.ResponseWriter, request *http.Request){
	query := request.URL.Query()

	id := query.Get("id");
	name := query.Get("name");
	
	fmt.Println(id);
	fmt.Println(name);

	helpers.JSONResponse(response, http.StatusOK, models.APIResponse{
		Message: "GET product", 
		Success: true,
	});
};

func handleCreateProduct(response http.ResponseWriter, request *http.Request){
	var product models.Product

		err := json.NewDecoder(request.Body).Decode(&product);

		if err != nil {
			helpers.JSONResponse(response, http.StatusBadRequest, models.APIResponse{
				Message: err.Error(),
				Success: false,
			})
			return
		}
		
		fmt.Println(product.Name);
		fmt.Println(product.Price);

		helpers.JSONResponse(response, http.StatusCreated, models.APIResponse{
			Message: "POST product",
			Success: true,
		})
};

func handleUpdateProduct(response http.ResponseWriter, request *http.Request){
	helpers.JSONResponse(response, http.StatusCreated, models.APIResponse{
		Message: "PUT product",
		Success: true,
	})
};

func handleDeleteProduct(response http.ResponseWriter, request *http.Request){
	helpers.JSONResponse(response, http.StatusCreated, models.APIResponse{
		Message: "DELETE product",
		Success: true,
	})
};