package server

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
)

// Handler handler
type Handler struct {
	db       *DB
	document DocumentRepository
	asset    AssetRepository
}

// NewHandler ハンドラを生成する
func NewHandler(db *DB, document DocumentRepository, asset AssetRepository) Handler {
	return Handler{
		db:       db,
		document: document,
		asset:    asset,
	}
}

// GetLogin ログインページ
func (h Handler) GetLogin(c Context) error {
	return c.Render(http.StatusOK, TmplLogin, LoginView{
		Errors: ValidationResult{},
	})
}

// LoginParam ログイン
type LoginParam struct {
	Login    string `form:"login" validate:"required"`
	Password string `form:"password" validate:"required"`
}

// PostLogin ログイン処理
func (h Handler) PostLogin(c Context) error {
	var prm LoginParam
	if err := c.Bind(&prm); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	errors := ValidationResult{}
	if err := c.Validate(prm); err != nil {
		errors = ReportValidation(err)
	}

	user, ok, err := h.db.VerifyUser(prm.Login, prm.Password)
	if err != nil {
		return err
	}

	// success login
	if ok {
		sess, _ := getSession(c)
		sess.Values[SessionKeyUserID] = user.ID
		if err := saveSession(c, sess); err != nil {
			return err
		}
		return c.Redirect(http.StatusSeeOther, "/")
	}

	return c.Render(http.StatusOK, TmplLogin, LoginView{
		Errors: errors,
	})
}

// GetLogout is a handler to logout users
func (h Handler) GetLogout(c Context) error {
	// TODO: logout process
	return c.Render(http.StatusOK, TmplLogout, nil)
}

// GetIndex is a handler show index of webpage
func (h Handler) GetIndex(c Context) error {
	return c.String(http.StatusOK, "it's works!")
}

// GetDocument is a handler of get document
func (h Handler) GetDocument(c Context) error {
	title := c.Param("title")

	doc, err := h.document.Get(title)
	if err == DocumentNotFound {
		if !c.IsLoggedIn {
			return c.String(http.StatusNotFound, "document not found")
		}
		return c.Render(http.StatusOK, TmplEdit, EditView{
			Title: title,
		})
	}
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, doc.Content)
}

// PutDocumentParam is a parameter used as PutDocument handler
type PutDocumentParam struct {
	Content string `json:"content" validate:"required"`
}

// PutDocument is a handler to save document
func (h Handler) PutDocument(c Context) error {
	title := c.Param("title")

	prm := PutDocumentParam{}
	if err := c.Bind(&prm); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	if err := c.Validate(prm); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	doc := Document{
		Title:   title,
		Content: prm.Content,
	}

	if err := h.document.Put(doc); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, struct{}{})
}

// GetAsset is a handler
func (h Handler) GetAsset(c Context) error {
	id := c.Param("id")

	asset, err := h.asset.Get(id)
	if err == AssetNotFound {
		return c.String(http.StatusNotFound, "asset not found")
	}
	if err != nil {
		return err
	}
	defer asset.Close()

	h.decolateAssetResponse(c, asset)
	return c.Stream(http.StatusOK, asset.ContentType, asset)
}

// GetFormatedAsset is a handler
func (h Handler) GetFormatedAsset(c Context) error {
	id := c.Param("id")
	// format := c.Param("format")

	asset, err := h.asset.Get(id)
	if err == AssetNotFound {
		return c.String(http.StatusNotFound, "asset not found")
	}
	if err != nil {
		return err
	}
	defer asset.Close()

	h.decolateAssetResponse(c, asset)
	return c.Stream(http.StatusOK, asset.ContentType, asset)
}

func (h Handler) decolateAssetResponse(c Context, asset Asset) error {
	resp := c.Response()
	header := resp.Header()
	// for download
	download := c.QueryParam("download")
	if download == "true" || download == "yes" {
		encodedFileName := url.QueryEscape(asset.FileName)
		hValue := fmt.Sprintf(`attachment;filename*="UTF-8''%s"`, encodedFileName)
		header.Set(echo.HeaderContentDisposition, hValue)
	}
	return nil
}

func (h Handler) UploadAsset(c Context) error {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return err
	}

	fileName := fileHeader.Filename
	contentType := fileHeader.Header.Get("Content-Type")
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	sa := Asset{
		Stream:      file,
		FileName:    fileName,
		ContentType: contentType,
	}
	id, err := h.asset.Add(sa)
	if err != nil {
		return err
	}

	defer file.Close()

	return c.String(http.StatusOK, id)
}
