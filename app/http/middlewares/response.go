package middlewares

//type Response struct {
//	Message string `json:"message"`
//	Code    int64  `json:"code"`
//	Data    any    `json:"data"`
//}
//
//func responseOk(data any) Response {
//	return Response{
//		Message: "ok",
//		Code:    http.StatusOK,
//		Data:    data,
//	}
//}
//
//func responseError(message string) Response {
//	return Response{
//		Message: message,
//		Code:    http.StatusInternalServerError,
//		Data:    nil,
//	}
//}
//
//// ExceptionHandlerFunc 异常处理
//type ExceptionHandlerFunc func(ctx *gin.Context) (any, error)
//
//// Wrapper 中间件
//func Wrapper(handlerFunc ExceptionHandlerFunc) gin.HandlerFunc {
//	return func(context *gin.Context) {
//		data, err := handlerFunc(context)
//		if err != nil {
//			context.JSON(http.StatusOK, responseError(err.Error()))
//			return
//		}
//
//		context.JSON(http.StatusOK, responseOk(data))
//	}
//}
