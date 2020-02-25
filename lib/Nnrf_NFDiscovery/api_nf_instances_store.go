/*
 * NRF NFDiscovery Service
 *
 * NRF NFDiscovery  Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnrf_NFDiscovery

import (
	"free5gc/lib/openapi/common"
	"free5gc/lib/openapi/models"

	"context"
	"github.com/antihax/optional"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type NFInstancesStoreApiService service

/*
NFInstancesStoreApiService Search a collection of NF Instances
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param targetNfType Type of the target NF
 * @param requesterNfType Type of the requester NF
 * @param optional nil or *SearchNFInstancesParamOpts - Optional Parameters:
 * @param "ServiceNames" (optional.Interface of []string) -  Name of the service offered by the NF
 * @param "RequesterNfInstanceFqdn" (optional.String) -  FQDN of the requester NF
 * @param "TargetPlmnList" (optional.Interface of []models.PlmnId) -  Id of the PLMN where the target NF is located
 * @param "RequesterPlmnList" (optional.Interface of []models.PlmnId) -  Id of the PLMN where the NF issuing the Discovery request is located
 * @param "TargetNfInstanceId" (optional.Interface of string) -  Identity of the NF instance being discovered
 * @param "TargetNfFqdn" (optional.String) -  FQDN of the NF instance being discovered
 * @param "HnrfUri" (optional.String) -  Uri of the home NRF
 * @param "Snssais" (optional.Interface of []models.Snssai) -  Slice info of the target NF
 * @param "Dnn" (optional.String) -  Dnn supported by the BSF, SMF or UPF
 * @param "NsiList" (optional.Interface of []string) -  NSI IDs that are served by the services being discovered
 * @param "SmfServingArea" (optional.String) -
 * @param "Tai" (optional.Interface of models.Tai) -  Tracking Area Identity
 * @param "AmfRegionId" (optional.String) -  AMF Region Identity
 * @param "AmfSetId" (optional.String) -  AMF Set Identity
 * @param "Guami" (optional.Interface of models.Guami) -  Guami used to search for an appropriate AMF
 * @param "Supi" (optional.String) -  SUPI of the user
 * @param "UeIpv4Address" (optional.String) -  IPv4 address of the UE
 * @param "IpDomain" (optional.String) -  IP domain of the UE, which supported by BSF
 * @param "UeIpv6Prefix" (optional.Interface of string) -  IPv6 prefix of the UE
 * @param "PgwInd" (optional.Bool) -  Combined PGW-C and SMF or a standalone SMF
 * @param "Pgw" (optional.String) -  PGW FQDN of a combined PGW-C and SMF
 * @param "Gpsi" (optional.String) -  GPSI of the user
 * @param "ExternalGroupIdentity" (optional.String) -  external group identifier of the user
 * @param "DataSet" (optional.Interface of models.DataSetId) -  data set supported by the NF
 * @param "RoutingIndicator" (optional.String) -  routing indicator in SUCI
 * @param "GroupIdList" (optional.Interface of []string) -  Group IDs of the NFs being discovered
 * @param "DnaiList" (optional.Interface of []string) -  Data network access identifiers of the NFs being discovered
 * @param "SupportedFeatures" (optional.String) -  Features required to be supported by the target NF
 * @param "UpfIwkEpsInd" (optional.Bool) -  UPF supporting interworking with EPS or not
 * @param "ChfSupportedPlmn" (optional.Interface of models.PlmnId) -  PLMN ID supported by a CHF
 * @param "PreferredLocality" (optional.String) -  preferred target NF location
 * @param "AccessType" (optional.Interface of models.AccessType) -  AccessType supported by the target NF
 * @param "IfNoneMatch" (optional.String) -  Validator for conditional requests, as described in IETF RFC 7232, 3.2
@return models.SearchResult
*/

