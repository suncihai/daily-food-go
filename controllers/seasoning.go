package controllers

import (
	"encoding/json"
	"go-daily-food/db"
	"go-daily-food/models"
	"net/http"
)

// GetSeasoningstore godoc
// @Summary Get seasoningstore request
// @Description Get seasoningstore from dailyfood database
// @Tags seasoningstore
// @Accept  json
// @Produce  json
// @Success 200 {array} models.SeasoningStore
// @Router /seasoningstore [get]
func GetSeasoningStore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	seasonings := []models.SeasoningStore{}
	
	rows, err := db.DB.Query("SELECT * FROM seasoningstore")
	if err != nil {
		RespondWithError(err, w)
	}

	defer rows.Close()

	for rows.Next() {
		var seasoning models.SeasoningStore
		err = rows.Scan(&seasoning.ID, &seasoning.Name, &seasoning.Src)
		if err != nil {
			RespondWithError(err, w)
		}
		seasonings = append(seasonings, seasoning)
	}
    
	res := models.SuccessRes{Success: true, Status: 200, Data: seasonings}
	RespondWithSuccess(res, w)
}

// GetSeasoing godoc
// @Summary Get seasoning request
// @Description Get seasoning from dailyfood database
// @Tags seasoning
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Seasoning
// @Router /seasoning [get]
func GetSeasoning(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	seasonings := []models.Seasoning{}
	
	rows, err := db.DB.Query("SELECT * FROM seasoning")
	if err != nil {
		RespondWithError(err, w)
	}

	defer rows.Close()

	for rows.Next() {
		var seasoning models.Seasoning
		err = rows.Scan(&seasoning.ID, &seasoning.Name, &seasoning.CreatedAt, &seasoning.IsEaten, &seasoning.SeasoningId, &seasoning.Quantity, &seasoning.OwnerId, &seasoning.OwnerName)
		if err != nil {
			RespondWithError(err, w)
		}
		seasonings = append(seasonings, seasoning)
	}
    
	res := models.SuccessRes{Success: true, Status: 200, Data: seasonings}
	RespondWithSuccess(res, w)
}

// @Summary Create a new seasoning
// @Description Create a new seasoning with the input paylod
// @Tags seasoning
// @Accept  json
// @Produce  json
// @Param seasoning body models.Food true "Create seasoning"
// @Success 200
// @Router /seasoning [post]
func CreateSeasoning(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	seasoning := models.Seasoning{}
	err := json.NewDecoder(r.Body).Decode(&seasoning)
	if err != nil {
		RespondWithError(err, w)
	} else { 
		_, err = db.DB.Exec("INSERT INTO seasoning (id, name, created_at, is_eaten, seasoning_id, quantity, owner_id, owner_name) VALUES (? ? ? ? ? ? ? ?)", seasoning.ID, seasoning.Name, seasoning.CreatedAt, seasoning.IsEaten, seasoning.SeasoningId, seasoning.Quantity, seasoning.OwnerId, seasoning.OwnerName)
		if err != nil {
			RespondWithError(err, w)
		}
		RespondWithSuccess(true, w)
	}
}