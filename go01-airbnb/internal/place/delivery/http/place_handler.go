package placehttp

import (
	"context"
	placemodel "go01-airbnb/internal/place/model"
	"go01-airbnb/pkg/common"
	"go01-airbnb/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PlaceUseCase interface {
	CreatePlace(context.Context, *placemodel.Place) error
	GetPlaces(context.Context, *common.Paging, *placemodel.Filter) ([]placemodel.Place, error)
	GetPlaceByID(context.Context, int) (*placemodel.Place, error)
	UpdatePlace(context.Context, common.Requester, int, *placemodel.Place) error
	DeletePlace(context.Context, common.Requester, int) error
}

type placeHandler struct {
	placeUC PlaceUseCase
	hasher  *utils.Hasher
}

func NewPlaceHandler(placeUC PlaceUseCase, hasher *utils.Hasher) *placeHandler {
	return &placeHandler{placeUC, hasher}
}

func (hdl *placeHandler) CreatePlace() gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet("user").(common.Requester)

		var place placemodel.Place

		if err := c.ShouldBind(&place); err != nil {
			panic(common.ErrBadRequest(err))
		}

		place.OwnerId = requester.GetUserId()

		if err := hdl.placeUC.CreatePlace(c.Request.Context(), &place); err != nil {
			panic(err)
		}

		// Encode id trước trả ra cho client
		place.FakeId = hdl.hasher.Encode(place.Id, common.DBTypePlace)

		c.JSON(http.StatusOK, common.Response(place))
	}
}

func (hdl *placeHandler) GetPlaces() gin.HandlerFunc {
	return func(c *gin.Context) {
		// paging
		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrBadRequest(err))
		}
		paging.Fulfill()

		// filter
		var filter placemodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrBadRequest(err))
		}

		data, err := hdl.placeUC.GetPlaces(c.Request.Context(), &paging, &filter)
		if err != nil {
			panic(err)
		}

		// Encode id trước trả ra cho client
		for i, v := range data {
			data[i].FakeId = hdl.hasher.Encode(v.Id, common.DBTypePlace)
			data[i].Owner.FakeId = hdl.hasher.Encode(v.Owner.Id, common.DBTypeUser)
		}

		c.JSON(http.StatusOK, common.ResponseWithPaging(data, paging))
	}
}

func (hdl *placeHandler) GetPlaceByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// id, err := strconv.Atoi(c.Param("id"))
		id, err := hdl.hasher.Decode(c.Param("id"))
		if err != nil {
			panic(common.ErrBadRequest(err))
		}

		data, err := hdl.placeUC.GetPlaceByID(c.Request.Context(), id)
		if err != nil {
			panic(err)
		}

		data.FakeId = hdl.hasher.Encode(data.Id, common.DBTypePlace)
		data.Owner.FakeId = hdl.hasher.Encode(data.Owner.Id, common.DBTypeUser)
		c.JSON(http.StatusOK, common.Response(data))
	}
}

func (hdl *placeHandler) UpdatePlace() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Lấy thông tin Requester
		requester := c.MustGet("user").(common.Requester)

		// id, err := strconv.Atoi(c.Param("id"))
		id, err := hdl.hasher.Decode(c.Param("id"))
		if err != nil {
			panic(common.ErrBadRequest(err))
		}

		var place placemodel.Place

		if err := c.ShouldBind(&place); err != nil {
			panic(common.ErrBadRequest(err))
		}

		if err := hdl.placeUC.UpdatePlace(c.Request.Context(), requester, id, &place); err != nil {
			panic(err)
		}

		place.FakeId = hdl.hasher.Encode(place.Id, common.DBTypePlace)
		c.JSON(http.StatusOK, common.Response(place))
	}
}

func (hdl *placeHandler) DeletePlace() gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet("user").(common.Requester)

		// id, err := strconv.Atoi(c.Param("id"))
		id, err := hdl.hasher.Decode(c.Param("id"))
		if err != nil {
			panic(common.ErrBadRequest(err))
		}

		if err := hdl.placeUC.DeletePlace(c.Request.Context(), requester, id); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.Response(true))
	}
}
