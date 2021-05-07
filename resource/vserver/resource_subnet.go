package vserver

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform/client"
	"github.com/vngcloud/terraform/client/vserver"
)

func ResourceSubnet() *schema.Resource {
	return &schema.Resource{
		Create: resourceSubnetCreate,
		Read:   resourceSubnetRead,
		//Update: resourceSubnetUpdate,
		Delete: resourceSubnetDelete,

		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cidr": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"network_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}
func resourceSubnetStateRefreshFunc(cli *client.Client, subnetID string, projectID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		resp, _, err := cli.VserverClient.SubnetRestControllerApi.GetSubnetUsingGET(context.TODO(), projectID, subnetID)
		if err != nil {
			return nil, "", fmt.Errorf("Error on network State Refresh: %s", err)
		}
		if !resp.Success {
			return nil, "", fmt.Errorf("Error describing instance: %s", resp.ErrorMsg)
		}
		subnet := resp.Subnets[0]
		return subnet, subnet.Status, nil
	}
}
func resourceSubnetCreate(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	subnet := vserver.CreateSubnetRequest{
		Name:      d.Get("name").(string),
		Cidr:      d.Get("cidr").(string),
		NetworkId: d.Get("network_id").(string),
	}
	cli := m.(*client.Client)
	resp, _, err := cli.VserverClient.SubnetRestControllerApi.CreateSubnetUsingPOST(context.TODO(), subnet, projectID)
	if err != nil {
		return err
	}
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	if !resp.Success {
		err := fmt.Errorf(resp.ErrorMsg)
		return err
	}
	stateConf := &resource.StateChangeConf{
		Pending:    subnetCreating,
		Target:     subnetCreated,
		Refresh:    resourceSubnetStateRefreshFunc(cli, resp.Subnets[0].Uuid, projectID),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      10 * time.Second,
		MinTimeout: 3 * time.Second,
	}
	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error waiting for instance (%s) to be created: %s", resp.Subnets[0].Uuid, err)
	}
	d.SetId(resp.Subnets[0].Uuid)
	return resourceSubnetRead(d, m)
}

func resourceSubnetRead(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	subnetID := d.Id()
	cli := m.(*client.Client)
	resp, _, err := cli.VserverClient.SubnetRestControllerApi.GetSubnetUsingGET(context.TODO(), projectID, subnetID)
	if err != nil {
		return err
	}
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	if !resp.Success {
		err := fmt.Errorf(resp.ErrorMsg)
		return err
	}
	if len(resp.Subnets) == 0 {
		d.SetId("")
	}
	return nil
}

func resourceSubnetUpdate(d *schema.ResourceData, m interface{}) error {
	// if d.HasChange("size") || d.HasChange("volume_type_id") {
	// 	projectID := d.Get("project_id").(string)
	// 	resizeVolume := vSubnet.ResizeVolumeRequest{
	// 		NewSize:         int32(d.Get("size").(int)),
	// 		VolumeId:        d.Id(),
	// 		NewVolumeTypeId: d.Get("volume_type_id").(string),
	// 	}
	// 	cli := m.(*client.Client)
	// 	resp, _, err := cli.VSubnetClient.VolumeRestControllerApi.ResizeVolumeUsingPUT(context.TODO(), projectID, resizeVolume)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	if !resp.Success {
	// 		err := fmt.Errorf(resp.ErrorMsg)
	// 		return err
	// 	}
	// 	respJSON, _ := json.Marshal(resp)
	// 	log.Printf("-------------------------------------\n")
	// 	log.Printf("%s\n", string(respJSON))
	// 	log.Printf("-------------------------------------\n")
	// 	return nil
	// }
	return resourceSubnetRead(d, m)

}

func resourceSubnetDelete(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	deleteSubnet := vserver.DeleteSubnetRequest{
		SubnetId:  d.Id(),
		NetworkId: d.Get("network_id").(string),
	}
	cli := m.(*client.Client)
	resp, _, err := cli.VserverClient.SubnetRestControllerApi.DeleteSubnetUsingDELETE(context.TODO(), deleteSubnet, projectID)
	if err != nil {
		return err
	}
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	if !resp.Success {
		err := fmt.Errorf(resp.ErrorMsg)
		return err
	}
	return resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		resp, _, err := cli.VserverClient.SubnetRestControllerApi.GetSubnetUsingGET(context.TODO(), projectID, d.Id())
		if err != nil {
			return resource.NonRetryableError(fmt.Errorf("Error describing instance: %s", err))
		}
		if !resp.Success {
			return resource.NonRetryableError(fmt.Errorf("Error describing instance: %s", resp.ErrorMsg))
		}
		if len(resp.Subnets) == 0 {
			d.SetId("")
			return nil
		}
		return resource.RetryableError(fmt.Errorf("Expected instance to be created but was in state %s", resp.Subnets[0].Status))
	})
}
