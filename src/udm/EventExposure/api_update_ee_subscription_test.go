/*
 * Nudm_EE
 *
 * Nudm Event Exposure Service
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package EventExposure_test

import (
	"testing"
)

// UpdateEeSubscription - Patch
func TestUpdateEeSubscription(t *testing.T) {
	/*go udm_handler.Handle()
	go func() { // udm server
		router := gin.Default()
		Nudm_EE_Server.AddService(router)

		udmLogPath := path_util.Gofree5gcPath("free5gc/udmsslkey.log")
		udmPemPath := path_util.Gofree5gcPath("free5gc/support/TLS/udm.pem")
		udmKeyPath := path_util.Gofree5gcPath("free5gc/support/TLS/udm.key")
		server, err := http2_util.NewServer(":29503", udmLogPath, router)
		if err == nil && server != nil {
			logger.InitLog.Infoln(server.ListenAndServeTLS(udmPemPath, udmKeyPath))
			assert.True(t, err == nil)
		}
	}()

	go func() { // fake udr server
		router := gin.Default()

		router.PATCH("/nudr-dr/v1/:ueIdentity/ee-subscriptions/:subscriptionId", func(c *gin.Context) {
			ueIdentity := c.Param("ueIdentity")
			fmt.Println("==========CreateEeSubscription - Subscribe==========")
			fmt.Println("ueIdentity: ", ueIdentity)
			var patchItems []models.PatchItem
			if err := c.ShouldBindJSON(&patchItems); err != nil {
				fmt.Println("fake udm server error: ", err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{})
				return
			}
			fmt.Println("patchItems: ", patchItems)
			c.JSON(http.StatusCreated, gin.H{})
		})

		udrLogPath := path_util.Gofree5gcPath("free5gc/udrsslkey.log")
		udrPemPath := path_util.Gofree5gcPath("free5gc/support/TLS/udr.pem")
		udrKeyPath := path_util.Gofree5gcPath("free5gc/support/TLS/udr.key")

		server, err := http2_util.NewServer(":29504", udrLogPath, router)
		if err == nil && server != nil {
			logger.InitLog.Infoln(server.ListenAndServeTLS(udrPemPath, udrKeyPath))
			assert.True(t, err == nil)
		}
	}()

	udm_context.Init()
	cfg := Nudm_EE_Client.NewConfiguration()
	cfg.SetBasePath("https://localhost:29503")
	clientAPI := Nudm_EE_Client.NewAPIClient(cfg)

	var patchItem []models.PatchItem
	subscriptionId := "Test001"
	ueIdentity := "SDM1234"
	resp, err := clientAPI.UpdateEESubscriptionApi.UpdateEeSubscription(context.Background(), ueIdentity, subscriptionId, patchItem)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("resp: ", resp)
	}
	*/
}
