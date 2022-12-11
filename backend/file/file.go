package file

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	Name    string `json:"name" binding:"required"`
	Content string `json:"content" binding:"required"`
	Format  string `json:"format"`
	PetID   int
}

type FileHandler struct {
	db *gorm.DB
}

func NewFileHandler(db *gorm.DB) *FileHandler {
	return &FileHandler{db: db}
}

func (h *FileHandler) Upload(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var fileList []File
	files := form.File["files[]"]

	for _, f := range files {
		openFile, err := f.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		file, _ := ioutil.ReadAll(openFile)
		name := f.Filename
		format := filepath.Ext(f.Filename)
		sEncoded := base64.StdEncoding.EncodeToString(file)

		f := File{
			Name:    name,
			Content: sEncoded,
			Format:  format,
			Model:   gorm.Model{},
			PetID:   id,
		}
		fileList = append(fileList, f)
	}

	result := h.db.Create(&fileList)
	if err := result.Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"ID": fileList[0].PetID,
	})
}

func (h *FileHandler) Download(c *gin.Context) {
	queryId := c.Request.URL.Query().Get("id")
	queryName := c.Request.URL.Query().Get("name")

	if queryId == "" && queryName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "query parameter missing",
		})
		return
	}

	var files []File
	if queryId != "" {
		id, err := strconv.Atoi(queryId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		result := h.db.Find(&files, id)
		if err := result.Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
	} else {
		name := queryName
		result := h.db.Where("name = ?", name).Find(&files)
		if err := result.Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	bytes, err := base64.StdEncoding.DecodeString(files[0].Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.Header("Content-Disposition", "attachment; filename="+files[0].Name)
	c.Data(http.StatusOK, "application/octet-stream", bytes)
}
