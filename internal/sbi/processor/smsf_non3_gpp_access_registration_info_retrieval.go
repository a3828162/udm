package processor

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetSmsfNon3gppAccess - retrieve the SMSF registration for non-3GPP access information
func (p *Processor) HandleGetSmsfNon3gppAccess(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}