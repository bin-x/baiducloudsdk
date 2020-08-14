package baiducloudsdk

import (
	"encoding/base64"
	"io/ioutil"
	 "net/url"
)

/**
 * 相同图检索—入库 same_hq_add api url
 */
const sameHqAddUrl = "https://aip.baidubce.com/rest/2.0/realtime_search/same_hq/add"

/**
 * 相同图检索—检索 same_hq_search api url
 */
const sameHqSearchUrl = "https://aip.baidubce.com/rest/2.0/realtime_search/same_hq/search"

/**
 * 相同图检索—更新 same_hq_update api url
 */
const sameHqUpdateUrl = "https://aip.baidubce.com/rest/2.0/realtime_search/same_hq/update"

/**
 * 相同图检索—删除 same_hq_delete api url
 */
const sameHqDeleteUrl = "https://aip.baidubce.com/rest/2.0/realtime_search/same_hq/delete"

/**
 * 相似图检索—入库 similar_add api url
 */
const similarAddUrl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/realtime_search/similar/add"

/**
 * 相似图检索—检索 similar_search api url
 */
const similarSearchUrl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/realtime_search/similar/search"

/**
 * 相似图检索—更新 similar_update api url
 */
const similarUpdateUrl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/realtime_search/similar/update"

/**
 * 相似图检索—删除 similar_delete api url
 */
const similarDeleteUrl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/realtime_search/similar/delete"

/**
 * 商品检索—入库 product_add api url
 */
const productAddUrl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/realtime_search/product/add"

/**
 * 商品检索—检索 product_search api url
 */
const productSearchUrl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/realtime_search/product/search"

/**
 * 商品检索—更新 product_update api url
 */
const productUpdateUrl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/realtime_search/product/update"

/**
 * 商品检索—删除 product_delete api url
 */
const productDeleteUrl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/realtime_search/product/delete"


/**
 * 绘本图片搜索—入库-image
 */
const picturebookAddUrl = "https://aip.baidubce.com/rest/2.0/imagesearch/v1/realtime_search/picturebook/add"

/**
 * 绘本图片搜索—入库-检索
 */
const picturebookSearchUrl = "https://aip.baidubce.com/rest/2.0/imagesearch/v1/realtime_search/picturebook/search"

/**
 * 绘本图片搜索—入库-删除
 */
const picturebookDeleteUrl = "https://aip.baidubce.com/rest/2.0/imagesearch/v1/realtime_search/picturebook/delete";

/**
 * 绘本图片搜索—入库-更新
 */
const picturebookUpdateUrl = "https://aip.baidubce.com/rest/2.0/imagesearch/v1/realtime_search/picturebook/update"


func  (client *HttpClient) RequestByImage(reqUrl string, imagePath string, options map[string]string)([]byte, error) {
	data:= url.Values{}
	image,err:=ioutil.ReadFile(imagePath)
	if err!=nil {
		return nil, err
	}
	imageBase64:=base64.StdEncoding.EncodeToString(image)
	data["image"] = []string{imageBase64}
	for k,v:=range options {
		data[k]= []string{v}
	}
	return client.Request(reqUrl,data,nil)
}

func  (client *HttpClient) RequestByUrl(reqUrl string, imageUrl string, options map[string]string)([]byte, error) {
	data:= url.Values{}
	data["url"] = []string{imageUrl}
	for k,v:=range options {
		data[k]= []string{v}
	}
	return client.Request(reqUrl,data,nil)
}

func  (client *HttpClient) RequestBySign(reqUrl string, contSign string, options map[string]string)([]byte, error) {
	data:= url.Values{}
	data["cont_sign"] = []string{contSign}
	for k,v:=range options {
		data[k]= []string{v}
	}
	return client.Request(reqUrl,data,nil)
}

func (client *HttpClient) SameHqAdd(imagePath string,brief string, options map[string]string) ([]byte, error) {
	options["brief"] = brief
	return client.RequestByImage(sameHqAddUrl,imagePath,options)
}

func (client *HttpClient) SameHqAddUrl(imageUrl string, brief string, options map[string]string) ([]byte, error) {
	options["brief"] = brief
	return client.RequestByUrl(sameHqAddUrl,imageUrl,options)
}

func (client *HttpClient) SameHqSearch(image string, options map[string]string) ([]byte, error) {
	return client.RequestByImage(sameHqSearchUrl,image,options)
}

func (client *HttpClient) SameHqSearchUrl(imageUrl string, options map[string]string) ([]byte, error) {
	return client.RequestByUrl(sameHqSearchUrl,imageUrl,options)
}

func (client *HttpClient) SameHqUpdate(image string, options map[string]string) ([]byte, error) {
	return client.RequestByImage(sameHqUpdateUrl,image,options)
}

func (client *HttpClient) SameHqUpdateUrl(imageUrl string, options map[string]string) ([]byte, error) {
	return client.RequestByUrl(sameHqUpdateUrl,imageUrl,options)
}

func (client *HttpClient) SameHqUpdateContSign(contSign string, options map[string]string) ([]byte, error) {
	return client.RequestBySign(sameHqUpdateUrl,contSign,options)
}

func (client *HttpClient) SameHqDeleteByImage(image string, options map[string]string) ([]byte, error) {
	return client.RequestByImage(sameHqDeleteUrl,image,options)
}

func (client *HttpClient) SameHqDeleteByUrl(imageUrl string, options map[string]string) ([]byte, error) {
	return client.RequestByUrl(sameHqDeleteUrl,imageUrl,options)
}

func (client *HttpClient) SameHqDeleteBySign(contSign string, options map[string]string) ([]byte, error) {
	return client.RequestBySign(sameHqDeleteUrl,contSign,options)
}

