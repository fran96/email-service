package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/fran96/email-service/src/docs"
	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Controller struct {
	mailService MailService
}

func NewController(mailService MailService) *Controller {
	return &Controller{mailService: mailService}
}

func (c *Controller) HttpRoutes() {
	r := chi.NewRouter()

	r.HandleFunc("/api", func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case "POST":
			c.post(w, req)
		default:
			http.Error(w, "allowed method: POST", http.StatusMethodNotAllowed)
		}
	})

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	// listen port 8080 for requests
	fmt.Println("listening for requests on port 8080....")
	log.Fatal(http.ListenAndServe(":8080", r))
}

type MailServiceError struct{}

// SendMail godoc
// @Summary Send mail
// @Description Send mail
// @Accept  json
// @Produce  json
// @Param  mail body Mail true "To, Subject, Message"
// @Success 201 {string} string "created"
// @Failure 400,404 {string} fail
// @Failure 500 {string} fail
// @Router /api [post]
func (c *Controller) post(w http.ResponseWriter, req *http.Request) {
	var mail Mail

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(reqBody, &mail); err != nil {
		panic(err)
	}

	_, err = c.mailService.SendMail(mail.To, mail.Subject, mail.Message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}
