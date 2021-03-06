package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mbredikhin/snippets"
)

func (h *Handler) createSnippet(c *gin.Context) {
	listID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid list id parameter")
		return
	}
	var input snippets.Snippet
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Snippet.Create(listID, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllSnippets(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	listID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid list id parameter")
		return
	}
	snippets, err := h.services.Snippet.GetAll(userID, listID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, snippets)
}

func (h *Handler) getSnippet(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	snippetID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id parameter")
		return
	}
	snippet, err := h.services.Snippet.GetByID(userID, snippetID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, snippet)
}

func (h *Handler) updateSnippet(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	snippetID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id parameter")
		return
	}
	var input snippets.UpdateSnippetInput
	if c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = h.services.Snippet.Update(userID, snippetID, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) deleteSnippet(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	snippetID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid snippet id parameter")
		return
	}
	if err := h.services.Snippet.Delete(userID, snippetID); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) getFavouriteSnippets(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	snippetIDs, err := h.services.Snippet.GetFavourites(userID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"data": snippetIDs,
	})
}

func (h *Handler) addFavouriteSnippet(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	snippetID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid snippet id parameter")
		return
	}
	if err := h.services.Snippet.AddFavourite(userID, snippetID); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) removeFavouriteSnippet(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	snippetID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid snippet id parameter")
		return
	}
	if err := h.services.Snippet.RemoveFavourite(userID, snippetID); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{"ok"})
}
