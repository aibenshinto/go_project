package crud

func InsertPhone(w http.ResponseWriter, r *http.Request) {
    var phone database.MobilePhone
    err := json.NewDecoder(r.Body).Decode(&phone)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Convert Images slice to JSON string
    imagesJSON, err := json.Marshal(phone.Images)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Insert phone into database
    _, err = database.DB.Exec("INSERT INTO phones (name, spec, images, price) VALUES (?, ?, ?, ?)", phone.Name, phone.Spec, string(imagesJSON), phone.Price)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Send response indicating success
    w.WriteHeader(http.StatusCreated)
}