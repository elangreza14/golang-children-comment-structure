package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
)

type Comment struct {
	Content       string    `json:"content"`
	ID            int64     `json:"id"`
	ParentID      int64     `json:"parent_id"`
	TotalChildren int64     `json:"total_children"`
	Children      []Comment `json:"children"`
}

func main() {
	newComment1 := Comment{
		Content:  "Andi",
		ID:       1,
		ParentID: 0,
	}
	newComment2 := Comment{
		Content:  "Candra",
		ID:       2,
		ParentID: 1,
	}
	newComment3 := Comment{
		Content:  "Budi",
		ID:       3,
		ParentID: 0,
	}
	newComment4 := Comment{
		Content:  "Dedi",
		ID:       4,
		ParentID: 2,
	}
	newComment5 := Comment{
		Content:  "Ega",
		ID:       5,
		ParentID: 2,
	}
	newComment6 := Comment{
		Content:  "Fall",
		ID:       6,
		ParentID: 5,
	}
	newComment7 := Comment{
		Content:  "Gargda",
		ID:       7,
		ParentID: 3,
	}
	newComment8 := Comment{
		Content:  "Hanif",
		ID:       8,
		ParentID: 3,
	}
	newComment9 := Comment{
		Content:  "Icaas",
		ID:       9,
		ParentID: 0,
	}

	childrenID := []int64{}
	baseComments := []Comment{
		newComment1,
		newComment2,
		newComment3,
		newComment4,
		newComment5,
		newComment6,
		newComment7, newComment8, newComment9}
	comments := []Comment{}
	for i := len(baseComments) - 1; i >= 0; i-- {
		// foundedID := -1
		foundedID := []int64{}
		for j := 0; j < len(comments); j++ {
			if comments[j].ParentID == baseComments[i].ID {
				// foundedID = j
				foundedID = append(foundedID, int64(j))
			}
		}

		if len(foundedID) > 0 {
			for l := 0; l < len(foundedID); l++ {
				baseComments[i].Children = append(baseComments[i].Children, comments[foundedID[l]])
				baseComments[i].TotalChildren = baseComments[i].TotalChildren + 1
				childrenID = append(childrenID, comments[foundedID[l]].ID)
			}
		}

		sort.Slice(baseComments[i].Children, func(j, k int) bool {
			return baseComments[i].Children[j].Content > baseComments[i].Children[k].Content
		})

		comments = append(comments, baseComments[i])

	}
	fmt.Println(comments)

	for k := 0; k < len(childrenID); k++ {
		myTest := func(s int64) bool { return childrenID[k] != s }
		comments = filter(comments, myTest)
	}

	byte, err := json.MarshalIndent((comments), " ", "  ")
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("./test_2.json", byte, 0644)
	if err != nil {
		log.Fatal(err)
	}

}

func filter(ss []Comment, test func(int64) bool) (ret []Comment) {
	for _, s := range ss {
		if test(s.ID) {
			ret = append(ret, s)
		}
	}
	return
}
