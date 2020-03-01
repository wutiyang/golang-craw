package rpcdemo

import "errors"

// 调用方式：Service.Method

type DemoService struct {

}

type Args struct {
	A,B int
}

// args及result参数是固定的，返回error 为rpc框架要求
func (DemoService) Div (args Args, result *float64) error {
	if args.B == 0 {
		return errors.New("division by zeror")
	}

	*result = float64(args.A) / float64(args.B)

	return nil
}
