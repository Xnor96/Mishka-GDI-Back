package application

import (
	"errors"
	"os"
	"time"

	"github.com/Mishka-GDI-Back/domain"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

const (
	accessTokenDuration  = 24 * time.Hour
	refreshTokenDuration = 7 * 24 * time.Hour
)

// TokenPair contiene los tokens generados para el usuario
type TokenPair struct {
	AccessToken  string
	RefreshToken string
	ExpiresIn    int64
	Username     string
	Rol          string
}

// Claims define los campos del JWT
type Claims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Rol      string `json:"rol"`
	Type     string `json:"type"`
	jwt.RegisteredClaims
}

type AuthService interface {
	Login(username, password string) (*TokenPair, error)
	RefreshToken(refreshToken string) (*TokenPair, error)
}

type authService struct {
	usuarioRepo domain.UsuarioRepository
}

func NewAuthService(usuarioRepo domain.UsuarioRepository) AuthService {
	return &authService{usuarioRepo: usuarioRepo}
}

func getJWTSecret() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "mishka-gdi-secret-key-2025"
	}
	return []byte(secret)
}

func generateToken(user *domain.Usuario, tokenType string, duration time.Duration) (string, error) {
	claims := Claims{
		UserID:   user.ID,
		Username: user.Username,
		Rol:      user.Rol,
		Type:     tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "mishka-gdi",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(getJWTSecret())
}

// ValidateToken valida un token JWT y retorna los claims
func ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("método de firma inválido")
		}
		return getJWTSecret(), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("token inválido")
	}
	return claims, nil
}

func (s *authService) Login(username, password string) (*TokenPair, error) {
	usuario, err := s.usuarioRepo.GetByUsername(username)
	if err != nil {
		return nil, errors.New("usuario o contraseña incorrectos")
	}
	if err = bcrypt.CompareHashAndPassword([]byte(usuario.PasswordHash), []byte(password)); err != nil {
		return nil, errors.New("usuario o contraseña incorrectos")
	}
	accessToken, err := generateToken(usuario, "access", accessTokenDuration)
	if err != nil {
		return nil, err
	}
	refreshToken, err := generateToken(usuario, "refresh", refreshTokenDuration)
	if err != nil {
		return nil, err
	}
	return &TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    int64(accessTokenDuration.Seconds()),
		Username:     usuario.Username,
		Rol:          usuario.Rol,
	}, nil
}

func (s *authService) RefreshToken(refreshTokenStr string) (*TokenPair, error) {
	claims, err := ValidateToken(refreshTokenStr)
	if err != nil {
		return nil, errors.New("refresh token inválido o expirado")
	}
	if claims.Type != "refresh" {
		return nil, errors.New("token type incorrecto")
	}
	usuario, err := s.usuarioRepo.GetByID(claims.UserID)
	if err != nil {
		return nil, errors.New("usuario no encontrado")
	}
	accessToken, err := generateToken(usuario, "access", accessTokenDuration)
	if err != nil {
		return nil, err
	}
	return &TokenPair{
		AccessToken: accessToken,
		ExpiresIn:   int64(accessTokenDuration.Seconds()),
		Username:    usuario.Username,
		Rol:         usuario.Rol,
	}, nil
}
