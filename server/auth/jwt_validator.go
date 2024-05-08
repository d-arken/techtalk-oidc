package auth

import (
	"context"
	"errors"
	"github.com/coreos/go-oidc"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

const OIDCClaimsContext = "OIDCClaims"

var ErrMissingAccessToken = errors.New("missing access token")

type OIDCClaims struct {
	Sub string `json:"sub"`
}

var config = &oidc.Config{
	ClientID:          os.Getenv("OIDC_CLIENT_ID"),
	SkipClientIDCheck: true,
}

var v = oidc.NewVerifier("https://darkmdev.kinde.com", oidc.NewRemoteKeySet(context.Background(), "https://darkmdev.kinde.com/.well-known/jwks"), config)


func JWTProtected() gin.HandlerFunc {
	return func(c *gin.Context) {

		accessToken := c.GetHeader("Authorization")

		if accessToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": ErrMissingAccessToken.Error()})
			c.Abort()
			return
		}

		idToken, err := v.Verify(c, accessToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			c.Abort()
			return
		}

		var claims OIDCClaims
		err = idToken.Claims(&claims)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			c.Abort()
			return
		}

		c.Set(OIDCClaimsContext, claims)
		c.Next()
	}
}
