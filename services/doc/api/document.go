/*
 * Copyright 2022 Baidu, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
 * except in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the
 * License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions
 * and limitations under the License.
 */

// document.go - the document APIs definition supported by the DOC service

// Package api defines all APIs supported by the DOC service of BCE.
package api

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// RegisterDocument - register document in doc service
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - regParam title and format of the document being registered
// RETURNS:
//     - *RegDocumentResp: id and document location in bos
//     - error: the return error if any occurs
func RegisterDocument(cli bce.Client, regParam *RegDocumentParam) (*RegDocumentResp, error) {
	if regParam == nil {
		return nil, errors.New("param cannot be nil")
	}
	playload, err := regParam.String()
	if err != nil {
		return nil, err
	}
	body, err := bce.NewBodyFromString(playload)
	if err != nil {
		return nil, err
	}

	req := &bce.BceRequest{}
	req.SetUri("/v2/document")
	req.SetParam("register", "")
	req.SetMethod(http.POST)
	req.SetHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE)
	req.SetBody(body)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	result := &RegDocumentResp{}
	if err := resp.ParseJsonBody(result); err != nil {
		return nil, err
	}
	return result, nil
}

// PublishDocument - publish document
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - documentId: id of document in doc service
// RETURNS:
//     - error: the return error if any occurs
func PublishDocument(cli bce.Client, documentId string) error {
	req := &bce.BceRequest{}
	urlPath := fmt.Sprintf("/v2/document/%s", documentId)
	req.SetUri(urlPath)
	req.SetParam("publish", "")
	req.SetMethod(http.PUT)
	req.SetHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE)
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}
	return nil

}

// QueryDocument - query document's status
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - documentId: id of document in doc service
//     - queryParam: enable/disable https of coverl url
// RETURNS:
//     - *QueryDocumentResp
//     - error: the return error if any occurs
func QueryDocument(cli bce.Client, documentId string, queryParam *QueryDocumentParam) (*QueryDocumentResp, error) {
	req := &bce.BceRequest{}
	urlPath := fmt.Sprintf("/v2/document/%s", documentId)
	req.SetUri(urlPath)
	if queryParam != nil {
		req.SetParam("https", strconv.FormatBool(queryParam.Https))
	}
	req.SetMethod(http.GET)
	req.SetHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE)
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	result := &QueryDocumentResp{}
	if err := resp.ParseJsonBody(result); err != nil {
		return nil, err
	}
	return result, nil
}

// ReadDocument - get document token for client sdk
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - documentId: id of document in doc service
//     - readParam: expiration time of the doc's html
// RETURNS:
//     - *ReadDocumentResp
//     - error: the return error if any occurs
func ReadDocument(cli bce.Client, documentId string, readParam *ReadDocumentParam) (*ReadDocumentResp, error) {
	req := &bce.BceRequest{}
	urlPath := fmt.Sprintf("/v2/document/%s", documentId)
	req.SetUri(urlPath)
	req.SetParam("read", "")
	if readParam != nil {
		req.SetParam("expireInSeconds", strconv.FormatInt(readParam.ExpireInSeconds, 10))
	}
	req.SetMethod(http.GET)
	req.SetHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE)
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	result := &ReadDocumentResp{}
	if err := resp.ParseJsonBody(result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetImages - Get the list of images generated by the document conversion
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - documentId: id of document in doc service
// RETURNS:
//     - *ImagesListResp
//     - error: the return error if any occurs
func GetImages(cli bce.Client, documentId string) (*GetImagesResp, error) {
	req := &bce.BceRequest{}
	urlPath := fmt.Sprintf("/v2/document/%s", documentId)
	req.SetUri(urlPath)
	req.SetParam("getImages", "")
	req.SetMethod(http.GET)
	req.SetHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	result := &GetImagesResp{}
	if err := resp.ParseJsonBody(result); err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteDocument - delete document in doc service
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - documentId: id of document in doc service
// RETURNS:
//     - error: the return error if any occurs
func DeleteDocument(cli bce.Client, documentId string) error {
	req := &bce.BceRequest{}
	urlPath := fmt.Sprintf("/v2/document/%s", documentId)
	req.SetUri(urlPath)
	req.SetMethod(http.DELETE)
	req.SetHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}
	return nil
}

// ListDocuments - list all documents
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - param: the optional arguments to list documents
// RETURNS:
//     - *ListDocumentsResp: the result docments list structure
//     - error: nil if ok otherwise the specific error
func ListDocuments(cli bce.Client, listParam *ListDocumentsParam) (*ListDocumentsResp, error) {
	err := listParam.Check()
	if err != nil {
		return nil, err
	}

	req := &bce.BceRequest{}
	req.SetUri("/v2/document/")
	req.SetMethod(http.GET)
	if listParam.Status != "" {
		req.SetParam("status", string(listParam.Status))
	}
	if listParam.Marker != "" {
		req.SetParam("marker", listParam.Marker)
	}
	if listParam.MaxSize != 0 {
		req.SetParam("maxSize", strconv.FormatInt(listParam.MaxSize, 10))
	}

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	result := &ListDocumentsResp{}
	if err := resp.ParseJsonBody(result); err != nil {
		return nil, err
	}
	return result, nil
}
