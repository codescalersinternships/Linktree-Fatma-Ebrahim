package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/codescalersinternships/Linktree-Fatma-Ebrahim/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUserAPIs(t *testing.T) {
	newtime := fmt.Sprintf("%v", time.Now())
	t.Run("test signup with new user", func(t *testing.T) {
		router := Linktreeserver()
		response := httptest.NewRecorder()

		user := models.User{
			Username: "testuser" + newtime,
			Email:    "testuser@example.com",
			Password: "test",
		}

		userbytes, err := json.Marshal(user)

		request, err := http.NewRequest("POST", "/linktree/signup", bytes.NewReader(userbytes))
		if err != nil {
			t.Fatal(err)
		}

		request.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(response, request)
		assert.Equal(t, http.StatusCreated, response.Code)

	})

	t.Run("test signup with exisiting user", func(t *testing.T) {
		router := Linktreeserver()
		response := httptest.NewRecorder()
		user := models.User{
			Username: "testuser" + newtime,
			Email:    "testuser@example.com",
			Password: "test",
		}

		userbytes, err := json.Marshal(user)

		request, err := http.NewRequest("POST", "/linktree/signup", bytes.NewReader(userbytes))
		if err != nil {
			t.Fatal(err)
		}

		request.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(response, request)
		assert.Equal(t, http.StatusBadRequest, response.Code)

	})

	t.Run("test login with existing user", func(t *testing.T) {
		router := Linktreeserver()
		response := httptest.NewRecorder()
		user := models.User{
			Username: "testuser" + newtime,
			Email:    "testuser@example.com",
			Password: "test",
		}

		userbytes, err := json.Marshal(user)

		request, err := http.NewRequest("POST", "/linktree/login", bytes.NewReader(userbytes))
		if err != nil {
			t.Fatal(err)
		}

		request.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(response, request)
		assert.Equal(t, http.StatusOK, response.Code)

	})

	t.Run("test login with new user", func(t *testing.T) {
		newtime := fmt.Sprintf("%v", time.Now())
		router := Linktreeserver()
		response := httptest.NewRecorder()
		user := models.User{
			Username: "testuser" + newtime,
			Email:    "testuser@example.com",
			Password: "test",
		}

		userbytes, err := json.Marshal(user)

		request, err := http.NewRequest("POST", "/linktree/login", bytes.NewReader(userbytes))
		if err != nil {
			t.Fatal(err)
		}

		request.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(response, request)
		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

}

func TestTreeAPIs(t *testing.T) {
	user, err := signupHelper()
	// tree := models.Linktree{}

	if err != nil {
		t.Fatal(err)
	}
	t.Run("test add linktree with authorized user", func(t *testing.T) {
		router := Linktreeserver()
		response := httptest.NewRecorder()
		tree := models.Linktree{
			Fullname: "Test User",
			Bio:      "This is a test bio",
			Links:    []models.Link{},
		}
		treebytes, err := json.Marshal(tree)

		request, err := http.NewRequest("POST", "/linktree", bytes.NewReader(treebytes))
		if err != nil {
			t.Fatal(err)
		}

		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("token", user.Token)
		router.ServeHTTP(response, request)

		err = json.Unmarshal(response.Body.Bytes(), &user)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, http.StatusCreated, response.Code)
	})

	t.Run("test get linktree by id with authorized user", func(t *testing.T) {
		tree := models.Linktree{
			Fullname: "Test User",
			Bio:      "This is a test bio",
			Links:    []models.Link{},
		}
		router := Linktreeserver()
		response := httptest.NewRecorder()
		tree_id := user.LinkTreeID.Hex()
		request, err := http.NewRequest("GET", "/linktree/"+tree_id, nil)
		if err != nil {
			t.Fatal(err)
		}
		request.Header.Set("token", user.Token)

		router.ServeHTTP(response, request)
		gottree := models.Linktree{}
		err = json.Unmarshal(response.Body.Bytes(), &gottree)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, tree.Bio, gottree.Bio)
		assert.Equal(t, tree.Fullname, gottree.Fullname)
	})
	t.Run("test add linktree with unauthorized user", func(t *testing.T) {
		router := Linktreeserver()
		response := httptest.NewRecorder()
		tree := models.Linktree{
			Fullname: "Test User",
			Bio:      "This is a test bio",
			Links:    []models.Link{},
		}
		treebytes, err := json.Marshal(tree)

		request, err := http.NewRequest("POST", "/linktree", bytes.NewReader(treebytes))
		if err != nil {
			t.Fatal(err)
		}
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("token", "testtoken")
		router.ServeHTTP(response, request)
		got := gin.H{}
		err = json.Unmarshal(response.Body.Bytes(), &got)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, got["error"], "Invalid token")
		assert.Equal(t, http.StatusUnauthorized, response.Code)
	})

	t.Run("test get linktree by id with unauthorized user", func(t *testing.T) {
		router := Linktreeserver()
		response := httptest.NewRecorder()
		tree_id := user.LinkTreeID.Hex()
		request, err := http.NewRequest("GET", "/linktree/"+tree_id, nil)
		if err != nil {
			t.Fatal(err)
		}
		request.Header.Set("token", "testtoken")

		router.ServeHTTP(response, request)
		got := gin.H{}
		err = json.Unmarshal(response.Body.Bytes(), &got)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, got["error"], "Invalid token")
		assert.Equal(t, http.StatusUnauthorized, response.Code)
	})

}

