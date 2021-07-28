package jwt

import "testing"

/**
 * @Author: mf
 * @email: 18539271635@163.com
 * @Date: 2021/7/28 4:37 下午
 * @Desc:
 */

func TestGenerateToken(t *testing.T) {
	token, err := GenerateToken(10000, "tom")
	if err != nil {
		t.Fatalf("GenerateToken Fail. err=[%s]", err.Error())
	}
	t.Logf("GenerateToken SUCCESS.\ntoken=[%s]", token)
}

func TestParseToken(t *testing.T) {
	//token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMDAwMCwidXNlcm5hbWUiOiJ0b20iLCJleHAiOjE2Mjc0Njg4NTUsImlhdCI6MTYyNzQ2MTY1NSwiaXNzIjoiTUYifQ.vKRY8oLor0IPyISnFOOflfQf3n8_LHRa4xrR37ue1hk"
	//claims, ok := ParseToken(token)
	//if !ok {
	//	t.Fatalf("ParseToken Fail.")
	//}
	//if claims.UserId != 10000 && claims.Username != "tom" {
	//	t.Fatalf("claims=[%v]", *claims)
	//}
	//t.Logf("claims=[%+v]", *claims)

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMDAwMCwidXNlcm5hbWUiOiJ0b20iLCJleHAiOjE2Mjc0NjI1MzksImlhdCI6MTYyNzQ2MjUwOSwiaXNzIjoiTUYifQ.wskVxsSbCfq5TNZu2z3r8_AkJW-DNrykI6wBKcjjTSw"
	claims, ok := ParseToken(token)
	if !ok {
		t.Fatalf("ParseToken Fail.")
	}
	if claims.UserId != 10000 && claims.Username != "tom" {
		t.Fatalf("claims=[%v]", *claims)
	}
	t.Logf("claims=[%+v]", *claims)
}