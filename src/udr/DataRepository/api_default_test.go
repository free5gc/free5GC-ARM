/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package DataRepository_test

import (
	"context"
	"flag"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/urfave/cli"
	"free5gc/lib/MongoDBLibrary"
	"free5gc/lib/Nudr_DataRepository"
	"free5gc/lib/openapi/models"
	"free5gc/src/udr/factory"
	"free5gc/src/udr/udr_handler"
	"free5gc/src/udr/udr_service"

	"github.com/antihax/optional"
	"github.com/google/go-cmp/cmp"
)

var Opt cmp.Option

var serverFlag = 0
var mongodbFlag = 0
var UDR = &udr_service.UDR{}

func runTestServer(t *testing.T) {
	if serverFlag == 0 {
		flag := flag.FlagSet{}
		cli := cli.NewContext(nil, &flag, nil)
		UDR.Initialize(cli)
		go func() {
			UDR.Start()
		}()
		serverFlag = 1

		// Run channel Handle Function for UDR
		go udr_handler.Handle()
		time.Sleep(100 * time.Millisecond)
	}
}

func connectMongoDB(t *testing.T) {
	if mongodbFlag == 0 {
		// Connect to MongoDB
		mongodb := factory.UdrConfig.Configuration.Mongodb
		SetMongoDB(mongodb.Name, mongodb.Url)
		Client = MongoDBLibrary.Client
		mongodbFlag = 1

		alwaysEqual := cmp.Comparer(func(_, _ interface{}) bool { return true })

		// This option handles slices and maps of any type.
		optTmp := cmp.FilterValues(func(x, y interface{}) bool {
			vx, vy := reflect.ValueOf(x), reflect.ValueOf(y)
			// fmt.Println(vx.Kind(), "and", vy.Kind())
			return (vx.IsValid() && vy.IsValid() && vx.Type() == vy.Type()) &&
				(vx.Kind() == reflect.Map || vx.Kind() == reflect.Slice)
		}, alwaysEqual)

		Opt = optTmp
	}
}

func setTestClient(t *testing.T) *Nudr_DataRepository.APIClient {
	sbi := factory.UdrConfig.Configuration.Sbi
	clientUrl := fmt.Sprintf("https://%s:%d", sbi.IPv4Addr, sbi.Port)

	configuration := Nudr_DataRepository.NewConfiguration()
	configuration.SetBasePath(clientUrl)
	return Nudr_DataRepository.NewAPIClient(configuration)
}

// ApplicationDataInfluenceDataGet -
func TestApplicationDataInfluenceDataGet(t *testing.T) {
	// c.JSON(http.StatusOK, gin.H{})
}

// ApplicationDataInfluenceDataInfluenceIdDelete -
func TestApplicationDataInfluenceDataInfluenceIdDelete(t *testing.T) {
	// c.JSON(http.StatusOK, gin.H{})
}

// ApplicationDataInfluenceDataInfluenceIdPatch -
func TestApplicationDataInfluenceDataInfluenceIdPatch(t *testing.T) {
	// c.JSON(http.StatusOK, gin.H{})
}

// ApplicationDataInfluenceDataInfluenceIdPut -
func TestApplicationDataInfluenceDataInfluenceIdPut(t *testing.T) {
	// c.JSON(http.StatusOK, gin.H{})
}

// ApplicationDataInfluenceDataSubsToNotifyGet -
func TestApplicationDataInfluenceDataSubsToNotifyGet(t *testing.T) {
	// c.JSON(http.StatusOK, gin.H{})
}

// ApplicationDataInfluenceDataSubsToNotifyPost -
func TestApplicationDataInfluenceDataSubsToNotifyPost(t *testing.T) {
	// c.JSON(http.StatusOK, gin.H{})
}

// ApplicationDataInfluenceDataSubsToNotifySubscriptionIdDelete -
func TestApplicationDataInfluenceDataSubsToNotifySubscriptionIdDelete(t *testing.T) {
	// c.JSON(http.StatusOK, gin.H{})
}

// ApplicationDataInfluenceDataSubsToNotifySubscriptionIdGet -
func TestApplicationDataInfluenceDataSubsToNotifySubscriptionIdGet(t *testing.T) {
	// c.JSON(http.StatusOK, gin.H{})
}

