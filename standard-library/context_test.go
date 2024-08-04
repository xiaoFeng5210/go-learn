package standardLibrary

import ("testing" "context" "time")

func TestContext(t *testing.T) {
	ctx := context.Background() // 起点
  timeoutCtx, cancel := context.WithTimeout(ctx, time.Second)
}

func SomeBusiness(ctx: context.Context) {

}

