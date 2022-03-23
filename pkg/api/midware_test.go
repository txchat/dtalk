package api

import "testing"

func Test_VerifyAddress(t *testing.T) {
	//sig := "5a3tPpKtUwVzOXAmEFuMRNdHJqJnkjbWUPKVYJDLJRV9+AksajpvT9UUSeNFVVL1W1F8EUDQt01bp11jtV8gbwA=#1610336258730*6XofpoSc#0375610055c57e011a0a51457e0ce451849a4ca588b0ff0beb0ba5d929ca2dd82b"
	sig := "sjL53g5zbfwREjuBDfTpyaQRfULgip2Ax0Es3tVEIuBCxvWryXm7EVRv/jEmmi6ZlMZZbEXeKlBxrFt41OCPWAE=#1626170987818458*2syd7kfaiw#036801a786cba366d5f62283173681dd5801740d3ef6fe1d4383680f3c0f8e7d4f"

	b, err := VerifyAddress(sig)
	if err != nil {
		t.Errorf("verify failed: %v", err)
		return
	}
	t.Logf("verify:%v", b)
}
