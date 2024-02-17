package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"resfulsimple/internal/models"
	"strconv"
	"strings"
)

// CustomerHandler CRUD для продуктів
func CustomerHandler(w http.ResponseWriter, r *http.Request) { //nolint:funlen
	switch r.Method {
	case http.MethodGet:
		if getCustomers(w, r) {
			return
		}
	case http.MethodPost:
		if createCustomer(w, r) {
			return
		}
	case http.MethodPut, http.MethodPatch:
		if updateCustomer(w, r) {
			return
		}
	case http.MethodDelete:
		if deleteCustomer(w, r) {
			return
		}
	default:
		http.Error(w, "Метод не підтримується", http.StatusMethodNotAllowed)
	}
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) bool {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Не вказано id", http.StatusBadRequest)
		return true
	}
	customers := models.GetCustomersData()
	pid, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, fmt.Errorf("can't parse customer id:%w", err).Error(), http.StatusBadRequest)
		return true
	}
	customers.Delete(pid)
	w.WriteHeader(http.StatusOK)
	return false
}

func updateCustomer(w http.ResponseWriter, r *http.Request) bool {
	//id := r.URL.Query().Get("id")
	//if id == "" {
	//	http.Error(w, "Не вказано id", http.StatusBadRequest)
	//	return true
	//}
	var customer models.Customer
	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return true
	}
	customers := models.GetCustomersData()
	if err := customers.Update(customer); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return true
	}
	w.WriteHeader(http.StatusAccepted)
	return false
}

func createCustomer(w http.ResponseWriter, r *http.Request) bool {
	var newCustomer models.Customer
	data, _ := io.ReadAll(r.Body)
	_ = r.Body.Close()
	if err := json.Unmarshal(data, &newCustomer); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("can't parse request: %v: %s", err, string(data))
		return true
	}
	customers := models.GetCustomersData()
	if customers.Find(newCustomer.ID) != nil {
		http.Error(w, "Користувач з таким id вже існує", http.StatusAlreadyReported)
		return true
	}
	customers.Add(newCustomer)
	w.WriteHeader(http.StatusCreated)
	return false
}

func getCustomers(w http.ResponseWriter, r *http.Request) bool {
    id := r.URL.Query().Get("id")

    pageStr := r.URL.Query().Get("page")
    pageSizeStr := r.URL.Query().Get("pageSize")

    headerKey := "Pagination-Key"
    headerValue := r.Header.Get(headerKey)

    switch id {
    case "":
        var page, pageSize int

        if pageStr != "" {
            page, _ = strconv.Atoi(pageStr)
        }

        if pageSizeStr != "" {
            pageSize, _ = strconv.Atoi(pageSizeStr)
        }

        if page == 0 && pageSize == 0 && headerValue != "" {

            headerValues := strings.Split(headerValue, ";")

            for _, val := range headerValues {
                parts := strings.Split(strings.TrimSpace(val), "=")
                if len(parts) == 2 {
                    switch parts[0] {
                    case "page":
                        page, _ = strconv.Atoi(parts[1])
                    case "pageSize":
                        pageSize, _ = strconv.Atoi(parts[1])
                    }
                }
            }

        }

        w.Header().Set("Content-Type", "application/json")
        if err := json.NewEncoder(w).Encode(models.GetCustomersData().Get()); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    default:
        pid, err := strconv.Atoi(id)
        if err != nil {
            http.Error(w, fmt.Errorf("can't parse customer id:%w", err).Error(), http.StatusBadRequest)
            return true
        }
        customer := models.GetCustomersData().Find(pid)
        if customer == nil {
            http.Error(w, "Користувач не знайдено", http.StatusNotFound)
            return true
        }
        w.Header().Set("Content-Type", "application/json")
        if err := json.NewEncoder(w).Encode(customer); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    }
    return false
}
