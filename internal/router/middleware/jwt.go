package midware

// JwtAuthMiddleware 基于 JWT 的认证中间件
//func JwtAuthMiddleware() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		// 从配置文件中读取 TokenHeader, 这里为 "Authorization"
//		tokenHeader := config.Get().JwtConfig.TokenHeader
//
//		// 客户端携带 Token 有三种方式 1.放在请求头 2.放在请求体 3.放在URI
//		// 这里假设 Token 放在 Header 的 Authorization 中，并使用 Bearer 开头
//		// 这里的具体实现方式要依据你的实际业务情况决定
//		authHeader := c.Request.Header.Get(tokenHeader)
//
//		if authHeader == "" {
//			resp.Unauthorized(c, "请求头中auth为空")
//			c.Abort()
//			return
//		}
//		fmt.Println("获取到的Authorization:", authHeader)
//
//		// 按空格分割
//		// 如果 s 中没有 str 子串，则将整个 s 作为 []string 的第一个元素返回。
//		// 参数 n 表示最多切分出几个子串，超出的部分将不再切分，最后一个 n 包含了所有剩下的不切分。
//		// 如果 n 为 0，则返回 nil；如果 n 小于 0，则不限制切分个数，全部切分
//		parts := strings.SplitN(authHeader, " ", 2)
//
//		// 从配置文件读取 tokenHead 这里为 "Bearer"
//		tokenHead := config.Get().JwtConfig.TokenHead
//
//		// 验证请求头的格式
//		if !(len(parts) == 2 && parts[0] == tokenHead) {
//			resp.Unauthorized(c, "请求头中auth格式有误")
//			c.Abort()
//			return
//		}
//		// parts[1] 是获取到的 tokenString，我们使用之前定义好的解析 JWT 的函数来解析它
//		mc, err := util.ParseJwtToken(parts[1])
//		if err != nil {
//			resp.Unauthorized(c, "无效的token")
//			c.Abort()
//			return
//		}
//		user, err := service.UserService.LoadUserByUsername(mc.Username)
//		if err != nil {
//			resp.Unauthorized(c, "用户不存在")
//			c.Abort()
//			return
//		}
//		// 将当前请求的 username 信息保存到请求的上下文 c 上
//		c.Set("user", user)
//		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
//	}
//}
