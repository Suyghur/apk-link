/*
@Time : 2020/9/3
@Author : #Suyghur,
@File : common
*/

package response

type PageResult struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
}
