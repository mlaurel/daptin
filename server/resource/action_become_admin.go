package resource

import (
  "github.com/pkg/errors"
)

type BecomeAdminActionPerformer struct {
  responseAttrs map[string]interface{}
  cruds         map[string]*DbResource
}

func (d *BecomeAdminActionPerformer) Name() string {
  return "__become_admin"
}

func (d *BecomeAdminActionPerformer) DoAction(request ActionRequest, inFieldMap map[string]interface{}) ([]ActionResponse, []error) {

  if !d.cruds["world"].CanBecomeAdmin() {
    return nil, []error{errors.New("Unauthorized")}
  }
  user := inFieldMap["user"].(map[string]interface{})

  responseAttrs := make(map[string]interface{})

  if d.cruds["world"].BecomeAdmin(user["id"].(int64)) {
    responseAttrs["location"] = "/"
    responseAttrs["window"] = "self"
    responseAttrs["delay"] = 0
  }

  actionResponse := NewActionResponse("client.redirect", responseAttrs)

  return []ActionResponse{actionResponse}, nil
}

func NewBecomeAdminPerformer(initConfig *CmsConfig, cruds  map[string]*DbResource) (ActionPerformerInterface, error) {

  handler := BecomeAdminActionPerformer{
    cruds: cruds,
  }

  return &handler, nil

}