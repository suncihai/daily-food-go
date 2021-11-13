package models

import "time"

type FoodStore struct {
	ID       int    `json:"id" example:"1"`
	Name     string `json:"name" example:"Pork"`
	Category string `json:"category" example:"meat"`
	Src      string `json:"src" example:"http://imageurl"`
}
type Food struct {
    ID        int        `json:"id" example:"1"`
	Name      string     `json:"name" example:"Pork"`
	Category  string     `json:"category" example:"meat"`
	CreatedAt time.Time  `json:"create_at" example:"2021/11/20 15:30:42"`
	IsEaten   bool       `json:"is_eaten" example:"false"`
	FoodId    int        `json:"food_id" example:"15"`
	Quantity  float32    `json:"quantity" example:"12"`
	OwnerId   int        `json:"owner_id" example:"12"`
	OwnerName string     `json:"owner_name" example:"Peter"`
	Price     float32    `json:"price" example:"5.18"`
	ShelfLife int        `json:"shelf_life" example:"30"`
}
type SeasoningStore struct {
	ID   int    `json:"id" example:"1"`
	Name string `json:"name" example:"Salt"`
	Src  string `json:"src" example:"http://imageurl"`
}
type Seasoning struct {
	ID           int        `json:"id" example:"1"`
	Name         string     `json:"name" example:"Salt"`
	CreatedAt    time.Time  `json:"create_at" example:"2021/11/20 15:30:42"`
	IsEaten      bool       `json:"is_eaten" example:"false"`
	SeasoningId  int        `json:"seasoning_id" example:"15"`
	Quantity     float32    `json:"quantity" example:"12"`
	OwnerId      int        `json:"owner_id" example:"12"`
	OwnerName    string     `json:"owner_name" example:"Peter"`
}
type Recipe struct {
	ID          int       `json:"id" example:"1"`
	Name        string    `json:"name" exampe:"Kun Pao Chicken"`
	Type        string    `json:"type" exampe:"Thai Food"`
	Foods       string    `json:"foods" exampe:"Kun Pao Chicken"`
	Seasonings  string    `json:"seasonings" exampe:"Salt"`
	OwnerId     int       `json:"owner_id" example:"12"`
	OwnerName   string    `json:"owner_name" example:"Peter"`
	Like        float32   `json:"like" example:"3.8"`
	Difficulty  float32   `json:"difficulty" example:"4.2"`
}
type SuccessRes struct {
	Success bool `json:"success"`
	Status int64 `json:"status"`
	Data interface{} `json:"data"`
 }