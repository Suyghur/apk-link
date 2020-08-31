/*
@Time : 2020/8/31
@Author : #Suyghur,
@File : common
*/

package request

type PageInfo struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

type GetById struct {
	Id float64 `json:"id"`
}

type IdsReq struct {
	Ids []int `json:"ids"`
}
