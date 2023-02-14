package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pqppq/go-related/frameworks/gin/example/repository"
	"go.uber.org/zap"
)

type BookHandler struct {
	repo   *repository.BookRepo
	logger *zap.Logger
}

func NewBookHandler(repo *repository.BookRepo, logger *zap.Logger) *BookHandler {
	return &BookHandler{
		repo:   repo,
		logger: logger,
	}
}

func writeJsonResponse(c *gin.Context, status int, response gin.H) {
	c.JSON(status, response)
}

func (h *BookHandler) InternalServerError(c *gin.Context) {
	writeJsonResponse(c, http.StatusInternalServerError,
		gin.H{"message": "invalid id"},
	)
}

func (h *BookHandler) InvalidRequest(c *gin.Context) {
	writeJsonResponse(c, http.StatusBadRequest,
		gin.H{"message": "invalid id"},
	)
}

func (h *BookHandler) Success(c *gin.Context) {
	writeJsonResponse(c, http.StatusOK,
		gin.H{"message": "ok"},
	)
}

func (h *BookHandler) ShowBook(c *gin.Context) {
	rowId := c.Param("id")
	id, err := strconv.Atoi(rowId)
	if err != nil {
		h.InternalServerError(c)
		return
	}
	book, err := h.repo.Get(id)
	if err != nil {
		h.logger.Info(
			"failed to get book data",
			zap.Int("id", id))
		h.InternalServerError(c)
		return
	}

	writeJsonResponse(c, http.StatusOK,
		gin.H{"book": book},
	)
}

func (h *BookHandler) ShowBookList(c *gin.Context) {
	books, err := h.repo.GetAll()
	if err != nil {
		h.logger.Info(
			"failed to get all book data")
		h.InternalServerError(c)
		return
	}

	writeJsonResponse(c, http.StatusOK,
		gin.H{"books": books},
	)
}

func (h *BookHandler) AddBook(c *gin.Context) {
	var title string
	err := c.Bind(&title)
	if err != nil {
		h.logger.Info(
			"failed to parse title to insert")
		h.InvalidRequest(c)
		return
	}

	err = h.repo.Create(title)
	if err != nil {
		h.InternalServerError(c)
		return
	}

	h.Success(c)
}

func (h *BookHandler) UpdateBook(c *gin.Context) {
	rowId := c.Param("id")
	title := c.Param("title")
	id, err := strconv.Atoi(rowId)
	if err != nil {
		h.logger.Info(
			"failed to parse id or title to update")
		h.InvalidRequest(c)
		return
	}
	if err != nil {
		h.InvalidRequest(c)
		return
	}

	err = h.repo.Update(id, title)
	if err != nil {
		h.InternalServerError(c)
		return
	}

	h.Success(c)
}

func (h *BookHandler) DeleteBook(c *gin.Context) {
	rowId := c.Param("id")
	id, err := strconv.Atoi(rowId)
	if err != nil {
		h.logger.Info(
			"failed to parse id  to delete")
		h.InvalidRequest(c)
		return
	}

	err = h.repo.Delete(id)
	if err != nil {
		h.InternalServerError(c)
		return
	}

	h.Success(c)
}
