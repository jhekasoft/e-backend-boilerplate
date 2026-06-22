package crud

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type (
	GetResponse[M CRUDModel] struct {
		Data M
	}

	CreateResponse[M CRUDModel] GetResponse[M]

	UpdateResponse[M CRUDModel] GetResponse[M]

	ListResponse[M CRUDModel] struct {
		Data   []M
		Offset int
		Limit  int
		Total  int64
	}
)

type Handler[
	M CRUDModel,
	F CRUDListFilter,
	CR CRUDCreateRequest[M],
	UR CRUDUpdateRequest[M],
	LFR CRUDListFilterRequest[F],
] struct {
	service CRUDService[M, F]
}

func NewHandler[
	M CRUDModel,
	F CRUDListFilter,
	CR CRUDCreateRequest[M],
	UR CRUDUpdateRequest[M],
	LFR CRUDListFilterRequest[F],
](service CRUDService[M, F]) *Handler[M, F, CR, UR, LFR] {
	return &Handler[M, F, CR, UR, LFR]{service}
}

func (h *Handler[M, F, CR, UR, LFR]) GetService() CRUDService[M, F] {
	return h.service
}

func (h *Handler[M, F, CR, UR, LFR]) CreateItem(c echo.Context) error {
	var req CR
	err := c.Bind(&req)
	if err != nil {
		return err
	}
	if err = c.Validate(req); err != nil {
		return err
	}

	item := req.ToModel()
	createdItem, err := h.service.Create(item)
	if err != nil {
		return err
	}

	resp := CreateResponse[M]{Data: *createdItem}
	return c.JSON(http.StatusOK, resp)
}

func (h *Handler[M, F, CR, UR, LFR]) UpdateItem(c echo.Context) error {
	// Get ID parameter
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return err
	}

	var req UR
	err = c.Bind(&req)
	if err != nil {
		return err
	}
	if err = c.Validate(req); err != nil {
		return err
	}

	item := req.ToModel()

	updatedItem, err := h.service.Update(uint(id), item)
	if err != nil {
		return err
	}

	resp := UpdateResponse[M]{Data: *updatedItem}

	return c.JSON(http.StatusOK, resp)
}

func (h *Handler[M, F, CR, UR, LFR]) GetItem(c echo.Context) error {
	// Get ID parameter
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return err
	}

	item, err := h.service.Get(uint(id))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}
	if err != nil {
		return err
	}

	resp := GetResponse[M]{Data: *item}

	return c.JSON(http.StatusOK, resp)
}

func (h *Handler[M, F, CR, UR, LFR]) GetList(c echo.Context) error {
	var req LFR
	err := c.Bind(&req)
	if err != nil {
		return err
	}

	filter := req.ToFilter()

	list, total, err := h.GetService().GetManyWithTotal(filter)
	if err != nil {
		return err
	}

	resp := ListResponse[M]{
		Data:   list,
		Offset: filter.GetOffset(),
		Limit:  filter.GetLimit(),
		Total:  total,
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *Handler[M, F, CR, UR, LFR]) DeleteItem(c echo.Context) error {
	// Get ID parameter
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return err
	}

	err = h.service.Delete(uint(id))
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
