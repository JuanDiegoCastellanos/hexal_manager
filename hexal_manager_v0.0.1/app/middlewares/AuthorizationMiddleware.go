package middlewares

//
//import (
//	"github.com/labstack/echo/v4"
//	"net/http"
//	"strings"
//)
//
//func Authorize(next echo.HandlerFunc) echo.HandlerFunc {
//	return func(c echo.Context) error {
//		token := c.Request().Header.Get("x-token")
//		if token == "" {
//			return c.JSON(http.StatusUnauthorized, "Token not provided")
//		}
//		// Aquí podrías implementar la lógica para validar el token, por ejemplo,
//		// verificar que es válido y que el usuario que hace la solicitud tiene
//		// permisos suficientes para acceder a los recursos solicitados.
//
//		return next(c)
//	}
//}
//
//type JWTClaims struct {
//	UserID uint `json:"user_id"`
//	jwt.StandardClaims
//}
//
//func Authorize(next echo.HandlerFunc) echo.HandlerFunc {
//	return func(c echo.Context) error {
//		// Get the token from the authorization header
//		authHeader := c.Request().Header.Get("Authorization")
//		if authHeader == "" {
//			return c.JSON(http.StatusUnauthorized, "Missing authorization header")
//		}
//		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
//
//		// Parse the token and validate it
//		token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
//			// Replace with your secret key
//			return []byte("my_secret_key"), nil
//		})
//		if err != nil {
//			return c.JSON(http.StatusUnauthorized, err.Error())
//		}
//
//		// Check if the token is valid
//		if !token.Valid {
//			return c.JSON(http.StatusUnauthorized, "Invalid token")
//		}
//
//		// Get the claims from the token
//		claims := token.Claims.(*JWTClaims)
//
//		// Check if the user has enough permissions to access the resource
//		if !hasPermissions(claims.UserID, c.Path(), c.Request().Method) {
//			return c.JSON(http.StatusForbidden, "Insufficient permissions")
//		}
//
//		// Set the user ID in the context for the next handler
//		c.Set("user_id", claims.UserID)
//
//		// Call the next handler
//		return next(c)
//	}
//}
//
//// This function should be implemented to check if the user has enough permissions to access the resource
//func hasPermissions(userID uint, path, method string) bool {
//	// Replace with your logic to check the user permissions
//	return true
//}
