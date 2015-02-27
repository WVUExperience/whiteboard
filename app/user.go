package main

import (
    "appengine"
    "appengine/user"
    "strings"    
)

type EmbeddedUser struct {
    Name, Email, ID, Logout string
}

func GetEmbeddedUser(u *user.User, c appengine.Context) *EmbeddedUser {
    if u == nil {
        return nil
    } else {
        logout, _ := user.LogoutURL(c, "/")
        return &EmbeddedUser{
            Name: u.String(),
            Email: u.Email,
            ID: u.ID,
            Logout: logout,
        }
    }
}

func IsWVUStudent(email string) bool {
    return strings.Split(email, "@")[1] == "mix.wvu.edu"
}

func IsCampaignStaff(email string) bool {
    staffEmails := []string{"absiford@mix.wvu.edu", "whardy1@mix.wvu.edu",
    "adsutherland@mix.wvu.edu", "kalaska@mix.wvu.edu", "ilkovtoniuk@mix.wvu.edu",
    "jtheenan@mix.wvu.edu", "reelkins@mix.wvu.edu", "krash@mix.wvu.edu",
    "cmorlock@mix.wvu.edu", "sricha12@mix.wvu.edu", "tcmccloud@mix.wvu.edu",
    "aionderik@mix.wvu.edu", "jleach6@mix.wvu.edu", "jamcguire@mix.wvu.edu",
    "mwrogers@mix.wvu.edu", "sjcrandall@mix.wvu.edu", "aherric1@mix.wvu.edu",
    "jlriseberg@mix.wvu.edu", "nrmcdill@mix.wvu.edu", "clammer1@mix.wvu.edu",
    "rlburky@mix.wvu.edu", "aleccneu@gmail.com", "scottbraxton2015@gmail.com"}
    for _, v := range staffEmails {
        if email == v {
            return true
        }
    }
    return false
}