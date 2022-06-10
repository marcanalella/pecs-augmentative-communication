package auth

import (
	err "errors"
	"net/http"
	"time"

	"github.com/pecs/pecs-be/internal/entity"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/pecs/pecs-be/internal/errors"

	log "github.com/sirupsen/logrus"
)

type TokenDetails struct {
	AccessToken  string    `json:"accessToken"`
	RefreshToken string    `json:"-"`
	AccessUUID   uuid.UUID `json:"-"`
	RefreshUUID  uuid.UUID `json:"-"`
	AtExpires    int64     `json:"-"`
	RtExpires    int64     `json:"-"`
}

//generate access and refresh jwt tokens
func (s service) generateJWT(identity Identity) (*TokenDetails, error) {
	td := &TokenDetails{}

	td.AtExpires = time.Now().Add(time.Minute * 100000000).Unix() //TODO externalise expiration time access token
	td.AccessUUID, _ = uuid.NewUUID()

	td.RefreshUUID, _ = uuid.NewUUID()
	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix() //TODO externalise expiration time refresh token

	//access token
	atClaims := jwt.MapClaims{
		"id":          identity.GetID(),
		"iss":         "pecs",
		"sub":         identity.GetName(),
		"exp":         td.AtExpires,
		"access_uuid": td.AccessUUID,
		"role":        identity.GetRole(),
	}

	aToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims).SignedString([]byte(s.signinAccessTokenKey))

	if err != nil {
		return nil, err
	}

	td.AccessToken = aToken

	//refresh token
	rtClaims := jwt.MapClaims{
		"refresh_uuid": td.RefreshUUID,
		"id":           identity.GetID(),
		"exp":          td.RtExpires,
	}

	rToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims).SignedString([]byte(s.signinRefreshTokenKey)) //TODO Externalise this Key

	if err != nil {
		return nil, err
	}

	td.RefreshToken = rToken

	return td, nil
}

func verifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := r.Header.Get("Authorization")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//check if the token signing method is "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Errorf("Unexpected signing method: %v", token.Header["alg"])
			return nil, errors.NewHTTPError(err.New("Signin method is invalid"), http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), http.StatusText(http.StatusInternalServerError), "Unexpected signing method")
		}

		return []byte("myatkey"), nil //TODO get through s.signinkaccesstokenkey
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func TokenValid(r *http.Request) (*jwt.Token, error) {
	token, err := verifyToken(r)

	if err != nil {
		return nil, err
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return nil, err
	}

	return token, nil
}

// RefreshToken get the refresh token from the body and return a new jwt token
func (s service) RefreshToken(r *http.Request) (*TokenDetails, error) {
	refreshToken, err := r.Cookie("refresh_token")

	if err != nil {
		return nil, errors.NewHTTPError(err, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), http.StatusText(http.StatusInternalServerError), "Unable to read refresh token")
	}

	token, err := jwt.Parse(refreshToken.Value, func(token *jwt.Token) (interface{}, error) {
		//check if the token signing method is "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Errorf("Unexpected signing method: %v", token.Header["alg"])
			return nil, errors.NewHTTPError(nil, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), http.StatusText(http.StatusInternalServerError), "Unexpected signing method")
		}

		return []byte(s.signinRefreshTokenKey), nil
	})

	if err != nil {
		log.Infof("Refresh token expired %v", err)
		return nil, errors.NewHTTPError(err, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), http.StatusText(http.StatusUnauthorized), "Refresh token is not valid or expired")
	}

	//is token valid?
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		log.Error("Refresh token is not valid")
		return nil, errors.NewHTTPError(nil, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), http.StatusText(http.StatusUnauthorized), "Refresh token is not valid")
	}

	claims, ok := token.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims

	if ok && token.Valid {
		//Get refresh UUID
		_, ok := claims["refresh_uuid"].(string)
		if !ok {
			log.Error("failed to take refreshUuid")
			return nil, nil //TODO HANDLE ERROR
		}

		//Get user refreshUUID
		refreshUuid, ok := claims["refresh_uuid"].(string)
		log.Infof("refreshUuid %v", refreshUuid)

		if !ok {
			log.Error("Failed to take refreshUuid")
			return nil, errors.NewHTTPError(nil, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), http.StatusText(http.StatusInternalServerError), "Failed to take refreshUuid")
		}

		//todo Delete the previous Refresh Token with refreshUuid

		//Get user UUID
		id, ok := claims["id"].(string)
		log.Infof("id %v", id)

		if !ok {
			log.Error("Failed to take id")
			return nil, errors.NewHTTPError(nil, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), http.StatusText(http.StatusInternalServerError), "Failed to take id")
		}

		//userID, err := uuid.FromBytes([]byte(id)) //convert string to uuid
		userID, err := uuid.Parse(id)
		if err != nil {
			//log.Error("Invalid UUID %v", err)
			return nil, errors.NewHTTPError(err, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), http.StatusText(http.StatusInternalServerError), "Invalid UUID")
		}

		var user entity.User
		var ts *TokenDetails

		user, err = s.repo.GetById(userID)

		//Create new pairs of refresh and access tokens
		ts, err = s.generateJWT(user)

		if err != nil {
			log.Info("failed to refresh jwt")
			return nil, err
		}
		return ts, nil
	}

	return nil, errors.NewHTTPError(err, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), http.StatusText(http.StatusUnauthorized), "Token is expired")
}
