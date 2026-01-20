package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"go-http/config"
	"go-http/helpers"
	"go-http/models"
)


func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	var productId string
	if len(parts) == 3 && parts[1] == "products" {
		productId = parts[2]
	} else {
		productId = ""
	}

	switch r.Method {
	case http.MethodGet:
		handleGetProducts(w, r);
	case http.MethodPost:
		handleCreateProduct(w, r);
	case http.MethodPut:
		if(productId == ""){
			helpers.JSONResponse(w, http.StatusBadRequest, models.APIResponse{
				Message: "Product id is required",
				Success: false,
			})
			return
		}
		handleUpdateProduct(w, r, productId);
	case http.MethodDelete:
		handleDeleteProduct(w, r);
	default:
		helpers.JSONResponse(w, http.StatusMethodNotAllowed, models.APIResponse{
			Message: "Method not allowed",
			Success: false,
		});
	}
};

func handleGetProducts(w http.ResponseWriter, r *http.Request){
	q := r.URL.Query();

	priceStr := q.Get("price")

	var price interface{} = nil
	if priceStr != "" {
		p, err := strconv.Atoi(priceStr)
		if err != nil {
			http.Error(w, "Invalid price", http.StatusBadRequest)
			return
		}
		price = p
	}

	query := `SELECT id, name, price
FROM products
WHERE ($1 = '' OR name ILIKE '%' || $1 || '%')
  AND (price = $2::int OR $2 IS NULL)`

	rows, err := config.DB.Query(query, q.Get("name"), price);

	if(err != nil){
		helpers.JSONResponse(w, http.StatusInternalServerError, models.APIResponse{
			Message: err.Error(), 
			Success: false,
			Data: nil,
		});
	};

	/*
		`defer`	: 
		▪️Function yang akan selalu di eksekusi walaupun terjadi error di dalam function tsb
		▪️Mirip seperti `finally` di Javascript

		`rows.Close()`:
		▪️`rows` menahan connection di database. Jika tidak ditutup/di close, maka: 
			➡️ Connection akan habis
			➡️ Membuat aplikasi menjadi error: too many connections

		  Dengan `rows.Close()` membuat connection database kita di kembalikan ke pool nya.

		✅ Dengan `defer rows.Close()`, membuat connection pool akan selalu dikembalikan walaupun eksekusi function nya gagal ataupun berhasil
	*/
	defer rows.Close();

	products := []models.Product{};

	for rows.Next() {
		var p models.Product;
		err := rows.Scan(&p.Id, &p.Name, &p.Price);
		if err != nil {
			helpers.JSONResponse(w, http.StatusInternalServerError, models.APIResponse{
				Message: err.Error(), 
				Success: false,
				Data: nil,
			});
			return;
		}
		products = append(products, p);
	}

	helpers.JSONResponse(w, http.StatusOK, models.APIResponse{
		Message: "GET products successful", 
		Success: true,
		Data: products,
	});
};

func handleCreateProduct(w http.ResponseWriter, r *http.Request){
	defer r.Body.Close()
	
	var product models.Product

	err := json.NewDecoder(r.Body).Decode(&product)
	
	if err != nil {
		helpers.JSONResponse(w, http.StatusBadRequest, models.APIResponse{
			Message: err.Error(),
			Success: false,
		})
		return
	}

	err = config.DB.QueryRow(`
		INSERT INTO products (name, price)
		VALUES ($1, $2)
		RETURNING id
	`, product.Name, product.Price).Scan(&product.Id)

	if err != nil {
		helpers.JSONResponse(w, http.StatusInternalServerError, models.APIResponse{
			Message: err.Error(),
			Success: false,
			Data: nil, 
		})
		return
	}

	helpers.JSONResponse(w, http.StatusCreated, models.APIResponse{
		Message: "Create product successful",
		Success: true,
		Data: product, 
	})
};

func handleUpdateProduct(w http.ResponseWriter, r *http.Request, productId string){
	defer r.Body.Close()
	
	var product models.Product

	err := json.NewDecoder(r.Body).Decode(&product)
	
	if err != nil {
		helpers.JSONResponse(w, http.StatusBadRequest, models.APIResponse{
			Message: err.Error(),
			Success: false,
		})
		return
	};

	result, err := config.DB.Exec(`
		UPDATE products
		SET name=$1, price=$2
		WHERE id=$3
	`, product.Name, product.Price, productId)

	if err != nil {
		helpers.JSONResponse(w, http.StatusInternalServerError, models.APIResponse{
			Message: err.Error(),
			Success: false,
		})
		return
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		helpers.JSONResponse(w, http.StatusNotFound, models.APIResponse{
			Message: "Product not found",
			Success: false,
		})
		return
	}

	helpers.JSONResponse(w, http.StatusCreated, models.APIResponse{
		Message: fmt.Sprintf("Update product with id %s successful", productId),
		Success: true,
	})
};

func handleDeleteProduct(w http.ResponseWriter, r *http.Request){
	helpers.JSONResponse(w, http.StatusCreated, models.APIResponse{
		Message: "DELETE product",
		Success: true,
	})
};