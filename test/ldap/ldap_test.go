package ldap_test

import (
	"log"
	"testing"

	"github.com/go-ldap/ldap/v3"
)

func TestLdap(t *testing.T) {
	// 连接 LDAP 服务器
	l, err := ldap.DialURL("ldap://192.168.1.105:389")
	if err != nil {
		log.Fatal("连接失败:", err)
	}
	defer l.Close()

	// 管理员绑定
	err = l.Bind("cn=admin,dc=qqlx,dc=net", "123456")
	if err != nil {
		log.Fatal("管理员绑定失败:", err)
	}

	// 创建用户
	userDN := "uid=jenkins_user,ou=people,dc=qqlx,dc=net"
	userReq := ldap.NewAddRequest(userDN, nil)
	userReq.Attribute("objectClass", []string{"inetOrgPerson", "organizationalPerson", "person", "top"})
	userReq.Attribute("uid", []string{"jenkins_user"})
	userReq.Attribute("cn", []string{"Jenkins User"})
	userReq.Attribute("sn", []string{"User"})
	userReq.Attribute("userPassword", []string{"123456"})

	if err := l.Add(userReq); err != nil {
		log.Fatal("创建用户失败:", err)
	}

	// 创建组并添加用户到组
	groupDN := "cn=jenkins_admins,ou=groups,dc=qqlx,dc=net"
	groupReq := ldap.NewAddRequest(groupDN, nil)
	groupReq.Attribute("objectClass", []string{"groupOfNames", "top"})
	groupReq.Attribute("cn", []string{"jenkins_admins"})
	groupReq.Attribute("member", []string{userDN})

	if err := l.Add(groupReq); err != nil {
		log.Fatal("创建组失败:", err)
	}

	log.Println("用户和组创建成功！")
}
