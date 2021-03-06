package mcir

import (
	"fmt"
	"net/http"

	"github.com/cloud-barista/poc-cicd-tumblebug/src/core/common"
	"github.com/cloud-barista/poc-cicd-tumblebug/src/core/mcir"
	"github.com/labstack/echo/v4"
)

// RestPostSecurityGroup godoc
// @Summary Create Security Group
// @Description Create Security Group
// @Tags [MCIR] Security group management
// @Accept  json
// @Produce  json
// @Param nsId path string true "Namespace ID"
// @Param securityGroupReq body mcir.TbSecurityGroupReq true "Details for an securityGroup object"
// @Success 200 {object} mcir.TbSecurityGroupInfo
// @Failure 404 {object} common.SimpleMsg
// @Failure 500 {object} common.SimpleMsg
// @Router /ns/{nsId}/resources/securityGroup [post]
func RestPostSecurityGroup(c echo.Context) error {

	nsId := c.Param("nsId")

	u := &mcir.TbSecurityGroupReq{}
	if err := c.Bind(u); err != nil {
		return err
	}

	fmt.Println("[POST SecurityGroup")
	//fmt.Println("[Creating SecurityGroup]")
	//content, responseCode, _, err := CreateSecurityGroup(nsId, u)
	content, err := mcir.CreateSecurityGroup(nsId, u)
	if err != nil {
		common.CBLog.Error(err)
		/*
			mapA := map[string]string{
				"message": "Failed to create a SecurityGroup"}
		*/
		mapA := map[string]string{"message": err.Error()}
		return c.JSON(http.StatusInternalServerError, &mapA)
	}
	return c.JSON(http.StatusCreated, content)
}

/* function RestPutSecurityGroup not yet implemented
// RestPutSecurityGroup godoc
// @Summary Update Security Group
// @Description Update Security Group
// @Tags [MCIR] Security group management
// @Accept  json
// @Produce  json
// @Param securityGroupInfo body mcir.TbSecurityGroupInfo true "Details for an securityGroup object"
// @Success 200 {object} mcir.TbSecurityGroupInfo
// @Failure 404 {object} common.SimpleMsg
// @Failure 500 {object} common.SimpleMsg
// @Router /ns/{nsId}/resources/securityGroup/{securityGroupId} [put]
*/
func RestPutSecurityGroup(c echo.Context) error {
	//nsId := c.Param("nsId")

	return nil
}

// RestGetSecurityGroup godoc
// @Summary Get Security Group
// @Description Get Security Group
// @Tags [MCIR] Security group management
// @Accept  json
// @Produce  json
// @Param nsId path string true "Namespace ID"
// @Param securityGroupId path string true "Security Group ID"
// @Success 200 {object} mcir.TbSecurityGroupInfo
// @Failure 404 {object} common.SimpleMsg
// @Failure 500 {object} common.SimpleMsg
// @Router /ns/{nsId}/resources/securityGroup/{securityGroupId} [get]
func RestGetSecurityGroup(c echo.Context) error {
	// This is a dummy function for Swagger.
	return nil
}

// Response structure for RestGetAllSecurityGroup
type RestGetAllSecurityGroupResponse struct {
	SecurityGroup []mcir.TbSecurityGroupInfo `json:"securityGroup"`
}

// RestGetAllSecurityGroup godoc
// @Summary List all Security Groups or Security Groups' ID
// @Description List all Security Groups or Security Groups' ID
// @Tags [MCIR] Security group management
// @Accept  json
// @Produce  json
// @Param nsId path string true "Namespace ID"
// @Param option query string false "Option" Enums(id)
// @Success 200 {object} JSONResult{[DEFAULT]=RestGetAllSecurityGroupResponse,[ID]=common.IdList} "Different return structures by the given option param"
// @Failure 404 {object} common.SimpleMsg
// @Failure 500 {object} common.SimpleMsg
// @Router /ns/{nsId}/resources/securityGroup [get]
func RestGetAllSecurityGroup(c echo.Context) error {
	// This is a dummy function for Swagger.
	return nil
}

// RestDelSecurityGroup godoc
// @Summary Delete Security Group
// @Description Delete Security Group
// @Tags [MCIR] Security group management
// @Accept  json
// @Produce  json
// @Param nsId path string true "Namespace ID"
// @Param securityGroupId path string true "Security Group ID"
// @Success 200 {object} common.SimpleMsg
// @Failure 404 {object} common.SimpleMsg
// @Router /ns/{nsId}/resources/securityGroup/{securityGroupId} [delete]
func RestDelSecurityGroup(c echo.Context) error {
	// This is a dummy function for Swagger.
	return nil
}

// RestDelAllSecurityGroup godoc
// @Summary Delete all Security Groups
// @Description Delete all Security Groups
// @Tags [MCIR] Security group management
// @Accept  json
// @Produce  json
// @Param nsId path string true "Namespace ID"
// @Success 200 {object} common.SimpleMsg
// @Failure 404 {object} common.SimpleMsg
// @Router /ns/{nsId}/resources/securityGroup [delete]
func RestDelAllSecurityGroup(c echo.Context) error {
	// This is a dummy function for Swagger.
	return nil
}
