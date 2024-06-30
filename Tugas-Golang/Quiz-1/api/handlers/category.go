package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"Quiz-1/api/models"
	"Quiz-1/db"
	"Quiz-1/utils"
	"github.com/julienschmidt/httprouter"
)

func GetCategories(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	rows, err := db.DB.Query("SELECT id, name, created_at, updated_at FROM categories")
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer func() {
		if err := rows.Close(); err != nil {
			utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		}
	}()

	var categories []models.Category
	for rows.Next() {
		var cat models.Category
		var createdAt, updatedAt interface{}
		if err := rows.Scan(&cat.ID, &cat.Name, &createdAt, &updatedAt); err != nil {
			utils.ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}

		cat.CreatedAt, err = utils.ParseTime(createdAt)
		if err != nil {
			utils.ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}

		cat.UpdatedAt, err = utils.ParseTime(updatedAt)
		if err != nil {
			utils.ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}

		categories = append(categories, cat)
	}

	utils.ResponseFormatter(w, http.StatusOK, "Categories retrieved successfully", categories)
}

func CreateCategory(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var cat models.Category
	if err := json.NewDecoder(r.Body).Decode(&cat); err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	if cat.Name == "" {
		utils.ResponseError(w, http.StatusBadRequest, "Name is required")
		return
	}

	result, err := db.DB.Exec("INSERT INTO categories (name, created_at, updated_at) VALUES (?, ?, ?)", cat.Name, time.Now(), time.Now())
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	cat.ID = int(id)
	cat.CreatedAt = time.Now()
	cat.UpdatedAt = time.Now()

	utils.ResponseFormatter(w, http.StatusCreated, "Category created successfully", cat)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, "Invalid category ID")
		return
	}

	var cat models.Category
	if err := json.NewDecoder(r.Body).Decode(&cat); err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	if cat.Name == "" {
		utils.ResponseError(w, http.StatusBadRequest, "Name is required")
		return
	}

	_, err = db.DB.Exec("UPDATE categories SET name = ?, updated_at = ? WHERE id = ?", cat.Name, time.Now(), id)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	cat.ID = id
	cat.UpdatedAt = time.Now()

	utils.ResponseFormatter(w, http.StatusOK, "Category updated successfully", cat)
}

func DeleteCategory(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, "Invalid category ID")
		return
	}

	_, err = db.DB.Exec("DELETE FROM categories WHERE id = ?", id)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.ResponseFormatter(w, http.StatusOK, "Category deleted successfully", nil)
}

func GetArticlesByCategoryID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	categoryID, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, "Invalid category ID")
		return
	}

	query := "SELECT id, title, content, image_url, article_length, created_at, updated_at, category_id FROM articles WHERE category_id = ?"
	var conditions []string
	var args []interface{}
	args = append(args, categoryID)

	// Handle query parameters for filtering
	if title := r.URL.Query().Get("title"); title != "" {
		conditions = append(conditions, "title LIKE ?")
		args = append(args, "%"+title+"%")
	}
	if minYear := r.URL.Query().Get("minYear"); minYear != "" {
		conditions = append(conditions, "YEAR(created_at) >= ?")
		args = append(args, minYear)
	}
	if maxYear := r.URL.Query().Get("maxYear"); maxYear != "" {
		conditions = append(conditions, "YEAR(created_at) <= ?")
		args = append(args, maxYear)
	}
	if minWord := r.URL.Query().Get("minWord"); minWord != "" {
		conditions = append(conditions, "(LENGTH(content) - LENGTH(REPLACE(content, ' ', '')) + 1) >= ?")
		args = append(args, minWord)
	}
	if maxWord := r.URL.Query().Get("maxWord"); maxWord != "" {
		conditions = append(conditions, "(LENGTH(content) - LENGTH(REPLACE(content, ' ', '')) + 1) <= ?")
		args = append(args, maxWord)
	}

	// Combine conditions with WHERE clause if any
	if len(conditions) > 0 {
		query += " AND " + strings.Join(conditions, " AND ")
	}

	if sortByTitle := r.URL.Query().Get("sortByTitle"); sortByTitle != "" {
		if sortByTitle == "asc" {
			query += " ORDER BY title ASC"
		} else if sortByTitle == "desc" {
			query += " ORDER BY title DESC"
		}
	}

	rows, err := db.DB.Query(query, args...)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer func() {
		if err := rows.Close(); err != nil {
			utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		}
	}()

	var articles []models.Article
	for rows.Next() {
		var art models.Article
		var createdAt, updatedAt interface{}
		if err := rows.Scan(&art.ID, &art.Title, &art.Content, &art.ImageURL, &art.ArticleLength, &createdAt, &updatedAt, &art.CategoryID); err != nil {
			utils.ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}

		art.CreatedAt, err = utils.ParseTime(createdAt)
		if err != nil {
			utils.ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}

		art.UpdatedAt, err = utils.ParseTime(updatedAt)
		if err != nil {
			utils.ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}

		articles = append(articles, art)
	}

	utils.ResponseFormatter(w, http.StatusOK, "Articles retrieved successfully", articles)
}