// ApplicationDataInfluenceDataSubsToNotifySubscriptionIdPut -
func TestApplicationDataInfluenceDataSubsToNotifySubscriptionIdPut(t *testing.T) {
	// c.JSON(http.StatusOK, gin.H{})
}

// ApplicationDataPfdsAppIdDelete -
func TestApplicationDataPfdsAppIdDelete(t *testing.T) {
	// c.JSON(http.StatusOK, gin.H{})
}

// ApplicationDataPfdsAppIdGet -
func TestApplicationDataPfdsAppIdGet(t *testing.T) {
	// c.JSON(http.StatusOK, gin.H{})
}

// ApplicationDataPfdsAppIdPut -
func TestApplicationDataPfdsAppIdPut(t *testing.T) {
	// c.JSON(http.StatusOK, gin.H{})
}

// ApplicationDataPfdsGet -
func TestApplicationDataPfdsGet(t *testing.T) {
	// c.JSON(http.StatusOK, gin.H{})
}

// ExposureDataSubsToNotifyPost -
func TestExposureDataSubsToNotifyPost(t *testing.T) {
	// c.JSON(http.StatusOK, gin.H{})
}

// ExposureDataSubsToNotifySubIdDelete - Deletes a subcription for notifications
func TestExposureDataSubsToNotifySubIdDelete(t *testing.T) {
	// c.JSON(http.StatusOK, gin.H{})
}

// ExposureDataSubsToNotifySubIdPut - updates a subcription for notifications
func TestExposureDataSubsToNotifySubIdPut(t *testing.T) {
	// c.JSON(http.StatusOK, gin.H{})
}

// PolicyDataBdtDataBdtReferenceIdDelete -
func TestPolicyDataBdtDataBdtReferenceIdDelete(t *testing.T) {
	// c.JSON(http.StatusOK, gin.H{})
}

// PolicyDataBdtDataBdtReferenceIdGet -
func TestPolicyDataBdtDataBdtReferenceIdGet(t *testing.T) {
	// c.JSON(http.StatusOK, gin.H{})
}

// PolicyDataBdtDataBdtReferenceIdPut -
func TestPolicyDataBdtDataBdtReferenceIdPut(t *testing.T) {
	// c.JSON(http.StatusOK, gin.H{})
}

// PolicyDataBdtDataGet -
func TestPolicyDataBdtDataGet(t *testing.T) {
	// c.JSON(http.StatusOK, gin.H{})
}

// PolicyDataPlmnsPlmnIdUePolicySetGet -
func TestPolicyDataPlmnsPlmnIdUePolicySetGet(t *testing.T) {
	// c.JSON(http.StatusOK, gin.H{})
}

// PolicyDataSponsorConnectivityDataSponsorIdGet -
func TestPolicyDataSponsorConnectivityDataSponsorIdGet(t *testing.T) {
	// c.JSON(http.StatusOK, gin.H{})
}

// PolicyDataSubsToNotifyPost -
func TestPolicyDataSubsToNotifyPost(t *testing.T) {
	// c.JSON(http.StatusOK, gin.H{})
}

// PolicyDataSubsToNotifySubsIdDelete -
func TestPolicyDataSubsToNotifySubsIdDelete(t *testing.T) {
	// c.JSON(http.StatusOK, gin.H{})
}

// PolicyDataSubsToNotifySubsIdPut -
func TestPolicyDataSubsToNotifySubsIdPut(t *testing.T) {
	// c.JSON(http.StatusOK, gin.H{})
}

