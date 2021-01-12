package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	contact "github.com/Akshit8/go-boilerplate"
	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson"
)

var mh *contact.MongoHandler

func registerRoutes() http.Handler {
	r := chi.NewRouter()
	r.Route("/contacts", func(r chi.Router) {
		r.Get("/", getAllContact)
		r.Get("/{phonenumber}", getContact)
		r.Post("/", addContact)
		r.Put("/{phonenumber}", updateContact)
		r.Delete("/{phonenumber}", deleteContact)
	})
	return r
}

func main() {
	// mongoDBconnectionString := os.Getenv("MONGODB_CONNECTION_STRING")
	mongoDBconnectionString := "mongodb+srv://akshit:akshit2853@cluster0.y2ty6.mongodb.net/contactdb?retryWrites=true&w=majority"
	mh = contact.NewHandler(mongoDBconnectionString)
	r := registerRoutes()
	log.Fatal(http.ListenAndServe(":8000", r))
}

func getContact(w http.ResponseWriter, r *http.Request) {
	phoneNumber := chi.URLParam(r, "phonenumber")
	if phoneNumber == "" {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	contact := &contact.Contact{}
	err := mh.GetOne(contact, bson.M{"phoneNumber": phoneNumber})
	if err != nil {
		http.Error(w, fmt.Sprintf("Contact with phonenumber: %s not found", phoneNumber), 404)
		return
	}
	json.NewEncoder(w).Encode(contact)
}

func getAllContact(w http.ResponseWriter, r *http.Request) {
	contacts := mh.Get(bson.M{})
	json.NewEncoder(w).Encode(contacts)
}

func addContact(w http.ResponseWriter, r *http.Request) {
	existingContact := &contact.Contact{}
	var contact contact.Contact
	json.NewDecoder(r.Body).Decode(&contact)
	contact.CreatedOn = time.Now()
	err := mh.GetOne(existingContact, bson.M{"phoneNumber": contact.PhoneNumber})
	if err == nil {
		http.Error(w, fmt.Sprintf("Contact with phonenumber: %s already exist", contact.PhoneNumber), 400)
		return
	}
	_, err = mh.AddOne(&contact)
	if err != nil {
		http.Error(w, fmt.Sprint(err), 400)
		return
	}
	w.Write([]byte("Contact created successfully"))
	w.WriteHeader(201)
}

func deleteContact(w http.ResponseWriter, r *http.Request) {
	existingContact := &contact.Contact{}
	phoneNumber := chi.URLParam(r, "phonenumber")
	if phoneNumber == "" {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	err := mh.GetOne(existingContact, bson.M{"phoneNumber": phoneNumber})
	if err != nil {
		http.Error(w, fmt.Sprintf("Contact with phonenumber: %s does not exist", phoneNumber), 400)
		return
	}
	_, err = mh.RemoveOne(bson.M{"phoneNumber": phoneNumber})
	if err != nil {
		http.Error(w, fmt.Sprint(err), 400)
		return
	}
	w.Write([]byte("Contact deleted"))
	w.WriteHeader(200)
}

func updateContact(w http.ResponseWriter, r *http.Request) {
	phoneNumber := chi.URLParam(r, "phonenumber")
	if phoneNumber == "" {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	contact := &contact.Contact{}
	json.NewDecoder(r.Body).Decode(contact)
	_, err := mh.Update(bson.M{"phoneNumber": phoneNumber}, contact)
	if err != nil {
		http.Error(w, fmt.Sprint(err), 400)
		return
	}
	w.Write([]byte("Contact update successful"))
	w.WriteHeader(200)
}
