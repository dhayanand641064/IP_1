package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gocql/gocql"
)

type Profile struct {
	Name         string
	RollNumber   int
	Email        string
	GitHub       string
	Role         string
	IsAdmin      bool
	Projects     []string
	CodingLinks  []string
	Skills       []string
	Interests    []string
	Achievements []string
}

var Article = "The mountain gave birth to the mouse"

func GetProfileHandler(w http.ResponseWriter, r *http.Request, session *gocql.Session, name string) {
	profile, err := getProfile(session, name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, "Failed to marshal profile to JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func getProfile(session *gocql.Session, name string) (Profile, error) {
	var profile Profile
	query := "SELECT * FROM dhaya5 WHERE column1_name = ?"
	iter := session.Query(query, name).Iter()

	var result map[string]interface{}
	result = make(map[string]interface{})

	if iter.MapScan(result) {
		profile.Name = result["column1_name"].(string)
		profile.RollNumber = result["column2_roll_number"].(int)
		profile.Email = result["column3_email"].(string)
		profile.GitHub = result["column4_github"].(string)
		profile.Role = result["column5_role"].(string)
		profile.IsAdmin = result["column6_is_admin"].(bool)
		profile.Projects = result["column10_projects"].([]string)
		profile.CodingLinks = result["column11_coding_links"].([]string)
		profile.Skills = result["column7_skills"].([]string)
		profile.Interests = result["column8_interests"].([]string)
		profile.Achievements = result["column9_achievements"].([]string)

		return profile, nil
	}

	if err := iter.Close(); err != nil {
		return profile, fmt.Errorf("Failed to retrieve profile: %v", err)
	}

	return profile, fmt.Errorf("Profile not found")
}
