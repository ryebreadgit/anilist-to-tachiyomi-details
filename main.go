package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/buger/jsonparser"
	"github.com/imroc/req/v2"
)

func errHandle(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func anilist(id int, query string) (string, error) {
	client := req.C()

	postBod := &PostBodyStruct{Query: query, Variables: VariablesStruct{Id: id}}

	resp, err := client.R().SetBodyJsonMarshal(postBod).SetHeader("Content-Type", "application/json").Post("https://graphql.anilist.co")

	if resp.StatusCode != 200 {
		return "", errors.New(resp.Status)
	}

	errHandle(err)
	return resp.ToString()
}

func getManga(id int) ([]byte, error) {
	query := `
	query ($id: Int) {
		Media(id: $id, type : MANGA) {
			id
			title {
				romaji
			}
			description
			genres
			status
			coverImage {
				extraLarge
				large
				medium
			}
			staff{
				nodes {
					id
				}
				edges {
					id
					role
				}
			}
		}
	}`

	data, err := anilist(id, query)
	return []byte(data), err
}

func getMangaka(id int) ([]byte, error) {
	query := `query ($id: Int) {
		Staff(id: $id) {
		id
			name {
				full
			}
		}
	}`

	data, err := anilist(id, query)

	if err != nil {
		return nil, err
	}

	res, err := jsonparser.GetString([]byte(data), "data", "Staff", "name", "full")

	return []byte(res), err
}

func descParse(data string) string {
	res := strings.ReplaceAll(data, "\n<br>", "\n")
	res = strings.ReplaceAll(res, "<br>\n", "\n")
	res = strings.ReplaceAll(res, "<br>", "\n")
	res = strings.ReplaceAll(res, "<i>", "")
	res = strings.ReplaceAll(res, "</i>", "")
	res = strings.ReplaceAll(res, "< !--link-->", "")
	res = strings.ReplaceAll(res, "<b>", "")
	res = strings.ReplaceAll(res, "</b>", "")
	return res
}

func makeComicInfo(details DetailsStruct) ComicInfoStruct {
	comicInfo := ComicInfoStruct{}
	return comicInfo
}

func main() {
	reader := bufio.NewReader(os.Stdin) //create new reader, assuming bufio imported
	fmt.Print("Enter Manga ID: ")
	userInput, _ := reader.ReadString('\n')
	userInput = strings.TrimSuffix(userInput, "\n")
	userInput = strings.TrimSuffix(userInput, "\r")

	mangaID, err := strconv.Atoi(userInput)
	errHandle(err)

	baseInfo, err := getManga(mangaID)
	errHandle(err)

	var tempStr string

	details := DetailsStruct{
		StatusValues: []string{"0 = Unknown", "1 = Ongoing", "2 = Completed", "3 = Licensed"},
		AnilistID:    mangaID,
	}

	details.Title, err = jsonparser.GetString(baseInfo, "data", "Media", "title", "romaji")
	errHandle(err)

	tempStr, err = jsonparser.GetString(baseInfo, "data", "Media", "description")
	errHandle(err)

	details.Description = descParse(tempStr)

	jsonparser.ArrayEach(baseInfo, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		details.Genre = append(details.Genre, string(value))
	}, "data", "Media", "genres")

	tempStr, err = jsonparser.GetString(baseInfo, "data", "Media", "status")
	errHandle(err)

	switch tempStr { // Set status appropriately
	case "RELEASING":
		details.Status = "1"
	case "FINISHED":
		details.Status = "2"
	case "NOT_YET_RELEASED":
		details.Status = "3"
	default:
		details.Status = "0"
	}

	// Get author info
	count := 0
	jsonparser.ArrayEach(baseInfo, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		errHandle(err)

		authrole, err := jsonparser.GetString(value, "role")
		errHandle(err)
		authid, err := jsonparser.GetInt(baseInfo, "data", "Media", "staff", "nodes", "["+strconv.Itoa(count)+"]", "id")
		errHandle(err)

		if authrole == "Story & Art" {
			authname, err := getMangaka(int(authid))
			errHandle(err)
			details.Author = string(authname)
			details.Artist = string(authname)
		} else if authrole == "Story" {
			authname, err := getMangaka(int(authid))
			errHandle(err)
			details.Author = string(authname)
		} else if authrole == "Art" {
			authname, err := getMangaka(int(authid))
			errHandle(err)
			details.Artist = string(authname)
		}
		count++
	}, "data", "Media", "staff", "edges")

	errHandle(err)

	var coverImage string

	tempStr, err = jsonparser.GetString(baseInfo, "data", "Media", "coverImage", "medium")
	errHandle(err)
	if tempStr != "" {
		coverImage = tempStr
	}
	tempStr, err = jsonparser.GetString(baseInfo, "data", "Media", "coverImage", "large")
	errHandle(err)
	if tempStr != "" {
		coverImage = tempStr
	}
	tempStr, err = jsonparser.GetString(baseInfo, "data", "Media", "coverImage", "extraLarge")
	errHandle(err)
	if tempStr != "" {
		coverImage = tempStr
	}

	req.C().R().SetOutputFile("./cover.jpg").Get(coverImage)

	f, err := os.Create("details.json")

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	defer f.Close()

	data, _ := json.MarshalIndent(details, "", "\t")

	_, err2 := f.WriteString(string(data))

	if err2 != nil {
		fmt.Println(err2.Error())
		os.Exit(1)
	}

	fmt.Printf("Successfully downloaded manga details for \"%v\"!\n", details.Title)

	fmt.Print("Press Enter to Exit...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

}
