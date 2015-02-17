package main

import (
    "strings"    
)

func IsWVUStudent(email string) bool {
    return strings.Split(email, "@")[1] == "mix.wvu.edu"
}

func IsCampaignStaff(email string) bool {
    staffEmails := []string{"absiford@mix.wvu.edu", "whardy1@mix.wvu.edu"}
    for _, v := range staffEmails {
        if email == v {
            return true
        }
    }
    return false
}