package controllers

import (
	"apiInvitation/src/users/application"
	_"apiInvitation/src/users/domain/entities"
	"apiInvitation/src/users/infraestructure/utils"
	"bytes"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UploadPictureUserController struct {
    uploadUserUseCase *application.UploadPictureUserUseCase
}

func NewUploadPictureUserController(uploadPictureUseCase *application.UploadPictureUserUseCase) *UploadPictureUserController {
    return &UploadPictureUserController{uploadUserUseCase: uploadPictureUseCase}
}

func (uuc *UploadPictureUserController) UpdatePictureUser(g *gin.Context) {
    idParam := g.Param("id")
    id, err := strconv.ParseInt(idParam, 10, 32)
    if err != nil {
        g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    // Manejar la imagen del perfil
    file, header, err := g.Request.FormFile("profile_picture")
    if err != nil {
        g.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get profile picture"})
        return
    }
    defer file.Close()

    // Leer el archivo en bytes
    buf := new(bytes.Buffer)
    _, err = io.Copy(buf, file)
    if err != nil {
        g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read profile picture"})
        return
    }

    // Subir la imagen a S3
    url, err := utils.UploadToS3(buf.Bytes(), header.Filename)
    if err != nil {
        g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload profile picture"})
        return
    }

    // Actualizar la URL de la imagen en la base de datos
    idCheck := int32(id)
    err = uuc.uploadUserUseCase.Execute(idCheck, url)
    if err != nil {
        g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile picture URL"})
        return
    }

    response := gin.H{
        "data": gin.H{
            "type": "users",
            "id":   idCheck,
            "attributes": gin.H{
                "profile_picture": url,
            },
        },
    }
    
    g.JSON(http.StatusOK, response)
}