type SearchNFInstancesParamOpts struct {
	ServiceNames            optional.Interface
	RequesterNfInstanceFqdn optional.String
	TargetPlmnList          optional.Interface
	RequesterPlmnList       optional.Interface
	TargetNfInstanceId      optional.Interface
	TargetNfFqdn            optional.String
	HnrfUri                 optional.String
	Snssais                 optional.Interface
	Dnn                     optional.String
	NsiList                 optional.Interface
	SmfServingArea          optional.String
	Tai                     optional.Interface
	AmfRegionId             optional.String
	AmfSetId                optional.String
	Guami                   optional.Interface
	Supi                    optional.String
	UeIpv4Address           optional.String
	IpDomain                optional.String
	UeIpv6Prefix            optional.Interface
	PgwInd                  optional.Bool
	Pgw                     optional.String
	Gpsi                    optional.String
	ExternalGroupIdentity   optional.String
	DataSet                 optional.Interface
	RoutingIndicator        optional.String
	GroupIdList             optional.Interface
	DnaiList                optional.Interface
	SupportedFeatures       optional.String
	UpfIwkEpsInd            optional.Bool
	ChfSupportedPlmn        optional.Interface
	PreferredLocality       optional.String
	AccessType              optional.Interface
	IfNoneMatch             optional.String
}