// PolicyDataUesUeIdAmDataGet -
func TestPolicyDataUesUeIdAmDataGet(t *testing.T) {
	runTestServer(t)

	connectMongoDB(t)

	// Drop old data
	collection := Client.Database("free5gc").Collection("policyData.ues.amData")
	collection.DeleteOne(context.TODO(), bson.M{"ueId": "imsi-0123456789"})

	// Set client and set url
	client := setTestClient(t)

	// Set test data
	ueId := "imsi-0123456789"
	testData := models.AmPolicyData{
		SubscCats: []string{"abc", "bcd"},
	}
	insertTestData := toBsonM(testData)
	insertTestData["ueId"] = ueId
	collection.InsertOne(context.TODO(), insertTestData)

	{
		// Check test data (Use RESTful GET)
		amPolicyData, res, err := client.DefaultApi.PolicyDataUesUeIdAmDataGet(context.TODO(), ueId)
		if err != nil {
			log.Panic(err)
		}

		if status := res.StatusCode; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
		if reflect.DeepEqual(testData, amPolicyData) != true {
			t.Errorf("handler returned unexpected body: got %v want %v",
				amPolicyData, testData)
		}
	}

	// Clean test data
	collection.DeleteOne(context.TODO(), bson.M{"ueId": "imsi-0123456789"})

	// TEST END
}

// PolicyDataUesUeIdOperatorSpecificDataGet -
func TestPolicyDataUesUeIdOperatorSpecificDataGet(t *testing.T) {
	runTestServer(t)

	connectMongoDB(t)

	// Drop old data
	collection := Client.Database("free5gc").Collection("policyData.ues.operatorSpecificData")
	collection.DeleteOne(context.TODO(), bson.M{"ueId": "imsi-0123456789"})

	// Set client and set url
	client := setTestClient(t)

	// Set test data
	ueId := "imsi-0123456789"
	testData := map[string]models.OperatorSpecificDataContainer{
		"name_test": {
			StringTypeElements: map[string]string{
				"a": "b",
			},
		},
	}
	insertTestData := map[string]interface{}{"operatorSpecificDataContainerMap": testData}
	insertTestData["ueId"] = ueId
	collection.InsertOne(context.TODO(), insertTestData)

	{
		// Check test data (Use RESTful GET)
		var policyDataUesUeIdOperatorSpecificDataGetParamOpts Nudr_DataRepository.PolicyDataUesUeIdOperatorSpecificDataGetParamOpts
		operatorSpecificDataContainerMap, res, err := client.DefaultApi.PolicyDataUesUeIdOperatorSpecificDataGet(context.TODO(), ueId, &policyDataUesUeIdOperatorSpecificDataGetParamOpts)
		if err != nil {
			log.Panic(err)
		}

		if status := res.StatusCode; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
		if reflect.DeepEqual(testData, operatorSpecificDataContainerMap) != true {
			t.Errorf("handler returned unexpected body: got %v want %v",
				operatorSpecificDataContainerMap, testData)
		}
	}

	// Clean test data
	collection.DeleteOne(context.TODO(), bson.M{"ueId": "imsi-0123456789"})

	// TEST END
}

