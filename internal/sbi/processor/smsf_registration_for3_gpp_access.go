package processor

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UpdateSMSFReg3GPP - register as SMSF for 3GPP access
func (p *Processor) HandleUpdateSMSFReg3GPP(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}