func (a *NFInstancesStoreApiService) SearchNFInstances(ctx context.Context, targetNfType models.NfType, requesterNfType models.NfType, localVarOptionals *SearchNFInstancesParamOpts) (models.SearchResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  models.SearchResult
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/nf-instances"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("target-nf-type", common.ParameterToString(targetNfType, ""))
	localVarQueryParams.Add("requester-nf-type", common.ParameterToString(requesterNfType, ""))
	if localVarOptionals != nil && localVarOptionals.ServiceNames.IsSet() {
		localVarQueryParams.Add("service-names", common.ParameterToString(localVarOptionals.ServiceNames.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.RequesterNfInstanceFqdn.IsSet() {
		localVarQueryParams.Add("requester-nf-instance-fqdn", common.ParameterToString(localVarOptionals.RequesterNfInstanceFqdn.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TargetPlmnList.IsSet() {
		localVarQueryParams.Add("target-plmn-list", common.ParameterToString(localVarOptionals.TargetPlmnList.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.RequesterPlmnList.IsSet() {
		localVarQueryParams.Add("requester-plmn-list", common.ParameterToString(localVarOptionals.RequesterPlmnList.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.TargetNfInstanceId.IsSet() {
		localVarQueryParams.Add("target-nf-instance-id", common.ParameterToString(localVarOptionals.TargetNfInstanceId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TargetNfFqdn.IsSet() {
		localVarQueryParams.Add("target-nf-fqdn", common.ParameterToString(localVarOptionals.TargetNfFqdn.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.HnrfUri.IsSet() {
		localVarQueryParams.Add("hnrf-uri", common.ParameterToString(localVarOptionals.HnrfUri.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Snssais.IsSet() {
		localVarQueryParams.Add("snssais", common.ParameterToString(localVarOptionals.Snssais.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.Dnn.IsSet() {
		localVarQueryParams.Add("dnn", common.ParameterToString(localVarOptionals.Dnn.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.NsiList.IsSet() {
		localVarQueryParams.Add("nsi-list", common.ParameterToString(localVarOptionals.NsiList.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.SmfServingArea.IsSet() {
		localVarQueryParams.Add("smf-serving-area", common.ParameterToString(localVarOptionals.SmfServingArea.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Tai.IsSet() {
		localVarQueryParams.Add("tai", common.ParameterToString(localVarOptionals.Tai.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AmfRegionId.IsSet() {
		localVarQueryParams.Add("amf-region-id", common.ParameterToString(localVarOptionals.AmfRegionId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AmfSetId.IsSet() {
		localVarQueryParams.Add("amf-set-id", common.ParameterToString(localVarOptionals.AmfSetId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Guami.IsSet() {
		localVarQueryParams.Add("guami", common.ParameterToString(localVarOptionals.Guami.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Supi.IsSet() {
		localVarQueryParams.Add("supi", common.ParameterToString(localVarOptionals.Supi.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.UeIpv4Address.IsSet() {
		localVarQueryParams.Add("ue-ipv4-address", common.ParameterToString(localVarOptionals.UeIpv4Address.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IpDomain.IsSet() {
		localVarQueryParams.Add("ip-domain", common.ParameterToString(localVarOptionals.IpDomain.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.UeIpv6Prefix.IsSet() {
		localVarQueryParams.Add("ue-ipv6-prefix", common.ParameterToString(localVarOptionals.UeIpv6Prefix.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PgwInd.IsSet() {
		localVarQueryParams.Add("pgw-ind", common.ParameterToString(localVarOptionals.PgwInd.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Pgw.IsSet() {
		localVarQueryParams.Add("pgw", common.ParameterToString(localVarOptionals.Pgw.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Gpsi.IsSet() {
		localVarQueryParams.Add("gpsi", common.ParameterToString(localVarOptionals.Gpsi.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ExternalGroupIdentity.IsSet() {
		localVarQueryParams.Add("external-group-identity", common.ParameterToString(localVarOptionals.ExternalGroupIdentity.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.DataSet.IsSet() {
		localVarQueryParams.Add("data-set", common.ParameterToString(localVarOptionals.DataSet.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RoutingIndicator.IsSet() {
		localVarQueryParams.Add("routing-indicator", common.ParameterToString(localVarOptionals.RoutingIndicator.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.GroupIdList.IsSet() {
		localVarQueryParams.Add("group-id-list", common.ParameterToString(localVarOptionals.GroupIdList.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.DnaiList.IsSet() {
		localVarQueryParams.Add("dnai-list", common.ParameterToString(localVarOptionals.DnaiList.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.SupportedFeatures.IsSet() {
		localVarQueryParams.Add("supported-features", common.ParameterToString(localVarOptionals.SupportedFeatures.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.UpfIwkEpsInd.IsSet() {
		localVarQueryParams.Add("upf-iwk-eps-ind", common.ParameterToString(localVarOptionals.UpfIwkEpsInd.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ChfSupportedPlmn.IsSet() {
		localVarQueryParams.Add("chf-supported-plmn", common.ParameterToString(localVarOptionals.ChfSupportedPlmn.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PreferredLocality.IsSet() {
		localVarQueryParams.Add("preferred-locality", common.ParameterToString(localVarOptionals.PreferredLocality.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AccessType.IsSet() {
		localVarQueryParams.Add("access-type", common.ParameterToString(localVarOptionals.AccessType.Value(), ""))
	}

	localVarHTTPContentTypes := []string{"application/json"}

	localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := common.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	if localVarOptionals != nil && localVarOptionals.IfNoneMatch.IsSet() {
		localVarHeaderParams["If-None-Match"] = common.ParameterToString(localVarOptionals.IfNoneMatch.Value(), "")
	}

	r, err := common.PrepareRequest(ctx, a.client.cfg, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := common.CallAPI(a.client.cfg, r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	apiError := common.GenericOpenAPIError{
		RawBody:     localVarBody,
		ErrorStatus: localVarHTTPResponse.Status,
	}

	switch localVarHTTPResponse.StatusCode {
	case 200:
		err = common.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
		}
		return localVarReturnValue, localVarHTTPResponse, nil
	case 307:
		return localVarReturnValue, localVarHTTPResponse, nil
	case 400:
		var v models.ProblemDetails
		err = common.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHTTPResponse, apiError
	case 403:
		var v models.ProblemDetails
		err = common.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHTTPResponse, apiError
	case 404:
		var v models.ProblemDetails
		err = common.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHTTPResponse, apiError
	case 411:
		var v models.ProblemDetails
		err = common.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHTTPResponse, apiError
	case 413:
		var v models.ProblemDetails
		err = common.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHTTPResponse, apiError
	case 415:
		var v models.ProblemDetails
		err = common.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHTTPResponse, apiError
	case 500:
		var v models.ProblemDetails
		err = common.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHTTPResponse, apiError
	case 501:
		var v models.ProblemDetails
		err = common.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHTTPResponse, apiError
	case 503:
		var v models.ProblemDetails
		err = common.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHTTPResponse, apiError
	default:
		return localVarReturnValue, localVarHTTPResponse, nil
	}
}
