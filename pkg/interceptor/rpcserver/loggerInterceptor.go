package rpcserver

/**
* @Description rpc service logger interceptor
* @Author Mikael
* @Date 2021/1/9 13:35
* @Version 1.0
**/

//func LoggerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
//
//	resp, err = handler(ctx, req)
//	if err != nil {
//		causeErr := errors.Cause(err)                // err类型
//		if e, ok := causeErr.(*xerr.CodeError); ok { //自定义错误类型
//			logx.WithContext(ctx).Errorf("【RPC-SRV-ERR】 %+v", err)
//
//			//转成grpc err
//			err = status.Error(codes.Code(e.GetErrCode()), e.GetErrMsg())
//		} else {
//			logx.WithContext(ctx).Errorf("【RPC-SRV-ERR】 %+v", err)
//		}
//
//	}
//
//	return resp, err
//}
