package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"Quiz-1/utils"
	"github.com/julienschmidt/httprouter"
)

func CalculateShape(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	shapeType := ps.ByName("shape")
	calcType := r.URL.Query().Get("hitung")

	var result string
	var err error

	switch shapeType {
	case "persegi":
		sisi, err := strconv.Atoi(r.URL.Query().Get("sisi"))
		if err != nil {
			utils.ResponseError(w, http.StatusBadRequest, "Invalid value for 'sisi'")
			return
		}
		result, err = calculatePersegi(sisi, calcType)
	case "persegi-panjang":
		panjang, err := strconv.Atoi(r.URL.Query().Get("panjang"))
		if err != nil {
			utils.ResponseError(w, http.StatusBadRequest, "Invalid value for 'panjang'")
			return
		}
		lebar, err := strconv.Atoi(r.URL.Query().Get("lebar"))
		if err != nil {
			utils.ResponseError(w, http.StatusBadRequest, "Invalid value for 'lebar'")
			return
		}
		result, err = calculatePersegiPanjang(panjang, lebar, calcType)
	case "lingkaran":
		jariJari, err := strconv.Atoi(r.URL.Query().Get("jariJari"))
		if err != nil {
			utils.ResponseError(w, http.StatusBadRequest, "Invalid value for 'jariJari'")
			return
		}
		result, err = calculateLingkaran(jariJari, calcType)
	case "kubus":
		sisi, err := strconv.Atoi(r.URL.Query().Get("sisi"))
		if err != nil {
			utils.ResponseError(w, http.StatusBadRequest, "Invalid value for 'sisi'")
			return
		}
		result, err = calculateKubus(sisi, calcType)
	case "balok":
		panjang, err := strconv.Atoi(r.URL.Query().Get("panjang"))
		if err != nil {
			utils.ResponseError(w, http.StatusBadRequest, "Invalid value for 'panjang'")
			return
		}
		lebar, err := strconv.Atoi(r.URL.Query().Get("lebar"))
		if err != nil {
			utils.ResponseError(w, http.StatusBadRequest, "Invalid value for 'lebar'")
			return
		}
		tinggi, err := strconv.Atoi(r.URL.Query().Get("tinggi"))
		if err != nil {
			utils.ResponseError(w, http.StatusBadRequest, "Invalid value for 'tinggi'")
			return
		}
		result, err = calculateBalok(panjang, lebar, tinggi, calcType)
	case "tabung":
		jariJari, err := strconv.Atoi(r.URL.Query().Get("jariJari"))
		if err != nil {
			utils.ResponseError(w, http.StatusBadRequest, "Invalid value for 'jariJari'")
			return
		}
		tinggi, err := strconv.Atoi(r.URL.Query().Get("tinggi"))
		if err != nil {
			utils.ResponseError(w, http.StatusBadRequest, "Invalid value for 'tinggi'")
			return
		}
		result, err = calculateTabung(jariJari, tinggi, calcType)
	default:
		utils.ResponseError(w, http.StatusBadRequest, "Invalid shape type")
		return
	}

	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.ResponseFormatter(w, http.StatusOK, "Calculation successful", result)
}

func calculatePersegi(sisi int, calcType string) (string, error) {
	switch calcType {
	case "luas":
		return strconv.Itoa(sisi * sisi), nil
	case "keliling":
		return strconv.Itoa(4 * sisi), nil
	default:
		return "", fmt.Errorf("invalid calculation type for persegi")
	}
}

func calculatePersegiPanjang(panjang, lebar int, calcType string) (string, error) {
	switch calcType {
	case "luas":
		return strconv.Itoa(panjang * lebar), nil
	case "keliling":
		return strconv.Itoa(2 * (panjang + lebar)), nil
	default:
		return "", fmt.Errorf("invalid calculation type for persegi-panjang")
	}
}

func calculateLingkaran(jariJari int, calcType string) (string, error) {
	switch calcType {
	case "luas":
		return strconv.FormatFloat(float64(jariJari*jariJari)*3.14, 'f', 2, 64), nil
	case "keliling":
		return strconv.FormatFloat(2*float64(jariJari)*3.14, 'f', 2, 64), nil
	default:
		return "", fmt.Errorf("invalid calculation type for lingkaran")
	}
}

func calculateKubus(sisi int, calcType string) (string, error) {
	switch calcType {
	case "volume":
		return strconv.Itoa(sisi * sisi * sisi), nil
	case "luasPermukaan":
		return strconv.Itoa(6 * sisi * sisi), nil
	default:
		return "", fmt.Errorf("invalid calculation type for kubus")
	}
}

func calculateBalok(panjang, lebar, tinggi int, calcType string) (string, error) {
	switch calcType {
	case "volume":
		return strconv.Itoa(panjang * lebar * tinggi), nil
	case "luasPermukaan":
		return strconv.Itoa(2 * (panjang*lebar + panjang*tinggi + lebar*tinggi)), nil
	default:
		return "", fmt.Errorf("invalid calculation type for balok")
	}
}

func calculateTabung(jariJari, tinggi int, calcType string) (string, error) {
	switch calcType {
	case "volume":
		return strconv.FormatFloat(3.14*float64(jariJari)*float64(jariJari)*float64(tinggi), 'f', 2, 64), nil
	case "luasPermukaan":
		return strconv.FormatFloat(2*3.14*float64(jariJari)*(float64(jariJari)+float64(tinggi)), 'f', 2, 64), nil
	default:
		return "", fmt.Errorf("invalid calculation type for tabung")
	}
}
