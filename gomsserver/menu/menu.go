package menu


// Check 检查用户菜单权限
func Check(path, method string, userType uint8) bool {
	//fmt.Printf("path:%s,method:%s,userType:%d", path, method, userType)
	if path == "" {
		return false
	}

	for _, menuItem := range items {
		// 没匹配上路径
		if menuItem.Path != path {
			continue
		}

		// 不需要验证请求方法和用户权限
		if len(menuItem.Methods) == 0 && len(menuItem.Types) == 0 {
			return true
		}

		// 验证请求方法
		if len(menuItem.Methods) != 0 {
			for _, menuMethod := range menuItem.Methods {
				// 没匹配上方法
				if menuMethod.Method != method {
					continue
				}

				// 不需要验证用户权限
				if len(menuMethod.Types) == 0 {
					return true
				}
				// 验证用户权限
				return checkType(menuMethod.Types, userType)
			}
			return false
		}

		// 验证用户权限
		if len(menuItem.Types) != 0 {
			return checkType(menuItem.Types, userType)
		}

		return false
	}

	return false
}


func checkType(userTypes []uint8, userType uint8) bool {
	for _, ut := range userTypes {
		if ut == userType {
			return true
		}
	}
	return false
}