// PolicyDataUesUeIdOperatorSpecificDataPatch - Need to be fixed
func TestPolicyDataUesUeIdOperatorSpecificDataPatch(t *testing.T) {
	runTestServer(t)

	connectMongoDB(t)

	// Drop old data
	collection := Client.Database("free5gc").Collection("policyData.ues.operatorSpecificData")
	collection.DeleteOne(context.TODO(), bson.M{"ueId": "imsi-0123456789"})

	// Set client and set url
	client := setTestClient(t)

	// Set test data
	ueId := "imsi-0123456789"
	testData := map[string]models.OperatorSpecificDataContainer{
		"name_test": {
			StringTypeElements: map[string]string{
				"a": "b",
			},
		},
	}
	insertTestData := map[string]interface{}{"operatorSpecificDataContainerMap": testData}
	insertTestData["ueId"] = ueId
	collection.InsertOne(context.TODO(), insertTestData)

	{
		// Check test data (Use RESTful GET)
		var policyDataUesUeIdOperatorSpecificDataGetParamOpts Nudr_DataRepository.PolicyDataUesUeIdOperatorSpecificDataGetParamOpts
		operatorSpecificDataContainerMap, res, err := client.DefaultApi.PolicyDataUesUeIdOperatorSpecificDataGet(context.TODO(), ueId, &policyDataUesUeIdOperatorSpecificDataGetParamOpts)
		if err != nil {
			log.Panic(err)
		}

		if status := res.StatusCode; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
		if reflect.DeepEqual(testData, operatorSpecificDataContainerMap) != true {
			t.Errorf("handler returned unexpected body: got %v want %v",
				operatorSpecificDataContainerMap, testData)
		}
	}

	patchItemArray := []models.PatchItem{
		{
			Op:   models.PatchOperation_REPLACE,
			Path: "/name_test",
			Value: models.OperatorSpecificDataContainer{
				StringTypeElements: map[string]string{
					"c": "d",
				},
			},
		},
	}
	patchData := map[string]models.OperatorSpecificDataContainer{
		"name_test": {
			StringTypeElements: map[string]string{
				"c": "d",
			},
		},
	}

	{
		// Patch data (Use RESTful PATCH)
		res, err := client.DefaultApi.PolicyDataUesUeIdOperatorSpecificDataPatch(context.TODO(), ueId, patchItemArray)
		if err != nil {
			log.Panic(err)
		}

		if status := res.StatusCode; status != http.StatusNoContent {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusNoContent)
		}
	}

	{
		// Check patch data (Use RESTful GET)
		var policyDataUesUeIdOperatorSpecificDataGetParamOpts Nudr_DataRepository.PolicyDataUesUeIdOperatorSpecificDataGetParamOpts
		operatorSpecificDataContainerMap, res, err := client.DefaultApi.PolicyDataUesUeIdOperatorSpecificDataGet(context.TODO(), ueId, &policyDataUesUeIdOperatorSpecificDataGetParamOpts)
		if err != nil {
			log.Panic(err)
		}

		if status := res.StatusCode; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		if reflect.DeepEqual(patchData, operatorSpecificDataContainerMap) != true {
			t.Errorf("handler returned unexpected body: got %v want %v",
				operatorSpecificDataContainerMap, patchData)
		}
	}

	// Clean test data
	collection.DeleteOne(context.TODO(), bson.M{"ueId": "imsi-0123456789"})

	// TEST END
}

// PolicyDataUesUeIdOperatorSpecificDataPut -
func TestPolicyDataUesUeIdOperatorSpecificDataPut(t *testing.T) {
	runTestServer(t)

	connectMongoDB(t)

	// Drop old data
	collection := Client.Database("free5gc").Collection("policyData.ues.operatorSpecificData")
	collection.DeleteOne(context.TODO(), bson.M{"ueId": "imsi-0123456789"})

	// Set client and set url
	client := setTestClient(t)

	// Set test data
	ueId := "imsi-0123456789"
	testData := map[string]models.OperatorSpecificDataContainer{
		"name_test": {
			StringTypeElements: map[string]string{
				"a": "b",
			},
		},
	}

	{
		// Insert test data (Use RESTful PUT)
		_, res, err := client.DefaultApi.PolicyDataUesUeIdOperatorSpecificDataPut(context.TODO(), ueId, testData)
		if err != nil {
			log.Panic(err)
		}

		if status := res.StatusCode; status != http.StatusNoContent && status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v or %v",
				status, http.StatusNoContent, http.StatusOK)
		}
	}

	{
		// Check test data (Use RESTful GET)
		var policyDataUesUeIdOperatorSpecificDataGetParamOpts Nudr_DataRepository.PolicyDataUesUeIdOperatorSpecificDataGetParamOpts
		operatorSpecificDataContainerMap, res, err := client.DefaultApi.PolicyDataUesUeIdOperatorSpecificDataGet(context.TODO(), ueId, &policyDataUesUeIdOperatorSpecificDataGetParamOpts)
		if err != nil {
			log.Panic(err)
		}

		if status := res.StatusCode; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
		if reflect.DeepEqual(testData, operatorSpecificDataContainerMap) != true {
			t.Errorf("handler returned unexpected body: got %v want %v",
				operatorSpecificDataContainerMap, testData)
		}
	}

	// Clean test data
	collection.DeleteOne(context.TODO(), bson.M{"ueId": "imsi-0123456789"})

	// TEST END
}

