package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"Quiz-1/api/models"
	"Quiz-1/db"
	"Quiz-1/utils"
	"github.com/julienschmidt/httprouter"
)

func GetArticles(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	query := "SELECT id, title, content, image_url, article_length, created_at, updated_at, category_id FROM articles"
	var conditions []string
	var args []interface{}

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
		query += " WHERE " + strings.Join(conditions, " AND ")
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
		var categoryID sql.NullInt64
		if err := rows.Scan(&art.ID, &art.Title, &art.Content, &art.ImageURL, &art.ArticleLength, &createdAt, &updatedAt, &categoryID); err != nil {
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

		if categoryID.Valid {
			art.CategoryID = int(categoryID.Int64)
		} else {
			art.CategoryID = 0
		}

		articles = append(articles, art)
	}

	utils.ResponseFormatter(w, http.StatusOK, "Articles retrieved successfully", articles)
}

func CreateArticle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var art models.Article
	if err := json.NewDecoder(r.Body).Decode(&art); err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := validateArticleFields(art); err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	if !categoryExists(art.CategoryID) {
		utils.ResponseError(w, http.StatusBadRequest, "Category not found")
		return
	}

	art.ArticleLength = calculateArticleLength(art.Content)
	art.CreatedAt = time.Now()
	art.UpdatedAt = time.Now()

	result, err := db.DB.Exec("INSERT INTO articles (title, content, image_url, article_length, created_at, updated_at, category_id) VALUES (?, ?, ?, ?, ?, ?, ?)", art.Title, art.Content, art.ImageURL, art.ArticleLength, art.CreatedAt, art.UpdatedAt, art.CategoryID)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	art.ID = int(id)

	utils.ResponseFormatter(w, http.StatusCreated, "Article created successfully", art)
}

func UpdateArticle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, "Invalid article ID")
		return
	}

	var art models.Article
	if err := json.NewDecoder(r.Body).Decode(&art); err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := validateArticleFields(art); err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	if !categoryExists(art.CategoryID) {
		utils.ResponseError(w, http.StatusBadRequest, "Category not found")
		return
	}

	art.ArticleLength = calculateArticleLength(art.Content)
	art.UpdatedAt = time.Now()

	_, err = db.DB.Exec("UPDATE articles SET title = ?, content = ?, image_url = ?, article_length = ?, updated_at = ?, category_id = ? WHERE id = ?", art.Title, art.Content, art.ImageURL, art.ArticleLength, art.UpdatedAt, art.CategoryID, id)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	art.ID = id

	utils.ResponseFormatter(w, http.StatusOK, "Article updated successfully", art)
}

func DeleteArticle(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, "Invalid article ID")
		return
	}

	_, err = db.DB.Exec("DELETE FROM articles WHERE id = ?", id)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.ResponseFormatter(w, http.StatusOK, "Article deleted successfully", nil)
}

func validateArticleFields(art models.Article) error {
	var errs []string
	if art.Title == "" {
		errs = append(errs, "title is required")
	}
	if art.Content == "" {
		errs = append(errs, "content is required")
	}
	if art.ImageURL == "" {
		errs = append(errs, "image_url is required")
	}
	if art.CategoryID == 0 {
		errs = append(errs, "category_id is required")
	}
	if len(strings.Fields(art.Title)) > 3 {
		errs = append(errs, "title must not exceed 3 words")
	}
	if !isValidURL(art.ImageURL) {
		errs = append(errs, "invalid URL for image_url")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func calculateArticleLength(content string) string {
	wordCount := len(strings.Fields(content))
	charCount := len(content)

	if wordCount <= 100 || charCount <= 400 {
		return "pendek"
	}
	if wordCount <= 200 || charCount <= 800 {
		return "sedang"
	}
	return "panjang"
}

func isValidURL(urlStr string) bool {
	u, err := url.Parse(urlStr)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func categoryExists(categoryID int) bool {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM categories WHERE id = ?)"
	err := db.DB.QueryRow(query, categoryID).Scan(&exists)
	return err == nil && exists
}
