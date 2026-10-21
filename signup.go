package main

import (
	"fmt"
	"net/http"
	"golang.org/x/crypto/bcrypt"
)

func (s *Server) signup(w http.ResponseWriter, r *http.Request) {

    if r.Method == http.MethodGet {
        fmt.Fprintln(w, `
        <!DOCTYPE html>
        <html>
        <head>
            <title>Sign Up</title>
        </head>

        <body>

            <h1>Create Account</h1>

            <form action="/signup" method="POST">

                <label>Username</label>
                <input type="text" name="username">

                <br><br>

                <label>Email</label>
                <input type="email" name="email">

                <br><br>

                <label>Password</label>
                <input type="password" name="password">

                <br><br>

                <button type="submit">Sign Up</button>

            </form>

        </body>
        </html>
        `)

        return
    }

    if r.Method != http.MethodPost {
        http.Error(
            w,
            "Method not allowed",
            http.StatusMethodNotAllowed,
        )
        return
    }

    username := r.FormValue("username")
    email := r.FormValue("email")
    password := r.FormValue("password")

    passwordHash, err := bcrypt.GenerateFromPassword(
        []byte(password),
        bcrypt.DefaultCost,
    )

    if err != nil {
        http.Error(
            w,
            "Unable to create account",
            http.StatusInternalServerError,
        )
        return
    }

    err = s.CreateUser(
        username,
        email,
        string(passwordHash),
    )

    if err != nil {
        http.Error(
            w,
            "Unable to create account",
            http.StatusInternalServerError,
        )
        return
    }

    fmt.Fprintln(w, "Account created successfully")
}