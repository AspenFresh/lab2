package main

import (
    "fmt"
)


type Role struct {
    Name        string   
    Permissions []string 
}


type User struct {
    ID       int
    Name     string
    Email    string
    Password string
    Role     Role
}

func (u User) HasAccess(permission string) bool {
    for _, p := range u.Role.Permissions {
        if p == permission {
            return true
        }
    }
    return false
}


func main() {
    
    adminRole := Role{
        Name:        "Admin",
        Permissions: []string{"create", "read", "update", "delete"},
    }

    userRole := Role{
        Name:        "User",
        Permissions: []string{"read"},
    }

 
    admin := User{
        ID:       1,
        Name:     "Maryna",
        Email:    "maryna@example.com",
        Password: "123456",
        Role:     adminRole,
    }

    user := User{
        ID:       2,
        Name:     "Ivan",
        Email:    "ivan@example.com",
        Password: "qwerty",
        Role:     userRole,
    }

   
    fmt.Printf("Can user %s create? %v\n", admin.Name, admin.HasAccess("create"))
    fmt.Printf("Can user %s delete? %v\n", user.Name, user.HasAccess("delete"))
    fmt.Printf("Can user %s read? %v\n", user.Name, user.HasAccess("read"))
}