package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Mushus/apartment/model"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type ApartmentAPI struct {
	DB *gorm.DB
}

func (h ApartmentAPI) List(c echo.Context) error {
	var models []model.Apartment
	db := h.DB.Find(&models)
	if db.RecordNotFound() {
		return c.JSON(http.StatusOK, []model.ApartmentView{})
	}
	if err := db.Error; err != nil {
		log.Print(err)
		return nil
	}

	v := make([]model.ApartmentView, len(models))
	for i, dbm := range models {
		v[i] = dbm.ToViewModel()
	}

	return c.JSON(http.StatusOK, v)
}

func (h ApartmentAPI) Create(c echo.Context) error {
	var apartment model.ApartmentParam
	if err := c.Bind(&apartment); err != nil {
		log.Print(err)
		return nil
	}

	if err := c.Validate(&apartment); err != nil {
		log.Print(err)
		return nil
	}

	v, err := h.save(0, apartment)
	if err != nil {
		log.Print(err)
		return nil
	}

	return c.JSON(http.StatusOK, v)
}

func (h ApartmentAPI) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Print(err)
		return nil
	}

	var apartment model.ApartmentParam
	if err := c.Bind(&apartment); err != nil {
		log.Print(err)
		return nil
	}
	if err := c.Validate(&apartment); err != nil {
		log.Print(err)
		return nil
	}

	v, err := h.save(id, apartment)
	if err != nil {
		log.Print(err)
		return nil
	}
	return c.JSON(http.StatusOK, v)
}

func (h ApartmentAPI) save(id int, p model.ApartmentParam) (model.ApartmentView, error) {
	m := p.ToDBModel()
	if id > 0 {
		m.ID = id
	}

	err := h.DB.Save(&m).Error
	if err != nil {
		return model.ApartmentView{}, nil
	}

	return m.ToViewModel(), nil
}

func (h ApartmentAPI) Show(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Print(err)
		return nil
	}

	var m model.Apartment
	err = h.DB.Where("id = ?", id).First(&m).Error
	if err != nil {
		log.Print(err)
		return nil
	}

	return c.JSON(http.StatusOK, m.ToViewModel())
}
