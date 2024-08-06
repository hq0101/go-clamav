package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hq0101/go-clamav/pkg/clamav"
	"net/http"
)

type ClamdController struct {
}

func NewClamd() *ClamdController {
	return &ClamdController{}
}

// Ping godoc
// @Summary Ping Check the server's state. It should reply with "PONG".
// @Tags ClamAV
// @Success 200 {string} string PONG
// @Failure 500 {string} string
// @Router /ping [get]
func (c *ClamdController) Ping(ctx *gin.Context) {
	clamClient := clamav.NewClamClient(GetCfg().ClamdNetworkType, GetCfg().ClamdAddress, GetCfg().ClamdConnTimeout, GetCfg().ClamdReadTimeout)
	content, err := clamClient.Ping()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, content)
}

// Version godoc
// @Summary Get the ClamAV version
// @Tags ClamAV
// @Success 200 {string} string "ClamAV 0.103.11/27353/Wed"
// @Failure 500 {string} string
// @Router /version [get]
func (c *ClamdController) Version(ctx *gin.Context) {
	clamClient := clamav.NewClamClient(GetCfg().ClamdNetworkType, GetCfg().ClamdAddress, GetCfg().ClamdConnTimeout, GetCfg().ClamdReadTimeout)
	content, err := clamClient.Version()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, content)
}

// VersionCommands godoc
// @Summary Get the ClamAV version
// @Tags ClamAV
// @Success 200 {string} string "ClamAV 0.103.11/27353/Wed"
// @Failure 500 {string} string
// @Router /versioncommands [get]
func (c *ClamdController) VersionCommands(ctx *gin.Context) {
	clamClient := clamav.NewClamClient(GetCfg().ClamdNetworkType, GetCfg().ClamdAddress, GetCfg().ClamdConnTimeout, GetCfg().ClamdReadTimeout)
	content, err := clamClient.VersionCommands()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, content)
}

// Stats godoc
// @Summary Get ClamAV stats
// @Tags ClamAV
// @Success 200 {object} cli.ClamdStats
// @Failure 500 {string} string
// @Router /stats [get]
func (c *ClamdController) Stats(ctx *gin.Context) {
	clamClient := clamav.NewClamClient(GetCfg().ClamdNetworkType, GetCfg().ClamdAddress, GetCfg().ClamdConnTimeout, GetCfg().ClamdReadTimeout)
	content, err := clamClient.Stats()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, content)
}

// Reload godoc
// @Summary Reload the virus database
// @Tags ClamAV
// @Success 200 {string} string RELOADING
// @Failure 500 {string} string
// @Router /reload [post]
func (c *ClamdController) Reload(ctx *gin.Context) {
	clamClient := clamav.NewClamClient(GetCfg().ClamdNetworkType, GetCfg().ClamdAddress, GetCfg().ClamdConnTimeout, GetCfg().ClamdReadTimeout)
	content, err := clamClient.Reload()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, content)
}

// Shutdown godoc
// @Summary Shut down the ClamAV server
// @Tags ClamAV
// @Success 200 {string} string "OK"
// @Failure 500 {string} string
// @Router /shutdown [post]
func (c *ClamdController) Shutdown(ctx *gin.Context) {
	clamClient := clamav.NewClamClient(GetCfg().ClamdNetworkType, GetCfg().ClamdAddress, GetCfg().ClamdConnTimeout, GetCfg().ClamdReadTimeout)
	content, err := clamClient.Shutdown()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, content)
}

// Scan godoc
// @Summary Scan a file or directory
// @Tags ClamAV
// @Param file query string true "File path to scan"
// @Success 200 {array} cli.ScanResult
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /scan [get]
func (c *ClamdController) Scan(ctx *gin.Context) {
	filePath := ctx.Query("file")
	if filePath == "" {
		ctx.JSON(http.StatusBadRequest, "File path is required")
		return
	}
	clamClient := clamav.NewClamClient(GetCfg().ClamdNetworkType, GetCfg().ClamdAddress, GetCfg().ClamdConnTimeout, GetCfg().ClamdReadTimeout)
	content, err := clamClient.Scan(filePath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, content)
}

// Contscan godoc
// @Summary Continuously scan a file or directory
// @Tags ClamAV
// @Param file query string true "File path to scan"
// @Success 200 {array} cli.ScanResult
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /contscan [get]
func (c *ClamdController) Contscan(ctx *gin.Context) {
	filePath := ctx.Query("file")
	if filePath == "" {
		ctx.JSON(http.StatusBadRequest, "File path is required")
		return
	}
	clamClient := clamav.NewClamClient(GetCfg().ClamdNetworkType, GetCfg().ClamdAddress, GetCfg().ClamdConnTimeout, GetCfg().ClamdReadTimeout)
	content, err := clamClient.ContScan(filePath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, content)
}

// MultiScan godoc
// @Summary Multithreaded scan of a file or directory
// @Tags ClamAV
// @Param file query string true "File path to scan"
// @Success 200 {array} cli.ScanResult
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /multiscan [get]
func (c *ClamdController) MultiScan(ctx *gin.Context) {
	filePath := ctx.Query("file")
	if filePath == "" {
		ctx.JSON(http.StatusBadRequest, "File path is required")
		return
	}
	clamClient := clamav.NewClamClient(GetCfg().ClamdNetworkType, GetCfg().ClamdAddress, GetCfg().ClamdConnTimeout, GetCfg().ClamdReadTimeout)
	content, err := clamClient.MultiScan(filePath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, content)
}

// AllMatchScan godoc
// @Summary AllMatchScan scan a file or directory
// @Tags ClamAV
// @Param file query string true "File path to scan"
// @Success 200 {array} cli.ScanResult
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /allmatchscan [get]
func (c *ClamdController) AllMatchScan(ctx *gin.Context) {
	filePath := ctx.Query("file")
	if filePath == "" {
		ctx.JSON(http.StatusBadRequest, "File path is required")
		return
	}
	clamClient := clamav.NewClamClient(GetCfg().ClamdNetworkType, GetCfg().ClamdAddress, GetCfg().ClamdConnTimeout, GetCfg().ClamdReadTimeout)
	content, err := clamClient.AllMatchScan(filePath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, content)
}

// Instream godoc
// @Summary Scan data stream
// @Tags ClamAV
// @Accept  multipart/form-data
// @Produce application/json
// @Param file formData file true "File to upload"
// @Success 200 {array} cli.ScanResult
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /instream [post]
func (c *ClamdController) Instream(ctx *gin.Context) {
	data, err := ctx.GetRawData()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Failed to read raw data")
		return
	}
	clamClient := clamav.NewClamClient(GetCfg().ClamdNetworkType, GetCfg().ClamdAddress, GetCfg().ClamdConnTimeout, GetCfg().ClamdReadTimeout)
	content, err := clamClient.Instream(data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, content)
}
