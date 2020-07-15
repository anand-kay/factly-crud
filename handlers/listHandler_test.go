package handlers

import "testing"

func TestPaginate(t *testing.T) {
	var userInfos []userInfo
	var userInfo userInfo

	userInfo.ID = "1"
	userInfo.UserName = "one"
	userInfo.Email = "one@gmail.com"
	userInfos = append(userInfos, userInfo)

	userInfo.ID = "2"
	userInfo.UserName = "two"
	userInfo.Email = "two@gmail.com"
	userInfos = append(userInfos, userInfo)

	pagesA := paginate(userInfos)
	if len(pagesA[1]) != 2 {
		t.Error("Error in pagination")
	}

	userInfo.ID = "3"
	userInfo.UserName = "two"
	userInfo.Email = "two@gmail.com"
	userInfos = append(userInfos, userInfo)

	userInfo.ID = "4"
	userInfo.UserName = "two"
	userInfo.Email = "two@gmail.com"
	userInfos = append(userInfos, userInfo)

	userInfo.ID = "5"
	userInfo.UserName = "two"
	userInfo.Email = "two@gmail.com"
	userInfos = append(userInfos, userInfo)

	userInfo.ID = "6"
	userInfo.UserName = "two"
	userInfo.Email = "two@gmail.com"
	userInfos = append(userInfos, userInfo)

	userInfo.ID = "7"
	userInfo.UserName = "two"
	userInfo.Email = "two@gmail.com"
	userInfos = append(userInfos, userInfo)

	userInfo.ID = "8"
	userInfo.UserName = "two"
	userInfo.Email = "two@gmail.com"
	userInfos = append(userInfos, userInfo)

	userInfo.ID = "9"
	userInfo.UserName = "two"
	userInfo.Email = "two@gmail.com"
	userInfos = append(userInfos, userInfo)

	userInfo.ID = "10"
	userInfo.UserName = "two"
	userInfo.Email = "two@gmail.com"
	userInfos = append(userInfos, userInfo)

	pagesB := paginate(userInfos)
	if len(pagesB[1]) != 10 {
		t.Error("Error in pagination")
	}

	userInfo.ID = "11"
	userInfo.UserName = "two"
	userInfo.Email = "two@gmail.com"
	userInfos = append(userInfos, userInfo)

	userInfo.ID = "12"
	userInfo.UserName = "two"
	userInfo.Email = "two@gmail.com"
	userInfos = append(userInfos, userInfo)

	userInfo.ID = "13"
	userInfo.UserName = "two"
	userInfo.Email = "two@gmail.com"
	userInfos = append(userInfos, userInfo)

	pagesC := paginate(userInfos)
	if len(pagesC[2]) != 3 {
		t.Error("Error in pagination")
	}
}
