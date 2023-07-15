package shopee

import (
	"fmt"
	"github.com/Epur/ext-sdk/model"
	"strconv"
)

//倒序

func (p *Api) GetOrderListDesc(Body model.BodyMap) *model.Client {

	var (
		endx   int64
		result = GetOrderListResponse{}
		c      *model.Client
	)

	Body.Set("sort_direction", "DESC")
	end, _ := strconv.ParseInt(Body.Get("time_to"), 10, 64)
	start, _ := strconv.ParseInt(Body.Get("time_from"), 10, 64)

	if end-HALFMONTYDAY <= start {
		endx = start
	} else {
		endx = end - HALFMONTYDAY
	}

	for endx >= start && start < end {

		Body.Set("time_from", fmt.Sprintf("%d", endx)).
			Set("time_to", fmt.Sprintf("%d", end))

		c = p.GetOrderList(Body)
		if c.Err != nil {
			return c
		}

		result.List = append(result.List, c.Response.Response.DataTo.(GetOrderListResponse).List...)

		if endx <= start {
			break
		}
		end, endx = endx, endx-HALFMONTYDAY
		if endx <= start {
			endx = start
		}
	}
	result.Total = len(result.List)
	c.Response.Response.DataTo = result

	return c
}

//正序

func (p *Api) GetOrderListAsc(Body model.BodyMap) *model.Client {

	var (
		endx   int64
		result = GetOrderListResponse{}
		c      *model.Client
	)

	Body.Set("sort_direction", "DESC")
	end, _ := strconv.ParseInt(Body.Get("time_to"), 10, 64)
	start, _ := strconv.ParseInt(Body.Get("time_from"), 10, 64)

	if start+HALFMONTYDAY > end {
		endx = end
	} else {
		endx = start + HALFMONTYDAY
	}

	for endx <= end && start < end {

		Body.Set("time_from", fmt.Sprintf("%d", start)).
			Set("time_to", fmt.Sprintf("%d", endx))

		c = p.GetOrderList(Body)
		if c.Err != nil {
			return c
		}

		result.List = append(result.List, c.Response.Response.DataTo.(GetOrderListResponse).List...)

		if endx >= end {
			break
		}
		start, endx = endx, endx+HALFMONTYDAY
		if endx >= end {
			endx = end
		}
	}
	result.Total = len(result.List)
	c.Response.Response.DataTo = result

	return c
}
