package cdk

import (
	"sync"
	"testing"

	"github.com/txchat/dtalk/pkg/util"
	"github.com/txchat/dtalk/service/backend/model/types"
)

func TestCreateCdkOrder(t *testing.T) {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			t.Log("start req", util.ToString(i))
			req := &types.CreateCdkOrderReq{
				PersonId: util.ToString(i),
				CdkId:    util.ToString(i),
				Number:   util.ToInt64(i),
			}

			resp, err := srv.CreateCdkOrderSvc(req)
			if err != nil {
				t.Log(err)
			} else {
				t.Log(resp)
			}
			t.Log("end req", util.ToString(i))
			wg.Done()
		}(i)

	}
	wg.Wait()

}
