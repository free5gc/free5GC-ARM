/*
 * NSSF NS Selection
 *
 * NSSF Network Slice Selection Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package nssf_producer

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"free5gc/lib/http_wrapper"
	. "free5gc/lib/openapi/models"
	"free5gc/src/nssf/logger"
	"free5gc/src/nssf/nssf_handler/nssf_message"
	"free5gc/src/nssf/plugin"
	"free5gc/src/nssf/util"
)

// Parse NSSelectionGet query parameter
func parseQueryParameter(query url.Values) (p plugin.NsselectionQueryParameter, err error) {

	if query.Get("nf-type") != "" {
		p.NfType = new(NfType)
		*p.NfType = NfType(query.Get("nf-type"))
	}

	p.NfId = query.Get("nf-id")

	if query.Get("slice-info-request-for-registration") != "" {
		p.SliceInfoRequestForRegistration = new(SliceInfoForRegistration)
		err = json.NewDecoder(strings.NewReader(query.Get("slice-info-request-for-registration"))).Decode(p.SliceInfoRequestForRegistration)
		if err != nil {
			return
		}
	}

	if query.Get("slice-info-request-for-pdu-session") != "" {
		p.SliceInfoRequestForPduSession = new(SliceInfoForPduSession)
		err = json.NewDecoder(strings.NewReader(query.Get("slice-info-request-for-pdu-session"))).Decode(p.SliceInfoRequestForPduSession)
		if err != nil {
			return
		}
	}

	if query.Get("home-plmn-id") != "" {
		p.HomePlmnId = new(PlmnId)
		err = json.NewDecoder(strings.NewReader(query.Get("home-plmn-id"))).Decode(p.HomePlmnId)
		if err != nil {
			return
		}
	}

	if query.Get("tai") != "" {
		p.Tai = new(Tai)
		err = json.NewDecoder(strings.NewReader(query.Get("tai"))).Decode(p.Tai)
		if err != nil {
			return
		}
	}

	if query.Get("supported-features") != "" {
		p.SupportedFeatures = query.Get("supported-features")
	}

	return
}

// Check if the NF service consumer is authorized
// TODO: Check if the NF service consumer is legal with local configuration, or possibly after querying NRF through
//       `nf-id` e.g. Whether the V-NSSF is authorized
func checkNfServiceConsumer(nfType NfType) error {
	if nfType != NfType_AMF && nfType != NfType_NSSF {
		return fmt.Errorf("`nf-type`:'%s' is not authorized to retrieve the slice selection information", string(nfType))
	}

	return nil
}

// NSSelectionGet - Retrieve the Network Slice Selection Information
func NSSelectionGet(responseChan chan nssf_message.HandlerResponseMessage, query url.Values) {

	logger.Nsselection.Infof("Request received - NSSelectionGet")

	var (
		isValidRequest bool = true
		status         int
		a              AuthorizedNetworkSliceInfo
		d              ProblemDetails
	)

	// TODO: Record request times of the NF service consumer and response with ProblemDetails of 429 Too Many Requests
	//       if the consumer has sent too many requests in a configured amount of time
	// TODO: Check URI length and response with ProblemDetails of 414 URI Too Long if URI is too long

	// Parse query parameter
	p, err := parseQueryParameter(query)
	if err != nil {
		problemDetail := "[Query Parameter] " + err.Error()
		status = http.StatusBadRequest
		d = ProblemDetails{
			Title:  util.MALFORMED_REQUEST,
			Status: http.StatusBadRequest,
			Detail: problemDetail,
		}
		isValidRequest = false
	}

	// Check permission of NF service consumer
	err = checkNfServiceConsumer(*p.NfType)
	if err != nil {
		problemDetail := err.Error()
		status = http.StatusForbidden
		d = ProblemDetails{
			Title:  util.UNAUTHORIZED_CONSUMER,
			Status: http.StatusForbidden,
			Detail: problemDetail,
		}
		isValidRequest = false
	}

	if isValidRequest {
		if p.SliceInfoRequestForRegistration != nil {
			// Network slice information is requested during the Registration procedure
			status = nsselectionForRegistration(p, &a, &d)
		} else {
			// Network slice information is requested during the PDU session establishment procedure
			status = nsselectionForPduSession(p, &a, &d)
		}
	}

	if status == http.StatusOK {
		responseChan <- nssf_message.HandlerResponseMessage{
			HttpResponse: &http_wrapper.Response{
				Header: nil,
				Status: status,
				Body:   a,
			},
		}
	} else {
		responseChan <- nssf_message.HandlerResponseMessage{
			HttpResponse: &http_wrapper.Response{
				Header: nil,
				Status: status,
				Body:   d,
			},
		}
	}
}
