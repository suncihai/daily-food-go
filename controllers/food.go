package controllers

import (
	"encoding/json"
	"go-daily-food/db"
	"go-daily-food/models"
	"net/http"
)

// GetFoodstore godoc
// @Summary Get foodstore request
// @Description Get foodstore from dailyfood database
// @Tags foodstore
// @Accept  json
// @Produce  json
// @Success 200 {array} models.FoodStore
// @Router /foodstore [get]
func GetFoodStore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	foods := []models.FoodStore{}
	
	rows, err := db.DB.Query("SELECT * FROM foodstore")
	if err != nil {
		RespondWithError(err, w)
	}

	defer rows.Close()

	for rows.Next() {
		var food models.FoodStore
		err = rows.Scan(&food.ID, &food.Name, &food.Category, &food.ShelfLife, &food.Src)
		if err != nil {
			RespondWithError(err, w)
		}
		foods = append(foods, food)
	}
    
	res := models.SuccessRes{Success: true, Status: 200, Data: foods}
	RespondWithSuccess(res, w)
}

// GetFood godoc
// @Summary Get food request
// @Description Get food from dailyfood database
// @Tags food
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Food
// @Router /food [get]
func GetFood(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	foods := []models.Food{}
	
	rows, err := db.DB.Query("SELECT * FROM food")
	if err != nil {
		RespondWithError(err, w)
	}

	defer rows.Close()

	for rows.Next() {
		var food models.Food
		err = rows.Scan(&food.ID, &food.Name, &food.Category, &food.CreatedAt, &food.IsEaten, &food.FoodId, &food.Quantity, &food.OwnerId, &food.OwnerName, &food.Price, &food.ShelfLife, &food.Src)
		if err != nil {
			RespondWithError(err, w)
		}
		foods = append(foods, food)
	}
    
	res := models.SuccessRes{Success: true, Status: 200, Data: foods}
	RespondWithSuccess(res, w)
}

// @Summary Create a new food
// @Description Create a new food with the input paylod
// @Tags food
// @Accept  json
// @Produce  json
// @Param food body models.Food true "Create food"
// @Success 200
// @Router /food [post]
func CreateFood(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	food := models.Food{}
	err := json.NewDecoder(r.Body).Decode(&food)
	if err != nil {
		RespondWithError(err, w)
	} else { 
		_, err = db.DB.Exec(`INSERT INTO food (name, category, created_at, is_eaten, food_id, quantity, owner_id, owner_name, price, shelf_life, src) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`, food.Name, food.Category, food.CreatedAt, food.IsEaten, food.FoodId, food.Quantity, food.OwnerId, food.OwnerName, food.Price, food.ShelfLife, food.Src)
		if err != nil {
			RespondWithError(err, w)
		}
		res := models.SuccessRes{Success: true, Status: 200}
		RespondWithSuccess(res, w)
	}
}

// @Summary Edit a food
// @Description Edit a food with the input paylod
// @Tags food
// @Accept  json
// @Produce  json
// @Param food body models.Food true "Edit food"
// @Success 200
// @Router /food [put]
func EditFood(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	food := models.Food{}
	err := json.NewDecoder(r.Body).Decode(&food)
	if err != nil {
		RespondWithError(err, w)
	} else { 
		_, err = db.DB.Exec(`UPDATE food SET name = ?, category = ?, created_at = ?, is_eaten = ?, food_id = ?, quantity = ?, owner_id = ?, owner_name = ?, price = ?, shelf_life = ?, src = ? WHERE id = ?`, food.Name, food.Category, food.CreatedAt, food.IsEaten, food.FoodId, food.Quantity, food.OwnerId, food.OwnerName, food.Price, food.ShelfLife, food.Src, food.ID)
		if err != nil {
			RespondWithError(err, w)
		}
		res := models.SuccessRes{Success: true, Status: 200}
		RespondWithSuccess(res, w)
	}
}

// @Summary Delete a food
// @Description Delete a food that is eaten
// @Tags food
// @Accept  json
// @Produce  json
// @Param food body models.Food true "Delete food"
// @Success 200
// @Router /delete-food [post]
func DeleteFood(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	food := models.Food{}
	err := json.NewDecoder(r.Body).Decode(&food)
	if err != nil {
		RespondWithError(err, w)
	} else { 
		_, err = db.DB.Exec(`DELETE FROM food WHERE id = ?`, food.ID)
		if err != nil {
			RespondWithError(err, w)
		}
		res := models.SuccessRes{Success: true, Status: 200}
		RespondWithSuccess(res, w)
	}
}