// PolicyDataUesUeIdSmDataGet -
func TestPolicyDataUesUeIdSmDataGet(t *testing.T) {
	runTestServer(t)

	connectMongoDB(t)

	// Drop old data
	collection := Client.Database("free5gc").Collection("policyData.ues.smData")
	collection.DeleteOne(context.TODO(), bson.M{"ueId": "imsi-0123456789"})

	// Set client and set url
	client := setTestClient(t)

	// Set test data
	ueId := "imsi-0123456789"
	testData := models.SmPolicyData{
		SmPolicySnssaiData: map[string]models.SmPolicySnssaiData{
			"Snssai": {
				Snssai: &models.Snssai{
					Sst: 1,
				},
			},
		},
	}
	insertTestData := toBsonM(testData)
	insertTestData["ueId"] = ueId
	collection.InsertOne(context.TODO(), insertTestData)

	{
		// Check test data (Use RESTful GET)
		var policyDataUesUeIdSmDataGetParamOpts Nudr_DataRepository.PolicyDataUesUeIdSmDataGetParamOpts
		smPolicyData, res, err := client.DefaultApi.PolicyDataUesUeIdSmDataGet(context.TODO(), ueId, &policyDataUesUeIdSmDataGetParamOpts)
		if err != nil {
			log.Panic(err)
		}

		if status := res.StatusCode; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
		if cmp.Equal(testData, smPolicyData, Opt) != true {
			t.Errorf("handler returned unexpected body: got %v want %v",
				smPolicyData, testData)
		}
	}

	// Clean test data
	collection.DeleteOne(context.TODO(), bson.M{"ueId": "imsi-0123456789"})

	// TEST END
}

// PolicyDataUesUeIdSmDataPatch - Need to be fixed
func TmpTestPolicyDataUesUeIdSmDataPatch(t *testing.T) {
	runTestServer(t)

	connectMongoDB(t)

	// Drop old data
	collection := Client.Database("free5gc").Collection("policyData.ues.smData")
	collection.DeleteOne(context.TODO(), bson.M{"ueId": "imsi-0123456789"})

	// Set client and set url
	client := setTestClient(t)

	// Set test data
	ueId := "imsi-0123456789"
	testData := models.SmPolicyData{
		SmPolicySnssaiData: map[string]models.SmPolicySnssaiData{
			"Snssai": {
				Snssai: &models.Snssai{
					Sst: 1,
				},
			},
		},
	}
	insertTestData := toBsonM(testData)
	insertTestData["ueId"] = ueId
	collection.InsertOne(context.TODO(), insertTestData)

	{
		// Check test data (Use RESTful GET)
		var policyDataUesUeIdSmDataGetParamOpts Nudr_DataRepository.PolicyDataUesUeIdSmDataGetParamOpts
		uePolicySet, res, err := client.DefaultApi.PolicyDataUesUeIdSmDataGet(context.TODO(), ueId, &policyDataUesUeIdSmDataGetParamOpts)
		if err != nil {
			log.Panic(err)
		}

		if status := res.StatusCode; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		if reflect.DeepEqual(testData, uePolicySet) != true {
			t.Errorf("handler returned unexpected body: got %v want %v",
				uePolicySet, testData)
		}
	}

	patchData := models.SmPolicyData{
		SmPolicySnssaiData: map[string]models.SmPolicySnssaiData{
			"Snssai": {
				Snssai: &models.Snssai{
					Sst: 12,
				},
			},
		},
	}

	{
		// Patch data (Use RESTful PATCH)
		var policyDataUesUeIdSmDataPatchParamOpts Nudr_DataRepository.PolicyDataUesUeIdSmDataPatchParamOpts
		policyDataUesUeIdSmDataPatchParamOpts.RequestBody = optional.NewInterface(patchData)
		res, err := client.DefaultApi.PolicyDataUesUeIdSmDataPatch(context.TODO(), ueId, &policyDataUesUeIdSmDataPatchParamOpts)
		if err != nil {
			log.Panic(err)
		}

		if status := res.StatusCode; status != http.StatusNoContent {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusNoContent)
		}
	}

	{
		// Check patch data (Use RESTful GET)
		var policyDataUesUeIdSmDataGetParamOpts Nudr_DataRepository.PolicyDataUesUeIdSmDataGetParamOpts
		uePolicySet, res, err := client.DefaultApi.PolicyDataUesUeIdSmDataGet(context.TODO(), ueId, &policyDataUesUeIdSmDataGetParamOpts)
		if err != nil {
			log.Panic(err)
		}

		if status := res.StatusCode; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		if reflect.DeepEqual(patchData, uePolicySet) != true {
			t.Errorf("handler returned unexpected body: got %v want %v",
				uePolicySet, patchData)
		}
	}

	// Clean test data
	collection.DeleteOne(context.TODO(), bson.M{"ueId": "imsi-0123456789"})

	// TEST END
}

