package api

import (
	"auth0_demo/internal/config"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

// _ "github.com/byron-ojua/starter-project/internal/api/docs"

//go:generate swag init --parseDependency --parseInternal

// @title Auth0 Demo API
// @version 1.0
// @description This is the API for the Auth0 demo
// @termsOfService TBD
//
// @contact.name Byron Ojua-Nice
// @contact.url http://firstlaunch.dev
// @contact.email byronojua@firstlaunch.dev
//
// @license.name TBD
// @license.url TBD
//
// @host localhost:8080
// @BasePath /
//
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Use the format: `Bearer <your_token>` (Bearer must be added before the token)
//
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

const API_VERSION = "1.0"

// error declarations
const (
	//the error returned to the user when bad user token being provided
	ErrUnauthorized = "bad user token"

	//the error returned to the user when no token is provided
	ErrNoToke = "no authentication token provided"

	//the error returned to the user when a user is not authorized to perform an action
	ErrForbidden = "forbidden"

	//the error returned when no resource ID is provided for a request pertaining to a specific resource
	ErrorNoResourceID = "no resource id provided"

	//the error returned to the user when the requested record in out of date
	ErrCalibrationOutOfDate = "out of date record"

	//error signifying that a feature is not supported for a specific product
	ErrUnsupportedFeature = "unsupported product feature"

	//err signifying that a vehicle does not have an attached device
	ErrNoVehicleDevice = "no device attached to vehicle"
)

// ErrorsResponse represents the structure for API error responses.
type ErrorsResponse struct {
	Errors []string `json:"errors"` // Array of error messages
}

// Api is the interface for the API
type Api interface {
	RunLocal() error                    // Run the API locally
	RunAzureFunction(port string) error // Run the API as an Azure Function
}

// env is the environment for the API
type env struct {
	logger    *zap.SugaredLogger  `validate:"required"` // Logger
	validator *validator.Validate `validate:"required"` // Validator
	config    *config.Config      `validate:"required"` // Config
	api       *gin.Engine         `validate:"required"` // API
}

// RunLocal runs the API locally
func (e *env) RunLocal() error {
	return http.ListenAndServe("localhost:8080", e.api)
}

// Add this method to your env struct implementation
func (e *env) RunAzureFunction(port string) error {
	// Configure your server for Azure Functions
	return http.ListenAndServe(":"+port, e.api)
}

// New creates a new API instance
func New(logger *zap.SugaredLogger, cfg *config.Config) (Api, error) {
	// Create a new validator
	validator := validator.New()

	e := &env{
		validator: validator,
		logger:    logger,
		config:    cfg,
	}

	r := gin.Default()

	// Define route prefix (empty for local, "/api" for Azure Functions)
	routePrefix := ""
	if _, exists := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); exists {
		routePrefix = "/api"
	}

	//CORS setup
	r.Use(cors.New(cors.Config{
		AllowHeaders:    []string{"Authorization", "Content-Type"},
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
	}))

	r.GET(routePrefix+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Setup the routes. Make inline return with message "hello from server 2"
	r.GET(routePrefix+"/ping", func(c *gin.Context) {
		logger.Info("ping request received")
		c.JSON(http.StatusOK, gin.H{
			"message": "hello from server 2",
		})
	})

	// Protected endpoints
	// Replace "your-secret-key" with the actual secret key used to sign your JWTs.
	// protected := r.Group("/protected")
	// protected.Use(e.AuthMiddleware())
	// {
	// 	protected.GET("/profile", func(c *gin.Context) {
	// 		// Here you can retrieve claims from the context if needed:
	// 		claims, _ := c.Get("claims")
	// 		c.JSON(http.StatusOK, gin.H{
	// 			"message": "This is a protected endpoint.",
	// 			"claims":  claims,
	// 		})
	// 	})
	// }

	e.api = r

	return e, nil
}

// AuthMiddleware returns a Gin middleware that checks for a valid Bearer token.
// Replace 'your-secret-key' with your actual secret or use your configuration.
// func (e *env) AuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// Get the Authorization header value
// 		authHeader := c.GetHeader("Authorization")
// 		if authHeader == "" {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": ErrNoToke})
// 			return
// 		}

// 		// Split the header to get the token part
// 		parts := strings.Split(authHeader, " ")
// 		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": ErrUnauthorized})
// 			return
// 		}

// 		tokenString := parts[1]

// 		// Parse and validate the token
// 		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 			// Validate the algorithm, for example using HMAC
// 			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 			}

// 			return []byte(env.config), nil
// 		})

// 		if err != nil || !token.Valid {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
// 			return
// 		}

// 		// Optionally, you can set token claims in the context for later use
// 		c.Set("claims", token.Claims)
// 		c.Next()
// 	}
// }
