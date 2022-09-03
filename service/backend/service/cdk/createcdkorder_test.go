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
			t.Log("start req", util.MustToString(i))
			req := &types.CreateCdkOrderReq{
				PersonId: util.MustToString(i),
				CdkId:    util.MustToString(i),
				Number:   util.MustToInt64(i),
			}

			resp, err := srv.CreateCdkOrderSvc(req)
			if err != nil {
				t.Log(err)
			} else {
				t.Log(resp)
			}
			t.Log("end req", util.MustToString(i))
			wg.Done()
		}(i)

	}
	wg.Wait()

}