// PolicyDataUesUeIdSmDataUsageMonIdDelete -
func TestPolicyDataUesUeIdSmDataUsageMonIdDelete(t *testing.T) {
	runTestServer(t)

	connectMongoDB(t)

	// Drop old data
	collection := Client.Database("free5gc").Collection("policyData.ues.smData.usageMonData")
	collection.DeleteOne(context.TODO(), bson.M{"ueId": "imsi-0123456789"})

	// Set client and set url
	client := setTestClient(t)

	// Set test data
	ueId := "imsi-0123456789"
	usageMonId := "usageMonId_test"
	testData := models.UsageMonData{
		LimitId: "LimitId_test",
	}
	insertTestData := toBsonM(testData)
	insertTestData["ueId"] = ueId
	insertTestData["usageMonId"] = usageMonId
	collection.InsertOne(context.TODO(), insertTestData)

	{
		// Check test data (Use RESTful GET)
		usageMonData, res, err := client.DefaultApi.PolicyDataUesUeIdSmDataUsageMonIdGet(context.TODO(), ueId, usageMonId)
		if err != nil {
			log.Panic(err)
		}

		if status := res.StatusCode; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
		if cmp.Equal(testData, usageMonData, Opt) != true {
			t.Errorf("handler returned unexpected body: got %v want %v",
				usageMonData, testData)
		}
	}

	{
		// delete test data (Use RESTful DELETE)
		res, err := client.DefaultApi.PolicyDataUesUeIdSmDataUsageMonIdDelete(context.TODO(), ueId, usageMonId)
		if err != nil {
			log.Panic(err)
		}

		if status := res.StatusCode; status != http.StatusNoContent {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusNoContent)
		}
	}

	{
		// Check test data (Use RESTful GET)
		usageMonData, res, err := client.DefaultApi.PolicyDataUesUeIdSmDataUsageMonIdGet(context.TODO(), ueId, usageMonId)
		if err != nil {
			log.Panic(err)
		}

		if status := res.StatusCode; status != http.StatusNoContent {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusNoContent)
		}

		var empty models.UsageMonData
		if reflect.DeepEqual(empty, usageMonData) != true {
			t.Errorf("handler returned unexpected body: got %v want %v",
				usageMonData, empty)
		}
	}

	// Clean test data
	collection.DeleteOne(context.TODO(), bson.M{"ueId": "imsi-0123456789"})

	// TEST END
}

// PolicyDataUesUeIdSmDataUsageMonIdGet -
func TestPolicyDataUesUeIdSmDataUsageMonIdGet(t *testing.T) {
	runTestServer(t)

	connectMongoDB(t)

	// Drop old data
	collection := Client.Database("free5gc").Collection("policyData.ues.smData.usageMonData")
	collection.DeleteOne(context.TODO(), bson.M{"ueId": "imsi-0123456789"})

	// Set client and set url
	client := setTestClient(t)

	// Set test data
	ueId := "imsi-0123456789"
	usageMonId := "usageMonId_test"
	testData := models.UsageMonData{
		LimitId: "LimitId_test",
	}
	insertTestData := toBsonM(testData)
	insertTestData["ueId"] = ueId
	insertTestData["usageMonId"] = usageMonId
	collection.InsertOne(context.TODO(), insertTestData)

	{
		// Check test data (Use RESTful GET)
		usageMonData, res, err := client.DefaultApi.PolicyDataUesUeIdSmDataUsageMonIdGet(context.TODO(), ueId, usageMonId)
		if err != nil {
			log.Panic(err)
		}

		if status := res.StatusCode; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
		if cmp.Equal(testData, usageMonData, Opt) != true {
			t.Errorf("handler returned unexpected body: got %v want %v",
				usageMonData, testData)
		}
	}

	// Clean test data
	collection.DeleteOne(context.TODO(), bson.M{"ueId": "imsi-0123456789"})

	// TEST END
}

