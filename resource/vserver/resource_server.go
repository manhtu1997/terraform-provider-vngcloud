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

func ResourceServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerCreate,
		Read:   resourceServerRead,
		Update: resourceServerUpdate,
		Delete: resourceServerDelete,

		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"attach_floating": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"data_disk_encryption_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"data_disk_size": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"data_disk_type_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"data_volume_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"encryption_volume": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"flavor_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"image_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"is_poc": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"network_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"os_licence": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"period": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"root_disk_encryption_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"root_disk_size": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"root_disk_type_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"security_group": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"source_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssh_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
func resourceServerStateRefreshFunc(cli *client.Client, serverID string, projectID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		resp, _, err := cli.VserverClient.ServerRestControllerApi.GetServerUsingGET(context.TODO(), projectID, serverID)
		if err != nil {
			return nil, "", fmt.Errorf("Error on network State Refresh: %s", err)
		}
		if !resp.Success {
			return nil, "", fmt.Errorf("Error describing instance: %s", resp.ErrorMsg)
		}
		server := resp.Servers[0]
		return server, server.Status, nil
	}
}
func resourceServerCreate(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	securityGroupInterface := d.Get("security_group").([]interface{})
	var securityGroup []string
	for _, s := range securityGroupInterface {
		securityGroup = append(securityGroup, s.(string))
	}
	server := vserver.CreateServerRequest{
		AttachFloating:         d.Get("attach_floating").(bool),
		DataDiskEncryptionType: d.Get("data_disk_encryption_type").(string),
		DataDiskSize:           int32(d.Get("data_disk_size").(int)),
		DataDiskTypeId:         d.Get("data_disk_type_id").(string),
		DataVolumeName:         d.Get("data_volume_name").(string),
		EncryptionVolume:       d.Get("encryption_volume").(bool),
		FlavorId:               d.Get("flavor_id").(string),
		ImageId:                d.Get("image_id").(string),
		IsPoc:                  d.Get("is_poc").(bool),
		Name:                   d.Get("name").(string),
		NetworkId:              d.Get("flavor_id").(string),
		OsLicence:              d.Get("os_licence").(bool),
		Period:                 int32(d.Get("period").(int)),
		RootDiskEncryptionType: d.Get("root_disk_encryption_type").(string),
		RootDiskSize:           int32(d.Get("root_disk_size").(int)),
		RootDiskTypeId:         d.Get("root_disk_type_id").(string),
		SecurityGroup:          securityGroup,
		SourceType:             d.Get("source_type").(string),
		SshKeyId:               d.Get("ssh_key").(string),
		SubnetId:               d.Get("subnet_id").(string),
	}
	cli := m.(*client.Client)
	resp, _, err := cli.VserverClient.ServerRestControllerApi.CreateServerUsingPOST(context.TODO(), server, projectID)
	if err != nil {
		return err
	}
	if !resp.Success {
		err := fmt.Errorf(resp.ErrorMsg)
		return err
	}
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	stateConf := &resource.StateChangeConf{
		Pending:    serverCreating,
		Target:     serverCreated,
		Refresh:    resourceServerStateRefreshFunc(cli, resp.Servers[0].Uuid, projectID),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      10 * time.Second,
		MinTimeout: 3 * time.Second,
	}
	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error waiting for instance (%s) to be created: %s", resp.Servers[0].Uuid, err)
	}
	d.SetId(resp.Servers[0].Uuid)
	return resourceServerRead(d, m)
}

func resourceServerRead(d *schema.ResourceData, m interface{}) error {

	projectID := d.Get("project_id").(string)
	serverID := d.Id()
	cli := m.(*client.Client)
	resp, _, err := cli.VserverClient.ServerRestControllerApi.GetServerUsingGET(context.TODO(), projectID, serverID)
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
	if len(resp.Servers) == 0 {
		d.SetId("")
	}
	d.Set("status", resp.Servers[0].Status)
	return nil
}

func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {

	return resourceServerRead(d, m)

}

func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	deleteServer := vserver.DeleteServerRequest{
		ServerId:    d.Id(),
		ForceDelete: true,
	}
	cli := m.(*client.Client)
	resp, _, err := cli.VserverClient.ServerRestControllerApi.DeleteServerInTrashUsingDELETE(context.TODO(), deleteServer, projectID)
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
		resp, _, err := cli.VserverClient.ServerRestControllerApi.GetServerUsingGET(context.TODO(), projectID, d.Id())
		if err != nil {
			return resource.NonRetryableError(fmt.Errorf("Error describing instance: %s", err))
		}
		if !resp.Success {
			return resource.NonRetryableError(fmt.Errorf("Error describing instance: %s", resp.ErrorMsg))
		}
		if len(resp.Servers) == 0 {
			d.SetId("")
			return nil
		}
		return resource.RetryableError(fmt.Errorf("Expected instance to be created but was in state %s", resp.Servers[0].Status))
	})
}
