package http

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/DarkSoul94/metamemorys_backend/global_const"
	"github.com/DarkSoul94/metamemorys_backend/metamemory"
	"github.com/DarkSoul94/metamemorys_backend/models"
	"github.com/gin-gonic/gin"
)

// Handler ...
type MetaHandler struct {
	metaUC metamemory.MetaUsecase
}

// NewHandler ...
func NewMetaHandler(uc metamemory.MetaUsecase) *MetaHandler {
	return &MetaHandler{
		metaUC: uc,
	}
}

func (h *MetaHandler) GetMemberList(ctx *gin.Context) {
	var outMembers []outMember = make([]outMember, 0)

	user, _ := ctx.Get(global_const.CtxUserKey)

	members, err := h.metaUC.GetMemberList(user.(*models.User).ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}

	for _, member := range members {
		outMembers = append(outMembers, h.toOutMember(member))
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"status": "success", "members": outMembers})
}

func (h *MetaHandler) CreateMember(ctx *gin.Context) {
	var member newMember

	if err := ctx.BindJSON(&member); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}

	user, _ := ctx.Get(global_const.CtxUserKey)

	id, err := h.metaUC.CreateMember(h.toModelsMember(member, user.(*models.User)))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"status": "success", "id": id})
}

func (h *MetaHandler) CreateFile(ctx *gin.Context) {
	var newFiles inpFiles

	if err := ctx.BindJSON(&newFiles); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}

	user, _ := ctx.Get(global_const.CtxUserKey)

	h.metaUC.CreateFiles(user.(*models.User).ID, newFiles.MemberID, h.toModelFiles(newFiles))

	ctx.JSON(http.StatusOK, map[string]interface{}{"status": "success"})
}

var invalidID error = errors.New("Invalid member id")

func (h *MetaHandler) GetMemberFiles(ctx *gin.Context) {
	memberID, ok := ctx.GetQuery("memberID")
	if !ok {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"status": "error", "error": invalidID})
		return
	}

	id, err := strconv.ParseUint(memberID, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"status": "error", "error": invalidID})
		return
	}

	files, err := h.metaUC.GetMemberFiles(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"status": "error", "error": err})
		return
	}

	var outFiles []outFile
	for _, file := range files {
		outFiles = append(outFiles, h.toOutFile(file))
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{"status": "success", "files": outFiles})
}
