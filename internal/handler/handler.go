package handler

import (
	"net/http"

	"github.com/a6slab/shorter_url/internal/dto"
	"github.com/a6slab/shorter_url/internal/repository"
	"github.com/a6slab/shorter_url/internal/utils"
	"github.com/gin-gonic/gin"
)

type ShorterHandler interface {
	ShortenURL(ctx *gin.Context)
	GetRedirectURL(ctx *gin.Context)

	HealthCheck(ctx *gin.Context)
}

type shorterHandler struct {
	repo repository.ShorterURLRepository
}

func NewShorterHandler(repo *repository.ShorterURLRepository) ShorterHandler {
	return &shorterHandler{
		repo: *repo,
	}
}

func (s *shorterHandler) HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"ping": "pong"})
}

func (s *shorterHandler) ShortenURL(ctx *gin.Context) {
	var request dto.ShortenRequest
	ctx.BindJSON(&request)

	shortURL := utils.GenerateShortURL(request.LongURL)
	data, err := s.repo.ShortenURL(request.LongURL, shortURL, request.ExpiryDate, request.UserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to shorten url"})
	}

	ctx.JSON(http.StatusCreated, gin.H{"result": data})
}

func (s *shorterHandler) GetRedirectURL(ctx *gin.Context) {
	shortURLKey := ctx.Param("short_url_key")

	redirectURL, err := s.repo.GetRedirectURL(shortURLKey)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "not found redirect url"})
	}
	ctx.JSON(http.StatusOK, gin.H{"redirect_url": redirectURL})
}
