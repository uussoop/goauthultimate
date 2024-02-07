# goauthultimate - A Go library for easy and customizable authentication

## Overview

goauthultimate is a Go library that makes it easy to add authentication to your web applications. It handles login, registration, password reset and more out of the box, while allowing high customizability via callback functions.

Key features:

- Supports username/password, email/password, phone/password, phone, email and wallet based authentication
- Customizable via callback functions for user lookup, registration etc.
- Integration with Gin for easy middleware usage
- Password reset flow
- Registration and login endpoints provided

## Usage

```go
import "github.com/gin-gonic/gin"
import "github.com/uussoop/goauthultimate"

func main() {

  // Define callback functions
  utilities := goauthultimate.UtilityFuncs{
    GetUser: func(identifier *string) {}, 
    UserExists: func(identifier *string) bool {},
    ...
  }

  // Create router
  r := gin.Default()

  // Create config
  conf := goauthultimate.AuthConfig{
    Utilities: utilities,
    Router: r,
    Authtype: goauthultimate.Mail,
    Base: "auth",
  }

  // Use middleware
  r.Use(goauthultimate.GoAuthMiddleware(&conf))

  // Start server
  r.Run()
}
```

The key steps are:

1. Define callback functions for user lookup, registration etc.
2. Create a router.
3. Create an AuthConfig with the callbacks and options. 
4. Use the middleware on the router.
5. Start the server.

The library handles the endpoints, validation, error handling etc automatically.

## Authentication Types

The following authentication types are supported via the `Authtype` option:

- `UserPassword` - Username & password based login
- `MailPassword` - Email & password based login
- `PhoneNumberPassword` - Phone number & password based login  
- `Phone` - Phone number based login (sends verification code)
- `Mail` - Email based login (sends verification code)
- `Wallet` - Wallet address based login

## Customization

The key customization points are:

- `UtilityFuncs` - Callback functions for user managementlogic
- `Authtype` - Authentication type to use
- `Base` - Base path for auth routes

Additionally, the middleware and endpoints are highly customizable by modifying the library code itself.

## Conclusion

goauthultimate makes adding authentication to Go applications easy while providing flexibility. Give it a try on your next project!


## TODO

- [ ] Implement SSO
- [ ] Implement password reset endpoint
- [ ] Write tests for middleware
- [ ] Add documentation for customizing error message
- [ ] Add ratelimiting to prevent brute force attacks
- [ ] Audit code for security vulnerabilities