// PolicyDataUesUeIdSmDataUsageMonIdPut -
func TestPolicyDataUesUeIdSmDataUsageMonIdPut(t *testing.T) {
	runTestServer(t)

	connectMongoDB(t)

	// Drop old data
	collection := Client.Database("free5gc").Collection("policyData.ues.smData.usageMonData")
	collection.DeleteOne(context.TODO(), bson.M{"ueId": "imsi-0123456789"})

	// Set client and set url
	client := setTestClient(t)

	// Set test data
	ueId := "imsi-0123456789"
	usageMonId := "usageMonId_test"
	testData := models.UsageMonData{
		LimitId: "LimitId_test",
	}

	{
		// Insert test data (Use RESTful PUT)
		var policyDataUesUeIdSmDataUsageMonIdPutParamOpts Nudr_DataRepository.PolicyDataUesUeIdSmDataUsageMonIdPutParamOpts
		policyDataUesUeIdSmDataUsageMonIdPutParamOpts.UsageMonData = optional.NewInterface(testData)
		res, err := client.DefaultApi.PolicyDataUesUeIdSmDataUsageMonIdPut(context.TODO(), ueId, usageMonId, &policyDataUesUeIdSmDataUsageMonIdPutParamOpts)
		if err != nil {
			log.Panic(err)
		}

		if status := res.StatusCode; status != http.StatusNoContent && status != http.StatusCreated {
			t.Errorf("handler returned wrong status code: got %v want %v or %v",
				status, http.StatusNoContent, http.StatusCreated)
		}
	}

	{
		// Check test data (Use RESTful GET)
		usageMonData, res, err := client.DefaultApi.PolicyDataUesUeIdSmDataUsageMonIdGet(context.TODO(), ueId, usageMonId)
		if err != nil {
			log.Panic(err)
		}

		if status := res.StatusCode; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		if cmp.Equal(testData, usageMonData, Opt) != true {
			t.Errorf("handler returned unexpected body: got %v want %v",
				usageMonData, testData)
		}
	}

	// Clean test data
	collection.DeleteOne(context.TODO(), bson.M{"ueId": "imsi-0123456789"})

	// TEST END
}

// PolicyDataUesUeIdUePolicySetGet -
func TestPolicyDataUesUeIdUePolicySetGet(t *testing.T) {
	runTestServer(t)

	connectMongoDB(t)

	// Drop old data
	collection := Client.Database("free5gc").Collection("policyData.ues.uePolicySet")
	collection.DeleteOne(context.TODO(), bson.M{"ueId": "imsi-0123456789"})

	// Set client and set url
	client := setTestClient(t)

	// Set test data
	ueId := "imsi-0123456789"
	testData := models.UePolicySet{
		SubscCats: []string{"abc", "bcd"},
		Upsis:     []string{"abcd", "bcde"},
	}
	insertTestData := toBsonM(testData)
	insertTestData["ueId"] = ueId
	collection.InsertOne(context.TODO(), insertTestData)

	{
		// Check test data (Use RESTful GET)
		uePolicySet, res, err := client.DefaultApi.PolicyDataUesUeIdUePolicySetGet(context.TODO(), ueId)
		if err != nil {
			log.Panic(err)
		}

		if status := res.StatusCode; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
		if cmp.Equal(testData, uePolicySet, Opt) != true {
			t.Errorf("handler returned unexpected body: got %v want %v",
				uePolicySet, testData)
		}
	}

	// Clean test data
	collection.DeleteOne(context.TODO(), bson.M{"ueId": "imsi-0123456789"})

	// TEST END
}