// 相似图片
func (client *HttpClient) SimilarAdd(imagePath string,brief string, options map[string]string) ([]byte, error) {
	options["brief"] = brief
	return client.RequestByImage(similarAddUrl,imagePath,options)
}

func (client *HttpClient) SimilarAddUrl(imageUrl string, brief string, options map[string]string) ([]byte, error) {
	options["brief"] = brief
	return client.RequestByUrl(similarAddUrl,imageUrl,options)
}

func (client *HttpClient) SimilarSearch(image string, options map[string]string) ([]byte, error) {
	return client.RequestByImage(similarSearchUrl,image,options)
}

func (client *HttpClient) SimilarSearchUrl(imageUrl string, options map[string]string) ([]byte, error) {
	return client.RequestByUrl(similarSearchUrl,imageUrl,options)
}

func (client *HttpClient) SimilarUpdate(image string, options map[string]string) ([]byte, error) {
	return client.RequestByImage(similarUpdateUrl,image,options)
}

func (client *HttpClient) SimilarUpdateUrl(imageUrl string, options map[string]string) ([]byte, error) {
	return client.RequestByUrl(similarUpdateUrl,imageUrl,options)
}

func (client *HttpClient) SimilarUpdateContSign(contSign string, options map[string]string) ([]byte, error) {
	return client.RequestBySign(similarUpdateUrl,contSign,options)
}

func (client *HttpClient) SimilarDeleteByImage(image string, options map[string]string) ([]byte, error) {
	return client.RequestByImage(similarDeleteUrl,image,options)
}

func (client *HttpClient) SimilarDeleteByUrl(imageUrl string, options map[string]string) ([]byte, error) {
	return client.RequestByUrl(similarDeleteUrl,imageUrl,options)
}

func (client *HttpClient) SimilarDeleteBySign(contSign string, options map[string]string) ([]byte, error) {
	return client.RequestBySign(similarDeleteUrl,contSign,options)
}


// 商品检索
func (client *HttpClient) ProductAdd(imagePath string,brief string, options map[string]string) ([]byte, error) {
	options["brief"] = brief
	return client.RequestByImage(productAddUrl,imagePath,options)
}

func (client *HttpClient) ProductAddUrl(imageUrl string, brief string, options map[string]string) ([]byte, error) {
	options["brief"] = brief
	return client.RequestByUrl(productAddUrl,imageUrl,options)
}

func (client *HttpClient) ProductSearch(image string, options map[string]string) ([]byte, error) {
	return client.RequestByImage(productSearchUrl,image,options)
}

func (client *HttpClient) ProductSearchUrl(imageUrl string, options map[string]string) ([]byte, error) {
	return client.RequestByUrl(productSearchUrl,imageUrl,options)
}

func (client *HttpClient) ProductUpdate(image string, options map[string]string) ([]byte, error) {
	return client.RequestByImage(productUpdateUrl,image,options)
}

func (client *HttpClient) ProductUpdateUrl(imageUrl string, options map[string]string) ([]byte, error) {
	return client.RequestByUrl(productUpdateUrl,imageUrl,options)
}

func (client *HttpClient) ProductUpdateContSign(contSign string, options map[string]string) ([]byte, error) {
	return client.RequestBySign(productUpdateUrl,contSign,options)
}

func (client *HttpClient) ProductDeleteByImage(image string, options map[string]string) ([]byte, error) {
	return client.RequestByImage(productDeleteUrl,image,options)
}

func (client *HttpClient) ProductDeleteByUrl(imageUrl string, options map[string]string) ([]byte, error) {
	return client.RequestByUrl(productDeleteUrl,imageUrl,options)
}

func (client *HttpClient) ProductDeleteBySign(contSign string, options map[string]string) ([]byte, error) {
	return client.RequestBySign(productDeleteUrl,contSign,options)
}

// 绘本图片搜索
func (client *HttpClient) PictureBookAdd(imagePath string,brief string, options map[string]string) ([]byte, error) {
	options["brief"] = brief
	return client.RequestByImage(picturebookAddUrl,imagePath,options)
}

func (client *HttpClient) PictureBookAddUrl(imageUrl string, brief string, options map[string]string) ([]byte, error) {
	options["brief"] = brief
	return client.RequestByUrl(picturebookAddUrl,imageUrl,options)
}

func (client *HttpClient) PictureBookSearch(image string, options map[string]string) ([]byte, error) {
	return client.RequestByImage(picturebookSearchUrl,image,options)
}

func (client *HttpClient) PictureBookSearchUrl(imageUrl string, options map[string]string) ([]byte, error) {
	return client.RequestByUrl(picturebookSearchUrl,imageUrl,options)
}

func (client *HttpClient) PictureBookUpdate(image string, options map[string]string) ([]byte, error) {
	return client.RequestByImage(picturebookUpdateUrl,image,options)
}

func (client *HttpClient) PictureBookUpdateUrl(imageUrl string, options map[string]string) ([]byte, error) {
	return client.RequestByUrl(picturebookUpdateUrl,imageUrl,options)
}

func (client *HttpClient) PictureBookUpdateContSign(contSign string, options map[string]string) ([]byte, error) {
	return client.RequestBySign(picturebookUpdateUrl,contSign,options)
}

func (client *HttpClient) PictureBookDeleteByImage(image string, options map[string]string) ([]byte, error) {
	return client.RequestByImage(picturebookDeleteUrl,image,options)
}

func (client *HttpClient) PictureBookDeleteByUrl(imageUrl string, options map[string]string) ([]byte, error) {
	return client.RequestByUrl(picturebookDeleteUrl,imageUrl,options)
}

func (client *HttpClient) PictureBookDeleteBySign(contSign string, options map[string]string) ([]byte, error) {
	return client.RequestBySign(picturebookDeleteUrl,contSign,options)
}