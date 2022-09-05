package val

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func retrieveIDFromURL(urlStr string, model string) (int64, error) {
	url, err := url.Parse(urlStr)
	if err != nil {
		return 0, err
	}

	var id string
	urlSlice := strings.Split(url.Path, "/")
	for i, v := range urlSlice {
		if v == model {
			if len(urlSlice) > i+1 {
				id = urlSlice[i+1]
			}
		}
	}

	i, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}

	return int64(i), nil
}

func formToMap(r *http.Request, params []string) map[string]string {
	formMap := make(map[string]string)

	for _, k := range params {
		formMap[k] = r.PostFormValue(k)
	}

	return formMap
}

func checkFormParams(r *http.Request, params []string) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}

	for _, k := range params {
		_, e := r.PostForm[k]
		if !e {
			return fmt.Errorf("%s must exists", k)
		}

		if r.PostFormValue(k) == "" {
			return fmt.Errorf("%s cannot be black", k)
		}
	}

	return nil
}

func checkPasswordConfirmation(password string, passwordConfirmation string) error {
	if password != passwordConfirmation {
		return fmt.Errorf("password confirmation does not match")
	}
	return nil
}

func checkFormPassword(password string, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(password),
	)
	if err != nil {
		return err
	}

	return nil
}

func checkDateAfter(date time.Time, model string) error {
	if time.Now().After(date) {
		return fmt.Errorf("%s date has already passed", model)
	}

	return nil
}

func checkUserID(formUserID string, userID int64) error {
	id, err := strconv.Atoi(formUserID)
	if err != nil {
		return err
	}

	if userID != int64(id) {
		return fmt.Errorf("user id's do not match")
	}

	return nil
}

func checkURLQueries(queries string, permitted map[string]string) (map[string]string, error) {
	cleanQueries := make(map[string]string)
	for _, v := range strings.Split(queries, "&") {
		query := strings.Split(v, "=")
		queryVal := strings.ReplaceAll(query[1], "+", " ")

		if _, ok := permitted[query[0]]; ok {
			switch query[0] {
			case "order_by":
				cleanQueries["ord_"+query[0]] = queryVal
				break
			case "desc":
				cleanQueries["ord_"+query[0]] = queryVal
				break
			case "asc":
				break
			case "location":
				cleanQueries["sel_arr_"+query[0]] = queryVal
			default:
				cleanQueries["sel_"+query[0]] = queryVal
				break
			}
		} else {
			return cleanQueries, fmt.Errorf("parameter \"%s\" unrecognized or not permitted", query[0])
		}
	}

	return cleanQueries, nil
}