// PolicyDataUesUeIdUePolicySetPatch -
func TestPolicyDataUesUeIdUePolicySetPatch(t *testing.T) {
	runTestServer(t)

	connectMongoDB(t)

	// Drop old data
	collection := Client.Database("free5gc").Collection("policyData.ues.uePolicySet")
	collection.DeleteOne(context.TODO(), bson.M{"ueId": "imsi-0123456789"})

	// Set client and set url
	client := setTestClient(t)

	// Set test data
	ueId := "imsi-0123456789"
	testData := models.UePolicySet{
		SubscCats: []string{"abc", "bcd"},
		Upsis:     []string{"abcd", "bcde"},
	}
	insertTestData := toBsonM(testData)
	insertTestData["ueId"] = ueId
	collection.InsertOne(context.TODO(), insertTestData)

	{
		// Check test data (Use RESTful GET)
		uePolicySet, res, err := client.DefaultApi.PolicyDataUesUeIdUePolicySetGet(context.TODO(), ueId)
		if err != nil {
			log.Panic(err)
		}

		if status := res.StatusCode; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		if cmp.Equal(testData, uePolicySet, Opt) != true {
			t.Errorf("handler returned unexpected body: got %v want %v",
				uePolicySet, testData)
		}
	}

	patchData := models.UePolicySet{
		SubscCats: []string{"abcd", "bcde"},
		Upsis:     []string{"abcde", "bcdef"},
	}

	{
		// Patch data (Use RESTful PATCH)
		var policyDataUesUeIdUePolicySetPatchParamOpts Nudr_DataRepository.PolicyDataUesUeIdUePolicySetPatchParamOpts
		policyDataUesUeIdUePolicySetPatchParamOpts.UePolicySet = optional.NewInterface(patchData)
		res, err := client.DefaultApi.PolicyDataUesUeIdUePolicySetPatch(context.TODO(), ueId, &policyDataUesUeIdUePolicySetPatchParamOpts)
		if err != nil {
			log.Panic(err)
		}

		if status := res.StatusCode; status != http.StatusNoContent {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusNoContent)
		}
	}

	{
		// Check patch data (Use RESTful GET)
		uePolicySet, res, err := client.DefaultApi.PolicyDataUesUeIdUePolicySetGet(context.TODO(), ueId)
		if err != nil {
			log.Panic(err)
		}

		if status := res.StatusCode; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		if cmp.Equal(patchData, uePolicySet, Opt) != true {
			t.Errorf("handler returned unexpected body: got %v want %v",
				uePolicySet, patchData)
		}
	}

	// Clean test data
	collection.DeleteOne(context.TODO(), bson.M{"ueId": "imsi-0123456789"})

	// TEST END
}

// PolicyDataUesUeIdUePolicySetPut -
func TestPolicyDataUesUeIdUePolicySetPut(t *testing.T) {
	runTestServer(t)

	connectMongoDB(t)

	// Drop old data
	collection := Client.Database("free5gc").Collection("policyData.ues.uePolicySet")
	collection.DeleteOne(context.TODO(), bson.M{"ueId": "imsi-0123456789"})

	// Set client and set url
	client := setTestClient(t)

	// Set test data
	ueId := "imsi-0123456789"
	testData := models.UePolicySet{
		SubscCats: []string{"abc", "bcd"},
		Upsis:     []string{"abcd", "bcde"},
	}

	{
		// Insert test data (Use RESTful PUT)
		var policyDataUesUeIdUePolicySetPutParamOpts Nudr_DataRepository.PolicyDataUesUeIdUePolicySetPutParamOpts
		policyDataUesUeIdUePolicySetPutParamOpts.UePolicySet = optional.NewInterface(testData)
		res, err := client.DefaultApi.PolicyDataUesUeIdUePolicySetPut(context.TODO(), ueId, &policyDataUesUeIdUePolicySetPutParamOpts)
		if err != nil {
			log.Panic(err)
		}

		if status := res.StatusCode; status != http.StatusNoContent && status != http.StatusCreated {
			t.Errorf("handler returned wrong status code: got %v want %v or %v",
				status, http.StatusNoContent, http.StatusCreated)
		}
	}

	{
		// Check test data (Use RESTful GET)
		uePolicySet, res, err := client.DefaultApi.PolicyDataUesUeIdUePolicySetGet(context.TODO(), ueId)
		if err != nil {
			log.Panic(err)
		}

		if status := res.StatusCode; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		if cmp.Equal(testData, uePolicySet, Opt) != true {
			t.Errorf("handler returned unexpected body: got %v want %v",
				uePolicySet, testData)
		}
	}

	// Clean test data
	collection.DeleteOne(context.TODO(), bson.M{"ueId": "imsi-0123456789"})

	// TEST END
}