func TestLinkAPIs(t *testing.T) {
	user, err := signupHelper()

	if err != nil {
		t.Fatal(err)
	}
	_, err = addtreeHelper(&user)
	if err != nil {
		t.Fatal(err)
	}
	t.Run("test add links to tree with authorized user", func(t *testing.T) {
		router := Linktreeserver()
		response := httptest.NewRecorder()
		link := models.Link{
			Name:   "Google",
			Link:   "https://www.google.com/",
			Visits: 0,
		}

		linkbytes, err := json.Marshal(link)
		if err != nil {
			t.Fatal(err)
		}

		endpoint := fmt.Sprintf("/linktree/%s/addlink", user.LinkTreeID.Hex())
		request, err := http.NewRequest("POST", endpoint, bytes.NewReader(linkbytes))
		if err != nil {
			t.Fatal(err)
		}
		request.Header.Set("Content-Type", "application/json")

		request.Header.Set("token", user.Token)
		router.ServeHTTP(response, request)
		assert.Equal(t, http.StatusCreated, response.Code)
	})

	t.Run("test add links to tree with unauthorized user", func(t *testing.T) {
		router := Linktreeserver()
		response := httptest.NewRecorder()
		link := models.Link{}
		linkbytes, err := json.Marshal(link)
		if err != nil {
			t.Fatal(err)
		}

		endpoint := fmt.Sprintf("/linktree/%s/addlink", user.LinkTreeID.Hex())
		request, err := http.NewRequest("POST", endpoint, bytes.NewReader(linkbytes))
		if err != nil {
			t.Fatal(err)
		}
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("token", "testtoken")
		router.ServeHTTP(response, request)
		got := gin.H{}
		err = json.Unmarshal(response.Body.Bytes(), &got)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, got["error"], "Invalid token")
		assert.Equal(t, http.StatusUnauthorized, response.Code)
	})

	t.Run("test add links to tree with invalid tree id", func(t *testing.T) {
		router := Linktreeserver()
		response := httptest.NewRecorder()
		link := models.Link{}
		linkbytes, err := json.Marshal(link)
		if err != nil {
			t.Fatal(err)
		}

		endpoint := fmt.Sprintf("/linktree/%s/addlink", primitive.NewObjectID().Hex())
		request, err := http.NewRequest("POST", endpoint, bytes.NewReader(linkbytes))
		if err != nil {
			t.Fatal(err)
		}
		request.Header.Set("Content-Type", "application/json")

		request.Header.Set("token", user.Token)
		router.ServeHTTP(response, request)
		got := gin.H{}
		err = json.Unmarshal(response.Body.Bytes(), &got)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, got["message"], "mongo: no documents in result")
		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("test update bio to tree with authorized user", func(t *testing.T) {
		router := Linktreeserver()
		response := httptest.NewRecorder()
		bio := "New Bio"

		biobytes, err := json.Marshal(bio)
		if err != nil {
			t.Fatal(err)
		}

		endpoint := fmt.Sprintf("/linktree/%s/addbio", user.LinkTreeID.Hex())
		request, err := http.NewRequest("POST", endpoint, bytes.NewReader(biobytes))
		if err != nil {
			t.Fatal(err)
		}
		request.Header.Set("Content-Type", "application/json")

		request.Header.Set("token", user.Token)
		router.ServeHTTP(response, request)
		gottree := models.Linktree{}
		err = json.Unmarshal(response.Body.Bytes(), &gottree)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, bio, gottree.Bio)
	})

	t.Run("test update fullname to tree with authorized user", func(t *testing.T) {
		router := Linktreeserver()
		response := httptest.NewRecorder()
		fullname := "New Fullname"

		fullnamebytes, err := json.Marshal(fullname)
		if err != nil {
			t.Fatal(err)
		}

		endpoint := fmt.Sprintf("/linktree/%s/addfullname", user.LinkTreeID.Hex())
		request, err := http.NewRequest("POST", endpoint, bytes.NewReader(fullnamebytes))
		if err != nil {
			t.Fatal(err)
		}
		request.Header.Set("Content-Type", "application/json")

		request.Header.Set("token", user.Token)
		router.ServeHTTP(response, request)
		gottree := models.Linktree{}
		err = json.Unmarshal(response.Body.Bytes(), &gottree)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, fullname, gottree.Fullname)
	})

	t.Run("test edit links with authorized user", func(t *testing.T) {
		tree, err := gettreeHelper(&user)
		if err != nil {
			t.Fatal(err)
		}
		router := Linktreeserver()
		response := httptest.NewRecorder()

		link := models.Link{
			ID:     tree.Links[0].ID,
			Name:   "facebook",
			Link:   "https://www.facebook.com/",
			Visits: 0,
		}

		linkbytes, err := json.Marshal(link)
		if err != nil {
			t.Fatal(err)
		}

		endpoint := fmt.Sprintf("/linktree/%s/updatelink", user.LinkTreeID.Hex())
		request, err := http.NewRequest("POST", endpoint, bytes.NewReader(linkbytes))
		if err != nil {
			t.Fatal(err)
		}
		request.Header.Set("Content-Type", "application/json")

		request.Header.Set("token", user.Token)
		router.ServeHTTP(response, request)
		gottree := models.Linktree{}
		err = json.Unmarshal(response.Body.Bytes(), &gottree)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, link.Name, gottree.Links[0].Name)
		assert.Equal(t, link.Link, gottree.Links[0].Link)

	})

	var deletedlink primitive.ObjectID
	t.Run("test delete link with authorized user", func(t *testing.T) {
		tree, err := gettreeHelper(&user)
		deletedlink = tree.Links[0].ID
		if err != nil {
			t.Fatal(err)
		}
		router := Linktreeserver()
		response := httptest.NewRecorder()

		link := models.Link{
			ID: tree.Links[0].ID,
		}
		linkbytes, err := json.Marshal(link)
		if err != nil {
			t.Fatal(err)
		}

		endpoint := fmt.Sprintf("/linktree/%s/deletelink", user.LinkTreeID.Hex())
		request, err := http.NewRequest("DELETE", endpoint, bytes.NewReader(linkbytes))
		if err != nil {
			t.Fatal(err)
		}
		request.Header.Set("Content-Type", "application/json")

		request.Header.Set("token", user.Token)
		router.ServeHTTP(response, request)
		gottree := models.Linktree{}
		err = json.Unmarshal(response.Body.Bytes(), &gottree)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, 2, len(gottree.Links))
		assert.NotContains(t, gottree.Links, link)

	})
	t.Run("test delete previously deletedlink with authorized user", func(t *testing.T) {
		router := Linktreeserver()
		response := httptest.NewRecorder()

		link := models.Link{
			ID: deletedlink,
		}

		linkbytes, err := json.Marshal(link)
		if err != nil {
			t.Fatal(err)
		}

		endpoint := fmt.Sprintf("/linktree/%s/deletelink", user.LinkTreeID.Hex())
		request, err := http.NewRequest("DELETE", endpoint, bytes.NewReader(linkbytes))
		if err != nil {
			t.Fatal(err)
		}
		request.Header.Set("Content-Type", "application/json")

		request.Header.Set("token", user.Token)
		router.ServeHTTP(response, request)
		got := gin.H{}
		err = json.Unmarshal(response.Body.Bytes(), &got)
		if err != nil {
			t.Fatal(err)
		}
		tree, err := gettreeHelper(&user)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, got["message"], "link not found in linktree")
		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, 2, len(tree.Links))

	})
}

func signupHelper() (models.User, error) {
	newtime := fmt.Sprintf("%v", time.Now())
	router := Linktreeserver()
	response := httptest.NewRecorder()

	user := models.User{
		Username: "testuser" + newtime,
		Email:    "testuser@example.com",
		Password: "test",
	}

	userbytes, err := json.Marshal(user)
	request, err := http.NewRequest("POST", "/linktree/signup", bytes.NewReader(userbytes))
	if err != nil {
		return models.User{}, err
	}
	request.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(response, request)

	if response.Code != http.StatusCreated {
		return models.User{}, fmt.Errorf("signup failed")
	}
	err = json.Unmarshal(response.Body.Bytes(), &user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func addtreeHelper(user *models.User) (models.Linktree, error) {

	router := Linktreeserver()
	response := httptest.NewRecorder()
	tree := models.Linktree{
		Fullname: "Test User",
		Bio:      "This is a test bio",
		Links: []models.Link{
			{
				Name:   "link1",
				Link:   "https://www.link1.com/",
				Visits: 5,
			},
			{
				Name:   "link2",
				Link:   "https://www.link2.com/",
				Visits: 5,
			},
		},
	}
	treebytes, err := json.Marshal(tree)

	request, err := http.NewRequest("POST", "/linktree", bytes.NewReader(treebytes))
	if err != nil {
		return models.Linktree{}, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("token", user.Token)
	router.ServeHTTP(response, request)

	err = json.Unmarshal(response.Body.Bytes(), &user)
	if err != nil {
		return models.Linktree{}, err
	}
	return tree, nil
}

func gettreeHelper(user *models.User) (models.Linktree, error) {

	router := Linktreeserver()
	response := httptest.NewRecorder()
	tree_id := user.LinkTreeID.Hex()
	request, err := http.NewRequest("GET", "/linktree/"+tree_id, nil)
	if err != nil {
		return models.Linktree{}, err
	}
	request.Header.Set("token", user.Token)

	router.ServeHTTP(response, request)
	gottree := models.Linktree{}
	err = json.Unmarshal(response.Body.Bytes(), &gottree)
	if err != nil {
		return models.Linktree{}, err
	}
	return gottree, nil
}
